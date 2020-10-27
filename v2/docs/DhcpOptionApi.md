# \DhcpOptionApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDhcpOptions**](DhcpOptionApi.md#CreateDhcpOptions) | **Post** /CreateDhcpOptions | 
[**DeleteDhcpOptions**](DhcpOptionApi.md#DeleteDhcpOptions) | **Post** /DeleteDhcpOptions | 
[**ReadDhcpOptions**](DhcpOptionApi.md#ReadDhcpOptions) | **Post** /ReadDhcpOptions | 



## CreateDhcpOptions

> CreateDhcpOptionsResponse CreateDhcpOptions(ctx).CreateDhcpOptionsRequest(createDhcpOptionsRequest).Execute()



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
    createDhcpOptionsRequest := *openapiclient.NewCreateDhcpOptionsRequest() // CreateDhcpOptionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DhcpOptionApi.CreateDhcpOptions(context.Background()).CreateDhcpOptionsRequest(createDhcpOptionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpOptionApi.CreateDhcpOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDhcpOptions`: CreateDhcpOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DhcpOptionApi.CreateDhcpOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDhcpOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDhcpOptionsRequest** | [**CreateDhcpOptionsRequest**](CreateDhcpOptionsRequest.md) |  | 

### Return type

[**CreateDhcpOptionsResponse**](CreateDhcpOptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDhcpOptions

> DeleteDhcpOptionsResponse DeleteDhcpOptions(ctx).DeleteDhcpOptionsRequest(deleteDhcpOptionsRequest).Execute()



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
    deleteDhcpOptionsRequest := *openapiclient.NewDeleteDhcpOptionsRequest("DhcpOptionsSetId_example") // DeleteDhcpOptionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DhcpOptionApi.DeleteDhcpOptions(context.Background()).DeleteDhcpOptionsRequest(deleteDhcpOptionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpOptionApi.DeleteDhcpOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDhcpOptions`: DeleteDhcpOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DhcpOptionApi.DeleteDhcpOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDhcpOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDhcpOptionsRequest** | [**DeleteDhcpOptionsRequest**](DeleteDhcpOptionsRequest.md) |  | 

### Return type

[**DeleteDhcpOptionsResponse**](DeleteDhcpOptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDhcpOptions

> ReadDhcpOptionsResponse ReadDhcpOptions(ctx).ReadDhcpOptionsRequest(readDhcpOptionsRequest).Execute()



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
    readDhcpOptionsRequest := *openapiclient.NewReadDhcpOptionsRequest() // ReadDhcpOptionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DhcpOptionApi.ReadDhcpOptions(context.Background()).ReadDhcpOptionsRequest(readDhcpOptionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpOptionApi.ReadDhcpOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDhcpOptions`: ReadDhcpOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DhcpOptionApi.ReadDhcpOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadDhcpOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readDhcpOptionsRequest** | [**ReadDhcpOptionsRequest**](ReadDhcpOptionsRequest.md) |  | 

### Return type

[**ReadDhcpOptionsResponse**](ReadDhcpOptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

