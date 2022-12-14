package directoryroles

import "github.com/upbound/upjet/pkg/config"

const group = "directoryroles"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_custom_directory_role", func(r *config.Resource) {
		r.Kind = "CustomDirectoryRole"
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
