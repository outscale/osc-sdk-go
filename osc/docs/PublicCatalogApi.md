# \PublicCatalogApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPublicCatalog**](PublicCatalogApi.md#ReadPublicCatalog) | **Post** /ReadPublicCatalog | 



## ReadPublicCatalog

> ReadPublicCatalogResponse ReadPublicCatalog(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadPublicCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadPublicCatalogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPublicCatalogRequest** | [**optional.Interface of ReadPublicCatalogRequest**](ReadPublicCatalogRequest.md)|  | 

### Return type

[**ReadPublicCatalogResponse**](ReadPublicCatalogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

