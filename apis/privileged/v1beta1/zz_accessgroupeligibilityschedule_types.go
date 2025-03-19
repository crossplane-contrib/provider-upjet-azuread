// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessGroupEligibilityScheduleInitParameters struct {

	// The type of assignment to the group. Can be either member or owner.
	// The ID of the assignment to the group
	AssignmentType *string `json:"assignmentType,omitempty" tf:"assignment_type,omitempty"`

	// The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
	// The duration of the assignment, formatted as an ISO8601 duration string (e.g. P3D for 3 days)
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// 01-01T01:02:03Z).
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The Object ID of the Azure AD group to which the principal will be assigned.
	// The ID of the Group representing the scope of the assignment
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/groups/v1beta2.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in groups to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in groups to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The justification for this assignment. May be required by the role policy.
	// The justification for the assignment
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// Is this assigment permanently valid.
	// Is the assignment permanent
	PermanentAssignment *bool `json:"permanentAssignment,omitempty" tf:"permanent_assignment,omitempty"`

	// The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
	// The ID of the Principal assigned to the schedule
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/users/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// Reference to a User in users to populate principalId.
	// +kubebuilder:validation:Optional
	PrincipalIDRef *v1.Reference `json:"principalIdRef,omitempty" tf:"-"`

	// Selector for a User in users to populate principalId.
	// +kubebuilder:validation:Optional
	PrincipalIDSelector *v1.Selector `json:"principalIdSelector,omitempty" tf:"-"`

	// 01-01T01:02:03Z). If not provided, the assignment is immediately valid.
	// The date that this assignment starts, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// The ticket number in the ticket system approving this assignment. May be required by the role policy.
	// The ticket number authorising the assignment
	TicketNumber *string `json:"ticketNumber,omitempty" tf:"ticket_number,omitempty"`

	// The ticket system containing the ticket number approving this assignment. May be required by the role policy.
	// The ticket system authorising the assignment
	TicketSystem *string `json:"ticketSystem,omitempty" tf:"ticket_system,omitempty"`
}

type AccessGroupEligibilityScheduleObservation struct {

	// The type of assignment to the group. Can be either member or owner.
	// The ID of the assignment to the group
	AssignmentType *string `json:"assignmentType,omitempty" tf:"assignment_type,omitempty"`

	// The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
	// The duration of the assignment, formatted as an ISO8601 duration string (e.g. P3D for 3 days)
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// 01-01T01:02:03Z).
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The Object ID of the Azure AD group to which the principal will be assigned.
	// The ID of the Group representing the scope of the assignment
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// (String) The ID of this request.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The justification for this assignment. May be required by the role policy.
	// The justification for the assignment
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// Is this assigment permanently valid.
	// Is the assignment permanent
	PermanentAssignment *bool `json:"permanentAssignment,omitempty" tf:"permanent_assignment,omitempty"`

	// The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
	// The ID of the Principal assigned to the schedule
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// 01-01T01:02:03Z). If not provided, the assignment is immediately valid.
	// The date that this assignment starts, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// (String) The provisioning status of this request.
	// The status of the schedule
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The ticket number in the ticket system approving this assignment. May be required by the role policy.
	// The ticket number authorising the assignment
	TicketNumber *string `json:"ticketNumber,omitempty" tf:"ticket_number,omitempty"`

	// The ticket system containing the ticket number approving this assignment. May be required by the role policy.
	// The ticket system authorising the assignment
	TicketSystem *string `json:"ticketSystem,omitempty" tf:"ticket_system,omitempty"`
}

type AccessGroupEligibilityScheduleParameters struct {

	// The type of assignment to the group. Can be either member or owner.
	// The ID of the assignment to the group
	// +kubebuilder:validation:Optional
	AssignmentType *string `json:"assignmentType,omitempty" tf:"assignment_type,omitempty"`

	// The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
	// The duration of the assignment, formatted as an ISO8601 duration string (e.g. P3D for 3 days)
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// 01-01T01:02:03Z).
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)
	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The Object ID of the Azure AD group to which the principal will be assigned.
	// The ID of the Group representing the scope of the assignment
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/groups/v1beta2.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in groups to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in groups to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The justification for this assignment. May be required by the role policy.
	// The justification for the assignment
	// +kubebuilder:validation:Optional
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// Is this assigment permanently valid.
	// Is the assignment permanent
	// +kubebuilder:validation:Optional
	PermanentAssignment *bool `json:"permanentAssignment,omitempty" tf:"permanent_assignment,omitempty"`

	// The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
	// The ID of the Principal assigned to the schedule
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/users/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// Reference to a User in users to populate principalId.
	// +kubebuilder:validation:Optional
	PrincipalIDRef *v1.Reference `json:"principalIdRef,omitempty" tf:"-"`

	// Selector for a User in users to populate principalId.
	// +kubebuilder:validation:Optional
	PrincipalIDSelector *v1.Selector `json:"principalIdSelector,omitempty" tf:"-"`

	// 01-01T01:02:03Z). If not provided, the assignment is immediately valid.
	// The date that this assignment starts, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)
	// +kubebuilder:validation:Optional
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// The ticket number in the ticket system approving this assignment. May be required by the role policy.
	// The ticket number authorising the assignment
	// +kubebuilder:validation:Optional
	TicketNumber *string `json:"ticketNumber,omitempty" tf:"ticket_number,omitempty"`

	// The ticket system containing the ticket number approving this assignment. May be required by the role policy.
	// The ticket system authorising the assignment
	// +kubebuilder:validation:Optional
	TicketSystem *string `json:"ticketSystem,omitempty" tf:"ticket_system,omitempty"`
}

// AccessGroupEligibilityScheduleSpec defines the desired state of AccessGroupEligibilitySchedule
type AccessGroupEligibilityScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessGroupEligibilityScheduleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AccessGroupEligibilityScheduleInitParameters `json:"initProvider,omitempty"`
}

// AccessGroupEligibilityScheduleStatus defines the observed state of AccessGroupEligibilitySchedule.
type AccessGroupEligibilityScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessGroupEligibilityScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessGroupEligibilitySchedule is the Schema for the AccessGroupEligibilitySchedules API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type AccessGroupEligibilitySchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.assignmentType) || (has(self.initProvider) && has(self.initProvider.assignmentType))",message="spec.forProvider.assignmentType is a required parameter"
	Spec   AccessGroupEligibilityScheduleSpec   `json:"spec"`
	Status AccessGroupEligibilityScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessGroupEligibilityScheduleList contains a list of AccessGroupEligibilitySchedules
type AccessGroupEligibilityScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessGroupEligibilitySchedule `json:"items"`
}

// Repository type metadata.
var (
	AccessGroupEligibilitySchedule_Kind             = "AccessGroupEligibilitySchedule"
	AccessGroupEligibilitySchedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessGroupEligibilitySchedule_Kind}.String()
	AccessGroupEligibilitySchedule_KindAPIVersion   = AccessGroupEligibilitySchedule_Kind + "." + CRDGroupVersion.String()
	AccessGroupEligibilitySchedule_GroupVersionKind = CRDGroupVersion.WithKind(AccessGroupEligibilitySchedule_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessGroupEligibilitySchedule{}, &AccessGroupEligibilityScheduleList{})
}
