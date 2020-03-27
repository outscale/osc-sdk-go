# \VolumeApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](VolumeApi.md#CreateVolume) | **Post** /CreateVolume | 
[**DeleteVolume**](VolumeApi.md#DeleteVolume) | **Post** /DeleteVolume | 
[**LinkVolume**](VolumeApi.md#LinkVolume) | **Post** /LinkVolume | 
[**ReadVolumes**](VolumeApi.md#ReadVolumes) | **Post** /ReadVolumes | 
[**UnlinkVolume**](VolumeApi.md#UnlinkVolume) | **Post** /UnlinkVolume | 



## CreateVolume

> CreateVolumeResponse CreateVolume(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateVolumeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVolumeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVolumeRequest** | [**optional.Interface of CreateVolumeRequest**](CreateVolumeRequest.md)|  | 

### Return type

[**CreateVolumeResponse**](CreateVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> DeleteVolumeResponse DeleteVolume(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteVolumeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVolumeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVolumeRequest** | [**optional.Interface of DeleteVolumeRequest**](DeleteVolumeRequest.md)|  | 

### Return type

[**DeleteVolumeResponse**](DeleteVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkVolume

> LinkVolumeResponse LinkVolume(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkVolumeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkVolumeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkVolumeRequest** | [**optional.Interface of LinkVolumeRequest**](LinkVolumeRequest.md)|  | 

### Return type

[**LinkVolumeResponse**](LinkVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVolumes

> ReadVolumesResponse ReadVolumes(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVolumesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVolumesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVolumesRequest** | [**optional.Interface of ReadVolumesRequest**](ReadVolumesRequest.md)|  | 

### Return type

[**ReadVolumesResponse**](ReadVolumesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkVolume

> UnlinkVolumeResponse UnlinkVolume(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkVolumeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkVolumeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkVolumeRequest** | [**optional.Interface of UnlinkVolumeRequest**](UnlinkVolumeRequest.md)|  | 

### Return type

[**UnlinkVolumeResponse**](UnlinkVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

