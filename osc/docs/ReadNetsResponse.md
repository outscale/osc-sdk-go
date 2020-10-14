# ReadNetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nets** | Pointer to [**[]Net**](Net.md) | Information about the described Nets. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadNetsResponse

`func NewReadNetsResponse() *ReadNetsResponse`

NewReadNetsResponse instantiates a new ReadNetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadNetsResponseWithDefaults

`func NewReadNetsResponseWithDefaults() *ReadNetsResponse`

NewReadNetsResponseWithDefaults instantiates a new ReadNetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNets

`func (o *ReadNetsResponse) GetNets() []Net`

GetNets returns the Nets field if non-nil, zero value otherwise.

### GetNetsOk

`func (o *ReadNetsResponse) GetNetsOk() (*[]Net, bool)`

GetNetsOk returns a tuple with the Nets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNets

`func (o *ReadNetsResponse) SetNets(v []Net)`

SetNets sets Nets field to given value.

### HasNets

`func (o *ReadNetsResponse) HasNets() bool`

HasNets returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadNetsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNetsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadNetsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadNetsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


