# \NetAccessPointApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetAccessPoint**](NetAccessPointApi.md#CreateNetAccessPoint) | **Post** /CreateNetAccessPoint | 
[**DeleteNetAccessPoint**](NetAccessPointApi.md#DeleteNetAccessPoint) | **Post** /DeleteNetAccessPoint | 
[**ReadNetAccessPointServices**](NetAccessPointApi.md#ReadNetAccessPointServices) | **Post** /ReadNetAccessPointServices | 
[**ReadNetAccessPoints**](NetAccessPointApi.md#ReadNetAccessPoints) | **Post** /ReadNetAccessPoints | 
[**UpdateNetAccessPoint**](NetAccessPointApi.md#UpdateNetAccessPoint) | **Post** /UpdateNetAccessPoint | 



## CreateNetAccessPoint

> CreateNetAccessPointResponse CreateNetAccessPoint(ctx).CreateNetAccessPointRequest(createNetAccessPointRequest).Execute()



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
    createNetAccessPointRequest := *openapiclient.NewCreateNetAccessPointRequest("NetId_example", "ServiceName_example") // CreateNetAccessPointRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetAccessPointApi.CreateNetAccessPoint(context.Background()).CreateNetAccessPointRequest(createNetAccessPointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetAccessPointApi.CreateNetAccessPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetAccessPoint`: CreateNetAccessPointResponse
    fmt.Fprintf(os.Stdout, "Response from `NetAccessPointApi.CreateNetAccessPoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetAccessPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetAccessPointRequest** | [**CreateNetAccessPointRequest**](CreateNetAccessPointRequest.md) |  | 

### Return type

[**CreateNetAccessPointResponse**](CreateNetAccessPointResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetAccessPoint

> DeleteNetAccessPointResponse DeleteNetAccessPoint(ctx).DeleteNetAccessPointRequest(deleteNetAccessPointRequest).Execute()



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
    deleteNetAccessPointRequest := *openapiclient.NewDeleteNetAccessPointRequest("NetAccessPointId_example") // DeleteNetAccessPointRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetAccessPointApi.DeleteNetAccessPoint(context.Background()).DeleteNetAccessPointRequest(deleteNetAccessPointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetAccessPointApi.DeleteNetAccessPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetAccessPoint`: DeleteNetAccessPointResponse
    fmt.Fprintf(os.Stdout, "Response from `NetAccessPointApi.DeleteNetAccessPoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetAccessPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNetAccessPointRequest** | [**DeleteNetAccessPointRequest**](DeleteNetAccessPointRequest.md) |  | 

### Return type

[**DeleteNetAccessPointResponse**](DeleteNetAccessPointResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNetAccessPointServices

> ReadNetAccessPointServicesResponse ReadNetAccessPointServices(ctx).ReadNetAccessPointServicesRequest(readNetAccessPointServicesRequest).Execute()



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
    readNetAccessPointServicesRequest := *openapiclient.NewReadNetAccessPointServicesRequest() // ReadNetAccessPointServicesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetAccessPointApi.ReadNetAccessPointServices(context.Background()).ReadNetAccessPointServicesRequest(readNetAccessPointServicesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetAccessPointApi.ReadNetAccessPointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNetAccessPointServices`: ReadNetAccessPointServicesResponse
    fmt.Fprintf(os.Stdout, "Response from `NetAccessPointApi.ReadNetAccessPointServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadNetAccessPointServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNetAccessPointServicesRequest** | [**ReadNetAccessPointServicesRequest**](ReadNetAccessPointServicesRequest.md) |  | 

### Return type

[**ReadNetAccessPointServicesResponse**](ReadNetAccessPointServicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNetAccessPoints

> ReadNetAccessPointsResponse ReadNetAccessPoints(ctx).ReadNetAccessPointsRequest(readNetAccessPointsRequest).Execute()



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
    readNetAccessPointsRequest := *openapiclient.NewReadNetAccessPointsRequest() // ReadNetAccessPointsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetAccessPointApi.ReadNetAccessPoints(context.Background()).ReadNetAccessPointsRequest(readNetAccessPointsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetAccessPointApi.ReadNetAccessPoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNetAccessPoints`: ReadNetAccessPointsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetAccessPointApi.ReadNetAccessPoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadNetAccessPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNetAccessPointsRequest** | [**ReadNetAccessPointsRequest**](ReadNetAccessPointsRequest.md) |  | 

### Return type

[**ReadNetAccessPointsResponse**](ReadNetAccessPointsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetAccessPoint

> UpdateNetAccessPointResponse UpdateNetAccessPoint(ctx).UpdateNetAccessPointRequest(updateNetAccessPointRequest).Execute()



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
    updateNetAccessPointRequest := *openapiclient.NewUpdateNetAccessPointRequest("NetAccessPointId_example") // UpdateNetAccessPointRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetAccessPointApi.UpdateNetAccessPoint(context.Background()).UpdateNetAccessPointRequest(updateNetAccessPointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetAccessPointApi.UpdateNetAccessPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetAccessPoint`: UpdateNetAccessPointResponse
    fmt.Fprintf(os.Stdout, "Response from `NetAccessPointApi.UpdateNetAccessPoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetAccessPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNetAccessPointRequest** | [**UpdateNetAccessPointRequest**](UpdateNetAccessPointRequest.md) |  | 

### Return type

[**UpdateNetAccessPointResponse**](UpdateNetAccessPointResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

