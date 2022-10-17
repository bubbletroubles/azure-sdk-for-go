# Release History

## 2.0.0-beta.2 (2022-06-24)
### Features Added

- New const `DataSourceTypeIngestion`


## 2.0.0-beta.1 (2022-05-24)
### Breaking Changes

- Function `*TablesClient.Update` has been removed
- Struct `TablesClientUpdateOptions` has been removed
- Field `Etag` of struct `Workspace` has been removed

### Features Added

- New const `CreatedByTypeKey`
- New const `ColumnTypeEnumReal`
- New const `SourceEnumMicrosoft`
- New const `CreatedByTypeApplication`
- New const `IdentityTypeKey`
- New const `ProvisioningStateEnumSucceeded`
- New const `ColumnDataTypeHintEnumGUID`
- New const `ColumnTypeEnumString`
- New const `TableTypeEnumMicrosoft`
- New const `TableTypeEnumRestoredLogs`
- New const `TableTypeEnumCustomLog`
- New const `ColumnTypeEnumInt`
- New const `ColumnTypeEnumGUID`
- New const `TableTypeEnumSearchResults`
- New const `ProvisioningStateEnumInProgress`
- New const `IdentityTypeApplication`
- New const `TableSubTypeEnumDataCollectionRuleBased`
- New const `ColumnDataTypeHintEnumIP`
- New const `TablePlanEnumBasic`
- New const `ColumnTypeEnumLong`
- New const `ColumnTypeEnumDateTime`
- New const `ColumnTypeEnumBoolean`
- New const `CreatedByTypeManagedIdentity`
- New const `ProvisioningStateEnumUpdating`
- New const `IdentityTypeManagedIdentity`
- New const `ColumnDataTypeHintEnumArmPath`
- New const `SourceEnumCustomer`
- New const `CreatedByTypeUser`
- New const `TableSubTypeEnumClassic`
- New const `TableSubTypeEnumAny`
- New const `IdentityTypeUser`
- New const `ColumnDataTypeHintEnumURI`
- New const `ColumnTypeEnumDynamic`
- New const `TablePlanEnumAnalytics`
- New function `TagsResource.MarshalJSON() ([]byte, error)`
- New function `LogAnalyticsQueryPackQueryPropertiesRelated.MarshalJSON() ([]byte, error)`
- New function `*LogAnalyticsQueryPackProperties.UnmarshalJSON([]byte) error`
- New function `PossibleCreatedByTypeValues() []CreatedByType`
- New function `PossibleSourceEnumValues() []SourceEnum`
- New function `*SystemData.UnmarshalJSON([]byte) error`
- New function `*TablesClient.BeginUpdate(context.Context, string, string, string, Table, *TablesClientBeginUpdateOptions) (*runtime.Poller[TablesClientUpdateResponse], error)`
- New function `LogAnalyticsQueryPackQuerySearchPropertiesRelated.MarshalJSON() ([]byte, error)`
- New function `LogAnalyticsQueryPackProperties.MarshalJSON() ([]byte, error)`
- New function `LogAnalyticsQueryPackQuerySearchProperties.MarshalJSON() ([]byte, error)`
- New function `SystemDataAutoGenerated.MarshalJSON() ([]byte, error)`
- New function `PossibleProvisioningStateEnumValues() []ProvisioningStateEnum`
- New function `PossibleColumnDataTypeHintEnumValues() []ColumnDataTypeHintEnum`
- New function `PossibleTableSubTypeEnumValues() []TableSubTypeEnum`
- New function `Schema.MarshalJSON() ([]byte, error)`
- New function `*SearchResults.UnmarshalJSON([]byte) error`
- New function `*LogAnalyticsQueryPackQueryProperties.UnmarshalJSON([]byte) error`
- New function `SystemData.MarshalJSON() ([]byte, error)`
- New function `PossibleTableTypeEnumValues() []TableTypeEnum`
- New function `PossibleColumnTypeEnumValues() []ColumnTypeEnum`
- New function `QueryPacksResource.MarshalJSON() ([]byte, error)`
- New function `*SystemDataAutoGenerated.UnmarshalJSON([]byte) error`
- New function `*TablesClient.BeginCreateOrUpdate(context.Context, string, string, string, Table, *TablesClientBeginCreateOrUpdateOptions) (*runtime.Poller[TablesClientCreateOrUpdateResponse], error)`
- New function `RestoredLogs.MarshalJSON() ([]byte, error)`
- New function `LogAnalyticsQueryPackQueryProperties.MarshalJSON() ([]byte, error)`
- New function `*RestoredLogs.UnmarshalJSON([]byte) error`
- New function `LogAnalyticsQueryPackQuery.MarshalJSON() ([]byte, error)`
- New function `SearchResults.MarshalJSON() ([]byte, error)`
- New function `*TablesClient.BeginDelete(context.Context, string, string, string, *TablesClientBeginDeleteOptions) (*runtime.Poller[TablesClientDeleteResponse], error)`
- New function `PossibleTablePlanEnumValues() []TablePlanEnum`
- New function `*TablesClient.Migrate(context.Context, string, string, string, *TablesClientMigrateOptions) (TablesClientMigrateResponse, error)`
- New function `LogAnalyticsQueryPack.MarshalJSON() ([]byte, error)`
- New struct `AzureResourceProperties`
- New struct `Column`
- New struct `LogAnalyticsQueryPack`
- New struct `LogAnalyticsQueryPackListResult`
- New struct `LogAnalyticsQueryPackProperties`
- New struct `LogAnalyticsQueryPackQuery`
- New struct `LogAnalyticsQueryPackQueryListResult`
- New struct `LogAnalyticsQueryPackQueryProperties`
- New struct `LogAnalyticsQueryPackQueryPropertiesRelated`
- New struct `LogAnalyticsQueryPackQuerySearchProperties`
- New struct `LogAnalyticsQueryPackQuerySearchPropertiesRelated`
- New struct `QueriesClientDeleteOptions`
- New struct `QueriesClientDeleteResponse`
- New struct `QueriesClientGetOptions`
- New struct `QueriesClientGetResponse`
- New struct `QueriesClientListOptions`
- New struct `QueriesClientListResponse`
- New struct `QueriesClientPutOptions`
- New struct `QueriesClientPutResponse`
- New struct `QueriesClientSearchOptions`
- New struct `QueriesClientSearchResponse`
- New struct `QueriesClientUpdateOptions`
- New struct `QueriesClientUpdateResponse`
- New struct `QueryPacksClientCreateOrUpdateOptions`
- New struct `QueryPacksClientCreateOrUpdateResponse`
- New struct `QueryPacksClientDeleteOptions`
- New struct `QueryPacksClientDeleteResponse`
- New struct `QueryPacksClientGetOptions`
- New struct `QueryPacksClientGetResponse`
- New struct `QueryPacksClientListByResourceGroupOptions`
- New struct `QueryPacksClientListByResourceGroupResponse`
- New struct `QueryPacksClientListOptions`
- New struct `QueryPacksClientListResponse`
- New struct `QueryPacksClientUpdateTagsOptions`
- New struct `QueryPacksClientUpdateTagsResponse`
- New struct `QueryPacksResource`
- New struct `RestoredLogs`
- New struct `ResultStatistics`
- New struct `Schema`
- New struct `SearchResults`
- New struct `SystemData`
- New struct `SystemDataAutoGenerated`
- New struct `TablesClientBeginCreateOrUpdateOptions`
- New struct `TablesClientBeginDeleteOptions`
- New struct `TablesClientBeginUpdateOptions`
- New struct `TablesClientCreateOrUpdateResponse`
- New struct `TablesClientDeleteResponse`
- New struct `TablesClientMigrateOptions`
- New struct `TablesClientMigrateResponse`
- New struct `TagsResource`
- New field `SystemData` in struct `Workspace`
- New field `ETag` in struct `Workspace`
- New field `DefaultDataCollectionRuleResourceID` in struct `WorkspaceProperties`
- New field `Schema` in struct `TableProperties`
- New field `SearchResults` in struct `TableProperties`
- New field `LastPlanModifiedDate` in struct `TableProperties`
- New field `ProvisioningState` in struct `TableProperties`
- New field `ResultStatistics` in struct `TableProperties`
- New field `TotalRetentionInDays` in struct `TableProperties`
- New field `ArchiveRetentionInDays` in struct `TableProperties`
- New field `Plan` in struct `TableProperties`
- New field `RestoredLogs` in struct `TableProperties`
- New field `SystemData` in struct `Table`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).