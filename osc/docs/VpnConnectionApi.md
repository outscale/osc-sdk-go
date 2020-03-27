# \VpnConnectionApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVpnConnection**](VpnConnectionApi.md#CreateVpnConnection) | **Post** /CreateVpnConnection | 
[**CreateVpnConnectionRoute**](VpnConnectionApi.md#CreateVpnConnectionRoute) | **Post** /CreateVpnConnectionRoute | 
[**DeleteVpnConnection**](VpnConnectionApi.md#DeleteVpnConnection) | **Post** /DeleteVpnConnection | 
[**DeleteVpnConnectionRoute**](VpnConnectionApi.md#DeleteVpnConnectionRoute) | **Post** /DeleteVpnConnectionRoute | 
[**ReadVpnConnections**](VpnConnectionApi.md#ReadVpnConnections) | **Post** /ReadVpnConnections | 



## CreateVpnConnection

> CreateVpnConnectionResponse CreateVpnConnection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateVpnConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVpnConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVpnConnectionRequest** | [**optional.Interface of CreateVpnConnectionRequest**](CreateVpnConnectionRequest.md)|  | 

### Return type

[**CreateVpnConnectionResponse**](CreateVpnConnectionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpnConnectionRoute

> CreateVpnConnectionRouteResponse CreateVpnConnectionRoute(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateVpnConnectionRouteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVpnConnectionRouteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVpnConnectionRouteRequest** | [**optional.Interface of CreateVpnConnectionRouteRequest**](CreateVpnConnectionRouteRequest.md)|  | 

### Return type

[**CreateVpnConnectionRouteResponse**](CreateVpnConnectionRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpnConnection

> DeleteVpnConnectionResponse DeleteVpnConnection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteVpnConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVpnConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVpnConnectionRequest** | [**optional.Interface of DeleteVpnConnectionRequest**](DeleteVpnConnectionRequest.md)|  | 

### Return type

[**DeleteVpnConnectionResponse**](DeleteVpnConnectionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpnConnectionRoute

> DeleteVpnConnectionRouteResponse DeleteVpnConnectionRoute(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteVpnConnectionRouteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVpnConnectionRouteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVpnConnectionRouteRequest** | [**optional.Interface of DeleteVpnConnectionRouteRequest**](DeleteVpnConnectionRouteRequest.md)|  | 

### Return type

[**DeleteVpnConnectionRouteResponse**](DeleteVpnConnectionRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVpnConnections

> ReadVpnConnectionsResponse ReadVpnConnections(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVpnConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVpnConnectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVpnConnectionsRequest** | [**optional.Interface of ReadVpnConnectionsRequest**](ReadVpnConnectionsRequest.md)|  | 

### Return type

[**ReadVpnConnectionsResponse**](ReadVpnConnectionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

