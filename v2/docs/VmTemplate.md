# VmTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCores** | **int32** | The number of vCores. | 
**CpuGeneration** | **string** | The processor generation. | 
**CpuPerformance** | Pointer to **string** | The performance of the VMs. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time of creation of the VM template. | [optional] 
**Description** | Pointer to **string** | The description of the VM template. | [optional] 
**ImageId** | **string** | The ID of the OMI. | 
**KeypairName** | Pointer to **string** | The name of the keypair. | [optional] 
**Ram** | **int32** | The amount of RAM. | 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the VM template. | [optional] 
**VmTemplateId** | **string** | The ID of the VM template. | 
**VmTemplateName** | **string** | The name of the VM template. | 

## Methods

### NewVmTemplate

`func NewVmTemplate(cpuCores int32, cpuGeneration string, imageId string, ram int32, vmTemplateId string, vmTemplateName string, ) *VmTemplate`

NewVmTemplate instantiates a new VmTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmTemplateWithDefaults

`func NewVmTemplateWithDefaults() *VmTemplate`

NewVmTemplateWithDefaults instantiates a new VmTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCores

`func (o *VmTemplate) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *VmTemplate) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *VmTemplate) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.


### GetCpuGeneration

`func (o *VmTemplate) GetCpuGeneration() string`

GetCpuGeneration returns the CpuGeneration field if non-nil, zero value otherwise.

### GetCpuGenerationOk

`func (o *VmTemplate) GetCpuGenerationOk() (*string, bool)`

GetCpuGenerationOk returns a tuple with the CpuGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuGeneration

`func (o *VmTemplate) SetCpuGeneration(v string)`

SetCpuGeneration sets CpuGeneration field to given value.


### GetCpuPerformance

`func (o *VmTemplate) GetCpuPerformance() string`

GetCpuPerformance returns the CpuPerformance field if non-nil, zero value otherwise.

### GetCpuPerformanceOk

`func (o *VmTemplate) GetCpuPerformanceOk() (*string, bool)`

GetCpuPerformanceOk returns a tuple with the CpuPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPerformance

`func (o *VmTemplate) SetCpuPerformance(v string)`

SetCpuPerformance sets CpuPerformance field to given value.

### HasCpuPerformance

`func (o *VmTemplate) HasCpuPerformance() bool`

HasCpuPerformance returns a boolean if a field has been set.

### GetCreationDate

`func (o *VmTemplate) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *VmTemplate) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *VmTemplate) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *VmTemplate) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *VmTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VmTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VmTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VmTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageId

`func (o *VmTemplate) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *VmTemplate) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *VmTemplate) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetKeypairName

`func (o *VmTemplate) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *VmTemplate) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *VmTemplate) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *VmTemplate) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetRam

`func (o *VmTemplate) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *VmTemplate) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *VmTemplate) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetTags

`func (o *VmTemplate) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VmTemplate) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VmTemplate) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VmTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmTemplateId

`func (o *VmTemplate) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *VmTemplate) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *VmTemplate) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.


### GetVmTemplateName

`func (o *VmTemplate) GetVmTemplateName() string`

GetVmTemplateName returns the VmTemplateName field if non-nil, zero value otherwise.

### GetVmTemplateNameOk

`func (o *VmTemplate) GetVmTemplateNameOk() (*string, bool)`

GetVmTemplateNameOk returns a tuple with the VmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateName

`func (o *VmTemplate) SetVmTemplateName(v string)`

SetVmTemplateName sets VmTemplateName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


