#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# We have to delete privilegedaccessgroupassignmentschedule before deleting referenced resources.
# https://github.com/hashicorp/terraform-provider-azuread/issues/1399
# package-runtime 2025-06-21T20:59:21Z	DEBUG	provider-azuread	Successfully requested deletion of external resource	{"controller": "managed/identitygovernance.azuread.upbound.io/v1beta1, kind=privilegedaccessgroupassignmentschedule", "request": {"name":"privilegedaccessgroupassignmentschedule"}, "uid": "836eaaeb-3fe1-4cfe-b22e-d16bea623737", "version": "2304934", "external-name": "ebffd8ea-3912-43f1-9cbf-7d3c6fa39a26_member_5fa13e1b-5063-4f66-bc68-9b96e8d45e8c", "deletion-timestamp": "2025-06-21 20:56:44 +0000 UTC"}
# package-runtime E0621 20:59:55.149605       1 runtime.go:79] Observed a panic: "invalid memory address or nil pointer dereference" (runtime error: invalid memory address or nil pointer dereference)
# package-runtime goroutine 2865 [running]:
# package-runtime k8s.io/apimachinery/pkg/util/runtime.logPanic({0x21760e0, 0x3b66500})
# package-runtime 	k8s.io/apimachinery@v0.30.0/pkg/util/runtime/runtime.go:75 +0x85
# package-runtime github.com/crossplane/upjet/pkg/controller.(*panicHandler).recoverIfPanic(0xc00157bf38)
# package-runtime 	github.com/crossplane/upjet@v1.5.1/pkg/controller/external_async_tfpluginfw.go:154 +0x64
# package-runtime panic({0x21760e0?, 0x3b66500?})
# package-runtime 	runtime/panic.go:791 +0x132
# package-runtime github.com/manicminer/hamilton/msgraph.(*PrivilegedAccessGroupAssignmentScheduleRequestsClient).Create(0xc0019de620, {0x28b1a30, 0xc001a84b60}, {0xc001a5a060, {0xc00105ef7a, 0x6}, {0x249b08a, 0xb}, 0x0, 0x0, ...})
# package-runtime 	github.com/manicminer/hamilton@v0.71.0/msgraph/privilegedaccess_groups_assignmentschedule_requests.go:75 +0x281
# package-runtime github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance.revokeAssignmentRequest({0x28b1a30, 0xc001a84b60}, {0xc00185c408, {0x28b3498, 0xc0005950c8}, 0xc001abfa80, 0x0, {0x28b3930, 0x3beb440}}, 0xc0019de620, ...)
# package-runtime 	github.com/hashicorp/terraform-provider-azuread@v1.6.1-0.20230727144955-0adfe586f500/internal/services/identitygovernance/privileged_access_group_assignment_schedule_resource.go:294 +0x1d0
# package-runtime github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance.(*PrivilegedAccessGroupAssignmentScheduleResource).Delete.PrivilegedAccessGroupAssignmentScheduleResource.Delete.func1({0x28b1a30, 0xc001a84b60}, {0xc00185c408, {0x28b3498, 0xc0005950c8}, 0xc001abfa80, 0x0, {0x28b3930, 0x3beb440}})
# package-runtime 	github.com/hashicorp/terraform-provider-azuread@v1.6.1-0.20230727144955-0adfe586f500/internal/services/identitygovernance/privileged_access_group_assignment_schedule_resource.go:271 +0x315
# package-runtime github.com/hashicorp/terraform-provider-azuread/internal/sdk.(*ResourceWrapper).Resource.func4({0x28b1a30, 0xc001a84b60}, 0xc001a84b60?, {0x2150820?, 0xc00185c408?})
# package-runtime 	github.com/hashicorp/terraform-provider-azuread@v1.6.1-0.20230727144955-0adfe586f500/internal/sdk/wrapper_resource.go:72 +0x145
# package-runtime github.com/hashicorp/terraform-provider-azuread/internal/sdk.(*ResourceWrapper).Resource.(*ResourceWrapper).diagnosticsWrapper.diagnosticsWrapper.func12({0x28b1a30?, 0xc001a84b60?}, 0x45d964b800?, {0x2150820?, 0xc00185c408?})
# package-runtime 	github.com/hashicorp/terraform-provider-azuread@v1.6.1-0.20230727144955-0adfe586f500/internal/sdk/wrapper_resource.go:190 +0x59
# package-runtime github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema.(*Resource).delete(0xc0006f0000, {0x28b1a30, 0xc000262930}, 0xc001abfa80, {0x2150820, 0xc00185c408})
# package-runtime 	github.com/hashicorp/terraform-plugin-sdk/v2@v2.30.0/helper/schema/resource.go:829 +0x119
# package-runtime github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema.(*Resource).Apply(0xc0006f0000, {0x28b1a30, 0xc000262930}, 0xc0018305b0, 0xc001f0fa80, {0x2150820, 0xc00185c408})
# package-runtime 	github.com/hashicorp/terraform-plugin-sdk/v2@v2.30.0/helper/schema/resource.go:878 +0x5e5
# package-runtime github.com/crossplane/upjet/pkg/controller.(*terraformPluginSDKExternal).Delete(0xc001a6e0f0, {0x28b1a30, 0xc000262930}, {0xc001154d60?, 0xc000d4b3e0?})
# package-runtime 	github.com/crossplane/upjet@v1.5.1/pkg/controller/external_tfpluginsdk.go:738 +0x14c
# package-runtime github.com/crossplane/upjet/pkg/controller.(*terraformPluginSDKAsyncExternal).Delete.func1()
# package-runtime 	github.com/crossplane/upjet@v1.5.1/pkg/controller/external_async_tfpluginsdk.go:236 +0x1d7
# package-runtime created by github.com/crossplane/upjet/pkg/controller.(*terraformPluginSDKAsyncExternal).Delete in goroutine 1010
# package-runtime 	github.com/crossplane/upjet@v1.5.1/pkg/controller/external_async_tfpluginsdk.go:215 +0x22c

${KUBECTL} delete  privilegedaccessgroupassignmentschedule.identitygovernance.azuread.upbound.io -l testing.upbound.io/example-name=privilegedaccessgroupassignmentschedule

