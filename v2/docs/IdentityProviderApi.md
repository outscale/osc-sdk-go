# \IdentityProviderApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableOutscaleLogin**](IdentityProviderApi.md#DisableOutscaleLogin) | **Post** /DisableOutscaleLogin | 
[**DisableOutscaleLoginForUsers**](IdentityProviderApi.md#DisableOutscaleLoginForUsers) | **Post** /DisableOutscaleLoginForUsers | 
[**DisableOutscaleLoginPerUsers**](IdentityProviderApi.md#DisableOutscaleLoginPerUsers) | **Post** /DisableOutscaleLoginPerUsers | 
[**EnableOutscaleLogin**](IdentityProviderApi.md#EnableOutscaleLogin) | **Post** /EnableOutscaleLogin | 
[**EnableOutscaleLoginForUsers**](IdentityProviderApi.md#EnableOutscaleLoginForUsers) | **Post** /EnableOutscaleLoginForUsers | 
[**EnableOutscaleLoginPerUsers**](IdentityProviderApi.md#EnableOutscaleLoginPerUsers) | **Post** /EnableOutscaleLoginPerUsers | 



## DisableOutscaleLogin

> DisableOutscaleLoginResponse DisableOutscaleLogin(ctx).DisableOutscaleLoginRequest(disableOutscaleLoginRequest).Execute()





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
    disableOutscaleLoginRequest := *openapiclient.NewDisableOutscaleLoginRequest() // DisableOutscaleLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProviderApi.DisableOutscaleLogin(context.Background()).DisableOutscaleLoginRequest(disableOutscaleLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.DisableOutscaleLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableOutscaleLogin`: DisableOutscaleLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.DisableOutscaleLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableOutscaleLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableOutscaleLoginRequest** | [**DisableOutscaleLoginRequest**](DisableOutscaleLoginRequest.md) |  | 

### Return type

[**DisableOutscaleLoginResponse**](DisableOutscaleLoginResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableOutscaleLoginForUsers

> DisableOutscaleLoginResponse DisableOutscaleLoginForUsers(ctx).DisableOutscaleLoginRequest(disableOutscaleLoginRequest).Execute()





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
    disableOutscaleLoginRequest := *openapiclient.NewDisableOutscaleLoginRequest() // DisableOutscaleLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProviderApi.DisableOutscaleLoginForUsers(context.Background()).DisableOutscaleLoginRequest(disableOutscaleLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.DisableOutscaleLoginForUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableOutscaleLoginForUsers`: DisableOutscaleLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.DisableOutscaleLoginForUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableOutscaleLoginForUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableOutscaleLoginRequest** | [**DisableOutscaleLoginRequest**](DisableOutscaleLoginRequest.md) |  | 

### Return type

[**DisableOutscaleLoginResponse**](DisableOutscaleLoginResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableOutscaleLoginPerUsers

> DisableOutscaleLoginPerUsersResponse DisableOutscaleLoginPerUsers(ctx).DisableOutscaleLoginPerUsersRequest(disableOutscaleLoginPerUsersRequest).Execute()





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
    disableOutscaleLoginPerUsersRequest := *openapiclient.NewDisableOutscaleLoginPerUsersRequest() // DisableOutscaleLoginPerUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProviderApi.DisableOutscaleLoginPerUsers(context.Background()).DisableOutscaleLoginPerUsersRequest(disableOutscaleLoginPerUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.DisableOutscaleLoginPerUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableOutscaleLoginPerUsers`: DisableOutscaleLoginPerUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.DisableOutscaleLoginPerUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableOutscaleLoginPerUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableOutscaleLoginPerUsersRequest** | [**DisableOutscaleLoginPerUsersRequest**](DisableOutscaleLoginPerUsersRequest.md) |  | 

### Return type

[**DisableOutscaleLoginPerUsersResponse**](DisableOutscaleLoginPerUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableOutscaleLogin

> EnableOutscaleLoginResponse EnableOutscaleLogin(ctx).EnableOutscaleLoginRequest(enableOutscaleLoginRequest).Execute()





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
    enableOutscaleLoginRequest := *openapiclient.NewEnableOutscaleLoginRequest() // EnableOutscaleLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProviderApi.EnableOutscaleLogin(context.Background()).EnableOutscaleLoginRequest(enableOutscaleLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.EnableOutscaleLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableOutscaleLogin`: EnableOutscaleLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.EnableOutscaleLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableOutscaleLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableOutscaleLoginRequest** | [**EnableOutscaleLoginRequest**](EnableOutscaleLoginRequest.md) |  | 

### Return type

[**EnableOutscaleLoginResponse**](EnableOutscaleLoginResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableOutscaleLoginForUsers

> EnableOutscaleLoginForUsersResponse EnableOutscaleLoginForUsers(ctx).EnableOutscaleLoginForUsersRequest(enableOutscaleLoginForUsersRequest).Execute()





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
    enableOutscaleLoginForUsersRequest := *openapiclient.NewEnableOutscaleLoginForUsersRequest() // EnableOutscaleLoginForUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProviderApi.EnableOutscaleLoginForUsers(context.Background()).EnableOutscaleLoginForUsersRequest(enableOutscaleLoginForUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.EnableOutscaleLoginForUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableOutscaleLoginForUsers`: EnableOutscaleLoginForUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.EnableOutscaleLoginForUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableOutscaleLoginForUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableOutscaleLoginForUsersRequest** | [**EnableOutscaleLoginForUsersRequest**](EnableOutscaleLoginForUsersRequest.md) |  | 

### Return type

[**EnableOutscaleLoginForUsersResponse**](EnableOutscaleLoginForUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableOutscaleLoginPerUsers

> EnableOutscaleLoginPerUsersResponse EnableOutscaleLoginPerUsers(ctx).EnableOutscaleLoginPerUsersRequest(enableOutscaleLoginPerUsersRequest).Execute()





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
    enableOutscaleLoginPerUsersRequest := *openapiclient.NewEnableOutscaleLoginPerUsersRequest() // EnableOutscaleLoginPerUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProviderApi.EnableOutscaleLoginPerUsers(context.Background()).EnableOutscaleLoginPerUsersRequest(enableOutscaleLoginPerUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.EnableOutscaleLoginPerUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableOutscaleLoginPerUsers`: EnableOutscaleLoginPerUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.EnableOutscaleLoginPerUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableOutscaleLoginPerUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableOutscaleLoginPerUsersRequest** | [**EnableOutscaleLoginPerUsersRequest**](EnableOutscaleLoginPerUsersRequest.md) |  | 

### Return type

[**EnableOutscaleLoginPerUsersResponse**](EnableOutscaleLoginPerUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

