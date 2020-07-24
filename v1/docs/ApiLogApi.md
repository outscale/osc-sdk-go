# \ApiLogApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadApiLogs**](ApiLogApi.md#ReadApiLogs) | **Post** /ReadApiLogs | 



## ReadApiLogs

> ReadApiLogsResponse ReadApiLogs(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadApiLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadApiLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readApiLogsRequest** | [**optional.Interface of ReadApiLogsRequest**](ReadApiLogsRequest.md)|  | 

### Return type

[**ReadApiLogsResponse**](ReadApiLogsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

