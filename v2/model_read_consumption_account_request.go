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

// ReadConsumptionAccountRequest struct for ReadConsumptionAccountRequest
type ReadConsumptionAccountRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The beginning of the time period, in ISO 8601 date format (for example, `2020-06-14`). The date-time format is also accepted, but only with a time set to midnight (for example, `2020-06-14T00:00:00.000Z`).
	FromDate string `json:"FromDate"`
	// By default or if false, returns only the consumption of the specific account that sends this request. If true, returns either the overall consumption of your paying account and all linked accounts (if the account that sends this request is a paying account) or returns nothing (if the account that sends this request is a linked account).
	Overall *bool `json:"Overall,omitempty"`
	// If true, the response also includes the unit price of the consumed resource (`UnitPrice`) and the total price of the consumed resource during the specified time period (`Price`), in the currency of the Region's catalog.
	ShowPrice *bool `json:"ShowPrice,omitempty"`
	// The end of the time period, in ISO 8601 date format (for example, `2020-06-30`). The date-time format is also accepted, but only with a time set to midnight (for example, `2020-06-30T00:00:00.000Z`).
	ToDate string `json:"ToDate"`
}

// NewReadConsumptionAccountRequest instantiates a new ReadConsumptionAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadConsumptionAccountRequest(fromDate string, toDate string) *ReadConsumptionAccountRequest {
	this := ReadConsumptionAccountRequest{}
	this.FromDate = fromDate
	var overall bool = false
	this.Overall = &overall
	this.ToDate = toDate
	return &this
}

// NewReadConsumptionAccountRequestWithDefaults instantiates a new ReadConsumptionAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadConsumptionAccountRequestWithDefaults() *ReadConsumptionAccountRequest {
	this := ReadConsumptionAccountRequest{}
	var overall bool = false
	this.Overall = &overall
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadConsumptionAccountRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadConsumptionAccountRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadConsumptionAccountRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFromDate returns the FromDate field value
func (o *ReadConsumptionAccountRequest) GetFromDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountRequest) GetFromDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromDate, true
}

// SetFromDate sets field value
func (o *ReadConsumptionAccountRequest) SetFromDate(v string) {
	o.FromDate = v
}

// GetOverall returns the Overall field value if set, zero value otherwise.
func (o *ReadConsumptionAccountRequest) GetOverall() bool {
	if o == nil || o.Overall == nil {
		var ret bool
		return ret
	}
	return *o.Overall
}

// GetOverallOk returns a tuple with the Overall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountRequest) GetOverallOk() (*bool, bool) {
	if o == nil || o.Overall == nil {
		return nil, false
	}
	return o.Overall, true
}

// HasOverall returns a boolean if a field has been set.
func (o *ReadConsumptionAccountRequest) HasOverall() bool {
	if o != nil && o.Overall != nil {
		return true
	}

	return false
}

// SetOverall gets a reference to the given bool and assigns it to the Overall field.
func (o *ReadConsumptionAccountRequest) SetOverall(v bool) {
	o.Overall = &v
}

// GetShowPrice returns the ShowPrice field value if set, zero value otherwise.
func (o *ReadConsumptionAccountRequest) GetShowPrice() bool {
	if o == nil || o.ShowPrice == nil {
		var ret bool
		return ret
	}
	return *o.ShowPrice
}

// GetShowPriceOk returns a tuple with the ShowPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountRequest) GetShowPriceOk() (*bool, bool) {
	if o == nil || o.ShowPrice == nil {
		return nil, false
	}
	return o.ShowPrice, true
}

// HasShowPrice returns a boolean if a field has been set.
func (o *ReadConsumptionAccountRequest) HasShowPrice() bool {
	if o != nil && o.ShowPrice != nil {
		return true
	}

	return false
}

// SetShowPrice gets a reference to the given bool and assigns it to the ShowPrice field.
func (o *ReadConsumptionAccountRequest) SetShowPrice(v bool) {
	o.ShowPrice = &v
}

// GetToDate returns the ToDate field value
func (o *ReadConsumptionAccountRequest) GetToDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountRequest) GetToDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToDate, true
}

// SetToDate sets field value
func (o *ReadConsumptionAccountRequest) SetToDate(v string) {
	o.ToDate = v
}

func (o ReadConsumptionAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["FromDate"] = o.FromDate
	}
	if o.Overall != nil {
		toSerialize["Overall"] = o.Overall
	}
	if o.ShowPrice != nil {
		toSerialize["ShowPrice"] = o.ShowPrice
	}
	if true {
		toSerialize["ToDate"] = o.ToDate
	}
	return json.Marshal(toSerialize)
}

type NullableReadConsumptionAccountRequest struct {
	value *ReadConsumptionAccountRequest
	isSet bool
}

func (v NullableReadConsumptionAccountRequest) Get() *ReadConsumptionAccountRequest {
	return v.value
}

func (v *NullableReadConsumptionAccountRequest) Set(val *ReadConsumptionAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadConsumptionAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadConsumptionAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadConsumptionAccountRequest(val *ReadConsumptionAccountRequest) *NullableReadConsumptionAccountRequest {
	return &NullableReadConsumptionAccountRequest{value: val, isSet: true}
}

func (v NullableReadConsumptionAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadConsumptionAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
