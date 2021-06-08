# Phase1Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpdTimeoutAction** | Pointer to **string** | The action to carry out after a Dead Peer Detection (DPD) timeout occurs. | [optional] 
**DpdTimeoutSeconds** | Pointer to **int32** | The maximum waiting time for a Dead Peer Detection (DPD) response before considering the peer as dead, in seconds. | [optional] 
**IkeVersions** | Pointer to **[]string** | The Internet Key Exchange (IKE) versions allowed for the VPN tunnel. | [optional] 
**Phase1DhGroupNumbers** | Pointer to **[]int32** | The Diffie-Hellman (DH) group numbers allowed for the VPN tunnel for phase 1. | [optional] 
**Phase1EncryptionAlgorithms** | Pointer to **[]string** | The encryption algorithms allowed for the VPN tunnel for phase 1. | [optional] 
**Phase1IntegrityAlgorithms** | Pointer to **[]string** | The integrity algorithms allowed for the VPN tunnel for phase 1. | [optional] 
**Phase1LifetimeSeconds** | Pointer to **int32** | The lifetime for phase 1 of the IKE negotiation process, in seconds. | [optional] 
**ReplayWindowSize** | Pointer to **int32** | The number of packets in an IKE replay window. | [optional] 
**StartupAction** | Pointer to **string** | The action to carry out when establishing tunnels for a VPN connection. | [optional] 

## Methods

### NewPhase1Options

`func NewPhase1Options() *Phase1Options`

NewPhase1Options instantiates a new Phase1Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhase1OptionsWithDefaults

`func NewPhase1OptionsWithDefaults() *Phase1Options`

NewPhase1OptionsWithDefaults instantiates a new Phase1Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpdTimeoutAction

`func (o *Phase1Options) GetDpdTimeoutAction() string`

GetDpdTimeoutAction returns the DpdTimeoutAction field if non-nil, zero value otherwise.

### GetDpdTimeoutActionOk

`func (o *Phase1Options) GetDpdTimeoutActionOk() (*string, bool)`

GetDpdTimeoutActionOk returns a tuple with the DpdTimeoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeoutAction

`func (o *Phase1Options) SetDpdTimeoutAction(v string)`

SetDpdTimeoutAction sets DpdTimeoutAction field to given value.

### HasDpdTimeoutAction

`func (o *Phase1Options) HasDpdTimeoutAction() bool`

HasDpdTimeoutAction returns a boolean if a field has been set.

### GetDpdTimeoutSeconds

`func (o *Phase1Options) GetDpdTimeoutSeconds() int32`

GetDpdTimeoutSeconds returns the DpdTimeoutSeconds field if non-nil, zero value otherwise.

### GetDpdTimeoutSecondsOk

`func (o *Phase1Options) GetDpdTimeoutSecondsOk() (*int32, bool)`

GetDpdTimeoutSecondsOk returns a tuple with the DpdTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeoutSeconds

`func (o *Phase1Options) SetDpdTimeoutSeconds(v int32)`

SetDpdTimeoutSeconds sets DpdTimeoutSeconds field to given value.

### HasDpdTimeoutSeconds

`func (o *Phase1Options) HasDpdTimeoutSeconds() bool`

HasDpdTimeoutSeconds returns a boolean if a field has been set.

### GetIkeVersions

`func (o *Phase1Options) GetIkeVersions() []string`

GetIkeVersions returns the IkeVersions field if non-nil, zero value otherwise.

### GetIkeVersionsOk

`func (o *Phase1Options) GetIkeVersionsOk() (*[]string, bool)`

GetIkeVersionsOk returns a tuple with the IkeVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersions

`func (o *Phase1Options) SetIkeVersions(v []string)`

SetIkeVersions sets IkeVersions field to given value.

### HasIkeVersions

`func (o *Phase1Options) HasIkeVersions() bool`

HasIkeVersions returns a boolean if a field has been set.

### GetPhase1DhGroupNumbers

`func (o *Phase1Options) GetPhase1DhGroupNumbers() []int32`

GetPhase1DhGroupNumbers returns the Phase1DhGroupNumbers field if non-nil, zero value otherwise.

### GetPhase1DhGroupNumbersOk

`func (o *Phase1Options) GetPhase1DhGroupNumbersOk() (*[]int32, bool)`

GetPhase1DhGroupNumbersOk returns a tuple with the Phase1DhGroupNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1DhGroupNumbers

