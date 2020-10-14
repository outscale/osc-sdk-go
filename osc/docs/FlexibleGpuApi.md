# \FlexibleGpuApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlexibleGpu**](FlexibleGpuApi.md#CreateFlexibleGpu) | **Post** /CreateFlexibleGpu | 
[**DeleteFlexibleGpu**](FlexibleGpuApi.md#DeleteFlexibleGpu) | **Post** /DeleteFlexibleGpu | 
[**LinkFlexibleGpu**](FlexibleGpuApi.md#LinkFlexibleGpu) | **Post** /LinkFlexibleGpu | 
[**ReadFlexibleGpuCatalog**](FlexibleGpuApi.md#ReadFlexibleGpuCatalog) | **Post** /ReadFlexibleGpuCatalog | 
[**ReadFlexibleGpus**](FlexibleGpuApi.md#ReadFlexibleGpus) | **Post** /ReadFlexibleGpus | 
[**UnlinkFlexibleGpu**](FlexibleGpuApi.md#UnlinkFlexibleGpu) | **Post** /UnlinkFlexibleGpu | 
[**UpdateFlexibleGpu**](FlexibleGpuApi.md#UpdateFlexibleGpu) | **Post** /UpdateFlexibleGpu | 



## CreateFlexibleGpu

> CreateFlexibleGpuResponse CreateFlexibleGpu(ctx).CreateFlexibleGpuRequest(createFlexibleGpuRequest).Execute()



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
    createFlexibleGpuRequest := openapiclient.CreateFlexibleGpuRequest{DeleteOnVmDeletion: false, DryRun: false, Generation: "Generation_example", ModelName: "ModelName_example", SubregionName: "SubregionName_example"} // CreateFlexibleGpuRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlexibleGpuApi.CreateFlexibleGpu(context.Background()).CreateFlexibleGpuRequest(createFlexibleGpuRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexibleGpuApi.CreateFlexibleGpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlexibleGpu`: CreateFlexibleGpuResponse
    fmt.Fprintf(os.Stdout, "Response from `FlexibleGpuApi.CreateFlexibleGpu`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlexibleGpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFlexibleGpuRequest** | [**CreateFlexibleGpuRequest**](CreateFlexibleGpuRequest.md) |  | 

### Return type

[**CreateFlexibleGpuResponse**](CreateFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlexibleGpu

> DeleteFlexibleGpuResponse DeleteFlexibleGpu(ctx).DeleteFlexibleGpuRequest(deleteFlexibleGpuRequest).Execute()



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
    deleteFlexibleGpuRequest := openapiclient.DeleteFlexibleGpuRequest{DryRun: false, FlexibleGpuId: "FlexibleGpuId_example"} // DeleteFlexibleGpuRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlexibleGpuApi.DeleteFlexibleGpu(context.Background()).DeleteFlexibleGpuRequest(deleteFlexibleGpuRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexibleGpuApi.DeleteFlexibleGpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFlexibleGpu`: DeleteFlexibleGpuResponse
    fmt.Fprintf(os.Stdout, "Response from `FlexibleGpuApi.DeleteFlexibleGpu`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlexibleGpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteFlexibleGpuRequest** | [**DeleteFlexibleGpuRequest**](DeleteFlexibleGpuRequest.md) |  | 

### Return type

[**DeleteFlexibleGpuResponse**](DeleteFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkFlexibleGpu

> LinkFlexibleGpuResponse LinkFlexibleGpu(ctx).LinkFlexibleGpuRequest(linkFlexibleGpuRequest).Execute()



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
    linkFlexibleGpuRequest := openapiclient.LinkFlexibleGpuRequest{DryRun: false, FlexibleGpuId: "FlexibleGpuId_example", VmId: "VmId_example"} // LinkFlexibleGpuRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlexibleGpuApi.LinkFlexibleGpu(context.Background()).LinkFlexibleGpuRequest(linkFlexibleGpuRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexibleGpuApi.LinkFlexibleGpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkFlexibleGpu`: LinkFlexibleGpuResponse
    fmt.Fprintf(os.Stdout, "Response from `FlexibleGpuApi.LinkFlexibleGpu`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkFlexibleGpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkFlexibleGpuRequest** | [**LinkFlexibleGpuRequest**](LinkFlexibleGpuRequest.md) |  | 

### Return type

[**LinkFlexibleGpuResponse**](LinkFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadFlexibleGpuCatalog

> ReadFlexibleGpuCatalogResponse ReadFlexibleGpuCatalog(ctx).ReadFlexibleGpuCatalogRequest(readFlexibleGpuCatalogRequest).Execute()



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
    readFlexibleGpuCatalogRequest := openapiclient.ReadFlexibleGpuCatalogRequest{DryRun: false} // ReadFlexibleGpuCatalogRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlexibleGpuApi.ReadFlexibleGpuCatalog(context.Background()).ReadFlexibleGpuCatalogRequest(readFlexibleGpuCatalogRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexibleGpuApi.ReadFlexibleGpuCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadFlexibleGpuCatalog`: ReadFlexibleGpuCatalogResponse
    fmt.Fprintf(os.Stdout, "Response from `FlexibleGpuApi.ReadFlexibleGpuCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadFlexibleGpuCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readFlexibleGpuCatalogRequest** | [**ReadFlexibleGpuCatalogRequest**](ReadFlexibleGpuCatalogRequest.md) |  | 

### Return type

[**ReadFlexibleGpuCatalogResponse**](ReadFlexibleGpuCatalogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadFlexibleGpus

> ReadFlexibleGpusResponse ReadFlexibleGpus(ctx).ReadFlexibleGpusRequest(readFlexibleGpusRequest).Execute()



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
    readFlexibleGpusRequest := openapiclient.ReadFlexibleGpusRequest{DryRun: false, Filters: openapiclient.FiltersFlexibleGpu{DeleteOnVmDeletion: false, FlexibleGpuIds: []string{"FlexibleGpuIds_example"), Generations: []string{"Generations_example"), ModelNames: []string{"ModelNames_example"), States: []string{"States_example"), SubregionNames: []string{"SubregionNames_example"), VmIds: []string{"VmIds_example")}} // ReadFlexibleGpusRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlexibleGpuApi.ReadFlexibleGpus(context.Background()).ReadFlexibleGpusRequest(readFlexibleGpusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexibleGpuApi.ReadFlexibleGpus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadFlexibleGpus`: ReadFlexibleGpusResponse
    fmt.Fprintf(os.Stdout, "Response from `FlexibleGpuApi.ReadFlexibleGpus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadFlexibleGpusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readFlexibleGpusRequest** | [**ReadFlexibleGpusRequest**](ReadFlexibleGpusRequest.md) |  | 

### Return type

[**ReadFlexibleGpusResponse**](ReadFlexibleGpusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkFlexibleGpu

> UnlinkFlexibleGpuResponse UnlinkFlexibleGpu(ctx).UnlinkFlexibleGpuRequest(unlinkFlexibleGpuRequest).Execute()



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
    unlinkFlexibleGpuRequest := openapiclient.UnlinkFlexibleGpuRequest{DryRun: false, FlexibleGpuId: "FlexibleGpuId_example"} // UnlinkFlexibleGpuRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlexibleGpuApi.UnlinkFlexibleGpu(context.Background()).UnlinkFlexibleGpuRequest(unlinkFlexibleGpuRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexibleGpuApi.UnlinkFlexibleGpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkFlexibleGpu`: UnlinkFlexibleGpuResponse
    fmt.Fprintf(os.Stdout, "Response from `FlexibleGpuApi.UnlinkFlexibleGpu`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkFlexibleGpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkFlexibleGpuRequest** | [**UnlinkFlexibleGpuRequest**](UnlinkFlexibleGpuRequest.md) |  | 

### Return type

[**UnlinkFlexibleGpuResponse**](UnlinkFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlexibleGpu

> UpdateFlexibleGpuResponse UpdateFlexibleGpu(ctx).UpdateFlexibleGpuRequest(updateFlexibleGpuRequest).Execute()



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
    updateFlexibleGpuRequest := openapiclient.UpdateFlexibleGpuRequest{DeleteOnVmDeletion: false, DryRun: false, FlexibleGpuId: "FlexibleGpuId_example"} // UpdateFlexibleGpuRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlexibleGpuApi.UpdateFlexibleGpu(context.Background()).UpdateFlexibleGpuRequest(updateFlexibleGpuRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexibleGpuApi.UpdateFlexibleGpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlexibleGpu`: UpdateFlexibleGpuResponse
    fmt.Fprintf(os.Stdout, "Response from `FlexibleGpuApi.UpdateFlexibleGpu`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlexibleGpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFlexibleGpuRequest** | [**UpdateFlexibleGpuRequest**](UpdateFlexibleGpuRequest.md) |  | 

### Return type

[**UpdateFlexibleGpuResponse**](UpdateFlexibleGpuResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

