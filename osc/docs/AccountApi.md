# \AccountApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAuthentication**](AccountApi.md#CheckAuthentication) | **Post** /CheckAuthentication | 
[**CreateAccount**](AccountApi.md#CreateAccount) | **Post** /CreateAccount | 
[**ReadAccounts**](AccountApi.md#ReadAccounts) | **Post** /ReadAccounts | 
[**ReadConsumptionAccount**](AccountApi.md#ReadConsumptionAccount) | **Post** /ReadConsumptionAccount | 
[**ResetAccountPassword**](AccountApi.md#ResetAccountPassword) | **Post** /ResetAccountPassword | 
[**SendResetPasswordEmail**](AccountApi.md#SendResetPasswordEmail) | **Post** /SendResetPasswordEmail | 
[**UpdateAccount**](AccountApi.md#UpdateAccount) | **Post** /UpdateAccount | 



## CheckAuthentication

> CheckAuthenticationResponse CheckAuthentication(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CheckAuthenticationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CheckAuthenticationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAuthenticationRequest** | [**optional.Interface of CheckAuthenticationRequest**](CheckAuthenticationRequest.md)|  | 

### Return type

[**CheckAuthenticationResponse**](CheckAuthenticationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> CreateAccountResponse CreateAccount(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountRequest** | [**optional.Interface of CreateAccountRequest**](CreateAccountRequest.md)|  | 

### Return type

[**CreateAccountResponse**](CreateAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccounts

> ReadAccountsResponse ReadAccounts(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readAccountsRequest** | [**optional.Interface of ReadAccountsRequest**](ReadAccountsRequest.md)|  | 

### Return type

[**ReadAccountsResponse**](ReadAccountsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadConsumptionAccount

> ReadConsumptionAccountResponse ReadConsumptionAccount(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadConsumptionAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadConsumptionAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readConsumptionAccountRequest** | [**optional.Interface of ReadConsumptionAccountRequest**](ReadConsumptionAccountRequest.md)|  | 

### Return type

[**ReadConsumptionAccountResponse**](ReadConsumptionAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## UpdateAccount

> UpdateAccountResponse UpdateAccount(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccountRequest** | [**optional.Interface of UpdateAccountRequest**](UpdateAccountRequest.md)|  | 

### Return type

[**UpdateAccountResponse**](UpdateAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

