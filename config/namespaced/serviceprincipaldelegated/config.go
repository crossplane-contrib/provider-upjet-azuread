// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package serviceprincipaldelegated

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_service_principal_delegated_permission_grant", func(r *config.Resource) {
		r.Kind = "PermissionGrant"
		r.References["service_principal_object_id"] = config.Reference{
			TerraformName: "azuread_service_principal",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("object_id",true)`,
		}
		r.References["resource_service_principal_object_id"] = config.Reference{
			TerraformName: "azuread_service_principal",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("object_id",true)`,
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "serviceprincipaldelegated"
	})
}
