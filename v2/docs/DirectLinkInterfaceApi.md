# \DirectLinkInterfaceApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectLinkInterface**](DirectLinkInterfaceApi.md#CreateDirectLinkInterface) | **Post** /CreateDirectLinkInterface | 
[**DeleteDirectLinkInterface**](DirectLinkInterfaceApi.md#DeleteDirectLinkInterface) | **Post** /DeleteDirectLinkInterface | 
[**ReadDirectLinkInterfaces**](DirectLinkInterfaceApi.md#ReadDirectLinkInterfaces) | **Post** /ReadDirectLinkInterfaces | 
[**UpdateDirectLinkInterface**](DirectLinkInterfaceApi.md#UpdateDirectLinkInterface) | **Post** /UpdateDirectLinkInterface | 



## CreateDirectLinkInterface

> CreateDirectLinkInterfaceResponse CreateDirectLinkInterface(ctx).CreateDirectLinkInterfaceRequest(createDirectLinkInterfaceRequest).Execute()



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
    createDirectLinkInterfaceRequest := *openapiclient.NewCreateDirectLinkInterfaceRequest("DirectLinkId_example", *openapiclient.NewDirectLinkInterface(int32(123), "DirectLinkInterfaceName_example", "VirtualGatewayId_example", int32(123))) // CreateDirectLinkInterfaceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectLinkInterfaceApi.CreateDirectLinkInterface(context.Background()).CreateDirectLinkInterfaceRequest(createDirectLinkInterfaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectLinkInterfaceApi.CreateDirectLinkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDirectLinkInterface`: CreateDirectLinkInterfaceResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectLinkInterfaceApi.CreateDirectLinkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectLinkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDirectLinkInterfaceRequest** | [**CreateDirectLinkInterfaceRequest**](CreateDirectLinkInterfaceRequest.md) |  | 

### Return type

[**CreateDirectLinkInterfaceResponse**](CreateDirectLinkInterfaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDirectLinkInterface

> DeleteDirectLinkInterfaceResponse DeleteDirectLinkInterface(ctx).DeleteDirectLinkInterfaceRequest(deleteDirectLinkInterfaceRequest).Execute()



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
    deleteDirectLinkInterfaceRequest := *openapiclient.NewDeleteDirectLinkInterfaceRequest("DirectLinkInterfaceId_example") // DeleteDirectLinkInterfaceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectLinkInterfaceApi.DeleteDirectLinkInterface(context.Background()).DeleteDirectLinkInterfaceRequest(deleteDirectLinkInterfaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectLinkInterfaceApi.DeleteDirectLinkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDirectLinkInterface`: DeleteDirectLinkInterfaceResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectLinkInterfaceApi.DeleteDirectLinkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectLinkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDirectLinkInterfaceRequest** | [**DeleteDirectLinkInterfaceRequest**](DeleteDirectLinkInterfaceRequest.md) |  | 

### Return type

[**DeleteDirectLinkInterfaceResponse**](DeleteDirectLinkInterfaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDirectLinkInterfaces

> ReadDirectLinkInterfacesResponse ReadDirectLinkInterfaces(ctx).ReadDirectLinkInterfacesRequest(readDirectLinkInterfacesRequest).Execute()



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
    readDirectLinkInterfacesRequest := *openapiclient.NewReadDirectLinkInterfacesRequest() // ReadDirectLinkInterfacesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectLinkInterfaceApi.ReadDirectLinkInterfaces(context.Background()).ReadDirectLinkInterfacesRequest(readDirectLinkInterfacesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectLinkInterfaceApi.ReadDirectLinkInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDirectLinkInterfaces`: ReadDirectLinkInterfacesResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectLinkInterfaceApi.ReadDirectLinkInterfaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadDirectLinkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readDirectLinkInterfacesRequest** | [**ReadDirectLinkInterfacesRequest**](ReadDirectLinkInterfacesRequest.md) |  | 

### Return type

[**ReadDirectLinkInterfacesResponse**](ReadDirectLinkInterfacesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDirectLinkInterface

> UpdateDirectLinkInterfaceResponse UpdateDirectLinkInterface(ctx).UpdateDirectLinkInterfaceRequest(updateDirectLinkInterfaceRequest).Execute()



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
    updateDirectLinkInterfaceRequest := *openapiclient.NewUpdateDirectLinkInterfaceRequest("DirectLinkInterfaceId_example", int32(123)) // UpdateDirectLinkInterfaceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectLinkInterfaceApi.UpdateDirectLinkInterface(context.Background()).UpdateDirectLinkInterfaceRequest(updateDirectLinkInterfaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectLinkInterfaceApi.UpdateDirectLinkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDirectLinkInterface`: UpdateDirectLinkInterfaceResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectLinkInterfaceApi.UpdateDirectLinkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDirectLinkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDirectLinkInterfaceRequest** | [**UpdateDirectLinkInterfaceRequest**](UpdateDirectLinkInterfaceRequest.md) |  | 

### Return type

[**UpdateDirectLinkInterfaceResponse**](UpdateDirectLinkInterfaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

