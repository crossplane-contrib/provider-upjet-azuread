// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package groups

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "groups"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"owners"},
		}
	})
	p.AddResourceConfigurator("azuread_group_member", func(r *config.Resource) {
		r.References["group_object_id"] = config.Reference{
			TerraformName: "azuread_group",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("object_id",true)`,
		}
		r.References["member_object_id"] = config.Reference{
			TerraformName: "azuread_user",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("object_id",true)`,
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "groups"
	})
}
