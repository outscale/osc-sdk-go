/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// Phase1Options Information about Phase 1 of the Internet Key Exchange (IKE) negotiation. When Phase 1 finishes successfully, peers proceed to Phase 2 negotiations.
type Phase1Options struct {
	// The action to carry out after a Dead Peer Detection (DPD) timeout occurs.
	DpdTimeoutAction *string `json:"DpdTimeoutAction,omitempty"`
	// The maximum waiting time for a Dead Peer Detection (DPD) response before considering the peer as dead, in seconds.
	DpdTimeoutSeconds *int32 `json:"DpdTimeoutSeconds,omitempty"`
	// The Internet Key Exchange (IKE) versions allowed for the VPN tunnel.
	IkeVersions *[]string `json:"IkeVersions,omitempty"`
	// The Diffie-Hellman (DH) group numbers allowed for the VPN tunnel for phase 1.
	Phase1DhGroupNumbers *[]int32 `json:"Phase1DhGroupNumbers,omitempty"`
	// The encryption algorithms allowed for the VPN tunnel for phase 1.
	Phase1EncryptionAlgorithms *[]string `json:"Phase1EncryptionAlgorithms,omitempty"`
	// The integrity algorithms allowed for the VPN tunnel for phase 1.
	Phase1IntegrityAlgorithms *[]string `json:"Phase1IntegrityAlgorithms,omitempty"`
	// The lifetime for phase 1 of the IKE negotiation process, in seconds.
	Phase1LifetimeSeconds *int32 `json:"Phase1LifetimeSeconds,omitempty"`
	// The number of packets in an IKE replay window.
	ReplayWindowSize *int32 `json:"ReplayWindowSize,omitempty"`
	// The action to carry out when establishing tunnels for a VPN connection.
	StartupAction *string `json:"StartupAction,omitempty"`
}

// NewPhase1Options instantiates a new Phase1Options object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhase1Options() *Phase1Options {
	this := Phase1Options{}
	return &this
}

// NewPhase1OptionsWithDefaults instantiates a new Phase1Options object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhase1OptionsWithDefaults() *Phase1Options {
	this := Phase1Options{}
	return &this
}

// GetDpdTimeoutAction returns the DpdTimeoutAction field value if set, zero value otherwise.
func (o *Phase1Options) GetDpdTimeoutAction() string {
	if o == nil || o.DpdTimeoutAction == nil {
		var ret string
		return ret
	}
	return *o.DpdTimeoutAction
}

// GetDpdTimeoutActionOk returns a tuple with the DpdTimeoutAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetDpdTimeoutActionOk() (*string, bool) {
	if o == nil || o.DpdTimeoutAction == nil {
		return nil, false
	}
	return o.DpdTimeoutAction, true
}

// HasDpdTimeoutAction returns a boolean if a field has been set.
func (o *Phase1Options) HasDpdTimeoutAction() bool {
	if o != nil && o.DpdTimeoutAction != nil {
		return true
	}

	return false
}

// SetDpdTimeoutAction gets a reference to the given string and assigns it to the DpdTimeoutAction field.
func (o *Phase1Options) SetDpdTimeoutAction(v string) {
	o.DpdTimeoutAction = &v
}

// GetDpdTimeoutSeconds returns the DpdTimeoutSeconds field value if set, zero value otherwise.
func (o *Phase1Options) GetDpdTimeoutSeconds() int32 {
	if o == nil || o.DpdTimeoutSeconds == nil {
		var ret int32
		return ret
	}
	return *o.DpdTimeoutSeconds
}

// GetDpdTimeoutSecondsOk returns a tuple with the DpdTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetDpdTimeoutSecondsOk() (*int32, bool) {
	if o == nil || o.DpdTimeoutSeconds == nil {
		return nil, false
	}
	return o.DpdTimeoutSeconds, true
}

// HasDpdTimeoutSeconds returns a boolean if a field has been set.
func (o *Phase1Options) HasDpdTimeoutSeconds() bool {
	if o != nil && o.DpdTimeoutSeconds != nil {
		return true
	}

	return false
}

// SetDpdTimeoutSeconds gets a reference to the given int32 and assigns it to the DpdTimeoutSeconds field.
func (o *Phase1Options) SetDpdTimeoutSeconds(v int32) {
	o.DpdTimeoutSeconds = &v
}

// GetIkeVersions returns the IkeVersions field value if set, zero value otherwise.
func (o *Phase1Options) GetIkeVersions() []string {
	if o == nil || o.IkeVersions == nil {
		var ret []string
		return ret
	}
	return *o.IkeVersions
}

