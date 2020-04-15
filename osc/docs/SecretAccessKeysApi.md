# \SecretAccessKeysApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadSecretAccessKey**](SecretAccessKeysApi.md#ReadSecretAccessKey) | **Post** /ReadSecretAccessKey | 



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

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

