# DeleteVmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Vms** | Pointer to [**[]VmState**](VmState.md) | Information about one or more terminated VMs. | [optional] 

## Methods

### NewDeleteVmsResponse

`func NewDeleteVmsResponse() *DeleteVmsResponse`

NewDeleteVmsResponse instantiates a new DeleteVmsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVmsResponseWithDefaults

`func NewDeleteVmsResponseWithDefaults() *DeleteVmsResponse`

NewDeleteVmsResponseWithDefaults instantiates a new DeleteVmsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *DeleteVmsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *DeleteVmsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *DeleteVmsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *DeleteVmsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVms

`func (o *DeleteVmsResponse) GetVms() []VmState`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *DeleteVmsResponse) GetVmsOk() (*[]VmState, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *DeleteVmsResponse) SetVms(v []VmState)`

SetVms sets Vms field to given value.

### HasVms

`func (o *DeleteVmsResponse) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


