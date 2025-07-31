// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package conditionalaccess

import "github.com/crossplane/upjet/v2/pkg/config"

const group = "conditionalaccess"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_named_location", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_conditional_access_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
