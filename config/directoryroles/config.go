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
	p.AddResourceConfigurator("azuread_directory_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group

		config.MoveToStatus(r.TerraformResource, "template_id")

		//r.LateInitializer = config.LateInitializer{
		//	IgnoredFields: []string{
		//		"template_id",
		//	},
		//}
	})
	p.AddResourceConfigurator("azuread_directory_role_assignment", func(r *config.Resource) {
		r.References["role_id"] = config.Reference{
			Type:      "Role",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("template_id",true)`,
		}
		r.References["principal_object_id"] = config.Reference{
			Type: "github.com/upbound/provider-azuread/apis/users/v1beta1.User",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"directory_scope_object_id",
			},
		}
	})
	p.AddResourceConfigurator("azuread_directory_role_member", func(r *config.Resource) {
		r.References["role_object_id"] = config.Reference{
			Type: "Role",
		}
		r.References["member_object_id"] = config.Reference{
			Type: "github.com/upbound/provider-azuread/apis/users/v1beta1.User",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
