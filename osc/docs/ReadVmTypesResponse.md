# ReadVmTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmTypes** | Pointer to [**[]VmType**](VmType.md) | Information about one or more VM types. | [optional] 

## Methods

### NewReadVmTypesResponse

`func NewReadVmTypesResponse() *ReadVmTypesResponse`

NewReadVmTypesResponse instantiates a new ReadVmTypesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmTypesResponseWithDefaults

`func NewReadVmTypesResponseWithDefaults() *ReadVmTypesResponse`

NewReadVmTypesResponseWithDefaults instantiates a new ReadVmTypesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadVmTypesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmTypesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVmTypesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVmTypesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmTypes

`func (o *ReadVmTypesResponse) GetVmTypes() []VmType`

GetVmTypes returns the VmTypes field if non-nil, zero value otherwise.

### GetVmTypesOk

`func (o *ReadVmTypesResponse) GetVmTypesOk() (*[]VmType, bool)`

GetVmTypesOk returns a tuple with the VmTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTypes

`func (o *ReadVmTypesResponse) SetVmTypes(v []VmType)`

SetVmTypes sets VmTypes field to given value.

### HasVmTypes

`func (o *ReadVmTypesResponse) HasVmTypes() bool`

HasVmTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


