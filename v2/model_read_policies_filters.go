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

// ReadPoliciesFilters One or more filters.
type ReadPoliciesFilters struct {
	// If set to true, lists only the policies attached to a user.
	OnlyLinked *bool `json:"OnlyLinked,omitempty"`
	// The path prefix you can use to filter the results, set to a slash (`/`) by default.
	PathPrefix *string `json:"PathPrefix,omitempty"`
	// The scope to filter policies (`OWS` \\| `LOCAL`).
	Scope *string `json:"Scope,omitempty"`
}

// NewReadPoliciesFilters instantiates a new ReadPoliciesFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadPoliciesFilters() *ReadPoliciesFilters {
	this := ReadPoliciesFilters{}
	return &this
}

// NewReadPoliciesFiltersWithDefaults instantiates a new ReadPoliciesFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadPoliciesFiltersWithDefaults() *ReadPoliciesFilters {
	this := ReadPoliciesFilters{}
	return &this
}

// GetOnlyLinked returns the OnlyLinked field value if set, zero value otherwise.
func (o *ReadPoliciesFilters) GetOnlyLinked() bool {
	if o == nil || o.OnlyLinked == nil {
		var ret bool
		return ret
	}
	return *o.OnlyLinked
}

// GetOnlyLinkedOk returns a tuple with the OnlyLinked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadPoliciesFilters) GetOnlyLinkedOk() (*bool, bool) {
	if o == nil || o.OnlyLinked == nil {
		return nil, false
	}
	return o.OnlyLinked, true
}

// HasOnlyLinked returns a boolean if a field has been set.
func (o *ReadPoliciesFilters) HasOnlyLinked() bool {
	if o != nil && o.OnlyLinked != nil {
		return true
	}

	return false
}

// SetOnlyLinked gets a reference to the given bool and assigns it to the OnlyLinked field.
func (o *ReadPoliciesFilters) SetOnlyLinked(v bool) {
	o.OnlyLinked = &v
}

// GetPathPrefix returns the PathPrefix field value if set, zero value otherwise.
func (o *ReadPoliciesFilters) GetPathPrefix() string {
	if o == nil || o.PathPrefix == nil {
		var ret string
		return ret
	}
	return *o.PathPrefix
}

// GetPathPrefixOk returns a tuple with the PathPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadPoliciesFilters) GetPathPrefixOk() (*string, bool) {
	if o == nil || o.PathPrefix == nil {
		return nil, false
	}
	return o.PathPrefix, true
}

// HasPathPrefix returns a boolean if a field has been set.
func (o *ReadPoliciesFilters) HasPathPrefix() bool {
	if o != nil && o.PathPrefix != nil {
		return true
	}

	return false
}

// SetPathPrefix gets a reference to the given string and assigns it to the PathPrefix field.
func (o *ReadPoliciesFilters) SetPathPrefix(v string) {
	o.PathPrefix = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ReadPoliciesFilters) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadPoliciesFilters) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ReadPoliciesFilters) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ReadPoliciesFilters) SetScope(v string) {
	o.Scope = &v
}

func (o ReadPoliciesFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OnlyLinked != nil {
		toSerialize["OnlyLinked"] = o.OnlyLinked
	}
	if o.PathPrefix != nil {
		toSerialize["PathPrefix"] = o.PathPrefix
	}
	if o.Scope != nil {
		toSerialize["Scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableReadPoliciesFilters struct {
	value *ReadPoliciesFilters
	isSet bool
}

func (v NullableReadPoliciesFilters) Get() *ReadPoliciesFilters {
	return v.value
}

func (v *NullableReadPoliciesFilters) Set(val *ReadPoliciesFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableReadPoliciesFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableReadPoliciesFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadPoliciesFilters(val *ReadPoliciesFilters) *NullableReadPoliciesFilters {
	return &NullableReadPoliciesFilters{value: val, isSet: true}
}

func (v NullableReadPoliciesFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadPoliciesFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
