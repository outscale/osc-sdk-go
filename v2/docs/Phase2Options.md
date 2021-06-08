# Phase2Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase2DhGroupNumbers** | Pointer to **[]int32** | The Diffie-Hellman (DH) group numbers allowed for the VPN tunnel for phase 2. | [optional] 
**Phase2EncryptionAlgorithms** | Pointer to **[]string** | The encryption algorithms allowed for the VPN tunnel for phase 2. | [optional] 
**Phase2IntegrityAlgorithms** | Pointer to **[]string** | The integrity algorithms allowed for the VPN tunnel for phase 2. | [optional] 
**Phase2LifetimeSeconds** | Pointer to **int32** | The lifetime for phase 2 of the Internet Key Exchange (IKE) negociation process, in seconds. | [optional] 
**PreSharedKey** | Pointer to **string** | The pre-shared key to establish the initial authentication between the client gateway and the virtual gateway. | [optional] 

## Methods

### NewPhase2Options

`func NewPhase2Options() *Phase2Options`

NewPhase2Options instantiates a new Phase2Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhase2OptionsWithDefaults

`func NewPhase2OptionsWithDefaults() *Phase2Options`

NewPhase2OptionsWithDefaults instantiates a new Phase2Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase2DhGroupNumbers

`func (o *Phase2Options) GetPhase2DhGroupNumbers() []int32`

GetPhase2DhGroupNumbers returns the Phase2DhGroupNumbers field if non-nil, zero value otherwise.

### GetPhase2DhGroupNumbersOk

`func (o *Phase2Options) GetPhase2DhGroupNumbersOk() (*[]int32, bool)`

GetPhase2DhGroupNumbersOk returns a tuple with the Phase2DhGroupNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2DhGroupNumbers

`func (o *Phase2Options) SetPhase2DhGroupNumbers(v []int32)`

SetPhase2DhGroupNumbers sets Phase2DhGroupNumbers field to given value.

### HasPhase2DhGroupNumbers

`func (o *Phase2Options) HasPhase2DhGroupNumbers() bool`

HasPhase2DhGroupNumbers returns a boolean if a field has been set.

### GetPhase2EncryptionAlgorithms

`func (o *Phase2Options) GetPhase2EncryptionAlgorithms() []string`

GetPhase2EncryptionAlgorithms returns the Phase2EncryptionAlgorithms field if non-nil, zero value otherwise.

### GetPhase2EncryptionAlgorithmsOk

`func (o *Phase2Options) GetPhase2EncryptionAlgorithmsOk() (*[]string, bool)`

GetPhase2EncryptionAlgorithmsOk returns a tuple with the Phase2EncryptionAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2EncryptionAlgorithms

`func (o *Phase2Options) SetPhase2EncryptionAlgorithms(v []string)`

SetPhase2EncryptionAlgorithms sets Phase2EncryptionAlgorithms field to given value.

### HasPhase2EncryptionAlgorithms

`func (o *Phase2Options) HasPhase2EncryptionAlgorithms() bool`

HasPhase2EncryptionAlgorithms returns a boolean if a field has been set.

### GetPhase2IntegrityAlgorithms

`func (o *Phase2Options) GetPhase2IntegrityAlgorithms() []string`

GetPhase2IntegrityAlgorithms returns the Phase2IntegrityAlgorithms field if non-nil, zero value otherwise.

### GetPhase2IntegrityAlgorithmsOk

`func (o *Phase2Options) GetPhase2IntegrityAlgorithmsOk() (*[]string, bool)`

GetPhase2IntegrityAlgorithmsOk returns a tuple with the Phase2IntegrityAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2IntegrityAlgorithms

`func (o *Phase2Options) SetPhase2IntegrityAlgorithms(v []string)`

SetPhase2IntegrityAlgorithms sets Phase2IntegrityAlgorithms field to given value.

### HasPhase2IntegrityAlgorithms

`func (o *Phase2Options) HasPhase2IntegrityAlgorithms() bool`

HasPhase2IntegrityAlgorithms returns a boolean if a field has been set.

### GetPhase2LifetimeSeconds

`func (o *Phase2Options) GetPhase2LifetimeSeconds() int32`

GetPhase2LifetimeSeconds returns the Phase2LifetimeSeconds field if non-nil, zero value otherwise.

### GetPhase2LifetimeSecondsOk

`func (o *Phase2Options) GetPhase2LifetimeSecondsOk() (*int32, bool)`

GetPhase2LifetimeSecondsOk returns a tuple with the Phase2LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2LifetimeSeconds

`func (o *Phase2Options) SetPhase2LifetimeSeconds(v int32)`

SetPhase2LifetimeSeconds sets Phase2LifetimeSeconds field to given value.

### HasPhase2LifetimeSeconds

`func (o *Phase2Options) HasPhase2LifetimeSeconds() bool`

HasPhase2LifetimeSeconds returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *Phase2Options) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *Phase2Options) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *Phase2Options) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *Phase2Options) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


