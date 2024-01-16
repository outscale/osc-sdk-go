# Keypair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeypairFingerprint** | Pointer to **string** | The MD5 public key fingerprint as specified in section 4 of RFC 4716. | [optional] 
**KeypairName** | Pointer to **string** | The name of the keypair. | [optional] 
**KeypairType** | Pointer to **string** | The type of the keypair (&#x60;ssh-rsa&#x60;, &#x60;ssh-ed25519&#x60;, &#x60;ecdsa-sha2-nistp256&#x60;, &#x60;ecdsa-sha2-nistp384&#x60;, or &#x60;ecdsa-sha2-nistp521&#x60;). | [optional] 

## Methods

### NewKeypair

`func NewKeypair() *Keypair`

NewKeypair instantiates a new Keypair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeypairWithDefaults

`func NewKeypairWithDefaults() *Keypair`

NewKeypairWithDefaults instantiates a new Keypair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypairFingerprint

`func (o *Keypair) GetKeypairFingerprint() string`

GetKeypairFingerprint returns the KeypairFingerprint field if non-nil, zero value otherwise.

### GetKeypairFingerprintOk

`func (o *Keypair) GetKeypairFingerprintOk() (*string, bool)`

GetKeypairFingerprintOk returns a tuple with the KeypairFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairFingerprint

`func (o *Keypair) SetKeypairFingerprint(v string)`

SetKeypairFingerprint sets KeypairFingerprint field to given value.

### HasKeypairFingerprint

`func (o *Keypair) HasKeypairFingerprint() bool`

HasKeypairFingerprint returns a boolean if a field has been set.

### GetKeypairName

`func (o *Keypair) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *Keypair) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *Keypair) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *Keypair) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetKeypairType

`func (o *Keypair) GetKeypairType() string`

GetKeypairType returns the KeypairType field if non-nil, zero value otherwise.

### GetKeypairTypeOk

`func (o *Keypair) GetKeypairTypeOk() (*string, bool)`

GetKeypairTypeOk returns a tuple with the KeypairType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairType

`func (o *Keypair) SetKeypairType(v string)`

SetKeypairType sets KeypairType field to given value.

### HasKeypairType

`func (o *Keypair) HasKeypairType() bool`

HasKeypairType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


