/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/upbound/provider-azuread/internal/controller/applications/application"
	group "github.com/upbound/provider-azuread/internal/controller/groups/group"
	member "github.com/upbound/provider-azuread/internal/controller/groups/member"
	invitation "github.com/upbound/provider-azuread/internal/controller/invitations/invitation"
	principal "github.com/upbound/provider-azuread/internal/controller/principals/principal"
	providerconfig "github.com/upbound/provider-azuread/internal/controller/providerconfig"
	user "github.com/upbound/provider-azuread/internal/controller/users/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		group.Setup,
		member.Setup,
		invitation.Setup,
		principal.Setup,
		providerconfig.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
