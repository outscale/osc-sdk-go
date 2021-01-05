# DecryptCiphertextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterKeyId** | Pointer to **string** | The ID of the master key used to decrypt the data. | [optional] 
**Plaintext** | Pointer to **string** | The decrypted ciphertext, encoded in base64. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewDecryptCiphertextResponse

`func NewDecryptCiphertextResponse() *DecryptCiphertextResponse`

NewDecryptCiphertextResponse instantiates a new DecryptCiphertextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptCiphertextResponseWithDefaults

`func NewDecryptCiphertextResponseWithDefaults() *DecryptCiphertextResponse`

NewDecryptCiphertextResponseWithDefaults instantiates a new DecryptCiphertextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterKeyId

`func (o *DecryptCiphertextResponse) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *DecryptCiphertextResponse) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *DecryptCiphertextResponse) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.

### HasMasterKeyId

`func (o *DecryptCiphertextResponse) HasMasterKeyId() bool`

HasMasterKeyId returns a boolean if a field has been set.

### GetPlaintext

`func (o *DecryptCiphertextResponse) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *DecryptCiphertextResponse) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *DecryptCiphertextResponse) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.

### HasPlaintext

`func (o *DecryptCiphertextResponse) HasPlaintext() bool`

HasPlaintext returns a boolean if a field has been set.

### GetResponseContext

`func (o *DecryptCiphertextResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *DecryptCiphertextResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *DecryptCiphertextResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *DecryptCiphertextResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


