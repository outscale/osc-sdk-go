# \AccessKeyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessKey**](AccessKeyApi.md#CreateAccessKey) | **Post** /CreateAccessKey | 
[**DeleteAccessKey**](AccessKeyApi.md#DeleteAccessKey) | **Post** /DeleteAccessKey | 
[**UpdateAccessKey**](AccessKeyApi.md#UpdateAccessKey) | **Post** /UpdateAccessKey | 



## CreateAccessKey

> CreateAccessKeyResponse CreateAccessKey(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CreateAccessKeyResponse**](CreateAccessKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessKey

> DeleteAccessKeyResponse DeleteAccessKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteAccessKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteAccessKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAccessKeyRequest** | [**optional.Interface of DeleteAccessKeyRequest**](DeleteAccessKeyRequest.md)|  | 

### Return type

[**DeleteAccessKeyResponse**](DeleteAccessKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessKey

> UpdateAccessKeyResponse UpdateAccessKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateAccessKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAccessKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccessKeyRequest** | [**optional.Interface of UpdateAccessKeyRequest**](UpdateAccessKeyRequest.md)|  | 

### Return type

[**UpdateAccessKeyResponse**](UpdateAccessKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

