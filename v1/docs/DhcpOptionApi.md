# \DhcpOptionApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDhcpOptions**](DhcpOptionApi.md#CreateDhcpOptions) | **Post** /CreateDhcpOptions | 
[**DeleteDhcpOptions**](DhcpOptionApi.md#DeleteDhcpOptions) | **Post** /DeleteDhcpOptions | 
[**ReadDhcpOptions**](DhcpOptionApi.md#ReadDhcpOptions) | **Post** /ReadDhcpOptions | 



## CreateDhcpOptions

> CreateDhcpOptionsResponse CreateDhcpOptions(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDhcpOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDhcpOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDhcpOptionsRequest** | [**optional.Interface of CreateDhcpOptionsRequest**](CreateDhcpOptionsRequest.md)|  | 

### Return type

[**CreateDhcpOptionsResponse**](CreateDhcpOptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDhcpOptions

> DeleteDhcpOptionsResponse DeleteDhcpOptions(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteDhcpOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDhcpOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDhcpOptionsRequest** | [**optional.Interface of DeleteDhcpOptionsRequest**](DeleteDhcpOptionsRequest.md)|  | 

### Return type

[**DeleteDhcpOptionsResponse**](DeleteDhcpOptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDhcpOptions

> ReadDhcpOptionsResponse ReadDhcpOptions(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadDhcpOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadDhcpOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readDhcpOptionsRequest** | [**optional.Interface of ReadDhcpOptionsRequest**](ReadDhcpOptionsRequest.md)|  | 

### Return type

[**ReadDhcpOptionsResponse**](ReadDhcpOptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

