# \ClientGatewayApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientGateway**](ClientGatewayApi.md#CreateClientGateway) | **Post** /CreateClientGateway | 
[**DeleteClientGateway**](ClientGatewayApi.md#DeleteClientGateway) | **Post** /DeleteClientGateway | 
[**ReadClientGateways**](ClientGatewayApi.md#ReadClientGateways) | **Post** /ReadClientGateways | 



## CreateClientGateway

> CreateClientGatewayResponse CreateClientGateway(ctx).CreateClientGatewayRequest(createClientGatewayRequest).Execute()



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
    createClientGatewayRequest := *openapiclient.NewCreateClientGatewayRequest(int32(123), "ConnectionType_example", "PublicIp_example") // CreateClientGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientGatewayApi.CreateClientGateway(context.Background()).CreateClientGatewayRequest(createClientGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientGatewayApi.CreateClientGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClientGateway`: CreateClientGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientGatewayApi.CreateClientGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClientGatewayRequest** | [**CreateClientGatewayRequest**](CreateClientGatewayRequest.md) |  | 

### Return type

[**CreateClientGatewayResponse**](CreateClientGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientGateway

> DeleteClientGatewayResponse DeleteClientGateway(ctx).DeleteClientGatewayRequest(deleteClientGatewayRequest).Execute()



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
    deleteClientGatewayRequest := *openapiclient.NewDeleteClientGatewayRequest("ClientGatewayId_example") // DeleteClientGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientGatewayApi.DeleteClientGateway(context.Background()).DeleteClientGatewayRequest(deleteClientGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientGatewayApi.DeleteClientGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClientGateway`: DeleteClientGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientGatewayApi.DeleteClientGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteClientGatewayRequest** | [**DeleteClientGatewayRequest**](DeleteClientGatewayRequest.md) |  | 

### Return type

[**DeleteClientGatewayResponse**](DeleteClientGatewayResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadClientGateways

> ReadClientGatewaysResponse ReadClientGateways(ctx).ReadClientGatewaysRequest(readClientGatewaysRequest).Execute()



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
    readClientGatewaysRequest := *openapiclient.NewReadClientGatewaysRequest() // ReadClientGatewaysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientGatewayApi.ReadClientGateways(context.Background()).ReadClientGatewaysRequest(readClientGatewaysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientGatewayApi.ReadClientGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadClientGateways`: ReadClientGatewaysResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientGatewayApi.ReadClientGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadClientGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readClientGatewaysRequest** | [**ReadClientGatewaysRequest**](ReadClientGatewaysRequest.md) |  | 

### Return type

[**ReadClientGatewaysResponse**](ReadClientGatewaysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

