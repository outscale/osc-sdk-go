# \LoadBalancerApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancer**](LoadBalancerApi.md#CreateLoadBalancer) | **Post** /CreateLoadBalancer | 
[**DeleteLoadBalancer**](LoadBalancerApi.md#DeleteLoadBalancer) | **Post** /DeleteLoadBalancer | 
[**DeregisterVmsInLoadBalancer**](LoadBalancerApi.md#DeregisterVmsInLoadBalancer) | **Post** /DeregisterVmsInLoadBalancer | 
[**ReadLoadBalancers**](LoadBalancerApi.md#ReadLoadBalancers) | **Post** /ReadLoadBalancers | 
[**RegisterVmsInLoadBalancer**](LoadBalancerApi.md#RegisterVmsInLoadBalancer) | **Post** /RegisterVmsInLoadBalancer | 
[**UpdateLoadBalancer**](LoadBalancerApi.md#UpdateLoadBalancer) | **Post** /UpdateLoadBalancer | 



## CreateLoadBalancer

> CreateLoadBalancerResponse CreateLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerRequest** | [**optional.Interface of CreateLoadBalancerRequest**](CreateLoadBalancerRequest.md)|  | 

### Return type

[**CreateLoadBalancerResponse**](CreateLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancer

> DeleteLoadBalancerResponse DeleteLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerRequest** | [**optional.Interface of DeleteLoadBalancerRequest**](DeleteLoadBalancerRequest.md)|  | 

### Return type

[**DeleteLoadBalancerResponse**](DeleteLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeregisterVmsInLoadBalancer

> DeregisterVmsInLoadBalancerResponse DeregisterVmsInLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeregisterVmsInLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeregisterVmsInLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deregisterVmsInLoadBalancerRequest** | [**optional.Interface of DeregisterVmsInLoadBalancerRequest**](DeregisterVmsInLoadBalancerRequest.md)|  | 

### Return type

[**DeregisterVmsInLoadBalancerResponse**](DeregisterVmsInLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadLoadBalancers

> ReadLoadBalancersResponse ReadLoadBalancers(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadLoadBalancersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadLoadBalancersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLoadBalancersRequest** | [**optional.Interface of ReadLoadBalancersRequest**](ReadLoadBalancersRequest.md)|  | 

### Return type

[**ReadLoadBalancersResponse**](ReadLoadBalancersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterVmsInLoadBalancer

> RegisterVmsInLoadBalancerResponse RegisterVmsInLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterVmsInLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterVmsInLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerVmsInLoadBalancerRequest** | [**optional.Interface of RegisterVmsInLoadBalancerRequest**](RegisterVmsInLoadBalancerRequest.md)|  | 

### Return type

[**RegisterVmsInLoadBalancerResponse**](RegisterVmsInLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancer

> UpdateLoadBalancerResponse UpdateLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLoadBalancerRequest** | [**optional.Interface of UpdateLoadBalancerRequest**](UpdateLoadBalancerRequest.md)|  | 

### Return type

[**UpdateLoadBalancerResponse**](UpdateLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

