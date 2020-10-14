# LinkFlexibleGpuRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**FlexibleGpuId** | **string** | The ID of the fGPU you want to attach. | 
**VmId** | **string** | The ID of the VM you want to attach the fGPU to. | 

## Methods

### NewLinkFlexibleGpuRequest

`func NewLinkFlexibleGpuRequest(flexibleGpuId string, vmId string, ) *LinkFlexibleGpuRequest`

NewLinkFlexibleGpuRequest instantiates a new LinkFlexibleGpuRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkFlexibleGpuRequestWithDefaults

`func NewLinkFlexibleGpuRequestWithDefaults() *LinkFlexibleGpuRequest`

NewLinkFlexibleGpuRequestWithDefaults instantiates a new LinkFlexibleGpuRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *LinkFlexibleGpuRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkFlexibleGpuRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkFlexibleGpuRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkFlexibleGpuRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFlexibleGpuId

`func (o *LinkFlexibleGpuRequest) GetFlexibleGpuId() string`

GetFlexibleGpuId returns the FlexibleGpuId field if non-nil, zero value otherwise.

### GetFlexibleGpuIdOk

`func (o *LinkFlexibleGpuRequest) GetFlexibleGpuIdOk() (*string, bool)`

GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleGpuId

`func (o *LinkFlexibleGpuRequest) SetFlexibleGpuId(v string)`

SetFlexibleGpuId sets FlexibleGpuId field to given value.


### GetVmId

`func (o *LinkFlexibleGpuRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *LinkFlexibleGpuRequest) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *LinkFlexibleGpuRequest) SetVmId(v string)`

SetVmId sets VmId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


