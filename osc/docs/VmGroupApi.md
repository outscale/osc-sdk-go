# \VmGroupApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVmGroup**](VmGroupApi.md#CreateVmGroup) | **Post** /CreateVmGroup | 
[**DeleteVmGroup**](VmGroupApi.md#DeleteVmGroup) | **Post** /DeleteVmGroup | 
[**ReadVmGroups**](VmGroupApi.md#ReadVmGroups) | **Post** /ReadVmGroups | 
[**ScaleDownVmGroup**](VmGroupApi.md#ScaleDownVmGroup) | **Post** /ScaleDownVmGroup | 
[**ScaleUpVmGroup**](VmGroupApi.md#ScaleUpVmGroup) | **Post** /ScaleUpVmGroup | 
[**UpdateVmGroup**](VmGroupApi.md#UpdateVmGroup) | **Post** /UpdateVmGroup | 



## CreateVmGroup

> CreateVmGroupResponse CreateVmGroup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateVmGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVmGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVmGroupRequest** | [**optional.Interface of CreateVmGroupRequest**](CreateVmGroupRequest.md)|  | 

### Return type

[**CreateVmGroupResponse**](CreateVmGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVmGroup

> DeleteVmGroupResponse DeleteVmGroup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteVmGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVmGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVmGroupRequest** | [**optional.Interface of DeleteVmGroupRequest**](DeleteVmGroupRequest.md)|  | 

### Return type

[**DeleteVmGroupResponse**](DeleteVmGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmGroups

> ReadVmGroupsResponse ReadVmGroups(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVmGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVmGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmGroupsRequest** | [**optional.Interface of ReadVmGroupsRequest**](ReadVmGroupsRequest.md)|  | 

### Return type

[**ReadVmGroupsResponse**](ReadVmGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScaleDownVmGroup

> ScaleDownVmGroupResponse ScaleDownVmGroup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScaleDownVmGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScaleDownVmGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scaleDownVmGroupRequest** | [**optional.Interface of ScaleDownVmGroupRequest**](ScaleDownVmGroupRequest.md)|  | 

### Return type

[**ScaleDownVmGroupResponse**](ScaleDownVmGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScaleUpVmGroup

> ScaleUpVmGroupResponse ScaleUpVmGroup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScaleUpVmGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScaleUpVmGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scaleUpVmGroupRequest** | [**optional.Interface of ScaleUpVmGroupRequest**](ScaleUpVmGroupRequest.md)|  | 

### Return type

[**ScaleUpVmGroupResponse**](ScaleUpVmGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVmGroup

> UpdateVmGroupResponse UpdateVmGroup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateVmGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateVmGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVmGroupRequest** | [**optional.Interface of UpdateVmGroupRequest**](UpdateVmGroupRequest.md)|  | 

### Return type

[**UpdateVmGroupResponse**](UpdateVmGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

