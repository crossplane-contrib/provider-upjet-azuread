// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package directoryroles

import "github.com/crossplane/upjet/pkg/config"

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
	})
	p.AddResourceConfigurator("azuread_directory_role_assignment", func(r *config.Resource) {
		r.References["role_id"] = config.Reference{
			TerraformName: "azuread_directory_role",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("template_id",true)`,
		}
		r.References["principal_object_id"] = config.Reference{
			TerraformName: "azuread_user",
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
	p.AddResourceConfigurator("azuread_directory_role_eligibility_schedule_request", func(r *config.Resource) {
		r.References["role_definition_id"] = config.Reference{
			TerraformName: "azuread_directory_role",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("template_id",true)`,
		}
		r.References["principal_id"] = config.Reference{
			TerraformName: "azuread_user",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("object_id",true)`,
		}

		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
