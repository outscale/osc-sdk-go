# \ListenerApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListenerRule**](ListenerApi.md#CreateListenerRule) | **Post** /CreateListenerRule | 
[**CreateLoadBalancerListeners**](ListenerApi.md#CreateLoadBalancerListeners) | **Post** /CreateLoadBalancerListeners | 
[**DeleteListenerRule**](ListenerApi.md#DeleteListenerRule) | **Post** /DeleteListenerRule | 
[**DeleteLoadBalancerListeners**](ListenerApi.md#DeleteLoadBalancerListeners) | **Post** /DeleteLoadBalancerListeners | 
[**ReadListenerRules**](ListenerApi.md#ReadListenerRules) | **Post** /ReadListenerRules | 
[**UpdateListenerRule**](ListenerApi.md#UpdateListenerRule) | **Post** /UpdateListenerRule | 



## CreateListenerRule

> CreateListenerRuleResponse CreateListenerRule(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateListenerRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateListenerRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createListenerRuleRequest** | [**optional.Interface of CreateListenerRuleRequest**](CreateListenerRuleRequest.md)|  | 

### Return type

[**CreateListenerRuleResponse**](CreateListenerRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## DeleteListenerRule

> DeleteListenerRuleResponse DeleteListenerRule(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteListenerRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteListenerRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteListenerRuleRequest** | [**optional.Interface of DeleteListenerRuleRequest**](DeleteListenerRuleRequest.md)|  | 

### Return type

[**DeleteListenerRuleResponse**](DeleteListenerRuleResponse.md)

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


## ReadListenerRules

> ReadListenerRulesResponse ReadListenerRules(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadListenerRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadListenerRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readListenerRulesRequest** | [**optional.Interface of ReadListenerRulesRequest**](ReadListenerRulesRequest.md)|  | 

### Return type

[**ReadListenerRulesResponse**](ReadListenerRulesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListenerRule

> UpdateListenerRuleResponse UpdateListenerRule(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateListenerRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateListenerRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateListenerRuleRequest** | [**optional.Interface of UpdateListenerRuleRequest**](UpdateListenerRuleRequest.md)|  | 

### Return type

[**UpdateListenerRuleResponse**](UpdateListenerRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

