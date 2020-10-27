# CreateDhcpOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptionsSet** | Pointer to [**DhcpOptionsSet**](DhcpOptionsSet.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateDhcpOptionsResponse

`func NewCreateDhcpOptionsResponse() *CreateDhcpOptionsResponse`

NewCreateDhcpOptionsResponse instantiates a new CreateDhcpOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDhcpOptionsResponseWithDefaults

`func NewCreateDhcpOptionsResponseWithDefaults() *CreateDhcpOptionsResponse`

NewCreateDhcpOptionsResponseWithDefaults instantiates a new CreateDhcpOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpOptionsSet

`func (o *CreateDhcpOptionsResponse) GetDhcpOptionsSet() DhcpOptionsSet`

GetDhcpOptionsSet returns the DhcpOptionsSet field if non-nil, zero value otherwise.

### GetDhcpOptionsSetOk

`func (o *CreateDhcpOptionsResponse) GetDhcpOptionsSetOk() (*DhcpOptionsSet, bool)`

GetDhcpOptionsSetOk returns a tuple with the DhcpOptionsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsSet

`func (o *CreateDhcpOptionsResponse) SetDhcpOptionsSet(v DhcpOptionsSet)`

SetDhcpOptionsSet sets DhcpOptionsSet field to given value.

### HasDhcpOptionsSet

`func (o *CreateDhcpOptionsResponse) HasDhcpOptionsSet() bool`

HasDhcpOptionsSet returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateDhcpOptionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateDhcpOptionsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateDhcpOptionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateDhcpOptionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


