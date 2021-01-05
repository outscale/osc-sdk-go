# EncryptPlaintextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**EncryptionContext** | Pointer to **map[string]string** | Information about the encryption context. | [optional] 
**MasterKeyId** | **string** | The ID of the master key used to encrypt the data. | 
**Plaintext** | **string** | The plaintext you want to encrypt, encoded in base64.&lt;br /&gt; This base64-encoded plaintext must contain between 1 and 4096 characters. | 

## Methods

### NewEncryptPlaintextRequest

`func NewEncryptPlaintextRequest(masterKeyId string, plaintext string, ) *EncryptPlaintextRequest`

NewEncryptPlaintextRequest instantiates a new EncryptPlaintextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptPlaintextRequestWithDefaults

`func NewEncryptPlaintextRequestWithDefaults() *EncryptPlaintextRequest`

NewEncryptPlaintextRequestWithDefaults instantiates a new EncryptPlaintextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *EncryptPlaintextRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *EncryptPlaintextRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *EncryptPlaintextRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *EncryptPlaintextRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetEncryptionContext

`func (o *EncryptPlaintextRequest) GetEncryptionContext() map[string]string`

GetEncryptionContext returns the EncryptionContext field if non-nil, zero value otherwise.

### GetEncryptionContextOk

`func (o *EncryptPlaintextRequest) GetEncryptionContextOk() (*map[string]string, bool)`

GetEncryptionContextOk returns a tuple with the EncryptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionContext

`func (o *EncryptPlaintextRequest) SetEncryptionContext(v map[string]string)`

SetEncryptionContext sets EncryptionContext field to given value.

### HasEncryptionContext

`func (o *EncryptPlaintextRequest) HasEncryptionContext() bool`

HasEncryptionContext returns a boolean if a field has been set.

### GetMasterKeyId

`func (o *EncryptPlaintextRequest) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *EncryptPlaintextRequest) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *EncryptPlaintextRequest) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.


### GetPlaintext

`func (o *EncryptPlaintextRequest) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *EncryptPlaintextRequest) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *EncryptPlaintextRequest) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


