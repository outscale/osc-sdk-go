# \ApiAccessPolicyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadApiAccessPolicy**](ApiAccessPolicyApi.md#ReadApiAccessPolicy) | **Post** /ReadApiAccessPolicy | 
[**UpdateApiAccessPolicy**](ApiAccessPolicyApi.md#UpdateApiAccessPolicy) | **Post** /UpdateApiAccessPolicy | 



## ReadApiAccessPolicy

> ReadApiAccessPolicyResponse ReadApiAccessPolicy(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadApiAccessPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadApiAccessPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readApiAccessPolicyRequest** | [**optional.Interface of ReadApiAccessPolicyRequest**](ReadApiAccessPolicyRequest.md)|  | 

### Return type

[**ReadApiAccessPolicyResponse**](ReadApiAccessPolicyResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiAccessPolicy

> UpdateApiAccessPolicyResponse UpdateApiAccessPolicy(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateApiAccessPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApiAccessPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateApiAccessPolicyRequest** | [**optional.Interface of UpdateApiAccessPolicyRequest**](UpdateApiAccessPolicyRequest.md)|  | 

### Return type

[**UpdateApiAccessPolicyResponse**](UpdateApiAccessPolicyResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

