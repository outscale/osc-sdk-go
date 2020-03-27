# \VmApi

All URIs are relative to *https://api.eu-west-2.outscale.com/oapi/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVms**](VmApi.md#CreateVms) | **Post** /CreateVms | 
[**DeleteVms**](VmApi.md#DeleteVms) | **Post** /DeleteVms | 
[**ReadAdminPassword**](VmApi.md#ReadAdminPassword) | **Post** /ReadAdminPassword | 
[**ReadConsoleOutput**](VmApi.md#ReadConsoleOutput) | **Post** /ReadConsoleOutput | 
[**ReadVmTypes**](VmApi.md#ReadVmTypes) | **Post** /ReadVmTypes | 
[**ReadVms**](VmApi.md#ReadVms) | **Post** /ReadVms | 
[**ReadVmsState**](VmApi.md#ReadVmsState) | **Post** /ReadVmsState | 
[**RebootVms**](VmApi.md#RebootVms) | **Post** /RebootVms | 
[**StartVms**](VmApi.md#StartVms) | **Post** /StartVms | 
[**StopVms**](VmApi.md#StopVms) | **Post** /StopVms | 
[**UpdateVm**](VmApi.md#UpdateVm) | **Post** /UpdateVm | 



## CreateVms

> CreateVmsResponse CreateVms(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVmsRequest** | [**optional.Interface of CreateVmsRequest**](CreateVmsRequest.md)|  | 

### Return type

[**CreateVmsResponse**](CreateVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVms

> DeleteVmsResponse DeleteVms(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVmsRequest** | [**optional.Interface of DeleteVmsRequest**](DeleteVmsRequest.md)|  | 

### Return type

[**DeleteVmsResponse**](DeleteVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAdminPassword

> ReadAdminPasswordResponse ReadAdminPassword(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadAdminPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadAdminPasswordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readAdminPasswordRequest** | [**optional.Interface of ReadAdminPasswordRequest**](ReadAdminPasswordRequest.md)|  | 

### Return type

[**ReadAdminPasswordResponse**](ReadAdminPasswordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadConsoleOutput

> ReadConsoleOutputResponse ReadConsoleOutput(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadConsoleOutputOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadConsoleOutputOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readConsoleOutputRequest** | [**optional.Interface of ReadConsoleOutputRequest**](ReadConsoleOutputRequest.md)|  | 

### Return type

[**ReadConsoleOutputResponse**](ReadConsoleOutputResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmTypes

> ReadVmTypesResponse ReadVmTypes(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVmTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVmTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmTypesRequest** | [**optional.Interface of ReadVmTypesRequest**](ReadVmTypesRequest.md)|  | 

### Return type

[**ReadVmTypesResponse**](ReadVmTypesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVms

> ReadVmsResponse ReadVms(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmsRequest** | [**optional.Interface of ReadVmsRequest**](ReadVmsRequest.md)|  | 

### Return type

[**ReadVmsResponse**](ReadVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmsState

> ReadVmsStateResponse ReadVmsState(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadVmsStateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadVmsStateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmsStateRequest** | [**optional.Interface of ReadVmsStateRequest**](ReadVmsStateRequest.md)|  | 

### Return type

[**ReadVmsStateResponse**](ReadVmsStateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootVms

> RebootVmsResponse RebootVms(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RebootVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RebootVmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rebootVmsRequest** | [**optional.Interface of RebootVmsRequest**](RebootVmsRequest.md)|  | 

### Return type

[**RebootVmsResponse**](RebootVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVms

> StartVmsResponse StartVms(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StartVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StartVmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startVmsRequest** | [**optional.Interface of StartVmsRequest**](StartVmsRequest.md)|  | 

### Return type

[**StartVmsResponse**](StartVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVms

> StopVmsResponse StopVms(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StopVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StopVmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopVmsRequest** | [**optional.Interface of StopVmsRequest**](StopVmsRequest.md)|  | 

### Return type

[**StopVmsResponse**](StopVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVm

> UpdateVmResponse UpdateVm(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateVmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateVmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVmRequest** | [**optional.Interface of UpdateVmRequest**](UpdateVmRequest.md)|  | 

### Return type

[**UpdateVmResponse**](UpdateVmResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

