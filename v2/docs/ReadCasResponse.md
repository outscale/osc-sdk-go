# ReadCasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cas** | Pointer to [**[]Ca**](Ca.md) | Information about one or more CAs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadCasResponse

`func NewReadCasResponse() *ReadCasResponse`

NewReadCasResponse instantiates a new ReadCasResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadCasResponseWithDefaults

`func NewReadCasResponseWithDefaults() *ReadCasResponse`

NewReadCasResponseWithDefaults instantiates a new ReadCasResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCas

`func (o *ReadCasResponse) GetCas() []Ca`

GetCas returns the Cas field if non-nil, zero value otherwise.

### GetCasOk

`func (o *ReadCasResponse) GetCasOk() (*[]Ca, bool)`

GetCasOk returns a tuple with the Cas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCas

`func (o *ReadCasResponse) SetCas(v []Ca)`

SetCas sets Cas field to given value.

### HasCas

`func (o *ReadCasResponse) HasCas() bool`

HasCas returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadCasResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadCasResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadCasResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadCasResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


