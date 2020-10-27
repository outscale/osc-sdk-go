# \ProductTypeApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadProductTypes**](ProductTypeApi.md#ReadProductTypes) | **Post** /ReadProductTypes | 



## ReadProductTypes

> ReadProductTypesResponse ReadProductTypes(ctx).ReadProductTypesRequest(readProductTypesRequest).Execute()



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
    readProductTypesRequest := *openapiclient.NewReadProductTypesRequest() // ReadProductTypesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductTypeApi.ReadProductTypes(context.Background()).ReadProductTypesRequest(readProductTypesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeApi.ReadProductTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadProductTypes`: ReadProductTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductTypeApi.ReadProductTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadProductTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readProductTypesRequest** | [**ReadProductTypesRequest**](ReadProductTypesRequest.md) |  | 

### Return type

[**ReadProductTypesResponse**](ReadProductTypesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

