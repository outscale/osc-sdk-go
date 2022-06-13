# LinkNicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceNumber** | **int32** | The index of the VM device for the NIC attachment (between &#x60;1&#x60; and &#x60;7&#x60;, both included). | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | **string** | The ID of the NIC you want to attach. | 
**VmId** | **string** | The ID of the VM to which you want to attach the NIC. | 

## Methods

### NewLinkNicRequest

`func NewLinkNicRequest(deviceNumber int32, nicId string, vmId string, ) *LinkNicRequest`

NewLinkNicRequest instantiates a new LinkNicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkNicRequestWithDefaults

`func NewLinkNicRequestWithDefaults() *LinkNicRequest`

NewLinkNicRequestWithDefaults instantiates a new LinkNicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceNumber

`func (o *LinkNicRequest) GetDeviceNumber() int32`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *LinkNicRequest) GetDeviceNumberOk() (*int32, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNumber

`func (o *LinkNicRequest) SetDeviceNumber(v int32)`

SetDeviceNumber sets DeviceNumber field to given value.


### GetDryRun

`func (o *LinkNicRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkNicRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkNicRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkNicRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNicId

`func (o *LinkNicRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *LinkNicRequest) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *LinkNicRequest) SetNicId(v string)`

SetNicId sets NicId field to given value.


### GetVmId

`func (o *LinkNicRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *LinkNicRequest) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *LinkNicRequest) SetVmId(v string)`

SetVmId sets VmId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


