/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkPublicIp Information about the public IP association.
type LinkPublicIp struct {
	// (Required in a Net) The ID representing the association of the public IP with the VM or the NIC.
	LinkPublicIpId *string `json:"LinkPublicIpId,omitempty"`
	// The name of the public DNS.
	PublicDnsName *string `json:"PublicDnsName,omitempty"`
	// The public IP associated with the NIC.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The account ID of the owner of the public IP.
	PublicIpAccountId *string `json:"PublicIpAccountId,omitempty"`
	// The allocation ID of the public IP.
	PublicIpId *string `json:"PublicIpId,omitempty"`
}

// NewLinkPublicIp instantiates a new LinkPublicIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkPublicIp() *LinkPublicIp {
	this := LinkPublicIp{}
	return &this
}

// NewLinkPublicIpWithDefaults instantiates a new LinkPublicIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkPublicIpWithDefaults() *LinkPublicIp {
	this := LinkPublicIp{}
	return &this
}

// GetLinkPublicIpId returns the LinkPublicIpId field value if set, zero value otherwise.
func (o *LinkPublicIp) GetLinkPublicIpId() string {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret
	}
	return *o.LinkPublicIpId
}

// GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetLinkPublicIpIdOk() (*string, bool) {
	if o == nil || o.LinkPublicIpId == nil {
		return nil, false
	}
	return o.LinkPublicIpId, true
}

// HasLinkPublicIpId returns a boolean if a field has been set.
func (o *LinkPublicIp) HasLinkPublicIpId() bool {
	if o != nil && o.LinkPublicIpId != nil {
		return true
	}

	return false
}

// SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.
func (o *LinkPublicIp) SetLinkPublicIpId(v string) {
	o.LinkPublicIpId = &v
}

// GetPublicDnsName returns the PublicDnsName field value if set, zero value otherwise.
func (o *LinkPublicIp) GetPublicDnsName() string {
	if o == nil || o.PublicDnsName == nil {
		var ret string
		return ret
	}
	return *o.PublicDnsName
}

// GetPublicDnsNameOk returns a tuple with the PublicDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetPublicDnsNameOk() (*string, bool) {
	if o == nil || o.PublicDnsName == nil {
		return nil, false
	}
	return o.PublicDnsName, true
}

// HasPublicDnsName returns a boolean if a field has been set.
func (o *LinkPublicIp) HasPublicDnsName() bool {
	if o != nil && o.PublicDnsName != nil {
		return true
	}

	return false
}

// SetPublicDnsName gets a reference to the given string and assigns it to the PublicDnsName field.
func (o *LinkPublicIp) SetPublicDnsName(v string) {
	o.PublicDnsName = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *LinkPublicIp) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *LinkPublicIp) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *LinkPublicIp) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicIpAccountId returns the PublicIpAccountId field value if set, zero value otherwise.
func (o *LinkPublicIp) GetPublicIpAccountId() string {
	if o == nil || o.PublicIpAccountId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpAccountId
}

// GetPublicIpAccountIdOk returns a tuple with the PublicIpAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetPublicIpAccountIdOk() (*string, bool) {
	if o == nil || o.PublicIpAccountId == nil {
		return nil, false
	}
	return o.PublicIpAccountId, true
}

// HasPublicIpAccountId returns a boolean if a field has been set.
func (o *LinkPublicIp) HasPublicIpAccountId() bool {
	if o != nil && o.PublicIpAccountId != nil {
		return true
	}

	return false
}

// SetPublicIpAccountId gets a reference to the given string and assigns it to the PublicIpAccountId field.
func (o *LinkPublicIp) SetPublicIpAccountId(v string) {
	o.PublicIpAccountId = &v
}

// GetPublicIpId returns the PublicIpId field value if set, zero value otherwise.
func (o *LinkPublicIp) GetPublicIpId() string {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpId
}

// GetPublicIpIdOk returns a tuple with the PublicIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetPublicIpIdOk() (*string, bool) {
	if o == nil || o.PublicIpId == nil {
		return nil, false
	}
	return o.PublicIpId, true
}

// HasPublicIpId returns a boolean if a field has been set.
func (o *LinkPublicIp) HasPublicIpId() bool {
	if o != nil && o.PublicIpId != nil {
		return true
	}

	return false
}

// SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.
func (o *LinkPublicIp) SetPublicIpId(v string) {
	o.PublicIpId = &v
}

func (o LinkPublicIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkPublicIpId != nil {
		toSerialize["LinkPublicIpId"] = o.LinkPublicIpId
	}
	if o.PublicDnsName != nil {
		toSerialize["PublicDnsName"] = o.PublicDnsName
	}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.PublicIpAccountId != nil {
		toSerialize["PublicIpAccountId"] = o.PublicIpAccountId
	}
	if o.PublicIpId != nil {
		toSerialize["PublicIpId"] = o.PublicIpId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkPublicIp struct {
	value *LinkPublicIp
	isSet bool
}

func (v NullableLinkPublicIp) Get() *LinkPublicIp {
	return v.value
}

func (v *NullableLinkPublicIp) Set(val *LinkPublicIp) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkPublicIp) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkPublicIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkPublicIp(val *LinkPublicIp) *NullableLinkPublicIp {
	return &NullableLinkPublicIp{value: val, isSet: true}
}

func (v NullableLinkPublicIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkPublicIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
