/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"context"
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/xpprovider"
	"github.com/pkg/errors"

	"github.com/upbound/provider-azuread/config/administrativeunits"
	"github.com/upbound/provider-azuread/config/app"
	"github.com/upbound/provider-azuread/config/applications"
	"github.com/upbound/provider-azuread/config/conditionalaccess"
	"github.com/upbound/provider-azuread/config/directoryroles"
	"github.com/upbound/provider-azuread/config/groups"
	"github.com/upbound/provider-azuread/config/invitations"
	"github.com/upbound/provider-azuread/config/policies"
	"github.com/upbound/provider-azuread/config/serviceprincipaldelegated"
	"github.com/upbound/provider-azuread/config/serviceprincipals"
	"github.com/upbound/provider-azuread/config/synchronization"
	"github.com/upbound/provider-azuread/config/users"
)

const (
	resourcePrefix = "azuread"
	modulePath     = "github.com/upbound/provider-azuread"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// workaround for the TF Google v2.41.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(ctx context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p, err = xpprovider.GetProviderSchema(ctx)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithDefaultResourceOptions(
			resourceConfigurator(),
		),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(p),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		invitations.Configure,
		applications.Configure,
		groups.Configure,
		users.Configure,
		serviceprincipals.Configure,
		policies.Configure,
		administrativeunits.Configure,
		synchronization.Configure,
		conditionalaccess.Configure,
		directoryroles.Configure,
		app.Configure,
		serviceprincipaldelegated.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]ujconfig.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}
