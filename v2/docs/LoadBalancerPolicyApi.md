# \LoadBalancerPolicyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancerPolicy**](LoadBalancerPolicyApi.md#CreateLoadBalancerPolicy) | **Post** /CreateLoadBalancerPolicy | 
[**DeleteLoadBalancerPolicy**](LoadBalancerPolicyApi.md#DeleteLoadBalancerPolicy) | **Post** /DeleteLoadBalancerPolicy | 



## CreateLoadBalancerPolicy

> CreateLoadBalancerPolicyResponse CreateLoadBalancerPolicy(ctx).CreateLoadBalancerPolicyRequest(createLoadBalancerPolicyRequest).Execute()



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
    createLoadBalancerPolicyRequest := *openapiclient.NewCreateLoadBalancerPolicyRequest("LoadBalancerName_example", "PolicyName_example", "PolicyType_example") // CreateLoadBalancerPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerPolicyApi.CreateLoadBalancerPolicy(context.Background()).CreateLoadBalancerPolicyRequest(createLoadBalancerPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerPolicyApi.CreateLoadBalancerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancerPolicy`: CreateLoadBalancerPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerPolicyApi.CreateLoadBalancerPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerPolicyRequest** | [**CreateLoadBalancerPolicyRequest**](CreateLoadBalancerPolicyRequest.md) |  | 

### Return type

[**CreateLoadBalancerPolicyResponse**](CreateLoadBalancerPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancerPolicy

> DeleteLoadBalancerPolicyResponse DeleteLoadBalancerPolicy(ctx).DeleteLoadBalancerPolicyRequest(deleteLoadBalancerPolicyRequest).Execute()



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
    deleteLoadBalancerPolicyRequest := *openapiclient.NewDeleteLoadBalancerPolicyRequest("LoadBalancerName_example", "PolicyName_example") // DeleteLoadBalancerPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerPolicyApi.DeleteLoadBalancerPolicy(context.Background()).DeleteLoadBalancerPolicyRequest(deleteLoadBalancerPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerPolicyApi.DeleteLoadBalancerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancerPolicy`: DeleteLoadBalancerPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerPolicyApi.DeleteLoadBalancerPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerPolicyRequest** | [**DeleteLoadBalancerPolicyRequest**](DeleteLoadBalancerPolicyRequest.md) |  | 

### Return type

[**DeleteLoadBalancerPolicyResponse**](DeleteLoadBalancerPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

