/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/upbound/provider-azuread/internal/controller/applications/application"
	certificate "github.com/upbound/provider-azuread/internal/controller/applications/certificate"
	group "github.com/upbound/provider-azuread/internal/controller/groups/group"
	member "github.com/upbound/provider-azuread/internal/controller/groups/member"
	invitation "github.com/upbound/provider-azuread/internal/controller/invitations/invitation"
	claimsmappingpolicy "github.com/upbound/provider-azuread/internal/controller/policies/claimsmappingpolicy"
	providerconfig "github.com/upbound/provider-azuread/internal/controller/providerconfig"
	certificateserviceprincipals "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/certificate"
	claimsmappingpolicyassignment "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/claimsmappingpolicyassignment"
	password "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/password"
	principal "github.com/upbound/provider-azuread/internal/controller/serviceprincipals/principal"
	user "github.com/upbound/provider-azuread/internal/controller/users/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		certificate.Setup,
		group.Setup,
		member.Setup,
		invitation.Setup,
		claimsmappingpolicy.Setup,
		providerconfig.Setup,
		certificateserviceprincipals.Setup,
		claimsmappingpolicyassignment.Setup,
		password.Setup,
		principal.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
