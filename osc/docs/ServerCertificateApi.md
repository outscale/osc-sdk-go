# \ServerCertificateApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerCertificate**](ServerCertificateApi.md#CreateServerCertificate) | **Post** /CreateServerCertificate | 
[**DeleteServerCertificate**](ServerCertificateApi.md#DeleteServerCertificate) | **Post** /DeleteServerCertificate | 
[**ReadServerCertificates**](ServerCertificateApi.md#ReadServerCertificates) | **Post** /ReadServerCertificates | 
[**UpdateServerCertificate**](ServerCertificateApi.md#UpdateServerCertificate) | **Post** /UpdateServerCertificate | 



## CreateServerCertificate

> CreateServerCertificateResponse CreateServerCertificate(ctx).CreateServerCertificateRequest(createServerCertificateRequest).Execute()



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
    createServerCertificateRequest := openapiclient.CreateServerCertificateRequest{Body: "Body_example", Chain: "Chain_example", DryRun: false, Name: "Name_example", Path: "Path_example", PrivateKey: "PrivateKey_example"} // CreateServerCertificateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerCertificateApi.CreateServerCertificate(context.Background()).CreateServerCertificateRequest(createServerCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerCertificateApi.CreateServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServerCertificate`: CreateServerCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerCertificateApi.CreateServerCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerCertificateRequest** | [**CreateServerCertificateRequest**](CreateServerCertificateRequest.md) |  | 

### Return type

[**CreateServerCertificateResponse**](CreateServerCertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerCertificate

> DeleteServerCertificateResponse DeleteServerCertificate(ctx).DeleteServerCertificateRequest(deleteServerCertificateRequest).Execute()



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
    deleteServerCertificateRequest := openapiclient.DeleteServerCertificateRequest{DryRun: false, Name: "Name_example"} // DeleteServerCertificateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerCertificateApi.DeleteServerCertificate(context.Background()).DeleteServerCertificateRequest(deleteServerCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerCertificateApi.DeleteServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServerCertificate`: DeleteServerCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerCertificateApi.DeleteServerCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteServerCertificateRequest** | [**DeleteServerCertificateRequest**](DeleteServerCertificateRequest.md) |  | 

### Return type

[**DeleteServerCertificateResponse**](DeleteServerCertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadServerCertificates

> ReadServerCertificatesResponse ReadServerCertificates(ctx).ReadServerCertificatesRequest(readServerCertificatesRequest).Execute()



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
    readServerCertificatesRequest := openapiclient.ReadServerCertificatesRequest{DryRun: false, Filters: openapiclient.FiltersServerCertificate{Paths: "Paths_example"}} // ReadServerCertificatesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerCertificateApi.ReadServerCertificates(context.Background()).ReadServerCertificatesRequest(readServerCertificatesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerCertificateApi.ReadServerCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadServerCertificates`: ReadServerCertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerCertificateApi.ReadServerCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadServerCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readServerCertificatesRequest** | [**ReadServerCertificatesRequest**](ReadServerCertificatesRequest.md) |  | 

### Return type

[**ReadServerCertificatesResponse**](ReadServerCertificatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerCertificate

> UpdateServerCertificateResponse UpdateServerCertificate(ctx).UpdateServerCertificateRequest(updateServerCertificateRequest).Execute()



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
    updateServerCertificateRequest := openapiclient.UpdateServerCertificateRequest{DryRun: false, Name: "Name_example", NewName: "NewName_example", NewPath: "NewPath_example"} // UpdateServerCertificateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServerCertificateApi.UpdateServerCertificate(context.Background()).UpdateServerCertificateRequest(updateServerCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerCertificateApi.UpdateServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerCertificate`: UpdateServerCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerCertificateApi.UpdateServerCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateServerCertificateRequest** | [**UpdateServerCertificateRequest**](UpdateServerCertificateRequest.md) |  | 

### Return type

[**UpdateServerCertificateResponse**](UpdateServerCertificateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

