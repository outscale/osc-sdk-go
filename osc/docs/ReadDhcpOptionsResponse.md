# ReadDhcpOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptionsSets** | Pointer to [**[]DhcpOptionsSet**](DhcpOptionsSet.md) | Information about one or more DHCP options sets. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadDhcpOptionsResponse

`func NewReadDhcpOptionsResponse() *ReadDhcpOptionsResponse`

NewReadDhcpOptionsResponse instantiates a new ReadDhcpOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadDhcpOptionsResponseWithDefaults

`func NewReadDhcpOptionsResponseWithDefaults() *ReadDhcpOptionsResponse`

NewReadDhcpOptionsResponseWithDefaults instantiates a new ReadDhcpOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpOptionsSets

`func (o *ReadDhcpOptionsResponse) GetDhcpOptionsSets() []DhcpOptionsSet`

GetDhcpOptionsSets returns the DhcpOptionsSets field if non-nil, zero value otherwise.

### GetDhcpOptionsSetsOk

`func (o *ReadDhcpOptionsResponse) GetDhcpOptionsSetsOk() (*[]DhcpOptionsSet, bool)`

GetDhcpOptionsSetsOk returns a tuple with the DhcpOptionsSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsSets

`func (o *ReadDhcpOptionsResponse) SetDhcpOptionsSets(v []DhcpOptionsSet)`

SetDhcpOptionsSets sets DhcpOptionsSets field to given value.

### HasDhcpOptionsSets

`func (o *ReadDhcpOptionsResponse) HasDhcpOptionsSets() bool`

HasDhcpOptionsSets returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadDhcpOptionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadDhcpOptionsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadDhcpOptionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadDhcpOptionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


