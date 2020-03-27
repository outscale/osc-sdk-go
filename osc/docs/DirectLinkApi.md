# \DirectLinkApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectLink**](DirectLinkApi.md#CreateDirectLink) | **Post** /CreateDirectLink | 
[**DeleteDirectLink**](DirectLinkApi.md#DeleteDirectLink) | **Post** /DeleteDirectLink | 
[**ReadDirectLinks**](DirectLinkApi.md#ReadDirectLinks) | **Post** /ReadDirectLinks | 



## CreateDirectLink

> CreateDirectLinkResponse CreateDirectLink(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDirectLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDirectLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDirectLinkRequest** | [**optional.Interface of CreateDirectLinkRequest**](CreateDirectLinkRequest.md)|  | 

### Return type

[**CreateDirectLinkResponse**](CreateDirectLinkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDirectLink

> DeleteDirectLinkResponse DeleteDirectLink(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteDirectLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDirectLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDirectLinkRequest** | [**optional.Interface of DeleteDirectLinkRequest**](DeleteDirectLinkRequest.md)|  | 

### Return type

[**DeleteDirectLinkResponse**](DeleteDirectLinkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDirectLinks

> ReadDirectLinksResponse ReadDirectLinks(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadDirectLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadDirectLinksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readDirectLinksRequest** | [**optional.Interface of ReadDirectLinksRequest**](ReadDirectLinksRequest.md)|  | 

### Return type

[**ReadDirectLinksResponse**](ReadDirectLinksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

