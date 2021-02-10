# \NicApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNic**](NicApi.md#CreateNic) | **Post** /CreateNic | 
[**DeleteNic**](NicApi.md#DeleteNic) | **Post** /DeleteNic | 
[**LinkNic**](NicApi.md#LinkNic) | **Post** /LinkNic | 
[**LinkPrivateIps**](NicApi.md#LinkPrivateIps) | **Post** /LinkPrivateIps | 
[**ReadNics**](NicApi.md#ReadNics) | **Post** /ReadNics | 
[**UnlinkNic**](NicApi.md#UnlinkNic) | **Post** /UnlinkNic | 
[**UnlinkPrivateIps**](NicApi.md#UnlinkPrivateIps) | **Post** /UnlinkPrivateIps | 
[**UpdateNic**](NicApi.md#UpdateNic) | **Post** /UpdateNic | 



## CreateNic

> CreateNicResponse CreateNic(ctx).CreateNicRequest(createNicRequest).Execute()



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
    createNicRequest := *openapiclient.NewCreateNicRequest("SubnetId_example") // CreateNicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.CreateNic(context.Background()).CreateNicRequest(createNicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.CreateNic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNic`: CreateNicResponse
    fmt.Fprintf(os.Stdout, "Response from `NicApi.CreateNic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNicRequest** | [**CreateNicRequest**](CreateNicRequest.md) |  | 

### Return type

[**CreateNicResponse**](CreateNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNic

> DeleteNicResponse DeleteNic(ctx).DeleteNicRequest(deleteNicRequest).Execute()



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
    deleteNicRequest := *openapiclient.NewDeleteNicRequest("NicId_example") // DeleteNicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.DeleteNic(context.Background()).DeleteNicRequest(deleteNicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DeleteNic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNic`: DeleteNicResponse
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DeleteNic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNicRequest** | [**DeleteNicRequest**](DeleteNicRequest.md) |  | 

### Return type

[**DeleteNicResponse**](DeleteNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkNic

> LinkNicResponse LinkNic(ctx).LinkNicRequest(linkNicRequest).Execute()



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
    linkNicRequest := *openapiclient.NewLinkNicRequest(int32(123), "NicId_example", "VmId_example") // LinkNicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.LinkNic(context.Background()).LinkNicRequest(linkNicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.LinkNic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkNic`: LinkNicResponse
    fmt.Fprintf(os.Stdout, "Response from `NicApi.LinkNic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkNicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkNicRequest** | [**LinkNicRequest**](LinkNicRequest.md) |  | 

### Return type

[**LinkNicResponse**](LinkNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkPrivateIps

> LinkPrivateIpsResponse LinkPrivateIps(ctx).LinkPrivateIpsRequest(linkPrivateIpsRequest).Execute()



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
    linkPrivateIpsRequest := *openapiclient.NewLinkPrivateIpsRequest("NicId_example") // LinkPrivateIpsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.LinkPrivateIps(context.Background()).LinkPrivateIpsRequest(linkPrivateIpsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.LinkPrivateIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkPrivateIps`: LinkPrivateIpsResponse
    fmt.Fprintf(os.Stdout, "Response from `NicApi.LinkPrivateIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkPrivateIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkPrivateIpsRequest** | [**LinkPrivateIpsRequest**](LinkPrivateIpsRequest.md) |  | 

### Return type

[**LinkPrivateIpsResponse**](LinkPrivateIpsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNics

> ReadNicsResponse ReadNics(ctx).ReadNicsRequest(readNicsRequest).Execute()



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
    readNicsRequest := *openapiclient.NewReadNicsRequest() // ReadNicsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.ReadNics(context.Background()).ReadNicsRequest(readNicsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.ReadNics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNics`: ReadNicsResponse
    fmt.Fprintf(os.Stdout, "Response from `NicApi.ReadNics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadNicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNicsRequest** | [**ReadNicsRequest**](ReadNicsRequest.md) |  | 

### Return type

[**ReadNicsResponse**](ReadNicsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkNic

> UnlinkNicResponse UnlinkNic(ctx).UnlinkNicRequest(unlinkNicRequest).Execute()



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
    unlinkNicRequest := *openapiclient.NewUnlinkNicRequest("LinkNicId_example") // UnlinkNicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.UnlinkNic(context.Background()).UnlinkNicRequest(unlinkNicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.UnlinkNic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkNic`: UnlinkNicResponse
    fmt.Fprintf(os.Stdout, "Response from `NicApi.UnlinkNic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkNicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkNicRequest** | [**UnlinkNicRequest**](UnlinkNicRequest.md) |  | 

### Return type

[**UnlinkNicResponse**](UnlinkNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkPrivateIps

> UnlinkPrivateIpsResponse UnlinkPrivateIps(ctx).UnlinkPrivateIpsRequest(unlinkPrivateIpsRequest).Execute()



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
    unlinkPrivateIpsRequest := *openapiclient.NewUnlinkPrivateIpsRequest("NicId_example", []string{"PrivateIps_example"}) // UnlinkPrivateIpsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.UnlinkPrivateIps(context.Background()).UnlinkPrivateIpsRequest(unlinkPrivateIpsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.UnlinkPrivateIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkPrivateIps`: UnlinkPrivateIpsResponse
    fmt.Fprintf(os.Stdout, "Response from `NicApi.UnlinkPrivateIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkPrivateIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkPrivateIpsRequest** | [**UnlinkPrivateIpsRequest**](UnlinkPrivateIpsRequest.md) |  | 

### Return type

[**UnlinkPrivateIpsResponse**](UnlinkPrivateIpsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNic

> UpdateNicResponse UpdateNic(ctx).UpdateNicRequest(updateNicRequest).Execute()



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
    updateNicRequest := *openapiclient.NewUpdateNicRequest("NicId_example") // UpdateNicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.UpdateNic(context.Background()).UpdateNicRequest(updateNicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.UpdateNic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNic`: UpdateNicResponse
    fmt.Fprintf(os.Stdout, "Response from `NicApi.UpdateNic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNicRequest** | [**UpdateNicRequest**](UpdateNicRequest.md) |  | 

### Return type

[**UpdateNicResponse**](UpdateNicResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

