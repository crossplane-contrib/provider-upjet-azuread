// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"strings"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/pkg/config"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/crossplane/upjet/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/xpprovider"
	"github.com/pkg/errors"

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

const (
	resourcePrefix = "azuread"
	modulePath     = "github.com/upbound/provider-azuread"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// oldSingletonListAPIs is a newline-delimited list of Terraform resource
// names with converted singleton list APIs with at least CRD API version
// containing the old singleton list API. This is to prevent the API
// conversion for the newly added resources whose CRD APIs will already
// use embedded objects instead of the singleton lists and thus, will
// not possess a CRD API version with the singleton list. Thus, for
// the newly added resources (resources added after the singleton lists
// have been converted), we do not need the CRD API conversion
// functions that convert between singleton lists and embedded objects,
// but we need only the Terraform conversion functions.
// This list is immutable and represents the set of resources with the
// already generated CRD API versions with now converted singleton lists.
// Because new resources should never have singleton lists in their
// generated APIs, there should be no need to add them to this list.
// However, bugs might result in exceptions in the future.
// Please see:
// https://github.com/crossplane-contrib/provider-upjet-azuread/pull/123
// for more context on singleton list to embedded object conversions.
//
//go:embed old-singleton-list-apis.txt
var oldSingletonListAPIs string

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
	sdkProvider, err := xpprovider.GetProviderSchema(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get the Terraform SDK provider")
	}

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
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	bumpVersionsWithEmbeddedLists(pc)
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

func bumpVersionsWithEmbeddedLists(pc *ujconfig.Provider) {
	l := strings.Split(strings.TrimSpace(oldSingletonListAPIs), "\n")
	oldSLAPIs := make(map[string]struct{}, len(l))
	for _, n := range l {
		oldSLAPIs[n] = struct{}{}
	}

	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}
		if _, ok := oldSLAPIs[name]; ok {
			r.Version = "v1beta2"
			r.PreviousVersions = []string{"v1beta1"}
			// we would like to set the storage version to v1beta1 to facilitate
			// downgrades.
			r.SetCRDStorageVersion("v1beta1")
			// because the controller reconciles on the API version with the singleton list API,
			// no need for a Terraform conversion.
			r.ControllerReconcileVersion = "v1beta1"
			r.Conversions = []conversion.Conversion{
				conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes(), r.CRDListConversionPaths()...),
				conversion.NewSingletonListConversion("v1beta1", "v1beta2", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToEmbeddedObject),
				conversion.NewSingletonListConversion("v1beta2", "v1beta1", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToSingletonList)}
		} else {
			// the controller will be reconciling on the CRD API version
			// with the converted API (with embedded objects in place of
			// singleton lists), so we need the appropriate Terraform
			// converter in this case.
			r.TerraformConversions = []config.TerraformConversion{
				config.NewTFSingletonConversion(),
			}
		}
		pc.Resources[name] = r
	}
}
