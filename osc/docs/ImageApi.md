# \ImageApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](ImageApi.md#CreateImage) | **Post** /CreateImage | 
[**CreateImageExportTask**](ImageApi.md#CreateImageExportTask) | **Post** /CreateImageExportTask | 
[**DeleteImage**](ImageApi.md#DeleteImage) | **Post** /DeleteImage | 
[**ReadImageExportTasks**](ImageApi.md#ReadImageExportTasks) | **Post** /ReadImageExportTasks | 
[**ReadImages**](ImageApi.md#ReadImages) | **Post** /ReadImages | 
[**UpdateImage**](ImageApi.md#UpdateImage) | **Post** /UpdateImage | 



## CreateImage

> CreateImageResponse CreateImage(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImageRequest** | [**optional.Interface of CreateImageRequest**](CreateImageRequest.md)|  | 

### Return type

[**CreateImageResponse**](CreateImageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateImageExportTask

> CreateImageExportTaskResponse CreateImageExportTask(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateImageExportTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateImageExportTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImageExportTaskRequest** | [**optional.Interface of CreateImageExportTaskRequest**](CreateImageExportTaskRequest.md)|  | 

### Return type

[**CreateImageExportTaskResponse**](CreateImageExportTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImageResponse DeleteImage(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteImageRequest** | [**optional.Interface of DeleteImageRequest**](DeleteImageRequest.md)|  | 

### Return type

[**DeleteImageResponse**](DeleteImageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadImageExportTasks

> ReadImageExportTasksResponse ReadImageExportTasks(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadImageExportTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadImageExportTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readImageExportTasksRequest** | [**optional.Interface of ReadImageExportTasksRequest**](ReadImageExportTasksRequest.md)|  | 

### Return type

[**ReadImageExportTasksResponse**](ReadImageExportTasksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadImages

> ReadImagesResponse ReadImages(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadImagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readImagesRequest** | [**optional.Interface of ReadImagesRequest**](ReadImagesRequest.md)|  | 

### Return type

[**ReadImagesResponse**](ReadImagesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImage

> UpdateImageResponse UpdateImage(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateImageRequest** | [**optional.Interface of UpdateImageRequest**](UpdateImageRequest.md)|  | 

### Return type

[**UpdateImageResponse**](UpdateImageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

