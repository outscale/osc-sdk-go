# CreateVmGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmGroup** | Pointer to [**VmGroup**](VmGroup.md) |  | [optional] 

## Methods

### NewCreateVmGroupResponse

`func NewCreateVmGroupResponse() *CreateVmGroupResponse`

NewCreateVmGroupResponse instantiates a new CreateVmGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVmGroupResponseWithDefaults

`func NewCreateVmGroupResponseWithDefaults() *CreateVmGroupResponse`

NewCreateVmGroupResponseWithDefaults instantiates a new CreateVmGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateVmGroupResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateVmGroupResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateVmGroupResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateVmGroupResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmGroup

`func (o *CreateVmGroupResponse) GetVmGroup() VmGroup`

GetVmGroup returns the VmGroup field if non-nil, zero value otherwise.

### GetVmGroupOk

`func (o *CreateVmGroupResponse) GetVmGroupOk() (*VmGroup, bool)`

GetVmGroupOk returns a tuple with the VmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroup

`func (o *CreateVmGroupResponse) SetVmGroup(v VmGroup)`

SetVmGroup sets VmGroup field to given value.

### HasVmGroup

`func (o *CreateVmGroupResponse) HasVmGroup() bool`

HasVmGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


