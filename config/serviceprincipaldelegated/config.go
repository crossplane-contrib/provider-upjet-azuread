package serviceprincipaldelegated

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_service_principal_delegated_permission_grant", func(r *config.Resource) {
		r.Kind = "PermissionGrant"
		r.References["service_principal_object_id"] = config.Reference{
			Type: "github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1.Principal",
		}
		r.References["resource_service_principal_object_id"] = config.Reference{
			Type: "github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1.Principal",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "serviceprincipaldelegated"
	})
}
