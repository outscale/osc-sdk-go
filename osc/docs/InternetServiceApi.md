# \InternetServiceApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternetService**](InternetServiceApi.md#CreateInternetService) | **Post** /CreateInternetService | 
[**DeleteInternetService**](InternetServiceApi.md#DeleteInternetService) | **Post** /DeleteInternetService | 
[**LinkInternetService**](InternetServiceApi.md#LinkInternetService) | **Post** /LinkInternetService | 
[**ReadInternetServices**](InternetServiceApi.md#ReadInternetServices) | **Post** /ReadInternetServices | 
[**UnlinkInternetService**](InternetServiceApi.md#UnlinkInternetService) | **Post** /UnlinkInternetService | 



## CreateInternetService

> CreateInternetServiceResponse CreateInternetService(ctx).CreateInternetServiceRequest(createInternetServiceRequest).Execute()



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
    createInternetServiceRequest := openapiclient.CreateInternetServiceRequest{DryRun: false} // CreateInternetServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetServiceApi.CreateInternetService(context.Background()).CreateInternetServiceRequest(createInternetServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetServiceApi.CreateInternetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInternetService`: CreateInternetServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InternetServiceApi.CreateInternetService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInternetServiceRequest** | [**CreateInternetServiceRequest**](CreateInternetServiceRequest.md) |  | 

### Return type

[**CreateInternetServiceResponse**](CreateInternetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInternetService

> DeleteInternetServiceResponse DeleteInternetService(ctx).DeleteInternetServiceRequest(deleteInternetServiceRequest).Execute()



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
    deleteInternetServiceRequest := openapiclient.DeleteInternetServiceRequest{DryRun: false, InternetServiceId: "InternetServiceId_example"} // DeleteInternetServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetServiceApi.DeleteInternetService(context.Background()).DeleteInternetServiceRequest(deleteInternetServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetServiceApi.DeleteInternetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInternetService`: DeleteInternetServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InternetServiceApi.DeleteInternetService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInternetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteInternetServiceRequest** | [**DeleteInternetServiceRequest**](DeleteInternetServiceRequest.md) |  | 

### Return type

[**DeleteInternetServiceResponse**](DeleteInternetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkInternetService

> LinkInternetServiceResponse LinkInternetService(ctx).LinkInternetServiceRequest(linkInternetServiceRequest).Execute()



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
    linkInternetServiceRequest := openapiclient.LinkInternetServiceRequest{DryRun: false, InternetServiceId: "InternetServiceId_example", NetId: "NetId_example"} // LinkInternetServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetServiceApi.LinkInternetService(context.Background()).LinkInternetServiceRequest(linkInternetServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetServiceApi.LinkInternetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkInternetService`: LinkInternetServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InternetServiceApi.LinkInternetService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkInternetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkInternetServiceRequest** | [**LinkInternetServiceRequest**](LinkInternetServiceRequest.md) |  | 

### Return type

[**LinkInternetServiceResponse**](LinkInternetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInternetServices

> ReadInternetServicesResponse ReadInternetServices(ctx).ReadInternetServicesRequest(readInternetServicesRequest).Execute()



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
    readInternetServicesRequest := openapiclient.ReadInternetServicesRequest{DryRun: false, Filters: openapiclient.FiltersInternetService{InternetServiceIds: []string{"InternetServiceIds_example"), LinkNetIds: []string{"LinkNetIds_example"), LinkStates: []string{"LinkStates_example"), TagKeys: []string{"TagKeys_example"), TagValues: []string{"TagValues_example"), Tags: []string{"Tags_example")}} // ReadInternetServicesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetServiceApi.ReadInternetServices(context.Background()).ReadInternetServicesRequest(readInternetServicesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetServiceApi.ReadInternetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadInternetServices`: ReadInternetServicesResponse
    fmt.Fprintf(os.Stdout, "Response from `InternetServiceApi.ReadInternetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadInternetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readInternetServicesRequest** | [**ReadInternetServicesRequest**](ReadInternetServicesRequest.md) |  | 

### Return type

[**ReadInternetServicesResponse**](ReadInternetServicesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkInternetService

> UnlinkInternetServiceResponse UnlinkInternetService(ctx).UnlinkInternetServiceRequest(unlinkInternetServiceRequest).Execute()



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
    unlinkInternetServiceRequest := openapiclient.UnlinkInternetServiceRequest{DryRun: false, InternetServiceId: "InternetServiceId_example", NetId: "NetId_example"} // UnlinkInternetServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetServiceApi.UnlinkInternetService(context.Background()).UnlinkInternetServiceRequest(unlinkInternetServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetServiceApi.UnlinkInternetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkInternetService`: UnlinkInternetServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InternetServiceApi.UnlinkInternetService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkInternetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkInternetServiceRequest** | [**UnlinkInternetServiceRequest**](UnlinkInternetServiceRequest.md) |  | 

### Return type

[**UnlinkInternetServiceResponse**](UnlinkInternetServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

