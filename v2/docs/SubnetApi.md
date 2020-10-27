# \SubnetApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnet**](SubnetApi.md#CreateSubnet) | **Post** /CreateSubnet | 
[**DeleteSubnet**](SubnetApi.md#DeleteSubnet) | **Post** /DeleteSubnet | 
[**ReadSubnets**](SubnetApi.md#ReadSubnets) | **Post** /ReadSubnets | 
[**UpdateSubnet**](SubnetApi.md#UpdateSubnet) | **Post** /UpdateSubnet | 



## CreateSubnet

> CreateSubnetResponse CreateSubnet(ctx).CreateSubnetRequest(createSubnetRequest).Execute()



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
    createSubnetRequest := *openapiclient.NewCreateSubnetRequest("IpRange_example", "NetId_example") // CreateSubnetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubnetApi.CreateSubnet(context.Background()).CreateSubnetRequest(createSubnetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubnetApi.CreateSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubnet`: CreateSubnetResponse
    fmt.Fprintf(os.Stdout, "Response from `SubnetApi.CreateSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubnetRequest** | [**CreateSubnetRequest**](CreateSubnetRequest.md) |  | 

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

> DeleteSubnetResponse DeleteSubnet(ctx).DeleteSubnetRequest(deleteSubnetRequest).Execute()



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
    deleteSubnetRequest := *openapiclient.NewDeleteSubnetRequest("SubnetId_example") // DeleteSubnetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubnetApi.DeleteSubnet(context.Background()).DeleteSubnetRequest(deleteSubnetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubnetApi.DeleteSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubnet`: DeleteSubnetResponse
    fmt.Fprintf(os.Stdout, "Response from `SubnetApi.DeleteSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSubnetRequest** | [**DeleteSubnetRequest**](DeleteSubnetRequest.md) |  | 

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

> ReadSubnetsResponse ReadSubnets(ctx).ReadSubnetsRequest(readSubnetsRequest).Execute()



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
    readSubnetsRequest := *openapiclient.NewReadSubnetsRequest() // ReadSubnetsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubnetApi.ReadSubnets(context.Background()).ReadSubnetsRequest(readSubnetsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubnetApi.ReadSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSubnets`: ReadSubnetsResponse
    fmt.Fprintf(os.Stdout, "Response from `SubnetApi.ReadSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSubnetsRequest** | [**ReadSubnetsRequest**](ReadSubnetsRequest.md) |  | 

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

> UpdateSubnetResponse UpdateSubnet(ctx).UpdateSubnetRequest(updateSubnetRequest).Execute()



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
    updateSubnetRequest := *openapiclient.NewUpdateSubnetRequest(false, "SubnetId_example") // UpdateSubnetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubnetApi.UpdateSubnet(context.Background()).UpdateSubnetRequest(updateSubnetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubnetApi.UpdateSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubnet`: UpdateSubnetResponse
    fmt.Fprintf(os.Stdout, "Response from `SubnetApi.UpdateSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSubnetRequest** | [**UpdateSubnetRequest**](UpdateSubnetRequest.md) |  | 

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

