# \ServerCertificateApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerCertificate**](ServerCertificateApi.md#CreateServerCertificate) | **Post** /CreateServerCertificate | 
[**DeleteServerCertificate**](ServerCertificateApi.md#DeleteServerCertificate) | **Post** /DeleteServerCertificate | 
[**ReadServerCertificates**](ServerCertificateApi.md#ReadServerCertificates) | **Post** /ReadServerCertificates | 
[**UpdateServerCertificate**](ServerCertificateApi.md#UpdateServerCertificate) | **Post** /UpdateServerCertificate | 



## CreateServerCertificate

> CreateServerCertificateResponse CreateServerCertificate(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServerCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServerCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerCertificateRequest** | [**optional.Interface of CreateServerCertificateRequest**](CreateServerCertificateRequest.md)|  | 

### Return type

[**CreateServerCertificateResponse**](CreateServerCertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerCertificate

> DeleteServerCertificateResponse DeleteServerCertificate(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteServerCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteServerCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteServerCertificateRequest** | [**optional.Interface of DeleteServerCertificateRequest**](DeleteServerCertificateRequest.md)|  | 

### Return type

[**DeleteServerCertificateResponse**](DeleteServerCertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadServerCertificates

> ReadServerCertificatesResponse ReadServerCertificates(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadServerCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadServerCertificatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readServerCertificatesRequest** | [**optional.Interface of ReadServerCertificatesRequest**](ReadServerCertificatesRequest.md)|  | 

### Return type

[**ReadServerCertificatesResponse**](ReadServerCertificatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerCertificate

> UpdateServerCertificateResponse UpdateServerCertificate(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateServerCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServerCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateServerCertificateRequest** | [**optional.Interface of UpdateServerCertificateRequest**](UpdateServerCertificateRequest.md)|  | 

### Return type

[**UpdateServerCertificateResponse**](UpdateServerCertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

