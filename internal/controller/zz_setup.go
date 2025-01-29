// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	member "github.com/upbound/provider-azuread/internal/controller/administrativeunits/member"
	unit "github.com/upbound/provider-azuread/internal/controller/administrativeunits/unit"
	roleassignment "github.com/upbound/provider-azuread/internal/controller/app/roleassignment"
	application "github.com/upbound/provider-azuread/internal/controller/applications/application"
	certificate "github.com/upbound/provider-azuread/internal/controller/applications/certificate"
	federatedidentitycredential "github.com/upbound/provider-azuread/internal/controller/applications/federatedidentitycredential"
	password "github.com/upbound/provider-azuread/internal/controller/applications/password"
	preauthorized "github.com/upbound/provider-azuread/internal/controller/applications/preauthorized"
	accesspolicy "github.com/upbound/provider-azuread/internal/controller/conditionalaccess/accesspolicy"
	location "github.com/upbound/provider-azuread/internal/controller/conditionalaccess/location"
	customdirectoryrole "github.com/upbound/provider-azuread/internal/controller/directoryroles/customdirectoryrole"
	role "github.com/upbound/provider-azuread/internal/controller/directoryroles/role"
	roleassignmentdirectoryroles "github.com/upbound/provider-azuread/internal/controller/directoryroles/roleassignment"
	group "github.com/upbound/provider-azuread/internal/controller/groups/group"
	membergroups "github.com/upbound/provider-azuread/internal/controller/groups/member"
	invitation "github.com/upbound/provider-azuread/internal/controller/invitations/invitation"
	claimsmappingpolicy "github.com/upbound/provider-azuread/internal/controller/policies/claimsmappingpolicy"
	accessgroupassignmentschedule "github.com/upbound/provider-azuread/internal/controller/privileged/accessgroupassignmentschedule"
	accessgroupeligibilityschedule "github.com/upbound/provider-azuread/internal/controller/privileged/accessgroupeligibilityschedule"
	providerconfig "github.com/upbound/provider-azuread/internal/controller/providerconfig"
	permissiongrant "github.com/upbound/provider-azuread/internal/controller/serviceprincipaldelegated/permissiongrant"
	certificateserviceprincipals "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/certificate"
	claimsmappingpolicyassignment "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/claimsmappingpolicyassignment"
	passwordserviceprincipals "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/password"
	principal "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/principal"
	tokensigningcertificate "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/tokensigningcertificate"
	job "github.com/upbound/provider-azuread/internal/controller/synchronization/job"
	secret "github.com/upbound/provider-azuread/internal/controller/synchronization/secret"
	user "github.com/upbound/provider-azuread/internal/controller/users/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		member.Setup,
		unit.Setup,
		roleassignment.Setup,
		application.Setup,
		certificate.Setup,
		federatedidentitycredential.Setup,
		password.Setup,
		preauthorized.Setup,
		accesspolicy.Setup,
		location.Setup,
		customdirectoryrole.Setup,
		role.Setup,
		roleassignmentdirectoryroles.Setup,
		group.Setup,
		membergroups.Setup,
		invitation.Setup,
		claimsmappingpolicy.Setup,
		accessgroupassignmentschedule.Setup,
		accessgroupeligibilityschedule.Setup,
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
