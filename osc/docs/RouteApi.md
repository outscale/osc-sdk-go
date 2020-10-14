# \RouteApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoute**](RouteApi.md#CreateRoute) | **Post** /CreateRoute | 
[**DeleteRoute**](RouteApi.md#DeleteRoute) | **Post** /DeleteRoute | 
[**UpdateRoute**](RouteApi.md#UpdateRoute) | **Post** /UpdateRoute | 



## CreateRoute

> CreateRouteResponse CreateRoute(ctx).CreateRouteRequest(createRouteRequest).Execute()



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
    createRouteRequest := openapiclient.CreateRouteRequest{DestinationIpRange: "DestinationIpRange_example", DryRun: false, GatewayId: "GatewayId_example", NatServiceId: "NatServiceId_example", NetPeeringId: "NetPeeringId_example", NicId: "NicId_example", RouteTableId: "RouteTableId_example", VmId: "VmId_example"} // CreateRouteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.CreateRoute(context.Background()).CreateRouteRequest(createRouteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.CreateRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoute`: CreateRouteResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.CreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteRequest** | [**CreateRouteRequest**](CreateRouteRequest.md) |  | 

### Return type

[**CreateRouteResponse**](CreateRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteRouteResponse DeleteRoute(ctx).DeleteRouteRequest(deleteRouteRequest).Execute()



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
    deleteRouteRequest := openapiclient.DeleteRouteRequest{DestinationIpRange: "DestinationIpRange_example", DryRun: false, RouteTableId: "RouteTableId_example"} // DeleteRouteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.DeleteRoute(context.Background()).DeleteRouteRequest(deleteRouteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.DeleteRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoute`: DeleteRouteResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.DeleteRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRouteRequest** | [**DeleteRouteRequest**](DeleteRouteRequest.md) |  | 

### Return type

[**DeleteRouteResponse**](DeleteRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoute

> UpdateRouteResponse UpdateRoute(ctx).UpdateRouteRequest(updateRouteRequest).Execute()



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
    updateRouteRequest := openapiclient.UpdateRouteRequest{DestinationIpRange: "DestinationIpRange_example", DryRun: false, GatewayId: "GatewayId_example", NatServiceId: "NatServiceId_example", NetPeeringId: "NetPeeringId_example", NicId: "NicId_example", RouteTableId: "RouteTableId_example", VmId: "VmId_example"} // UpdateRouteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.UpdateRoute(context.Background()).UpdateRouteRequest(updateRouteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.UpdateRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoute`: UpdateRouteResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRouteRequest** | [**UpdateRouteRequest**](UpdateRouteRequest.md) |  | 

### Return type

[**UpdateRouteResponse**](UpdateRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

