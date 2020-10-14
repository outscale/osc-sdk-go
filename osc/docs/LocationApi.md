# \LocationApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadLocations**](LocationApi.md#ReadLocations) | **Post** /ReadLocations | 



## ReadLocations

> ReadLocationsResponse ReadLocations(ctx).ReadLocationsRequest(readLocationsRequest).Execute()



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
    readLocationsRequest := openapiclient.ReadLocationsRequest{DryRun: false} // ReadLocationsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LocationApi.ReadLocations(context.Background()).ReadLocationsRequest(readLocationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.ReadLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadLocations`: ReadLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationApi.ReadLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLocationsRequest** | [**ReadLocationsRequest**](ReadLocationsRequest.md) |  | 

### Return type

[**ReadLocationsResponse**](ReadLocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

