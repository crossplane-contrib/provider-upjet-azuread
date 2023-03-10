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

type DynamicMembershipObservation struct {
}

type DynamicMembershipParameters struct {

	// Whether rule processing is "On" (true) or "Paused" (false).
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// The rule that determines membership of this group. For more information, see official documentation on membership rules syntax.
	// Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'
	// +kubebuilder:validation:Required
	Rule *string `json:"rule" tf:"rule,omitempty"`
}

type GroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The SMTP address for the group.
	// The SMTP address for the group
	Mail *string `json:"mail,omitempty" tf:"mail,omitempty"`

	// The object ID of the group.
	// The object ID of the group
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
	// The on-premises FQDN, also called dnsDomainName, synchronized from the on-premises directory when Azure AD Connect is used
	OnpremisesDomainName *string `json:"onpremisesDomainName,omitempty" tf:"onpremises_domain_name,omitempty"`

	// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
	// The on-premises NetBIOS name, synchronized from the on-premises directory when Azure AD Connect is used
	OnpremisesNetbiosName *string `json:"onpremisesNetbiosName,omitempty" tf:"onpremises_netbios_name,omitempty"`

	// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
	// The on-premises SAM account name, synchronized from the on-premises directory when Azure AD Connect is used
	OnpremisesSamAccountName *string `json:"onpremisesSamAccountName,omitempty" tf:"onpremises_sam_account_name,omitempty"`

	// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
	// The on-premises security identifier (SID), synchronized from the on-premises directory when Azure AD Connect is used
	OnpremisesSecurityIdentifier *string `json:"onpremisesSecurityIdentifier,omitempty" tf:"onpremises_security_identifier,omitempty"`

	// Whether this group is synchronised from an on-premises directory (true), no longer synchronised (false), or has never been synchronised (null).
	// Whether this group is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null)
	OnpremisesSyncEnabled *bool `json:"onpremisesSyncEnabled,omitempty" tf:"onpremises_sync_enabled,omitempty"`

	// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
	// The preferred language for a Microsoft 365 group, in ISO 639-1 notation
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// List of email addresses for the group that direct to the same group mailbox.
	// Email addresses for the group that direct to the same group mailbox
	ProxyAddresses []*string `json:"proxyAddresses,omitempty" tf:"proxy_addresses,omitempty"`
}

