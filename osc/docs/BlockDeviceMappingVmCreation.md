# BlockDeviceMappingVmCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bsu** | Pointer to [**BsuToCreate**](BsuToCreate.md) |  | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. | [optional] 
**NoDevice** | Pointer to **string** | Removes the device which is included in the block device mapping of the OMI. | [optional] 
**VirtualDeviceName** | Pointer to **string** | The name of the virtual device (ephemeralN). | [optional] 

## Methods

### NewBlockDeviceMappingVmCreation

`func NewBlockDeviceMappingVmCreation() *BlockDeviceMappingVmCreation`

NewBlockDeviceMappingVmCreation instantiates a new BlockDeviceMappingVmCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockDeviceMappingVmCreationWithDefaults

`func NewBlockDeviceMappingVmCreationWithDefaults() *BlockDeviceMappingVmCreation`

NewBlockDeviceMappingVmCreationWithDefaults instantiates a new BlockDeviceMappingVmCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsu

`func (o *BlockDeviceMappingVmCreation) GetBsu() BsuToCreate`

GetBsu returns the Bsu field if non-nil, zero value otherwise.

### GetBsuOk

`func (o *BlockDeviceMappingVmCreation) GetBsuOk() (*BsuToCreate, bool)`

GetBsuOk returns a tuple with the Bsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsu

`func (o *BlockDeviceMappingVmCreation) SetBsu(v BsuToCreate)`

SetBsu sets Bsu field to given value.

### HasBsu

`func (o *BlockDeviceMappingVmCreation) HasBsu() bool`

HasBsu returns a boolean if a field has been set.

### GetDeviceName

`func (o *BlockDeviceMappingVmCreation) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *BlockDeviceMappingVmCreation) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *BlockDeviceMappingVmCreation) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *BlockDeviceMappingVmCreation) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetNoDevice

`func (o *BlockDeviceMappingVmCreation) GetNoDevice() string`

GetNoDevice returns the NoDevice field if non-nil, zero value otherwise.

### GetNoDeviceOk

`func (o *BlockDeviceMappingVmCreation) GetNoDeviceOk() (*string, bool)`

GetNoDeviceOk returns a tuple with the NoDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDevice

`func (o *BlockDeviceMappingVmCreation) SetNoDevice(v string)`

SetNoDevice sets NoDevice field to given value.

### HasNoDevice

`func (o *BlockDeviceMappingVmCreation) HasNoDevice() bool`

HasNoDevice returns a boolean if a field has been set.

### GetVirtualDeviceName

`func (o *BlockDeviceMappingVmCreation) GetVirtualDeviceName() string`

GetVirtualDeviceName returns the VirtualDeviceName field if non-nil, zero value otherwise.

### GetVirtualDeviceNameOk

`func (o *BlockDeviceMappingVmCreation) GetVirtualDeviceNameOk() (*string, bool)`

GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceName

`func (o *BlockDeviceMappingVmCreation) SetVirtualDeviceName(v string)`

SetVirtualDeviceName sets VirtualDeviceName field to given value.

### HasVirtualDeviceName

`func (o *BlockDeviceMappingVmCreation) HasVirtualDeviceName() bool`

HasVirtualDeviceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


