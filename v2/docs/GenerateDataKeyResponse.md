# GenerateDataKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **string** | The encrypted data key, encoded in base64. | [optional] 
**Plaintext** | Pointer to **string** | The decrypted data key, encoded in base64. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewGenerateDataKeyResponse

`func NewGenerateDataKeyResponse() *GenerateDataKeyResponse`

NewGenerateDataKeyResponse instantiates a new GenerateDataKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDataKeyResponseWithDefaults

`func NewGenerateDataKeyResponseWithDefaults() *GenerateDataKeyResponse`

NewGenerateDataKeyResponseWithDefaults instantiates a new GenerateDataKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *GenerateDataKeyResponse) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *GenerateDataKeyResponse) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *GenerateDataKeyResponse) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *GenerateDataKeyResponse) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.

### GetPlaintext

`func (o *GenerateDataKeyResponse) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *GenerateDataKeyResponse) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *GenerateDataKeyResponse) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.

### HasPlaintext

`func (o *GenerateDataKeyResponse) HasPlaintext() bool`

HasPlaintext returns a boolean if a field has been set.

### GetResponseContext

`func (o *GenerateDataKeyResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *GenerateDataKeyResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *GenerateDataKeyResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *GenerateDataKeyResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


