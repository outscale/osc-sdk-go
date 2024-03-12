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

// DhcpOptionsSet Information about the DHCP options set.
type DhcpOptionsSet struct {
	// If true, the DHCP options set is a default one. If false, it is not.
	Default *bool `json:"Default,omitempty"`
	// The ID of the DHCP options set.
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty"`
	// The domain name.
	DomainName *string `json:"DomainName,omitempty"`
	// One or more IPs for the domain name servers.
	DomainNameServers *[]string `json:"DomainNameServers,omitempty"`
	// One or more IPs for the log servers.
	LogServers *[]string `json:"LogServers,omitempty"`
	// One or more IPs for the NTP servers.
	NtpServers *[]string `json:"NtpServers,omitempty"`
	// One or more tags associated with the DHCP options set.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewDhcpOptionsSet instantiates a new DhcpOptionsSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpOptionsSet() *DhcpOptionsSet {
	this := DhcpOptionsSet{}
	return &this
}

// NewDhcpOptionsSetWithDefaults instantiates a new DhcpOptionsSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpOptionsSetWithDefaults() *DhcpOptionsSet {
	this := DhcpOptionsSet{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *DhcpOptionsSet) SetDefault(v bool) {
	o.Default = &v
}

// GetDhcpOptionsSetId returns the DhcpOptionsSetId field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDhcpOptionsSetId() string {
	if o == nil || o.DhcpOptionsSetId == nil {
		var ret string
		return ret
	}
	return *o.DhcpOptionsSetId
}

// GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDhcpOptionsSetIdOk() (*string, bool) {
	if o == nil || o.DhcpOptionsSetId == nil {
		return nil, false
	}
	return o.DhcpOptionsSetId, true
}

// HasDhcpOptionsSetId returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDhcpOptionsSetId() bool {
	if o != nil && o.DhcpOptionsSetId != nil {
		return true
	}

	return false
}

// SetDhcpOptionsSetId gets a reference to the given string and assigns it to the DhcpOptionsSetId field.
func (o *DhcpOptionsSet) SetDhcpOptionsSetId(v string) {
	o.DhcpOptionsSetId = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *DhcpOptionsSet) SetDomainName(v string) {
	o.DomainName = &v
}

// GetDomainNameServers returns the DomainNameServers field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDomainNameServers() []string {
	if o == nil || o.DomainNameServers == nil {
		var ret []string
		return ret
	}
	return *o.DomainNameServers
}

// GetDomainNameServersOk returns a tuple with the DomainNameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDomainNameServersOk() (*[]string, bool) {
	if o == nil || o.DomainNameServers == nil {
		return nil, false
	}
	return o.DomainNameServers, true
}

// HasDomainNameServers returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDomainNameServers() bool {
	if o != nil && o.DomainNameServers != nil {
		return true
	}

	return false
}

// SetDomainNameServers gets a reference to the given []string and assigns it to the DomainNameServers field.
func (o *DhcpOptionsSet) SetDomainNameServers(v []string) {
	o.DomainNameServers = &v
}

// GetLogServers returns the LogServers field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetLogServers() []string {
	if o == nil || o.LogServers == nil {
		var ret []string
		return ret
	}
	return *o.LogServers
}

// GetLogServersOk returns a tuple with the LogServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetLogServersOk() (*[]string, bool) {
	if o == nil || o.LogServers == nil {
		return nil, false
	}
	return o.LogServers, true
}

// HasLogServers returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasLogServers() bool {
	if o != nil && o.LogServers != nil {
		return true
	}

	return false
}

// SetLogServers gets a reference to the given []string and assigns it to the LogServers field.
func (o *DhcpOptionsSet) SetLogServers(v []string) {
	o.LogServers = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetNtpServers() []string {
	if o == nil || o.NtpServers == nil {
		var ret []string
		return ret
	}
	return *o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *DhcpOptionsSet) SetNtpServers(v []string) {
	o.NtpServers = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *DhcpOptionsSet) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o DhcpOptionsSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil {
		toSerialize["Default"] = o.Default
	}
	if o.DhcpOptionsSetId != nil {
		toSerialize["DhcpOptionsSetId"] = o.DhcpOptionsSetId
	}
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.DomainNameServers != nil {
		toSerialize["DomainNameServers"] = o.DomainNameServers
	}
	if o.LogServers != nil {
		toSerialize["LogServers"] = o.LogServers
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpOptionsSet struct {
	value *DhcpOptionsSet
	isSet bool
}

func (v NullableDhcpOptionsSet) Get() *DhcpOptionsSet {
	return v.value
}

func (v *NullableDhcpOptionsSet) Set(val *DhcpOptionsSet) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpOptionsSet) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpOptionsSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpOptionsSet(val *DhcpOptionsSet) *NullableDhcpOptionsSet {
	return &NullableDhcpOptionsSet{value: val, isSet: true}
}

func (v NullableDhcpOptionsSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpOptionsSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
