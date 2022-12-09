package groups

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "groups"
	})
	p.AddResourceConfigurator("azuread_group_member", func(r *config.Resource) {
		r.References["group_object_id"] = config.Reference{
			Type: "Group",
		}
		r.References["member_object_id"] = config.Reference{
			Type: "github.com/upbound/provider-azuread/apis/users/v1beta1.User",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "groups"
	})
}
