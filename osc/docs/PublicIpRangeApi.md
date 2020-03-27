# \PublicIpRangeApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPublicIpRanges**](PublicIpRangeApi.md#ReadPublicIpRanges) | **Post** /ReadPublicIpRanges | 



## ReadPublicIpRanges

> ReadPublicIpRangesResponse ReadPublicIpRanges(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadPublicIpRangesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadPublicIpRangesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPublicIpRangesRequest** | [**optional.Interface of ReadPublicIpRangesRequest**](ReadPublicIpRangesRequest.md)|  | 

### Return type

[**ReadPublicIpRangesResponse**](ReadPublicIpRangesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

