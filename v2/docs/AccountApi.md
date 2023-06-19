# \AccountApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAuthentication**](AccountApi.md#CheckAuthentication) | **Post** /CheckAuthentication | 
[**CreateAccount**](AccountApi.md#CreateAccount) | **Post** /CreateAccount | 
[**ReadAccounts**](AccountApi.md#ReadAccounts) | **Post** /ReadAccounts | 
[**ReadConsumptionAccount**](AccountApi.md#ReadConsumptionAccount) | **Post** /ReadConsumptionAccount | 
[**ResetAccountPassword**](AccountApi.md#ResetAccountPassword) | **Post** /ResetAccountPassword | 
[**SendResetPasswordEmail**](AccountApi.md#SendResetPasswordEmail) | **Post** /SendResetPasswordEmail | 
[**UpdateAccount**](AccountApi.md#UpdateAccount) | **Post** /UpdateAccount | 



## CheckAuthentication

> CheckAuthenticationResponse CheckAuthentication(ctx).CheckAuthenticationRequest(checkAuthenticationRequest).Execute()



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
    checkAuthenticationRequest := *openapiclient.NewCheckAuthenticationRequest("Login_example", "Password_example") // CheckAuthenticationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CheckAuthentication(context.Background()).CheckAuthenticationRequest(checkAuthenticationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CheckAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAuthentication`: CheckAuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CheckAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAuthenticationRequest** | [**CheckAuthenticationRequest**](CheckAuthenticationRequest.md) |  | 

### Return type

[**CheckAuthenticationResponse**](CheckAuthenticationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> CreateAccountResponse CreateAccount(ctx).CreateAccountRequest(createAccountRequest).Execute()



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
    createAccountRequest := *openapiclient.NewCreateAccountRequest("City_example", "CompanyName_example", "Country_example", "CustomerId_example", "Email_example", "FirstName_example", "LastName_example", "ZipCode_example") // CreateAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CreateAccount(context.Background()).CreateAccountRequest(createAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: CreateAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountRequest** | [**CreateAccountRequest**](CreateAccountRequest.md) |  | 

### Return type

[**CreateAccountResponse**](CreateAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccounts

> ReadAccountsResponse ReadAccounts(ctx).ReadAccountsRequest(readAccountsRequest).Execute()



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
    readAccountsRequest := *openapiclient.NewReadAccountsRequest() // ReadAccountsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.ReadAccounts(context.Background()).ReadAccountsRequest(readAccountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ReadAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccounts`: ReadAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ReadAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readAccountsRequest** | [**ReadAccountsRequest**](ReadAccountsRequest.md) |  | 

### Return type

[**ReadAccountsResponse**](ReadAccountsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadConsumptionAccount

> ReadConsumptionAccountResponse ReadConsumptionAccount(ctx).ReadConsumptionAccountRequest(readConsumptionAccountRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    readConsumptionAccountRequest := *openapiclient.NewReadConsumptionAccountRequest(time.Now(), time.Now()) // ReadConsumptionAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.ReadConsumptionAccount(context.Background()).ReadConsumptionAccountRequest(readConsumptionAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ReadConsumptionAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadConsumptionAccount`: ReadConsumptionAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ReadConsumptionAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadConsumptionAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readConsumptionAccountRequest** | [**ReadConsumptionAccountRequest**](ReadConsumptionAccountRequest.md) |  | 

### Return type

[**ReadConsumptionAccountResponse**](ReadConsumptionAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetAccountPassword

> ResetAccountPasswordResponse ResetAccountPassword(ctx).ResetAccountPasswordRequest(resetAccountPasswordRequest).Execute()



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
    resetAccountPasswordRequest := *openapiclient.NewResetAccountPasswordRequest("Password_example", "Token_example") // ResetAccountPasswordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.ResetAccountPassword(context.Background()).ResetAccountPasswordRequest(resetAccountPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ResetAccountPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetAccountPassword`: ResetAccountPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ResetAccountPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetAccountPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetAccountPasswordRequest** | [**ResetAccountPasswordRequest**](ResetAccountPasswordRequest.md) |  | 

### Return type

[**ResetAccountPasswordResponse**](ResetAccountPasswordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendResetPasswordEmail

> SendResetPasswordEmailResponse SendResetPasswordEmail(ctx).SendResetPasswordEmailRequest(sendResetPasswordEmailRequest).Execute()



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
    sendResetPasswordEmailRequest := *openapiclient.NewSendResetPasswordEmailRequest("Email_example") // SendResetPasswordEmailRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.SendResetPasswordEmail(context.Background()).SendResetPasswordEmailRequest(sendResetPasswordEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.SendResetPasswordEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendResetPasswordEmail`: SendResetPasswordEmailResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.SendResetPasswordEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendResetPasswordEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendResetPasswordEmailRequest** | [**SendResetPasswordEmailRequest**](SendResetPasswordEmailRequest.md) |  | 

### Return type

[**SendResetPasswordEmailResponse**](SendResetPasswordEmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> UpdateAccountResponse UpdateAccount(ctx).UpdateAccountRequest(updateAccountRequest).Execute()



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
    updateAccountRequest := *openapiclient.NewUpdateAccountRequest() // UpdateAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.UpdateAccount(context.Background()).UpdateAccountRequest(updateAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: UpdateAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccountRequest** | [**UpdateAccountRequest**](UpdateAccountRequest.md) |  | 

### Return type

[**UpdateAccountResponse**](UpdateAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

