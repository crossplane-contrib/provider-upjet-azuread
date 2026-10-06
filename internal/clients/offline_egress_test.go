// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfazureclient "github.com/hashicorp/terraform-provider-azuread/xpprovider"
)

// graphHost is the Microsoft Graph endpoint of the public cloud, which is the
// environment an offline configuration with no ProviderConfig defaults to.
const graphHost = "graph.microsoft.com"

// loginHost is the endpoint the Azure SDK requests access tokens from.
const loginHost = "login.microsoftonline.com"

// realAuthClient is the Azure SDK's own token HTTP client, captured during
// package variable initialisation - that is, before any test can call
// EnableOfflineAuthentication, whose sync.Once swap cannot be undone. Holding
// on to it lets TestConfigureEgressesWithoutOfflineAuthentication put the real
// client back regardless of what ran first.
var realAuthClient = auth.Client

// egress is the tripwire armed by TestMain, shared by every test in the
// package. Tests that read it must not run in parallel with one another.
var egress *proxyRecorder

// proxyRecorder is an HTTP proxy that records every request routed through it
// and forwards none of them.
//
// Both HTTP transports the Azure SDK builds - the one that fetches access
// tokens (go-azure-sdk/sdk/auth) and the one that talks to Microsoft Graph
// (go-azure-sdk/sdk/client) - construct their own http.Transport with
// Proxy: http.ProxyFromEnvironment rather than using http.DefaultTransport.
// Pointing HTTPS_PROXY at a recorder therefore puts a tripwire on the process's
// real network boundary: nothing is stubbed or injected, and a request cannot
// reach Azure without being recorded here first. The same tripwire works
// unchanged against the compiled provider binary.
type proxyRecorder struct {
	url string

	mu   sync.Mutex
	seen []string
}

func (p *proxyRecorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CONNECT carries its target in Host; an absolute-form proxy request
	// carries it in URL.
	target := r.Host
	if target == "" && r.URL != nil {
		target = r.URL.Host
	}
	p.mu.Lock()
	p.seen = append(p.seen, r.Method+" "+target)
	p.mu.Unlock()
	http.Error(w, "blocked by the egress tripwire", http.StatusBadGateway)
}

func (p *proxyRecorder) records() []string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return append([]string(nil), p.seen...)
}

func (p *proxyRecorder) reset() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.seen = nil
}

// reached reports whether any recorded request targeted host.
func (p *proxyRecorder) reached(host string) bool {
	for _, r := range p.records() {
		if strings.Contains(r, host) {
			return true
		}
	}
	return false
}

func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

func runTests(m *testing.M) int {
	lc := &net.ListenConfig{}
	l, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot listen for the egress tripwire: %v\n", err)
		return 1
	}
	defer l.Close() //nolint:errcheck // the process is exiting

	egress = &proxyRecorder{url: "http://" + l.Addr().String()}
	srv := &http.Server{Handler: egress, ReadHeaderTimeout: 5 * time.Second}
	go srv.Serve(l)   //nolint:errcheck // Serve always returns a non-nil error
	defer srv.Close() //nolint:errcheck // the process is exiting

	// net/http resolves the proxy environment exactly once per process, so
	// this has to happen before anything in the package makes a request.
	os.Setenv("HTTP_PROXY", egress.url)  //nolint:errcheck // must outlive any single test, so t.Setenv will not do
	os.Setenv("HTTPS_PROXY", egress.url) //nolint:errcheck // must outlive any single test, so t.Setenv will not do
	os.Setenv("NO_PROXY", "")            //nolint:errcheck // must outlive any single test, so t.Setenv will not do

	return m.Run()
}

// TestOfflineConfigureDoesNotEgress asserts that configuring the AzureAD
// Terraform provider for offline diffs makes no outbound request whatsoever.
//
// This is worth asserting because the Microsoft Graph egress guard is installed
// on the provider's Meta only *after* p.Configure returns - see
// configureOffline - so anything Configure itself calls goes out unguarded.
// clients.Client.build does contain such a call: when the access token carries
// no oid claim it queries Graph to discover the authenticated principal's
// object ID. What keeps that branch unreachable is the oid claim
// offlineAccessToken mints into the token, and that invariant - not the guard -
// is what this test locks in. See
// TestOfflineConfigureWithoutObjectIDClaimDoesNotEgress for what happens when
// the claim is taken away.
func TestOfflineConfigureDoesNotEgress(t *testing.T) {
	EnableOfflineAuthentication()
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureAD provider schema: %v", err)
	}
	if _, err := configureOffline(context.Background(), p, nil); err != nil {
		t.Fatalf("cannot configure the AzureAD provider offline: %v", err)
	}

	if got := egress.records(); len(got) > 0 {
		t.Errorf("configuring the provider offline reached the network: %v", got)
	}
}

