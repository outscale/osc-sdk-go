# \NetApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNet**](NetApi.md#CreateNet) | **Post** /CreateNet | 
[**DeleteNet**](NetApi.md#DeleteNet) | **Post** /DeleteNet | 
[**ReadNets**](NetApi.md#ReadNets) | **Post** /ReadNets | 
[**UpdateNet**](NetApi.md#UpdateNet) | **Post** /UpdateNet | 



## CreateNet

> CreateNetResponse CreateNet(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetRequest** | [**optional.Interface of CreateNetRequest**](CreateNetRequest.md)|  | 

### Return type

[**CreateNetResponse**](CreateNetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNet

> DeleteNetResponse DeleteNet(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteNetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteNetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNetRequest** | [**optional.Interface of DeleteNetRequest**](DeleteNetRequest.md)|  | 

### Return type

[**DeleteNetResponse**](DeleteNetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNets

> ReadNetsResponse ReadNets(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadNetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadNetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNetsRequest** | [**optional.Interface of ReadNetsRequest**](ReadNetsRequest.md)|  | 

### Return type

[**ReadNetsResponse**](ReadNetsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNet

> UpdateNetResponse UpdateNet(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateNetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateNetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNetRequest** | [**optional.Interface of UpdateNetRequest**](UpdateNetRequest.md)|  | 

### Return type

[**UpdateNetResponse**](UpdateNetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

