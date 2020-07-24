# \InternetServiceApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternetService**](InternetServiceApi.md#CreateInternetService) | **Post** /CreateInternetService | 
[**DeleteInternetService**](InternetServiceApi.md#DeleteInternetService) | **Post** /DeleteInternetService | 
[**LinkInternetService**](InternetServiceApi.md#LinkInternetService) | **Post** /LinkInternetService | 
[**ReadInternetServices**](InternetServiceApi.md#ReadInternetServices) | **Post** /ReadInternetServices | 
[**UnlinkInternetService**](InternetServiceApi.md#UnlinkInternetService) | **Post** /UnlinkInternetService | 



## CreateInternetService

> CreateInternetServiceResponse CreateInternetService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateInternetServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInternetServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInternetServiceRequest** | [**optional.Interface of CreateInternetServiceRequest**](CreateInternetServiceRequest.md)|  | 

### Return type

[**CreateInternetServiceResponse**](CreateInternetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInternetService

> DeleteInternetServiceResponse DeleteInternetService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteInternetServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteInternetServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteInternetServiceRequest** | [**optional.Interface of DeleteInternetServiceRequest**](DeleteInternetServiceRequest.md)|  | 

### Return type

[**DeleteInternetServiceResponse**](DeleteInternetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkInternetService

> LinkInternetServiceResponse LinkInternetService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkInternetServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkInternetServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkInternetServiceRequest** | [**optional.Interface of LinkInternetServiceRequest**](LinkInternetServiceRequest.md)|  | 

### Return type

[**LinkInternetServiceResponse**](LinkInternetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInternetServices

> ReadInternetServicesResponse ReadInternetServices(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadInternetServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadInternetServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readInternetServicesRequest** | [**optional.Interface of ReadInternetServicesRequest**](ReadInternetServicesRequest.md)|  | 

### Return type

[**ReadInternetServicesResponse**](ReadInternetServicesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkInternetService

> UnlinkInternetServiceResponse UnlinkInternetService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkInternetServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkInternetServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkInternetServiceRequest** | [**optional.Interface of UnlinkInternetServiceRequest**](UnlinkInternetServiceRequest.md)|  | 

### Return type

[**UnlinkInternetServiceResponse**](UnlinkInternetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

