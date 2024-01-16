# KeypairCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeypairFingerprint** | Pointer to **string** | The MD5 public key fingerprint, as specified in section 4 of RFC 4716. | [optional] 
**KeypairName** | Pointer to **string** | The name of the keypair. | [optional] 
**KeypairType** | Pointer to **string** | The type of the keypair (&#x60;ssh-rsa&#x60;, &#x60;ssh-ed25519&#x60;, &#x60;ecdsa-sha2-nistp256&#x60;, &#x60;ecdsa-sha2-nistp384&#x60;, or &#x60;ecdsa-sha2-nistp521&#x60;). | [optional] 
**PrivateKey** | Pointer to **string** | The private key, returned only if you are creating a keypair (not if you are importing). When you save this private key in a .rsa file, make sure you replace the &#x60;\\n&#x60; escape sequences with real line breaks. | [optional] 

## Methods

### NewKeypairCreated

`func NewKeypairCreated() *KeypairCreated`

NewKeypairCreated instantiates a new KeypairCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeypairCreatedWithDefaults

`func NewKeypairCreatedWithDefaults() *KeypairCreated`

NewKeypairCreatedWithDefaults instantiates a new KeypairCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypairFingerprint

`func (o *KeypairCreated) GetKeypairFingerprint() string`

GetKeypairFingerprint returns the KeypairFingerprint field if non-nil, zero value otherwise.

### GetKeypairFingerprintOk

`func (o *KeypairCreated) GetKeypairFingerprintOk() (*string, bool)`

GetKeypairFingerprintOk returns a tuple with the KeypairFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairFingerprint

`func (o *KeypairCreated) SetKeypairFingerprint(v string)`

SetKeypairFingerprint sets KeypairFingerprint field to given value.

### HasKeypairFingerprint

`func (o *KeypairCreated) HasKeypairFingerprint() bool`

HasKeypairFingerprint returns a boolean if a field has been set.

### GetKeypairName

`func (o *KeypairCreated) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *KeypairCreated) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *KeypairCreated) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *KeypairCreated) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetKeypairType

`func (o *KeypairCreated) GetKeypairType() string`

GetKeypairType returns the KeypairType field if non-nil, zero value otherwise.

### GetKeypairTypeOk

`func (o *KeypairCreated) GetKeypairTypeOk() (*string, bool)`

GetKeypairTypeOk returns a tuple with the KeypairType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairType

`func (o *KeypairCreated) SetKeypairType(v string)`

SetKeypairType sets KeypairType field to given value.

### HasKeypairType

`func (o *KeypairCreated) HasKeypairType() bool`

HasKeypairType returns a boolean if a field has been set.

### GetPrivateKey

`func (o *KeypairCreated) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KeypairCreated) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *KeypairCreated) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *KeypairCreated) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


