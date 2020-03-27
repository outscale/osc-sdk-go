# \NetAccessPointApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetAccessPoint**](NetAccessPointApi.md#CreateNetAccessPoint) | **Post** /CreateNetAccessPoint | 
[**DeleteNetAccessPoint**](NetAccessPointApi.md#DeleteNetAccessPoint) | **Post** /DeleteNetAccessPoint | 
[**ReadNetAccessPointServices**](NetAccessPointApi.md#ReadNetAccessPointServices) | **Post** /ReadNetAccessPointServices | 
[**ReadNetAccessPoints**](NetAccessPointApi.md#ReadNetAccessPoints) | **Post** /ReadNetAccessPoints | 
[**UpdateNetAccessPoint**](NetAccessPointApi.md#UpdateNetAccessPoint) | **Post** /UpdateNetAccessPoint | 



## CreateNetAccessPoint

> CreateNetAccessPointResponse CreateNetAccessPoint(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNetAccessPointOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNetAccessPointOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetAccessPointRequest** | [**optional.Interface of CreateNetAccessPointRequest**](CreateNetAccessPointRequest.md)|  | 

### Return type

[**CreateNetAccessPointResponse**](CreateNetAccessPointResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetAccessPoint

> DeleteNetAccessPointResponse DeleteNetAccessPoint(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteNetAccessPointOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteNetAccessPointOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNetAccessPointRequest** | [**optional.Interface of DeleteNetAccessPointRequest**](DeleteNetAccessPointRequest.md)|  | 

### Return type

[**DeleteNetAccessPointResponse**](DeleteNetAccessPointResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNetAccessPointServices

> ReadNetAccessPointServicesResponse ReadNetAccessPointServices(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadNetAccessPointServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadNetAccessPointServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNetAccessPointServicesRequest** | [**optional.Interface of ReadNetAccessPointServicesRequest**](ReadNetAccessPointServicesRequest.md)|  | 

### Return type

[**ReadNetAccessPointServicesResponse**](ReadNetAccessPointServicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNetAccessPoints

> ReadNetAccessPointsResponse ReadNetAccessPoints(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadNetAccessPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadNetAccessPointsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNetAccessPointsRequest** | [**optional.Interface of ReadNetAccessPointsRequest**](ReadNetAccessPointsRequest.md)|  | 

### Return type

[**ReadNetAccessPointsResponse**](ReadNetAccessPointsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetAccessPoint

> UpdateNetAccessPointResponse UpdateNetAccessPoint(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateNetAccessPointOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateNetAccessPointOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNetAccessPointRequest** | [**optional.Interface of UpdateNetAccessPointRequest**](UpdateNetAccessPointRequest.md)|  | 

### Return type

[**UpdateNetAccessPointResponse**](UpdateNetAccessPointResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

