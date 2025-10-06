// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package policies

import "github.com/crossplane/upjet/v2/pkg/config"

const group = "policies"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_claims_mapping_policy", func(r *config.Resource) {
		r.Kind = "ClaimsMappingPolicy"
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_group_role_management_policy", func(r *config.Resource) {
		r.Kind = "GroupRoleManagementPolicy"
		r.ShortGroup = group
		r.References["group_id"] = config.Reference{
			TerraformName: "azuread_group",
		}
	})
}
