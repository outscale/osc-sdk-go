# \ApiAccessPolicyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadApiAccessPolicy**](ApiAccessPolicyApi.md#ReadApiAccessPolicy) | **Post** /ReadApiAccessPolicy | 
[**UpdateApiAccessPolicy**](ApiAccessPolicyApi.md#UpdateApiAccessPolicy) | **Post** /UpdateApiAccessPolicy | 



## ReadApiAccessPolicy

> ReadApiAccessPolicyResponse ReadApiAccessPolicy(ctx).ReadApiAccessPolicyRequest(readApiAccessPolicyRequest).Execute()



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
    readApiAccessPolicyRequest := *openapiclient.NewReadApiAccessPolicyRequest() // ReadApiAccessPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiAccessPolicyApi.ReadApiAccessPolicy(context.Background()).ReadApiAccessPolicyRequest(readApiAccessPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAccessPolicyApi.ReadApiAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadApiAccessPolicy`: ReadApiAccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiAccessPolicyApi.ReadApiAccessPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadApiAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readApiAccessPolicyRequest** | [**ReadApiAccessPolicyRequest**](ReadApiAccessPolicyRequest.md) |  | 

### Return type

[**ReadApiAccessPolicyResponse**](ReadApiAccessPolicyResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiAccessPolicy

> UpdateApiAccessPolicyResponse UpdateApiAccessPolicy(ctx).UpdateApiAccessPolicyRequest(updateApiAccessPolicyRequest).Execute()



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
    updateApiAccessPolicyRequest := *openapiclient.NewUpdateApiAccessPolicyRequest(int64(123), false) // UpdateApiAccessPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiAccessPolicyApi.UpdateApiAccessPolicy(context.Background()).UpdateApiAccessPolicyRequest(updateApiAccessPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAccessPolicyApi.UpdateApiAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiAccessPolicy`: UpdateApiAccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiAccessPolicyApi.UpdateApiAccessPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateApiAccessPolicyRequest** | [**UpdateApiAccessPolicyRequest**](UpdateApiAccessPolicyRequest.md) |  | 

### Return type

[**UpdateApiAccessPolicyResponse**](UpdateApiAccessPolicyResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

