package serviceprincipals

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_service_principal", func(r *config.Resource) {
		r.References["application_id"] = config.Reference{
			Type:      "github.com/upbound/provider-azuread/apis/applications/v1alpha1.Application",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("application_id",true)`,
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "serviceprincipals"

		config.MoveToStatus(r.TerraformResource, "features")
	})
	p.AddResourceConfigurator("azuread_claims_mapping_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "serviceprincipals"
	})
}
