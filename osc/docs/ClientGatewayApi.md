# \ClientGatewayApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientGateway**](ClientGatewayApi.md#CreateClientGateway) | **Post** /CreateClientGateway | 
[**DeleteClientGateway**](ClientGatewayApi.md#DeleteClientGateway) | **Post** /DeleteClientGateway | 
[**ReadClientGateways**](ClientGatewayApi.md#ReadClientGateways) | **Post** /ReadClientGateways | 



## CreateClientGateway

> CreateClientGatewayResponse CreateClientGateway(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateClientGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateClientGatewayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClientGatewayRequest** | [**optional.Interface of CreateClientGatewayRequest**](CreateClientGatewayRequest.md)|  | 

### Return type

[**CreateClientGatewayResponse**](CreateClientGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientGateway

> DeleteClientGatewayResponse DeleteClientGateway(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteClientGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteClientGatewayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteClientGatewayRequest** | [**optional.Interface of DeleteClientGatewayRequest**](DeleteClientGatewayRequest.md)|  | 

### Return type

[**DeleteClientGatewayResponse**](DeleteClientGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadClientGateways

> ReadClientGatewaysResponse ReadClientGateways(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadClientGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadClientGatewaysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readClientGatewaysRequest** | [**optional.Interface of ReadClientGatewaysRequest**](ReadClientGatewaysRequest.md)|  | 

### Return type

[**ReadClientGatewaysResponse**](ReadClientGatewaysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

