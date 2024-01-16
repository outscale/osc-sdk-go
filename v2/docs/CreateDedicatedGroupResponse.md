# CreateDedicatedGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedGroup** | Pointer to [**DedicatedGroup**](DedicatedGroup.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateDedicatedGroupResponse

`func NewCreateDedicatedGroupResponse() *CreateDedicatedGroupResponse`

NewCreateDedicatedGroupResponse instantiates a new CreateDedicatedGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDedicatedGroupResponseWithDefaults

`func NewCreateDedicatedGroupResponseWithDefaults() *CreateDedicatedGroupResponse`

NewCreateDedicatedGroupResponseWithDefaults instantiates a new CreateDedicatedGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedGroup

`func (o *CreateDedicatedGroupResponse) GetDedicatedGroup() DedicatedGroup`

GetDedicatedGroup returns the DedicatedGroup field if non-nil, zero value otherwise.

### GetDedicatedGroupOk

`func (o *CreateDedicatedGroupResponse) GetDedicatedGroupOk() (*DedicatedGroup, bool)`

GetDedicatedGroupOk returns a tuple with the DedicatedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedGroup

`func (o *CreateDedicatedGroupResponse) SetDedicatedGroup(v DedicatedGroup)`

SetDedicatedGroup sets DedicatedGroup field to given value.

### HasDedicatedGroup

`func (o *CreateDedicatedGroupResponse) HasDedicatedGroup() bool`

HasDedicatedGroup returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateDedicatedGroupResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateDedicatedGroupResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateDedicatedGroupResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateDedicatedGroupResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


