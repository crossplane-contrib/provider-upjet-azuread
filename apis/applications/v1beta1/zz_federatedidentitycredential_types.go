/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FederatedIdentityCredentialInitParameters struct {

	// List of audiences that can appear in the external token. This specifies what should be accepted in the aud claim of incoming tokens.
	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	Audiences []*string `json:"audiences,omitempty" tf:"audiences,omitempty"`

	// A description for the federated identity credential.
	// A description for the federated identity credential
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
	// A unique display name for the federated identity credential
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type FederatedIdentityCredentialObservation struct {

	// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
	// The object ID of the application for which this federated identity credential should be created
	ApplicationObjectID *string `json:"applicationObjectId,omitempty" tf:"application_object_id,omitempty"`

	// List of audiences that can appear in the external token. This specifies what should be accepted in the aud claim of incoming tokens.
	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	Audiences []*string `json:"audiences,omitempty" tf:"audiences,omitempty"`

	// A UUID used to uniquely identify this federated identity credential.
	// A UUID used to uniquely identify this federated identity credential
	CredentialID *string `json:"credentialId,omitempty" tf:"credential_id,omitempty"`

	// A description for the federated identity credential.
	// A description for the federated identity credential
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
	// A unique display name for the federated identity credential
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type FederatedIdentityCredentialParameters struct {

	// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
	// The object ID of the application for which this federated identity credential should be created
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationObjectID *string `json:"applicationObjectId,omitempty" tf:"application_object_id,omitempty"`

	// Reference to a Application to populate applicationObjectId.
	// +kubebuilder:validation:Optional
	ApplicationObjectIDRef *v1.Reference `json:"applicationObjectIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationObjectId.
	// +kubebuilder:validation:Optional
	ApplicationObjectIDSelector *v1.Selector `json:"applicationObjectIdSelector,omitempty" tf:"-"`

	// List of audiences that can appear in the external token. This specifies what should be accepted in the aud claim of incoming tokens.
	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	// +kubebuilder:validation:Optional
	Audiences []*string `json:"audiences,omitempty" tf:"audiences,omitempty"`

	// A description for the federated identity credential.
	// A description for the federated identity credential
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
	// A unique display name for the federated identity credential
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

// FederatedIdentityCredentialSpec defines the desired state of FederatedIdentityCredential
type FederatedIdentityCredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FederatedIdentityCredentialParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FederatedIdentityCredentialInitParameters `json:"initProvider,omitempty"`
}

// FederatedIdentityCredentialStatus defines the observed state of FederatedIdentityCredential.
type FederatedIdentityCredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FederatedIdentityCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FederatedIdentityCredential is the Schema for the FederatedIdentityCredentials API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type FederatedIdentityCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.audiences) || has(self.initProvider.audiences)",message="audiences is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || has(self.initProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuer) || has(self.initProvider.issuer)",message="issuer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subject) || has(self.initProvider.subject)",message="subject is a required parameter"
	Spec   FederatedIdentityCredentialSpec   `json:"spec"`
	Status FederatedIdentityCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FederatedIdentityCredentialList contains a list of FederatedIdentityCredentials
type FederatedIdentityCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedIdentityCredential `json:"items"`
}

// Repository type metadata.
var (
	FederatedIdentityCredential_Kind             = "FederatedIdentityCredential"
	FederatedIdentityCredential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FederatedIdentityCredential_Kind}.String()
	FederatedIdentityCredential_KindAPIVersion   = FederatedIdentityCredential_Kind + "." + CRDGroupVersion.String()
	FederatedIdentityCredential_GroupVersionKind = CRDGroupVersion.WithKind(FederatedIdentityCredential_Kind)
)

func init() {
	SchemeBuilder.Register(&FederatedIdentityCredential{}, &FederatedIdentityCredentialList{})
}
