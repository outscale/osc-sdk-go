# FiltersMasterKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptions** | Pointer to **[]string** | The descriptions of the master keys. | [optional] 
**MasterKeyIds** | Pointer to **[]string** | The IDs of the master keys. | [optional] 
**States** | Pointer to **[]string** | The states of the master keys (&#x60;Enabled&#x60; \\| &#x60;Disabled&#x60; \\| &#x60;PendingDeletion&#x60;). | [optional] 

## Methods

### NewFiltersMasterKey

`func NewFiltersMasterKey() *FiltersMasterKey`

NewFiltersMasterKey instantiates a new FiltersMasterKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersMasterKeyWithDefaults

`func NewFiltersMasterKeyWithDefaults() *FiltersMasterKey`

NewFiltersMasterKeyWithDefaults instantiates a new FiltersMasterKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptions

`func (o *FiltersMasterKey) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersMasterKey) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *FiltersMasterKey) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *FiltersMasterKey) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetMasterKeyIds

`func (o *FiltersMasterKey) GetMasterKeyIds() []string`

GetMasterKeyIds returns the MasterKeyIds field if non-nil, zero value otherwise.

### GetMasterKeyIdsOk

`func (o *FiltersMasterKey) GetMasterKeyIdsOk() (*[]string, bool)`

GetMasterKeyIdsOk returns a tuple with the MasterKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyIds

`func (o *FiltersMasterKey) SetMasterKeyIds(v []string)`

SetMasterKeyIds sets MasterKeyIds field to given value.

### HasMasterKeyIds

`func (o *FiltersMasterKey) HasMasterKeyIds() bool`

HasMasterKeyIds returns a boolean if a field has been set.

### GetStates

`func (o *FiltersMasterKey) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersMasterKey) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersMasterKey) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersMasterKey) HasStates() bool`

HasStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


