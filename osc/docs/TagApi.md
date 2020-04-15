# \TagApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTags**](TagApi.md#CreateTags) | **Post** /CreateTags | 
[**DeleteTags**](TagApi.md#DeleteTags) | **Post** /DeleteTags | 
[**ReadTags**](TagApi.md#ReadTags) | **Post** /ReadTags | 



## CreateTags

> CreateTagsResponse CreateTags(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagsRequest** | [**optional.Interface of CreateTagsRequest**](CreateTagsRequest.md)|  | 

### Return type

[**CreateTagsResponse**](CreateTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTags

> DeleteTagsResponse DeleteTags(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTagsRequest** | [**optional.Interface of DeleteTagsRequest**](DeleteTagsRequest.md)|  | 

### Return type

[**DeleteTagsResponse**](DeleteTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTags

> ReadTagsResponse ReadTags(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readTagsRequest** | [**optional.Interface of ReadTagsRequest**](ReadTagsRequest.md)|  | 

### Return type

[**ReadTagsResponse**](ReadTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

