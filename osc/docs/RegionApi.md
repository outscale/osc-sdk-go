# \RegionApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadRegions**](RegionApi.md#ReadRegions) | **Post** /ReadRegions | 



## ReadRegions

> ReadRegionsResponse ReadRegions(ctx).ReadRegionsRequest(readRegionsRequest).Execute()



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
    readRegionsRequest := openapiclient.ReadRegionsRequest{DryRun: false} // ReadRegionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionApi.ReadRegions(context.Background()).ReadRegionsRequest(readRegionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ReadRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRegions`: ReadRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RegionApi.ReadRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readRegionsRequest** | [**ReadRegionsRequest**](ReadRegionsRequest.md) |  | 

### Return type

[**ReadRegionsResponse**](ReadRegionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

