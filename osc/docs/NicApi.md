# \NicApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNic**](NicApi.md#CreateNic) | **Post** /CreateNic | 
[**DeleteNic**](NicApi.md#DeleteNic) | **Post** /DeleteNic | 
[**LinkNic**](NicApi.md#LinkNic) | **Post** /LinkNic | 
[**LinkPrivateIps**](NicApi.md#LinkPrivateIps) | **Post** /LinkPrivateIps | 
[**ReadNics**](NicApi.md#ReadNics) | **Post** /ReadNics | 
[**UnlinkNic**](NicApi.md#UnlinkNic) | **Post** /UnlinkNic | 
[**UnlinkPrivateIps**](NicApi.md#UnlinkPrivateIps) | **Post** /UnlinkPrivateIps | 
[**UpdateNic**](NicApi.md#UpdateNic) | **Post** /UpdateNic | 



## CreateNic

> CreateNicResponse CreateNic(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNicRequest** | [**optional.Interface of CreateNicRequest**](CreateNicRequest.md)|  | 

### Return type

[**CreateNicResponse**](CreateNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNic

> DeleteNicResponse DeleteNic(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteNicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteNicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNicRequest** | [**optional.Interface of DeleteNicRequest**](DeleteNicRequest.md)|  | 

### Return type

[**DeleteNicResponse**](DeleteNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkNic

> LinkNicResponse LinkNic(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkNicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkNicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkNicRequest** | [**optional.Interface of LinkNicRequest**](LinkNicRequest.md)|  | 

### Return type

[**LinkNicResponse**](LinkNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkPrivateIps

> LinkPrivateIpsResponse LinkPrivateIps(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkPrivateIpsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkPrivateIpsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkPrivateIpsRequest** | [**optional.Interface of LinkPrivateIpsRequest**](LinkPrivateIpsRequest.md)|  | 

### Return type

[**LinkPrivateIpsResponse**](LinkPrivateIpsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNics

> ReadNicsResponse ReadNics(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadNicsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadNicsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNicsRequest** | [**optional.Interface of ReadNicsRequest**](ReadNicsRequest.md)|  | 

### Return type

[**ReadNicsResponse**](ReadNicsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkNic

> UnlinkNicResponse UnlinkNic(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkNicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkNicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkNicRequest** | [**optional.Interface of UnlinkNicRequest**](UnlinkNicRequest.md)|  | 

### Return type

[**UnlinkNicResponse**](UnlinkNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkPrivateIps

> UnlinkPrivateIpsResponse UnlinkPrivateIps(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkPrivateIpsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkPrivateIpsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkPrivateIpsRequest** | [**optional.Interface of UnlinkPrivateIpsRequest**](UnlinkPrivateIpsRequest.md)|  | 

### Return type

[**UnlinkPrivateIpsResponse**](UnlinkPrivateIpsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNic

> UpdateNicResponse UpdateNic(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateNicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateNicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNicRequest** | [**optional.Interface of UpdateNicRequest**](UpdateNicRequest.md)|  | 

### Return type

[**UpdateNicResponse**](UpdateNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

