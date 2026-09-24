// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfazureclient "github.com/hashicorp/terraform-provider-azuread/xpprovider"
	"sigs.k8s.io/controller-runtime/pkg/client"

	namespacedv1beta1 "github.com/upbound/provider-azuread/v2/apis/namespaced/v1beta1"
)

const (
	// Terraform Provider configuration key for the Azure CLI authenticator,
	// which defaults to true and shells out to the `az` binary.
	keyUseCLI = "use_cli"

	// Placeholder identities used when configuring the Terraform provider for
	// offline diffs. They are only ever observed locally: nothing derived from
	// them leaves the process, because the token they end up in is minted by
	// offlineTokenClient rather than by Azure.
	offlineTenantID     = "00000000-0000-0000-0000-000000000000"
	offlineClientID     = "00000000-0000-0000-0000-000000000001"
	offlineObjectID     = "00000000-0000-0000-0000-000000000002"
	offlineClientSecret = "offline-diff-placeholder"

	// errRegisterRequestMiddleware is returned when the egress guard cannot be
	// installed. Offline diffs rely on it, so failing to install it is fatal
	// rather than a degradation.
	errRegisterRequestMiddleware = "cannot register the offline request middleware: provider meta is not an AzureAD client"

	// fmtErrOfflineRequestBlocked reports a Microsoft Graph call that the
	// offline diff server refused to send. The request's query string is
	// deliberately left out: it carries $filter values taken from the
	// resource's own fields.
	fmtErrOfflineRequestBlocked = "the offline diff server blocked an outbound Microsoft Graph request (%s %s): computing this diff requires calling Azure, which is not possible offline. On Group, Application and Unit resources this is usually caused by spec.forProvider.preventDuplicateNames being true"

	// offlineTokenLifetime is what the offline token claims for expires_in.
	// The value is immaterial: when the cached token lapses, the Azure SDK
	// simply asks offlineTokenClient for another one.
	offlineTokenLifetime = "3600"
)

// EnableOfflineAuthentication replaces the HTTP clients that the Azure SDK uses
// to obtain access tokens with an in-process stub, so that configuring the
// AzureAD Terraform provider neither requires credentials nor reaches Azure.
//
// This is for the offline diff server, which computes the difference between a
// desired and an actual managed resource state and never reconciles the
// external resource. It must not be called in a provider that reconciles, as it
// disables authentication for the whole process: both auth.Client and
// auth.MetadataClient are package-level variables in the Azure SDK.
//
// It is safe to call more than once; only the first call takes effect.
func EnableOfflineAuthentication() {
	offlineAuthOnce.Do(func() {
		auth.Client = offlineTokenClient{}
		auth.MetadataClient = offlineTokenClient{}
	})
}

var offlineAuthOnce sync.Once

// OfflineTerraformSetupBuilder returns a Terraform setup that configures the
// AzureAD Terraform provider without credentials and without contacting Azure.
// The returned terraform.Setup carries a fully built *clients.Client as its
// Meta, which the Terraform provider's CustomizeDiff functions require, but one
// holding an access token that no Azure API will accept.
//
// Use it in place of TerraformSetupBuilder in the diff server only. The
// ProviderConfig is still resolved, so that configuration which affects the
// diff - the cloud environment above all - is honoured; only its credentials
// are ignored.
func OfflineTerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	EnableOfflineAuthentication()
	return func(ctx context.Context, client client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, err
		}

		return configureOffline(ctx, tfProvider, pcSpec)
	}
}

// configureOffline configures the Terraform provider from pcSpec without
// credentials, and installs the egress guard on the resulting client.
func configureOffline(ctx context.Context, tfProvider *schema.Provider, pcSpec *namespacedv1beta1.ProviderConfigSpec) (terraform.Setup, error) {
	ps := terraform.Setup{
		Configuration: offlineConfiguration(pcSpec),
	}
	if err := configureNoForkAzureClient(ctx, &ps, *tfProvider); err != nil {
		return terraform.Setup{}, errors.Wrap(err, "failed to configure the no-fork Azure client")
	}
	// The stubbed HTTP clients cover token acquisition only. Microsoft Graph
	// itself is reached through a separate client, so a CustomizeDiff function
	// that calls the API - see denyOutboundRequest - would otherwise retry
	// against an unreachable endpoint until it timed out.
	if !tfazureclient.RegisterRequestMiddleware(ps.Meta, denyOutboundRequest) {
		return terraform.Setup{}, errors.New(errRegisterRequestMiddleware)
	}
	return ps, nil
}

