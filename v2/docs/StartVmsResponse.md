# StartVmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Vms** | Pointer to [**[]VmState**](VmState.md) | Information about one or more started VMs. | [optional] 

## Methods

### NewStartVmsResponse

`func NewStartVmsResponse() *StartVmsResponse`

NewStartVmsResponse instantiates a new StartVmsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartVmsResponseWithDefaults

`func NewStartVmsResponseWithDefaults() *StartVmsResponse`

NewStartVmsResponseWithDefaults instantiates a new StartVmsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *StartVmsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *StartVmsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *StartVmsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *StartVmsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVms

`func (o *StartVmsResponse) GetVms() []VmState`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *StartVmsResponse) GetVmsOk() (*[]VmState, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *StartVmsResponse) SetVms(v []VmState)`

SetVms sets Vms field to given value.

### HasVms

`func (o *StartVmsResponse) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


