# BsuCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM. | [optional] 
**LinkDate** | Pointer to **time.Time** | The date and time of attachment of the volume to the VM, in ISO 8601 date-time format. | [optional] 
**State** | Pointer to **string** | The state of the volume. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume. | [optional] 

## Methods

### NewBsuCreated

`func NewBsuCreated() *BsuCreated`

NewBsuCreated instantiates a new BsuCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBsuCreatedWithDefaults

`func NewBsuCreatedWithDefaults() *BsuCreated`

NewBsuCreatedWithDefaults instantiates a new BsuCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnVmDeletion

`func (o *BsuCreated) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *BsuCreated) GetDeleteOnVmDeletionOk() (*bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnVmDeletion

`func (o *BsuCreated) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion sets DeleteOnVmDeletion field to given value.

### HasDeleteOnVmDeletion

`func (o *BsuCreated) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### GetLinkDate

`func (o *BsuCreated) GetLinkDate() time.Time`

GetLinkDate returns the LinkDate field if non-nil, zero value otherwise.

### GetLinkDateOk

`func (o *BsuCreated) GetLinkDateOk() (*time.Time, bool)`

GetLinkDateOk returns a tuple with the LinkDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDate

`func (o *BsuCreated) SetLinkDate(v time.Time)`

SetLinkDate sets LinkDate field to given value.

### HasLinkDate

`func (o *BsuCreated) HasLinkDate() bool`

HasLinkDate returns a boolean if a field has been set.

### GetState

`func (o *BsuCreated) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BsuCreated) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BsuCreated) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BsuCreated) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVolumeId

`func (o *BsuCreated) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *BsuCreated) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *BsuCreated) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *BsuCreated) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


