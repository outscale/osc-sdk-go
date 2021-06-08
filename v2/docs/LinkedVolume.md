# LinkedVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM. | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. | [optional] 
**State** | Pointer to **string** | The state of the attachment of the volume (&#x60;attaching&#x60; \\| &#x60;detaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detached&#x60;). | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume. | [optional] 

## Methods

### NewLinkedVolume

`func NewLinkedVolume() *LinkedVolume`

NewLinkedVolume instantiates a new LinkedVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedVolumeWithDefaults

`func NewLinkedVolumeWithDefaults() *LinkedVolume`

NewLinkedVolumeWithDefaults instantiates a new LinkedVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnVmDeletion

`func (o *LinkedVolume) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *LinkedVolume) GetDeleteOnVmDeletionOk() (*bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnVmDeletion

`func (o *LinkedVolume) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion sets DeleteOnVmDeletion field to given value.

### HasDeleteOnVmDeletion

`func (o *LinkedVolume) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### GetDeviceName

`func (o *LinkedVolume) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *LinkedVolume) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *LinkedVolume) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *LinkedVolume) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetState

`func (o *LinkedVolume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LinkedVolume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LinkedVolume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LinkedVolume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVmId

`func (o *LinkedVolume) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *LinkedVolume) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *LinkedVolume) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *LinkedVolume) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetVolumeId

`func (o *LinkedVolume) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *LinkedVolume) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *LinkedVolume) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *LinkedVolume) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


