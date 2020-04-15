# \ResetPasswordEmailApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendResetPasswordEmail**](ResetPasswordEmailApi.md#SendResetPasswordEmail) | **Post** /SendResetPasswordEmail | 



## SendResetPasswordEmail

> SendResetPasswordEmailResponse SendResetPasswordEmail(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SendResetPasswordEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendResetPasswordEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendResetPasswordEmailRequest** | [**optional.Interface of SendResetPasswordEmailRequest**](SendResetPasswordEmailRequest.md)|  | 

### Return type

[**SendResetPasswordEmailResponse**](SendResetPasswordEmailResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

