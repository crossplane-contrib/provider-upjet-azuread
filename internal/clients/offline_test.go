// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfazureclient "github.com/hashicorp/terraform-provider-azuread/xpprovider"
)

// keyDisplayName is the azuread_group attribute the offline diff tests change.
const keyDisplayName = "display_name"

// TestOfflineConfigure asserts that the AzureAD Terraform provider can be
// configured with no credentials and no access to Azure. The assertion is
// meaningful because the offline configuration's client secret is a
// placeholder: had the token request left the process, Azure would have
// rejected it and Configure would have failed.
func TestOfflineConfigure(t *testing.T) {
	EnableOfflineAuthentication()
	if _, ok := auth.Client.(offlineTokenClient); !ok {
		t.Fatalf("EnableOfflineAuthentication did not install the offline token client, auth.Client is %T", auth.Client)
	}

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureAD provider schema: %v", err)
	}

	ps, err := configureOffline(context.Background(), p, nil)
	if err != nil {
		t.Fatalf("cannot configure the AzureAD provider offline: %v", err)
	}
	if ps.Meta == nil {
		t.Fatal("offline configuration left terraform.Setup.Meta unset, the provider's CustomizeDiff functions need it")
	}
}

// TestOfflineDiff asserts that an offline configured provider can produce a
// diff, by making the call the upjet Terraform plugin SDK external client makes
// on the diff server's behalf. azuread_group is used because its CustomizeDiff
// function dereferences terraform.Setup.Meta.
func TestOfflineDiff(t *testing.T) {
	EnableOfflineAuthentication()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureAD provider schema: %v", err)
	}
	ps, err := configureOffline(context.Background(), p, nil)
	if err != nil {
		t.Fatalf("cannot configure the AzureAD provider offline: %v", err)
	}

	r := p.ResourcesMap["azuread_group"]
	state := &tfsdk.InstanceState{
		ID: offlineObjectID,
		Attributes: map[string]string{
			"id":               offlineObjectID,
			"object_id":        offlineObjectID,
			keyDisplayName:     "old-name",
			"description":      "before",
			"security_enabled": valTrue,
			"mail_enabled":     "false",
		},
	}
	config := &tfsdk.ResourceConfig{
		Config: map[string]any{
			keyDisplayName:     "new-name",
			"description":      "after",
			"security_enabled": true,
			"mail_enabled":     false,
		},
	}

	diff, err := schema.InternalMap(r.Schema).Diff(context.Background(), state, config, r.CustomizeDiff, ps.Meta, false)
	if err != nil {
		t.Fatalf("cannot diff azuread_group offline: %v", err)
	}

	got := map[string][2]string{}
	for k, a := range diff.Attributes {
		if a.Old != a.New {
			got[k] = [2]string{a.Old, a.New}
		}
	}
	want := map[string][2]string{
		keyDisplayName:            {"old-name", "new-name"},
		"description":             {"before", "after"},
		"prevent_duplicate_names": {"", "false"},
		"writeback_enabled":       {"", "false"},
	}
	if d := cmp.Diff(want, got); d != "" {
		t.Errorf("unexpected diff for azuread_group -want, +got:\n%s", d)
	}
}

// TestOfflineDiffBlocksOutboundRequest asserts that a CustomizeDiff function
// which calls Microsoft Graph fails immediately with the guard's error rather
// than retrying against an API that will not accept the offline access token.
// azuread_group reaches the API only when prevent_duplicate_names is set and
// the display name changes on a resource that already has an ID.
func TestOfflineDiffBlocksOutboundRequest(t *testing.T) {
	EnableOfflineAuthentication()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureAD provider schema: %v", err)
	}
	ps, err := configureOffline(context.Background(), p, nil)
	if err != nil {
		t.Fatalf("cannot configure the AzureAD provider offline: %v", err)
	}

	r := p.ResourcesMap["azuread_group"]
	state := &tfsdk.InstanceState{
		ID: offlineObjectID,
		Attributes: map[string]string{
			"id":                      offlineObjectID,
			"object_id":               offlineObjectID,
			keyDisplayName:            "old-name",
			"security_enabled":        valTrue,
			"mail_enabled":            "false",
			"prevent_duplicate_names": valTrue,
		},
	}
	config := &tfsdk.ResourceConfig{
		Config: map[string]any{
			keyDisplayName:            "new-name",
			"security_enabled":        true,
			"mail_enabled":            false,
			"prevent_duplicate_names": true,
		},
	}

	start := time.Now()
	_, err = schema.InternalMap(r.Schema).Diff(context.Background(), state, config, r.CustomizeDiff, ps.Meta, false)
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected the offline diff to fail, the duplicate name check must reach Microsoft Graph")
	}
	if !strings.Contains(err.Error(), "blocked an outbound Microsoft Graph request") {
		t.Errorf("expected the request guard's error, got: %v", err)
	}
	// The guard's value is that it fails at once. Without it the SDK retries
	// with backoff for minutes before giving up.
	if elapsed > 10*time.Second {
		t.Errorf("expected the blocked request to fail immediately, took %s", elapsed)
	}
}
