# \SubnetApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnet**](SubnetApi.md#CreateSubnet) | **Post** /CreateSubnet | 
[**DeleteSubnet**](SubnetApi.md#DeleteSubnet) | **Post** /DeleteSubnet | 
[**ReadSubnets**](SubnetApi.md#ReadSubnets) | **Post** /ReadSubnets | 
[**UpdateSubnet**](SubnetApi.md#UpdateSubnet) | **Post** /UpdateSubnet | 



## CreateSubnet

> CreateSubnetResponse CreateSubnet(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSubnetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSubnetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubnetRequest** | [**optional.Interface of CreateSubnetRequest**](CreateSubnetRequest.md)|  | 

### Return type

[**CreateSubnetResponse**](CreateSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubnet

> DeleteSubnetResponse DeleteSubnet(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteSubnetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSubnetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSubnetRequest** | [**optional.Interface of DeleteSubnetRequest**](DeleteSubnetRequest.md)|  | 

### Return type

[**DeleteSubnetResponse**](DeleteSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSubnets

> ReadSubnetsResponse ReadSubnets(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadSubnetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadSubnetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSubnetsRequest** | [**optional.Interface of ReadSubnetsRequest**](ReadSubnetsRequest.md)|  | 

### Return type

[**ReadSubnetsResponse**](ReadSubnetsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubnet

> UpdateSubnetResponse UpdateSubnet(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateSubnetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSubnetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSubnetRequest** | [**optional.Interface of UpdateSubnetRequest**](UpdateSubnetRequest.md)|  | 

### Return type

[**UpdateSubnetResponse**](UpdateSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

