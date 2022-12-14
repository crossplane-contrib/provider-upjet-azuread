/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	member "github.com/upbound/provider-azuread/internal/controller/administrativeunits/member"
	unit "github.com/upbound/provider-azuread/internal/controller/administrativeunits/unit"
	application "github.com/upbound/provider-azuread/internal/controller/applications/application"
	certificate "github.com/upbound/provider-azuread/internal/controller/applications/certificate"
	federatedidentitycredential "github.com/upbound/provider-azuread/internal/controller/applications/federatedidentitycredential"
	password "github.com/upbound/provider-azuread/internal/controller/applications/password"
	preauthorized "github.com/upbound/provider-azuread/internal/controller/applications/preauthorized"
	policy "github.com/upbound/provider-azuread/internal/controller/conditionalaccess/policy"
	group "github.com/upbound/provider-azuread/internal/controller/groups/group"
	membergroups "github.com/upbound/provider-azuread/internal/controller/groups/member"
	invitation "github.com/upbound/provider-azuread/internal/controller/invitations/invitation"
	claimsmappingpolicy "github.com/upbound/provider-azuread/internal/controller/policies/claimsmappingpolicy"
	providerconfig "github.com/upbound/provider-azuread/internal/controller/providerconfig"
	certificateserviceprincipals "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/certificate"
	claimsmappingpolicyassignment "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/claimsmappingpolicyassignment"
	passwordserviceprincipals "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/password"
	principal "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/principal"
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
		application.Setup,
		certificate.Setup,
		federatedidentitycredential.Setup,
		password.Setup,
		preauthorized.Setup,
		policy.Setup,
		group.Setup,
		membergroups.Setup,
		invitation.Setup,
		claimsmappingpolicy.Setup,
		providerconfig.Setup,
		certificateserviceprincipals.Setup,
		claimsmappingpolicyassignment.Setup,
		passwordserviceprincipals.Setup,
		principal.Setup,
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
