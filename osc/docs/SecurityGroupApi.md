# \SecurityGroupApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityGroup**](SecurityGroupApi.md#CreateSecurityGroup) | **Post** /CreateSecurityGroup | 
[**DeleteSecurityGroup**](SecurityGroupApi.md#DeleteSecurityGroup) | **Post** /DeleteSecurityGroup | 
[**ReadSecurityGroups**](SecurityGroupApi.md#ReadSecurityGroups) | **Post** /ReadSecurityGroups | 



## CreateSecurityGroup

> CreateSecurityGroupResponse CreateSecurityGroup(ctx).CreateSecurityGroupRequest(createSecurityGroupRequest).Execute()



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
    createSecurityGroupRequest := openapiclient.CreateSecurityGroupRequest{Description: "Description_example", DryRun: false, NetId: "NetId_example", SecurityGroupName: "SecurityGroupName_example"} // CreateSecurityGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupApi.CreateSecurityGroup(context.Background()).CreateSecurityGroupRequest(createSecurityGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupApi.CreateSecurityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityGroup`: CreateSecurityGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupApi.CreateSecurityGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSecurityGroupRequest** | [**CreateSecurityGroupRequest**](CreateSecurityGroupRequest.md) |  | 

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

> DeleteSecurityGroupResponse DeleteSecurityGroup(ctx).DeleteSecurityGroupRequest(deleteSecurityGroupRequest).Execute()



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
    deleteSecurityGroupRequest := openapiclient.DeleteSecurityGroupRequest{DryRun: false, SecurityGroupId: "SecurityGroupId_example", SecurityGroupName: "SecurityGroupName_example"} // DeleteSecurityGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupApi.DeleteSecurityGroup(context.Background()).DeleteSecurityGroupRequest(deleteSecurityGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupApi.DeleteSecurityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSecurityGroup`: DeleteSecurityGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupApi.DeleteSecurityGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSecurityGroupRequest** | [**DeleteSecurityGroupRequest**](DeleteSecurityGroupRequest.md) |  | 

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

> ReadSecurityGroupsResponse ReadSecurityGroups(ctx).ReadSecurityGroupsRequest(readSecurityGroupsRequest).Execute()



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
    readSecurityGroupsRequest := openapiclient.ReadSecurityGroupsRequest{DryRun: false, Filters: openapiclient.FiltersSecurityGroup{AccountIds: []string{"AccountIds_example"), NetIds: []string{"NetIds_example"), SecurityGroupIds: []string{"SecurityGroupIds_example"), SecurityGroupNames: []string{"SecurityGroupNames_example"), TagKeys: []string{"TagKeys_example"), TagValues: []string{"TagValues_example"), Tags: []string{"Tags_example")}} // ReadSecurityGroupsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupApi.ReadSecurityGroups(context.Background()).ReadSecurityGroupsRequest(readSecurityGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupApi.ReadSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSecurityGroups`: ReadSecurityGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupApi.ReadSecurityGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSecurityGroupsRequest** | [**ReadSecurityGroupsRequest**](ReadSecurityGroupsRequest.md) |  | 

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

