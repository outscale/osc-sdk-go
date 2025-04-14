# \PolicyApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](PolicyApi.md#CreatePolicy) | **Post** /CreatePolicy | 
[**CreatePolicyVersion**](PolicyApi.md#CreatePolicyVersion) | **Post** /CreatePolicyVersion | 
[**DeletePolicy**](PolicyApi.md#DeletePolicy) | **Post** /DeletePolicy | 
[**DeletePolicyVersion**](PolicyApi.md#DeletePolicyVersion) | **Post** /DeletePolicyVersion | 
[**DeleteUserGroupPolicy**](PolicyApi.md#DeleteUserGroupPolicy) | **Post** /DeleteUserGroupPolicy | 
[**DeleteUserPolicy**](PolicyApi.md#DeleteUserPolicy) | **Post** /DeleteUserPolicy | 
[**LinkManagedPolicyToUserGroup**](PolicyApi.md#LinkManagedPolicyToUserGroup) | **Post** /LinkManagedPolicyToUserGroup | 
[**LinkPolicy**](PolicyApi.md#LinkPolicy) | **Post** /LinkPolicy | 
[**PutUserGroupPolicy**](PolicyApi.md#PutUserGroupPolicy) | **Post** /PutUserGroupPolicy | 
[**PutUserPolicy**](PolicyApi.md#PutUserPolicy) | **Post** /PutUserPolicy | 
[**ReadEntitiesLinkedToPolicy**](PolicyApi.md#ReadEntitiesLinkedToPolicy) | **Post** /ReadEntitiesLinkedToPolicy | 
[**ReadLinkedPolicies**](PolicyApi.md#ReadLinkedPolicies) | **Post** /ReadLinkedPolicies | 
[**ReadManagedPoliciesLinkedToUserGroup**](PolicyApi.md#ReadManagedPoliciesLinkedToUserGroup) | **Post** /ReadManagedPoliciesLinkedToUserGroup | 
[**ReadPolicies**](PolicyApi.md#ReadPolicies) | **Post** /ReadPolicies | 
[**ReadPolicy**](PolicyApi.md#ReadPolicy) | **Post** /ReadPolicy | 
[**ReadPolicyVersion**](PolicyApi.md#ReadPolicyVersion) | **Post** /ReadPolicyVersion | 
[**ReadPolicyVersions**](PolicyApi.md#ReadPolicyVersions) | **Post** /ReadPolicyVersions | 
[**ReadUserGroupPolicies**](PolicyApi.md#ReadUserGroupPolicies) | **Post** /ReadUserGroupPolicies | 
[**ReadUserGroupPolicy**](PolicyApi.md#ReadUserGroupPolicy) | **Post** /ReadUserGroupPolicy | 
[**ReadUserPolicies**](PolicyApi.md#ReadUserPolicies) | **Post** /ReadUserPolicies | 
[**ReadUserPolicy**](PolicyApi.md#ReadUserPolicy) | **Post** /ReadUserPolicy | 
[**SetDefaultPolicyVersion**](PolicyApi.md#SetDefaultPolicyVersion) | **Post** /SetDefaultPolicyVersion | 
[**UnlinkManagedPolicyFromUserGroup**](PolicyApi.md#UnlinkManagedPolicyFromUserGroup) | **Post** /UnlinkManagedPolicyFromUserGroup | 
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


## DeleteUserGroupPolicy

> DeleteUserGroupPolicyResponse DeleteUserGroupPolicy(ctx).DeleteUserGroupPolicyRequest(deleteUserGroupPolicyRequest).Execute()



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
    deleteUserGroupPolicyRequest := *openapiclient.NewDeleteUserGroupPolicyRequest("PolicyName_example", "UserGroupName_example") // DeleteUserGroupPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.DeleteUserGroupPolicy(context.Background()).DeleteUserGroupPolicyRequest(deleteUserGroupPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeleteUserGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserGroupPolicy`: DeleteUserGroupPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.DeleteUserGroupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserGroupPolicyRequest** | [**DeleteUserGroupPolicyRequest**](DeleteUserGroupPolicyRequest.md) |  | 

### Return type

[**DeleteUserGroupPolicyResponse**](DeleteUserGroupPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserPolicy

> DeleteUserPolicyResponse DeleteUserPolicy(ctx).DeleteUserPolicyRequest(deleteUserPolicyRequest).Execute()



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
    deleteUserPolicyRequest := *openapiclient.NewDeleteUserPolicyRequest("PolicyName_example", "UserName_example") // DeleteUserPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.DeleteUserPolicy(context.Background()).DeleteUserPolicyRequest(deleteUserPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeleteUserPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserPolicy`: DeleteUserPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.DeleteUserPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserPolicyRequest** | [**DeleteUserPolicyRequest**](DeleteUserPolicyRequest.md) |  | 

### Return type

[**DeleteUserPolicyResponse**](DeleteUserPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkManagedPolicyToUserGroup

> LinkManagedPolicyToUserGroupResponse LinkManagedPolicyToUserGroup(ctx).LinkManagedPolicyToUserGroupRequest(linkManagedPolicyToUserGroupRequest).Execute()



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
    linkManagedPolicyToUserGroupRequest := *openapiclient.NewLinkManagedPolicyToUserGroupRequest("PolicyOrn_example", "UserGroupName_example") // LinkManagedPolicyToUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.LinkManagedPolicyToUserGroup(context.Background()).LinkManagedPolicyToUserGroupRequest(linkManagedPolicyToUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.LinkManagedPolicyToUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkManagedPolicyToUserGroup`: LinkManagedPolicyToUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.LinkManagedPolicyToUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkManagedPolicyToUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkManagedPolicyToUserGroupRequest** | [**LinkManagedPolicyToUserGroupRequest**](LinkManagedPolicyToUserGroupRequest.md) |  | 

### Return type

[**LinkManagedPolicyToUserGroupResponse**](LinkManagedPolicyToUserGroupResponse.md)

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


## PutUserGroupPolicy

> PutUserGroupPolicyResponse PutUserGroupPolicy(ctx).PutUserGroupPolicyRequest(putUserGroupPolicyRequest).Execute()



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
    putUserGroupPolicyRequest := *openapiclient.NewPutUserGroupPolicyRequest("PolicyDocument_example", "PolicyName_example", "UserGroupName_example") // PutUserGroupPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.PutUserGroupPolicy(context.Background()).PutUserGroupPolicyRequest(putUserGroupPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.PutUserGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUserGroupPolicy`: PutUserGroupPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.PutUserGroupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUserGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putUserGroupPolicyRequest** | [**PutUserGroupPolicyRequest**](PutUserGroupPolicyRequest.md) |  | 

### Return type

[**PutUserGroupPolicyResponse**](PutUserGroupPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUserPolicy

> PutUserPolicyResponse PutUserPolicy(ctx).PutUserPolicyRequest(putUserPolicyRequest).Execute()



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
    putUserPolicyRequest := *openapiclient.NewPutUserPolicyRequest("PolicyDocument_example", "PolicyName_example", "UserName_example") // PutUserPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.PutUserPolicy(context.Background()).PutUserPolicyRequest(putUserPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.PutUserPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUserPolicy`: PutUserPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.PutUserPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUserPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putUserPolicyRequest** | [**PutUserPolicyRequest**](PutUserPolicyRequest.md) |  | 

### Return type

[**PutUserPolicyResponse**](PutUserPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEntitiesLinkedToPolicy

> ReadEntitiesLinkedToPolicyResponse ReadEntitiesLinkedToPolicy(ctx).ReadEntitiesLinkedToPolicyRequest(readEntitiesLinkedToPolicyRequest).Execute()



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
    readEntitiesLinkedToPolicyRequest := *openapiclient.NewReadEntitiesLinkedToPolicyRequest() // ReadEntitiesLinkedToPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadEntitiesLinkedToPolicy(context.Background()).ReadEntitiesLinkedToPolicyRequest(readEntitiesLinkedToPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadEntitiesLinkedToPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadEntitiesLinkedToPolicy`: ReadEntitiesLinkedToPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadEntitiesLinkedToPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadEntitiesLinkedToPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readEntitiesLinkedToPolicyRequest** | [**ReadEntitiesLinkedToPolicyRequest**](ReadEntitiesLinkedToPolicyRequest.md) |  | 

### Return type

[**ReadEntitiesLinkedToPolicyResponse**](ReadEntitiesLinkedToPolicyResponse.md)

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
    readLinkedPoliciesRequest := *openapiclient.NewReadLinkedPoliciesRequest("UserName_example") // ReadLinkedPoliciesRequest |  (optional)

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


## ReadManagedPoliciesLinkedToUserGroup

> ReadManagedPoliciesLinkedToUserGroupResponse ReadManagedPoliciesLinkedToUserGroup(ctx).ReadManagedPoliciesLinkedToUserGroupRequest(readManagedPoliciesLinkedToUserGroupRequest).Execute()



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
    readManagedPoliciesLinkedToUserGroupRequest := *openapiclient.NewReadManagedPoliciesLinkedToUserGroupRequest("UserGroupName_example") // ReadManagedPoliciesLinkedToUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadManagedPoliciesLinkedToUserGroup(context.Background()).ReadManagedPoliciesLinkedToUserGroupRequest(readManagedPoliciesLinkedToUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadManagedPoliciesLinkedToUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadManagedPoliciesLinkedToUserGroup`: ReadManagedPoliciesLinkedToUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadManagedPoliciesLinkedToUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadManagedPoliciesLinkedToUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readManagedPoliciesLinkedToUserGroupRequest** | [**ReadManagedPoliciesLinkedToUserGroupRequest**](ReadManagedPoliciesLinkedToUserGroupRequest.md) |  | 

### Return type

[**ReadManagedPoliciesLinkedToUserGroupResponse**](ReadManagedPoliciesLinkedToUserGroupResponse.md)

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


## ReadUserGroupPolicies

> ReadUserGroupPoliciesResponse ReadUserGroupPolicies(ctx).ReadUserGroupPoliciesRequest(readUserGroupPoliciesRequest).Execute()



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
    readUserGroupPoliciesRequest := *openapiclient.NewReadUserGroupPoliciesRequest("UserGroupName_example") // ReadUserGroupPoliciesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadUserGroupPolicies(context.Background()).ReadUserGroupPoliciesRequest(readUserGroupPoliciesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadUserGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserGroupPolicies`: ReadUserGroupPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadUserGroupPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadUserGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readUserGroupPoliciesRequest** | [**ReadUserGroupPoliciesRequest**](ReadUserGroupPoliciesRequest.md) |  | 

### Return type

[**ReadUserGroupPoliciesResponse**](ReadUserGroupPoliciesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserGroupPolicy

> ReadUserGroupPolicyResponse ReadUserGroupPolicy(ctx).ReadUserGroupPolicyRequest(readUserGroupPolicyRequest).Execute()



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
    readUserGroupPolicyRequest := *openapiclient.NewReadUserGroupPolicyRequest("PolicyName_example", "UserGroupName_example") // ReadUserGroupPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadUserGroupPolicy(context.Background()).ReadUserGroupPolicyRequest(readUserGroupPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadUserGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserGroupPolicy`: ReadUserGroupPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadUserGroupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadUserGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readUserGroupPolicyRequest** | [**ReadUserGroupPolicyRequest**](ReadUserGroupPolicyRequest.md) |  | 

### Return type

[**ReadUserGroupPolicyResponse**](ReadUserGroupPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserPolicies

> ReadUserPoliciesResponse ReadUserPolicies(ctx).ReadUserPoliciesRequest(readUserPoliciesRequest).Execute()



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
    readUserPoliciesRequest := *openapiclient.NewReadUserPoliciesRequest("UserName_example") // ReadUserPoliciesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadUserPolicies(context.Background()).ReadUserPoliciesRequest(readUserPoliciesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadUserPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserPolicies`: ReadUserPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadUserPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadUserPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readUserPoliciesRequest** | [**ReadUserPoliciesRequest**](ReadUserPoliciesRequest.md) |  | 

### Return type

[**ReadUserPoliciesResponse**](ReadUserPoliciesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserPolicy

> ReadUserPolicyResponse ReadUserPolicy(ctx).ReadUserPolicyRequest(readUserPolicyRequest).Execute()



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
    readUserPolicyRequest := *openapiclient.NewReadUserPolicyRequest("PolicyName_example", "UserName_example") // ReadUserPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ReadUserPolicy(context.Background()).ReadUserPolicyRequest(readUserPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReadUserPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserPolicy`: ReadUserPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReadUserPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadUserPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readUserPolicyRequest** | [**ReadUserPolicyRequest**](ReadUserPolicyRequest.md) |  | 

### Return type

[**ReadUserPolicyResponse**](ReadUserPolicyResponse.md)

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


## UnlinkManagedPolicyFromUserGroup

> UnlinkManagedPolicyFromUserGroupResponse UnlinkManagedPolicyFromUserGroup(ctx).UnlinkManagedPolicyFromUserGroupRequest(unlinkManagedPolicyFromUserGroupRequest).Execute()



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
    unlinkManagedPolicyFromUserGroupRequest := *openapiclient.NewUnlinkManagedPolicyFromUserGroupRequest("PolicyOrn_example", "UserGroupName_example") // UnlinkManagedPolicyFromUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.UnlinkManagedPolicyFromUserGroup(context.Background()).UnlinkManagedPolicyFromUserGroupRequest(unlinkManagedPolicyFromUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.UnlinkManagedPolicyFromUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkManagedPolicyFromUserGroup`: UnlinkManagedPolicyFromUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.UnlinkManagedPolicyFromUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkManagedPolicyFromUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlinkManagedPolicyFromUserGroupRequest** | [**UnlinkManagedPolicyFromUserGroupRequest**](UnlinkManagedPolicyFromUserGroupRequest.md) |  | 

### Return type

[**UnlinkManagedPolicyFromUserGroupResponse**](UnlinkManagedPolicyFromUserGroupResponse.md)

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

