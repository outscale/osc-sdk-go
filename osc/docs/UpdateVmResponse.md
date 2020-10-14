# UpdateVmResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Vm** | Pointer to [**Vm**](Vm.md) |  | [optional] 

## Methods

### NewUpdateVmResponse

`func NewUpdateVmResponse() *UpdateVmResponse`

NewUpdateVmResponse instantiates a new UpdateVmResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVmResponseWithDefaults

`func NewUpdateVmResponseWithDefaults() *UpdateVmResponse`

NewUpdateVmResponseWithDefaults instantiates a new UpdateVmResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *UpdateVmResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UpdateVmResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UpdateVmResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UpdateVmResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVm

`func (o *UpdateVmResponse) GetVm() Vm`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *UpdateVmResponse) GetVmOk() (*Vm, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *UpdateVmResponse) SetVm(v Vm)`

SetVm sets Vm field to given value.

### HasVm

`func (o *UpdateVmResponse) HasVm() bool`

HasVm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


