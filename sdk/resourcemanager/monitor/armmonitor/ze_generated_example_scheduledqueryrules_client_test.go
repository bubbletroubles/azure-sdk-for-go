//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("b67f7fec-69fc-4974-9099-a26bd6ffeda3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"Rac46PostSwapRG",
		"logalertfoo",
		armmonitor.LogSearchRuleResource{
			Location: to.Ptr("eastus"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.LogSearchRule{
				Description: to.Ptr("log alert description"),
				Action: &armmonitor.AlertingAction{
					ODataType: to.Ptr("Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction"),
					AznsAction: &armmonitor.AzNsActionGroup{
						ActionGroup:          []*string{},
						CustomWebhookPayload: to.Ptr("{}"),
						EmailSubject:         to.Ptr("Email Header"),
					},
					Severity: to.Ptr(armmonitor.AlertSeverityOne),
					Trigger: &armmonitor.TriggerCondition{
						MetricTrigger: &armmonitor.LogMetricTrigger{
							MetricColumn:      to.Ptr("Computer"),
							MetricTriggerType: to.Ptr(armmonitor.MetricTriggerTypeConsecutive),
							Threshold:         to.Ptr[float64](5),
							ThresholdOperator: to.Ptr(armmonitor.ConditionalOperatorGreaterThan),
						},
						Threshold:         to.Ptr[float64](3),
						ThresholdOperator: to.Ptr(armmonitor.ConditionalOperatorGreaterThan),
					},
				},
				Enabled: to.Ptr(armmonitor.EnabledTrue),
				Schedule: &armmonitor.Schedule{
					FrequencyInMinutes:  to.Ptr[int32](15),
					TimeWindowInMinutes: to.Ptr[int32](15),
				},
				Source: &armmonitor.Source{
					DataSourceID: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.OperationalInsights/workspaces/sampleWorkspace"),
					Query:        to.Ptr("Heartbeat | summarize AggregatedValue = count() by bin(TimeGenerated, 5m)"),
					QueryType:    to.Ptr(armmonitor.QueryTypeResultCount),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/getScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("b67f7fec-69fc-4974-9099-a26bd6ffeda3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"Rac46PostSwapRG",
		"logalertfoo",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/patchScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"my-resource-group",
		"logalertfoo",
		armmonitor.LogSearchRuleResourcePatch{
			Properties: &armmonitor.LogSearchRulePatch{
				Enabled: to.Ptr(armmonitor.EnabledTrue),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/deleteScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("b67f7fec-69fc-4974-9099-a26bd6ffeda3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"Rac46PostSwapRG",
		"logalertfoo",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/listScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(&armmonitor.ScheduledQueryRulesClientListBySubscriptionOptions{Filter: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/listScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("gigtest",
		&armmonitor.ScheduledQueryRulesClientListByResourceGroupOptions{Filter: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}