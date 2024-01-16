# ReadDedicatedGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedGroups** | Pointer to [**[]DedicatedGroup**](DedicatedGroup.md) | Information about one or more dedicated groups. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadDedicatedGroupsResponse

`func NewReadDedicatedGroupsResponse() *ReadDedicatedGroupsResponse`

NewReadDedicatedGroupsResponse instantiates a new ReadDedicatedGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadDedicatedGroupsResponseWithDefaults

`func NewReadDedicatedGroupsResponseWithDefaults() *ReadDedicatedGroupsResponse`

NewReadDedicatedGroupsResponseWithDefaults instantiates a new ReadDedicatedGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedGroups

`func (o *ReadDedicatedGroupsResponse) GetDedicatedGroups() []DedicatedGroup`

GetDedicatedGroups returns the DedicatedGroups field if non-nil, zero value otherwise.

### GetDedicatedGroupsOk

`func (o *ReadDedicatedGroupsResponse) GetDedicatedGroupsOk() (*[]DedicatedGroup, bool)`

GetDedicatedGroupsOk returns a tuple with the DedicatedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedGroups

`func (o *ReadDedicatedGroupsResponse) SetDedicatedGroups(v []DedicatedGroup)`

SetDedicatedGroups sets DedicatedGroups field to given value.

### HasDedicatedGroups

`func (o *ReadDedicatedGroupsResponse) HasDedicatedGroups() bool`

HasDedicatedGroups returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadDedicatedGroupsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadDedicatedGroupsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadDedicatedGroupsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadDedicatedGroupsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


