/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MemberObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MemberParameters struct {

	// The object ID of the group you want to add the member to
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	GroupObjectID *string `json:"groupObjectId,omitempty" tf:"group_object_id,omitempty"`

	// Reference to a Group to populate groupObjectId.
	// +kubebuilder:validation:Optional
	GroupObjectIDRef *v1.Reference `json:"groupObjectIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupObjectId.
	// +kubebuilder:validation:Optional
	GroupObjectIDSelector *v1.Selector `json:"groupObjectIdSelector,omitempty" tf:"-"`

	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/users/v1alpha1.User
	// +kubebuilder:validation:Optional
	MemberObjectID *string `json:"memberObjectId,omitempty" tf:"member_object_id,omitempty"`

	// Reference to a User in users to populate memberObjectId.
	// +kubebuilder:validation:Optional
	MemberObjectIDRef *v1.Reference `json:"memberObjectIdRef,omitempty" tf:"-"`

	// Selector for a User in users to populate memberObjectId.
	// +kubebuilder:validation:Optional
	MemberObjectIDSelector *v1.Selector `json:"memberObjectIdSelector,omitempty" tf:"-"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Member is the Schema for the Members API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
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