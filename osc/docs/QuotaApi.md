# \QuotaApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadQuotas**](QuotaApi.md#ReadQuotas) | **Post** /ReadQuotas | 



## ReadQuotas

> ReadQuotasResponse ReadQuotas(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadQuotasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadQuotasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readQuotasRequest** | [**optional.Interface of ReadQuotasRequest**](ReadQuotasRequest.md)|  | 

### Return type

[**ReadQuotasResponse**](ReadQuotasResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

