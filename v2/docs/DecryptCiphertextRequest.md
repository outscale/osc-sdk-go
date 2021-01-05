# DecryptCiphertextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | **string** | The ciphertext you want to decrypt. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**EncryptionContext** | Pointer to **map[string]string** | Information about the encryption context. | [optional] 

## Methods

### NewDecryptCiphertextRequest

`func NewDecryptCiphertextRequest(ciphertext string, ) *DecryptCiphertextRequest`

NewDecryptCiphertextRequest instantiates a new DecryptCiphertextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptCiphertextRequestWithDefaults

`func NewDecryptCiphertextRequestWithDefaults() *DecryptCiphertextRequest`

NewDecryptCiphertextRequestWithDefaults instantiates a new DecryptCiphertextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *DecryptCiphertextRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *DecryptCiphertextRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *DecryptCiphertextRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### GetDryRun

`func (o *DecryptCiphertextRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DecryptCiphertextRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DecryptCiphertextRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DecryptCiphertextRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetEncryptionContext

`func (o *DecryptCiphertextRequest) GetEncryptionContext() map[string]string`

GetEncryptionContext returns the EncryptionContext field if non-nil, zero value otherwise.

### GetEncryptionContextOk

`func (o *DecryptCiphertextRequest) GetEncryptionContextOk() (*map[string]string, bool)`

GetEncryptionContextOk returns a tuple with the EncryptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionContext

`func (o *DecryptCiphertextRequest) SetEncryptionContext(v map[string]string)`

SetEncryptionContext sets EncryptionContext field to given value.

### HasEncryptionContext

`func (o *DecryptCiphertextRequest) HasEncryptionContext() bool`

HasEncryptionContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


