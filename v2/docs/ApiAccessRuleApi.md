# \ApiAccessRuleApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiAccessRule**](ApiAccessRuleApi.md#CreateApiAccessRule) | **Post** /CreateApiAccessRule | 
[**DeleteApiAccessRule**](ApiAccessRuleApi.md#DeleteApiAccessRule) | **Post** /DeleteApiAccessRule | 
[**ReadApiAccessRules**](ApiAccessRuleApi.md#ReadApiAccessRules) | **Post** /ReadApiAccessRules | 
[**UpdateApiAccessRule**](ApiAccessRuleApi.md#UpdateApiAccessRule) | **Post** /UpdateApiAccessRule | 



## CreateApiAccessRule

> CreateApiAccessRuleResponse CreateApiAccessRule(ctx).CreateApiAccessRuleRequest(createApiAccessRuleRequest).Execute()



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
    createApiAccessRuleRequest := *openapiclient.NewCreateApiAccessRuleRequest() // CreateApiAccessRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiAccessRuleApi.CreateApiAccessRule(context.Background()).CreateApiAccessRuleRequest(createApiAccessRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAccessRuleApi.CreateApiAccessRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiAccessRule`: CreateApiAccessRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiAccessRuleApi.CreateApiAccessRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiAccessRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiAccessRuleRequest** | [**CreateApiAccessRuleRequest**](CreateApiAccessRuleRequest.md) |  | 

### Return type

[**CreateApiAccessRuleResponse**](CreateApiAccessRuleResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiAccessRule

> DeleteApiAccessRuleResponse DeleteApiAccessRule(ctx).DeleteApiAccessRuleRequest(deleteApiAccessRuleRequest).Execute()



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
    deleteApiAccessRuleRequest := *openapiclient.NewDeleteApiAccessRuleRequest("ApiAccessRuleId_example") // DeleteApiAccessRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiAccessRuleApi.DeleteApiAccessRule(context.Background()).DeleteApiAccessRuleRequest(deleteApiAccessRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAccessRuleApi.DeleteApiAccessRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiAccessRule`: DeleteApiAccessRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiAccessRuleApi.DeleteApiAccessRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiAccessRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteApiAccessRuleRequest** | [**DeleteApiAccessRuleRequest**](DeleteApiAccessRuleRequest.md) |  | 

### Return type

[**DeleteApiAccessRuleResponse**](DeleteApiAccessRuleResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadApiAccessRules

> ReadApiAccessRulesResponse ReadApiAccessRules(ctx).ReadApiAccessRulesRequest(readApiAccessRulesRequest).Execute()



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
    readApiAccessRulesRequest := *openapiclient.NewReadApiAccessRulesRequest() // ReadApiAccessRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiAccessRuleApi.ReadApiAccessRules(context.Background()).ReadApiAccessRulesRequest(readApiAccessRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAccessRuleApi.ReadApiAccessRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadApiAccessRules`: ReadApiAccessRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiAccessRuleApi.ReadApiAccessRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadApiAccessRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readApiAccessRulesRequest** | [**ReadApiAccessRulesRequest**](ReadApiAccessRulesRequest.md) |  | 

### Return type

[**ReadApiAccessRulesResponse**](ReadApiAccessRulesResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiAccessRule

> UpdateApiAccessRuleResponse UpdateApiAccessRule(ctx).UpdateApiAccessRuleRequest(updateApiAccessRuleRequest).Execute()



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
    updateApiAccessRuleRequest := *openapiclient.NewUpdateApiAccessRuleRequest("ApiAccessRuleId_example") // UpdateApiAccessRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiAccessRuleApi.UpdateApiAccessRule(context.Background()).UpdateApiAccessRuleRequest(updateApiAccessRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAccessRuleApi.UpdateApiAccessRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiAccessRule`: UpdateApiAccessRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiAccessRuleApi.UpdateApiAccessRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiAccessRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateApiAccessRuleRequest** | [**UpdateApiAccessRuleRequest**](UpdateApiAccessRuleRequest.md) |  | 

### Return type

[**UpdateApiAccessRuleResponse**](UpdateApiAccessRuleResponse.md)

### Authorization

[ApiKeyAuthSec](../README.md#ApiKeyAuthSec), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

