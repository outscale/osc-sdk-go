# CreateVmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Vms** | Pointer to [**[]Vm**](Vm.md) | Information about one or more created VMs. | [optional] 

## Methods

### NewCreateVmsResponse

`func NewCreateVmsResponse() *CreateVmsResponse`

NewCreateVmsResponse instantiates a new CreateVmsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVmsResponseWithDefaults

`func NewCreateVmsResponseWithDefaults() *CreateVmsResponse`

NewCreateVmsResponseWithDefaults instantiates a new CreateVmsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateVmsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateVmsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateVmsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateVmsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVms

`func (o *CreateVmsResponse) GetVms() []Vm`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *CreateVmsResponse) GetVmsOk() (*[]Vm, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *CreateVmsResponse) SetVms(v []Vm)`

SetVms sets Vms field to given value.

### HasVms

`func (o *CreateVmsResponse) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


