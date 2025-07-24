#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# We have to delete privilegedaccessgroupeligibilityschedule before deleting referenced resources.
# https://github.com/hashicorp/terraform-provider-azuread/issues/1399
# package-runtime 2025-06-21T20:18:22Z	DEBUG	provider-azuread	Successfully requested deletion of external resource	{"controller": "managed/identitygovernance.azuread.upbound.io/v1beta1, kind=privilegedaccessgroupeligibilityschedule", "request": {"name":"example"}, "uid": "d0381cd5-5691-4215-bc66-fc7c46824787", "version": "2298525", "external-name": "24219757-f21f-4235-a102-7a1f1bbc7fe3_member_f92d5141-746f-408a-b582-cdb6147618cb", "deletion-timestamp": "2025-06-21 20:09:19 +0000 UTC"}
# package-runtime E0621 20:18:45.080912       1 runtime.go:79] Observed a panic: "invalid memory address or nil pointer dereference" (runtime error: invalid memory address or nil pointer dereference)
# package-runtime goroutine 3253 [running]:
# package-runtime k8s.io/apimachinery/pkg/util/runtime.logPanic({0x21760e0, 0x3b66500})
# package-runtime 	k8s.io/apimachinery@v0.30.0/pkg/util/runtime/runtime.go:75 +0x85
# package-runtime github.com/crossplane/upjet/pkg/controller.(*panicHandler).recoverIfPanic(0xc00135ff38)
# package-runtime 	github.com/crossplane/upjet@v1.5.1/pkg/controller/external_async_tfpluginfw.go:154 +0x64
# package-runtime panic({0x21760e0?, 0x3b66500?})
# package-runtime 	runtime/panic.go:791 +0x132
# package-runtime github.com/manicminer/hamilton/msgraph.(*PrivilegedAccessGroupEligibilityScheduleRequestsClient).Create(0xc000366070, {0x28b1a30, 0xc0003677a0}, {0xc0012433a0, {0xc001862a9a, 0x6}, {0x249b08a, 0xb}, 0x0, 0x0, ...})
# package-runtime 	github.com/manicminer/hamilton@v0.71.0/msgraph/privilegedaccess_groups_eligibilityschedule_request.go:74 +0x281
# package-runtime github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance.revokeEligibilityRequest({0x28b1a30, 0xc0003677a0}, {0xc0011d5608, {0x28b3498, 0xc0003401c8}, 0xc00118ac00, 0x0, {0x28b3930, 0x3beb440}}, 0xc000366070, ...)
# package-runtime 	github.com/hashicorp/terraform-provider-azuread@v1.6.1-0.20230727144955-0adfe586f500/internal/services/identitygovernance/privileged_access_group_eligiblity_schedule_resource.go:294 +0x1d0
# package-runtime github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance.(*PrivilegedAccessGroupEligibilityScheduleResource).Delete.PrivilegedAccessGroupEligibilityScheduleResource.Delete.func1({0x28b1a30, 0xc0003677a0}, {0xc0011d5608, {0x28b3498, 0xc0003401c8}, 0xc00118ac00, 0x0, {0x28b3930, 0x3beb440}})
# package-runtime 	github.com/hashicorp/terraform-provider-azuread@v1.6.1-0.20230727144955-0adfe586f500/internal/services/identitygovernance/privileged_access_group_eligiblity_schedule_resource.go:271 +0x315
# package-runtime github.com/hashicorp/terraform-provider-azuread/internal/sdk.(*ResourceWrapper).Resource.func4({0x28b1a30, 0xc0003677a0}, 0xc0003677a0?, {0x2150820?, 0xc0011d5608?})
# package-runtime 	github.com/hashicorp/terraform-provider-azuread@v1.6.1-0.20230727144955-0adfe586f500/internal/sdk/wrapper_resource.go:72 +0x145
# package-runtime github.com/hashicorp/terraform-provider-azuread/internal/sdk.(*ResourceWrapper).Resource.(*ResourceWrapper).diagnosticsWrapper.diagnosticsWrapper.func12({0x28b1a30?, 0xc0003677a0?}, 0x45d964b800?, {0x2150820?, 0xc0011d5608?})
# package-runtime 	github.com/hashicorp/terraform-provider-azuread@v1.6.1-0.20230727144955-0adfe586f500/internal/sdk/wrapper_resource.go:190 +0x59
# package-runtime github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema.(*Resource).delete(0xc000199dc0, {0x28b1a30, 0xc000b46e70}, 0xc00118ac00, {0x2150820, 0xc0011d5608})
# package-runtime 	github.com/hashicorp/terraform-plugin-sdk/v2@v2.30.0/helper/schema/resource.go:829 +0x119
# package-runtime github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema.(*Resource).Apply(0xc000199dc0, {0x28b1a30, 0xc000b46e70}, 0xc0015961a0, 0xc000bfe280, {0x2150820, 0xc0011d5608})
# package-runtime 	github.com/hashicorp/terraform-plugin-sdk/v2@v2.30.0/helper/schema/resource.go:878 +0x5e5
# package-runtime github.com/crossplane/upjet/pkg/controller.(*terraformPluginSDKExternal).Delete(0xc000afc4b0, {0x28b1a30, 0xc000b46e70}, {0xc000da62c0?, 0x0?})
# package-runtime 	github.com/crossplane/upjet@v1.5.1/pkg/controller/external_tfpluginsdk.go:738 +0x14c
# package-runtime github.com/crossplane/upjet/pkg/controller.(*terraformPluginSDKAsyncExternal).Delete.func1()
# package-runtime 	github.com/crossplane/upjet@v1.5.1/pkg/controller/external_async_tfpluginsdk.go:236 +0x1d7
# package-runtime created by github.com/crossplane/upjet/pkg/controller.(*terraformPluginSDKAsyncExternal).Delete in goroutine 1183
# package-runtime 	github.com/crossplane/upjet@v1.5.1/pkg/controller/external_async_tfpluginsdk.go:215 +0x22c

${KUBECTL} delete  privilegedaccessgroupeligibilityschedule.identitygovernance.azuread.upbound.io -l testing.upbound.io/example-name=privilegedaccessgroupeligibilityschedule

