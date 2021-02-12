# \ListenerApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListenerRule**](ListenerApi.md#CreateListenerRule) | **Post** /CreateListenerRule | 
[**CreateLoadBalancerListeners**](ListenerApi.md#CreateLoadBalancerListeners) | **Post** /CreateLoadBalancerListeners | 
[**DeleteListenerRule**](ListenerApi.md#DeleteListenerRule) | **Post** /DeleteListenerRule | 
[**DeleteLoadBalancerListeners**](ListenerApi.md#DeleteLoadBalancerListeners) | **Post** /DeleteLoadBalancerListeners | 
[**ReadListenerRules**](ListenerApi.md#ReadListenerRules) | **Post** /ReadListenerRules | 
[**UpdateListenerRule**](ListenerApi.md#UpdateListenerRule) | **Post** /UpdateListenerRule | 



## CreateListenerRule

> CreateListenerRuleResponse CreateListenerRule(ctx).CreateListenerRuleRequest(createListenerRuleRequest).Execute()



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
    createListenerRuleRequest := *openapiclient.NewCreateListenerRuleRequest(*openapiclient.NewLoadBalancerLight("LoadBalancerName_example", int32(123)), *openapiclient.NewListenerRuleForCreation("ListenerRuleName_example", int32(123)), []string{"VmIds_example"}) // CreateListenerRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListenerApi.CreateListenerRule(context.Background()).CreateListenerRuleRequest(createListenerRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListenerApi.CreateListenerRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListenerRule`: CreateListenerRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ListenerApi.CreateListenerRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateListenerRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createListenerRuleRequest** | [**CreateListenerRuleRequest**](CreateListenerRuleRequest.md) |  | 

### Return type

[**CreateListenerRuleResponse**](CreateListenerRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancerListeners

> CreateLoadBalancerListenersResponse CreateLoadBalancerListeners(ctx).CreateLoadBalancerListenersRequest(createLoadBalancerListenersRequest).Execute()



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
    createLoadBalancerListenersRequest := *openapiclient.NewCreateLoadBalancerListenersRequest([]openapiclient.ListenerForCreation{*openapiclient.NewListenerForCreation(int32(123), int32(123), "LoadBalancerProtocol_example")}, "LoadBalancerName_example") // CreateLoadBalancerListenersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListenerApi.CreateLoadBalancerListeners(context.Background()).CreateLoadBalancerListenersRequest(createLoadBalancerListenersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListenerApi.CreateLoadBalancerListeners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancerListeners`: CreateLoadBalancerListenersResponse
    fmt.Fprintf(os.Stdout, "Response from `ListenerApi.CreateLoadBalancerListeners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerListenersRequest** | [**CreateLoadBalancerListenersRequest**](CreateLoadBalancerListenersRequest.md) |  | 

### Return type

[**CreateLoadBalancerListenersResponse**](CreateLoadBalancerListenersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListenerRule

> DeleteListenerRuleResponse DeleteListenerRule(ctx).DeleteListenerRuleRequest(deleteListenerRuleRequest).Execute()



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
    deleteListenerRuleRequest := *openapiclient.NewDeleteListenerRuleRequest("ListenerRuleName_example") // DeleteListenerRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListenerApi.DeleteListenerRule(context.Background()).DeleteListenerRuleRequest(deleteListenerRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListenerApi.DeleteListenerRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteListenerRule`: DeleteListenerRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ListenerApi.DeleteListenerRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListenerRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteListenerRuleRequest** | [**DeleteListenerRuleRequest**](DeleteListenerRuleRequest.md) |  | 

### Return type

[**DeleteListenerRuleResponse**](DeleteListenerRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancerListeners

> DeleteLoadBalancerListenersResponse DeleteLoadBalancerListeners(ctx).DeleteLoadBalancerListenersRequest(deleteLoadBalancerListenersRequest).Execute()



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
    deleteLoadBalancerListenersRequest := *openapiclient.NewDeleteLoadBalancerListenersRequest("LoadBalancerName_example", []int32{int32(123)}) // DeleteLoadBalancerListenersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListenerApi.DeleteLoadBalancerListeners(context.Background()).DeleteLoadBalancerListenersRequest(deleteLoadBalancerListenersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListenerApi.DeleteLoadBalancerListeners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancerListeners`: DeleteLoadBalancerListenersResponse
    fmt.Fprintf(os.Stdout, "Response from `ListenerApi.DeleteLoadBalancerListeners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerListenersRequest** | [**DeleteLoadBalancerListenersRequest**](DeleteLoadBalancerListenersRequest.md) |  | 

### Return type

[**DeleteLoadBalancerListenersResponse**](DeleteLoadBalancerListenersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadListenerRules

> ReadListenerRulesResponse ReadListenerRules(ctx).ReadListenerRulesRequest(readListenerRulesRequest).Execute()



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
    readListenerRulesRequest := *openapiclient.NewReadListenerRulesRequest() // ReadListenerRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListenerApi.ReadListenerRules(context.Background()).ReadListenerRulesRequest(readListenerRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListenerApi.ReadListenerRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadListenerRules`: ReadListenerRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `ListenerApi.ReadListenerRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadListenerRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readListenerRulesRequest** | [**ReadListenerRulesRequest**](ReadListenerRulesRequest.md) |  | 

### Return type

[**ReadListenerRulesResponse**](ReadListenerRulesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListenerRule

> UpdateListenerRuleResponse UpdateListenerRule(ctx).UpdateListenerRuleRequest(updateListenerRuleRequest).Execute()



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
    updateListenerRuleRequest := *openapiclient.NewUpdateListenerRuleRequest("ListenerRuleName_example") // UpdateListenerRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListenerApi.UpdateListenerRule(context.Background()).UpdateListenerRuleRequest(updateListenerRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListenerApi.UpdateListenerRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListenerRule`: UpdateListenerRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ListenerApi.UpdateListenerRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListenerRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateListenerRuleRequest** | [**UpdateListenerRuleRequest**](UpdateListenerRuleRequest.md) |  | 

### Return type

[**UpdateListenerRuleResponse**](UpdateListenerRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

