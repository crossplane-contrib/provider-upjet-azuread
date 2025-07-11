// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package identitygovernance

import "github.com/crossplane/upjet/pkg/config"

const group = "identitygovernance"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_privileged_access_group_assignment_schedule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "identitygovernance"
		r.ShortGroup = group
		r.Kind = "PrivilegedAccessGroupAssignmentSchedule"
	})
	p.AddResourceConfigurator("azuread_privileged_access_group_eligibility_schedule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "identitygovernance"
		r.ShortGroup = group
		r.Kind = "PrivilegedAccessGroupEligibilitySchedule"
	})
}
