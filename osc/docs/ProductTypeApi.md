# \ProductTypeApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadProductTypes**](ProductTypeApi.md#ReadProductTypes) | **Post** /ReadProductTypes | 



## ReadProductTypes

> ReadProductTypesResponse ReadProductTypes(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadProductTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadProductTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readProductTypesRequest** | [**optional.Interface of ReadProductTypesRequest**](ReadProductTypesRequest.md)|  | 

### Return type

[**ReadProductTypesResponse**](ReadProductTypesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

