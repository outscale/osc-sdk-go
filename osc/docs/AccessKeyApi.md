# \AccessKeyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessKey**](AccessKeyApi.md#CreateAccessKey) | **Post** /CreateAccessKey | 
[**DeleteAccessKey**](AccessKeyApi.md#DeleteAccessKey) | **Post** /DeleteAccessKey | 
[**ReadAccessKeys**](AccessKeyApi.md#ReadAccessKeys) | **Post** /ReadAccessKeys | 
[**ReadSecretAccessKey**](AccessKeyApi.md#ReadSecretAccessKey) | **Post** /ReadSecretAccessKey | 
[**UpdateAccessKey**](AccessKeyApi.md#UpdateAccessKey) | **Post** /UpdateAccessKey | 



## CreateAccessKey

> CreateAccessKeyResponse CreateAccessKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAccessKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAccessKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessKeyRequest** | [**optional.Interface of CreateAccessKeyRequest**](CreateAccessKeyRequest.md)|  | 

### Return type

[**CreateAccessKeyResponse**](CreateAccessKeyResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
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

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccessKeys

> ReadAccessKeysResponse ReadAccessKeys(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadAccessKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadAccessKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readAccessKeysRequest** | [**optional.Interface of ReadAccessKeysRequest**](ReadAccessKeysRequest.md)|  | 

### Return type

[**ReadAccessKeysResponse**](ReadAccessKeysResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSecretAccessKey

> ReadSecretAccessKeyResponse ReadSecretAccessKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadSecretAccessKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadSecretAccessKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSecretAccessKeyRequest** | [**optional.Interface of ReadSecretAccessKeyRequest**](ReadSecretAccessKeyRequest.md)|  | 

### Return type

[**ReadSecretAccessKeyResponse**](ReadSecretAccessKeyResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

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

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

