# ReadSubregionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Subregions** | Pointer to [**[]Subregion**](Subregion.md) | Information about one or more Subregions. | [optional] 

## Methods

### NewReadSubregionsResponse

`func NewReadSubregionsResponse() *ReadSubregionsResponse`

NewReadSubregionsResponse instantiates a new ReadSubregionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSubregionsResponseWithDefaults

`func NewReadSubregionsResponseWithDefaults() *ReadSubregionsResponse`

NewReadSubregionsResponseWithDefaults instantiates a new ReadSubregionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadSubregionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSubregionsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadSubregionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadSubregionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSubregions

`func (o *ReadSubregionsResponse) GetSubregions() []Subregion`

GetSubregions returns the Subregions field if non-nil, zero value otherwise.

### GetSubregionsOk

`func (o *ReadSubregionsResponse) GetSubregionsOk() (*[]Subregion, bool)`

GetSubregionsOk returns a tuple with the Subregions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregions

`func (o *ReadSubregionsResponse) SetSubregions(v []Subregion)`

SetSubregions sets Subregions field to given value.

### HasSubregions

`func (o *ReadSubregionsResponse) HasSubregions() bool`

HasSubregions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


