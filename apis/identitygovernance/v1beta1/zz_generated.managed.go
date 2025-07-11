// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PrivilegedAccessGroupAssignmentSchedule.
func (mg *PrivilegedAccessGroupAssignmentSchedule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PrivilegedAccessGroupEligibilitySchedule.
func (mg *PrivilegedAccessGroupEligibilitySchedule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
