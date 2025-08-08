// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package applications

import "github.com/crossplane/upjet/v2/pkg/config"

const group = "applications"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_application", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"tags",
			},
		}
	})
	p.AddResourceConfigurator("azuread_application_app_role", func(r *config.Resource) {
		r.References["application_id"] = config.Reference{
			TerraformName: "azuread_application",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_application_certificate", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_application_password", func(r *config.Resource) {
		r.References["application_id"] = config.Reference{
			TerraformName: "azuread_application",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_application_federated_identity_credential", func(r *config.Resource) {
		r.References["application_id"] = config.Reference{
			TerraformName: "azuread_application",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_application_pre_authorized", func(r *config.Resource) {
		r.References["application_id"] = config.Reference{
			TerraformName: "azuread_application",
		}
		r.References["authorized_client_id"] = config.Reference{
			TerraformName: "azuread_application",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("client_id",true)`,
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
