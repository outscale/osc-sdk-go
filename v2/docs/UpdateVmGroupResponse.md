# UpdateVmGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmGroup** | Pointer to [**VmGroup**](VmGroup.md) |  | [optional] 

## Methods

### NewUpdateVmGroupResponse

`func NewUpdateVmGroupResponse() *UpdateVmGroupResponse`

NewUpdateVmGroupResponse instantiates a new UpdateVmGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVmGroupResponseWithDefaults

`func NewUpdateVmGroupResponseWithDefaults() *UpdateVmGroupResponse`

NewUpdateVmGroupResponseWithDefaults instantiates a new UpdateVmGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *UpdateVmGroupResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UpdateVmGroupResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UpdateVmGroupResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UpdateVmGroupResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmGroup

`func (o *UpdateVmGroupResponse) GetVmGroup() VmGroup`

GetVmGroup returns the VmGroup field if non-nil, zero value otherwise.

### GetVmGroupOk

`func (o *UpdateVmGroupResponse) GetVmGroupOk() (*VmGroup, bool)`

GetVmGroupOk returns a tuple with the VmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroup

`func (o *UpdateVmGroupResponse) SetVmGroup(v VmGroup)`

SetVmGroup sets VmGroup field to given value.

### HasVmGroup

`func (o *UpdateVmGroupResponse) HasVmGroup() bool`

HasVmGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


