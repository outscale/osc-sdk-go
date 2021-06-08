# \ImageApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](ImageApi.md#CreateImage) | **Post** /CreateImage | 
[**CreateImageExportTask**](ImageApi.md#CreateImageExportTask) | **Post** /CreateImageExportTask | 
[**DeleteImage**](ImageApi.md#DeleteImage) | **Post** /DeleteImage | 
[**ReadImageExportTasks**](ImageApi.md#ReadImageExportTasks) | **Post** /ReadImageExportTasks | 
[**ReadImages**](ImageApi.md#ReadImages) | **Post** /ReadImages | 
[**UpdateImage**](ImageApi.md#UpdateImage) | **Post** /UpdateImage | 



## CreateImage

> CreateImageResponse CreateImage(ctx).CreateImageRequest(createImageRequest).Execute()



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
    createImageRequest := *openapiclient.NewCreateImageRequest() // CreateImageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageApi.CreateImage(context.Background()).CreateImageRequest(createImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.CreateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImage`: CreateImageResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.CreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImageRequest** | [**CreateImageRequest**](CreateImageRequest.md) |  | 

### Return type

[**CreateImageResponse**](CreateImageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateImageExportTask

> CreateImageExportTaskResponse CreateImageExportTask(ctx).CreateImageExportTaskRequest(createImageExportTaskRequest).Execute()



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
    createImageExportTaskRequest := *openapiclient.NewCreateImageExportTaskRequest("ImageId_example", *openapiclient.NewOsuExportToCreate("DiskImageFormat_example", "OsuBucket_example")) // CreateImageExportTaskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageApi.CreateImageExportTask(context.Background()).CreateImageExportTaskRequest(createImageExportTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.CreateImageExportTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImageExportTask`: CreateImageExportTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.CreateImageExportTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageExportTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImageExportTaskRequest** | [**CreateImageExportTaskRequest**](CreateImageExportTaskRequest.md) |  | 

### Return type

[**CreateImageExportTaskResponse**](CreateImageExportTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImageResponse DeleteImage(ctx).DeleteImageRequest(deleteImageRequest).Execute()



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
    deleteImageRequest := *openapiclient.NewDeleteImageRequest("ImageId_example") // DeleteImageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageApi.DeleteImage(context.Background()).DeleteImageRequest(deleteImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.DeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteImage`: DeleteImageResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.DeleteImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteImageRequest** | [**DeleteImageRequest**](DeleteImageRequest.md) |  | 

### Return type

[**DeleteImageResponse**](DeleteImageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadImageExportTasks

> ReadImageExportTasksResponse ReadImageExportTasks(ctx).ReadImageExportTasksRequest(readImageExportTasksRequest).Execute()



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
    readImageExportTasksRequest := *openapiclient.NewReadImageExportTasksRequest() // ReadImageExportTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageApi.ReadImageExportTasks(context.Background()).ReadImageExportTasksRequest(readImageExportTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.ReadImageExportTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadImageExportTasks`: ReadImageExportTasksResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.ReadImageExportTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadImageExportTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readImageExportTasksRequest** | [**ReadImageExportTasksRequest**](ReadImageExportTasksRequest.md) |  | 

### Return type

[**ReadImageExportTasksResponse**](ReadImageExportTasksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadImages

> ReadImagesResponse ReadImages(ctx).ReadImagesRequest(readImagesRequest).Execute()



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
    readImagesRequest := *openapiclient.NewReadImagesRequest() // ReadImagesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageApi.ReadImages(context.Background()).ReadImagesRequest(readImagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.ReadImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadImages`: ReadImagesResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.ReadImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readImagesRequest** | [**ReadImagesRequest**](ReadImagesRequest.md) |  | 

### Return type

[**ReadImagesResponse**](ReadImagesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImage

> UpdateImageResponse UpdateImage(ctx).UpdateImageRequest(updateImageRequest).Execute()



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
    updateImageRequest := *openapiclient.NewUpdateImageRequest("ImageId_example", *openapiclient.NewPermissionsOnResourceCreation()) // UpdateImageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageApi.UpdateImage(context.Background()).UpdateImageRequest(updateImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.UpdateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImage`: UpdateImageResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.UpdateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateImageRequest** | [**UpdateImageRequest**](UpdateImageRequest.md) |  | 

### Return type

[**UpdateImageResponse**](UpdateImageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

