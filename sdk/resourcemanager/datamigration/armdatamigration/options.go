//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

// FilesClientCreateOrUpdateOptions contains the optional parameters for the FilesClient.CreateOrUpdate method.
type FilesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// FilesClientDeleteOptions contains the optional parameters for the FilesClient.Delete method.
type FilesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// FilesClientGetOptions contains the optional parameters for the FilesClient.Get method.
type FilesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FilesClientListOptions contains the optional parameters for the FilesClient.NewListPager method.
type FilesClientListOptions struct {
	// placeholder for future optional parameters
}

// FilesClientReadOptions contains the optional parameters for the FilesClient.Read method.
type FilesClientReadOptions struct {
	// placeholder for future optional parameters
}

// FilesClientReadWriteOptions contains the optional parameters for the FilesClient.ReadWrite method.
type FilesClientReadWriteOptions struct {
	// placeholder for future optional parameters
}

// FilesClientUpdateOptions contains the optional parameters for the FilesClient.Update method.
type FilesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientCreateOrUpdateOptions contains the optional parameters for the ProjectsClient.CreateOrUpdate method.
type ProjectsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientDeleteOptions contains the optional parameters for the ProjectsClient.Delete method.
type ProjectsClientDeleteOptions struct {
	// Delete the resource even if it contains running tasks
	DeleteRunningTasks *bool
}

// ProjectsClientGetOptions contains the optional parameters for the ProjectsClient.Get method.
type ProjectsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientListOptions contains the optional parameters for the ProjectsClient.NewListPager method.
type ProjectsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientUpdateOptions contains the optional parameters for the ProjectsClient.Update method.
type ProjectsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ResourceSKUsClientListSKUsOptions contains the optional parameters for the ResourceSKUsClient.NewListSKUsPager method.
type ResourceSKUsClientListSKUsOptions struct {
	// placeholder for future optional parameters
}

// ServiceTasksClientCancelOptions contains the optional parameters for the ServiceTasksClient.Cancel method.
type ServiceTasksClientCancelOptions struct {
	// placeholder for future optional parameters
}

// ServiceTasksClientCreateOrUpdateOptions contains the optional parameters for the ServiceTasksClient.CreateOrUpdate method.
type ServiceTasksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ServiceTasksClientDeleteOptions contains the optional parameters for the ServiceTasksClient.Delete method.
type ServiceTasksClientDeleteOptions struct {
	// Delete the resource even if it contains running tasks
	DeleteRunningTasks *bool
}

// ServiceTasksClientGetOptions contains the optional parameters for the ServiceTasksClient.Get method.
type ServiceTasksClientGetOptions struct {
	// Expand the response
	Expand *string
}

// ServiceTasksClientListOptions contains the optional parameters for the ServiceTasksClient.NewListPager method.
type ServiceTasksClientListOptions struct {
	// Filter tasks by task type
	TaskType *string
}

// ServiceTasksClientUpdateOptions contains the optional parameters for the ServiceTasksClient.Update method.
type ServiceTasksClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServicesClient.BeginCreateOrUpdate method.
type ServicesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientBeginDeleteOptions contains the optional parameters for the ServicesClient.BeginDelete method.
type ServicesClientBeginDeleteOptions struct {
	// Delete the resource even if it contains running tasks
	DeleteRunningTasks *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientBeginStartOptions contains the optional parameters for the ServicesClient.BeginStart method.
type ServicesClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientBeginStopOptions contains the optional parameters for the ServicesClient.BeginStop method.
type ServicesClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientBeginUpdateOptions contains the optional parameters for the ServicesClient.BeginUpdate method.
type ServicesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientCheckChildrenNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckChildrenNameAvailability
// method.
type ServicesClientCheckChildrenNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientCheckNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckNameAvailability
// method.
type ServicesClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientCheckStatusOptions contains the optional parameters for the ServicesClient.CheckStatus method.
type ServicesClientCheckStatusOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
type ServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.NewListByResourceGroupPager
// method.
type ServicesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListOptions contains the optional parameters for the ServicesClient.NewListPager method.
type ServicesClientListOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListSKUsOptions contains the optional parameters for the ServicesClient.NewListSKUsPager method.
type ServicesClientListSKUsOptions struct {
	// placeholder for future optional parameters
}

// TasksClientCancelOptions contains the optional parameters for the TasksClient.Cancel method.
type TasksClientCancelOptions struct {
	// placeholder for future optional parameters
}

// TasksClientCommandOptions contains the optional parameters for the TasksClient.Command method.
type TasksClientCommandOptions struct {
	// placeholder for future optional parameters
}

// TasksClientCreateOrUpdateOptions contains the optional parameters for the TasksClient.CreateOrUpdate method.
type TasksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TasksClientDeleteOptions contains the optional parameters for the TasksClient.Delete method.
type TasksClientDeleteOptions struct {
	// Delete the resource even if it contains running tasks
	DeleteRunningTasks *bool
}

// TasksClientGetOptions contains the optional parameters for the TasksClient.Get method.
type TasksClientGetOptions struct {
	// Expand the response
	Expand *string
}

// TasksClientListOptions contains the optional parameters for the TasksClient.NewListPager method.
type TasksClientListOptions struct {
	// Filter tasks by task type
	TaskType *string
}

// TasksClientUpdateOptions contains the optional parameters for the TasksClient.Update method.
type TasksClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// UsagesClientListOptions contains the optional parameters for the UsagesClient.NewListPager method.
type UsagesClientListOptions struct {
	// placeholder for future optional parameters
}
