# \VmTemplateApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVmTemplate**](VmTemplateApi.md#CreateVmTemplate) | **Post** /CreateVmTemplate | 
[**DeleteVmTemplate**](VmTemplateApi.md#DeleteVmTemplate) | **Post** /DeleteVmTemplate | 
[**ReadVmTemplates**](VmTemplateApi.md#ReadVmTemplates) | **Post** /ReadVmTemplates | 
[**UpdateVmTemplate**](VmTemplateApi.md#UpdateVmTemplate) | **Post** /UpdateVmTemplate | 



## CreateVmTemplate

> CreateVmTemplateResponse CreateVmTemplate(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateVmTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVmTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVmTemplateRequest** | [**optional.Interface of CreateVmTemplateRequest**](CreateVmTemplateRequest.md)|  | 

### Return type

[**CreateVmTemplateResponse**](CreateVmTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVmTemplate

> DeleteVmTemplateResponse DeleteVmTemplate(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteVmTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVmTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVmTemplateRequest** | [**optional.Interface of DeleteVmTemplateRequest**](DeleteVmTemplateRequest.md)|  | 

### Return type

[**DeleteVmTemplateResponse**](DeleteVmTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmTemplates

> ReadVmTemplatesResponse ReadVmTemplates(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVmTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVmTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmTemplatesRequest** | [**optional.Interface of ReadVmTemplatesRequest**](ReadVmTemplatesRequest.md)|  | 

### Return type

[**ReadVmTemplatesResponse**](ReadVmTemplatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVmTemplate

> UpdateVmTemplateResponse UpdateVmTemplate(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateVmTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateVmTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVmTemplateRequest** | [**optional.Interface of UpdateVmTemplateRequest**](UpdateVmTemplateRequest.md)|  | 

### Return type

[**UpdateVmTemplateResponse**](UpdateVmTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

