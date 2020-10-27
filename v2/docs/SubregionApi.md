# \SubregionApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadSubregions**](SubregionApi.md#ReadSubregions) | **Post** /ReadSubregions | 



## ReadSubregions

> ReadSubregionsResponse ReadSubregions(ctx).ReadSubregionsRequest(readSubregionsRequest).Execute()



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
    readSubregionsRequest := *openapiclient.NewReadSubregionsRequest() // ReadSubregionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubregionApi.ReadSubregions(context.Background()).ReadSubregionsRequest(readSubregionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubregionApi.ReadSubregions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSubregions`: ReadSubregionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SubregionApi.ReadSubregions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadSubregionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSubregionsRequest** | [**ReadSubregionsRequest**](ReadSubregionsRequest.md) |  | 

### Return type

[**ReadSubregionsResponse**](ReadSubregionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