// GetIkeVersionsOk returns a tuple with the IkeVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetIkeVersionsOk() (*[]string, bool) {
	if o == nil || o.IkeVersions == nil {
		return nil, false
	}
	return o.IkeVersions, true
}

// HasIkeVersions returns a boolean if a field has been set.
func (o *Phase1Options) HasIkeVersions() bool {
	if o != nil && o.IkeVersions != nil {
		return true
	}

	return false
}

// SetIkeVersions gets a reference to the given []string and assigns it to the IkeVersions field.
func (o *Phase1Options) SetIkeVersions(v []string) {
	o.IkeVersions = &v
}

// GetPhase1DhGroupNumbers returns the Phase1DhGroupNumbers field value if set, zero value otherwise.
func (o *Phase1Options) GetPhase1DhGroupNumbers() []int32 {
	if o == nil || o.Phase1DhGroupNumbers == nil {
		var ret []int32
		return ret
	}
	return *o.Phase1DhGroupNumbers
}

// GetPhase1DhGroupNumbersOk returns a tuple with the Phase1DhGroupNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetPhase1DhGroupNumbersOk() (*[]int32, bool) {
	if o == nil || o.Phase1DhGroupNumbers == nil {
		return nil, false
	}
	return o.Phase1DhGroupNumbers, true
}

// HasPhase1DhGroupNumbers returns a boolean if a field has been set.
func (o *Phase1Options) HasPhase1DhGroupNumbers() bool {
	if o != nil && o.Phase1DhGroupNumbers != nil {
		return true
	}

	return false
}

// SetPhase1DhGroupNumbers gets a reference to the given []int32 and assigns it to the Phase1DhGroupNumbers field.
func (o *Phase1Options) SetPhase1DhGroupNumbers(v []int32) {
	o.Phase1DhGroupNumbers = &v
}

// GetPhase1EncryptionAlgorithms returns the Phase1EncryptionAlgorithms field value if set, zero value otherwise.
func (o *Phase1Options) GetPhase1EncryptionAlgorithms() []string {
	if o == nil || o.Phase1EncryptionAlgorithms == nil {
		var ret []string
		return ret
	}
	return *o.Phase1EncryptionAlgorithms
}

// GetPhase1EncryptionAlgorithmsOk returns a tuple with the Phase1EncryptionAlgorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetPhase1EncryptionAlgorithmsOk() (*[]string, bool) {
	if o == nil || o.Phase1EncryptionAlgorithms == nil {
		return nil, false
	}
	return o.Phase1EncryptionAlgorithms, true
}

// HasPhase1EncryptionAlgorithms returns a boolean if a field has been set.
func (o *Phase1Options) HasPhase1EncryptionAlgorithms() bool {
	if o != nil && o.Phase1EncryptionAlgorithms != nil {
		return true
	}

	return false
}

// SetPhase1EncryptionAlgorithms gets a reference to the given []string and assigns it to the Phase1EncryptionAlgorithms field.
func (o *Phase1Options) SetPhase1EncryptionAlgorithms(v []string) {
	o.Phase1EncryptionAlgorithms = &v
}

// GetPhase1IntegrityAlgorithms returns the Phase1IntegrityAlgorithms field value if set, zero value otherwise.
func (o *Phase1Options) GetPhase1IntegrityAlgorithms() []string {
	if o == nil || o.Phase1IntegrityAlgorithms == nil {
		var ret []string
		return ret
	}
	return *o.Phase1IntegrityAlgorithms
}

// GetPhase1IntegrityAlgorithmsOk returns a tuple with the Phase1IntegrityAlgorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetPhase1IntegrityAlgorithmsOk() (*[]string, bool) {
	if o == nil || o.Phase1IntegrityAlgorithms == nil {
		return nil, false
	}
	return o.Phase1IntegrityAlgorithms, true
}

// HasPhase1IntegrityAlgorithms returns a boolean if a field has been set.
func (o *Phase1Options) HasPhase1IntegrityAlgorithms() bool {
	if o != nil && o.Phase1IntegrityAlgorithms != nil {
		return true
	}

	return false
}

// SetPhase1IntegrityAlgorithms gets a reference to the given []string and assigns it to the Phase1IntegrityAlgorithms field.
func (o *Phase1Options) SetPhase1IntegrityAlgorithms(v []string) {
	o.Phase1IntegrityAlgorithms = &v
}

// GetPhase1LifetimeSeconds returns the Phase1LifetimeSeconds field value if set, zero value otherwise.
func (o *Phase1Options) GetPhase1LifetimeSeconds() int32 {
	if o == nil || o.Phase1LifetimeSeconds == nil {
		var ret int32
		return ret
	}
	return *o.Phase1LifetimeSeconds
}

