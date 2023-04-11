# UpdateVmTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A new description for the VM template. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | New tags for your VM template. | [optional] 
**VmTemplateId** | **string** | The ID of the VM template you want to update. | 
**VmTemplateName** | Pointer to **string** | A new name for your VM template. | [optional] 

## Methods

### NewUpdateVmTemplateRequest

`func NewUpdateVmTemplateRequest(vmTemplateId string, ) *UpdateVmTemplateRequest`

NewUpdateVmTemplateRequest instantiates a new UpdateVmTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVmTemplateRequestWithDefaults

`func NewUpdateVmTemplateRequestWithDefaults() *UpdateVmTemplateRequest`

NewUpdateVmTemplateRequestWithDefaults instantiates a new UpdateVmTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateVmTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVmTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVmTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVmTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateVmTemplateRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateVmTemplateRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateVmTemplateRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateVmTemplateRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVmTemplateRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVmTemplateRequest) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVmTemplateRequest) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVmTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmTemplateId

`func (o *UpdateVmTemplateRequest) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *UpdateVmTemplateRequest) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *UpdateVmTemplateRequest) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.


### GetVmTemplateName

`func (o *UpdateVmTemplateRequest) GetVmTemplateName() string`

GetVmTemplateName returns the VmTemplateName field if non-nil, zero value otherwise.

### GetVmTemplateNameOk

`func (o *UpdateVmTemplateRequest) GetVmTemplateNameOk() (*string, bool)`

GetVmTemplateNameOk returns a tuple with the VmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateName

`func (o *UpdateVmTemplateRequest) SetVmTemplateName(v string)`

SetVmTemplateName sets VmTemplateName field to given value.

### HasVmTemplateName

`func (o *UpdateVmTemplateRequest) HasVmTemplateName() bool`

HasVmTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


