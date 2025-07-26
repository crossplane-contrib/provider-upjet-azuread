// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	member "github.com/upbound/provider-azuread/internal/controller/cluster/administrativeunits/member"
	unit "github.com/upbound/provider-azuread/internal/controller/cluster/administrativeunits/unit"
	roleassignment "github.com/upbound/provider-azuread/internal/controller/cluster/app/roleassignment"
	application "github.com/upbound/provider-azuread/internal/controller/cluster/applications/application"
	approle "github.com/upbound/provider-azuread/internal/controller/cluster/applications/approle"
	certificate "github.com/upbound/provider-azuread/internal/controller/cluster/applications/certificate"
	federatedidentitycredential "github.com/upbound/provider-azuread/internal/controller/cluster/applications/federatedidentitycredential"
	password "github.com/upbound/provider-azuread/internal/controller/cluster/applications/password"
	preauthorized "github.com/upbound/provider-azuread/internal/controller/cluster/applications/preauthorized"
	accesspolicy "github.com/upbound/provider-azuread/internal/controller/cluster/conditionalaccess/accesspolicy"
	location "github.com/upbound/provider-azuread/internal/controller/cluster/conditionalaccess/location"
	customdirectoryrole "github.com/upbound/provider-azuread/internal/controller/cluster/directoryroles/customdirectoryrole"
	role "github.com/upbound/provider-azuread/internal/controller/cluster/directoryroles/role"
	roleassignmentdirectoryroles "github.com/upbound/provider-azuread/internal/controller/cluster/directoryroles/roleassignment"
	roleeligibilityschedulerequest "github.com/upbound/provider-azuread/internal/controller/cluster/directoryroles/roleeligibilityschedulerequest"
	group "github.com/upbound/provider-azuread/internal/controller/cluster/groups/group"
	membergroups "github.com/upbound/provider-azuread/internal/controller/cluster/groups/member"
	privilegedaccessgroupassignmentschedule "github.com/upbound/provider-azuread/internal/controller/cluster/identitygovernance/privilegedaccessgroupassignmentschedule"
	privilegedaccessgroupeligibilityschedule "github.com/upbound/provider-azuread/internal/controller/cluster/identitygovernance/privilegedaccessgroupeligibilityschedule"
	invitation "github.com/upbound/provider-azuread/internal/controller/cluster/invitations/invitation"
	claimsmappingpolicy "github.com/upbound/provider-azuread/internal/controller/cluster/policies/claimsmappingpolicy"
	providerconfig "github.com/upbound/provider-azuread/internal/controller/cluster/providerconfig"
	permissiongrant "github.com/upbound/provider-azuread/internal/controller/cluster/serviceprincipaldelegated/permissiongrant"
	certificateserviceprincipals "github.com/upbound/provider-azuread/internal/controller/cluster/serviceprincipals/certificate"
	claimsmappingpolicyassignment "github.com/upbound/provider-azuread/internal/controller/cluster/serviceprincipals/claimsmappingpolicyassignment"
	passwordserviceprincipals "github.com/upbound/provider-azuread/internal/controller/cluster/serviceprincipals/password"
	principal "github.com/upbound/provider-azuread/internal/controller/cluster/serviceprincipals/principal"
	tokensigningcertificate "github.com/upbound/provider-azuread/internal/controller/cluster/serviceprincipals/tokensigningcertificate"
	job "github.com/upbound/provider-azuread/internal/controller/cluster/synchronization/job"
	secret "github.com/upbound/provider-azuread/internal/controller/cluster/synchronization/secret"
	user "github.com/upbound/provider-azuread/internal/controller/cluster/users/user"
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
