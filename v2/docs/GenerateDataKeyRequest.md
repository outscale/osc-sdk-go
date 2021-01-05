# GenerateDataKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**EncryptionContext** | Pointer to **map[string]string** | Information about the encryption context. | [optional] 
**GeneratePlaintext** | Pointer to **bool** | If &#x60;true&#x60;, the data key is generated in plaintext and ciphertext. If &#x60;false&#x60;, the data key is generated only in ciphertext. | [optional] 
**MasterKeyId** | **string** | The ID of the master key used to generate the data key. | 
**Size** | **int32** | The size of the data key you want to generate, in bytes (between 1 and 1024). | 

## Methods

### NewGenerateDataKeyRequest

`func NewGenerateDataKeyRequest(masterKeyId string, size int32, ) *GenerateDataKeyRequest`

NewGenerateDataKeyRequest instantiates a new GenerateDataKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDataKeyRequestWithDefaults

`func NewGenerateDataKeyRequestWithDefaults() *GenerateDataKeyRequest`

NewGenerateDataKeyRequestWithDefaults instantiates a new GenerateDataKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *GenerateDataKeyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *GenerateDataKeyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *GenerateDataKeyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *GenerateDataKeyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetEncryptionContext

`func (o *GenerateDataKeyRequest) GetEncryptionContext() map[string]string`

GetEncryptionContext returns the EncryptionContext field if non-nil, zero value otherwise.

### GetEncryptionContextOk

`func (o *GenerateDataKeyRequest) GetEncryptionContextOk() (*map[string]string, bool)`

GetEncryptionContextOk returns a tuple with the EncryptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionContext

`func (o *GenerateDataKeyRequest) SetEncryptionContext(v map[string]string)`

SetEncryptionContext sets EncryptionContext field to given value.

### HasEncryptionContext

`func (o *GenerateDataKeyRequest) HasEncryptionContext() bool`

HasEncryptionContext returns a boolean if a field has been set.

### GetGeneratePlaintext

`func (o *GenerateDataKeyRequest) GetGeneratePlaintext() bool`

GetGeneratePlaintext returns the GeneratePlaintext field if non-nil, zero value otherwise.

### GetGeneratePlaintextOk

`func (o *GenerateDataKeyRequest) GetGeneratePlaintextOk() (*bool, bool)`

GetGeneratePlaintextOk returns a tuple with the GeneratePlaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratePlaintext

`func (o *GenerateDataKeyRequest) SetGeneratePlaintext(v bool)`

SetGeneratePlaintext sets GeneratePlaintext field to given value.

### HasGeneratePlaintext

`func (o *GenerateDataKeyRequest) HasGeneratePlaintext() bool`

HasGeneratePlaintext returns a boolean if a field has been set.

### GetMasterKeyId

`func (o *GenerateDataKeyRequest) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *GenerateDataKeyRequest) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *GenerateDataKeyRequest) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.


### GetSize

`func (o *GenerateDataKeyRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GenerateDataKeyRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GenerateDataKeyRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


