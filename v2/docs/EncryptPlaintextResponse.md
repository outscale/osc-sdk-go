# EncryptPlaintextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **string** | The encrypted plaintext. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewEncryptPlaintextResponse

`func NewEncryptPlaintextResponse() *EncryptPlaintextResponse`

NewEncryptPlaintextResponse instantiates a new EncryptPlaintextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptPlaintextResponseWithDefaults

`func NewEncryptPlaintextResponseWithDefaults() *EncryptPlaintextResponse`

NewEncryptPlaintextResponseWithDefaults instantiates a new EncryptPlaintextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *EncryptPlaintextResponse) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *EncryptPlaintextResponse) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *EncryptPlaintextResponse) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *EncryptPlaintextResponse) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.

### GetResponseContext

`func (o *EncryptPlaintextResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *EncryptPlaintextResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *EncryptPlaintextResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *EncryptPlaintextResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


