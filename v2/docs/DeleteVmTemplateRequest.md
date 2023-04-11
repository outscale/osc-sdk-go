# DeleteVmTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VmTemplateId** | **string** | The ID of the VM template you want to delete.  | 

## Methods

### NewDeleteVmTemplateRequest

`func NewDeleteVmTemplateRequest(vmTemplateId string, ) *DeleteVmTemplateRequest`

NewDeleteVmTemplateRequest instantiates a new DeleteVmTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVmTemplateRequestWithDefaults

`func NewDeleteVmTemplateRequestWithDefaults() *DeleteVmTemplateRequest`

NewDeleteVmTemplateRequestWithDefaults instantiates a new DeleteVmTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteVmTemplateRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteVmTemplateRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteVmTemplateRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteVmTemplateRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVmTemplateId

`func (o *DeleteVmTemplateRequest) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *DeleteVmTemplateRequest) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *DeleteVmTemplateRequest) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


