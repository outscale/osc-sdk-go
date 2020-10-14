# \ApiLogApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadApiLogs**](ApiLogApi.md#ReadApiLogs) | **Post** /ReadApiLogs | 



## ReadApiLogs

> ReadApiLogsResponse ReadApiLogs(ctx).ReadApiLogsRequest(readApiLogsRequest).Execute()



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
    readApiLogsRequest := openapiclient.ReadApiLogsRequest{DryRun: false, Filters: openapiclient.FiltersApiLog{QueryAccessKeys: []string{"QueryAccessKeys_example"), QueryApiNames: []string{"QueryApiNames_example"), QueryCallNames: []string{"QueryCallNames_example"), QueryDateAfter: time.Now(), QueryDateBefore: time.Now(), QueryIpAddresses: []string{"QueryIpAddresses_example"), QueryUserAgents: []string{"QueryUserAgents_example"), RequestIds: []string{"RequestIds_example"), ResponseStatusCodes: []int32{123)}, NextPageToken: "NextPageToken_example", ResultsPerPage: 123, With: openapiclient.With{AccountId: false, CallDuration: false, QueryAccessKey: false, QueryApiName: false, QueryApiVersion: false, QueryCallName: false, QueryDate: false, QueryHeaderRaw: false, QueryHeaderSize: false, QueryIpAddress: false, QueryPayloadRaw: false, QueryPayloadSize: false, QueryUserAgent: false, RequestId: false, ResponseSize: false, ResponseStatusCode: false}} // ReadApiLogsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiLogApi.ReadApiLogs(context.Background()).ReadApiLogsRequest(readApiLogsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiLogApi.ReadApiLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadApiLogs`: ReadApiLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiLogApi.ReadApiLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadApiLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readApiLogsRequest** | [**ReadApiLogsRequest**](ReadApiLogsRequest.md) |  | 

### Return type

[**ReadApiLogsResponse**](ReadApiLogsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

