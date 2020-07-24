# \SecurityGroupRuleApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityGroupRule**](SecurityGroupRuleApi.md#CreateSecurityGroupRule) | **Post** /CreateSecurityGroupRule | 
[**DeleteSecurityGroupRule**](SecurityGroupRuleApi.md#DeleteSecurityGroupRule) | **Post** /DeleteSecurityGroupRule | 



## CreateSecurityGroupRule

> CreateSecurityGroupRuleResponse CreateSecurityGroupRule(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSecurityGroupRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSecurityGroupRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSecurityGroupRuleRequest** | [**optional.Interface of CreateSecurityGroupRuleRequest**](CreateSecurityGroupRuleRequest.md)|  | 

### Return type

[**CreateSecurityGroupRuleResponse**](CreateSecurityGroupRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityGroupRule

> DeleteSecurityGroupRuleResponse DeleteSecurityGroupRule(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteSecurityGroupRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSecurityGroupRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSecurityGroupRuleRequest** | [**optional.Interface of DeleteSecurityGroupRuleRequest**](DeleteSecurityGroupRuleRequest.md)|  | 

### Return type

[**DeleteSecurityGroupRuleResponse**](DeleteSecurityGroupRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

