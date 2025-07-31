// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package administrativeunits

import "github.com/crossplane/upjet/pkg/config"

const group = "administrativeunits"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_administrative_unit", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_administrative_unit_member", func(r *config.Resource) {
		r.Kind = "Member"
		r.References["administrative_unit_object_id"] = config.Reference{
			TerraformName: "azuread_administrative_unit",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
