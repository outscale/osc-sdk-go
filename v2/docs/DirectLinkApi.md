# \DirectLinkApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectLink**](DirectLinkApi.md#CreateDirectLink) | **Post** /CreateDirectLink | 
[**DeleteDirectLink**](DirectLinkApi.md#DeleteDirectLink) | **Post** /DeleteDirectLink | 
[**ReadDirectLinks**](DirectLinkApi.md#ReadDirectLinks) | **Post** /ReadDirectLinks | 



## CreateDirectLink

> CreateDirectLinkResponse CreateDirectLink(ctx).CreateDirectLinkRequest(createDirectLinkRequest).Execute()



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
    createDirectLinkRequest := *openapiclient.NewCreateDirectLinkRequest("Bandwidth_example", "DirectLinkName_example", "Location_example") // CreateDirectLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectLinkApi.CreateDirectLink(context.Background()).CreateDirectLinkRequest(createDirectLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectLinkApi.CreateDirectLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDirectLink`: CreateDirectLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectLinkApi.CreateDirectLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDirectLinkRequest** | [**CreateDirectLinkRequest**](CreateDirectLinkRequest.md) |  | 

### Return type

[**CreateDirectLinkResponse**](CreateDirectLinkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDirectLink

> DeleteDirectLinkResponse DeleteDirectLink(ctx).DeleteDirectLinkRequest(deleteDirectLinkRequest).Execute()



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
    deleteDirectLinkRequest := *openapiclient.NewDeleteDirectLinkRequest("DirectLinkId_example") // DeleteDirectLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectLinkApi.DeleteDirectLink(context.Background()).DeleteDirectLinkRequest(deleteDirectLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectLinkApi.DeleteDirectLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDirectLink`: DeleteDirectLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectLinkApi.DeleteDirectLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDirectLinkRequest** | [**DeleteDirectLinkRequest**](DeleteDirectLinkRequest.md) |  | 

### Return type

[**DeleteDirectLinkResponse**](DeleteDirectLinkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDirectLinks

> ReadDirectLinksResponse ReadDirectLinks(ctx).ReadDirectLinksRequest(readDirectLinksRequest).Execute()



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
    readDirectLinksRequest := *openapiclient.NewReadDirectLinksRequest() // ReadDirectLinksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectLinkApi.ReadDirectLinks(context.Background()).ReadDirectLinksRequest(readDirectLinksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectLinkApi.ReadDirectLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDirectLinks`: ReadDirectLinksResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectLinkApi.ReadDirectLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadDirectLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readDirectLinksRequest** | [**ReadDirectLinksRequest**](ReadDirectLinksRequest.md) |  | 

### Return type

[**ReadDirectLinksResponse**](ReadDirectLinksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

