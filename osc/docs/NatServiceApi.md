# \NatServiceApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNatService**](NatServiceApi.md#CreateNatService) | **Post** /CreateNatService | 
[**DeleteNatService**](NatServiceApi.md#DeleteNatService) | **Post** /DeleteNatService | 
[**ReadNatServices**](NatServiceApi.md#ReadNatServices) | **Post** /ReadNatServices | 



## CreateNatService

> CreateNatServiceResponse CreateNatService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNatServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNatServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNatServiceRequest** | [**optional.Interface of CreateNatServiceRequest**](CreateNatServiceRequest.md)|  | 

### Return type

[**CreateNatServiceResponse**](CreateNatServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNatService

> DeleteNatServiceResponse DeleteNatService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteNatServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteNatServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNatServiceRequest** | [**optional.Interface of DeleteNatServiceRequest**](DeleteNatServiceRequest.md)|  | 

### Return type

[**DeleteNatServiceResponse**](DeleteNatServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNatServices

> ReadNatServicesResponse ReadNatServices(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadNatServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadNatServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNatServicesRequest** | [**optional.Interface of ReadNatServicesRequest**](ReadNatServicesRequest.md)|  | 

### Return type

[**ReadNatServicesResponse**](ReadNatServicesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

