// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/upbound/provider-azuread/apis/administrativeunits/v1beta1"
	v1beta1app "github.com/upbound/provider-azuread/apis/app/v1beta1"
	v1beta1applications "github.com/upbound/provider-azuread/apis/applications/v1beta1"
	v1beta2 "github.com/upbound/provider-azuread/apis/applications/v1beta2"
	v1beta1conditionalaccess "github.com/upbound/provider-azuread/apis/conditionalaccess/v1beta1"
	v1beta2conditionalaccess "github.com/upbound/provider-azuread/apis/conditionalaccess/v1beta2"
	v1beta1directoryroles "github.com/upbound/provider-azuread/apis/directoryroles/v1beta1"
	v1beta1groups "github.com/upbound/provider-azuread/apis/groups/v1beta1"
	v1beta2groups "github.com/upbound/provider-azuread/apis/groups/v1beta2"
	v1beta1invitations "github.com/upbound/provider-azuread/apis/invitations/v1beta1"
	v1beta2invitations "github.com/upbound/provider-azuread/apis/invitations/v1beta2"
	v1beta1policies "github.com/upbound/provider-azuread/apis/policies/v1beta1"
	v1beta1privileged "github.com/upbound/provider-azuread/apis/privileged/v1beta1"
	v1beta1serviceprincipaldelegated "github.com/upbound/provider-azuread/apis/serviceprincipaldelegated/v1beta1"
	v1beta1serviceprincipals "github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1"
	v1beta2serviceprincipals "github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta2"
	v1beta1synchronization "github.com/upbound/provider-azuread/apis/synchronization/v1beta1"
	v1beta1users "github.com/upbound/provider-azuread/apis/users/v1beta1"
	v1alpha1 "github.com/upbound/provider-azuread/apis/v1alpha1"
	v1beta1apis "github.com/upbound/provider-azuread/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1app.SchemeBuilder.AddToScheme,
		v1beta1applications.SchemeBuilder.AddToScheme,
		v1beta2.SchemeBuilder.AddToScheme,
		v1beta1conditionalaccess.SchemeBuilder.AddToScheme,
		v1beta2conditionalaccess.SchemeBuilder.AddToScheme,
		v1beta1directoryroles.SchemeBuilder.AddToScheme,
		v1beta1groups.SchemeBuilder.AddToScheme,
		v1beta2groups.SchemeBuilder.AddToScheme,
		v1beta1invitations.SchemeBuilder.AddToScheme,
		v1beta2invitations.SchemeBuilder.AddToScheme,
		v1beta1policies.SchemeBuilder.AddToScheme,
		v1beta1privileged.SchemeBuilder.AddToScheme,
		v1beta1serviceprincipaldelegated.SchemeBuilder.AddToScheme,
		v1beta1serviceprincipals.SchemeBuilder.AddToScheme,
		v1beta2serviceprincipals.SchemeBuilder.AddToScheme,
		v1beta1synchronization.SchemeBuilder.AddToScheme,
		v1beta1users.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