// denyOutboundRequest refuses every Microsoft Graph request, failing the diff
// immediately instead of letting it retry against an API that will not accept
// the offline access token. The Azure SDK returns a request middleware's error
// to the caller without sending the request.
func denyOutboundRequest(req *http.Request) (*http.Request, error) {
	method, path := "unknown", "unknown"
	if req != nil {
		method = req.Method
		if req.URL != nil {
			path = req.URL.Path
		}
	}
	return nil, errors.Errorf(fmtErrOfflineRequestBlocked, method, path)
}

// offlineConfiguration returns the Terraform provider configuration to use for
// offline diffs. It never reads the ProviderConfig's credentials, which are
// assumed to be unavailable.
func offlineConfiguration(pcSpec *namespacedv1beta1.ProviderConfigSpec) map[string]any {
	cfg := map[string]any{
		keyTerraformFeatures: map[string]any{},
		keyTenantID:          offlineTenantID,
		keyClientID:          offlineClientID,
		// A non-empty client secret makes the Azure SDK pick the client
		// credentials authorizer, whose token request offlineTokenClient
		// answers locally. Were all credentials empty, the SDK would fall
		// through to the Azure CLI authorizer and shell out to `az`.
		keyClientSecret: offlineClientSecret,
		keyUseCLI:       false,
	}

	if pcSpec == nil {
		return cfg
	}
	if pcSpec.TenantID != nil && *pcSpec.TenantID != "" {
		cfg[keyTenantID] = *pcSpec.TenantID
	}
	if pcSpec.ClientID != nil && *pcSpec.ClientID != "" {
		cfg[keyClientID] = *pcSpec.ClientID
	}
	if pcSpec.Environment != nil && *pcSpec.Environment != "" {
		cfg[keyEnvironment] = *pcSpec.Environment
	}
	return cfg
}

// offlineTokenClient answers every access token request locally. It satisfies
// auth.HTTPClient, the interface behind the Azure SDK's auth.Client and
// auth.MetadataClient, and covers every authentication method the provider
// supports, because they all obtain their token over HTTP and parse the same
// response shape.
type offlineTokenClient struct{}

func (offlineTokenClient) Do(req *http.Request) (*http.Response, error) {
	token, err := offlineAccessToken()
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(map[string]any{
		"access_token": token,
		"token_type":   "Bearer",
		// The Azure SDK accepts expires_in as either a string or a number.
		"expires_in": offlineTokenLifetime,
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot marshal the offline token response")
	}

	return &http.Response{
		Status:        "200 OK",
		StatusCode:    http.StatusOK,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        http.Header{"Content-Type": []string{"application/json"}},
		Body:          io.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Request:       req,
	}, nil
}

// offlineAccessToken mints an unsigned JWT carrying the claims that configuring
// the provider depends on. The Azure SDK decodes the token's payload to read
// them and does not verify its signature. The `oid` claim matters most: were it
// absent, the provider would query Microsoft Graph to discover the object ID of
// the authenticated principal, which offline it cannot do.
func offlineAccessToken() (string, error) {
	payload, err := json.Marshal(map[string]any{
		"aud":   "https://graph.microsoft.com",
		"iss":   "https://sts.windows.net/" + offlineTenantID + "/",
		"oid":   offlineObjectID,
		"tid":   offlineTenantID,
		"appid": offlineClientID,
		"scp":   "",
	})
	if err != nil {
		return "", errors.Wrap(err, "cannot marshal the offline token claims")
	}

	enc := base64.RawURLEncoding.EncodeToString
	return enc([]byte(`{"alg":"none","typ":"JWT"}`)) + "." + enc(payload) + ".offline", nil
}
