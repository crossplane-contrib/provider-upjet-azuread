package serviceprincipals

import "github.com/crossplane/upjet/pkg/config"

const group = "serviceprincipals"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_service_principal", func(r *config.Resource) {
		r.References["application_id"] = config.Reference{
			Type:      "github.com/upbound/provider-azuread/apis/applications/v1beta1.Application",
			Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("application_id",true)`,
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group

		delete(r.TerraformResource.Schema, "features")

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"tags",
			},
		}
	})
	p.AddResourceConfigurator("azuread_service_principal_claims_mapping_policy_assignment", func(r *config.Resource) {
		r.Kind = "ClaimsMappingPolicyAssignment"
		r.References["service_principal_id"] = config.Reference{
			Type: "Principal",
		}
		r.References["claims_mapping_policy_id"] = config.Reference{
			Type: "github.com/upbound/provider-azuread/apis/policies/v1beta1.ClaimsMappingPolicy",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_service_principal_certificate", func(r *config.Resource) {
		r.Kind = "Certificate"
		r.References["service_principal_id"] = config.Reference{
			Type: "Principal",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_service_principal_password", func(r *config.Resource) {
		r.Kind = "Password"
		r.References["service_principal_id"] = config.Reference{
			Type: "Principal",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_service_principal_token_signing_certificate", func(r *config.Resource) {
		r.Kind = "TokenSigningCertificate"
		r.References["service_principal_id"] = config.Reference{
			TerraformName: "azuread_service_principal",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
