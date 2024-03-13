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

type MemberInitParameters struct {

	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	// The object ID of the administrative unit
	// +crossplane:generate:reference:type=Unit
	AdministrativeUnitObjectID *string `json:"administrativeUnitObjectId,omitempty" tf:"administrative_unit_object_id,omitempty"`

	// Reference to a Unit to populate administrativeUnitObjectId.
	// +kubebuilder:validation:Optional
	AdministrativeUnitObjectIDRef *v1.Reference `json:"administrativeUnitObjectIdRef,omitempty" tf:"-"`

	// Selector for a Unit to populate administrativeUnitObjectId.
	// +kubebuilder:validation:Optional
	AdministrativeUnitObjectIDSelector *v1.Selector `json:"administrativeUnitObjectIdSelector,omitempty" tf:"-"`

	// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	// The object ID of the member
	MemberObjectID *string `json:"memberObjectId,omitempty" tf:"member_object_id,omitempty"`
}

type MemberObservation struct {

	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	// The object ID of the administrative unit
	AdministrativeUnitObjectID *string `json:"administrativeUnitObjectId,omitempty" tf:"administrative_unit_object_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	// The object ID of the member
	MemberObjectID *string `json:"memberObjectId,omitempty" tf:"member_object_id,omitempty"`
}

type MemberParameters struct {

	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	// The object ID of the administrative unit
	// +crossplane:generate:reference:type=Unit
	// +kubebuilder:validation:Optional
	AdministrativeUnitObjectID *string `json:"administrativeUnitObjectId,omitempty" tf:"administrative_unit_object_id,omitempty"`

	// Reference to a Unit to populate administrativeUnitObjectId.
	// +kubebuilder:validation:Optional
	AdministrativeUnitObjectIDRef *v1.Reference `json:"administrativeUnitObjectIdRef,omitempty" tf:"-"`

	// Selector for a Unit to populate administrativeUnitObjectId.
	// +kubebuilder:validation:Optional
	AdministrativeUnitObjectIDSelector *v1.Selector `json:"administrativeUnitObjectIdSelector,omitempty" tf:"-"`

	// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	// The object ID of the member
	// +kubebuilder:validation:Optional
	MemberObjectID *string `json:"memberObjectId,omitempty" tf:"member_object_id,omitempty"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
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
	InitProvider MemberInitParameters `json:"initProvider,omitempty"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Member is the Schema for the Members API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MemberSpec   `json:"spec"`
	Status            MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}