// TestOfflineConfigureWithoutObjectIDClaimDoesNotEgress strips the oid claim
// from the offline access token to reach the Configure-time Microsoft Graph
// call, and shows that it still cannot leave the process.
//
// The call is entered - clients.Client.build logs that it is querying Graph for
// the authenticated service principal - but go-azure-sdk refuses to execute it,
// because configureNoForkAzureClient hands Configure a context built with
// context.WithoutCancel, which strips the deadline the SDK requires for a
// paged operation. So two independent things, not one, keep Configure from
// egressing: the oid claim, which keeps this branch unreachable, and the
// absent deadline, which would stop the request even if it were reached.
// Neither is the Graph guard, which configureOffline installs only afterwards.
func TestOfflineConfigureWithoutObjectIDClaimDoesNotEgress(t *testing.T) {
	EnableOfflineAuthentication()
	t.Cleanup(func() { auth.Client = offlineTokenClient{} })
	auth.Client = noObjectIDTokenClient{}
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureAD provider schema: %v", err)
	}
	_, err = configureOffline(context.Background(), p, nil)
	if err == nil {
		t.Fatal("expected configuring the provider to fail with no oid claim in the access token")
	}
	// Asserted so that this test starts failing, rather than silently passing
	// for a new reason, if the SDK ever drops the deadline requirement.
	if !strings.Contains(err.Error(), "must have a deadline attached") {
		t.Errorf("expected the discovery call to be refused for want of a deadline, got: %v", err)
	}

	if got := egress.records(); len(got) > 0 {
		t.Errorf("configuring the provider offline reached the network: %v", got)
	}
}

// TestEgressTripwireObservesUnguardedGraphCall is the positive control for the
// two tests above. They are only meaningful if the tripwire can see outbound
// Graph traffic at all, so this one configures the provider exactly as
// configureOffline does but *without* installing denyOutboundRequest, then runs
// the same azuread_group duplicate-name diff that
// TestOfflineDiffBlocksOutboundRequest uses. With nothing in the way, the
// request reaches the network boundary and the tripwire records it.
func TestEgressTripwireObservesUnguardedGraphCall(t *testing.T) {
	EnableOfflineAuthentication()
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureAD provider schema: %v", err)
	}
	ps := terraform.Setup{Configuration: offlineConfiguration(nil)}
	if err := configureNoForkAzureClient(context.Background(), &ps, *p); err != nil {
		t.Fatalf("cannot configure the AzureAD provider: %v", err)
	}

	r := p.ResourcesMap["azuread_group"]
	state := &tfsdk.InstanceState{
		ID: offlineObjectID,
		Attributes: map[string]string{
			"id":                     offlineObjectID,
			"object_id":              offlineObjectID,
			keyDisplayName:           "old-name",
			keySecurityEnabled:       valTrue,
			keyMailEnabled:           valFalse,
			keyPreventDuplicateNames: valTrue,
		},
	}
	config := &tfsdk.ResourceConfig{
		Config: map[string]any{
			keyDisplayName:           "new-name",
			keySecurityEnabled:       true,
			keyMailEnabled:           false,
			keyPreventDuplicateNames: true,
		},
	}

	// The tripwire answers every request with a 502, which go-azure-sdk treats
	// as retryable. A deadline caps how many times it retries - the SDK derives
	// its retry count from the context's remaining time - and keeps the test
	// from spending minutes re-sending a request it already recorded.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//nolint:errcheck // the call is expected to fail; what matters is that it was seen leaving
	schema.InternalMap(r.Schema).Diff(ctx, state, config, r.CustomizeDiff, ps.Meta, false)

	if !egress.reached(graphHost) {
		t.Errorf("expected the unguarded duplicate-name check to reach %s, recorded: %v", graphHost, egress.records())
	}
}

// noObjectIDTokenClient answers token requests with an otherwise valid offline
// token that carries no oid claim.
type noObjectIDTokenClient struct{}

func (noObjectIDTokenClient) Do(req *http.Request) (*http.Response, error) {
	payload, err := json.Marshal(map[string]any{
		"aud":   "https://" + graphHost,
		"iss":   "https://sts.windows.net/" + offlineTenantID + "/",
		"tid":   offlineTenantID,
		"appid": offlineClientID,
		// An empty scp keeps build() on the service principal branch, which
		// lists service principals filtered by the client ID. With "openid" it
		// would take the /me branch instead; both reach Graph.
		"scp": "",
	})
	if err != nil {
		return nil, err
	}
	enc := base64.RawURLEncoding.EncodeToString
	body, err := json.Marshal(map[string]any{
		"access_token": enc([]byte(`{"alg":"none","typ":"JWT"}`)) + "." + enc(payload) + ".offline",
		"token_type":   "Bearer",
		"expires_in":   offlineTokenLifetime,
	})
	if err != nil {
		return nil, err
	}

	return &http.Response{
		Status:        "200 OK",
		StatusCode:    http.StatusOK,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        http.Header{"Content-Type": []string{"application/json"}},
		Body:          io.NopCloser(strings.NewReader(string(body))),
		ContentLength: int64(len(body)),
		Request:       req,
	}, nil
}

// TestConfigureEgressesWithoutOfflineAuthentication is the positive control for
// TestOfflineConfigureDoesNotEgress at the Configure boundary specifically.
// Putting the Azure SDK's real token client back makes Configure attempt a
// genuine token request, which the tripwire records against loginHost. That is
// what makes the silence in TestOfflineConfigureDoesNotEgress meaningful: the
// tripwire does watch this code path, and EnableOfflineAuthentication is what
// keeps it quiet.
func TestConfigureEgressesWithoutOfflineAuthentication(t *testing.T) {
	EnableOfflineAuthentication()
	t.Cleanup(func() { auth.Client = offlineTokenClient{} })
	auth.Client = realAuthClient
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureAD provider schema: %v", err)
	}
	if _, err := configureOffline(context.Background(), p, nil); err == nil {
		t.Fatal("expected configuring the provider to fail without the offline token client")
	}

	if !egress.reached(loginHost) {
		t.Errorf("expected Configure to request a token from %s, recorded: %v", loginHost, egress.records())
	}
}
