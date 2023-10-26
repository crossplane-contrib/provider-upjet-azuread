/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// invitations
	//
	// No import documented
	"azuread_invitation": config.IdentifierFromProvider,

	// applications
	//
	// azuread_application can be imported using their object ID
	"azuread_application": config.IdentifierFromProvider,
	// azuread_application_certificate can be imported using the object ID of the associated application and the key ID of the certificate credential:
	// {ObjectId}/certificate/{CertificateKeyId}
	// 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
	"azuread_application_certificate": config.IdentifierFromProvider,
	// ***"azuread_application_certificate": config.TemplatedStringAsIdentifier("", "{{ .parameters.application_object_id }}/certificate/{{ .parameters.key_id }}"),
	// azuread_application_federated_identity_credential can be imported using the object ID of the associated application and the ID of the federated identity credential:
	// {ObjectId}/federatedIdentityCredential/{CredentialId}
	// 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111
	"azuread_application_federated_identity_credential": config.IdentifierFromProvider,
	// No import documented
	"azuread_application_password": config.IdentifierFromProvider,
	// azuread_application_pre_authorized imported using the object ID of the authorizing application and the application ID of the application being authorized:
	// {ObjectId}/preAuthorizedApplication/{ApplicationId}
	// 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111
	"azuread_application_pre_authorized": config.IdentifierFromProvider,

	// groups
	//
	// azuread_group can be imported using their object ID
	"azuread_group": config.IdentifierFromProvider,
	// azuread_group_member can be imported using the object ID of the group and the object ID of the member:
	// {GroupObjectID}/member/{MemberObjectID}
	// 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
	"azuread_group_member": config.IdentifierFromProvider,

	// users
	//
	// azuread_user can be imported using their object ID
	"azuread_user": config.IdentifierFromProvider,

	// serviceprincipals
	//
	// azuread_service_principal can be imported using their object ID
	"azuread_service_principal": config.IdentifierFromProvider,
	// azuread_service_principal_claims_mapping_policy_assignment can be imported using the id
	"azuread_service_principal_claims_mapping_policy_assignment": config.IdentifierFromProvider,
	// azuread_service_principal_certificate can be imported using the object ID of the associated
	// service principal and the key ID of the certificate credential
	"azuread_service_principal_certificate": config.IdentifierFromProvider,
	// No import documented
	"azuread_service_principal_password": config.IdentifierFromProvider,
	// azuread_service_principal_token_signing_certificate can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential
	// {ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}
	// 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111
	// TODO: For now API is not normalized but further external name normalization possible.
	"azuread_service_principal_token_signing_certificate": config.IdentifierFromProvider,

	// policies
	//
	// azuread_claims_mapping_policy can be imported using the id
	"azuread_claims_mapping_policy": config.IdentifierFromProvider,

	// administrativeunits
	//
	// azuread_administrative_unit can be imported using their object ID
	"azuread_administrative_unit": config.IdentifierFromProvider,
	// azuread_administrative_unit_member can be imported using the object ID of the administrative unit and the object ID of the member:
	// {AdministrativeUnitObjectID}/member/{MemberObjectID}
	// 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
	"azuread_administrative_unit_member": config.IdentifierFromProvider,

	// synchronization
	//
	// azuread_administrative_unit can be imported using the id:
	// {servicePrincipalId}/job/{jobId}
	// 00000000-0000-0000-0000-000000000000/job/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5
	"azuread_synchronization_job": config.IdentifierFromProvider,
	// No import documented
	"azuread_synchronization_secret": config.IdentifierFromProvider,

	// conditionalaccess
	//
	// azuread_named_location can be imported using the id
	"azuread_named_location": config.IdentifierFromProvider,
	// azuread_conditional_access_policy can be imported using the id
	"azuread_conditional_access_policy": config.IdentifierFromProvider,

	// directoryroles
	//
	// No import documented
	"azuread_custom_directory_role": config.IdentifierFromProvider,
	// No import documented
	"azuread_directory_role": config.IdentifierFromProvider,
	// azuread_directory_role_assignment can be imported using the ID of the assignment
	"azuread_directory_role_assignment": config.IdentifierFromProvider,

	// app
	//
	// azuread_app_role_assignment can be can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note: not the ID of the app role):
	// {ResourcePrincipalID}/appRoleAssignment/{AppRoleAssignmentID}
	// 00000000-0000-0000-0000-000000000000/appRoleAssignment/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
	"azuread_app_role_assignment": config.IdentifierFromProvider,

	// serviceprincipaldelegated
	//
	// azuread_service_principal_delegated_permission_grant can be imported using their ID
	"azuread_service_principal_delegated_permission_grant": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
			r.Version = "v1beta1"
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
