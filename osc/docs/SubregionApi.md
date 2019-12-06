# \SubregionApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadSubregions**](SubregionApi.md#ReadSubregions) | **Post** /ReadSubregions | 



## ReadSubregions

> ReadSubregionsResponse ReadSubregions(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadSubregionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadSubregionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSubregionsRequest** | [**optional.Interface of ReadSubregionsRequest**](ReadSubregionsRequest.md)|  | 

### Return type

[**ReadSubregionsResponse**](ReadSubregionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

