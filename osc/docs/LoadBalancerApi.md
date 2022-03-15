# \LoadBalancerApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancer**](LoadBalancerApi.md#CreateLoadBalancer) | **Post** /CreateLoadBalancer | 
[**CreateLoadBalancerTags**](LoadBalancerApi.md#CreateLoadBalancerTags) | **Post** /CreateLoadBalancerTags | 
[**DeleteLoadBalancer**](LoadBalancerApi.md#DeleteLoadBalancer) | **Post** /DeleteLoadBalancer | 
[**DeleteLoadBalancerTags**](LoadBalancerApi.md#DeleteLoadBalancerTags) | **Post** /DeleteLoadBalancerTags | 
[**DeregisterVmsInLoadBalancer**](LoadBalancerApi.md#DeregisterVmsInLoadBalancer) | **Post** /DeregisterVmsInLoadBalancer | 
[**LinkLoadBalancerBackendMachines**](LoadBalancerApi.md#LinkLoadBalancerBackendMachines) | **Post** /LinkLoadBalancerBackendMachines | 
[**ReadLoadBalancerTags**](LoadBalancerApi.md#ReadLoadBalancerTags) | **Post** /ReadLoadBalancerTags | 
[**ReadLoadBalancers**](LoadBalancerApi.md#ReadLoadBalancers) | **Post** /ReadLoadBalancers | 
[**ReadVmsHealth**](LoadBalancerApi.md#ReadVmsHealth) | **Post** /ReadVmsHealth | 
[**RegisterVmsInLoadBalancer**](LoadBalancerApi.md#RegisterVmsInLoadBalancer) | **Post** /RegisterVmsInLoadBalancer | 
[**UnlinkLoadBalancerBackendMachines**](LoadBalancerApi.md#UnlinkLoadBalancerBackendMachines) | **Post** /UnlinkLoadBalancerBackendMachines | 
[**UpdateLoadBalancer**](LoadBalancerApi.md#UpdateLoadBalancer) | **Post** /UpdateLoadBalancer | 



## CreateLoadBalancer

> CreateLoadBalancerResponse CreateLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerRequest** | [**optional.Interface of CreateLoadBalancerRequest**](CreateLoadBalancerRequest.md)|  | 

### Return type

[**CreateLoadBalancerResponse**](CreateLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancerTags

> CreateLoadBalancerTagsResponse CreateLoadBalancerTags(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateLoadBalancerTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLoadBalancerTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerTagsRequest** | [**optional.Interface of CreateLoadBalancerTagsRequest**](CreateLoadBalancerTagsRequest.md)|  | 

### Return type

[**CreateLoadBalancerTagsResponse**](CreateLoadBalancerTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancer

> DeleteLoadBalancerResponse DeleteLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerRequest** | [**optional.Interface of DeleteLoadBalancerRequest**](DeleteLoadBalancerRequest.md)|  | 

### Return type

[**DeleteLoadBalancerResponse**](DeleteLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancerTags

> DeleteLoadBalancerTagsResponse DeleteLoadBalancerTags(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteLoadBalancerTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteLoadBalancerTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerTagsRequest** | [**optional.Interface of DeleteLoadBalancerTagsRequest**](DeleteLoadBalancerTagsRequest.md)|  | 

### Return type

[**DeleteLoadBalancerTagsResponse**](DeleteLoadBalancerTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeregisterVmsInLoadBalancer

> DeregisterVmsInLoadBalancerResponse DeregisterVmsInLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeregisterVmsInLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeregisterVmsInLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deregisterVmsInLoadBalancerRequest** | [**optional.Interface of DeregisterVmsInLoadBalancerRequest**](DeregisterVmsInLoadBalancerRequest.md)|  | 

### Return type

[**DeregisterVmsInLoadBalancerResponse**](DeregisterVmsInLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkLoadBalancerBackendMachines

> LinkLoadBalancerBackendMachinesResponse LinkLoadBalancerBackendMachines(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinkLoadBalancerBackendMachinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LinkLoadBalancerBackendMachinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkLoadBalancerBackendMachinesRequest** | [**optional.Interface of LinkLoadBalancerBackendMachinesRequest**](LinkLoadBalancerBackendMachinesRequest.md)|  | 

### Return type

[**LinkLoadBalancerBackendMachinesResponse**](LinkLoadBalancerBackendMachinesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadLoadBalancerTags

> ReadLoadBalancerTagsResponse ReadLoadBalancerTags(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadLoadBalancerTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadLoadBalancerTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLoadBalancerTagsRequest** | [**optional.Interface of ReadLoadBalancerTagsRequest**](ReadLoadBalancerTagsRequest.md)|  | 

### Return type

[**ReadLoadBalancerTagsResponse**](ReadLoadBalancerTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadLoadBalancers

> ReadLoadBalancersResponse ReadLoadBalancers(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadLoadBalancersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadLoadBalancersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLoadBalancersRequest** | [**optional.Interface of ReadLoadBalancersRequest**](ReadLoadBalancersRequest.md)|  | 

### Return type

[**ReadLoadBalancersResponse**](ReadLoadBalancersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmsHealth

> ReadVmsHealthResponse ReadVmsHealth(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVmsHealthOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVmsHealthOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmsHealthRequest** | [**optional.Interface of ReadVmsHealthRequest**](ReadVmsHealthRequest.md)|  | 

### Return type

[**ReadVmsHealthResponse**](ReadVmsHealthResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterVmsInLoadBalancer

> RegisterVmsInLoadBalancerResponse RegisterVmsInLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterVmsInLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterVmsInLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerVmsInLoadBalancerRequest** | [**optional.Interface of RegisterVmsInLoadBalancerRequest**](RegisterVmsInLoadBalancerRequest.md)|  | 

### Return type

[**RegisterVmsInLoadBalancerResponse**](RegisterVmsInLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkLoadBalancerBackendMachines

> UnlinkLoadBalancerBackendMachinesResponse UnlinkLoadBalancerBackendMachines(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnlinkLoadBalancerBackendMachinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnlinkLoadBalancerBackendMachinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkLoadBalancerBackendMachinesRequest** | [**optional.Interface of UnlinkLoadBalancerBackendMachinesRequest**](UnlinkLoadBalancerBackendMachinesRequest.md)|  | 

### Return type

[**UnlinkLoadBalancerBackendMachinesResponse**](UnlinkLoadBalancerBackendMachinesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancer

> UpdateLoadBalancerResponse UpdateLoadBalancer(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateLoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateLoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLoadBalancerRequest** | [**optional.Interface of UpdateLoadBalancerRequest**](UpdateLoadBalancerRequest.md)|  | 

### Return type

[**UpdateLoadBalancerResponse**](UpdateLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

