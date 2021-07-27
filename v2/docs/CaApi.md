# \CaApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCa**](CaApi.md#CreateCa) | **Post** /CreateCa | 
[**DeleteCa**](CaApi.md#DeleteCa) | **Post** /DeleteCa | 
[**ReadCas**](CaApi.md#ReadCas) | **Post** /ReadCas | 
[**UpdateCa**](CaApi.md#UpdateCa) | **Post** /UpdateCa | 



## CreateCa

> CreateCaResponse CreateCa(ctx).CreateCaRequest(createCaRequest).Execute()



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
    createCaRequest := *openapiclient.NewCreateCaRequest("CaPem_example") // CreateCaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaApi.CreateCa(context.Background()).CreateCaRequest(createCaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaApi.CreateCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCa`: CreateCaResponse
    fmt.Fprintf(os.Stdout, "Response from `CaApi.CreateCa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCaRequest** | [**CreateCaRequest**](CreateCaRequest.md) |  | 

### Return type

[**CreateCaResponse**](CreateCaResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCa

> DeleteCaResponse DeleteCa(ctx).DeleteCaRequest(deleteCaRequest).Execute()



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
    deleteCaRequest := *openapiclient.NewDeleteCaRequest("CaId_example") // DeleteCaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaApi.DeleteCa(context.Background()).DeleteCaRequest(deleteCaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaApi.DeleteCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCa`: DeleteCaResponse
    fmt.Fprintf(os.Stdout, "Response from `CaApi.DeleteCa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteCaRequest** | [**DeleteCaRequest**](DeleteCaRequest.md) |  | 

### Return type

[**DeleteCaResponse**](DeleteCaResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCas

> ReadCasResponse ReadCas(ctx).ReadCasRequest(readCasRequest).Execute()



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
    readCasRequest := *openapiclient.NewReadCasRequest() // ReadCasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaApi.ReadCas(context.Background()).ReadCasRequest(readCasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaApi.ReadCas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCas`: ReadCasResponse
    fmt.Fprintf(os.Stdout, "Response from `CaApi.ReadCas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadCasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readCasRequest** | [**ReadCasRequest**](ReadCasRequest.md) |  | 

### Return type

[**ReadCasResponse**](ReadCasResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCa

> UpdateCaResponse UpdateCa(ctx).UpdateCaRequest(updateCaRequest).Execute()



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
    updateCaRequest := *openapiclient.NewUpdateCaRequest("CaId_example") // UpdateCaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaApi.UpdateCa(context.Background()).UpdateCaRequest(updateCaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaApi.UpdateCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCa`: UpdateCaResponse
    fmt.Fprintf(os.Stdout, "Response from `CaApi.UpdateCa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCaRequest** | [**UpdateCaRequest**](UpdateCaRequest.md) |  | 

### Return type

[**UpdateCaResponse**](UpdateCaResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

