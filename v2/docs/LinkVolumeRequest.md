# LinkVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | **string** | The name of the device. For a root device, you must use &#x60;/dev/sda1&#x60;. For other volumes, you must use &#x60;/dev/sdX&#x60;, &#x60;/dev/sdXX&#x60;, &#x60;/dev/xvdX&#x60;, or &#x60;/dev/xvdXX&#x60; (where the first &#x60;X&#x60; is a letter between &#x60;b&#x60; and &#x60;z&#x60;, and the second &#x60;X&#x60; is a letter between &#x60;a&#x60; and &#x60;z&#x60;). | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VmId** | **string** | The ID of the VM you want to attach the volume to. | 
**VolumeId** | **string** | The ID of the volume you want to attach. | 

## Methods

### NewLinkVolumeRequest

`func NewLinkVolumeRequest(deviceName string, vmId string, volumeId string, ) *LinkVolumeRequest`

NewLinkVolumeRequest instantiates a new LinkVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkVolumeRequestWithDefaults

`func NewLinkVolumeRequestWithDefaults() *LinkVolumeRequest`

NewLinkVolumeRequestWithDefaults instantiates a new LinkVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *LinkVolumeRequest) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *LinkVolumeRequest) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *LinkVolumeRequest) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetDryRun

`func (o *LinkVolumeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkVolumeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkVolumeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkVolumeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVmId

`func (o *LinkVolumeRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *LinkVolumeRequest) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *LinkVolumeRequest) SetVmId(v string)`

SetVmId sets VmId field to given value.


### GetVolumeId

`func (o *LinkVolumeRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *LinkVolumeRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *LinkVolumeRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


