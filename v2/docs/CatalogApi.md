# \CatalogApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadCatalog**](CatalogApi.md#ReadCatalog) | **Post** /ReadCatalog | 
[**ReadCatalogs**](CatalogApi.md#ReadCatalogs) | **Post** /ReadCatalogs | 
[**ReadUnitPrice**](CatalogApi.md#ReadUnitPrice) | **Post** /ReadUnitPrice | 



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


## ReadCatalogs

> ReadCatalogsResponse ReadCatalogs(ctx).ReadCatalogsRequest(readCatalogsRequest).Execute()





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
    readCatalogsRequest := *openapiclient.NewReadCatalogsRequest() // ReadCatalogsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ReadCatalogs(context.Background()).ReadCatalogsRequest(readCatalogsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ReadCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCatalogs`: ReadCatalogsResponse
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ReadCatalogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readCatalogsRequest** | [**ReadCatalogsRequest**](ReadCatalogsRequest.md) |  | 

### Return type

[**ReadCatalogsResponse**](ReadCatalogsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUnitPrice

> ReadUnitPriceResponse ReadUnitPrice(ctx).ReadUnitPriceRequest(readUnitPriceRequest).Execute()





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
    readUnitPriceRequest := *openapiclient.NewReadUnitPriceRequest("Operation_example", "Service_example", "Type_example") // ReadUnitPriceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ReadUnitPrice(context.Background()).ReadUnitPriceRequest(readUnitPriceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ReadUnitPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUnitPrice`: ReadUnitPriceResponse
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ReadUnitPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadUnitPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readUnitPriceRequest** | [**ReadUnitPriceRequest**](ReadUnitPriceRequest.md) |  | 

### Return type

[**ReadUnitPriceResponse**](ReadUnitPriceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

