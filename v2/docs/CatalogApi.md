# \CatalogApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadCatalog**](CatalogApi.md#ReadCatalog) | **Post** /ReadCatalog | 



## ReadCatalog

> ReadCatalogResponse ReadCatalog(ctx).ReadCatalogRequest(readCatalogRequest).Execute()



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
    readCatalogRequest := *openapiclient.NewReadCatalogRequest() // ReadCatalogRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ReadCatalog(context.Background()).ReadCatalogRequest(readCatalogRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ReadCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCatalog`: ReadCatalogResponse
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ReadCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readCatalogRequest** | [**ReadCatalogRequest**](ReadCatalogRequest.md) |  | 

### Return type

[**ReadCatalogResponse**](ReadCatalogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

