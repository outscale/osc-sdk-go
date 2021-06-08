# \SnapshotApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotApi.md#CreateSnapshot) | **Post** /CreateSnapshot | 
[**CreateSnapshotExportTask**](SnapshotApi.md#CreateSnapshotExportTask) | **Post** /CreateSnapshotExportTask | 
[**DeleteSnapshot**](SnapshotApi.md#DeleteSnapshot) | **Post** /DeleteSnapshot | 
[**ReadSnapshotExportTasks**](SnapshotApi.md#ReadSnapshotExportTasks) | **Post** /ReadSnapshotExportTasks | 
[**ReadSnapshots**](SnapshotApi.md#ReadSnapshots) | **Post** /ReadSnapshots | 
[**UpdateSnapshot**](SnapshotApi.md#UpdateSnapshot) | **Post** /UpdateSnapshot | 



## CreateSnapshot

> CreateSnapshotResponse CreateSnapshot(ctx).CreateSnapshotRequest(createSnapshotRequest).Execute()



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
    createSnapshotRequest := *openapiclient.NewCreateSnapshotRequest() // CreateSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnapshotApi.CreateSnapshot(context.Background()).CreateSnapshotRequest(createSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.CreateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnapshot`: CreateSnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotRequest** | [**CreateSnapshotRequest**](CreateSnapshotRequest.md) |  | 

### Return type

[**CreateSnapshotResponse**](CreateSnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshotExportTask

> CreateSnapshotExportTaskResponse CreateSnapshotExportTask(ctx).CreateSnapshotExportTaskRequest(createSnapshotExportTaskRequest).Execute()



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
    createSnapshotExportTaskRequest := *openapiclient.NewCreateSnapshotExportTaskRequest(*openapiclient.NewOsuExportToCreate("DiskImageFormat_example", "OsuBucket_example"), "SnapshotId_example") // CreateSnapshotExportTaskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnapshotApi.CreateSnapshotExportTask(context.Background()).CreateSnapshotExportTaskRequest(createSnapshotExportTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.CreateSnapshotExportTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnapshotExportTask`: CreateSnapshotExportTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.CreateSnapshotExportTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotExportTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotExportTaskRequest** | [**CreateSnapshotExportTaskRequest**](CreateSnapshotExportTaskRequest.md) |  | 

### Return type

[**CreateSnapshotExportTaskResponse**](CreateSnapshotExportTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshotResponse DeleteSnapshot(ctx).DeleteSnapshotRequest(deleteSnapshotRequest).Execute()



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
    deleteSnapshotRequest := *openapiclient.NewDeleteSnapshotRequest("SnapshotId_example") // DeleteSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnapshotApi.DeleteSnapshot(context.Background()).DeleteSnapshotRequest(deleteSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.DeleteSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSnapshot`: DeleteSnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.DeleteSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSnapshotRequest** | [**DeleteSnapshotRequest**](DeleteSnapshotRequest.md) |  | 

### Return type

[**DeleteSnapshotResponse**](DeleteSnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSnapshotExportTasks

> ReadSnapshotExportTasksResponse ReadSnapshotExportTasks(ctx).ReadSnapshotExportTasksRequest(readSnapshotExportTasksRequest).Execute()



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
    readSnapshotExportTasksRequest := *openapiclient.NewReadSnapshotExportTasksRequest() // ReadSnapshotExportTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnapshotApi.ReadSnapshotExportTasks(context.Background()).ReadSnapshotExportTasksRequest(readSnapshotExportTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.ReadSnapshotExportTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSnapshotExportTasks`: ReadSnapshotExportTasksResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.ReadSnapshotExportTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadSnapshotExportTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSnapshotExportTasksRequest** | [**ReadSnapshotExportTasksRequest**](ReadSnapshotExportTasksRequest.md) |  | 

### Return type

[**ReadSnapshotExportTasksResponse**](ReadSnapshotExportTasksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSnapshots

> ReadSnapshotsResponse ReadSnapshots(ctx).ReadSnapshotsRequest(readSnapshotsRequest).Execute()



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
    readSnapshotsRequest := *openapiclient.NewReadSnapshotsRequest() // ReadSnapshotsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnapshotApi.ReadSnapshots(context.Background()).ReadSnapshotsRequest(readSnapshotsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.ReadSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSnapshots`: ReadSnapshotsResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.ReadSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readSnapshotsRequest** | [**ReadSnapshotsRequest**](ReadSnapshotsRequest.md) |  | 

### Return type

[**ReadSnapshotsResponse**](ReadSnapshotsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> UpdateSnapshotResponse UpdateSnapshot(ctx).UpdateSnapshotRequest(updateSnapshotRequest).Execute()



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
    updateSnapshotRequest := *openapiclient.NewUpdateSnapshotRequest(*openapiclient.NewPermissionsOnResourceCreation(), "SnapshotId_example") // UpdateSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SnapshotApi.UpdateSnapshot(context.Background()).UpdateSnapshotRequest(updateSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.UpdateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSnapshot`: UpdateSnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.UpdateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSnapshotRequest** | [**UpdateSnapshotRequest**](UpdateSnapshotRequest.md) |  | 

### Return type

[**UpdateSnapshotResponse**](UpdateSnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

