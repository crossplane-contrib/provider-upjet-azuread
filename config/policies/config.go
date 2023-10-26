package policies

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_claims_mapping_policy", func(r *config.Resource) {
		r.Kind = "ClaimsMappingPolicy"
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "policies"
	})
}
