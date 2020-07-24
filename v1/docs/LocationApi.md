# \LocationApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadLocations**](LocationApi.md#ReadLocations) | **Post** /ReadLocations | 



## ReadLocations

> ReadLocationsResponse ReadLocations(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadLocationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLocationsRequest** | [**optional.Interface of ReadLocationsRequest**](ReadLocationsRequest.md)|  | 

### Return type

[**ReadLocationsResponse**](ReadLocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

