# \QuotaApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadQuotas**](QuotaApi.md#ReadQuotas) | **Post** /ReadQuotas | 



## ReadQuotas

> ReadQuotasResponse ReadQuotas(ctx).ReadQuotasRequest(readQuotasRequest).Execute()



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
    readQuotasRequest := *openapiclient.NewReadQuotasRequest() // ReadQuotasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotaApi.ReadQuotas(context.Background()).ReadQuotasRequest(readQuotasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.ReadQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadQuotas`: ReadQuotasResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.ReadQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readQuotasRequest** | [**ReadQuotasRequest**](ReadQuotasRequest.md) |  | 

### Return type

[**ReadQuotasResponse**](ReadQuotasResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

