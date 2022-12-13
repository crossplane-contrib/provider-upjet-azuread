package synchronization

import "github.com/upbound/upjet/pkg/config"

const group = "synchronization"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_synchronization_job", func(r *config.Resource) {
		r.References["service_principal_id"] = config.Reference{
			Type: "github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1.Principal",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})

}
