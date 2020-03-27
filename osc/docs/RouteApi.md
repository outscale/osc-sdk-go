# \RouteApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoute**](RouteApi.md#CreateRoute) | **Post** /CreateRoute | 
[**DeleteRoute**](RouteApi.md#DeleteRoute) | **Post** /DeleteRoute | 
[**UpdateRoute**](RouteApi.md#UpdateRoute) | **Post** /UpdateRoute | 



## CreateRoute

> CreateRouteResponse CreateRoute(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRouteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRouteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteRequest** | [**optional.Interface of CreateRouteRequest**](CreateRouteRequest.md)|  | 

### Return type

[**CreateRouteResponse**](CreateRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteRouteResponse DeleteRoute(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteRouteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteRouteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRouteRequest** | [**optional.Interface of DeleteRouteRequest**](DeleteRouteRequest.md)|  | 

### Return type

[**DeleteRouteResponse**](DeleteRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoute

> UpdateRouteResponse UpdateRoute(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateRouteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRouteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRouteRequest** | [**optional.Interface of UpdateRouteRequest**](UpdateRouteRequest.md)|  | 

### Return type

[**UpdateRouteResponse**](UpdateRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

