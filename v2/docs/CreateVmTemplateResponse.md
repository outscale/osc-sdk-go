# CreateVmTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmTemplate** | Pointer to [**VmTemplate**](VmTemplate.md) |  | [optional] 

## Methods

### NewCreateVmTemplateResponse

`func NewCreateVmTemplateResponse() *CreateVmTemplateResponse`

NewCreateVmTemplateResponse instantiates a new CreateVmTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVmTemplateResponseWithDefaults

`func NewCreateVmTemplateResponseWithDefaults() *CreateVmTemplateResponse`

NewCreateVmTemplateResponseWithDefaults instantiates a new CreateVmTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateVmTemplateResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateVmTemplateResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateVmTemplateResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateVmTemplateResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmTemplate

`func (o *CreateVmTemplateResponse) GetVmTemplate() VmTemplate`

GetVmTemplate returns the VmTemplate field if non-nil, zero value otherwise.

### GetVmTemplateOk

`func (o *CreateVmTemplateResponse) GetVmTemplateOk() (*VmTemplate, bool)`

GetVmTemplateOk returns a tuple with the VmTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplate

`func (o *CreateVmTemplateResponse) SetVmTemplate(v VmTemplate)`

SetVmTemplate sets VmTemplate field to given value.

### HasVmTemplate

`func (o *CreateVmTemplateResponse) HasVmTemplate() bool`

HasVmTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


