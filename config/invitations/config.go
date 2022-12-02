package invitations

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_invitation", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "invitations"
	})
}
