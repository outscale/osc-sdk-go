# ReadVmsStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmStates** | Pointer to [**[]VmStates**](VmStates.md) | Information about one or more VM states. | [optional] 

## Methods

### NewReadVmsStateResponse

`func NewReadVmsStateResponse() *ReadVmsStateResponse`

NewReadVmsStateResponse instantiates a new ReadVmsStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmsStateResponseWithDefaults

`func NewReadVmsStateResponseWithDefaults() *ReadVmsStateResponse`

NewReadVmsStateResponseWithDefaults instantiates a new ReadVmsStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ReadVmsStateResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadVmsStateResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadVmsStateResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadVmsStateResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadVmsStateResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmsStateResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVmsStateResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVmsStateResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmStates

`func (o *ReadVmsStateResponse) GetVmStates() []VmStates`

GetVmStates returns the VmStates field if non-nil, zero value otherwise.

### GetVmStatesOk

`func (o *ReadVmsStateResponse) GetVmStatesOk() (*[]VmStates, bool)`

GetVmStatesOk returns a tuple with the VmStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStates

`func (o *ReadVmsStateResponse) SetVmStates(v []VmStates)`

SetVmStates sets VmStates field to given value.

### HasVmStates

`func (o *ReadVmsStateResponse) HasVmStates() bool`

HasVmStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


