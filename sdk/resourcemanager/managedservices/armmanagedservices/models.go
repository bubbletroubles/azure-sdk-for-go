//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

import "time"

// Authorization - The Azure Active Directory principal identifier and Azure built-in role that describes the access the principal
// will receive on the delegated resource in the managed tenant.
type Authorization struct {
	// REQUIRED; The identifier of the Azure Active Directory principal.
	PrincipalID *string

	// REQUIRED; The identifier of the Azure built-in role that defines the permissions that the Azure Active Directory principal
	// will have on the projected scope.
	RoleDefinitionID *string

	// The delegatedRoleDefinitionIds field is required when the roleDefinitionId refers to the User Access Administrator Role.
	// It is the list of role definition ids which define all the permissions that the
	// user in the authorization can assign to other principals.
	DelegatedRoleDefinitionIDs []*string

	// The display name of the Azure Active Directory principal.
	PrincipalIDDisplayName *string
}

// EligibleApprover - Defines the Azure Active Directory principal that can approve any just-in-time access requests by the
// principal defined in the EligibleAuthorization.
type EligibleApprover struct {
	// REQUIRED; The identifier of the Azure Active Directory principal.
	PrincipalID *string

	// The display name of the Azure Active Directory principal.
	PrincipalIDDisplayName *string
}

// EligibleAuthorization - The Azure Active Directory principal identifier, Azure built-in role, and just-in-time access policy
// that describes the just-in-time access the principal will receive on the delegated resource in the
// managed tenant.
type EligibleAuthorization struct {
	// REQUIRED; The identifier of the Azure Active Directory principal.
	PrincipalID *string

	// REQUIRED; The identifier of the Azure built-in role that defines the permissions that the Azure Active Directory principal
	// will have on the projected scope.
	RoleDefinitionID *string

	// The just-in-time access policy setting.
	JustInTimeAccessPolicy *JustInTimeAccessPolicy

	// The display name of the Azure Active Directory principal.
	PrincipalIDDisplayName *string
}

// ErrorDefinition - The error response indicating why the incoming request wasn’t able to be processed
type ErrorDefinition struct {
	// REQUIRED; The error code.
	Code *string

	// REQUIRED; The error message indicating why the operation failed.
	Message *string

	// The internal error details.
	Details []*ErrorDefinition
}

// ErrorResponse - Error response.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition
}

// JustInTimeAccessPolicy - Just-in-time access policy setting.
type JustInTimeAccessPolicy struct {
	// REQUIRED; The multi-factor authorization provider to be used for just-in-time access requests.
	MultiFactorAuthProvider *MultiFactorAuthProvider

	// The list of managedByTenant approvers for the eligible authorization.
	ManagedByTenantApprovers []*EligibleApprover

	// The maximum access duration in ISO 8601 format for just-in-time access requests.
	MaximumActivationDuration *string
}

type MarketplaceRegistrationDefinition struct {
	// The details for the Managed Services offer’s plan in Azure Marketplace.
	Plan *Plan

	// The properties of the marketplace registration definition.
	Properties *MarketplaceRegistrationDefinitionProperties

	// READ-ONLY; The fully qualified path of the marketplace registration definition.
	ID *string

	// READ-ONLY; The name of the marketplace registration definition.
	Name *string

	// READ-ONLY; The type of the Azure resource (Microsoft.ManagedServices/marketplaceRegistrationDefinitions).
	Type *string
}

// MarketplaceRegistrationDefinitionList - The list of marketplace registration definitions.
type MarketplaceRegistrationDefinitionList struct {
	// READ-ONLY; The link to the next page of marketplace registration definitions.
	NextLink *string

	// READ-ONLY; The list of marketplace registration definitions.
	Value []*MarketplaceRegistrationDefinition
}

// MarketplaceRegistrationDefinitionProperties - The properties of the marketplace registration definition.
type MarketplaceRegistrationDefinitionProperties struct {
	// REQUIRED; The collection of authorization objects describing the access Azure Active Directory principals in the managedBy
	// tenant will receive on the delegated resource in the managed tenant.
	Authorizations []*Authorization

	// REQUIRED; The identifier of the managedBy tenant.
	ManagedByTenantID *string

	// The collection of eligible authorization objects describing the just-in-time access Azure Active Directory principals in
	// the managedBy tenant will receive on the delegated resource in the managed
	// tenant.
	EligibleAuthorizations []*EligibleAuthorization

	// The marketplace offer display name.
	OfferDisplayName *string

	// The marketplace plan display name.
	PlanDisplayName *string

	// The marketplace publisher display name.
	PublisherDisplayName *string
}

// Operation - The object that describes a single Microsoft.ManagedServices operation.
type Operation struct {
	// READ-ONLY; The object that represents the operation.
	Display *OperationDisplay

	// READ-ONLY; The operation name with the format: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// The description of the operation.
	Description *string

	// The operation type.
	Operation *string

	// The service provider.
	Provider *string

	// The resource on which the operation is performed.
	Resource *string
}

// OperationList - The list of the operations.
type OperationList struct {
	// READ-ONLY; The list of Microsoft.ManagedServices operations.
	Value []*Operation
}

