# FiltersSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIds** | Pointer to **[]string** | The account IDs of the owners of the security groups. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets specified when the security groups were created. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups. | [optional] 
**SecurityGroupNames** | Pointer to **[]string** | The names of the security groups. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the security groups. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the security groups. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the security groups, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 

## Methods

### NewFiltersSecurityGroup

`func NewFiltersSecurityGroup() *FiltersSecurityGroup`

NewFiltersSecurityGroup instantiates a new FiltersSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersSecurityGroupWithDefaults

`func NewFiltersSecurityGroupWithDefaults() *FiltersSecurityGroup`

NewFiltersSecurityGroupWithDefaults instantiates a new FiltersSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIds

`func (o *FiltersSecurityGroup) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *FiltersSecurityGroup) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *FiltersSecurityGroup) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *FiltersSecurityGroup) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersSecurityGroup) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersSecurityGroup) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersSecurityGroup) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersSecurityGroup) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *FiltersSecurityGroup) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *FiltersSecurityGroup) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *FiltersSecurityGroup) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *FiltersSecurityGroup) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSecurityGroupNames

`func (o *FiltersSecurityGroup) GetSecurityGroupNames() []string`

GetSecurityGroupNames returns the SecurityGroupNames field if non-nil, zero value otherwise.

### GetSecurityGroupNamesOk

`func (o *FiltersSecurityGroup) GetSecurityGroupNamesOk() (*[]string, bool)`

GetSecurityGroupNamesOk returns a tuple with the SecurityGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupNames

`func (o *FiltersSecurityGroup) SetSecurityGroupNames(v []string)`

SetSecurityGroupNames sets SecurityGroupNames field to given value.

### HasSecurityGroupNames

`func (o *FiltersSecurityGroup) HasSecurityGroupNames() bool`

HasSecurityGroupNames returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersSecurityGroup) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersSecurityGroup) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersSecurityGroup) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersSecurityGroup) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersSecurityGroup) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersSecurityGroup) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersSecurityGroup) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersSecurityGroup) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersSecurityGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersSecurityGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersSecurityGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersSecurityGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


