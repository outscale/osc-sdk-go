# \MasterKeyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMasterKey**](MasterKeyApi.md#CreateMasterKey) | **Post** /CreateMasterKey | 
[**DecryptCiphertext**](MasterKeyApi.md#DecryptCiphertext) | **Post** /DecryptCiphertext | 
[**DeleteMasterKey**](MasterKeyApi.md#DeleteMasterKey) | **Post** /DeleteMasterKey | 
[**EncryptPlaintext**](MasterKeyApi.md#EncryptPlaintext) | **Post** /EncryptPlaintext | 
[**GenerateDataKey**](MasterKeyApi.md#GenerateDataKey) | **Post** /GenerateDataKey | 
[**ReadMasterKeys**](MasterKeyApi.md#ReadMasterKeys) | **Post** /ReadMasterKeys | 
[**UndeleteMasterKey**](MasterKeyApi.md#UndeleteMasterKey) | **Post** /UndeleteMasterKey | 
[**UpdateMasterKey**](MasterKeyApi.md#UpdateMasterKey) | **Post** /UpdateMasterKey | 



## CreateMasterKey

> CreateMasterKeyResponse CreateMasterKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateMasterKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMasterKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMasterKeyRequest** | [**optional.Interface of CreateMasterKeyRequest**](CreateMasterKeyRequest.md)|  | 

### Return type

[**CreateMasterKeyResponse**](CreateMasterKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecryptCiphertext

> DecryptCiphertextResponse DecryptCiphertext(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DecryptCiphertextOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DecryptCiphertextOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decryptCiphertextRequest** | [**optional.Interface of DecryptCiphertextRequest**](DecryptCiphertextRequest.md)|  | 

### Return type

[**DecryptCiphertextResponse**](DecryptCiphertextResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMasterKey

> DeleteMasterKeyResponse DeleteMasterKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteMasterKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteMasterKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMasterKeyRequest** | [**optional.Interface of DeleteMasterKeyRequest**](DeleteMasterKeyRequest.md)|  | 

### Return type

[**DeleteMasterKeyResponse**](DeleteMasterKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncryptPlaintext

> EncryptPlaintextResponse EncryptPlaintext(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EncryptPlaintextOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EncryptPlaintextOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptPlaintextRequest** | [**optional.Interface of EncryptPlaintextRequest**](EncryptPlaintextRequest.md)|  | 

### Return type

[**EncryptPlaintextResponse**](EncryptPlaintextResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateDataKey

> GenerateDataKeyResponse GenerateDataKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GenerateDataKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GenerateDataKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateDataKeyRequest** | [**optional.Interface of GenerateDataKeyRequest**](GenerateDataKeyRequest.md)|  | 

### Return type

[**GenerateDataKeyResponse**](GenerateDataKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMasterKeys

> ReadMasterKeysResponse ReadMasterKeys(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadMasterKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadMasterKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readMasterKeysRequest** | [**optional.Interface of ReadMasterKeysRequest**](ReadMasterKeysRequest.md)|  | 

### Return type

[**ReadMasterKeysResponse**](ReadMasterKeysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeleteMasterKey

> UndeleteMasterKeyResponse UndeleteMasterKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UndeleteMasterKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UndeleteMasterKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **undeleteMasterKeyRequest** | [**optional.Interface of UndeleteMasterKeyRequest**](UndeleteMasterKeyRequest.md)|  | 

### Return type

[**UndeleteMasterKeyResponse**](UndeleteMasterKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMasterKey

> UpdateMasterKeyResponse UpdateMasterKey(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateMasterKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMasterKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMasterKeyRequest** | [**optional.Interface of UpdateMasterKeyRequest**](UpdateMasterKeyRequest.md)|  | 

### Return type

[**UpdateMasterKeyResponse**](UpdateMasterKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

