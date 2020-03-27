# \SecurityGroupApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityGroup**](SecurityGroupApi.md#CreateSecurityGroup) | **Post** /CreateSecurityGroup | 
[**DeleteSecurityGroup**](SecurityGroupApi.md#DeleteSecurityGroup) | **Post** /DeleteSecurityGroup | 
[**ReadSecurityGroups**](SecurityGroupApi.md#ReadSecurityGroups) | **Post** /ReadSecurityGroups | 



## CreateSecurityGroup

> CreateSecurityGroupResponse CreateSecurityGroup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSecurityGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSecurityGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSecurityGroupRequest** | [**optional.Interface of CreateSecurityGroupRequest**](CreateSecurityGroupRequest.md)|  | 

### Return type

[**CreateSecurityGroupResponse**](CreateSecurityGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityGroup

> DeleteSecurityGroupResponse DeleteSecurityGroup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteSecurityGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSecurityGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSecurityGroupRequest** | [**optional.Interface of DeleteSecurityGroupRequest**](DeleteSecurityGroupRequest.md)|  | 

### Return type

[**DeleteSecurityGroupResponse**](DeleteSecurityGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSecurityGroups

> ReadSecurityGroupsResponse ReadSecurityGroups(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadSecurityGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadSecurityGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSecurityGroupsRequest** | [**optional.Interface of ReadSecurityGroupsRequest**](ReadSecurityGroupsRequest.md)|  | 

### Return type

[**ReadSecurityGroupsResponse**](ReadSecurityGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

