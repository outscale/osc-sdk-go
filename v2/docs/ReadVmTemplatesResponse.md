# ReadVmTemplatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmTemplates** | Pointer to [**[]VmTemplate**](VmTemplate.md) | Information about one or more VM templates. | [optional] 

## Methods

### NewReadVmTemplatesResponse

`func NewReadVmTemplatesResponse() *ReadVmTemplatesResponse`

NewReadVmTemplatesResponse instantiates a new ReadVmTemplatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmTemplatesResponseWithDefaults

`func NewReadVmTemplatesResponseWithDefaults() *ReadVmTemplatesResponse`

NewReadVmTemplatesResponseWithDefaults instantiates a new ReadVmTemplatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadVmTemplatesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmTemplatesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVmTemplatesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVmTemplatesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmTemplates

`func (o *ReadVmTemplatesResponse) GetVmTemplates() []VmTemplate`

GetVmTemplates returns the VmTemplates field if non-nil, zero value otherwise.

### GetVmTemplatesOk

`func (o *ReadVmTemplatesResponse) GetVmTemplatesOk() (*[]VmTemplate, bool)`

GetVmTemplatesOk returns a tuple with the VmTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplates

`func (o *ReadVmTemplatesResponse) SetVmTemplates(v []VmTemplate)`

SetVmTemplates sets VmTemplates field to given value.

### HasVmTemplates

`func (o *ReadVmTemplatesResponse) HasVmTemplates() bool`

HasVmTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


