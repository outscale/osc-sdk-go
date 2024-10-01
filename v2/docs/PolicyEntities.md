# PolicyEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]MinimalPolicy**](MinimalPolicy.md) |  | [optional] 
**Groups** | Pointer to [**[]MinimalPolicy**](MinimalPolicy.md) |  | [optional] 
**HasMoreItems** | Pointer to **bool** | If true, there are more items to return using the &#x60;FirstItem&#x60; parameter in a new request. | [optional] 
**ItemsCount** | Pointer to **int32** | The number of entities the specified policy is linked to. | [optional] 
**MaxResultsLimit** | Pointer to **int32** | Indicates maximum results defined for the operation. | [optional] 
**MaxResultsTruncated** | Pointer to **bool** | If true, indicates whether requested page size is more than allowed. | [optional] 
**Users** | Pointer to [**[]MinimalPolicy**](MinimalPolicy.md) |  | [optional] 

## Methods

### NewPolicyEntities

`func NewPolicyEntities() *PolicyEntities`

NewPolicyEntities instantiates a new PolicyEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEntitiesWithDefaults

`func NewPolicyEntitiesWithDefaults() *PolicyEntities`

NewPolicyEntitiesWithDefaults instantiates a new PolicyEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *PolicyEntities) GetAccounts() []MinimalPolicy`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PolicyEntities) GetAccountsOk() (*[]MinimalPolicy, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PolicyEntities) SetAccounts(v []MinimalPolicy)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PolicyEntities) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetGroups

`func (o *PolicyEntities) GetGroups() []MinimalPolicy`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PolicyEntities) GetGroupsOk() (*[]MinimalPolicy, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PolicyEntities) SetGroups(v []MinimalPolicy)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PolicyEntities) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHasMoreItems

`func (o *PolicyEntities) GetHasMoreItems() bool`

GetHasMoreItems returns the HasMoreItems field if non-nil, zero value otherwise.

### GetHasMoreItemsOk

`func (o *PolicyEntities) GetHasMoreItemsOk() (*bool, bool)`

GetHasMoreItemsOk returns a tuple with the HasMoreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreItems

`func (o *PolicyEntities) SetHasMoreItems(v bool)`

SetHasMoreItems sets HasMoreItems field to given value.

### HasHasMoreItems

`func (o *PolicyEntities) HasHasMoreItems() bool`

HasHasMoreItems returns a boolean if a field has been set.

### GetItemsCount

`func (o *PolicyEntities) GetItemsCount() int32`

GetItemsCount returns the ItemsCount field if non-nil, zero value otherwise.

### GetItemsCountOk

`func (o *PolicyEntities) GetItemsCountOk() (*int32, bool)`

GetItemsCountOk returns a tuple with the ItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsCount

`func (o *PolicyEntities) SetItemsCount(v int32)`

SetItemsCount sets ItemsCount field to given value.

### HasItemsCount

`func (o *PolicyEntities) HasItemsCount() bool`

HasItemsCount returns a boolean if a field has been set.

### GetMaxResultsLimit

`func (o *PolicyEntities) GetMaxResultsLimit() int32`

GetMaxResultsLimit returns the MaxResultsLimit field if non-nil, zero value otherwise.

### GetMaxResultsLimitOk

`func (o *PolicyEntities) GetMaxResultsLimitOk() (*int32, bool)`

GetMaxResultsLimitOk returns a tuple with the MaxResultsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsLimit

`func (o *PolicyEntities) SetMaxResultsLimit(v int32)`

SetMaxResultsLimit sets MaxResultsLimit field to given value.

### HasMaxResultsLimit

`func (o *PolicyEntities) HasMaxResultsLimit() bool`

HasMaxResultsLimit returns a boolean if a field has been set.

### GetMaxResultsTruncated

`func (o *PolicyEntities) GetMaxResultsTruncated() bool`

GetMaxResultsTruncated returns the MaxResultsTruncated field if non-nil, zero value otherwise.

### GetMaxResultsTruncatedOk

`func (o *PolicyEntities) GetMaxResultsTruncatedOk() (*bool, bool)`

GetMaxResultsTruncatedOk returns a tuple with the MaxResultsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsTruncated

`func (o *PolicyEntities) SetMaxResultsTruncated(v bool)`

SetMaxResultsTruncated sets MaxResultsTruncated field to given value.

### HasMaxResultsTruncated

`func (o *PolicyEntities) HasMaxResultsTruncated() bool`

HasMaxResultsTruncated returns a boolean if a field has been set.

### GetUsers

`func (o *PolicyEntities) GetUsers() []MinimalPolicy`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PolicyEntities) GetUsersOk() (*[]MinimalPolicy, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PolicyEntities) SetUsers(v []MinimalPolicy)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PolicyEntities) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


