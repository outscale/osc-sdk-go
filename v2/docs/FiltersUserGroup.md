# FiltersUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathPrefix** | Pointer to **string** | The path prefix of the groups. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 
**UserGroupIds** | Pointer to **[]string** | The IDs of the user groups. | [optional] 

## Methods

### NewFiltersUserGroup

`func NewFiltersUserGroup() *FiltersUserGroup`

NewFiltersUserGroup instantiates a new FiltersUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersUserGroupWithDefaults

`func NewFiltersUserGroupWithDefaults() *FiltersUserGroup`

NewFiltersUserGroupWithDefaults instantiates a new FiltersUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPathPrefix

`func (o *FiltersUserGroup) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *FiltersUserGroup) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *FiltersUserGroup) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *FiltersUserGroup) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetUserGroupIds

`func (o *FiltersUserGroup) GetUserGroupIds() []string`

GetUserGroupIds returns the UserGroupIds field if non-nil, zero value otherwise.

### GetUserGroupIdsOk

`func (o *FiltersUserGroup) GetUserGroupIdsOk() (*[]string, bool)`

GetUserGroupIdsOk returns a tuple with the UserGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupIds

`func (o *FiltersUserGroup) SetUserGroupIds(v []string)`

SetUserGroupIds sets UserGroupIds field to given value.

### HasUserGroupIds

`func (o *FiltersUserGroup) HasUserGroupIds() bool`

HasUserGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


