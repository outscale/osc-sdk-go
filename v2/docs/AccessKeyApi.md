# \AccessKeyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessKey**](AccessKeyApi.md#CreateAccessKey) | **Post** /CreateAccessKey | 
[**DeleteAccessKey**](AccessKeyApi.md#DeleteAccessKey) | **Post** /DeleteAccessKey | 
[**ReadAccessKeys**](AccessKeyApi.md#ReadAccessKeys) | **Post** /ReadAccessKeys | 
[**ReadSecretAccessKey**](AccessKeyApi.md#ReadSecretAccessKey) | **Post** /ReadSecretAccessKey | 
[**UpdateAccessKey**](AccessKeyApi.md#UpdateAccessKey) | **Post** /UpdateAccessKey | 



## CreateAccessKey

> CreateAccessKeyResponse CreateAccessKey(ctx).CreateAccessKeyRequest(createAccessKeyRequest).Execute()



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
    createAccessKeyRequest := *openapiclient.NewCreateAccessKeyRequest() // CreateAccessKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyApi.CreateAccessKey(context.Background()).CreateAccessKeyRequest(createAccessKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyApi.CreateAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessKey`: CreateAccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyApi.CreateAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessKeyRequest** | [**CreateAccessKeyRequest**](CreateAccessKeyRequest.md) |  | 

### Return type

[**CreateAccessKeyResponse**](CreateAccessKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessKey

> DeleteAccessKeyResponse DeleteAccessKey(ctx).DeleteAccessKeyRequest(deleteAccessKeyRequest).Execute()



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
    deleteAccessKeyRequest := *openapiclient.NewDeleteAccessKeyRequest("AccessKeyId_example") // DeleteAccessKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyApi.DeleteAccessKey(context.Background()).DeleteAccessKeyRequest(deleteAccessKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyApi.DeleteAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessKey`: DeleteAccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyApi.DeleteAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAccessKeyRequest** | [**DeleteAccessKeyRequest**](DeleteAccessKeyRequest.md) |  | 

### Return type

[**DeleteAccessKeyResponse**](DeleteAccessKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccessKeys

> ReadAccessKeysResponse ReadAccessKeys(ctx).ReadAccessKeysRequest(readAccessKeysRequest).Execute()



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
    readAccessKeysRequest := *openapiclient.NewReadAccessKeysRequest() // ReadAccessKeysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyApi.ReadAccessKeys(context.Background()).ReadAccessKeysRequest(readAccessKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyApi.ReadAccessKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccessKeys`: ReadAccessKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyApi.ReadAccessKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadAccessKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readAccessKeysRequest** | [**ReadAccessKeysRequest**](ReadAccessKeysRequest.md) |  | 

### Return type

[**ReadAccessKeysResponse**](ReadAccessKeysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSecretAccessKey

> ReadSecretAccessKeyResponse ReadSecretAccessKey(ctx).ReadSecretAccessKeyRequest(readSecretAccessKeyRequest).Execute()



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
    readSecretAccessKeyRequest := *openapiclient.NewReadSecretAccessKeyRequest("AccessKeyId_example") // ReadSecretAccessKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyApi.ReadSecretAccessKey(context.Background()).ReadSecretAccessKeyRequest(readSecretAccessKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyApi.ReadSecretAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSecretAccessKey`: ReadSecretAccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyApi.ReadSecretAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadSecretAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSecretAccessKeyRequest** | [**ReadSecretAccessKeyRequest**](ReadSecretAccessKeyRequest.md) |  | 

### Return type

[**ReadSecretAccessKeyResponse**](ReadSecretAccessKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessKey

> UpdateAccessKeyResponse UpdateAccessKey(ctx).UpdateAccessKeyRequest(updateAccessKeyRequest).Execute()



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
    updateAccessKeyRequest := *openapiclient.NewUpdateAccessKeyRequest("AccessKeyId_example", "State_example") // UpdateAccessKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyApi.UpdateAccessKey(context.Background()).UpdateAccessKeyRequest(updateAccessKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyApi.UpdateAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessKey`: UpdateAccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyApi.UpdateAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccessKeyRequest** | [**UpdateAccessKeyRequest**](UpdateAccessKeyRequest.md) |  | 

### Return type

[**UpdateAccessKeyResponse**](UpdateAccessKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

