# \DedicatedGroupApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDedicatedGroup**](DedicatedGroupApi.md#CreateDedicatedGroup) | **Post** /CreateDedicatedGroup | 
[**DeleteDedicatedGroup**](DedicatedGroupApi.md#DeleteDedicatedGroup) | **Post** /DeleteDedicatedGroup | 
[**ReadDedicatedGroups**](DedicatedGroupApi.md#ReadDedicatedGroups) | **Post** /ReadDedicatedGroups | 
[**UpdateDedicatedGroup**](DedicatedGroupApi.md#UpdateDedicatedGroup) | **Post** /UpdateDedicatedGroup | 



## CreateDedicatedGroup

> CreateDedicatedGroupResponse CreateDedicatedGroup(ctx).CreateDedicatedGroupRequest(createDedicatedGroupRequest).Execute()



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
    createDedicatedGroupRequest := *openapiclient.NewCreateDedicatedGroupRequest(int32(123), "Name_example", "SubregionName_example") // CreateDedicatedGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DedicatedGroupApi.CreateDedicatedGroup(context.Background()).CreateDedicatedGroupRequest(createDedicatedGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedGroupApi.CreateDedicatedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDedicatedGroup`: CreateDedicatedGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `DedicatedGroupApi.CreateDedicatedGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDedicatedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDedicatedGroupRequest** | [**CreateDedicatedGroupRequest**](CreateDedicatedGroupRequest.md) |  | 

### Return type

[**CreateDedicatedGroupResponse**](CreateDedicatedGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDedicatedGroup

> DeleteDedicatedGroupResponse DeleteDedicatedGroup(ctx).DeleteDedicatedGroupRequest(deleteDedicatedGroupRequest).Execute()



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
    deleteDedicatedGroupRequest := *openapiclient.NewDeleteDedicatedGroupRequest("DedicatedGroupId_example") // DeleteDedicatedGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DedicatedGroupApi.DeleteDedicatedGroup(context.Background()).DeleteDedicatedGroupRequest(deleteDedicatedGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedGroupApi.DeleteDedicatedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDedicatedGroup`: DeleteDedicatedGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `DedicatedGroupApi.DeleteDedicatedGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDedicatedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDedicatedGroupRequest** | [**DeleteDedicatedGroupRequest**](DeleteDedicatedGroupRequest.md) |  | 

### Return type

[**DeleteDedicatedGroupResponse**](DeleteDedicatedGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDedicatedGroups

> ReadDedicatedGroupsResponse ReadDedicatedGroups(ctx).ReadDedicatedGroupsRequest(readDedicatedGroupsRequest).Execute()



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
    readDedicatedGroupsRequest := *openapiclient.NewReadDedicatedGroupsRequest() // ReadDedicatedGroupsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DedicatedGroupApi.ReadDedicatedGroups(context.Background()).ReadDedicatedGroupsRequest(readDedicatedGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedGroupApi.ReadDedicatedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDedicatedGroups`: ReadDedicatedGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `DedicatedGroupApi.ReadDedicatedGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadDedicatedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readDedicatedGroupsRequest** | [**ReadDedicatedGroupsRequest**](ReadDedicatedGroupsRequest.md) |  | 

### Return type

[**ReadDedicatedGroupsResponse**](ReadDedicatedGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDedicatedGroup

> UpdateDedicatedGroupResponse UpdateDedicatedGroup(ctx).UpdateDedicatedGroupRequest(updateDedicatedGroupRequest).Execute()



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
    updateDedicatedGroupRequest := *openapiclient.NewUpdateDedicatedGroupRequest("DedicatedGroupId_example", "Name_example") // UpdateDedicatedGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DedicatedGroupApi.UpdateDedicatedGroup(context.Background()).UpdateDedicatedGroupRequest(updateDedicatedGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedGroupApi.UpdateDedicatedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDedicatedGroup`: UpdateDedicatedGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `DedicatedGroupApi.UpdateDedicatedGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDedicatedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDedicatedGroupRequest** | [**UpdateDedicatedGroupRequest**](UpdateDedicatedGroupRequest.md) |  | 

### Return type

[**UpdateDedicatedGroupResponse**](UpdateDedicatedGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

