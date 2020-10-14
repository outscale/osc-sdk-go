# \TaskApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExportTask**](TaskApi.md#DeleteExportTask) | **Post** /DeleteExportTask | 



## DeleteExportTask

> DeleteExportTaskResponse DeleteExportTask(ctx).DeleteExportTaskRequest(deleteExportTaskRequest).Execute()



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
    deleteExportTaskRequest := openapiclient.DeleteExportTaskRequest{DryRun: false, ExportTaskId: "ExportTaskId_example"} // DeleteExportTaskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.DeleteExportTask(context.Background()).DeleteExportTaskRequest(deleteExportTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.DeleteExportTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteExportTask`: DeleteExportTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.DeleteExportTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExportTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteExportTaskRequest** | [**DeleteExportTaskRequest**](DeleteExportTaskRequest.md) |  | 

### Return type

[**DeleteExportTaskResponse**](DeleteExportTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

