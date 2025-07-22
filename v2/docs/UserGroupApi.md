# \UserGroupApi

All URIs are relative to *https://api.eu-west-2.outscale.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToUserGroup**](UserGroupApi.md#AddUserToUserGroup) | **Post** /AddUserToUserGroup | 
[**CreateUserGroup**](UserGroupApi.md#CreateUserGroup) | **Post** /CreateUserGroup | 
[**DeleteUserGroup**](UserGroupApi.md#DeleteUserGroup) | **Post** /DeleteUserGroup | 
[**ReadUserGroup**](UserGroupApi.md#ReadUserGroup) | **Post** /ReadUserGroup | 
[**ReadUserGroups**](UserGroupApi.md#ReadUserGroups) | **Post** /ReadUserGroups | 
[**ReadUserGroupsPerUser**](UserGroupApi.md#ReadUserGroupsPerUser) | **Post** /ReadUserGroupsPerUser | 
[**RemoveUserFromUserGroup**](UserGroupApi.md#RemoveUserFromUserGroup) | **Post** /RemoveUserFromUserGroup | 
[**UpdateUserGroup**](UserGroupApi.md#UpdateUserGroup) | **Post** /UpdateUserGroup | 



## AddUserToUserGroup

> AddUserToUserGroupResponse AddUserToUserGroup(ctx).AddUserToUserGroupRequest(addUserToUserGroupRequest).Execute()





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
    addUserToUserGroupRequest := *openapiclient.NewAddUserToUserGroupRequest("UserGroupName_example", "UserName_example") // AddUserToUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupApi.AddUserToUserGroup(context.Background()).AddUserToUserGroupRequest(addUserToUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.AddUserToUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserToUserGroup`: AddUserToUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.AddUserToUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUserToUserGroupRequest** | [**AddUserToUserGroupRequest**](AddUserToUserGroupRequest.md) |  | 

### Return type

[**AddUserToUserGroupResponse**](AddUserToUserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserGroup

> CreateUserGroupResponse CreateUserGroup(ctx).CreateUserGroupRequest(createUserGroupRequest).Execute()





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
    createUserGroupRequest := *openapiclient.NewCreateUserGroupRequest("UserGroupName_example") // CreateUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupApi.CreateUserGroup(context.Background()).CreateUserGroupRequest(createUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.CreateUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserGroup`: CreateUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.CreateUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserGroupRequest** | [**CreateUserGroupRequest**](CreateUserGroupRequest.md) |  | 

### Return type

[**CreateUserGroupResponse**](CreateUserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserGroup

> DeleteUserGroupResponse DeleteUserGroup(ctx).DeleteUserGroupRequest(deleteUserGroupRequest).Execute()





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
    deleteUserGroupRequest := *openapiclient.NewDeleteUserGroupRequest("UserGroupName_example") // DeleteUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupApi.DeleteUserGroup(context.Background()).DeleteUserGroupRequest(deleteUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.DeleteUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserGroup`: DeleteUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.DeleteUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserGroupRequest** | [**DeleteUserGroupRequest**](DeleteUserGroupRequest.md) |  | 

### Return type

[**DeleteUserGroupResponse**](DeleteUserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserGroup

> ReadUserGroupResponse ReadUserGroup(ctx).ReadUserGroupRequest(readUserGroupRequest).Execute()





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
    readUserGroupRequest := *openapiclient.NewReadUserGroupRequest("UserGroupName_example") // ReadUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupApi.ReadUserGroup(context.Background()).ReadUserGroupRequest(readUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.ReadUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserGroup`: ReadUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.ReadUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readUserGroupRequest** | [**ReadUserGroupRequest**](ReadUserGroupRequest.md) |  | 

### Return type

[**ReadUserGroupResponse**](ReadUserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserGroups

> ReadUserGroupsResponse ReadUserGroups(ctx).ReadUserGroupsRequest(readUserGroupsRequest).Execute()





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
    readUserGroupsRequest := *openapiclient.NewReadUserGroupsRequest() // ReadUserGroupsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupApi.ReadUserGroups(context.Background()).ReadUserGroupsRequest(readUserGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.ReadUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserGroups`: ReadUserGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.ReadUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readUserGroupsRequest** | [**ReadUserGroupsRequest**](ReadUserGroupsRequest.md) |  | 

### Return type

[**ReadUserGroupsResponse**](ReadUserGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserGroupsPerUser

> ReadUserGroupsPerUserResponse ReadUserGroupsPerUser(ctx).ReadUserGroupsPerUserRequest(readUserGroupsPerUserRequest).Execute()





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
    readUserGroupsPerUserRequest := *openapiclient.NewReadUserGroupsPerUserRequest("UserName_example") // ReadUserGroupsPerUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupApi.ReadUserGroupsPerUser(context.Background()).ReadUserGroupsPerUserRequest(readUserGroupsPerUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.ReadUserGroupsPerUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserGroupsPerUser`: ReadUserGroupsPerUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.ReadUserGroupsPerUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadUserGroupsPerUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readUserGroupsPerUserRequest** | [**ReadUserGroupsPerUserRequest**](ReadUserGroupsPerUserRequest.md) |  | 

### Return type

[**ReadUserGroupsPerUserResponse**](ReadUserGroupsPerUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromUserGroup

> RemoveUserFromUserGroupResponse RemoveUserFromUserGroup(ctx).RemoveUserFromUserGroupRequest(removeUserFromUserGroupRequest).Execute()





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
    removeUserFromUserGroupRequest := *openapiclient.NewRemoveUserFromUserGroupRequest("UserGroupName_example", "UserName_example") // RemoveUserFromUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupApi.RemoveUserFromUserGroup(context.Background()).RemoveUserFromUserGroupRequest(removeUserFromUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.RemoveUserFromUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveUserFromUserGroup`: RemoveUserFromUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.RemoveUserFromUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeUserFromUserGroupRequest** | [**RemoveUserFromUserGroupRequest**](RemoveUserFromUserGroupRequest.md) |  | 

### Return type

[**RemoveUserFromUserGroupResponse**](RemoveUserFromUserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserGroup

> UpdateUserGroupResponse UpdateUserGroup(ctx).UpdateUserGroupRequest(updateUserGroupRequest).Execute()





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
    updateUserGroupRequest := *openapiclient.NewUpdateUserGroupRequest("UserGroupName_example") // UpdateUserGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupApi.UpdateUserGroup(context.Background()).UpdateUserGroupRequest(updateUserGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupApi.UpdateUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserGroup`: UpdateUserGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupApi.UpdateUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserGroupRequest** | [**UpdateUserGroupRequest**](UpdateUserGroupRequest.md) |  | 

### Return type

[**UpdateUserGroupResponse**](UpdateUserGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

