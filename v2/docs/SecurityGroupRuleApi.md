# \SecurityGroupRuleApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityGroupRule**](SecurityGroupRuleApi.md#CreateSecurityGroupRule) | **Post** /CreateSecurityGroupRule | 
[**DeleteSecurityGroupRule**](SecurityGroupRuleApi.md#DeleteSecurityGroupRule) | **Post** /DeleteSecurityGroupRule | 



## CreateSecurityGroupRule

> CreateSecurityGroupRuleResponse CreateSecurityGroupRule(ctx).CreateSecurityGroupRuleRequest(createSecurityGroupRuleRequest).Execute()



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
    createSecurityGroupRuleRequest := *openapiclient.NewCreateSecurityGroupRuleRequest("Flow_example", "SecurityGroupId_example") // CreateSecurityGroupRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupRuleApi.CreateSecurityGroupRule(context.Background()).CreateSecurityGroupRuleRequest(createSecurityGroupRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupRuleApi.CreateSecurityGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityGroupRule`: CreateSecurityGroupRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupRuleApi.CreateSecurityGroupRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSecurityGroupRuleRequest** | [**CreateSecurityGroupRuleRequest**](CreateSecurityGroupRuleRequest.md) |  | 

### Return type

[**CreateSecurityGroupRuleResponse**](CreateSecurityGroupRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityGroupRule

> DeleteSecurityGroupRuleResponse DeleteSecurityGroupRule(ctx).DeleteSecurityGroupRuleRequest(deleteSecurityGroupRuleRequest).Execute()



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
    deleteSecurityGroupRuleRequest := *openapiclient.NewDeleteSecurityGroupRuleRequest("Flow_example", "SecurityGroupId_example") // DeleteSecurityGroupRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupRuleApi.DeleteSecurityGroupRule(context.Background()).DeleteSecurityGroupRuleRequest(deleteSecurityGroupRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupRuleApi.DeleteSecurityGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSecurityGroupRule`: DeleteSecurityGroupRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupRuleApi.DeleteSecurityGroupRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSecurityGroupRuleRequest** | [**DeleteSecurityGroupRuleRequest**](DeleteSecurityGroupRuleRequest.md) |  | 

### Return type

[**DeleteSecurityGroupRuleResponse**](DeleteSecurityGroupRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

