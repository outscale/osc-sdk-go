# \LoadBalancerApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancer**](LoadBalancerApi.md#CreateLoadBalancer) | **Post** /CreateLoadBalancer | 
[**CreateLoadBalancerTags**](LoadBalancerApi.md#CreateLoadBalancerTags) | **Post** /CreateLoadBalancerTags | 
[**DeleteLoadBalancer**](LoadBalancerApi.md#DeleteLoadBalancer) | **Post** /DeleteLoadBalancer | 
[**DeleteLoadBalancerTags**](LoadBalancerApi.md#DeleteLoadBalancerTags) | **Post** /DeleteLoadBalancerTags | 
[**DeregisterVmsInLoadBalancer**](LoadBalancerApi.md#DeregisterVmsInLoadBalancer) | **Post** /DeregisterVmsInLoadBalancer | 
[**LinkLoadBalancerBackendMachines**](LoadBalancerApi.md#LinkLoadBalancerBackendMachines) | **Post** /LinkLoadBalancerBackendMachines | 
[**ReadLoadBalancerTags**](LoadBalancerApi.md#ReadLoadBalancerTags) | **Post** /ReadLoadBalancerTags | 
[**ReadLoadBalancers**](LoadBalancerApi.md#ReadLoadBalancers) | **Post** /ReadLoadBalancers | 
[**ReadVmsHealth**](LoadBalancerApi.md#ReadVmsHealth) | **Post** /ReadVmsHealth | 
[**RegisterVmsInLoadBalancer**](LoadBalancerApi.md#RegisterVmsInLoadBalancer) | **Post** /RegisterVmsInLoadBalancer | 
[**UnlinkLoadBalancerBackendMachines**](LoadBalancerApi.md#UnlinkLoadBalancerBackendMachines) | **Post** /UnlinkLoadBalancerBackendMachines | 
[**UpdateLoadBalancer**](LoadBalancerApi.md#UpdateLoadBalancer) | **Post** /UpdateLoadBalancer | 



## CreateLoadBalancer

> CreateLoadBalancerResponse CreateLoadBalancer(ctx).CreateLoadBalancerRequest(createLoadBalancerRequest).Execute()



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
    createLoadBalancerRequest := *openapiclient.NewCreateLoadBalancerRequest([]openapiclient.ListenerForCreation{*openapiclient.NewListenerForCreation(int32(123), int32(123), "LoadBalancerProtocol_example")}, "LoadBalancerName_example") // CreateLoadBalancerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.CreateLoadBalancer(context.Background()).CreateLoadBalancerRequest(createLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.CreateLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancer`: CreateLoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.CreateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerRequest** | [**CreateLoadBalancerRequest**](CreateLoadBalancerRequest.md) |  | 

### Return type

[**CreateLoadBalancerResponse**](CreateLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancerTags

> CreateLoadBalancerTagsResponse CreateLoadBalancerTags(ctx).CreateLoadBalancerTagsRequest(createLoadBalancerTagsRequest).Execute()



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
    createLoadBalancerTagsRequest := *openapiclient.NewCreateLoadBalancerTagsRequest([]string{"LoadBalancerNames_example"}, []openapiclient.ResourceTag{*openapiclient.NewResourceTag("Key_example", "Value_example")}) // CreateLoadBalancerTagsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.CreateLoadBalancerTags(context.Background()).CreateLoadBalancerTagsRequest(createLoadBalancerTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.CreateLoadBalancerTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancerTags`: CreateLoadBalancerTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.CreateLoadBalancerTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoadBalancerTagsRequest** | [**CreateLoadBalancerTagsRequest**](CreateLoadBalancerTagsRequest.md) |  | 

### Return type

[**CreateLoadBalancerTagsResponse**](CreateLoadBalancerTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancer

> DeleteLoadBalancerResponse DeleteLoadBalancer(ctx).DeleteLoadBalancerRequest(deleteLoadBalancerRequest).Execute()



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
    deleteLoadBalancerRequest := *openapiclient.NewDeleteLoadBalancerRequest("LoadBalancerName_example") // DeleteLoadBalancerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.DeleteLoadBalancer(context.Background()).DeleteLoadBalancerRequest(deleteLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.DeleteLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancer`: DeleteLoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.DeleteLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerRequest** | [**DeleteLoadBalancerRequest**](DeleteLoadBalancerRequest.md) |  | 

### Return type

[**DeleteLoadBalancerResponse**](DeleteLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancerTags

> DeleteLoadBalancerTagsResponse DeleteLoadBalancerTags(ctx).DeleteLoadBalancerTagsRequest(deleteLoadBalancerTagsRequest).Execute()



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
    deleteLoadBalancerTagsRequest := *openapiclient.NewDeleteLoadBalancerTagsRequest([]string{"LoadBalancerNames_example"}, []openapiclient.ResourceLoadBalancerTag{*openapiclient.NewResourceLoadBalancerTag()}) // DeleteLoadBalancerTagsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.DeleteLoadBalancerTags(context.Background()).DeleteLoadBalancerTagsRequest(deleteLoadBalancerTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.DeleteLoadBalancerTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancerTags`: DeleteLoadBalancerTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.DeleteLoadBalancerTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteLoadBalancerTagsRequest** | [**DeleteLoadBalancerTagsRequest**](DeleteLoadBalancerTagsRequest.md) |  | 

### Return type

[**DeleteLoadBalancerTagsResponse**](DeleteLoadBalancerTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeregisterVmsInLoadBalancer

> DeregisterVmsInLoadBalancerResponse DeregisterVmsInLoadBalancer(ctx).DeregisterVmsInLoadBalancerRequest(deregisterVmsInLoadBalancerRequest).Execute()



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
    deregisterVmsInLoadBalancerRequest := *openapiclient.NewDeregisterVmsInLoadBalancerRequest([]string{"BackendVmIds_example"}, "LoadBalancerName_example") // DeregisterVmsInLoadBalancerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.DeregisterVmsInLoadBalancer(context.Background()).DeregisterVmsInLoadBalancerRequest(deregisterVmsInLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.DeregisterVmsInLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeregisterVmsInLoadBalancer`: DeregisterVmsInLoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.DeregisterVmsInLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterVmsInLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deregisterVmsInLoadBalancerRequest** | [**DeregisterVmsInLoadBalancerRequest**](DeregisterVmsInLoadBalancerRequest.md) |  | 

### Return type

[**DeregisterVmsInLoadBalancerResponse**](DeregisterVmsInLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkLoadBalancerBackendMachines

> LinkLoadBalancerBackendMachinesResponse LinkLoadBalancerBackendMachines(ctx).LinkLoadBalancerBackendMachinesRequest(linkLoadBalancerBackendMachinesRequest).Execute()



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
    linkLoadBalancerBackendMachinesRequest := *openapiclient.NewLinkLoadBalancerBackendMachinesRequest("LoadBalancerName_example") // LinkLoadBalancerBackendMachinesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.LinkLoadBalancerBackendMachines(context.Background()).LinkLoadBalancerBackendMachinesRequest(linkLoadBalancerBackendMachinesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.LinkLoadBalancerBackendMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkLoadBalancerBackendMachines`: LinkLoadBalancerBackendMachinesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.LinkLoadBalancerBackendMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkLoadBalancerBackendMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkLoadBalancerBackendMachinesRequest** | [**LinkLoadBalancerBackendMachinesRequest**](LinkLoadBalancerBackendMachinesRequest.md) |  | 

### Return type

[**LinkLoadBalancerBackendMachinesResponse**](LinkLoadBalancerBackendMachinesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadLoadBalancerTags

> ReadLoadBalancerTagsResponse ReadLoadBalancerTags(ctx).ReadLoadBalancerTagsRequest(readLoadBalancerTagsRequest).Execute()



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
    readLoadBalancerTagsRequest := *openapiclient.NewReadLoadBalancerTagsRequest([]string{"LoadBalancerNames_example"}) // ReadLoadBalancerTagsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.ReadLoadBalancerTags(context.Background()).ReadLoadBalancerTagsRequest(readLoadBalancerTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.ReadLoadBalancerTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadLoadBalancerTags`: ReadLoadBalancerTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.ReadLoadBalancerTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadLoadBalancerTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLoadBalancerTagsRequest** | [**ReadLoadBalancerTagsRequest**](ReadLoadBalancerTagsRequest.md) |  | 

### Return type

[**ReadLoadBalancerTagsResponse**](ReadLoadBalancerTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadLoadBalancers

> ReadLoadBalancersResponse ReadLoadBalancers(ctx).ReadLoadBalancersRequest(readLoadBalancersRequest).Execute()



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
    readLoadBalancersRequest := *openapiclient.NewReadLoadBalancersRequest() // ReadLoadBalancersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.ReadLoadBalancers(context.Background()).ReadLoadBalancersRequest(readLoadBalancersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.ReadLoadBalancers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadLoadBalancers`: ReadLoadBalancersResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.ReadLoadBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadLoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLoadBalancersRequest** | [**ReadLoadBalancersRequest**](ReadLoadBalancersRequest.md) |  | 

### Return type

[**ReadLoadBalancersResponse**](ReadLoadBalancersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmsHealth

> ReadVmsHealthResponse ReadVmsHealth(ctx).ReadVmsHealthRequest(readVmsHealthRequest).Execute()



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
    readVmsHealthRequest := *openapiclient.NewReadVmsHealthRequest("LoadBalancerName_example") // ReadVmsHealthRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.ReadVmsHealth(context.Background()).ReadVmsHealthRequest(readVmsHealthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.ReadVmsHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVmsHealth`: ReadVmsHealthResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.ReadVmsHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVmsHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmsHealthRequest** | [**ReadVmsHealthRequest**](ReadVmsHealthRequest.md) |  | 

### Return type

[**ReadVmsHealthResponse**](ReadVmsHealthResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterVmsInLoadBalancer

> RegisterVmsInLoadBalancerResponse RegisterVmsInLoadBalancer(ctx).RegisterVmsInLoadBalancerRequest(registerVmsInLoadBalancerRequest).Execute()



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
    registerVmsInLoadBalancerRequest := *openapiclient.NewRegisterVmsInLoadBalancerRequest([]string{"BackendVmIds_example"}, "LoadBalancerName_example") // RegisterVmsInLoadBalancerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.RegisterVmsInLoadBalancer(context.Background()).RegisterVmsInLoadBalancerRequest(registerVmsInLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.RegisterVmsInLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterVmsInLoadBalancer`: RegisterVmsInLoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.RegisterVmsInLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterVmsInLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerVmsInLoadBalancerRequest** | [**RegisterVmsInLoadBalancerRequest**](RegisterVmsInLoadBalancerRequest.md) |  | 

### Return type

[**RegisterVmsInLoadBalancerResponse**](RegisterVmsInLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkLoadBalancerBackendMachines

> UnlinkLoadBalancerBackendMachinesResponse UnlinkLoadBalancerBackendMachines(ctx).UnlinkLoadBalancerBackendMachinesRequest(unlinkLoadBalancerBackendMachinesRequest).Execute()



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
    unlinkLoadBalancerBackendMachinesRequest := *openapiclient.NewUnlinkLoadBalancerBackendMachinesRequest("LoadBalancerName_example") // UnlinkLoadBalancerBackendMachinesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.UnlinkLoadBalancerBackendMachines(context.Background()).UnlinkLoadBalancerBackendMachinesRequest(unlinkLoadBalancerBackendMachinesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.UnlinkLoadBalancerBackendMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkLoadBalancerBackendMachines`: UnlinkLoadBalancerBackendMachinesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.UnlinkLoadBalancerBackendMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkLoadBalancerBackendMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkLoadBalancerBackendMachinesRequest** | [**UnlinkLoadBalancerBackendMachinesRequest**](UnlinkLoadBalancerBackendMachinesRequest.md) |  | 

### Return type

[**UnlinkLoadBalancerBackendMachinesResponse**](UnlinkLoadBalancerBackendMachinesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancer

> UpdateLoadBalancerResponse UpdateLoadBalancer(ctx).UpdateLoadBalancerRequest(updateLoadBalancerRequest).Execute()



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
    updateLoadBalancerRequest := *openapiclient.NewUpdateLoadBalancerRequest("LoadBalancerName_example") // UpdateLoadBalancerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancerApi.UpdateLoadBalancer(context.Background()).UpdateLoadBalancerRequest(updateLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.UpdateLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancer`: UpdateLoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.UpdateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLoadBalancerRequest** | [**UpdateLoadBalancerRequest**](UpdateLoadBalancerRequest.md) |  | 

### Return type

[**UpdateLoadBalancerResponse**](UpdateLoadBalancerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

