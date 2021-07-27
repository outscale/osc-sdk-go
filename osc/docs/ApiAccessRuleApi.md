# \ApiAccessRuleApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiAccessRule**](ApiAccessRuleApi.md#CreateApiAccessRule) | **Post** /CreateApiAccessRule | 
[**DeleteApiAccessRule**](ApiAccessRuleApi.md#DeleteApiAccessRule) | **Post** /DeleteApiAccessRule | 
[**ReadApiAccessRules**](ApiAccessRuleApi.md#ReadApiAccessRules) | **Post** /ReadApiAccessRules | 
[**UpdateApiAccessRule**](ApiAccessRuleApi.md#UpdateApiAccessRule) | **Post** /UpdateApiAccessRule | 



## CreateApiAccessRule

> CreateApiAccessRuleResponse CreateApiAccessRule(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateApiAccessRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApiAccessRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiAccessRuleRequest** | [**optional.Interface of CreateApiAccessRuleRequest**](CreateApiAccessRuleRequest.md)|  | 

### Return type

[**CreateApiAccessRuleResponse**](CreateApiAccessRuleResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiAccessRule

> DeleteApiAccessRuleResponse DeleteApiAccessRule(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteApiAccessRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteApiAccessRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteApiAccessRuleRequest** | [**optional.Interface of DeleteApiAccessRuleRequest**](DeleteApiAccessRuleRequest.md)|  | 

### Return type

[**DeleteApiAccessRuleResponse**](DeleteApiAccessRuleResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadApiAccessRules

> ReadApiAccessRulesResponse ReadApiAccessRules(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadApiAccessRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadApiAccessRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readApiAccessRulesRequest** | [**optional.Interface of ReadApiAccessRulesRequest**](ReadApiAccessRulesRequest.md)|  | 

### Return type

[**ReadApiAccessRulesResponse**](ReadApiAccessRulesResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiAccessRule

> UpdateApiAccessRuleResponse UpdateApiAccessRule(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateApiAccessRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApiAccessRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateApiAccessRuleRequest** | [**optional.Interface of UpdateApiAccessRuleRequest**](UpdateApiAccessRuleRequest.md)|  | 

### Return type

[**UpdateApiAccessRuleResponse**](UpdateApiAccessRuleResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

