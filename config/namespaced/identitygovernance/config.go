// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package identitygovernance

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	group = "identitygovernance"

	groupResource     = "azuread_group"
	objectIDExtractor = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("object_id",true)`
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_privileged_access_group_assignment_schedule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "identitygovernance"
		r.ShortGroup = group
		r.Kind = "PrivilegedAccessGroupAssignmentSchedule"
		r.References["group_id"] = config.Reference{
			TerraformName: groupResource,
			Extractor:     objectIDExtractor,
		}
		r.References["principal_id"] = config.Reference{
			TerraformName: "azuread_user",
			Extractor:     objectIDExtractor,
		}
	})
	p.AddResourceConfigurator("azuread_privileged_access_group_eligibility_schedule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "identitygovernance"
		r.ShortGroup = group
		r.Kind = "PrivilegedAccessGroupEligibilitySchedule"
		r.References["group_id"] = config.Reference{
			TerraformName: groupResource,
			Extractor:     objectIDExtractor,
		}
		r.References["principal_id"] = config.Reference{
			TerraformName: "azuread_user",
			Extractor:     objectIDExtractor,
		}
	})
	p.AddResourceConfigurator("azuread_access_package_catalog", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "access"
		r.ShortGroup = group
		r.Kind = "AccessPackageCatalog"
	})
	p.AddResourceConfigurator("azuread_access_package", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "access"
		r.ShortGroup = group
		r.Kind = "AccessPackage"
		r.References["catalog_id"] = config.Reference{
			TerraformName: "azuread_access_package_catalog",
		}
	})
	p.AddResourceConfigurator("azuread_access_package_assignment_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "access"
		r.ShortGroup = group
		r.Kind = "AccessPackageAssignmentPolicy"
		r.References["access_package_id"] = config.Reference{
			TerraformName: "azuread_access_package",
		}
		// The subject blocks below all take the object ID of a directory
		// principal. upjet generates a single reference per field, so these
		// resolve against groups, matching the upstream Terraform examples
		// which use a group with a "groupMembers" subject type. Individual
		// users can still be targeted by setting objectId to a literal
		// object ID with a "singleUser" subject type.
		for _, path := range []string{
			"approval_settings.approval_stage.primary_approver.object_id",
			"approval_settings.approval_stage.alternative_approver.object_id",
			"assignment_review_settings.reviewer.object_id",
			"requestor_settings.requestor.object_id",
		} {
			r.References[path] = config.Reference{
				TerraformName: groupResource,
				Extractor:     objectIDExtractor,
			}
		}
	})
	p.AddResourceConfigurator("azuread_access_package_catalog_role_assignment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "access"
		r.ShortGroup = group
		r.Kind = "AccessPackageCatalogRoleAssignment"
		r.References["catalog_id"] = config.Reference{
			TerraformName: "azuread_access_package_catalog",
		}
		r.References["principal_object_id"] = config.Reference{
			TerraformName: groupResource,
			Extractor:     objectIDExtractor,
		}
	})
	p.AddResourceConfigurator("azuread_access_package_resource_catalog_association", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "access"
		r.ShortGroup = group
		r.Kind = "AccessPackageResourceCatalogAssociation"
		r.References["catalog_id"] = config.Reference{
			TerraformName: "azuread_access_package_catalog",
		}
		r.References["resource_origin_id"] = config.Reference{
			TerraformName: groupResource,
			Extractor:     objectIDExtractor,
		}
	})
	p.AddResourceConfigurator("azuread_access_package_resource_package_association", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "access"
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
