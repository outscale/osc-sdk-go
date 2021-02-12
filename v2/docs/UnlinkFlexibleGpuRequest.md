# UnlinkFlexibleGpuRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**FlexibleGpuId** | **string** | The ID of the fGPU you want to detach from your VM. | 

## Methods

### NewUnlinkFlexibleGpuRequest

`func NewUnlinkFlexibleGpuRequest(flexibleGpuId string, ) *UnlinkFlexibleGpuRequest`

NewUnlinkFlexibleGpuRequest instantiates a new UnlinkFlexibleGpuRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkFlexibleGpuRequestWithDefaults

`func NewUnlinkFlexibleGpuRequestWithDefaults() *UnlinkFlexibleGpuRequest`

NewUnlinkFlexibleGpuRequestWithDefaults instantiates a new UnlinkFlexibleGpuRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UnlinkFlexibleGpuRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkFlexibleGpuRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkFlexibleGpuRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkFlexibleGpuRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFlexibleGpuId

`func (o *UnlinkFlexibleGpuRequest) GetFlexibleGpuId() string`

GetFlexibleGpuId returns the FlexibleGpuId field if non-nil, zero value otherwise.

### GetFlexibleGpuIdOk

`func (o *UnlinkFlexibleGpuRequest) GetFlexibleGpuIdOk() (*string, bool)`

GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleGpuId

`func (o *UnlinkFlexibleGpuRequest) SetFlexibleGpuId(v string)`

SetFlexibleGpuId sets FlexibleGpuId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


