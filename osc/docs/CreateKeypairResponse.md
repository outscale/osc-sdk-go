# CreateKeypairResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypair** | Pointer to [**KeypairCreated**](KeypairCreated.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateKeypairResponse

`func NewCreateKeypairResponse() *CreateKeypairResponse`

NewCreateKeypairResponse instantiates a new CreateKeypairResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeypairResponseWithDefaults

`func NewCreateKeypairResponseWithDefaults() *CreateKeypairResponse`

NewCreateKeypairResponseWithDefaults instantiates a new CreateKeypairResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypair

`func (o *CreateKeypairResponse) GetKeypair() KeypairCreated`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *CreateKeypairResponse) GetKeypairOk() (*KeypairCreated, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *CreateKeypairResponse) SetKeypair(v KeypairCreated)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *CreateKeypairResponse) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateKeypairResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateKeypairResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateKeypairResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateKeypairResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


