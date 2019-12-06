# \NetPeeringApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptNetPeering**](NetPeeringApi.md#AcceptNetPeering) | **Post** /AcceptNetPeering | 
[**CreateNetPeering**](NetPeeringApi.md#CreateNetPeering) | **Post** /CreateNetPeering | 
[**DeleteNetPeering**](NetPeeringApi.md#DeleteNetPeering) | **Post** /DeleteNetPeering | 
[**ReadNetPeerings**](NetPeeringApi.md#ReadNetPeerings) | **Post** /ReadNetPeerings | 
[**RejectNetPeering**](NetPeeringApi.md#RejectNetPeering) | **Post** /RejectNetPeering | 



## AcceptNetPeering

> AcceptNetPeeringResponse AcceptNetPeering(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AcceptNetPeeringOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AcceptNetPeeringOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptNetPeeringRequest** | [**optional.Interface of AcceptNetPeeringRequest**](AcceptNetPeeringRequest.md)|  | 

### Return type

[**AcceptNetPeeringResponse**](AcceptNetPeeringResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetPeering

> CreateNetPeeringResponse CreateNetPeering(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNetPeeringOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNetPeeringOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetPeeringRequest** | [**optional.Interface of CreateNetPeeringRequest**](CreateNetPeeringRequest.md)|  | 

### Return type

[**CreateNetPeeringResponse**](CreateNetPeeringResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetPeering

> DeleteNetPeeringResponse DeleteNetPeering(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteNetPeeringOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteNetPeeringOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNetPeeringRequest** | [**optional.Interface of DeleteNetPeeringRequest**](DeleteNetPeeringRequest.md)|  | 

### Return type

[**DeleteNetPeeringResponse**](DeleteNetPeeringResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNetPeerings

> ReadNetPeeringsResponse ReadNetPeerings(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadNetPeeringsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadNetPeeringsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNetPeeringsRequest** | [**optional.Interface of ReadNetPeeringsRequest**](ReadNetPeeringsRequest.md)|  | 

### Return type

[**ReadNetPeeringsResponse**](ReadNetPeeringsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectNetPeering

> RejectNetPeeringResponse RejectNetPeering(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RejectNetPeeringOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RejectNetPeeringOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rejectNetPeeringRequest** | [**optional.Interface of RejectNetPeeringRequest**](RejectNetPeeringRequest.md)|  | 

### Return type

[**RejectNetPeeringResponse**](RejectNetPeeringResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

