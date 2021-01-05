# \VolumeApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](VolumeApi.md#CreateVolume) | **Post** /CreateVolume | 
[**DeleteVolume**](VolumeApi.md#DeleteVolume) | **Post** /DeleteVolume | 
[**LinkVolume**](VolumeApi.md#LinkVolume) | **Post** /LinkVolume | 
[**ReadVolumes**](VolumeApi.md#ReadVolumes) | **Post** /ReadVolumes | 
[**UnlinkVolume**](VolumeApi.md#UnlinkVolume) | **Post** /UnlinkVolume | 
[**UpdateVolume**](VolumeApi.md#UpdateVolume) | **Post** /UpdateVolume | 



## CreateVolume

> CreateVolumeResponse CreateVolume(ctx).CreateVolumeRequest(createVolumeRequest).Execute()



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
    createVolumeRequest := *openapiclient.NewCreateVolumeRequest("SubregionName_example") // CreateVolumeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumeApi.CreateVolume(context.Background()).CreateVolumeRequest(createVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.CreateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolume`: CreateVolumeResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumeApi.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVolumeRequest** | [**CreateVolumeRequest**](CreateVolumeRequest.md) |  | 

### Return type

[**CreateVolumeResponse**](CreateVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> DeleteVolumeResponse DeleteVolume(ctx).DeleteVolumeRequest(deleteVolumeRequest).Execute()



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
    deleteVolumeRequest := *openapiclient.NewDeleteVolumeRequest("VolumeId_example") // DeleteVolumeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumeApi.DeleteVolume(context.Background()).DeleteVolumeRequest(deleteVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.DeleteVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVolume`: DeleteVolumeResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumeApi.DeleteVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVolumeRequest** | [**DeleteVolumeRequest**](DeleteVolumeRequest.md) |  | 

### Return type

[**DeleteVolumeResponse**](DeleteVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkVolume

> LinkVolumeResponse LinkVolume(ctx).LinkVolumeRequest(linkVolumeRequest).Execute()



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
    linkVolumeRequest := *openapiclient.NewLinkVolumeRequest("DeviceName_example", "VmId_example", "VolumeId_example") // LinkVolumeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumeApi.LinkVolume(context.Background()).LinkVolumeRequest(linkVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.LinkVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkVolume`: LinkVolumeResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumeApi.LinkVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkVolumeRequest** | [**LinkVolumeRequest**](LinkVolumeRequest.md) |  | 

### Return type

[**LinkVolumeResponse**](LinkVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadVolumes

> ReadVolumesResponse ReadVolumes(ctx).ReadVolumesRequest(readVolumesRequest).Execute()



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
    readVolumesRequest := *openapiclient.NewReadVolumesRequest() // ReadVolumesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumeApi.ReadVolumes(context.Background()).ReadVolumesRequest(readVolumesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.ReadVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadVolumes`: ReadVolumesResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumeApi.ReadVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readVolumesRequest** | [**ReadVolumesRequest**](ReadVolumesRequest.md) |  | 

### Return type

[**ReadVolumesResponse**](ReadVolumesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkVolume

> UnlinkVolumeResponse UnlinkVolume(ctx).UnlinkVolumeRequest(unlinkVolumeRequest).Execute()



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
    unlinkVolumeRequest := *openapiclient.NewUnlinkVolumeRequest("VolumeId_example") // UnlinkVolumeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumeApi.UnlinkVolume(context.Background()).UnlinkVolumeRequest(unlinkVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.UnlinkVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkVolume`: UnlinkVolumeResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumeApi.UnlinkVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkVolumeRequest** | [**UnlinkVolumeRequest**](UnlinkVolumeRequest.md) |  | 

### Return type

[**UnlinkVolumeResponse**](UnlinkVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> UpdateVolumeResponse UpdateVolume(ctx).UpdateVolumeRequest(updateVolumeRequest).Execute()



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
    updateVolumeRequest := *openapiclient.NewUpdateVolumeRequest("VolumeId_example") // UpdateVolumeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumeApi.UpdateVolume(context.Background()).UpdateVolumeRequest(updateVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.UpdateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVolume`: UpdateVolumeResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumeApi.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVolumeRequest** | [**UpdateVolumeRequest**](UpdateVolumeRequest.md) |  | 

### Return type

[**UpdateVolumeResponse**](UpdateVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

