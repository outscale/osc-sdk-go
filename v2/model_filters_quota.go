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

// FiltersQuota One or more filters.
type FiltersQuota struct {
	// The group names of the quotas.
	Collections *[]string `json:"Collections,omitempty"`
	// The names of the quotas.
	QuotaNames *[]string `json:"QuotaNames,omitempty"`
	// The resource IDs if these are resource-specific quotas, `global` if they are not.
	QuotaTypes *[]string `json:"QuotaTypes,omitempty"`
	// The description of the quotas.
	ShortDescriptions *[]string `json:"ShortDescriptions,omitempty"`
}

// NewFiltersQuota instantiates a new FiltersQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersQuota() *FiltersQuota {
	this := FiltersQuota{}
	return &this
}

// NewFiltersQuotaWithDefaults instantiates a new FiltersQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersQuotaWithDefaults() *FiltersQuota {
	this := FiltersQuota{}
	return &this
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *FiltersQuota) GetCollections() []string {
	if o == nil || o.Collections == nil {
		var ret []string
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersQuota) GetCollectionsOk() (*[]string, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *FiltersQuota) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []string and assigns it to the Collections field.
func (o *FiltersQuota) SetCollections(v []string) {
	o.Collections = &v
}

// GetQuotaNames returns the QuotaNames field value if set, zero value otherwise.
func (o *FiltersQuota) GetQuotaNames() []string {
	if o == nil || o.QuotaNames == nil {
		var ret []string
		return ret
	}
	return *o.QuotaNames
}

// GetQuotaNamesOk returns a tuple with the QuotaNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersQuota) GetQuotaNamesOk() (*[]string, bool) {
	if o == nil || o.QuotaNames == nil {
		return nil, false
	}
	return o.QuotaNames, true
}

// HasQuotaNames returns a boolean if a field has been set.
func (o *FiltersQuota) HasQuotaNames() bool {
	if o != nil && o.QuotaNames != nil {
		return true
	}

	return false
}

// SetQuotaNames gets a reference to the given []string and assigns it to the QuotaNames field.
func (o *FiltersQuota) SetQuotaNames(v []string) {
	o.QuotaNames = &v
}

// GetQuotaTypes returns the QuotaTypes field value if set, zero value otherwise.
func (o *FiltersQuota) GetQuotaTypes() []string {
	if o == nil || o.QuotaTypes == nil {
		var ret []string
		return ret
	}
	return *o.QuotaTypes
}

// GetQuotaTypesOk returns a tuple with the QuotaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersQuota) GetQuotaTypesOk() (*[]string, bool) {
	if o == nil || o.QuotaTypes == nil {
		return nil, false
	}
	return o.QuotaTypes, true
}

// HasQuotaTypes returns a boolean if a field has been set.
func (o *FiltersQuota) HasQuotaTypes() bool {
	if o != nil && o.QuotaTypes != nil {
		return true
	}

	return false
}

// SetQuotaTypes gets a reference to the given []string and assigns it to the QuotaTypes field.
func (o *FiltersQuota) SetQuotaTypes(v []string) {
	o.QuotaTypes = &v
}

// GetShortDescriptions returns the ShortDescriptions field value if set, zero value otherwise.
func (o *FiltersQuota) GetShortDescriptions() []string {
	if o == nil || o.ShortDescriptions == nil {
		var ret []string
		return ret
	}
	return *o.ShortDescriptions
}

// GetShortDescriptionsOk returns a tuple with the ShortDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersQuota) GetShortDescriptionsOk() (*[]string, bool) {
	if o == nil || o.ShortDescriptions == nil {
		return nil, false
	}
	return o.ShortDescriptions, true
}

// HasShortDescriptions returns a boolean if a field has been set.
func (o *FiltersQuota) HasShortDescriptions() bool {
	if o != nil && o.ShortDescriptions != nil {
		return true
	}

	return false
}

// SetShortDescriptions gets a reference to the given []string and assigns it to the ShortDescriptions field.
func (o *FiltersQuota) SetShortDescriptions(v []string) {
	o.ShortDescriptions = &v
}

func (o FiltersQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collections != nil {
		toSerialize["Collections"] = o.Collections
	}
	if o.QuotaNames != nil {
		toSerialize["QuotaNames"] = o.QuotaNames
	}
	if o.QuotaTypes != nil {
		toSerialize["QuotaTypes"] = o.QuotaTypes
	}
	if o.ShortDescriptions != nil {
		toSerialize["ShortDescriptions"] = o.ShortDescriptions
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersQuota struct {
	value *FiltersQuota
	isSet bool
}

func (v NullableFiltersQuota) Get() *FiltersQuota {
	return v.value
}

func (v *NullableFiltersQuota) Set(val *FiltersQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersQuota(val *FiltersQuota) *NullableFiltersQuota {
	return &NullableFiltersQuota{value: val, isSet: true}
}

func (v NullableFiltersQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
