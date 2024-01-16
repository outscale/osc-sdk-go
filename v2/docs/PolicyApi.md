# \PolicyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](PolicyApi.md#CreatePolicy) | **Post** /CreatePolicy | 
[**CreatePolicyVersion**](PolicyApi.md#CreatePolicyVersion) | **Post** /CreatePolicyVersion | 
[**DeletePolicy**](PolicyApi.md#DeletePolicy) | **Post** /DeletePolicy | 
[**DeletePolicyVersion**](PolicyApi.md#DeletePolicyVersion) | **Post** /DeletePolicyVersion | 
[**LinkPolicy**](PolicyApi.md#LinkPolicy) | **Post** /LinkPolicy | 
[**ReadLinkedPolicies**](PolicyApi.md#ReadLinkedPolicies) | **Post** /ReadLinkedPolicies | 
[**ReadPolicies**](PolicyApi.md#ReadPolicies) | **Post** /ReadPolicies | 
[**ReadPolicy**](PolicyApi.md#ReadPolicy) | **Post** /ReadPolicy | 
[**ReadPolicyVersion**](PolicyApi.md#ReadPolicyVersion) | **Post** /ReadPolicyVersion | 
[**ReadPolicyVersions**](PolicyApi.md#ReadPolicyVersions) | **Post** /ReadPolicyVersions | 
[**SetDefaultPolicyVersion**](PolicyApi.md#SetDefaultPolicyVersion) | **Post** /SetDefaultPolicyVersion | 
[**UnlinkPolicy**](PolicyApi.md#UnlinkPolicy) | **Post** /UnlinkPolicy | 



## CreatePolicy

> CreatePolicyResponse CreatePolicy(ctx).CreatePolicyRequest(createPolicyRequest).Execute()



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
    createPolicyRequest := *openapiclient.NewCreatePolicyRequest("Document_example", "PolicyName_example") // CreatePolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.CreatePolicy(context.Background()).CreatePolicyRequest(createPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.CreatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicy`: CreatePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPolicyRequest** | [**CreatePolicyRequest**](CreatePolicyRequest.md) |  | 

### Return type

[**CreatePolicyResponse**](CreatePolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyVersion

> CreatePolicyVersionResponse CreatePolicyVersion(ctx).CreatePolicyVersionRequest(createPolicyVersionRequest).Execute()



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
    createPolicyVersionRequest := *openapiclient.NewCreatePolicyVersionRequest("Document_example", "PolicyOrn_example") // CreatePolicyVersionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.CreatePolicyVersion(context.Background()).CreatePolicyVersionRequest(createPolicyVersionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.CreatePolicyVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyVersion`: CreatePolicyVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.CreatePolicyVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPolicyVersionRequest** | [**CreatePolicyVersionRequest**](CreatePolicyVersionRequest.md) |  | 

### Return type

[**CreatePolicyVersionResponse**](CreatePolicyVersionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicyResponse DeletePolicy(ctx).DeletePolicyRequest(deletePolicyRequest).Execute()



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
    deletePolicyRequest := *openapiclient.NewDeletePolicyRequest("PolicyOrn_example") // DeletePolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.DeletePolicy(context.Background()).DeletePolicyRequest(deletePolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicy`: DeletePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.DeletePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePolicyRequest** | [**DeletePolicyRequest**](DeletePolicyRequest.md) |  | 

### Return type

[**DeletePolicyResponse**](DeletePolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyVersion

> DeletePolicyVersionResponse DeletePolicyVersion(ctx).DeletePolicyVersionRequest(deletePolicyVersionRequest).Execute()



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
    deletePolicyVersionRequest := *openapiclient.NewDeletePolicyVersionRequest("PolicyOrn_example", "VersionId_example") // DeletePolicyVersionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.DeletePolicyVersion(context.Background()).DeletePolicyVersionRequest(deletePolicyVersionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicyVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicyVersion`: DeletePolicyVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.DeletePolicyVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePolicyVersionRequest** | [**DeletePolicyVersionRequest**](DeletePolicyVersionRequest.md) |  | 

### Return type

[**DeletePolicyVersionResponse**](DeletePolicyVersionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkPolicy

> LinkPolicyResponse LinkPolicy(ctx).LinkPolicyRequest(linkPolicyRequest).Execute()



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
    linkPolicyRequest := *openapiclient.NewLinkPolicyRequest("PolicyOrn_example", "UserName_example") // LinkPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.LinkPolicy(context.Background()).LinkPolicyRequest(linkPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.LinkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkPolicy`: LinkPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.LinkPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkPolicyRequest** | [**LinkPolicyRequest**](LinkPolicyRequest.md) |  | 

### Return type

[**LinkPolicyResponse**](LinkPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadLinkedPolicies

> ReadLinkedPoliciesResponse ReadLinkedPolicies(ctx).ReadLinkedPoliciesRequest(readLinkedPoliciesRequest).Execute()



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
    readLinkedPoliciesRequest := *openapiclient.NewReadLinkedPoliciesRequest() // ReadLinkedPoliciesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadLinkedPolicies(context.Background()).ReadLinkedPoliciesRequest(readLinkedPoliciesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadLinkedPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadLinkedPolicies`: ReadLinkedPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadLinkedPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadLinkedPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLinkedPoliciesRequest** | [**ReadLinkedPoliciesRequest**](ReadLinkedPoliciesRequest.md) |  | 

### Return type

[**ReadLinkedPoliciesResponse**](ReadLinkedPoliciesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadPolicies

> ReadPoliciesResponse ReadPolicies(ctx).ReadPoliciesRequest(readPoliciesRequest).Execute()



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
    readPoliciesRequest := *openapiclient.NewReadPoliciesRequest() // ReadPoliciesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadPolicies(context.Background()).ReadPoliciesRequest(readPoliciesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPolicies`: ReadPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPoliciesRequest** | [**ReadPoliciesRequest**](ReadPoliciesRequest.md) |  | 

### Return type

[**ReadPoliciesResponse**](ReadPoliciesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadPolicy

> ReadPolicyResponse ReadPolicy(ctx).ReadPolicyRequest(readPolicyRequest).Execute()



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
    readPolicyRequest := *openapiclient.NewReadPolicyRequest("PolicyOrn_example") // ReadPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadPolicy(context.Background()).ReadPolicyRequest(readPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPolicy`: ReadPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPolicyRequest** | [**ReadPolicyRequest**](ReadPolicyRequest.md) |  | 

### Return type

[**ReadPolicyResponse**](ReadPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadPolicyVersion

> ReadPolicyVersionResponse ReadPolicyVersion(ctx).ReadPolicyVersionRequest(readPolicyVersionRequest).Execute()



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
    readPolicyVersionRequest := *openapiclient.NewReadPolicyVersionRequest("PolicyOrn_example", "VersionId_example") // ReadPolicyVersionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadPolicyVersion(context.Background()).ReadPolicyVersionRequest(readPolicyVersionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadPolicyVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPolicyVersion`: ReadPolicyVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadPolicyVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPolicyVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPolicyVersionRequest** | [**ReadPolicyVersionRequest**](ReadPolicyVersionRequest.md) |  | 

### Return type

[**ReadPolicyVersionResponse**](ReadPolicyVersionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadPolicyVersions

> ReadPolicyVersionsResponse ReadPolicyVersions(ctx).ReadPolicyVersionsRequest(readPolicyVersionsRequest).Execute()



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
    readPolicyVersionsRequest := *openapiclient.NewReadPolicyVersionsRequest("PolicyOrn_example") // ReadPolicyVersionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadPolicyVersions(context.Background()).ReadPolicyVersionsRequest(readPolicyVersionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadPolicyVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPolicyVersions`: ReadPolicyVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadPolicyVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPolicyVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readPolicyVersionsRequest** | [**ReadPolicyVersionsRequest**](ReadPolicyVersionsRequest.md) |  | 

### Return type

[**ReadPolicyVersionsResponse**](ReadPolicyVersionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultPolicyVersion

> SetDefaultPolicyVersionResponse SetDefaultPolicyVersion(ctx).SetDefaultPolicyVersionRequest(setDefaultPolicyVersionRequest).Execute()



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
    setDefaultPolicyVersionRequest := *openapiclient.NewSetDefaultPolicyVersionRequest("PolicyOrn_example", "VersionId_example") // SetDefaultPolicyVersionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.SetDefaultPolicyVersion(context.Background()).SetDefaultPolicyVersionRequest(setDefaultPolicyVersionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.SetDefaultPolicyVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDefaultPolicyVersion`: SetDefaultPolicyVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.SetDefaultPolicyVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultPolicyVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultPolicyVersionRequest** | [**SetDefaultPolicyVersionRequest**](SetDefaultPolicyVersionRequest.md) |  | 

### Return type

[**SetDefaultPolicyVersionResponse**](SetDefaultPolicyVersionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkPolicy

> UnlinkPolicyResponse UnlinkPolicy(ctx).UnlinkPolicyRequest(unlinkPolicyRequest).Execute()



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
    unlinkPolicyRequest := *openapiclient.NewUnlinkPolicyRequest("PolicyOrn_example", "UserName_example") // UnlinkPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.UnlinkPolicy(context.Background()).UnlinkPolicyRequest(unlinkPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.UnlinkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkPolicy`: UnlinkPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.UnlinkPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkPolicyRequest** | [**UnlinkPolicyRequest**](UnlinkPolicyRequest.md) |  | 

### Return type

[**UnlinkPolicyResponse**](UnlinkPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

