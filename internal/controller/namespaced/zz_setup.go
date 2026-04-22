// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	member "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/administrativeunits/member"
	unit "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/administrativeunits/unit"
	roleassignment "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/app/roleassignment"
	application "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/applications/application"
	approle "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/applications/approle"
	certificate "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/applications/certificate"
	federatedidentitycredential "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/applications/federatedidentitycredential"
	flexiblefederatedidentitycredential "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/applications/flexiblefederatedidentitycredential"
	password "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/applications/password"
	preauthorized "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/applications/preauthorized"
	accesspolicy "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/conditionalaccess/accesspolicy"
	location "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/conditionalaccess/location"
	customdirectoryrole "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/directoryroles/customdirectoryrole"
	role "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/directoryroles/role"
	roleassignmentdirectoryroles "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/directoryroles/roleassignment"
	roleeligibilityschedulerequest "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/directoryroles/roleeligibilityschedulerequest"
	group "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/groups/group"
	membergroups "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/groups/member"
	privilegedaccessgroupassignmentschedule "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/identitygovernance/privilegedaccessgroupassignmentschedule"
	privilegedaccessgroupeligibilityschedule "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/identitygovernance/privilegedaccessgroupeligibilityschedule"
	invitation "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/invitations/invitation"
	claimsmappingpolicy "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/policies/claimsmappingpolicy"
	grouprolemanagementpolicy "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/policies/grouprolemanagementpolicy"
	providerconfig "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/providerconfig"
	permissiongrant "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/serviceprincipaldelegated/permissiongrant"
	certificateserviceprincipals "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/serviceprincipals/certificate"
	claimsmappingpolicyassignment "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/serviceprincipals/claimsmappingpolicyassignment"
	passwordserviceprincipals "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/serviceprincipals/password"
	principal "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/serviceprincipals/principal"
	tokensigningcertificate "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/serviceprincipals/tokensigningcertificate"
	job "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/synchronization/job"
	secret "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/synchronization/secret"
	user "github.com/upbound/provider-azuread/v2/internal/controller/namespaced/users/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		member.Setup,
		unit.Setup,
		roleassignment.Setup,
		application.Setup,
		approle.Setup,
		certificate.Setup,
		federatedidentitycredential.Setup,
		flexiblefederatedidentitycredential.Setup,
		password.Setup,
		preauthorized.Setup,
		accesspolicy.Setup,
		location.Setup,
		customdirectoryrole.Setup,
		role.Setup,
		roleassignmentdirectoryroles.Setup,
		roleeligibilityschedulerequest.Setup,
		group.Setup,
		membergroups.Setup,
		privilegedaccessgroupassignmentschedule.Setup,
		privilegedaccessgroupeligibilityschedule.Setup,
		invitation.Setup,
		claimsmappingpolicy.Setup,
		grouprolemanagementpolicy.Setup,
		providerconfig.Setup,
		permissiongrant.Setup,
		certificateserviceprincipals.Setup,
		claimsmappingpolicyassignment.Setup,
		passwordserviceprincipals.Setup,
		principal.Setup,
		tokensigningcertificate.Setup,
		job.Setup,
		secret.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		member.SetupGated,
		unit.SetupGated,
		roleassignment.SetupGated,
		application.SetupGated,
		approle.SetupGated,
		certificate.SetupGated,
		federatedidentitycredential.SetupGated,
		flexiblefederatedidentitycredential.SetupGated,
		password.SetupGated,
		preauthorized.SetupGated,
		accesspolicy.SetupGated,
		location.SetupGated,
		customdirectoryrole.SetupGated,
		role.SetupGated,
		roleassignmentdirectoryroles.SetupGated,
		roleeligibilityschedulerequest.SetupGated,
		group.SetupGated,
		membergroups.SetupGated,
		privilegedaccessgroupassignmentschedule.SetupGated,
		privilegedaccessgroupeligibilityschedule.SetupGated,
		invitation.SetupGated,
		claimsmappingpolicy.SetupGated,
		grouprolemanagementpolicy.SetupGated,
		providerconfig.SetupGated,
		permissiongrant.SetupGated,
		certificateserviceprincipals.SetupGated,
		claimsmappingpolicyassignment.SetupGated,
		passwordserviceprincipals.SetupGated,
		principal.SetupGated,
		tokensigningcertificate.SetupGated,
		job.SetupGated,
		secret.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
