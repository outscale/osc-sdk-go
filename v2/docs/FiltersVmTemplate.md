# FiltersVmTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCores** | Pointer to **[]int32** | The number of vCores. | [optional] 
**CpuGenerations** | Pointer to **[]string** | The processor generations (for example, &#x60;v4&#x60;). | [optional] 
**CpuPerformances** | Pointer to **[]string** | The performances of the VMs. | [optional] 
**Descriptions** | Pointer to **[]string** | The descriptions of the VM templates. | [optional] 
**ImageIds** | Pointer to **[]string** | The IDs of the OMIs. | [optional] 
**KeypairNames** | Pointer to **[]string** | The names of the keypairs. | [optional] 
**Rams** | Pointer to **[]int32** | The amount of RAM. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the VM templates. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the VM templates. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the VM templates, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**VmTemplateIds** | Pointer to **[]string** | The IDs of the VM templates. | [optional] 
**VmTemplateNames** | Pointer to **[]string** | The names of the VM templates. | [optional] 

## Methods

### NewFiltersVmTemplate

`func NewFiltersVmTemplate() *FiltersVmTemplate`

NewFiltersVmTemplate instantiates a new FiltersVmTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersVmTemplateWithDefaults

`func NewFiltersVmTemplateWithDefaults() *FiltersVmTemplate`

NewFiltersVmTemplateWithDefaults instantiates a new FiltersVmTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCores

`func (o *FiltersVmTemplate) GetCpuCores() []int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *FiltersVmTemplate) GetCpuCoresOk() (*[]int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *FiltersVmTemplate) SetCpuCores(v []int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *FiltersVmTemplate) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCpuGenerations

`func (o *FiltersVmTemplate) GetCpuGenerations() []string`

GetCpuGenerations returns the CpuGenerations field if non-nil, zero value otherwise.

### GetCpuGenerationsOk

`func (o *FiltersVmTemplate) GetCpuGenerationsOk() (*[]string, bool)`

GetCpuGenerationsOk returns a tuple with the CpuGenerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuGenerations

`func (o *FiltersVmTemplate) SetCpuGenerations(v []string)`

SetCpuGenerations sets CpuGenerations field to given value.

### HasCpuGenerations

`func (o *FiltersVmTemplate) HasCpuGenerations() bool`

HasCpuGenerations returns a boolean if a field has been set.

### GetCpuPerformances

`func (o *FiltersVmTemplate) GetCpuPerformances() []string`

GetCpuPerformances returns the CpuPerformances field if non-nil, zero value otherwise.

### GetCpuPerformancesOk

`func (o *FiltersVmTemplate) GetCpuPerformancesOk() (*[]string, bool)`

GetCpuPerformancesOk returns a tuple with the CpuPerformances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPerformances

`func (o *FiltersVmTemplate) SetCpuPerformances(v []string)`

SetCpuPerformances sets CpuPerformances field to given value.

### HasCpuPerformances

`func (o *FiltersVmTemplate) HasCpuPerformances() bool`

HasCpuPerformances returns a boolean if a field has been set.

### GetDescriptions

`func (o *FiltersVmTemplate) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersVmTemplate) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *FiltersVmTemplate) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *FiltersVmTemplate) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetImageIds

`func (o *FiltersVmTemplate) GetImageIds() []string`

GetImageIds returns the ImageIds field if non-nil, zero value otherwise.

### GetImageIdsOk

`func (o *FiltersVmTemplate) GetImageIdsOk() (*[]string, bool)`

GetImageIdsOk returns a tuple with the ImageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIds

`func (o *FiltersVmTemplate) SetImageIds(v []string)`

SetImageIds sets ImageIds field to given value.

### HasImageIds

`func (o *FiltersVmTemplate) HasImageIds() bool`

HasImageIds returns a boolean if a field has been set.

### GetKeypairNames

`func (o *FiltersVmTemplate) GetKeypairNames() []string`

GetKeypairNames returns the KeypairNames field if non-nil, zero value otherwise.

### GetKeypairNamesOk

`func (o *FiltersVmTemplate) GetKeypairNamesOk() (*[]string, bool)`

GetKeypairNamesOk returns a tuple with the KeypairNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairNames

`func (o *FiltersVmTemplate) SetKeypairNames(v []string)`

SetKeypairNames sets KeypairNames field to given value.

### HasKeypairNames

`func (o *FiltersVmTemplate) HasKeypairNames() bool`

HasKeypairNames returns a boolean if a field has been set.

### GetRams

`func (o *FiltersVmTemplate) GetRams() []int32`

GetRams returns the Rams field if non-nil, zero value otherwise.

### GetRamsOk

`func (o *FiltersVmTemplate) GetRamsOk() (*[]int32, bool)`

GetRamsOk returns a tuple with the Rams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRams

`func (o *FiltersVmTemplate) SetRams(v []int32)`

SetRams sets Rams field to given value.

### HasRams

`func (o *FiltersVmTemplate) HasRams() bool`

HasRams returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersVmTemplate) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersVmTemplate) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersVmTemplate) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersVmTemplate) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersVmTemplate) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersVmTemplate) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersVmTemplate) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersVmTemplate) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersVmTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersVmTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersVmTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersVmTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmTemplateIds

`func (o *FiltersVmTemplate) GetVmTemplateIds() []string`

GetVmTemplateIds returns the VmTemplateIds field if non-nil, zero value otherwise.

### GetVmTemplateIdsOk

`func (o *FiltersVmTemplate) GetVmTemplateIdsOk() (*[]string, bool)`

GetVmTemplateIdsOk returns a tuple with the VmTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateIds

`func (o *FiltersVmTemplate) SetVmTemplateIds(v []string)`

SetVmTemplateIds sets VmTemplateIds field to given value.

### HasVmTemplateIds

`func (o *FiltersVmTemplate) HasVmTemplateIds() bool`

HasVmTemplateIds returns a boolean if a field has been set.

### GetVmTemplateNames

`func (o *FiltersVmTemplate) GetVmTemplateNames() []string`

GetVmTemplateNames returns the VmTemplateNames field if non-nil, zero value otherwise.

### GetVmTemplateNamesOk

`func (o *FiltersVmTemplate) GetVmTemplateNamesOk() (*[]string, bool)`

GetVmTemplateNamesOk returns a tuple with the VmTemplateNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateNames

`func (o *FiltersVmTemplate) SetVmTemplateNames(v []string)`

SetVmTemplateNames sets VmTemplateNames field to given value.

### HasVmTemplateNames

`func (o *FiltersVmTemplate) HasVmTemplateNames() bool`

HasVmTemplateNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


