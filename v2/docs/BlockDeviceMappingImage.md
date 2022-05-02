# BlockDeviceMappingImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bsu** | Pointer to [**BsuToCreate**](BsuToCreate.md) |  | [optional] 
**DeviceName** | Pointer to **string** | The device name for the volume. For a root device, you must use &#x60;/dev/sda1&#x60;. For other volumes, you must use &#x60;/dev/sdX&#x60; or &#x60;/dev/xvdX&#x60; (where &#x60;X&#x60; is a letter between &#x60;b&#x60; and &#x60;z&#x60;). | [optional] 
**VirtualDeviceName** | Pointer to **string** | The name of the virtual device (ephemeralN). | [optional] 

## Methods

### NewBlockDeviceMappingImage

`func NewBlockDeviceMappingImage() *BlockDeviceMappingImage`

NewBlockDeviceMappingImage instantiates a new BlockDeviceMappingImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockDeviceMappingImageWithDefaults

`func NewBlockDeviceMappingImageWithDefaults() *BlockDeviceMappingImage`

NewBlockDeviceMappingImageWithDefaults instantiates a new BlockDeviceMappingImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsu

`func (o *BlockDeviceMappingImage) GetBsu() BsuToCreate`

GetBsu returns the Bsu field if non-nil, zero value otherwise.

### GetBsuOk

`func (o *BlockDeviceMappingImage) GetBsuOk() (*BsuToCreate, bool)`

GetBsuOk returns a tuple with the Bsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsu

`func (o *BlockDeviceMappingImage) SetBsu(v BsuToCreate)`

SetBsu sets Bsu field to given value.

### HasBsu

`func (o *BlockDeviceMappingImage) HasBsu() bool`

HasBsu returns a boolean if a field has been set.

### GetDeviceName

`func (o *BlockDeviceMappingImage) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *BlockDeviceMappingImage) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *BlockDeviceMappingImage) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *BlockDeviceMappingImage) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetVirtualDeviceName

`func (o *BlockDeviceMappingImage) GetVirtualDeviceName() string`

GetVirtualDeviceName returns the VirtualDeviceName field if non-nil, zero value otherwise.

### GetVirtualDeviceNameOk

`func (o *BlockDeviceMappingImage) GetVirtualDeviceNameOk() (*string, bool)`

GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceName

`func (o *BlockDeviceMappingImage) SetVirtualDeviceName(v string)`

SetVirtualDeviceName sets VirtualDeviceName field to given value.

### HasVirtualDeviceName

`func (o *BlockDeviceMappingImage) HasVirtualDeviceName() bool`

HasVirtualDeviceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


