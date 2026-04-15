// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package accesspackages

import "github.com/crossplane/upjet/v2/pkg/config"

const group = "accesspackages"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_access_package_catalog", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AccessPackageCatalog"
	})
	p.AddResourceConfigurator("azuread_access_package", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AccessPackage"
		r.References["catalog_id"] = config.Reference{
			TerraformName: "azuread_access_package_catalog",
		}
	})
	p.AddResourceConfigurator("azuread_access_package_assignment_policy", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AccessPackageAssignmentPolicy"
		r.References["access_package_id"] = config.Reference{
			TerraformName: "azuread_access_package",
		}
	})
	p.AddResourceConfigurator("azuread_access_package_catalog_role_assignment", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AccessPackageCatalogRoleAssignment"
		r.References["catalog_id"] = config.Reference{
			TerraformName: "azuread_access_package_catalog",
		}
		r.References["principal_object_id"] = config.Reference{
			TerraformName: "azuread_user",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("object_id",true)`,
		}
	})
	p.AddResourceConfigurator("azuread_access_package_resource_catalog_association", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AccessPackageResourceCatalogAssociation"
		r.References["catalog_id"] = config.Reference{
			TerraformName: "azuread_access_package_catalog",
		}
	})
	p.AddResourceConfigurator("azuread_access_package_resource_package_association", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AccessPackageResourcePackageAssociation"
		r.References["access_package_id"] = config.Reference{
			TerraformName: "azuread_access_package",
		}
		r.References["catalog_resource_association_id"] = config.Reference{
			TerraformName: "azuread_access_package_resource_catalog_association",
		}
	})
}
