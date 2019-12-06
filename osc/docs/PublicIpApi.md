# \PublicIpApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePublicIp**](PublicIpApi.md#CreatePublicIp) | **Post** /CreatePublicIp | 
[**DeletePublicIp**](PublicIpApi.md#DeletePublicIp) | **Post** /DeletePublicIp | 
[**LinkPublicIp**](PublicIpApi.md#LinkPublicIp) | **Post** /LinkPublicIp | 
[**ReadPublicIps**](PublicIpApi.md#ReadPublicIps) | **Post** /ReadPublicIps | 
[**UnlinkPublicIp**](PublicIpApi.md#UnlinkPublicIp) | **Post** /UnlinkPublicIp | 



## CreatePublicIp

> CreatePublicIpResponse CreatePublicIp(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreatePublicIpOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePublicIpOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPublicIpRequest** | [**optional.Interface of CreatePublicIpRequest**](CreatePublicIpRequest.md)|  | 

### Return type

[**CreatePublicIpResponse**](CreatePublicIpResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicIp

> DeletePublicIpResponse DeletePublicIp(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeletePublicIpOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeletePublicIpOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePublicIpRequest** | [**optional.Interface of DeletePublicIpRequest**](DeletePublicIpRequest.md)|  | 

### Return type

[**DeletePublicIpResponse**](DeletePublicIpResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkPublicIp

> LinkPublicIpResponse LinkPublicIp(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkPublicIpOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkPublicIpOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkPublicIpRequest** | [**optional.Interface of LinkPublicIpRequest**](LinkPublicIpRequest.md)|  | 

### Return type

[**LinkPublicIpResponse**](LinkPublicIpResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadPublicIps

> ReadPublicIpsResponse ReadPublicIps(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadPublicIpsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadPublicIpsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPublicIpsRequest** | [**optional.Interface of ReadPublicIpsRequest**](ReadPublicIpsRequest.md)|  | 

### Return type

[**ReadPublicIpsResponse**](ReadPublicIpsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkPublicIp

> UnlinkPublicIpResponse UnlinkPublicIp(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkPublicIpOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkPublicIpOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkPublicIpRequest** | [**optional.Interface of UnlinkPublicIpRequest**](UnlinkPublicIpRequest.md)|  | 

### Return type

[**UnlinkPublicIpResponse**](UnlinkPublicIpResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

