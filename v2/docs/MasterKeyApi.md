# \MasterKeyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMasterKey**](MasterKeyApi.md#CreateMasterKey) | **Post** /CreateMasterKey | 
[**DecryptCiphertext**](MasterKeyApi.md#DecryptCiphertext) | **Post** /DecryptCiphertext | 
[**DeleteMasterKey**](MasterKeyApi.md#DeleteMasterKey) | **Post** /DeleteMasterKey | 
[**EncryptPlaintext**](MasterKeyApi.md#EncryptPlaintext) | **Post** /EncryptPlaintext | 
[**GenerateDataKey**](MasterKeyApi.md#GenerateDataKey) | **Post** /GenerateDataKey | 
[**ReadMasterKeys**](MasterKeyApi.md#ReadMasterKeys) | **Post** /ReadMasterKeys | 
[**UndeleteMasterKey**](MasterKeyApi.md#UndeleteMasterKey) | **Post** /UndeleteMasterKey | 
[**UpdateMasterKey**](MasterKeyApi.md#UpdateMasterKey) | **Post** /UpdateMasterKey | 



## CreateMasterKey

> CreateMasterKeyResponse CreateMasterKey(ctx).CreateMasterKeyRequest(createMasterKeyRequest).Execute()



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
    createMasterKeyRequest := *openapiclient.NewCreateMasterKeyRequest() // CreateMasterKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MasterKeyApi.CreateMasterKey(context.Background()).CreateMasterKeyRequest(createMasterKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterKeyApi.CreateMasterKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMasterKey`: CreateMasterKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `MasterKeyApi.CreateMasterKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMasterKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMasterKeyRequest** | [**CreateMasterKeyRequest**](CreateMasterKeyRequest.md) |  | 

### Return type

[**CreateMasterKeyResponse**](CreateMasterKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecryptCiphertext

> DecryptCiphertextResponse DecryptCiphertext(ctx).DecryptCiphertextRequest(decryptCiphertextRequest).Execute()



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
    decryptCiphertextRequest := *openapiclient.NewDecryptCiphertextRequest(string(123)) // DecryptCiphertextRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MasterKeyApi.DecryptCiphertext(context.Background()).DecryptCiphertextRequest(decryptCiphertextRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterKeyApi.DecryptCiphertext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecryptCiphertext`: DecryptCiphertextResponse
    fmt.Fprintf(os.Stdout, "Response from `MasterKeyApi.DecryptCiphertext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecryptCiphertextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decryptCiphertextRequest** | [**DecryptCiphertextRequest**](DecryptCiphertextRequest.md) |  | 

### Return type

[**DecryptCiphertextResponse**](DecryptCiphertextResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMasterKey

> DeleteMasterKeyResponse DeleteMasterKey(ctx).DeleteMasterKeyRequest(deleteMasterKeyRequest).Execute()



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
    deleteMasterKeyRequest := *openapiclient.NewDeleteMasterKeyRequest("MasterKeyId_example") // DeleteMasterKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MasterKeyApi.DeleteMasterKey(context.Background()).DeleteMasterKeyRequest(deleteMasterKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterKeyApi.DeleteMasterKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMasterKey`: DeleteMasterKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `MasterKeyApi.DeleteMasterKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMasterKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMasterKeyRequest** | [**DeleteMasterKeyRequest**](DeleteMasterKeyRequest.md) |  | 

### Return type

[**DeleteMasterKeyResponse**](DeleteMasterKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncryptPlaintext

> EncryptPlaintextResponse EncryptPlaintext(ctx).EncryptPlaintextRequest(encryptPlaintextRequest).Execute()



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
    encryptPlaintextRequest := *openapiclient.NewEncryptPlaintextRequest("MasterKeyId_example", string(123)) // EncryptPlaintextRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MasterKeyApi.EncryptPlaintext(context.Background()).EncryptPlaintextRequest(encryptPlaintextRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterKeyApi.EncryptPlaintext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncryptPlaintext`: EncryptPlaintextResponse
    fmt.Fprintf(os.Stdout, "Response from `MasterKeyApi.EncryptPlaintext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEncryptPlaintextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptPlaintextRequest** | [**EncryptPlaintextRequest**](EncryptPlaintextRequest.md) |  | 

### Return type

[**EncryptPlaintextResponse**](EncryptPlaintextResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateDataKey

> GenerateDataKeyResponse GenerateDataKey(ctx).GenerateDataKeyRequest(generateDataKeyRequest).Execute()



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
    generateDataKeyRequest := *openapiclient.NewGenerateDataKeyRequest("MasterKeyId_example", int32(123)) // GenerateDataKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MasterKeyApi.GenerateDataKey(context.Background()).GenerateDataKeyRequest(generateDataKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterKeyApi.GenerateDataKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDataKey`: GenerateDataKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `MasterKeyApi.GenerateDataKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDataKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateDataKeyRequest** | [**GenerateDataKeyRequest**](GenerateDataKeyRequest.md) |  | 

### Return type

[**GenerateDataKeyResponse**](GenerateDataKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMasterKeys

> ReadMasterKeysResponse ReadMasterKeys(ctx).ReadMasterKeysRequest(readMasterKeysRequest).Execute()



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
    readMasterKeysRequest := *openapiclient.NewReadMasterKeysRequest() // ReadMasterKeysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MasterKeyApi.ReadMasterKeys(context.Background()).ReadMasterKeysRequest(readMasterKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterKeyApi.ReadMasterKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMasterKeys`: ReadMasterKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `MasterKeyApi.ReadMasterKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadMasterKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readMasterKeysRequest** | [**ReadMasterKeysRequest**](ReadMasterKeysRequest.md) |  | 

### Return type

[**ReadMasterKeysResponse**](ReadMasterKeysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeleteMasterKey

> UndeleteMasterKeyResponse UndeleteMasterKey(ctx).UndeleteMasterKeyRequest(undeleteMasterKeyRequest).Execute()



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
    undeleteMasterKeyRequest := *openapiclient.NewUndeleteMasterKeyRequest("MasterKeyId_example") // UndeleteMasterKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MasterKeyApi.UndeleteMasterKey(context.Background()).UndeleteMasterKeyRequest(undeleteMasterKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterKeyApi.UndeleteMasterKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UndeleteMasterKey`: UndeleteMasterKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `MasterKeyApi.UndeleteMasterKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteMasterKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **undeleteMasterKeyRequest** | [**UndeleteMasterKeyRequest**](UndeleteMasterKeyRequest.md) |  | 

### Return type

[**UndeleteMasterKeyResponse**](UndeleteMasterKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMasterKey

> UpdateMasterKeyResponse UpdateMasterKey(ctx).UpdateMasterKeyRequest(updateMasterKeyRequest).Execute()



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
    updateMasterKeyRequest := *openapiclient.NewUpdateMasterKeyRequest("MasterKeyId_example") // UpdateMasterKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MasterKeyApi.UpdateMasterKey(context.Background()).UpdateMasterKeyRequest(updateMasterKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterKeyApi.UpdateMasterKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMasterKey`: UpdateMasterKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `MasterKeyApi.UpdateMasterKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMasterKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMasterKeyRequest** | [**UpdateMasterKeyRequest**](UpdateMasterKeyRequest.md) |  | 

### Return type

[**UpdateMasterKeyResponse**](UpdateMasterKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

