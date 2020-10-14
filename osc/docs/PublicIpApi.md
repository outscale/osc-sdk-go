# \PublicIpApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePublicIp**](PublicIpApi.md#CreatePublicIp) | **Post** /CreatePublicIp | 
[**DeletePublicIp**](PublicIpApi.md#DeletePublicIp) | **Post** /DeletePublicIp | 
[**LinkPublicIp**](PublicIpApi.md#LinkPublicIp) | **Post** /LinkPublicIp | 
[**ReadPublicIpRanges**](PublicIpApi.md#ReadPublicIpRanges) | **Post** /ReadPublicIpRanges | 
[**ReadPublicIps**](PublicIpApi.md#ReadPublicIps) | **Post** /ReadPublicIps | 
[**UnlinkPublicIp**](PublicIpApi.md#UnlinkPublicIp) | **Post** /UnlinkPublicIp | 



## CreatePublicIp

> CreatePublicIpResponse CreatePublicIp(ctx).CreatePublicIpRequest(createPublicIpRequest).Execute()



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
    createPublicIpRequest := openapiclient.CreatePublicIpRequest{DryRun: false} // CreatePublicIpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicIpApi.CreatePublicIp(context.Background()).CreatePublicIpRequest(createPublicIpRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIpApi.CreatePublicIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePublicIp`: CreatePublicIpResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicIpApi.CreatePublicIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPublicIpRequest** | [**CreatePublicIpRequest**](CreatePublicIpRequest.md) |  | 

### Return type

[**CreatePublicIpResponse**](CreatePublicIpResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicIp

> DeletePublicIpResponse DeletePublicIp(ctx).DeletePublicIpRequest(deletePublicIpRequest).Execute()



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
    deletePublicIpRequest := openapiclient.DeletePublicIpRequest{DryRun: false, PublicIp: "PublicIp_example", PublicIpId: "PublicIpId_example"} // DeletePublicIpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicIpApi.DeletePublicIp(context.Background()).DeletePublicIpRequest(deletePublicIpRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIpApi.DeletePublicIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePublicIp`: DeletePublicIpResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicIpApi.DeletePublicIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePublicIpRequest** | [**DeletePublicIpRequest**](DeletePublicIpRequest.md) |  | 

### Return type

[**DeletePublicIpResponse**](DeletePublicIpResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkPublicIp

> LinkPublicIpResponse LinkPublicIp(ctx).LinkPublicIpRequest(linkPublicIpRequest).Execute()



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
    linkPublicIpRequest := openapiclient.LinkPublicIpRequest{AllowRelink: false, DryRun: false, NicId: "NicId_example", PrivateIp: "PrivateIp_example", PublicIp: "PublicIp_example", PublicIpId: "PublicIpId_example", VmId: "VmId_example"} // LinkPublicIpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicIpApi.LinkPublicIp(context.Background()).LinkPublicIpRequest(linkPublicIpRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIpApi.LinkPublicIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkPublicIp`: LinkPublicIpResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicIpApi.LinkPublicIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkPublicIpRequest** | [**LinkPublicIpRequest**](LinkPublicIpRequest.md) |  | 

### Return type

[**LinkPublicIpResponse**](LinkPublicIpResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadPublicIpRanges

> ReadPublicIpRangesResponse ReadPublicIpRanges(ctx).ReadPublicIpRangesRequest(readPublicIpRangesRequest).Execute()



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
    readPublicIpRangesRequest := openapiclient.ReadPublicIpRangesRequest{DryRun: false} // ReadPublicIpRangesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicIpApi.ReadPublicIpRanges(context.Background()).ReadPublicIpRangesRequest(readPublicIpRangesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIpApi.ReadPublicIpRanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPublicIpRanges`: ReadPublicIpRangesResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicIpApi.ReadPublicIpRanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPublicIpRangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPublicIpRangesRequest** | [**ReadPublicIpRangesRequest**](ReadPublicIpRangesRequest.md) |  | 

### Return type

[**ReadPublicIpRangesResponse**](ReadPublicIpRangesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadPublicIps

> ReadPublicIpsResponse ReadPublicIps(ctx).ReadPublicIpsRequest(readPublicIpsRequest).Execute()



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
    readPublicIpsRequest := openapiclient.ReadPublicIpsRequest{DryRun: false, Filters: openapiclient.FiltersPublicIp{LinkPublicIpIds: []string{"LinkPublicIpIds_example"), NicAccountIds: []string{"NicAccountIds_example"), NicIds: []string{"NicIds_example"), Placements: []string{"Placements_example"), PrivateIps: []string{"PrivateIps_example"), PublicIpIds: []string{"PublicIpIds_example"), PublicIps: []string{"PublicIps_example"), TagKeys: []string{"TagKeys_example"), TagValues: []string{"TagValues_example"), Tags: []string{"Tags_example"), VmIds: []string{"VmIds_example")}} // ReadPublicIpsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicIpApi.ReadPublicIps(context.Background()).ReadPublicIpsRequest(readPublicIpsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIpApi.ReadPublicIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPublicIps`: ReadPublicIpsResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicIpApi.ReadPublicIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPublicIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPublicIpsRequest** | [**ReadPublicIpsRequest**](ReadPublicIpsRequest.md) |  | 

### Return type

[**ReadPublicIpsResponse**](ReadPublicIpsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkPublicIp

> UnlinkPublicIpResponse UnlinkPublicIp(ctx).UnlinkPublicIpRequest(unlinkPublicIpRequest).Execute()



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
    unlinkPublicIpRequest := openapiclient.UnlinkPublicIpRequest{DryRun: false, LinkPublicIpId: "LinkPublicIpId_example", PublicIp: "PublicIp_example"} // UnlinkPublicIpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicIpApi.UnlinkPublicIp(context.Background()).UnlinkPublicIpRequest(unlinkPublicIpRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIpApi.UnlinkPublicIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkPublicIp`: UnlinkPublicIpResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicIpApi.UnlinkPublicIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkPublicIpRequest** | [**UnlinkPublicIpRequest**](UnlinkPublicIpRequest.md) |  | 

### Return type

[**UnlinkPublicIpResponse**](UnlinkPublicIpResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

