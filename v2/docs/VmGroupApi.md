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

> CreateVmGroupResponse CreateVmGroup(ctx).CreateVmGroupRequest(createVmGroupRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createVmGroupRequest := *openapiclient.NewCreateVmGroupRequest([]string{"SecurityGroupIds_example"}, "SubnetId_example", int32(123), "VmGroupName_example", "VmTemplateId_example") // CreateVmGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmGroupApi.CreateVmGroup(context.Background()).CreateVmGroupRequest(createVmGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmGroupApi.CreateVmGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVmGroup`: CreateVmGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VmGroupApi.CreateVmGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVmGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVmGroupRequest** | [**CreateVmGroupRequest**](CreateVmGroupRequest.md) |  | 

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

> DeleteVmGroupResponse DeleteVmGroup(ctx).DeleteVmGroupRequest(deleteVmGroupRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteVmGroupRequest := *openapiclient.NewDeleteVmGroupRequest("VmGroupId_example") // DeleteVmGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmGroupApi.DeleteVmGroup(context.Background()).DeleteVmGroupRequest(deleteVmGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmGroupApi.DeleteVmGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVmGroup`: DeleteVmGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VmGroupApi.DeleteVmGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVmGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVmGroupRequest** | [**DeleteVmGroupRequest**](DeleteVmGroupRequest.md) |  | 

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

> ReadVmGroupsResponse ReadVmGroups(ctx).ReadVmGroupsRequest(readVmGroupsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    readVmGroupsRequest := *openapiclient.NewReadVmGroupsRequest() // ReadVmGroupsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmGroupApi.ReadVmGroups(context.Background()).ReadVmGroupsRequest(readVmGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmGroupApi.ReadVmGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVmGroups`: ReadVmGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `VmGroupApi.ReadVmGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVmGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmGroupsRequest** | [**ReadVmGroupsRequest**](ReadVmGroupsRequest.md) |  | 

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

> ScaleDownVmGroupResponse ScaleDownVmGroup(ctx).ScaleDownVmGroupRequest(scaleDownVmGroupRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    scaleDownVmGroupRequest := *openapiclient.NewScaleDownVmGroupRequest("VmGroupId_example", int32(123)) // ScaleDownVmGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmGroupApi.ScaleDownVmGroup(context.Background()).ScaleDownVmGroupRequest(scaleDownVmGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmGroupApi.ScaleDownVmGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScaleDownVmGroup`: ScaleDownVmGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VmGroupApi.ScaleDownVmGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScaleDownVmGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scaleDownVmGroupRequest** | [**ScaleDownVmGroupRequest**](ScaleDownVmGroupRequest.md) |  | 

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

> ScaleUpVmGroupResponse ScaleUpVmGroup(ctx).ScaleUpVmGroupRequest(scaleUpVmGroupRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    scaleUpVmGroupRequest := *openapiclient.NewScaleUpVmGroupRequest(int32(123), "VmGroupId_example") // ScaleUpVmGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmGroupApi.ScaleUpVmGroup(context.Background()).ScaleUpVmGroupRequest(scaleUpVmGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmGroupApi.ScaleUpVmGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScaleUpVmGroup`: ScaleUpVmGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VmGroupApi.ScaleUpVmGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScaleUpVmGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scaleUpVmGroupRequest** | [**ScaleUpVmGroupRequest**](ScaleUpVmGroupRequest.md) |  | 

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

> UpdateVmGroupResponse UpdateVmGroup(ctx).UpdateVmGroupRequest(updateVmGroupRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateVmGroupRequest := *openapiclient.NewUpdateVmGroupRequest("VmGroupId_example") // UpdateVmGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmGroupApi.UpdateVmGroup(context.Background()).UpdateVmGroupRequest(updateVmGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmGroupApi.UpdateVmGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVmGroup`: UpdateVmGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VmGroupApi.UpdateVmGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVmGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVmGroupRequest** | [**UpdateVmGroupRequest**](UpdateVmGroupRequest.md) |  | 

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

