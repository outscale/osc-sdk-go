# \RouteTableApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteTable**](RouteTableApi.md#CreateRouteTable) | **Post** /CreateRouteTable | 
[**DeleteRouteTable**](RouteTableApi.md#DeleteRouteTable) | **Post** /DeleteRouteTable | 
[**LinkRouteTable**](RouteTableApi.md#LinkRouteTable) | **Post** /LinkRouteTable | 
[**ReadRouteTables**](RouteTableApi.md#ReadRouteTables) | **Post** /ReadRouteTables | 
[**UnlinkRouteTable**](RouteTableApi.md#UnlinkRouteTable) | **Post** /UnlinkRouteTable | 
[**UpdateRouteTableLink**](RouteTableApi.md#UpdateRouteTableLink) | **Post** /UpdateRouteTableLink | 



## CreateRouteTable

> CreateRouteTableResponse CreateRouteTable(ctx).CreateRouteTableRequest(createRouteTableRequest).Execute()



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
    createRouteTableRequest := *openapiclient.NewCreateRouteTableRequest("NetId_example") // CreateRouteTableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteTableApi.CreateRouteTable(context.Background()).CreateRouteTableRequest(createRouteTableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteTableApi.CreateRouteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRouteTable`: CreateRouteTableResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteTableApi.CreateRouteTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteTableRequest** | [**CreateRouteTableRequest**](CreateRouteTableRequest.md) |  | 

### Return type

[**CreateRouteTableResponse**](CreateRouteTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteTable

> DeleteRouteTableResponse DeleteRouteTable(ctx).DeleteRouteTableRequest(deleteRouteTableRequest).Execute()



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
    deleteRouteTableRequest := *openapiclient.NewDeleteRouteTableRequest("RouteTableId_example") // DeleteRouteTableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteTableApi.DeleteRouteTable(context.Background()).DeleteRouteTableRequest(deleteRouteTableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteTableApi.DeleteRouteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRouteTable`: DeleteRouteTableResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteTableApi.DeleteRouteTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRouteTableRequest** | [**DeleteRouteTableRequest**](DeleteRouteTableRequest.md) |  | 

### Return type

[**DeleteRouteTableResponse**](DeleteRouteTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkRouteTable

> LinkRouteTableResponse LinkRouteTable(ctx).LinkRouteTableRequest(linkRouteTableRequest).Execute()



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
    linkRouteTableRequest := *openapiclient.NewLinkRouteTableRequest("RouteTableId_example", "SubnetId_example") // LinkRouteTableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteTableApi.LinkRouteTable(context.Background()).LinkRouteTableRequest(linkRouteTableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteTableApi.LinkRouteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkRouteTable`: LinkRouteTableResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteTableApi.LinkRouteTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkRouteTableRequest** | [**LinkRouteTableRequest**](LinkRouteTableRequest.md) |  | 

### Return type

[**LinkRouteTableResponse**](LinkRouteTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadRouteTables

> ReadRouteTablesResponse ReadRouteTables(ctx).ReadRouteTablesRequest(readRouteTablesRequest).Execute()



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
    readRouteTablesRequest := *openapiclient.NewReadRouteTablesRequest() // ReadRouteTablesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteTableApi.ReadRouteTables(context.Background()).ReadRouteTablesRequest(readRouteTablesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteTableApi.ReadRouteTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRouteTables`: ReadRouteTablesResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteTableApi.ReadRouteTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadRouteTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readRouteTablesRequest** | [**ReadRouteTablesRequest**](ReadRouteTablesRequest.md) |  | 

### Return type

[**ReadRouteTablesResponse**](ReadRouteTablesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkRouteTable

> UnlinkRouteTableResponse UnlinkRouteTable(ctx).UnlinkRouteTableRequest(unlinkRouteTableRequest).Execute()



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
    unlinkRouteTableRequest := *openapiclient.NewUnlinkRouteTableRequest("LinkRouteTableId_example") // UnlinkRouteTableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteTableApi.UnlinkRouteTable(context.Background()).UnlinkRouteTableRequest(unlinkRouteTableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteTableApi.UnlinkRouteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkRouteTable`: UnlinkRouteTableResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteTableApi.UnlinkRouteTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkRouteTableRequest** | [**UnlinkRouteTableRequest**](UnlinkRouteTableRequest.md) |  | 

### Return type

[**UnlinkRouteTableResponse**](UnlinkRouteTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteTableLink

> UpdateRouteTableLinkResponse UpdateRouteTableLink(ctx).UpdateRouteTableLinkRequest(updateRouteTableLinkRequest).Execute()



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
    updateRouteTableLinkRequest := *openapiclient.NewUpdateRouteTableLinkRequest("LinkRouteTableId_example", "RouteTableId_example") // UpdateRouteTableLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteTableApi.UpdateRouteTableLink(context.Background()).UpdateRouteTableLinkRequest(updateRouteTableLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteTableApi.UpdateRouteTableLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRouteTableLink`: UpdateRouteTableLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteTableApi.UpdateRouteTableLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteTableLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRouteTableLinkRequest** | [**UpdateRouteTableLinkRequest**](UpdateRouteTableLinkRequest.md) |  | 

### Return type

[**UpdateRouteTableLinkResponse**](UpdateRouteTableLinkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

