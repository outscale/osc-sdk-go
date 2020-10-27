# \NetPeeringApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptNetPeering**](NetPeeringApi.md#AcceptNetPeering) | **Post** /AcceptNetPeering | 
[**CreateNetPeering**](NetPeeringApi.md#CreateNetPeering) | **Post** /CreateNetPeering | 
[**DeleteNetPeering**](NetPeeringApi.md#DeleteNetPeering) | **Post** /DeleteNetPeering | 
[**ReadNetPeerings**](NetPeeringApi.md#ReadNetPeerings) | **Post** /ReadNetPeerings | 
[**RejectNetPeering**](NetPeeringApi.md#RejectNetPeering) | **Post** /RejectNetPeering | 



## AcceptNetPeering

> AcceptNetPeeringResponse AcceptNetPeering(ctx).AcceptNetPeeringRequest(acceptNetPeeringRequest).Execute()



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
    acceptNetPeeringRequest := *openapiclient.NewAcceptNetPeeringRequest("NetPeeringId_example") // AcceptNetPeeringRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetPeeringApi.AcceptNetPeering(context.Background()).AcceptNetPeeringRequest(acceptNetPeeringRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetPeeringApi.AcceptNetPeering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptNetPeering`: AcceptNetPeeringResponse
    fmt.Fprintf(os.Stdout, "Response from `NetPeeringApi.AcceptNetPeering`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptNetPeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptNetPeeringRequest** | [**AcceptNetPeeringRequest**](AcceptNetPeeringRequest.md) |  | 

### Return type

[**AcceptNetPeeringResponse**](AcceptNetPeeringResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetPeering

> CreateNetPeeringResponse CreateNetPeering(ctx).CreateNetPeeringRequest(createNetPeeringRequest).Execute()



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
    createNetPeeringRequest := *openapiclient.NewCreateNetPeeringRequest("AccepterNetId_example", "SourceNetId_example") // CreateNetPeeringRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetPeeringApi.CreateNetPeering(context.Background()).CreateNetPeeringRequest(createNetPeeringRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetPeeringApi.CreateNetPeering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetPeering`: CreateNetPeeringResponse
    fmt.Fprintf(os.Stdout, "Response from `NetPeeringApi.CreateNetPeering`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetPeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetPeeringRequest** | [**CreateNetPeeringRequest**](CreateNetPeeringRequest.md) |  | 

### Return type

[**CreateNetPeeringResponse**](CreateNetPeeringResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetPeering

> DeleteNetPeeringResponse DeleteNetPeering(ctx).DeleteNetPeeringRequest(deleteNetPeeringRequest).Execute()



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
    deleteNetPeeringRequest := *openapiclient.NewDeleteNetPeeringRequest("NetPeeringId_example") // DeleteNetPeeringRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetPeeringApi.DeleteNetPeering(context.Background()).DeleteNetPeeringRequest(deleteNetPeeringRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetPeeringApi.DeleteNetPeering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetPeering`: DeleteNetPeeringResponse
    fmt.Fprintf(os.Stdout, "Response from `NetPeeringApi.DeleteNetPeering`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetPeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNetPeeringRequest** | [**DeleteNetPeeringRequest**](DeleteNetPeeringRequest.md) |  | 

### Return type

[**DeleteNetPeeringResponse**](DeleteNetPeeringResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNetPeerings

> ReadNetPeeringsResponse ReadNetPeerings(ctx).ReadNetPeeringsRequest(readNetPeeringsRequest).Execute()



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
    readNetPeeringsRequest := *openapiclient.NewReadNetPeeringsRequest() // ReadNetPeeringsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetPeeringApi.ReadNetPeerings(context.Background()).ReadNetPeeringsRequest(readNetPeeringsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetPeeringApi.ReadNetPeerings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNetPeerings`: ReadNetPeeringsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetPeeringApi.ReadNetPeerings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadNetPeeringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNetPeeringsRequest** | [**ReadNetPeeringsRequest**](ReadNetPeeringsRequest.md) |  | 

### Return type

[**ReadNetPeeringsResponse**](ReadNetPeeringsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectNetPeering

> RejectNetPeeringResponse RejectNetPeering(ctx).RejectNetPeeringRequest(rejectNetPeeringRequest).Execute()



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
    rejectNetPeeringRequest := *openapiclient.NewRejectNetPeeringRequest("NetPeeringId_example") // RejectNetPeeringRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetPeeringApi.RejectNetPeering(context.Background()).RejectNetPeeringRequest(rejectNetPeeringRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetPeeringApi.RejectNetPeering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectNetPeering`: RejectNetPeeringResponse
    fmt.Fprintf(os.Stdout, "Response from `NetPeeringApi.RejectNetPeering`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRejectNetPeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rejectNetPeeringRequest** | [**RejectNetPeeringRequest**](RejectNetPeeringRequest.md) |  | 

### Return type

[**RejectNetPeeringResponse**](RejectNetPeeringResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