// Plan - The details for the Managed Services offer’s plan in Azure Marketplace.
type Plan struct {
	// REQUIRED; Azure Marketplace plan name.
	Name *string

	// REQUIRED; Azure Marketplace product code.
	Product *string

	// REQUIRED; Azure Marketplace publisher ID.
	Publisher *string

	// REQUIRED; Azure Marketplace plan's version.
	Version *string
}

// RegistrationAssignment - The registration assignment.
type RegistrationAssignment struct {
	// The properties of a registration assignment.
	Properties *RegistrationAssignmentProperties

	// READ-ONLY; The fully qualified path of the registration assignment.
	ID *string

	// READ-ONLY; The name of the registration assignment.
	Name *string

	// READ-ONLY; The metadata for the registration assignment resource.
	SystemData *SystemData

	// READ-ONLY; The type of the Azure resource (Microsoft.ManagedServices/registrationAssignments).
	Type *string
}

// RegistrationAssignmentList - The list of registration assignments.
type RegistrationAssignmentList struct {
	// READ-ONLY; The link to the next page of registration assignments.
	NextLink *string

	// READ-ONLY; The list of registration assignments.
	Value []*RegistrationAssignment
}

// RegistrationAssignmentProperties - The properties of the registration assignment.
type RegistrationAssignmentProperties struct {
	// REQUIRED; The fully qualified path of the registration definition.
	RegistrationDefinitionID *string

	// READ-ONLY; The current provisioning state of the registration assignment.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The registration definition associated with the registration assignment.
	RegistrationDefinition *RegistrationAssignmentPropertiesRegistrationDefinition
}

// RegistrationAssignmentPropertiesRegistrationDefinition - The registration definition associated with the registration assignment.
type RegistrationAssignmentPropertiesRegistrationDefinition struct {
	// The details for the Managed Services offer’s plan in Azure Marketplace.
	Plan *Plan

	// The properties of the registration definition associated with the registration assignment.
	Properties *RegistrationAssignmentPropertiesRegistrationDefinitionProperties

	// READ-ONLY; The fully qualified path of the registration definition.
	ID *string

	// READ-ONLY; The name of the registration definition.
	Name *string

	// READ-ONLY; The metadata for the registration definition resource.
	SystemData *SystemData

	// READ-ONLY; The type of the Azure resource (Microsoft.ManagedServices/registrationDefinitions).
	Type *string
}

// RegistrationAssignmentPropertiesRegistrationDefinitionProperties - The properties of the registration definition associated
// with the registration assignment.
type RegistrationAssignmentPropertiesRegistrationDefinitionProperties struct {
	// The collection of authorization objects describing the access Azure Active Directory principals in the managedBy tenant
	// will receive on the delegated resource in the managed tenant.
	Authorizations []*Authorization

	// The description of the registration definition.
	Description *string

	// The collection of eligible authorization objects describing the just-in-time access Azure Active Directory principals in
	// the managedBy tenant will receive on the delegated resource in the managed
	// tenant.
	EligibleAuthorizations []*EligibleAuthorization

	// The identifier of the managedBy tenant.
	ManagedByTenantID *string

	// The name of the managedBy tenant.
	ManagedByTenantName *string

	// The identifier of the managed tenant.
	ManageeTenantID *string

	// The name of the managed tenant.
	ManageeTenantName *string

	// The current provisioning state of the registration definition.
	ProvisioningState *ProvisioningState

	// The name of the registration definition.
	RegistrationDefinitionName *string
}

// RegistrationDefinition - The registration definition.
type RegistrationDefinition struct {
	// The details for the Managed Services offer’s plan in Azure Marketplace.
	Plan *Plan

	// The properties of a registration definition.
	Properties *RegistrationDefinitionProperties

	// READ-ONLY; The fully qualified path of the registration definition.
	ID *string

	// READ-ONLY; The name of the registration definition.
	Name *string

	// READ-ONLY; The metadata for the registration assignment resource.
	SystemData *SystemData

	// READ-ONLY; The type of the Azure resource (Microsoft.ManagedServices/registrationDefinitions).
	Type *string
}

// RegistrationDefinitionList - The list of registration definitions.
type RegistrationDefinitionList struct {
	// READ-ONLY; The link to the next page of registration definitions.
	NextLink *string

	// READ-ONLY; The list of registration definitions.
	Value []*RegistrationDefinition
}

// RegistrationDefinitionProperties - The properties of a registration definition.
type RegistrationDefinitionProperties struct {
	// REQUIRED; The collection of authorization objects describing the access Azure Active Directory principals in the managedBy
	// tenant will receive on the delegated resource in the managed tenant.
	Authorizations []*Authorization

	// REQUIRED; The identifier of the managedBy tenant.
	ManagedByTenantID *string

	// The description of the registration definition.
	Description *string

	// The collection of eligible authorization objects describing the just-in-time access Azure Active Directory principals in
	// the managedBy tenant will receive on the delegated resource in the managed
	// tenant.
	EligibleAuthorizations []*EligibleAuthorization

	// The name of the registration definition.
	RegistrationDefinitionName *string

	// READ-ONLY; The name of the managedBy tenant.
	ManagedByTenantName *string

	// READ-ONLY; The identifier of the managed tenant.
	ManageeTenantID *string

	// READ-ONLY; The name of the managed tenant.
	ManageeTenantName *string

	// READ-ONLY; The current provisioning state of the registration definition.
	ProvisioningState *ProvisioningState
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}
