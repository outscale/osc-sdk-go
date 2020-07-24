# \FlexibleGpuApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlexibleGpu**](FlexibleGpuApi.md#CreateFlexibleGpu) | **Post** /CreateFlexibleGpu | 
[**DeleteFlexibleGpu**](FlexibleGpuApi.md#DeleteFlexibleGpu) | **Post** /DeleteFlexibleGpu | 
[**LinkFlexibleGpu**](FlexibleGpuApi.md#LinkFlexibleGpu) | **Post** /LinkFlexibleGpu | 
[**ReadFlexibleGpuCatalog**](FlexibleGpuApi.md#ReadFlexibleGpuCatalog) | **Post** /ReadFlexibleGpuCatalog | 
[**ReadFlexibleGpus**](FlexibleGpuApi.md#ReadFlexibleGpus) | **Post** /ReadFlexibleGpus | 
[**UnlinkFlexibleGpu**](FlexibleGpuApi.md#UnlinkFlexibleGpu) | **Post** /UnlinkFlexibleGpu | 
[**UpdateFlexibleGpu**](FlexibleGpuApi.md#UpdateFlexibleGpu) | **Post** /UpdateFlexibleGpu | 



## CreateFlexibleGpu

> CreateFlexibleGpuResponse CreateFlexibleGpu(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFlexibleGpuOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFlexibleGpuOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFlexibleGpuRequest** | [**optional.Interface of CreateFlexibleGpuRequest**](CreateFlexibleGpuRequest.md)|  | 

### Return type

[**CreateFlexibleGpuResponse**](CreateFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlexibleGpu

> DeleteFlexibleGpuResponse DeleteFlexibleGpu(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteFlexibleGpuOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteFlexibleGpuOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteFlexibleGpuRequest** | [**optional.Interface of DeleteFlexibleGpuRequest**](DeleteFlexibleGpuRequest.md)|  | 

### Return type

[**DeleteFlexibleGpuResponse**](DeleteFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkFlexibleGpu

> LinkFlexibleGpuResponse LinkFlexibleGpu(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkFlexibleGpuOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkFlexibleGpuOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkFlexibleGpuRequest** | [**optional.Interface of LinkFlexibleGpuRequest**](LinkFlexibleGpuRequest.md)|  | 

### Return type

[**LinkFlexibleGpuResponse**](LinkFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadFlexibleGpuCatalog

> ReadFlexibleGpuCatalogResponse ReadFlexibleGpuCatalog(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadFlexibleGpuCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadFlexibleGpuCatalogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readFlexibleGpuCatalogRequest** | [**optional.Interface of ReadFlexibleGpuCatalogRequest**](ReadFlexibleGpuCatalogRequest.md)|  | 

### Return type

[**ReadFlexibleGpuCatalogResponse**](ReadFlexibleGpuCatalogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadFlexibleGpus

> ReadFlexibleGpusResponse ReadFlexibleGpus(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadFlexibleGpusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadFlexibleGpusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readFlexibleGpusRequest** | [**optional.Interface of ReadFlexibleGpusRequest**](ReadFlexibleGpusRequest.md)|  | 

### Return type

[**ReadFlexibleGpusResponse**](ReadFlexibleGpusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkFlexibleGpu

> UnlinkFlexibleGpuResponse UnlinkFlexibleGpu(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkFlexibleGpuOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkFlexibleGpuOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkFlexibleGpuRequest** | [**optional.Interface of UnlinkFlexibleGpuRequest**](UnlinkFlexibleGpuRequest.md)|  | 

### Return type

[**UnlinkFlexibleGpuResponse**](UnlinkFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlexibleGpu

> UpdateFlexibleGpuResponse UpdateFlexibleGpu(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateFlexibleGpuOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFlexibleGpuOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFlexibleGpuRequest** | [**optional.Interface of UpdateFlexibleGpuRequest**](UpdateFlexibleGpuRequest.md)|  | 

### Return type

[**UpdateFlexibleGpuResponse**](UpdateFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

