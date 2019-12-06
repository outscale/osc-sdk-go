# \TaskApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExportTask**](TaskApi.md#DeleteExportTask) | **Post** /DeleteExportTask | 



## DeleteExportTask

> DeleteExportTaskResponse DeleteExportTask(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteExportTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteExportTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteExportTaskRequest** | [**optional.Interface of DeleteExportTaskRequest**](DeleteExportTaskRequest.md)|  | 

### Return type

[**DeleteExportTaskResponse**](DeleteExportTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

