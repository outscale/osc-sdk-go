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

// Net Information about the Net.
type Net struct {
	// The ID of the DHCP options set (or `default` if you want to associate the default one).
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty"`
	// The IP range for the Net, in CIDR notation (for example, `10.0.0.0/16`).
	IpRange *string `json:"IpRange,omitempty"`
	// The ID of the Net.
	NetId *string `json:"NetId,omitempty"`
	// The state of the Net (`pending` \\| `available` \\| `deleting`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the Net.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The VM tenancy in a Net.
	Tenancy *string `json:"Tenancy,omitempty"`
}

// NewNet instantiates a new Net object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNet() *Net {
	this := Net{}
	return &this
}

// NewNetWithDefaults instantiates a new Net object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetWithDefaults() *Net {
	this := Net{}
	return &this
}

// GetDhcpOptionsSetId returns the DhcpOptionsSetId field value if set, zero value otherwise.
func (o *Net) GetDhcpOptionsSetId() string {
	if o == nil || o.DhcpOptionsSetId == nil {
		var ret string
		return ret
	}
	return *o.DhcpOptionsSetId
}

// GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Net) GetDhcpOptionsSetIdOk() (*string, bool) {
	if o == nil || o.DhcpOptionsSetId == nil {
		return nil, false
	}
	return o.DhcpOptionsSetId, true
}

// HasDhcpOptionsSetId returns a boolean if a field has been set.
func (o *Net) HasDhcpOptionsSetId() bool {
	if o != nil && o.DhcpOptionsSetId != nil {
		return true
	}

	return false
}

// SetDhcpOptionsSetId gets a reference to the given string and assigns it to the DhcpOptionsSetId field.
func (o *Net) SetDhcpOptionsSetId(v string) {
	o.DhcpOptionsSetId = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *Net) GetIpRange() string {
	if o == nil || o.IpRange == nil {
		var ret string
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Net) GetIpRangeOk() (*string, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *Net) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *Net) SetIpRange(v string) {
	o.IpRange = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *Net) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Net) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *Net) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *Net) SetNetId(v string) {
	o.NetId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Net) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Net) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Net) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Net) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Net) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Net) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Net) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *Net) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise.
func (o *Net) GetTenancy() string {
	if o == nil || o.Tenancy == nil {
		var ret string
		return ret
	}
	return *o.Tenancy
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Net) GetTenancyOk() (*string, bool) {
	if o == nil || o.Tenancy == nil {
		return nil, false
	}
	return o.Tenancy, true
}

// HasTenancy returns a boolean if a field has been set.
func (o *Net) HasTenancy() bool {
	if o != nil && o.Tenancy != nil {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given string and assigns it to the Tenancy field.
func (o *Net) SetTenancy(v string) {
	o.Tenancy = &v
}

func (o Net) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DhcpOptionsSetId != nil {
		toSerialize["DhcpOptionsSetId"] = o.DhcpOptionsSetId
	}
	if o.IpRange != nil {
		toSerialize["IpRange"] = o.IpRange
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.Tenancy != nil {
		toSerialize["Tenancy"] = o.Tenancy
	}
	return json.Marshal(toSerialize)
}

type NullableNet struct {
	value *Net
	isSet bool
}

func (v NullableNet) Get() *Net {
	return v.value
}

func (v *NullableNet) Set(val *Net) {
	v.value = val
	v.isSet = true
}

func (v NullableNet) IsSet() bool {
	return v.isSet
}

func (v *NullableNet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNet(val *Net) *NullableNet {
	return &NullableNet{value: val, isSet: true}
}

func (v NullableNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
