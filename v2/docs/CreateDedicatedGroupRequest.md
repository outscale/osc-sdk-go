# CreateDedicatedGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuGeneration** | **int32** | The processor generation for the VMs in the dedicated group (for example, &#x60;4&#x60;). | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Name** | **string** | A name for the dedicated group. | 
**SubregionName** | **string** | The Subregion in which you want to create the dedicated group. | 

## Methods

### NewCreateDedicatedGroupRequest

`func NewCreateDedicatedGroupRequest(cpuGeneration int32, name string, subregionName string, ) *CreateDedicatedGroupRequest`

NewCreateDedicatedGroupRequest instantiates a new CreateDedicatedGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDedicatedGroupRequestWithDefaults

`func NewCreateDedicatedGroupRequestWithDefaults() *CreateDedicatedGroupRequest`

NewCreateDedicatedGroupRequestWithDefaults instantiates a new CreateDedicatedGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuGeneration

`func (o *CreateDedicatedGroupRequest) GetCpuGeneration() int32`

GetCpuGeneration returns the CpuGeneration field if non-nil, zero value otherwise.

### GetCpuGenerationOk

`func (o *CreateDedicatedGroupRequest) GetCpuGenerationOk() (*int32, bool)`

GetCpuGenerationOk returns a tuple with the CpuGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuGeneration

`func (o *CreateDedicatedGroupRequest) SetCpuGeneration(v int32)`

SetCpuGeneration sets CpuGeneration field to given value.


### GetDryRun

`func (o *CreateDedicatedGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateDedicatedGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateDedicatedGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateDedicatedGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetName

`func (o *CreateDedicatedGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDedicatedGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDedicatedGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSubregionName

`func (o *CreateDedicatedGroupRequest) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *CreateDedicatedGroupRequest) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *CreateDedicatedGroupRequest) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


