//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoep

// EnergyServicesClientAddPartitionResponse contains the response from method EnergyServicesClient.BeginAddPartition.
type EnergyServicesClientAddPartitionResponse struct {
	// placeholder for future response values
}

// EnergyServicesClientCreateResponse contains the response from method EnergyServicesClient.BeginCreate.
type EnergyServicesClientCreateResponse struct {
	EnergyService
}

// EnergyServicesClientDeleteResponse contains the response from method EnergyServicesClient.BeginDelete.
type EnergyServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// EnergyServicesClientGetResponse contains the response from method EnergyServicesClient.Get.
type EnergyServicesClientGetResponse struct {
	EnergyService
}

// EnergyServicesClientListByResourceGroupResponse contains the response from method EnergyServicesClient.NewListByResourceGroupPager.
type EnergyServicesClientListByResourceGroupResponse struct {
	// The list of oep resources.
	EnergyServiceList
}

// EnergyServicesClientListBySubscriptionResponse contains the response from method EnergyServicesClient.NewListBySubscriptionPager.
type EnergyServicesClientListBySubscriptionResponse struct {
	// The list of oep resources.
	EnergyServiceList
}

// EnergyServicesClientListPartitionsResponse contains the response from method EnergyServicesClient.ListPartitions.
type EnergyServicesClientListPartitionsResponse struct {
	// List of data partitions.
	DataPartitionsListResult
}

// EnergyServicesClientRemovePartitionResponse contains the response from method EnergyServicesClient.BeginRemovePartition.
type EnergyServicesClientRemovePartitionResponse struct {
	// placeholder for future response values
}

// EnergyServicesClientUpdateResponse contains the response from method EnergyServicesClient.Update.
type EnergyServicesClientUpdateResponse struct {
	EnergyService
}

// LocationsClientCheckNameAvailabilityResponse contains the response from method LocationsClient.CheckNameAvailability.
type LocationsClientCheckNameAvailabilityResponse struct {
	// The check availability result.
	CheckNameAvailabilityResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}
