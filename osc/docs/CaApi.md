# \CaApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCa**](CaApi.md#CreateCa) | **Post** /CreateCa | 
[**DeleteCa**](CaApi.md#DeleteCa) | **Post** /DeleteCa | 
[**ReadCas**](CaApi.md#ReadCas) | **Post** /ReadCas | 
[**UpdateCa**](CaApi.md#UpdateCa) | **Post** /UpdateCa | 



## CreateCa

> CreateCaResponse CreateCa(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCaRequest** | [**optional.Interface of CreateCaRequest**](CreateCaRequest.md)|  | 

### Return type

[**CreateCaResponse**](CreateCaResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCa

> DeleteCaResponse DeleteCa(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteCaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteCaRequest** | [**optional.Interface of DeleteCaRequest**](DeleteCaRequest.md)|  | 

### Return type

[**DeleteCaResponse**](DeleteCaResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCas

> ReadCasResponse ReadCas(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadCasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadCasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readCasRequest** | [**optional.Interface of ReadCasRequest**](ReadCasRequest.md)|  | 

### Return type

[**ReadCasResponse**](ReadCasResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCa

> UpdateCaResponse UpdateCa(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateCaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCaRequest** | [**optional.Interface of UpdateCaRequest**](UpdateCaRequest.md)|  | 

### Return type

[**UpdateCaResponse**](UpdateCaResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

