# \DirectLinkInterfaceApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectLinkInterface**](DirectLinkInterfaceApi.md#CreateDirectLinkInterface) | **Post** /CreateDirectLinkInterface | 
[**DeleteDirectLinkInterface**](DirectLinkInterfaceApi.md#DeleteDirectLinkInterface) | **Post** /DeleteDirectLinkInterface | 
[**ReadDirectLinkInterfaces**](DirectLinkInterfaceApi.md#ReadDirectLinkInterfaces) | **Post** /ReadDirectLinkInterfaces | 



## CreateDirectLinkInterface

> CreateDirectLinkInterfaceResponse CreateDirectLinkInterface(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDirectLinkInterfaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDirectLinkInterfaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDirectLinkInterfaceRequest** | [**optional.Interface of CreateDirectLinkInterfaceRequest**](CreateDirectLinkInterfaceRequest.md)|  | 

### Return type

[**CreateDirectLinkInterfaceResponse**](CreateDirectLinkInterfaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDirectLinkInterface

> DeleteDirectLinkInterfaceResponse DeleteDirectLinkInterface(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteDirectLinkInterfaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDirectLinkInterfaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDirectLinkInterfaceRequest** | [**optional.Interface of DeleteDirectLinkInterfaceRequest**](DeleteDirectLinkInterfaceRequest.md)|  | 

### Return type

[**DeleteDirectLinkInterfaceResponse**](DeleteDirectLinkInterfaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDirectLinkInterfaces

> ReadDirectLinkInterfacesResponse ReadDirectLinkInterfaces(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadDirectLinkInterfacesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadDirectLinkInterfacesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readDirectLinkInterfacesRequest** | [**optional.Interface of ReadDirectLinkInterfacesRequest**](ReadDirectLinkInterfacesRequest.md)|  | 

### Return type

[**ReadDirectLinkInterfacesResponse**](ReadDirectLinkInterfacesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

