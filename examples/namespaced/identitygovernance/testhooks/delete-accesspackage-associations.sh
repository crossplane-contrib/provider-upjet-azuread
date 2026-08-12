#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# We have to fully delete the access package associations before deleting the access
# packages and catalogs they belong to.
#
# The upstream Terraform provider does not tolerate a 404 when the parent access package
# or catalog is already gone, in either the read or the delete path:
#
#   - internal/services/identitygovernance/identitygovernance.go:27
#     GetAccessPackageResourcesRoleScope turns a 404 on the parent access package into an
#     error, so the correct "roleScope == nil -> remove from state" branch in the read of
#     access_package_resource_package_association_resource.go:168 is never reached.
#   - access_package_resource_package_association_resource.go:205
#     delete treats a 404 as a failure.
#   - access_package_resource_catalog_association_resource.go:151,187
#     the same, with the catalog as the parent.
#
# Because the observe never reports "does not exist", the finalizer is never removed and
# the managed resource is stuck forever:
#
#   async delete failed: failed to delete the resource: [{0 Deleting Identity Governance
#   Entitlement Management Access Package Id Access Package Resource Role Scope (...)
#   unexpected status 404 (404 Not Found) with error: AccessPackageNotFound: The access
#   package was not found.}]
#
# Terraform normally hides this because `terraform destroy` tears down in reverse
# dependency order. Crossplane deletes managed resources concurrently, and uptest issues
# every delete with --wait=false, so the parent frequently disappears first.
#
# Deleting the associations here without --wait, while their parents still exist, keeps
# both the read and the delete paths on their happy path.
#
# Upstream issue: https://github.com/hashicorp/terraform-provider-azuread/issues/1923
# Remove this hook once that is fixed and the provider is updated.

# The resource package association references the catalog association, so it has to go first.
${KUBECTL} delete accesspackageresourcepackageassociation.identitygovernance.azuread.m.upbound.io \
  --all --namespace upbound-system --ignore-not-found --timeout 10m

${KUBECTL} delete accesspackageresourcecatalogassociation.identitygovernance.azuread.m.upbound.io \
  --all --namespace upbound-system --ignore-not-found --timeout 10m
