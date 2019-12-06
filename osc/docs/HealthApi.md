# \HealthApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadVmsHealth**](HealthApi.md#ReadVmsHealth) | **Post** /ReadVmsHealth | 



## ReadVmsHealth

> ReadVmsHealthResponse ReadVmsHealth(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVmsHealthOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVmsHealthOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmsHealthRequest** | [**optional.Interface of ReadVmsHealthRequest**](ReadVmsHealthRequest.md)|  | 

### Return type

[**ReadVmsHealthResponse**](ReadVmsHealthResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

