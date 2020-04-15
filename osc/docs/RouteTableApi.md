# \RouteTableApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteTable**](RouteTableApi.md#CreateRouteTable) | **Post** /CreateRouteTable | 
[**DeleteRouteTable**](RouteTableApi.md#DeleteRouteTable) | **Post** /DeleteRouteTable | 
[**LinkRouteTable**](RouteTableApi.md#LinkRouteTable) | **Post** /LinkRouteTable | 
[**ReadRouteTables**](RouteTableApi.md#ReadRouteTables) | **Post** /ReadRouteTables | 
[**UnlinkRouteTable**](RouteTableApi.md#UnlinkRouteTable) | **Post** /UnlinkRouteTable | 



## CreateRouteTable

> CreateRouteTableResponse CreateRouteTable(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRouteTableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRouteTableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteTableRequest** | [**optional.Interface of CreateRouteTableRequest**](CreateRouteTableRequest.md)|  | 

### Return type

[**CreateRouteTableResponse**](CreateRouteTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteTable

> DeleteRouteTableResponse DeleteRouteTable(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteRouteTableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteRouteTableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRouteTableRequest** | [**optional.Interface of DeleteRouteTableRequest**](DeleteRouteTableRequest.md)|  | 

### Return type

[**DeleteRouteTableResponse**](DeleteRouteTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkRouteTable

> LinkRouteTableResponse LinkRouteTable(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkRouteTableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkRouteTableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkRouteTableRequest** | [**optional.Interface of LinkRouteTableRequest**](LinkRouteTableRequest.md)|  | 

### Return type

[**LinkRouteTableResponse**](LinkRouteTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadRouteTables

> ReadRouteTablesResponse ReadRouteTables(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadRouteTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadRouteTablesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readRouteTablesRequest** | [**optional.Interface of ReadRouteTablesRequest**](ReadRouteTablesRequest.md)|  | 

### Return type

[**ReadRouteTablesResponse**](ReadRouteTablesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkRouteTable

> UnlinkRouteTableResponse UnlinkRouteTable(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkRouteTableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkRouteTableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkRouteTableRequest** | [**optional.Interface of UnlinkRouteTableRequest**](UnlinkRouteTableRequest.md)|  | 

### Return type

[**UnlinkRouteTableResponse**](UnlinkRouteTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

