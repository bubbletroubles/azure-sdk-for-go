//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	NameAvailabilityInformation
}

// AccountsClientCreateResponse contains the response from method AccountsClient.Create.
type AccountsClientCreateResponse struct {
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientEnableKeyVaultResponse contains the response from method AccountsClient.EnableKeyVault.
type AccountsClientEnableKeyVaultResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.ListByResourceGroup.
type AccountsClientListByResourceGroupResponse struct {
	AccountListResult
}

// AccountsClientListResponse contains the response from method AccountsClient.List.
type AccountsClientListResponse struct {
	AccountListResult
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	Account
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.CreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.Delete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	FirewallRule
}

// FirewallRulesClientListByAccountResponse contains the response from method FirewallRulesClient.ListByAccount.
type FirewallRulesClientListByAccountResponse struct {
	FirewallRuleListResult
}

// FirewallRulesClientUpdateResponse contains the response from method FirewallRulesClient.Update.
type FirewallRulesClientUpdateResponse struct {
	FirewallRule
}

// LocationsClientGetCapabilityResponse contains the response from method LocationsClient.GetCapability.
type LocationsClientGetCapabilityResponse struct {
	CapabilityInformation
}

// LocationsClientGetUsageResponse contains the response from method LocationsClient.GetUsage.
type LocationsClientGetUsageResponse struct {
	UsageListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// TrustedIDProvidersClientCreateOrUpdateResponse contains the response from method TrustedIDProvidersClient.CreateOrUpdate.
type TrustedIDProvidersClientCreateOrUpdateResponse struct {
	TrustedIDProvider
}

// TrustedIDProvidersClientDeleteResponse contains the response from method TrustedIDProvidersClient.Delete.
type TrustedIDProvidersClientDeleteResponse struct {
	// placeholder for future response values
}

// TrustedIDProvidersClientGetResponse contains the response from method TrustedIDProvidersClient.Get.
type TrustedIDProvidersClientGetResponse struct {
	TrustedIDProvider
}

// TrustedIDProvidersClientListByAccountResponse contains the response from method TrustedIDProvidersClient.ListByAccount.
type TrustedIDProvidersClientListByAccountResponse struct {
	TrustedIDProviderListResult
}

// TrustedIDProvidersClientUpdateResponse contains the response from method TrustedIDProvidersClient.Update.
type TrustedIDProvidersClientUpdateResponse struct {
	TrustedIDProvider
}

// VirtualNetworkRulesClientCreateOrUpdateResponse contains the response from method VirtualNetworkRulesClient.CreateOrUpdate.
type VirtualNetworkRulesClientCreateOrUpdateResponse struct {
	VirtualNetworkRule
}

// VirtualNetworkRulesClientDeleteResponse contains the response from method VirtualNetworkRulesClient.Delete.
type VirtualNetworkRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworkRulesClientGetResponse contains the response from method VirtualNetworkRulesClient.Get.
type VirtualNetworkRulesClientGetResponse struct {
	VirtualNetworkRule
}

// VirtualNetworkRulesClientListByAccountResponse contains the response from method VirtualNetworkRulesClient.ListByAccount.
type VirtualNetworkRulesClientListByAccountResponse struct {
	VirtualNetworkRuleListResult
}

// VirtualNetworkRulesClientUpdateResponse contains the response from method VirtualNetworkRulesClient.Update.
type VirtualNetworkRulesClientUpdateResponse struct {
	VirtualNetworkRule
}