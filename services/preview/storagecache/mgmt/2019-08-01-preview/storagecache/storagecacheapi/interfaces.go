package storagecacheapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/storagecache/mgmt/2019-08-01-preview/storagecache"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result storagecache.APIOperationListResultPage, err error)
	ListComplete(ctx context.Context) (result storagecache.APIOperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*storagecache.OperationsClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	List(ctx context.Context) (result storagecache.ResourceSkusResultPage, err error)
	ListComplete(ctx context.Context) (result storagecache.ResourceSkusResultIterator, err error)
}

var _ SkusClientAPI = (*storagecache.SkusClient)(nil)

// UsageModelsClientAPI contains the set of methods on the UsageModelsClient type.
type UsageModelsClientAPI interface {
	List(ctx context.Context) (result storagecache.UsageModelsResultPage, err error)
	ListComplete(ctx context.Context) (result storagecache.UsageModelsResultIterator, err error)
}

var _ UsageModelsClientAPI = (*storagecache.UsageModelsClient)(nil)

// CachesClientAPI contains the set of methods on the CachesClient type.
type CachesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, cacheName string, cache *storagecache.Cache) (result storagecache.CachesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesDeleteFuture, err error)
	Flush(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesFlushFuture, err error)
	Get(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.Cache, err error)
	List(ctx context.Context) (result storagecache.CachesListResultPage, err error)
	ListComplete(ctx context.Context) (result storagecache.CachesListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storagecache.CachesListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result storagecache.CachesListResultIterator, err error)
	Start(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, cacheName string, cache *storagecache.Cache) (result storagecache.Cache, err error)
	UpgradeFirmware(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesUpgradeFirmwareFuture, err error)
}

var _ CachesClientAPI = (*storagecache.CachesClient)(nil)

// StorageTargetsClientAPI contains the set of methods on the StorageTargetsClient type.
type StorageTargetsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, storagetarget *storagecache.StorageTarget) (result storagecache.StorageTargetsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result storagecache.StorageTargetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result storagecache.StorageTarget, err error)
	ListByCache(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.StorageTargetsResultPage, err error)
	ListByCacheComplete(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.StorageTargetsResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, storagetarget *storagecache.StorageTarget) (result storagecache.StorageTarget, err error)
}

var _ StorageTargetsClientAPI = (*storagecache.StorageTargetsClient)(nil)