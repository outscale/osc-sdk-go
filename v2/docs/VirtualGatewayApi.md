# \VirtualGatewayApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualGateway**](VirtualGatewayApi.md#CreateVirtualGateway) | **Post** /CreateVirtualGateway | 
[**DeleteVirtualGateway**](VirtualGatewayApi.md#DeleteVirtualGateway) | **Post** /DeleteVirtualGateway | 
[**LinkVirtualGateway**](VirtualGatewayApi.md#LinkVirtualGateway) | **Post** /LinkVirtualGateway | 
[**ReadVirtualGateways**](VirtualGatewayApi.md#ReadVirtualGateways) | **Post** /ReadVirtualGateways | 
[**UnlinkVirtualGateway**](VirtualGatewayApi.md#UnlinkVirtualGateway) | **Post** /UnlinkVirtualGateway | 
[**UpdateRoutePropagation**](VirtualGatewayApi.md#UpdateRoutePropagation) | **Post** /UpdateRoutePropagation | 



## CreateVirtualGateway

> CreateVirtualGatewayResponse CreateVirtualGateway(ctx).CreateVirtualGatewayRequest(createVirtualGatewayRequest).Execute()



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
    createVirtualGatewayRequest := *openapiclient.NewCreateVirtualGatewayRequest("ConnectionType_example") // CreateVirtualGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualGatewayApi.CreateVirtualGateway(context.Background()).CreateVirtualGatewayRequest(createVirtualGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualGatewayApi.CreateVirtualGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualGateway`: CreateVirtualGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualGatewayApi.CreateVirtualGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVirtualGatewayRequest** | [**CreateVirtualGatewayRequest**](CreateVirtualGatewayRequest.md) |  | 

### Return type

[**CreateVirtualGatewayResponse**](CreateVirtualGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualGateway

> DeleteVirtualGatewayResponse DeleteVirtualGateway(ctx).DeleteVirtualGatewayRequest(deleteVirtualGatewayRequest).Execute()



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
    deleteVirtualGatewayRequest := *openapiclient.NewDeleteVirtualGatewayRequest("VirtualGatewayId_example") // DeleteVirtualGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualGatewayApi.DeleteVirtualGateway(context.Background()).DeleteVirtualGatewayRequest(deleteVirtualGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualGatewayApi.DeleteVirtualGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualGateway`: DeleteVirtualGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualGatewayApi.DeleteVirtualGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVirtualGatewayRequest** | [**DeleteVirtualGatewayRequest**](DeleteVirtualGatewayRequest.md) |  | 

### Return type

[**DeleteVirtualGatewayResponse**](DeleteVirtualGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkVirtualGateway

> LinkVirtualGatewayResponse LinkVirtualGateway(ctx).LinkVirtualGatewayRequest(linkVirtualGatewayRequest).Execute()



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
    linkVirtualGatewayRequest := *openapiclient.NewLinkVirtualGatewayRequest("NetId_example", "VirtualGatewayId_example") // LinkVirtualGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualGatewayApi.LinkVirtualGateway(context.Background()).LinkVirtualGatewayRequest(linkVirtualGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualGatewayApi.LinkVirtualGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkVirtualGateway`: LinkVirtualGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualGatewayApi.LinkVirtualGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkVirtualGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkVirtualGatewayRequest** | [**LinkVirtualGatewayRequest**](LinkVirtualGatewayRequest.md) |  | 

### Return type

[**LinkVirtualGatewayResponse**](LinkVirtualGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVirtualGateways

> ReadVirtualGatewaysResponse ReadVirtualGateways(ctx).ReadVirtualGatewaysRequest(readVirtualGatewaysRequest).Execute()



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
    readVirtualGatewaysRequest := *openapiclient.NewReadVirtualGatewaysRequest() // ReadVirtualGatewaysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualGatewayApi.ReadVirtualGateways(context.Background()).ReadVirtualGatewaysRequest(readVirtualGatewaysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualGatewayApi.ReadVirtualGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVirtualGateways`: ReadVirtualGatewaysResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualGatewayApi.ReadVirtualGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVirtualGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVirtualGatewaysRequest** | [**ReadVirtualGatewaysRequest**](ReadVirtualGatewaysRequest.md) |  | 

### Return type

[**ReadVirtualGatewaysResponse**](ReadVirtualGatewaysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkVirtualGateway

> UnlinkVirtualGatewayResponse UnlinkVirtualGateway(ctx).UnlinkVirtualGatewayRequest(unlinkVirtualGatewayRequest).Execute()



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
    unlinkVirtualGatewayRequest := *openapiclient.NewUnlinkVirtualGatewayRequest("NetId_example", "VirtualGatewayId_example") // UnlinkVirtualGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualGatewayApi.UnlinkVirtualGateway(context.Background()).UnlinkVirtualGatewayRequest(unlinkVirtualGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualGatewayApi.UnlinkVirtualGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkVirtualGateway`: UnlinkVirtualGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualGatewayApi.UnlinkVirtualGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkVirtualGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkVirtualGatewayRequest** | [**UnlinkVirtualGatewayRequest**](UnlinkVirtualGatewayRequest.md) |  | 

### Return type

[**UnlinkVirtualGatewayResponse**](UnlinkVirtualGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoutePropagation

> UpdateRoutePropagationResponse UpdateRoutePropagation(ctx).UpdateRoutePropagationRequest(updateRoutePropagationRequest).Execute()



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
    updateRoutePropagationRequest := *openapiclient.NewUpdateRoutePropagationRequest(false, "RouteTableId_example", "VirtualGatewayId_example") // UpdateRoutePropagationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualGatewayApi.UpdateRoutePropagation(context.Background()).UpdateRoutePropagationRequest(updateRoutePropagationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualGatewayApi.UpdateRoutePropagation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoutePropagation`: UpdateRoutePropagationResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualGatewayApi.UpdateRoutePropagation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoutePropagationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRoutePropagationRequest** | [**UpdateRoutePropagationRequest**](UpdateRoutePropagationRequest.md) |  | 

### Return type

[**UpdateRoutePropagationResponse**](UpdateRoutePropagationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

