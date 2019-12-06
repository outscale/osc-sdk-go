# \SnapshotApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotApi.md#CreateSnapshot) | **Post** /CreateSnapshot | 
[**CreateSnapshotExportTask**](SnapshotApi.md#CreateSnapshotExportTask) | **Post** /CreateSnapshotExportTask | 
[**DeleteSnapshot**](SnapshotApi.md#DeleteSnapshot) | **Post** /DeleteSnapshot | 
[**ReadSnapshotExportTasks**](SnapshotApi.md#ReadSnapshotExportTasks) | **Post** /ReadSnapshotExportTasks | 
[**ReadSnapshots**](SnapshotApi.md#ReadSnapshots) | **Post** /ReadSnapshots | 
[**UpdateSnapshot**](SnapshotApi.md#UpdateSnapshot) | **Post** /UpdateSnapshot | 



## CreateSnapshot

> CreateSnapshotResponse CreateSnapshot(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSnapshotOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotRequest** | [**optional.Interface of CreateSnapshotRequest**](CreateSnapshotRequest.md)|  | 

### Return type

[**CreateSnapshotResponse**](CreateSnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshotExportTask

> CreateSnapshotExportTaskResponse CreateSnapshotExportTask(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSnapshotExportTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSnapshotExportTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotExportTaskRequest** | [**optional.Interface of CreateSnapshotExportTaskRequest**](CreateSnapshotExportTaskRequest.md)|  | 

### Return type

[**CreateSnapshotExportTaskResponse**](CreateSnapshotExportTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshotResponse DeleteSnapshot(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSnapshotOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSnapshotRequest** | [**optional.Interface of DeleteSnapshotRequest**](DeleteSnapshotRequest.md)|  | 

### Return type

[**DeleteSnapshotResponse**](DeleteSnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSnapshotExportTasks

> ReadSnapshotExportTasksResponse ReadSnapshotExportTasks(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadSnapshotExportTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadSnapshotExportTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSnapshotExportTasksRequest** | [**optional.Interface of ReadSnapshotExportTasksRequest**](ReadSnapshotExportTasksRequest.md)|  | 

### Return type

[**ReadSnapshotExportTasksResponse**](ReadSnapshotExportTasksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSnapshots

> ReadSnapshotsResponse ReadSnapshots(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadSnapshotsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadSnapshotsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSnapshotsRequest** | [**optional.Interface of ReadSnapshotsRequest**](ReadSnapshotsRequest.md)|  | 

### Return type

[**ReadSnapshotsResponse**](ReadSnapshotsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> UpdateSnapshotResponse UpdateSnapshot(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSnapshotOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSnapshotRequest** | [**optional.Interface of UpdateSnapshotRequest**](UpdateSnapshotRequest.md)|  | 

### Return type

[**UpdateSnapshotResponse**](UpdateSnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