// GetPhase1LifetimeSecondsOk returns a tuple with the Phase1LifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetPhase1LifetimeSecondsOk() (*int32, bool) {
	if o == nil || o.Phase1LifetimeSeconds == nil {
		return nil, false
	}
	return o.Phase1LifetimeSeconds, true
}

// HasPhase1LifetimeSeconds returns a boolean if a field has been set.
func (o *Phase1Options) HasPhase1LifetimeSeconds() bool {
	if o != nil && o.Phase1LifetimeSeconds != nil {
		return true
	}

	return false
}

// SetPhase1LifetimeSeconds gets a reference to the given int32 and assigns it to the Phase1LifetimeSeconds field.
func (o *Phase1Options) SetPhase1LifetimeSeconds(v int32) {
	o.Phase1LifetimeSeconds = &v
}

// GetReplayWindowSize returns the ReplayWindowSize field value if set, zero value otherwise.
func (o *Phase1Options) GetReplayWindowSize() int32 {
	if o == nil || o.ReplayWindowSize == nil {
		var ret int32
		return ret
	}
	return *o.ReplayWindowSize
}

// GetReplayWindowSizeOk returns a tuple with the ReplayWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetReplayWindowSizeOk() (*int32, bool) {
	if o == nil || o.ReplayWindowSize == nil {
		return nil, false
	}
	return o.ReplayWindowSize, true
}

// HasReplayWindowSize returns a boolean if a field has been set.
func (o *Phase1Options) HasReplayWindowSize() bool {
	if o != nil && o.ReplayWindowSize != nil {
		return true
	}

	return false
}

// SetReplayWindowSize gets a reference to the given int32 and assigns it to the ReplayWindowSize field.
func (o *Phase1Options) SetReplayWindowSize(v int32) {
	o.ReplayWindowSize = &v
}

// GetStartupAction returns the StartupAction field value if set, zero value otherwise.
func (o *Phase1Options) GetStartupAction() string {
	if o == nil || o.StartupAction == nil {
		var ret string
		return ret
	}
	return *o.StartupAction
}

// GetStartupActionOk returns a tuple with the StartupAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phase1Options) GetStartupActionOk() (*string, bool) {
	if o == nil || o.StartupAction == nil {
		return nil, false
	}
	return o.StartupAction, true
}

// HasStartupAction returns a boolean if a field has been set.
func (o *Phase1Options) HasStartupAction() bool {
	if o != nil && o.StartupAction != nil {
		return true
	}

	return false
}

// SetStartupAction gets a reference to the given string and assigns it to the StartupAction field.
func (o *Phase1Options) SetStartupAction(v string) {
	o.StartupAction = &v
}

func (o Phase1Options) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DpdTimeoutAction != nil {
		toSerialize["DpdTimeoutAction"] = o.DpdTimeoutAction
	}
	if o.DpdTimeoutSeconds != nil {
		toSerialize["DpdTimeoutSeconds"] = o.DpdTimeoutSeconds
	}
	if o.IkeVersions != nil {
		toSerialize["IkeVersions"] = o.IkeVersions
	}
	if o.Phase1DhGroupNumbers != nil {
		toSerialize["Phase1DhGroupNumbers"] = o.Phase1DhGroupNumbers
	}
	if o.Phase1EncryptionAlgorithms != nil {
		toSerialize["Phase1EncryptionAlgorithms"] = o.Phase1EncryptionAlgorithms
	}
	if o.Phase1IntegrityAlgorithms != nil {
		toSerialize["Phase1IntegrityAlgorithms"] = o.Phase1IntegrityAlgorithms
	}
	if o.Phase1LifetimeSeconds != nil {
		toSerialize["Phase1LifetimeSeconds"] = o.Phase1LifetimeSeconds
	}
	if o.ReplayWindowSize != nil {
		toSerialize["ReplayWindowSize"] = o.ReplayWindowSize
	}
	if o.StartupAction != nil {
		toSerialize["StartupAction"] = o.StartupAction
	}
	return json.Marshal(toSerialize)
}

type NullablePhase1Options struct {
	value *Phase1Options
	isSet bool
}

func (v NullablePhase1Options) Get() *Phase1Options {
	return v.value
}

func (v *NullablePhase1Options) Set(val *Phase1Options) {
	v.value = val
	v.isSet = true
}

func (v NullablePhase1Options) IsSet() bool {
	return v.isSet
}

func (v *NullablePhase1Options) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhase1Options(val *Phase1Options) *NullablePhase1Options {
	return &NullablePhase1Options{value: val, isSet: true}
}

func (v NullablePhase1Options) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhase1Options) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
