# MasterKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **string** | The date and time of creation of the master key. | [optional] 
**DeletionDate** | Pointer to **string** | The date and time of scheduled deletion of the master key. | [optional] 
**Description** | Pointer to **string** |  The description of the master key. | [optional] 
**MasterKeyId** | Pointer to **string** | The ID of the master key. | [optional] 
**State** | Pointer to **string** | The state of the master key (&#x60;Enabled&#x60; \\| &#x60;Disabled&#x60; \\| &#x60;PendingDeletion&#x60;). | [optional] 

## Methods

### NewMasterKey

`func NewMasterKey() *MasterKey`

NewMasterKey instantiates a new MasterKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterKeyWithDefaults

`func NewMasterKeyWithDefaults() *MasterKey`

NewMasterKeyWithDefaults instantiates a new MasterKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *MasterKey) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *MasterKey) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *MasterKey) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *MasterKey) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDeletionDate

`func (o *MasterKey) GetDeletionDate() string`

GetDeletionDate returns the DeletionDate field if non-nil, zero value otherwise.

### GetDeletionDateOk

`func (o *MasterKey) GetDeletionDateOk() (*string, bool)`

GetDeletionDateOk returns a tuple with the DeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDate

`func (o *MasterKey) SetDeletionDate(v string)`

SetDeletionDate sets DeletionDate field to given value.

### HasDeletionDate

`func (o *MasterKey) HasDeletionDate() bool`

HasDeletionDate returns a boolean if a field has been set.

### GetDescription

`func (o *MasterKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MasterKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MasterKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MasterKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMasterKeyId

`func (o *MasterKey) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *MasterKey) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *MasterKey) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.

### HasMasterKeyId

`func (o *MasterKey) HasMasterKeyId() bool`

HasMasterKeyId returns a boolean if a field has been set.

### GetState

`func (o *MasterKey) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MasterKey) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MasterKey) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MasterKey) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


