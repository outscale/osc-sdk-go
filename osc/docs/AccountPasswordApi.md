# \AccountPasswordApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResetAccountPassword**](AccountPasswordApi.md#ResetAccountPassword) | **Post** /ResetAccountPassword | 



## ResetAccountPassword

> ResetAccountPasswordResponse ResetAccountPassword(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResetAccountPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ResetAccountPasswordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetAccountPasswordRequest** | [**optional.Interface of ResetAccountPasswordRequest**](ResetAccountPasswordRequest.md)|  | 

### Return type

[**ResetAccountPasswordResponse**](ResetAccountPasswordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

