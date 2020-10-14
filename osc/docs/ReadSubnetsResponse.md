# ReadSubnetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) | Information about one or more Subnets. | [optional] 

## Methods

### NewReadSubnetsResponse

`func NewReadSubnetsResponse() *ReadSubnetsResponse`

NewReadSubnetsResponse instantiates a new ReadSubnetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSubnetsResponseWithDefaults

`func NewReadSubnetsResponseWithDefaults() *ReadSubnetsResponse`

NewReadSubnetsResponseWithDefaults instantiates a new ReadSubnetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadSubnetsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSubnetsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadSubnetsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadSubnetsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSubnets

`func (o *ReadSubnetsResponse) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *ReadSubnetsResponse) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *ReadSubnetsResponse) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *ReadSubnetsResponse) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


