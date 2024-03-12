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

// FiltersCatalogs One or more filters.
type FiltersCatalogs struct {
	// By default or if set to true, only returns the current catalog. If false, returns the current catalog and past catalogs.
	CurrentCatalogOnly *bool `json:"CurrentCatalogOnly,omitempty"`
	// The beginning of the time period, in ISO 8601 date format (for example, `2020-06-14`). This date cannot be older than 3 years. You must specify the parameters `FromDate` and `ToDate` together.
	FromDate *string `json:"FromDate,omitempty"`
	// The end of the time period, in ISO 8601 date format (for example, `2020-06-30`). You must specify the parameters `FromDate` and `ToDate` together.
	ToDate *string `json:"ToDate,omitempty"`
}

// NewFiltersCatalogs instantiates a new FiltersCatalogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersCatalogs() *FiltersCatalogs {
	this := FiltersCatalogs{}
	return &this
}

// NewFiltersCatalogsWithDefaults instantiates a new FiltersCatalogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersCatalogsWithDefaults() *FiltersCatalogs {
	this := FiltersCatalogs{}
	return &this
}

// GetCurrentCatalogOnly returns the CurrentCatalogOnly field value if set, zero value otherwise.
func (o *FiltersCatalogs) GetCurrentCatalogOnly() bool {
	if o == nil || o.CurrentCatalogOnly == nil {
		var ret bool
		return ret
	}
	return *o.CurrentCatalogOnly
}

// GetCurrentCatalogOnlyOk returns a tuple with the CurrentCatalogOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersCatalogs) GetCurrentCatalogOnlyOk() (*bool, bool) {
	if o == nil || o.CurrentCatalogOnly == nil {
		return nil, false
	}
	return o.CurrentCatalogOnly, true
}

// HasCurrentCatalogOnly returns a boolean if a field has been set.
func (o *FiltersCatalogs) HasCurrentCatalogOnly() bool {
	if o != nil && o.CurrentCatalogOnly != nil {
		return true
	}

	return false
}

// SetCurrentCatalogOnly gets a reference to the given bool and assigns it to the CurrentCatalogOnly field.
func (o *FiltersCatalogs) SetCurrentCatalogOnly(v bool) {
	o.CurrentCatalogOnly = &v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *FiltersCatalogs) GetFromDate() string {
	if o == nil || o.FromDate == nil {
		var ret string
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersCatalogs) GetFromDateOk() (*string, bool) {
	if o == nil || o.FromDate == nil {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *FiltersCatalogs) HasFromDate() bool {
	if o != nil && o.FromDate != nil {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given string and assigns it to the FromDate field.
func (o *FiltersCatalogs) SetFromDate(v string) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *FiltersCatalogs) GetToDate() string {
	if o == nil || o.ToDate == nil {
		var ret string
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersCatalogs) GetToDateOk() (*string, bool) {
	if o == nil || o.ToDate == nil {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *FiltersCatalogs) HasToDate() bool {
	if o != nil && o.ToDate != nil {
		return true
	}

	return false
}

// SetToDate gets a reference to the given string and assigns it to the ToDate field.
func (o *FiltersCatalogs) SetToDate(v string) {
	o.ToDate = &v
}

func (o FiltersCatalogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentCatalogOnly != nil {
		toSerialize["CurrentCatalogOnly"] = o.CurrentCatalogOnly
	}
	if o.FromDate != nil {
		toSerialize["FromDate"] = o.FromDate
	}
	if o.ToDate != nil {
		toSerialize["ToDate"] = o.ToDate
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersCatalogs struct {
	value *FiltersCatalogs
	isSet bool
}

func (v NullableFiltersCatalogs) Get() *FiltersCatalogs {
	return v.value
}

func (v *NullableFiltersCatalogs) Set(val *FiltersCatalogs) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersCatalogs) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersCatalogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersCatalogs(val *FiltersCatalogs) *NullableFiltersCatalogs {
	return &NullableFiltersCatalogs{value: val, isSet: true}
}

func (v NullableFiltersCatalogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersCatalogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
