# ReadNicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nics** | Pointer to [**[]Nic**](Nic.md) | Information about one or more NICs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadNicsResponse

`func NewReadNicsResponse() *ReadNicsResponse`

NewReadNicsResponse instantiates a new ReadNicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadNicsResponseWithDefaults

`func NewReadNicsResponseWithDefaults() *ReadNicsResponse`

NewReadNicsResponseWithDefaults instantiates a new ReadNicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNics

`func (o *ReadNicsResponse) GetNics() []Nic`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *ReadNicsResponse) GetNicsOk() (*[]Nic, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *ReadNicsResponse) SetNics(v []Nic)`

SetNics sets Nics field to given value.

### HasNics

`func (o *ReadNicsResponse) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadNicsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNicsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadNicsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadNicsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


