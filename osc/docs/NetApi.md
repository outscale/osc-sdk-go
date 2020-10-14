# \NetApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNet**](NetApi.md#CreateNet) | **Post** /CreateNet | 
[**DeleteNet**](NetApi.md#DeleteNet) | **Post** /DeleteNet | 
[**ReadNets**](NetApi.md#ReadNets) | **Post** /ReadNets | 
[**UpdateNet**](NetApi.md#UpdateNet) | **Post** /UpdateNet | 



## CreateNet

> CreateNetResponse CreateNet(ctx).CreateNetRequest(createNetRequest).Execute()



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
    createNetRequest := openapiclient.CreateNetRequest{DryRun: false, IpRange: "IpRange_example", Tenancy: "Tenancy_example"} // CreateNetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetApi.CreateNet(context.Background()).CreateNetRequest(createNetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetApi.CreateNet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNet`: CreateNetResponse
    fmt.Fprintf(os.Stdout, "Response from `NetApi.CreateNet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetRequest** | [**CreateNetRequest**](CreateNetRequest.md) |  | 

### Return type

[**CreateNetResponse**](CreateNetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNet

> DeleteNetResponse DeleteNet(ctx).DeleteNetRequest(deleteNetRequest).Execute()



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
    deleteNetRequest := openapiclient.DeleteNetRequest{DryRun: false, NetId: "NetId_example"} // DeleteNetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetApi.DeleteNet(context.Background()).DeleteNetRequest(deleteNetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetApi.DeleteNet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNet`: DeleteNetResponse
    fmt.Fprintf(os.Stdout, "Response from `NetApi.DeleteNet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNetRequest** | [**DeleteNetRequest**](DeleteNetRequest.md) |  | 

### Return type

[**DeleteNetResponse**](DeleteNetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNets

> ReadNetsResponse ReadNets(ctx).ReadNetsRequest(readNetsRequest).Execute()



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
    readNetsRequest := openapiclient.ReadNetsRequest{DryRun: false, Filters: openapiclient.FiltersNet{DhcpOptionsSetIds: []string{"DhcpOptionsSetIds_example"), IpRanges: []string{"IpRanges_example"), IsDefault: false, NetIds: []string{"NetIds_example"), States: []string{"States_example"), TagKeys: []string{"TagKeys_example"), TagValues: []string{"TagValues_example"), Tags: []string{"Tags_example")}} // ReadNetsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetApi.ReadNets(context.Background()).ReadNetsRequest(readNetsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetApi.ReadNets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNets`: ReadNetsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetApi.ReadNets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadNetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNetsRequest** | [**ReadNetsRequest**](ReadNetsRequest.md) |  | 

### Return type

[**ReadNetsResponse**](ReadNetsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNet

> UpdateNetResponse UpdateNet(ctx).UpdateNetRequest(updateNetRequest).Execute()



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
    updateNetRequest := openapiclient.UpdateNetRequest{DhcpOptionsSetId: "DhcpOptionsSetId_example", DryRun: false, NetId: "NetId_example"} // UpdateNetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetApi.UpdateNet(context.Background()).UpdateNetRequest(updateNetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetApi.UpdateNet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNet`: UpdateNetResponse
    fmt.Fprintf(os.Stdout, "Response from `NetApi.UpdateNet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNetRequest** | [**UpdateNetRequest**](UpdateNetRequest.md) |  | 

### Return type

[**UpdateNetResponse**](UpdateNetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

