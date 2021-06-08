# BsuToUpdateVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If set to true, the volume is deleted when terminating the VM. If set to false, the volume is not deleted when terminating the VM. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume. | [optional] 

## Methods

### NewBsuToUpdateVm

`func NewBsuToUpdateVm() *BsuToUpdateVm`

NewBsuToUpdateVm instantiates a new BsuToUpdateVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBsuToUpdateVmWithDefaults

`func NewBsuToUpdateVmWithDefaults() *BsuToUpdateVm`

NewBsuToUpdateVmWithDefaults instantiates a new BsuToUpdateVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnVmDeletion

`func (o *BsuToUpdateVm) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *BsuToUpdateVm) GetDeleteOnVmDeletionOk() (*bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnVmDeletion

`func (o *BsuToUpdateVm) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion sets DeleteOnVmDeletion field to given value.

### HasDeleteOnVmDeletion

`func (o *BsuToUpdateVm) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### GetVolumeId

`func (o *BsuToUpdateVm) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *BsuToUpdateVm) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *BsuToUpdateVm) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *BsuToUpdateVm) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


