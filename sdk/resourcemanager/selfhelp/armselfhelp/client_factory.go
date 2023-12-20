//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armselfhelp

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	credential azcore.TokenCredential
	options    *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		credential: credential,
		options:    options.Clone(),
	}, nil
}

// NewCheckNameAvailabilityClient creates a new instance of CheckNameAvailabilityClient.
func (c *ClientFactory) NewCheckNameAvailabilityClient() *CheckNameAvailabilityClient {
	subClient, _ := NewCheckNameAvailabilityClient(c.credential, c.options)
	return subClient
}

// NewDiagnosticsClient creates a new instance of DiagnosticsClient.
func (c *ClientFactory) NewDiagnosticsClient() *DiagnosticsClient {
	subClient, _ := NewDiagnosticsClient(c.credential, c.options)
	return subClient
}

// NewDiscoverySolutionClient creates a new instance of DiscoverySolutionClient.
func (c *ClientFactory) NewDiscoverySolutionClient() *DiscoverySolutionClient {
	subClient, _ := NewDiscoverySolutionClient(c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewSolutionClient creates a new instance of SolutionClient.
func (c *ClientFactory) NewSolutionClient() *SolutionClient {
	subClient, _ := NewSolutionClient(c.credential, c.options)
	return subClient
}

// NewTroubleshootersClient creates a new instance of TroubleshootersClient.
func (c *ClientFactory) NewTroubleshootersClient() *TroubleshootersClient {
	subClient, _ := NewTroubleshootersClient(c.credential, c.options)
	return subClient
}
