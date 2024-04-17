// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package groups

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "groups"
	})
	p.AddResourceConfigurator("azuread_group_member", func(r *config.Resource) {
		r.References["group_object_id"] = config.Reference{
			TerraformName: "azuread_group",
		}
		r.References["member_object_id"] = config.Reference{
			TerraformName: "azuread_user",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = "groups"
	})
}
