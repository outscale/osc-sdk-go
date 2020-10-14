# \TagApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTags**](TagApi.md#CreateTags) | **Post** /CreateTags | 
[**DeleteTags**](TagApi.md#DeleteTags) | **Post** /DeleteTags | 
[**ReadTags**](TagApi.md#ReadTags) | **Post** /ReadTags | 



## CreateTags

> CreateTagsResponse CreateTags(ctx).CreateTagsRequest(createTagsRequest).Execute()



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
    createTagsRequest := openapiclient.CreateTagsRequest{DryRun: false, ResourceIds: []string{"ResourceIds_example"), Tags: []ResourceTag{openapiclient.ResourceTag{Key: "Key_example", Value: "Value_example"})} // CreateTagsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.CreateTags(context.Background()).CreateTagsRequest(createTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.CreateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTags`: CreateTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TagApi.CreateTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagsRequest** | [**CreateTagsRequest**](CreateTagsRequest.md) |  | 

### Return type

[**CreateTagsResponse**](CreateTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTags

> DeleteTagsResponse DeleteTags(ctx).DeleteTagsRequest(deleteTagsRequest).Execute()



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
    deleteTagsRequest := openapiclient.DeleteTagsRequest{DryRun: false, ResourceIds: []string{"ResourceIds_example"), Tags: []ResourceTag{openapiclient.ResourceTag{Key: "Key_example", Value: "Value_example"})} // DeleteTagsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.DeleteTags(context.Background()).DeleteTagsRequest(deleteTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.DeleteTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTags`: DeleteTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TagApi.DeleteTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTagsRequest** | [**DeleteTagsRequest**](DeleteTagsRequest.md) |  | 

### Return type

[**DeleteTagsResponse**](DeleteTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTags

> ReadTagsResponse ReadTags(ctx).ReadTagsRequest(readTagsRequest).Execute()



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
    readTagsRequest := openapiclient.ReadTagsRequest{DryRun: false, Filters: openapiclient.FiltersTag{Keys: []string{"Keys_example"), ResourceIds: []string{"ResourceIds_example"), ResourceTypes: []string{"ResourceTypes_example"), Values: []string{"Values_example")}} // ReadTagsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagApi.ReadTags(context.Background()).ReadTagsRequest(readTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.ReadTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTags`: ReadTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TagApi.ReadTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readTagsRequest** | [**ReadTagsRequest**](ReadTagsRequest.md) |  | 

### Return type

[**ReadTagsResponse**](ReadTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

