//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armgraphservices

// AccountsClientCreateAndUpdateResponse contains the response from method AccountsClient.BeginCreateAndUpdate.
type AccountsClientCreateAndUpdateResponse struct {
	// Account details
	AccountResource
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	// Account details
	AccountResource
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	// The list of accounts.
	AccountResourceList
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.NewListBySubscriptionPager.
type AccountsClientListBySubscriptionResponse struct {
	// The list of accounts.
	AccountResourceList
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	// Account details
	AccountResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}
