//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorder

// ManagementClientBeginCreateAddressOptions contains the optional parameters for the ManagementClient.BeginCreateAddress
// method.
type ManagementClientBeginCreateAddressOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagementClientBeginCreateOrderItemOptions contains the optional parameters for the ManagementClient.BeginCreateOrderItem
// method.
type ManagementClientBeginCreateOrderItemOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagementClientBeginDeleteAddressByNameOptions contains the optional parameters for the ManagementClient.BeginDeleteAddressByName
// method.
type ManagementClientBeginDeleteAddressByNameOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagementClientBeginDeleteOrderItemByNameOptions contains the optional parameters for the ManagementClient.BeginDeleteOrderItemByName
// method.
type ManagementClientBeginDeleteOrderItemByNameOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagementClientBeginReturnOrderItemOptions contains the optional parameters for the ManagementClient.BeginReturnOrderItem
// method.
type ManagementClientBeginReturnOrderItemOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagementClientBeginUpdateAddressOptions contains the optional parameters for the ManagementClient.BeginUpdateAddress
// method.
type ManagementClientBeginUpdateAddressOptions struct {
	// Defines the If-Match condition. The patch will be performed only if the ETag of the job on the server matches this value.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagementClientBeginUpdateOrderItemOptions contains the optional parameters for the ManagementClient.BeginUpdateOrderItem
// method.
type ManagementClientBeginUpdateOrderItemOptions struct {
	// Defines the If-Match condition. The patch will be performed only if the ETag of the order on the server matches this value.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagementClientCancelOrderItemOptions contains the optional parameters for the ManagementClient.CancelOrderItem method.
type ManagementClientCancelOrderItemOptions struct {
	// placeholder for future optional parameters
}

// ManagementClientGetAddressByNameOptions contains the optional parameters for the ManagementClient.GetAddressByName method.
type ManagementClientGetAddressByNameOptions struct {
	// placeholder for future optional parameters
}

// ManagementClientGetOrderByNameOptions contains the optional parameters for the ManagementClient.GetOrderByName method.
type ManagementClientGetOrderByNameOptions struct {
	// placeholder for future optional parameters
}

// ManagementClientGetOrderItemByNameOptions contains the optional parameters for the ManagementClient.GetOrderItemByName
// method.
type ManagementClientGetOrderItemByNameOptions struct {
	// $expand is supported on device details, forward shipping details and reverse shipping details parameters. Each of these
	// can be provided as a comma separated list. Device Details for order item
	// provides details on the devices of the product, Forward and Reverse Shipping details provide forward and reverse shipping
	// details respectively.
	Expand *string
}

// ManagementClientListAddressesAtResourceGroupLevelOptions contains the optional parameters for the ManagementClient.NewListAddressesAtResourceGroupLevelPager
// method.
type ManagementClientListAddressesAtResourceGroupLevelOptions struct {
	// $filter is supported to filter based on shipping address properties. Filter supports only equals operation.
	Filter *string

	// $skipToken is supported on Get list of addresses, which provides the next page in the list of address.
	SkipToken *string
}

// ManagementClientListAddressesAtSubscriptionLevelOptions contains the optional parameters for the ManagementClient.NewListAddressesAtSubscriptionLevelPager
// method.
type ManagementClientListAddressesAtSubscriptionLevelOptions struct {
	// $filter is supported to filter based on shipping address properties. Filter supports only equals operation.
	Filter *string

	// $skipToken is supported on Get list of addresses, which provides the next page in the list of addresses.
	SkipToken *string
}

// ManagementClientListConfigurationsOptions contains the optional parameters for the ManagementClient.NewListConfigurationsPager
// method.
type ManagementClientListConfigurationsOptions struct {
	// $skipToken is supported on list of configurations, which provides the next page in the list of configurations.
	SkipToken *string
}

// ManagementClientListOperationsOptions contains the optional parameters for the ManagementClient.NewListOperationsPager
// method.
type ManagementClientListOperationsOptions struct {
	// placeholder for future optional parameters
}

// ManagementClientListOrderAtResourceGroupLevelOptions contains the optional parameters for the ManagementClient.NewListOrderAtResourceGroupLevelPager
// method.
type ManagementClientListOrderAtResourceGroupLevelOptions struct {
	// $skipToken is supported on Get list of order, which provides the next page in the list of order.
	SkipToken *string
}

// ManagementClientListOrderAtSubscriptionLevelOptions contains the optional parameters for the ManagementClient.NewListOrderAtSubscriptionLevelPager
// method.
type ManagementClientListOrderAtSubscriptionLevelOptions struct {
	// $skipToken is supported on Get list of order, which provides the next page in the list of order.
	SkipToken *string
}

// ManagementClientListOrderItemsAtResourceGroupLevelOptions contains the optional parameters for the ManagementClient.NewListOrderItemsAtResourceGroupLevelPager
// method.
type ManagementClientListOrderItemsAtResourceGroupLevelOptions struct {
	// $expand is supported on device details, forward shipping details and reverse shipping details parameters. Each of these
	// can be provided as a comma separated list. Device Details for order item
	// provides details on the devices of the product, Forward and Reverse Shipping details provide forward and reverse shipping
	// details respectively.
	Expand *string

	// $filter is supported to filter based on order id. Filter supports only equals operation.
	Filter *string

	// $skipToken is supported on Get list of order items, which provides the next page in the list of order items.
	SkipToken *string
}

// ManagementClientListOrderItemsAtSubscriptionLevelOptions contains the optional parameters for the ManagementClient.NewListOrderItemsAtSubscriptionLevelPager
// method.
type ManagementClientListOrderItemsAtSubscriptionLevelOptions struct {
	// $expand is supported on device details, forward shipping details and reverse shipping details parameters. Each of these
	// can be provided as a comma separated list. Device Details for order item
	// provides details on the devices of the product, Forward and Reverse Shipping details provide forward and reverse shipping
	// details respectively.
	Expand *string

	// $filter is supported to filter based on order id. Filter supports only equals operation.
	Filter *string

	// $skipToken is supported on Get list of order items, which provides the next page in the list of order items.
	SkipToken *string
}

// ManagementClientListProductFamiliesMetadataOptions contains the optional parameters for the ManagementClient.NewListProductFamiliesMetadataPager
// method.
type ManagementClientListProductFamiliesMetadataOptions struct {
	// $skipToken is supported on list of product families metadata, which provides the next page in the list of product families
	// metadata.
	SkipToken *string
}

// ManagementClientListProductFamiliesOptions contains the optional parameters for the ManagementClient.NewListProductFamiliesPager
// method.
type ManagementClientListProductFamiliesOptions struct {
	// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
	Expand *string

	// $skipToken is supported on list of product families, which provides the next page in the list of product families.
	SkipToken *string
}
