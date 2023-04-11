# \VmTemplateApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVmTemplate**](VmTemplateApi.md#CreateVmTemplate) | **Post** /CreateVmTemplate | 
[**DeleteVmTemplate**](VmTemplateApi.md#DeleteVmTemplate) | **Post** /DeleteVmTemplate | 
[**ReadVmTemplates**](VmTemplateApi.md#ReadVmTemplates) | **Post** /ReadVmTemplates | 
[**UpdateVmTemplate**](VmTemplateApi.md#UpdateVmTemplate) | **Post** /UpdateVmTemplate | 



## CreateVmTemplate

> CreateVmTemplateResponse CreateVmTemplate(ctx).CreateVmTemplateRequest(createVmTemplateRequest).Execute()



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
    createVmTemplateRequest := *openapiclient.NewCreateVmTemplateRequest(int32(123), "CpuGeneration_example", "ImageId_example", int32(123), "VmTemplateName_example") // CreateVmTemplateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmTemplateApi.CreateVmTemplate(context.Background()).CreateVmTemplateRequest(createVmTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmTemplateApi.CreateVmTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVmTemplate`: CreateVmTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `VmTemplateApi.CreateVmTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVmTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVmTemplateRequest** | [**CreateVmTemplateRequest**](CreateVmTemplateRequest.md) |  | 

### Return type

[**CreateVmTemplateResponse**](CreateVmTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVmTemplate

> DeleteVmTemplateResponse DeleteVmTemplate(ctx).DeleteVmTemplateRequest(deleteVmTemplateRequest).Execute()



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
    deleteVmTemplateRequest := *openapiclient.NewDeleteVmTemplateRequest("VmTemplateId_example") // DeleteVmTemplateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmTemplateApi.DeleteVmTemplate(context.Background()).DeleteVmTemplateRequest(deleteVmTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmTemplateApi.DeleteVmTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVmTemplate`: DeleteVmTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `VmTemplateApi.DeleteVmTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVmTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVmTemplateRequest** | [**DeleteVmTemplateRequest**](DeleteVmTemplateRequest.md) |  | 

### Return type

[**DeleteVmTemplateResponse**](DeleteVmTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmTemplates

> ReadVmTemplatesResponse ReadVmTemplates(ctx).ReadVmTemplatesRequest(readVmTemplatesRequest).Execute()



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
    readVmTemplatesRequest := *openapiclient.NewReadVmTemplatesRequest() // ReadVmTemplatesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmTemplateApi.ReadVmTemplates(context.Background()).ReadVmTemplatesRequest(readVmTemplatesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmTemplateApi.ReadVmTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVmTemplates`: ReadVmTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `VmTemplateApi.ReadVmTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVmTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmTemplatesRequest** | [**ReadVmTemplatesRequest**](ReadVmTemplatesRequest.md) |  | 

### Return type

[**ReadVmTemplatesResponse**](ReadVmTemplatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVmTemplate

> UpdateVmTemplateResponse UpdateVmTemplate(ctx).UpdateVmTemplateRequest(updateVmTemplateRequest).Execute()



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
    updateVmTemplateRequest := *openapiclient.NewUpdateVmTemplateRequest("VmTemplateId_example") // UpdateVmTemplateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmTemplateApi.UpdateVmTemplate(context.Background()).UpdateVmTemplateRequest(updateVmTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmTemplateApi.UpdateVmTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVmTemplate`: UpdateVmTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `VmTemplateApi.UpdateVmTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVmTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVmTemplateRequest** | [**UpdateVmTemplateRequest**](UpdateVmTemplateRequest.md) |  | 

### Return type

[**UpdateVmTemplateResponse**](UpdateVmTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

