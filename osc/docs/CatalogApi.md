# \CatalogApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadCatalog**](CatalogApi.md#ReadCatalog) | **Post** /ReadCatalog | 
[**ReadCatalogs**](CatalogApi.md#ReadCatalogs) | **Post** /ReadCatalogs | 



## ReadCatalog

> ReadCatalogResponse ReadCatalog(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadCatalogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readCatalogRequest** | [**optional.Interface of ReadCatalogRequest**](ReadCatalogRequest.md)|  | 

### Return type

[**ReadCatalogResponse**](ReadCatalogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCatalogs

> ReadCatalogsResponse ReadCatalogs(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadCatalogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadCatalogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readCatalogsRequest** | [**optional.Interface of ReadCatalogsRequest**](ReadCatalogsRequest.md)|  | 

### Return type

[**ReadCatalogsResponse**](ReadCatalogsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

