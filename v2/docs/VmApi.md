# \VmApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVms**](VmApi.md#CreateVms) | **Post** /CreateVms | 
[**DeleteVms**](VmApi.md#DeleteVms) | **Post** /DeleteVms | 
[**ReadAdminPassword**](VmApi.md#ReadAdminPassword) | **Post** /ReadAdminPassword | 
[**ReadConsoleOutput**](VmApi.md#ReadConsoleOutput) | **Post** /ReadConsoleOutput | 
[**ReadVmTypes**](VmApi.md#ReadVmTypes) | **Post** /ReadVmTypes | 
[**ReadVms**](VmApi.md#ReadVms) | **Post** /ReadVms | 
[**ReadVmsState**](VmApi.md#ReadVmsState) | **Post** /ReadVmsState | 
[**RebootVms**](VmApi.md#RebootVms) | **Post** /RebootVms | 
[**StartVms**](VmApi.md#StartVms) | **Post** /StartVms | 
[**StopVms**](VmApi.md#StopVms) | **Post** /StopVms | 
[**UpdateVm**](VmApi.md#UpdateVm) | **Post** /UpdateVm | 



## CreateVms

> CreateVmsResponse CreateVms(ctx).CreateVmsRequest(createVmsRequest).Execute()



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
    createVmsRequest := *openapiclient.NewCreateVmsRequest("ImageId_example") // CreateVmsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.CreateVms(context.Background()).CreateVmsRequest(createVmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.CreateVms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVms`: CreateVmsResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.CreateVms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVmsRequest** | [**CreateVmsRequest**](CreateVmsRequest.md) |  | 

### Return type

[**CreateVmsResponse**](CreateVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVms

> DeleteVmsResponse DeleteVms(ctx).DeleteVmsRequest(deleteVmsRequest).Execute()



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
    deleteVmsRequest := *openapiclient.NewDeleteVmsRequest([]string{"VmIds_example")) // DeleteVmsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.DeleteVms(context.Background()).DeleteVmsRequest(deleteVmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.DeleteVms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVms`: DeleteVmsResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.DeleteVms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVmsRequest** | [**DeleteVmsRequest**](DeleteVmsRequest.md) |  | 

### Return type

[**DeleteVmsResponse**](DeleteVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAdminPassword

> ReadAdminPasswordResponse ReadAdminPassword(ctx).ReadAdminPasswordRequest(readAdminPasswordRequest).Execute()



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
    readAdminPasswordRequest := *openapiclient.NewReadAdminPasswordRequest("VmId_example") // ReadAdminPasswordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.ReadAdminPassword(context.Background()).ReadAdminPasswordRequest(readAdminPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.ReadAdminPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAdminPassword`: ReadAdminPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.ReadAdminPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadAdminPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readAdminPasswordRequest** | [**ReadAdminPasswordRequest**](ReadAdminPasswordRequest.md) |  | 

### Return type

[**ReadAdminPasswordResponse**](ReadAdminPasswordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadConsoleOutput

> ReadConsoleOutputResponse ReadConsoleOutput(ctx).ReadConsoleOutputRequest(readConsoleOutputRequest).Execute()



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
    readConsoleOutputRequest := *openapiclient.NewReadConsoleOutputRequest("VmId_example") // ReadConsoleOutputRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.ReadConsoleOutput(context.Background()).ReadConsoleOutputRequest(readConsoleOutputRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.ReadConsoleOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadConsoleOutput`: ReadConsoleOutputResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.ReadConsoleOutput`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadConsoleOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readConsoleOutputRequest** | [**ReadConsoleOutputRequest**](ReadConsoleOutputRequest.md) |  | 

### Return type

[**ReadConsoleOutputResponse**](ReadConsoleOutputResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmTypes

> ReadVmTypesResponse ReadVmTypes(ctx).ReadVmTypesRequest(readVmTypesRequest).Execute()



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
    readVmTypesRequest := *openapiclient.NewReadVmTypesRequest() // ReadVmTypesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.ReadVmTypes(context.Background()).ReadVmTypesRequest(readVmTypesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.ReadVmTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVmTypes`: ReadVmTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.ReadVmTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVmTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmTypesRequest** | [**ReadVmTypesRequest**](ReadVmTypesRequest.md) |  | 

### Return type

[**ReadVmTypesResponse**](ReadVmTypesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVms

> ReadVmsResponse ReadVms(ctx).ReadVmsRequest(readVmsRequest).Execute()



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
    readVmsRequest := *openapiclient.NewReadVmsRequest() // ReadVmsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.ReadVms(context.Background()).ReadVmsRequest(readVmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.ReadVms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVms`: ReadVmsResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.ReadVms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmsRequest** | [**ReadVmsRequest**](ReadVmsRequest.md) |  | 

### Return type

[**ReadVmsResponse**](ReadVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVmsState

> ReadVmsStateResponse ReadVmsState(ctx).ReadVmsStateRequest(readVmsStateRequest).Execute()



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
    readVmsStateRequest := *openapiclient.NewReadVmsStateRequest() // ReadVmsStateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.ReadVmsState(context.Background()).ReadVmsStateRequest(readVmsStateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.ReadVmsState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVmsState`: ReadVmsStateResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.ReadVmsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVmsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVmsStateRequest** | [**ReadVmsStateRequest**](ReadVmsStateRequest.md) |  | 

### Return type

[**ReadVmsStateResponse**](ReadVmsStateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootVms

> RebootVmsResponse RebootVms(ctx).RebootVmsRequest(rebootVmsRequest).Execute()



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
    rebootVmsRequest := *openapiclient.NewRebootVmsRequest([]string{"VmIds_example")) // RebootVmsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.RebootVms(context.Background()).RebootVmsRequest(rebootVmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.RebootVms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootVms`: RebootVmsResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.RebootVms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRebootVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rebootVmsRequest** | [**RebootVmsRequest**](RebootVmsRequest.md) |  | 

### Return type

[**RebootVmsResponse**](RebootVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVms

> StartVmsResponse StartVms(ctx).StartVmsRequest(startVmsRequest).Execute()



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
    startVmsRequest := *openapiclient.NewStartVmsRequest([]string{"VmIds_example")) // StartVmsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.StartVms(context.Background()).StartVmsRequest(startVmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.StartVms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartVms`: StartVmsResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.StartVms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startVmsRequest** | [**StartVmsRequest**](StartVmsRequest.md) |  | 

### Return type

[**StartVmsResponse**](StartVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVms

> StopVmsResponse StopVms(ctx).StopVmsRequest(stopVmsRequest).Execute()



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
    stopVmsRequest := *openapiclient.NewStopVmsRequest([]string{"VmIds_example")) // StopVmsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.StopVms(context.Background()).StopVmsRequest(stopVmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.StopVms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopVms`: StopVmsResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.StopVms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopVmsRequest** | [**StopVmsRequest**](StopVmsRequest.md) |  | 

### Return type

[**StopVmsResponse**](StopVmsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVm

> UpdateVmResponse UpdateVm(ctx).UpdateVmRequest(updateVmRequest).Execute()



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
    updateVmRequest := *openapiclient.NewUpdateVmRequest("VmId_example") // UpdateVmRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmApi.UpdateVm(context.Background()).UpdateVmRequest(updateVmRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmApi.UpdateVm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVm`: UpdateVmResponse
    fmt.Fprintf(os.Stdout, "Response from `VmApi.UpdateVm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVmRequest** | [**UpdateVmRequest**](UpdateVmRequest.md) |  | 

### Return type

[**UpdateVmResponse**](UpdateVmResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

