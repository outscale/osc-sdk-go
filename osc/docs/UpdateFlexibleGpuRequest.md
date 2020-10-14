# UpdateFlexibleGpuRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If &#x60;true&#x60;, the fGPU is deleted when the VM is terminated. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**FlexibleGpuId** | **string** | The ID of the fGPU you want to modify. | 

## Methods

### NewUpdateFlexibleGpuRequest

`func NewUpdateFlexibleGpuRequest(flexibleGpuId string, ) *UpdateFlexibleGpuRequest`

NewUpdateFlexibleGpuRequest instantiates a new UpdateFlexibleGpuRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFlexibleGpuRequestWithDefaults

`func NewUpdateFlexibleGpuRequestWithDefaults() *UpdateFlexibleGpuRequest`

NewUpdateFlexibleGpuRequestWithDefaults instantiates a new UpdateFlexibleGpuRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnVmDeletion

`func (o *UpdateFlexibleGpuRequest) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *UpdateFlexibleGpuRequest) GetDeleteOnVmDeletionOk() (*bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnVmDeletion

`func (o *UpdateFlexibleGpuRequest) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion sets DeleteOnVmDeletion field to given value.

### HasDeleteOnVmDeletion

`func (o *UpdateFlexibleGpuRequest) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateFlexibleGpuRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateFlexibleGpuRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateFlexibleGpuRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateFlexibleGpuRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFlexibleGpuId

`func (o *UpdateFlexibleGpuRequest) GetFlexibleGpuId() string`

GetFlexibleGpuId returns the FlexibleGpuId field if non-nil, zero value otherwise.

### GetFlexibleGpuIdOk

`func (o *UpdateFlexibleGpuRequest) GetFlexibleGpuIdOk() (*string, bool)`

GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleGpuId

`func (o *UpdateFlexibleGpuRequest) SetFlexibleGpuId(v string)`

SetFlexibleGpuId sets FlexibleGpuId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


