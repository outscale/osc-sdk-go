# \ListenerApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancerListeners**](ListenerApi.md#CreateLoadBalancerListeners) | **Post** /CreateLoadBalancerListeners | 
[**DeleteLoadBalancerListeners**](ListenerApi.md#DeleteLoadBalancerListeners) | **Post** /DeleteLoadBalancerListeners | 



## CreateLoadBalancerListeners

> CreateLoadBalancerListenersResponse CreateLoadBalancerListeners(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateLoadBalancerListenersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLoadBalancerListenersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerListenersRequest** | [**optional.Interface of CreateLoadBalancerListenersRequest**](CreateLoadBalancerListenersRequest.md)|  | 

### Return type

[**CreateLoadBalancerListenersResponse**](CreateLoadBalancerListenersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancerListeners

> DeleteLoadBalancerListenersResponse DeleteLoadBalancerListeners(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteLoadBalancerListenersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteLoadBalancerListenersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerListenersRequest** | [**optional.Interface of DeleteLoadBalancerListenersRequest**](DeleteLoadBalancerListenersRequest.md)|  | 

### Return type

[**DeleteLoadBalancerListenersResponse**](DeleteLoadBalancerListenersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

