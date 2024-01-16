/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.5
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// VpnOptionsToUpdate Information about the VPN options.
type VpnOptionsToUpdate struct {
	Phase2Options *Phase2OptionsToUpdate `json:"Phase2Options,omitempty"`
	// The range of inside IPs for the tunnel. This must be a /30 CIDR block from the 169.254.254.0/24 range.
	TunnelInsideIpRange *string `json:"TunnelInsideIpRange,omitempty"`
}

// NewVpnOptionsToUpdate instantiates a new VpnOptionsToUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpnOptionsToUpdate() *VpnOptionsToUpdate {
	this := VpnOptionsToUpdate{}
	return &this
}

// NewVpnOptionsToUpdateWithDefaults instantiates a new VpnOptionsToUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpnOptionsToUpdateWithDefaults() *VpnOptionsToUpdate {
	this := VpnOptionsToUpdate{}
	return &this
}

// GetPhase2Options returns the Phase2Options field value if set, zero value otherwise.
func (o *VpnOptionsToUpdate) GetPhase2Options() Phase2OptionsToUpdate {
	if o == nil || o.Phase2Options == nil {
		var ret Phase2OptionsToUpdate
		return ret
	}
	return *o.Phase2Options
}

// GetPhase2OptionsOk returns a tuple with the Phase2Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnOptionsToUpdate) GetPhase2OptionsOk() (*Phase2OptionsToUpdate, bool) {
	if o == nil || o.Phase2Options == nil {
		return nil, false
	}
	return o.Phase2Options, true
}

// HasPhase2Options returns a boolean if a field has been set.
func (o *VpnOptionsToUpdate) HasPhase2Options() bool {
	if o != nil && o.Phase2Options != nil {
		return true
	}

	return false
}

// SetPhase2Options gets a reference to the given Phase2OptionsToUpdate and assigns it to the Phase2Options field.
func (o *VpnOptionsToUpdate) SetPhase2Options(v Phase2OptionsToUpdate) {
	o.Phase2Options = &v
}

// GetTunnelInsideIpRange returns the TunnelInsideIpRange field value if set, zero value otherwise.
func (o *VpnOptionsToUpdate) GetTunnelInsideIpRange() string {
	if o == nil || o.TunnelInsideIpRange == nil {
		var ret string
		return ret
	}
	return *o.TunnelInsideIpRange
}

// GetTunnelInsideIpRangeOk returns a tuple with the TunnelInsideIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnOptionsToUpdate) GetTunnelInsideIpRangeOk() (*string, bool) {
	if o == nil || o.TunnelInsideIpRange == nil {
		return nil, false
	}
	return o.TunnelInsideIpRange, true
}

// HasTunnelInsideIpRange returns a boolean if a field has been set.
func (o *VpnOptionsToUpdate) HasTunnelInsideIpRange() bool {
	if o != nil && o.TunnelInsideIpRange != nil {
		return true
	}

	return false
}

// SetTunnelInsideIpRange gets a reference to the given string and assigns it to the TunnelInsideIpRange field.
func (o *VpnOptionsToUpdate) SetTunnelInsideIpRange(v string) {
	o.TunnelInsideIpRange = &v
}

func (o VpnOptionsToUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Phase2Options != nil {
		toSerialize["Phase2Options"] = o.Phase2Options
	}
	if o.TunnelInsideIpRange != nil {
		toSerialize["TunnelInsideIpRange"] = o.TunnelInsideIpRange
	}
	return json.Marshal(toSerialize)
}

type NullableVpnOptionsToUpdate struct {
	value *VpnOptionsToUpdate
	isSet bool
}

func (v NullableVpnOptionsToUpdate) Get() *VpnOptionsToUpdate {
	return v.value
}

func (v *NullableVpnOptionsToUpdate) Set(val *VpnOptionsToUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnOptionsToUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnOptionsToUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnOptionsToUpdate(val *VpnOptionsToUpdate) *NullableVpnOptionsToUpdate {
	return &NullableVpnOptionsToUpdate{value: val, isSet: true}
}

func (v NullableVpnOptionsToUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnOptionsToUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
