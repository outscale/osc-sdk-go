# \VpnConnectionApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVpnConnection**](VpnConnectionApi.md#CreateVpnConnection) | **Post** /CreateVpnConnection | 
[**CreateVpnConnectionRoute**](VpnConnectionApi.md#CreateVpnConnectionRoute) | **Post** /CreateVpnConnectionRoute | 
[**DeleteVpnConnection**](VpnConnectionApi.md#DeleteVpnConnection) | **Post** /DeleteVpnConnection | 
[**DeleteVpnConnectionRoute**](VpnConnectionApi.md#DeleteVpnConnectionRoute) | **Post** /DeleteVpnConnectionRoute | 
[**ReadVpnConnections**](VpnConnectionApi.md#ReadVpnConnections) | **Post** /ReadVpnConnections | 
[**UpdateVpnConnection**](VpnConnectionApi.md#UpdateVpnConnection) | **Post** /UpdateVpnConnection | 



## CreateVpnConnection

> CreateVpnConnectionResponse CreateVpnConnection(ctx).CreateVpnConnectionRequest(createVpnConnectionRequest).Execute()



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
    createVpnConnectionRequest := *openapiclient.NewCreateVpnConnectionRequest("ClientGatewayId_example", "ConnectionType_example", "VirtualGatewayId_example") // CreateVpnConnectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VpnConnectionApi.CreateVpnConnection(context.Background()).CreateVpnConnectionRequest(createVpnConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnConnectionApi.CreateVpnConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVpnConnection`: CreateVpnConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `VpnConnectionApi.CreateVpnConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVpnConnectionRequest** | [**CreateVpnConnectionRequest**](CreateVpnConnectionRequest.md) |  | 

### Return type

[**CreateVpnConnectionResponse**](CreateVpnConnectionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpnConnectionRoute

> CreateVpnConnectionRouteResponse CreateVpnConnectionRoute(ctx).CreateVpnConnectionRouteRequest(createVpnConnectionRouteRequest).Execute()



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
    createVpnConnectionRouteRequest := *openapiclient.NewCreateVpnConnectionRouteRequest("DestinationIpRange_example", "VpnConnectionId_example") // CreateVpnConnectionRouteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VpnConnectionApi.CreateVpnConnectionRoute(context.Background()).CreateVpnConnectionRouteRequest(createVpnConnectionRouteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnConnectionApi.CreateVpnConnectionRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVpnConnectionRoute`: CreateVpnConnectionRouteResponse
    fmt.Fprintf(os.Stdout, "Response from `VpnConnectionApi.CreateVpnConnectionRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnConnectionRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVpnConnectionRouteRequest** | [**CreateVpnConnectionRouteRequest**](CreateVpnConnectionRouteRequest.md) |  | 

### Return type

[**CreateVpnConnectionRouteResponse**](CreateVpnConnectionRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpnConnection

> DeleteVpnConnectionResponse DeleteVpnConnection(ctx).DeleteVpnConnectionRequest(deleteVpnConnectionRequest).Execute()



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
    deleteVpnConnectionRequest := *openapiclient.NewDeleteVpnConnectionRequest("VpnConnectionId_example") // DeleteVpnConnectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VpnConnectionApi.DeleteVpnConnection(context.Background()).DeleteVpnConnectionRequest(deleteVpnConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnConnectionApi.DeleteVpnConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVpnConnection`: DeleteVpnConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `VpnConnectionApi.DeleteVpnConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVpnConnectionRequest** | [**DeleteVpnConnectionRequest**](DeleteVpnConnectionRequest.md) |  | 

### Return type

[**DeleteVpnConnectionResponse**](DeleteVpnConnectionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpnConnectionRoute

> DeleteVpnConnectionRouteResponse DeleteVpnConnectionRoute(ctx).DeleteVpnConnectionRouteRequest(deleteVpnConnectionRouteRequest).Execute()



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
    deleteVpnConnectionRouteRequest := *openapiclient.NewDeleteVpnConnectionRouteRequest("DestinationIpRange_example", "VpnConnectionId_example") // DeleteVpnConnectionRouteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VpnConnectionApi.DeleteVpnConnectionRoute(context.Background()).DeleteVpnConnectionRouteRequest(deleteVpnConnectionRouteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnConnectionApi.DeleteVpnConnectionRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVpnConnectionRoute`: DeleteVpnConnectionRouteResponse
    fmt.Fprintf(os.Stdout, "Response from `VpnConnectionApi.DeleteVpnConnectionRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpnConnectionRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVpnConnectionRouteRequest** | [**DeleteVpnConnectionRouteRequest**](DeleteVpnConnectionRouteRequest.md) |  | 

### Return type

[**DeleteVpnConnectionRouteResponse**](DeleteVpnConnectionRouteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVpnConnections

> ReadVpnConnectionsResponse ReadVpnConnections(ctx).ReadVpnConnectionsRequest(readVpnConnectionsRequest).Execute()



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
    readVpnConnectionsRequest := *openapiclient.NewReadVpnConnectionsRequest() // ReadVpnConnectionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VpnConnectionApi.ReadVpnConnections(context.Background()).ReadVpnConnectionsRequest(readVpnConnectionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnConnectionApi.ReadVpnConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVpnConnections`: ReadVpnConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `VpnConnectionApi.ReadVpnConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVpnConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVpnConnectionsRequest** | [**ReadVpnConnectionsRequest**](ReadVpnConnectionsRequest.md) |  | 

### Return type

[**ReadVpnConnectionsResponse**](ReadVpnConnectionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVpnConnection

> UpdateVpnConnectionResponse UpdateVpnConnection(ctx).UpdateVpnConnectionRequest(updateVpnConnectionRequest).Execute()



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
    updateVpnConnectionRequest := *openapiclient.NewUpdateVpnConnectionRequest("VpnConnectionId_example") // UpdateVpnConnectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VpnConnectionApi.UpdateVpnConnection(context.Background()).UpdateVpnConnectionRequest(updateVpnConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnConnectionApi.UpdateVpnConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVpnConnection`: UpdateVpnConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `VpnConnectionApi.UpdateVpnConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVpnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVpnConnectionRequest** | [**UpdateVpnConnectionRequest**](UpdateVpnConnectionRequest.md) |  | 

### Return type

[**UpdateVpnConnectionResponse**](UpdateVpnConnectionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

