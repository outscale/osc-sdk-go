# \PublicCatalogApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPublicCatalog**](PublicCatalogApi.md#ReadPublicCatalog) | **Post** /ReadPublicCatalog | 



## ReadPublicCatalog

> ReadPublicCatalogResponse ReadPublicCatalog(ctx).ReadPublicCatalogRequest(readPublicCatalogRequest).Execute()



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
    readPublicCatalogRequest := *openapiclient.NewReadPublicCatalogRequest() // ReadPublicCatalogRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicCatalogApi.ReadPublicCatalog(context.Background()).ReadPublicCatalogRequest(readPublicCatalogRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicCatalogApi.ReadPublicCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPublicCatalog`: ReadPublicCatalogResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicCatalogApi.ReadPublicCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPublicCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPublicCatalogRequest** | [**ReadPublicCatalogRequest**](ReadPublicCatalogRequest.md) |  | 

### Return type

[**ReadPublicCatalogResponse**](ReadPublicCatalogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

