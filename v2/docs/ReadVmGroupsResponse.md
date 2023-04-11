# ReadVmGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmGroups** | Pointer to [**[]VmGroup**](VmGroup.md) | Information about one or more VM groups. | [optional] 

## Methods

### NewReadVmGroupsResponse

`func NewReadVmGroupsResponse() *ReadVmGroupsResponse`

NewReadVmGroupsResponse instantiates a new ReadVmGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmGroupsResponseWithDefaults

`func NewReadVmGroupsResponseWithDefaults() *ReadVmGroupsResponse`

NewReadVmGroupsResponseWithDefaults instantiates a new ReadVmGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadVmGroupsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmGroupsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVmGroupsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVmGroupsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmGroups

`func (o *ReadVmGroupsResponse) GetVmGroups() []VmGroup`

GetVmGroups returns the VmGroups field if non-nil, zero value otherwise.

### GetVmGroupsOk

`func (o *ReadVmGroupsResponse) GetVmGroupsOk() (*[]VmGroup, bool)`

GetVmGroupsOk returns a tuple with the VmGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroups

`func (o *ReadVmGroupsResponse) SetVmGroups(v []VmGroup)`

SetVmGroups sets VmGroups field to given value.

### HasVmGroups

`func (o *ReadVmGroupsResponse) HasVmGroups() bool`

HasVmGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


