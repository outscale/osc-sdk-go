# \VirtualGatewayApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualGateway**](VirtualGatewayApi.md#CreateVirtualGateway) | **Post** /CreateVirtualGateway | 
[**DeleteVirtualGateway**](VirtualGatewayApi.md#DeleteVirtualGateway) | **Post** /DeleteVirtualGateway | 
[**LinkVirtualGateway**](VirtualGatewayApi.md#LinkVirtualGateway) | **Post** /LinkVirtualGateway | 
[**ReadVirtualGateways**](VirtualGatewayApi.md#ReadVirtualGateways) | **Post** /ReadVirtualGateways | 
[**UnlinkVirtualGateway**](VirtualGatewayApi.md#UnlinkVirtualGateway) | **Post** /UnlinkVirtualGateway | 
[**UpdateRoutePropagation**](VirtualGatewayApi.md#UpdateRoutePropagation) | **Post** /UpdateRoutePropagation | 



## CreateVirtualGateway

> CreateVirtualGatewayResponse CreateVirtualGateway(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateVirtualGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVirtualGatewayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVirtualGatewayRequest** | [**optional.Interface of CreateVirtualGatewayRequest**](CreateVirtualGatewayRequest.md)|  | 

### Return type

[**CreateVirtualGatewayResponse**](CreateVirtualGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualGateway

> DeleteVirtualGatewayResponse DeleteVirtualGateway(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteVirtualGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVirtualGatewayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVirtualGatewayRequest** | [**optional.Interface of DeleteVirtualGatewayRequest**](DeleteVirtualGatewayRequest.md)|  | 

### Return type

[**DeleteVirtualGatewayResponse**](DeleteVirtualGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkVirtualGateway

> LinkVirtualGatewayResponse LinkVirtualGateway(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkVirtualGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkVirtualGatewayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkVirtualGatewayRequest** | [**optional.Interface of LinkVirtualGatewayRequest**](LinkVirtualGatewayRequest.md)|  | 

### Return type

[**LinkVirtualGatewayResponse**](LinkVirtualGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVirtualGateways

> ReadVirtualGatewaysResponse ReadVirtualGateways(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVirtualGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVirtualGatewaysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVirtualGatewaysRequest** | [**optional.Interface of ReadVirtualGatewaysRequest**](ReadVirtualGatewaysRequest.md)|  | 

### Return type

[**ReadVirtualGatewaysResponse**](ReadVirtualGatewaysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkVirtualGateway

> UnlinkVirtualGatewayResponse UnlinkVirtualGateway(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkVirtualGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkVirtualGatewayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkVirtualGatewayRequest** | [**optional.Interface of UnlinkVirtualGatewayRequest**](UnlinkVirtualGatewayRequest.md)|  | 

### Return type

[**UnlinkVirtualGatewayResponse**](UnlinkVirtualGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoutePropagation

> UpdateRoutePropagationResponse UpdateRoutePropagation(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateRoutePropagationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRoutePropagationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRoutePropagationRequest** | [**optional.Interface of UpdateRoutePropagationRequest**](UpdateRoutePropagationRequest.md)|  | 

### Return type

[**UpdateRoutePropagationResponse**](UpdateRoutePropagationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