type GroupParameters struct {

	// The object IDs of administrative units in which the group is a member. If specified, new groups will be created in the scope of the first administrative unit and added to the others. If empty, new groups will be created at the tenant level.
	// The administrative unit IDs in which the group should be. If empty, the group will be created at the tenant level.
	// +kubebuilder:validation:Optional
	AdministrativeUnitIds []*string `json:"administrativeUnitIds,omitempty" tf:"administrative_unit_ids,omitempty"`

	// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be true for security-enabled groups. Changing this forces a new resource to be created.
	// Indicates whether this group can be assigned to an Azure Active Directory role. This property can only be `true` for security-enabled groups.
	// +kubebuilder:validation:Optional
	AssignableToRole *bool `json:"assignableToRole,omitempty" tf:"assignable_to_role,omitempty"`

	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications.
	// +kubebuilder:validation:Optional
	AutoSubscribeNewMembers *bool `json:"autoSubscribeNewMembers,omitempty" tf:"auto_subscribe_new_members,omitempty"`

	// A set of behaviors for a Microsoft 365 group. Possible values are AllowOnlyMembersToPost, HideGroupInOutlook, SubscribeMembersToCalendarEventsDisabled, SubscribeNewGroupMembers and WelcomeEmailDisabled. See official documentation for more details. Changing this forces a new resource to be created.
	// The group behaviours for a Microsoft 365 group
	// +kubebuilder:validation:Optional
	Behaviors []*string `json:"behaviors,omitempty" tf:"behaviors,omitempty"`

	// The description for the group.
	// The description for the group
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name for the group.
	// The display name for the group
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// A dynamic_membership block as documented below. Required when types contains DynamicMembership. Cannot be used with the members property.
	// An optional block to configure dynamic membership for the group. Cannot be used with `members`
	// +kubebuilder:validation:Optional
	DynamicMembership []DynamicMembershipParameters `json:"dynamicMembership,omitempty" tf:"dynamic_membership,omitempty"`

	// Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
	// Indicates whether people external to the organization can send messages to the group.
	// +kubebuilder:validation:Optional
	ExternalSendersAllowed *bool `json:"externalSendersAllowed,omitempty" tf:"external_senders_allowed,omitempty"`

	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.
	// +kubebuilder:validation:Optional
	HideFromAddressLists *bool `json:"hideFromAddressLists,omitempty" tf:"hide_from_address_lists,omitempty"`

	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.
	// +kubebuilder:validation:Optional
	HideFromOutlookClients *bool `json:"hideFromOutlookClients,omitempty" tf:"hide_from_outlook_clients,omitempty"`

	// Whether the group is a mail enabled, with a shared group mailbox. At least one of mail_enabled or security_enabled must be specified. Only Microsoft 365 groups can be mail enabled (see the types property).
	// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. A group can be mail enabled _and_ security enabled
	// +kubebuilder:validation:Optional
	MailEnabled *bool `json:"mailEnabled,omitempty" tf:"mail_enabled,omitempty"`

	// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
	// The mail alias for the group, unique in the organisation
	// +kubebuilder:validation:Optional
	MailNickname *string `json:"mailNickname,omitempty" tf:"mail_nickname,omitempty"`

	// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the dynamic_membership block.
	// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals
	// +kubebuilder:validation:Optional
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// A set of object IDs of principals that will be granted ownership of the group. Supported object types are users or service principals. Groups cannot be created with no owners or have all their owners removed.
	// A set of owners who own this group. Supported object types are Users or Service Principals
	// +kubebuilder:validation:Optional
	Owners []*string `json:"owners,omitempty" tf:"owners,omitempty"`

	// If true, will return an error if an existing group is found with the same name. Defaults to false.
	// If `true`, will return an error if an existing group is found with the same name
	// +kubebuilder:validation:Optional
	PreventDuplicateNames *bool `json:"preventDuplicateNames,omitempty" tf:"prevent_duplicate_names,omitempty"`

	// A set of provisioning options for a Microsoft 365 group. The only supported value is Team. See official documentation for details. Changing this forces a new resource to be created.
	// The group provisioning options for a Microsoft 365 group
	// +kubebuilder:validation:Optional
	ProvisioningOptions []*string `json:"provisioningOptions,omitempty" tf:"provisioning_options,omitempty"`

	// Whether the group is a security group for controlling access to in-app resources. At least one of security_enabled or mail_enabled must be specified. A Microsoft 365 group can be security enabled and mail enabled (see the types property).
	// Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A group can be security enabled _and_ mail enabled
	// +kubebuilder:validation:Optional
	SecurityEnabled *bool `json:"securityEnabled,omitempty" tf:"security_enabled,omitempty"`

	// The colour theme for a Microsoft 365 group. Possible values are Blue, Green, Orange, Pink, Purple, Red or Teal. By default, no theme is set.
	// The colour theme for a Microsoft 365 group
	// +kubebuilder:validation:Optional
	Theme *string `json:"theme,omitempty" tf:"theme,omitempty"`

	// A set of group types to configure for the group. Supported values are DynamicMembership, which denotes a group with dynamic membership, and Unified, which specifies a Microsoft 365 group. Required when mail_enabled is true. Changing this forces a new resource to be created.
	// A set of group types to configure for the group. `Unified` specifies a Microsoft 365 group. Required when `mail_enabled` is true
	// +kubebuilder:validation:Optional
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`

	// The group join policy and group content visibility. Possible values are Private, Public, or Hiddenmembership. Only Microsoft 365 groups can have Hiddenmembership visibility and this value must be set when the group is created. By default, security groups will receive Private visibility and Microsoft 365 groups will receive Public visibility.
	// Specifies the group join policy and group content visibility
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
