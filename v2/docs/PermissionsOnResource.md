# PermissionsOnResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIds** | Pointer to **[]string** | One or more account IDs that the permission is associated with. | [optional] 
**GlobalPermission** | Pointer to **bool** | A global permission for all accounts.&lt;br /&gt; (Request) Set this parameter to true to make the resource public (if the parent parameter is &#x60;Additions&#x60;) or to make the resource private (if the parent parameter is &#x60;Removals&#x60;).&lt;br /&gt; (Response) If true, the resource is public. If false, the resource is private. | [optional] 

## Methods

### NewPermissionsOnResource

`func NewPermissionsOnResource() *PermissionsOnResource`

NewPermissionsOnResource instantiates a new PermissionsOnResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsOnResourceWithDefaults

`func NewPermissionsOnResourceWithDefaults() *PermissionsOnResource`

NewPermissionsOnResourceWithDefaults instantiates a new PermissionsOnResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIds

`func (o *PermissionsOnResource) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *PermissionsOnResource) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *PermissionsOnResource) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *PermissionsOnResource) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetGlobalPermission

`func (o *PermissionsOnResource) GetGlobalPermission() bool`

GetGlobalPermission returns the GlobalPermission field if non-nil, zero value otherwise.

### GetGlobalPermissionOk

`func (o *PermissionsOnResource) GetGlobalPermissionOk() (*bool, bool)`

GetGlobalPermissionOk returns a tuple with the GlobalPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermission

`func (o *PermissionsOnResource) SetGlobalPermission(v bool)`

SetGlobalPermission sets GlobalPermission field to given value.

### HasGlobalPermission

`func (o *PermissionsOnResource) HasGlobalPermission() bool`

HasGlobalPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


