//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

// AccessReviewDecisionIdentityClassification provides polymorphic access to related types.
// Call the interface's GetAccessReviewDecisionIdentity() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AccessReviewDecisionIdentity, *AccessReviewDecisionServicePrincipalIdentity, *AccessReviewDecisionUserIdentity
type AccessReviewDecisionIdentityClassification interface {
	// GetAccessReviewDecisionIdentity returns the AccessReviewDecisionIdentity content of the underlying type.
	GetAccessReviewDecisionIdentity() *AccessReviewDecisionIdentity
}

// AccessReviewDecisionInsightPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetAccessReviewDecisionInsightProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AccessReviewDecisionInsightProperties, *AccessReviewDecisionUserSignInInsightProperties
type AccessReviewDecisionInsightPropertiesClassification interface {
	// GetAccessReviewDecisionInsightProperties returns the AccessReviewDecisionInsightProperties content of the underlying type.
	GetAccessReviewDecisionInsightProperties() *AccessReviewDecisionInsightProperties
}

// AlertConfigurationPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetAlertConfigurationProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AlertConfigurationProperties, *AzureRolesAssignedOutsidePimAlertConfigurationProperties, *DuplicateRoleCreatedAlertConfigurationProperties,
// - *TooManyOwnersAssignedToResourceAlertConfigurationProperties, *TooManyPermanentOwnersAssignedToResourceAlertConfigurationProperties
type AlertConfigurationPropertiesClassification interface {
	// GetAlertConfigurationProperties returns the AlertConfigurationProperties content of the underlying type.
	GetAlertConfigurationProperties() *AlertConfigurationProperties
}

// AlertIncidentPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetAlertIncidentProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AlertIncidentProperties, *AzureRolesAssignedOutsidePimAlertIncidentProperties, *DuplicateRoleCreatedAlertIncidentProperties,
// - *TooManyOwnersAssignedToResourceAlertIncidentProperties, *TooManyPermanentOwnersAssignedToResourceAlertIncidentProperties
type AlertIncidentPropertiesClassification interface {
	// GetAlertIncidentProperties returns the AlertIncidentProperties content of the underlying type.
	GetAlertIncidentProperties() *AlertIncidentProperties
}

// RoleManagementPolicyRuleClassification provides polymorphic access to related types.
// Call the interface's GetRoleManagementPolicyRule() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RoleManagementPolicyApprovalRule, *RoleManagementPolicyAuthenticationContextRule, *RoleManagementPolicyEnablementRule,
// - *RoleManagementPolicyExpirationRule, *RoleManagementPolicyNotificationRule, *RoleManagementPolicyRule
type RoleManagementPolicyRuleClassification interface {
	// GetRoleManagementPolicyRule returns the RoleManagementPolicyRule content of the underlying type.
	GetRoleManagementPolicyRule() *RoleManagementPolicyRule
}
