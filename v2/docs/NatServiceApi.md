# \NatServiceApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNatService**](NatServiceApi.md#CreateNatService) | **Post** /CreateNatService | 
[**DeleteNatService**](NatServiceApi.md#DeleteNatService) | **Post** /DeleteNatService | 
[**ReadNatServices**](NatServiceApi.md#ReadNatServices) | **Post** /ReadNatServices | 



## CreateNatService

> CreateNatServiceResponse CreateNatService(ctx).CreateNatServiceRequest(createNatServiceRequest).Execute()



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
    createNatServiceRequest := *openapiclient.NewCreateNatServiceRequest("PublicIpId_example", "SubnetId_example") // CreateNatServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NatServiceApi.CreateNatService(context.Background()).CreateNatServiceRequest(createNatServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatServiceApi.CreateNatService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNatService`: CreateNatServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `NatServiceApi.CreateNatService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNatServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNatServiceRequest** | [**CreateNatServiceRequest**](CreateNatServiceRequest.md) |  | 

### Return type

[**CreateNatServiceResponse**](CreateNatServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNatService

> DeleteNatServiceResponse DeleteNatService(ctx).DeleteNatServiceRequest(deleteNatServiceRequest).Execute()



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
    deleteNatServiceRequest := *openapiclient.NewDeleteNatServiceRequest("NatServiceId_example") // DeleteNatServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NatServiceApi.DeleteNatService(context.Background()).DeleteNatServiceRequest(deleteNatServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatServiceApi.DeleteNatService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNatService`: DeleteNatServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `NatServiceApi.DeleteNatService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNatServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNatServiceRequest** | [**DeleteNatServiceRequest**](DeleteNatServiceRequest.md) |  | 

### Return type

[**DeleteNatServiceResponse**](DeleteNatServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNatServices

> ReadNatServicesResponse ReadNatServices(ctx).ReadNatServicesRequest(readNatServicesRequest).Execute()



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
    readNatServicesRequest := *openapiclient.NewReadNatServicesRequest() // ReadNatServicesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NatServiceApi.ReadNatServices(context.Background()).ReadNatServicesRequest(readNatServicesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatServiceApi.ReadNatServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNatServices`: ReadNatServicesResponse
    fmt.Fprintf(os.Stdout, "Response from `NatServiceApi.ReadNatServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadNatServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNatServicesRequest** | [**ReadNatServicesRequest**](ReadNatServicesRequest.md) |  | 

### Return type

[**ReadNatServicesResponse**](ReadNatServicesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

