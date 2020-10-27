# ReadProductTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTypes** | Pointer to [**[]ProductType**](ProductType.md) | Information about one or more product types. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadProductTypesResponse

`func NewReadProductTypesResponse() *ReadProductTypesResponse`

NewReadProductTypesResponse instantiates a new ReadProductTypesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadProductTypesResponseWithDefaults

`func NewReadProductTypesResponseWithDefaults() *ReadProductTypesResponse`

NewReadProductTypesResponseWithDefaults instantiates a new ReadProductTypesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTypes

`func (o *ReadProductTypesResponse) GetProductTypes() []ProductType`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *ReadProductTypesResponse) GetProductTypesOk() (*[]ProductType, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *ReadProductTypesResponse) SetProductTypes(v []ProductType)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *ReadProductTypesResponse) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadProductTypesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadProductTypesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadProductTypesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadProductTypesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


