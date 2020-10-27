# ReadKeypairsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypairs** | Pointer to [**[]Keypair**](Keypair.md) | Information about one or more keypairs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadKeypairsResponse

`func NewReadKeypairsResponse() *ReadKeypairsResponse`

NewReadKeypairsResponse instantiates a new ReadKeypairsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadKeypairsResponseWithDefaults

`func NewReadKeypairsResponseWithDefaults() *ReadKeypairsResponse`

NewReadKeypairsResponseWithDefaults instantiates a new ReadKeypairsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypairs

`func (o *ReadKeypairsResponse) GetKeypairs() []Keypair`

GetKeypairs returns the Keypairs field if non-nil, zero value otherwise.

### GetKeypairsOk

`func (o *ReadKeypairsResponse) GetKeypairsOk() (*[]Keypair, bool)`

GetKeypairsOk returns a tuple with the Keypairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairs

`func (o *ReadKeypairsResponse) SetKeypairs(v []Keypair)`

SetKeypairs sets Keypairs field to given value.

### HasKeypairs

`func (o *ReadKeypairsResponse) HasKeypairs() bool`

HasKeypairs returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadKeypairsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadKeypairsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadKeypairsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadKeypairsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


