# BlockDeviceMappingVmUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bsu** | Pointer to [**BsuToUpdateVm**](BsuToUpdateVm.md) |  | [optional] 
**DeviceName** | Pointer to **string** | The device name for the volume. For a root device, you must use &#x60;/dev/sda1&#x60;. For other volumes, you must use &#x60;/dev/sdX&#x60; or &#x60;/dev/xvdX&#x60; (where &#x60;X&#x60; is a letter between &#x60;b&#x60; and &#x60;z&#x60;). | [optional] 
**NoDevice** | Pointer to **string** | Removes the device which is included in the block device mapping of the OMI. | [optional] 
**VirtualDeviceName** | Pointer to **string** | The name of the virtual device (&#x60;ephemeralN&#x60;). | [optional] 

## Methods

### NewBlockDeviceMappingVmUpdate

`func NewBlockDeviceMappingVmUpdate() *BlockDeviceMappingVmUpdate`

NewBlockDeviceMappingVmUpdate instantiates a new BlockDeviceMappingVmUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockDeviceMappingVmUpdateWithDefaults

`func NewBlockDeviceMappingVmUpdateWithDefaults() *BlockDeviceMappingVmUpdate`

NewBlockDeviceMappingVmUpdateWithDefaults instantiates a new BlockDeviceMappingVmUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsu

`func (o *BlockDeviceMappingVmUpdate) GetBsu() BsuToUpdateVm`

GetBsu returns the Bsu field if non-nil, zero value otherwise.

### GetBsuOk

`func (o *BlockDeviceMappingVmUpdate) GetBsuOk() (*BsuToUpdateVm, bool)`

GetBsuOk returns a tuple with the Bsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsu

`func (o *BlockDeviceMappingVmUpdate) SetBsu(v BsuToUpdateVm)`

SetBsu sets Bsu field to given value.

### HasBsu

`func (o *BlockDeviceMappingVmUpdate) HasBsu() bool`

HasBsu returns a boolean if a field has been set.

### GetDeviceName

`func (o *BlockDeviceMappingVmUpdate) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *BlockDeviceMappingVmUpdate) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *BlockDeviceMappingVmUpdate) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *BlockDeviceMappingVmUpdate) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetNoDevice

`func (o *BlockDeviceMappingVmUpdate) GetNoDevice() string`

GetNoDevice returns the NoDevice field if non-nil, zero value otherwise.

### GetNoDeviceOk

`func (o *BlockDeviceMappingVmUpdate) GetNoDeviceOk() (*string, bool)`

GetNoDeviceOk returns a tuple with the NoDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDevice

`func (o *BlockDeviceMappingVmUpdate) SetNoDevice(v string)`

SetNoDevice sets NoDevice field to given value.

### HasNoDevice

`func (o *BlockDeviceMappingVmUpdate) HasNoDevice() bool`

HasNoDevice returns a boolean if a field has been set.

### GetVirtualDeviceName

`func (o *BlockDeviceMappingVmUpdate) GetVirtualDeviceName() string`

GetVirtualDeviceName returns the VirtualDeviceName field if non-nil, zero value otherwise.

### GetVirtualDeviceNameOk

`func (o *BlockDeviceMappingVmUpdate) GetVirtualDeviceNameOk() (*string, bool)`

GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceName

`func (o *BlockDeviceMappingVmUpdate) SetVirtualDeviceName(v string)`

SetVirtualDeviceName sets VirtualDeviceName field to given value.

### HasVirtualDeviceName

`func (o *BlockDeviceMappingVmUpdate) HasVirtualDeviceName() bool`

HasVirtualDeviceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


