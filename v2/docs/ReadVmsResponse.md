# ReadVmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Vms** | Pointer to [**[]Vm**](Vm.md) | Information about one or more VMs. | [optional] 

## Methods

### NewReadVmsResponse

`func NewReadVmsResponse() *ReadVmsResponse`

NewReadVmsResponse instantiates a new ReadVmsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmsResponseWithDefaults

`func NewReadVmsResponseWithDefaults() *ReadVmsResponse`

NewReadVmsResponseWithDefaults instantiates a new ReadVmsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ReadVmsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadVmsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadVmsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadVmsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadVmsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVmsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVmsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVms

`func (o *ReadVmsResponse) GetVms() []Vm`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *ReadVmsResponse) GetVmsOk() (*[]Vm, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *ReadVmsResponse) SetVms(v []Vm)`

SetVms sets Vms field to given value.

### HasVms

`func (o *ReadVmsResponse) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


