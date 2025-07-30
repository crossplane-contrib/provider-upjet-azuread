package config

import (
	"context"
	"github.com/crossplane/crossplane-runtime/pkg/errors"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/crossplane/upjet/pkg/schema/traverser"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/provider-azuread/config/namespaced/administrativeunits"
	"github.com/upbound/provider-azuread/config/namespaced/app"
	"github.com/upbound/provider-azuread/config/namespaced/applications"
	"github.com/upbound/provider-azuread/config/namespaced/conditionalaccess"
	"github.com/upbound/provider-azuread/config/namespaced/directoryroles"
	"github.com/upbound/provider-azuread/config/namespaced/groups"
	"github.com/upbound/provider-azuread/config/namespaced/identitygovernance"
	"github.com/upbound/provider-azuread/config/namespaced/invitations"
	"github.com/upbound/provider-azuread/config/namespaced/policies"
	"github.com/upbound/provider-azuread/config/namespaced/serviceprincipaldelegated"
	"github.com/upbound/provider-azuread/config/namespaced/serviceprincipals"
	"github.com/upbound/provider-azuread/config/namespaced/synchronization"
	"github.com/upbound/provider-azuread/config/namespaced/users"
)

// GetNamespacedProvider returns the namespaced provider configuration
func GetNamespacedProvider(ctx context.Context, sdkProvider *schema.Provider, generationProvider bool) (*ujconfig.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithDefaultResourceOptions(
			resourceConfigurator(),
		),
		ujconfig.WithRootGroup("azuread.m.upbound.io"),
		ujconfig.WithShortName("azuread"),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	registerTerraformConversions(pc)
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
		identitygovernance.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

func registerTerraformConversions(pc *ujconfig.Provider) {
	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}
		r.TerraformConversions = []ujconfig.TerraformConversion{
			ujconfig.NewTFSingletonConversion(),
		}
		pc.Resources[name] = r
	}
}