`func (o *Phase1Options) SetPhase1DhGroupNumbers(v []int32)`

SetPhase1DhGroupNumbers sets Phase1DhGroupNumbers field to given value.

### HasPhase1DhGroupNumbers

`func (o *Phase1Options) HasPhase1DhGroupNumbers() bool`

HasPhase1DhGroupNumbers returns a boolean if a field has been set.

### GetPhase1EncryptionAlgorithms

`func (o *Phase1Options) GetPhase1EncryptionAlgorithms() []string`

GetPhase1EncryptionAlgorithms returns the Phase1EncryptionAlgorithms field if non-nil, zero value otherwise.

### GetPhase1EncryptionAlgorithmsOk

`func (o *Phase1Options) GetPhase1EncryptionAlgorithmsOk() (*[]string, bool)`

GetPhase1EncryptionAlgorithmsOk returns a tuple with the Phase1EncryptionAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1EncryptionAlgorithms

`func (o *Phase1Options) SetPhase1EncryptionAlgorithms(v []string)`

SetPhase1EncryptionAlgorithms sets Phase1EncryptionAlgorithms field to given value.

### HasPhase1EncryptionAlgorithms

`func (o *Phase1Options) HasPhase1EncryptionAlgorithms() bool`

HasPhase1EncryptionAlgorithms returns a boolean if a field has been set.

### GetPhase1IntegrityAlgorithms

`func (o *Phase1Options) GetPhase1IntegrityAlgorithms() []string`

GetPhase1IntegrityAlgorithms returns the Phase1IntegrityAlgorithms field if non-nil, zero value otherwise.

### GetPhase1IntegrityAlgorithmsOk

`func (o *Phase1Options) GetPhase1IntegrityAlgorithmsOk() (*[]string, bool)`

GetPhase1IntegrityAlgorithmsOk returns a tuple with the Phase1IntegrityAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1IntegrityAlgorithms

`func (o *Phase1Options) SetPhase1IntegrityAlgorithms(v []string)`

SetPhase1IntegrityAlgorithms sets Phase1IntegrityAlgorithms field to given value.

### HasPhase1IntegrityAlgorithms

`func (o *Phase1Options) HasPhase1IntegrityAlgorithms() bool`

HasPhase1IntegrityAlgorithms returns a boolean if a field has been set.

### GetPhase1LifetimeSeconds

`func (o *Phase1Options) GetPhase1LifetimeSeconds() int32`

GetPhase1LifetimeSeconds returns the Phase1LifetimeSeconds field if non-nil, zero value otherwise.

### GetPhase1LifetimeSecondsOk

`func (o *Phase1Options) GetPhase1LifetimeSecondsOk() (*int32, bool)`

GetPhase1LifetimeSecondsOk returns a tuple with the Phase1LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1LifetimeSeconds

`func (o *Phase1Options) SetPhase1LifetimeSeconds(v int32)`

SetPhase1LifetimeSeconds sets Phase1LifetimeSeconds field to given value.

### HasPhase1LifetimeSeconds

`func (o *Phase1Options) HasPhase1LifetimeSeconds() bool`

HasPhase1LifetimeSeconds returns a boolean if a field has been set.

### GetReplayWindowSize

`func (o *Phase1Options) GetReplayWindowSize() int32`

GetReplayWindowSize returns the ReplayWindowSize field if non-nil, zero value otherwise.

### GetReplayWindowSizeOk

`func (o *Phase1Options) GetReplayWindowSizeOk() (*int32, bool)`

GetReplayWindowSizeOk returns a tuple with the ReplayWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayWindowSize

`func (o *Phase1Options) SetReplayWindowSize(v int32)`

SetReplayWindowSize sets ReplayWindowSize field to given value.

### HasReplayWindowSize

`func (o *Phase1Options) HasReplayWindowSize() bool`

HasReplayWindowSize returns a boolean if a field has been set.

### GetStartupAction

`func (o *Phase1Options) GetStartupAction() string`

GetStartupAction returns the StartupAction field if non-nil, zero value otherwise.

### GetStartupActionOk

`func (o *Phase1Options) GetStartupActionOk() (*string, bool)`

GetStartupActionOk returns a tuple with the StartupAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupAction

`func (o *Phase1Options) SetStartupAction(v string)`

SetStartupAction sets StartupAction field to given value.

### HasStartupAction

`func (o *Phase1Options) HasStartupAction() bool`

HasStartupAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


