# \LoadBalancerPolicyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancerPolicy**](LoadBalancerPolicyApi.md#CreateLoadBalancerPolicy) | **Post** /CreateLoadBalancerPolicy | 
[**DeleteLoadBalancerPolicy**](LoadBalancerPolicyApi.md#DeleteLoadBalancerPolicy) | **Post** /DeleteLoadBalancerPolicy | 



## CreateLoadBalancerPolicy

> CreateLoadBalancerPolicyResponse CreateLoadBalancerPolicy(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateLoadBalancerPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLoadBalancerPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerPolicyRequest** | [**optional.Interface of CreateLoadBalancerPolicyRequest**](CreateLoadBalancerPolicyRequest.md)|  | 

### Return type

[**CreateLoadBalancerPolicyResponse**](CreateLoadBalancerPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancerPolicy

> DeleteLoadBalancerPolicyResponse DeleteLoadBalancerPolicy(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteLoadBalancerPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteLoadBalancerPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerPolicyRequest** | [**optional.Interface of DeleteLoadBalancerPolicyRequest**](DeleteLoadBalancerPolicyRequest.md)|  | 

### Return type

[**DeleteLoadBalancerPolicyResponse**](DeleteLoadBalancerPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

