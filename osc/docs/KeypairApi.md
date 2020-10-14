# \KeypairApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeypair**](KeypairApi.md#CreateKeypair) | **Post** /CreateKeypair | 
[**DeleteKeypair**](KeypairApi.md#DeleteKeypair) | **Post** /DeleteKeypair | 
[**ReadKeypairs**](KeypairApi.md#ReadKeypairs) | **Post** /ReadKeypairs | 



## CreateKeypair

> CreateKeypairResponse CreateKeypair(ctx).CreateKeypairRequest(createKeypairRequest).Execute()



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
    createKeypairRequest := openapiclient.CreateKeypairRequest{DryRun: false, KeypairName: "KeypairName_example", PublicKey: "PublicKey_example"} // CreateKeypairRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeypairApi.CreateKeypair(context.Background()).CreateKeypairRequest(createKeypairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeypairApi.CreateKeypair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKeypair`: CreateKeypairResponse
    fmt.Fprintf(os.Stdout, "Response from `KeypairApi.CreateKeypair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKeypairRequest** | [**CreateKeypairRequest**](CreateKeypairRequest.md) |  | 

### Return type

[**CreateKeypairResponse**](CreateKeypairResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeypair

> DeleteKeypairResponse DeleteKeypair(ctx).DeleteKeypairRequest(deleteKeypairRequest).Execute()



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
    deleteKeypairRequest := openapiclient.DeleteKeypairRequest{DryRun: false, KeypairName: "KeypairName_example"} // DeleteKeypairRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeypairApi.DeleteKeypair(context.Background()).DeleteKeypairRequest(deleteKeypairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeypairApi.DeleteKeypair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKeypair`: DeleteKeypairResponse
    fmt.Fprintf(os.Stdout, "Response from `KeypairApi.DeleteKeypair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteKeypairRequest** | [**DeleteKeypairRequest**](DeleteKeypairRequest.md) |  | 

### Return type

[**DeleteKeypairResponse**](DeleteKeypairResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadKeypairs

> ReadKeypairsResponse ReadKeypairs(ctx).ReadKeypairsRequest(readKeypairsRequest).Execute()



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
    readKeypairsRequest := openapiclient.ReadKeypairsRequest{DryRun: false, Filters: openapiclient.FiltersKeypair{KeypairFingerprints: []string{"KeypairFingerprints_example"), KeypairNames: []string{"KeypairNames_example")}} // ReadKeypairsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeypairApi.ReadKeypairs(context.Background()).ReadKeypairsRequest(readKeypairsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeypairApi.ReadKeypairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadKeypairs`: ReadKeypairsResponse
    fmt.Fprintf(os.Stdout, "Response from `KeypairApi.ReadKeypairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadKeypairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readKeypairsRequest** | [**ReadKeypairsRequest**](ReadKeypairsRequest.md) |  | 

### Return type

[**ReadKeypairsResponse**](ReadKeypairsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

