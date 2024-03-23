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

// ReadLinkedPoliciesRequest struct for ReadLinkedPoliciesRequest
type ReadLinkedPoliciesRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool                      `json:"DryRun,omitempty"`
	Filters *ReadLinkedPoliciesFilters `json:"Filters,omitempty"`
	// The item starting the list of policies requested.
	FirstItem *int32 `json:"FirstItem,omitempty"`
	// The maximum number of items that can be returned in a single response (by default, 100).
	ResultsPerPage *int32 `json:"ResultsPerPage,omitempty"`
	// The name of the user the policies are linked to.
	UserName *string `json:"UserName,omitempty"`
}

// NewReadLinkedPoliciesRequest instantiates a new ReadLinkedPoliciesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadLinkedPoliciesRequest() *ReadLinkedPoliciesRequest {
	this := ReadLinkedPoliciesRequest{}
	return &this
}

// NewReadLinkedPoliciesRequestWithDefaults instantiates a new ReadLinkedPoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadLinkedPoliciesRequestWithDefaults() *ReadLinkedPoliciesRequest {
	this := ReadLinkedPoliciesRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadLinkedPoliciesRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLinkedPoliciesRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadLinkedPoliciesRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadLinkedPoliciesRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadLinkedPoliciesRequest) GetFilters() ReadLinkedPoliciesFilters {
	if o == nil || o.Filters == nil {
		var ret ReadLinkedPoliciesFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLinkedPoliciesRequest) GetFiltersOk() (*ReadLinkedPoliciesFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadLinkedPoliciesRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ReadLinkedPoliciesFilters and assigns it to the Filters field.
func (o *ReadLinkedPoliciesRequest) SetFilters(v ReadLinkedPoliciesFilters) {
	o.Filters = &v
}

// GetFirstItem returns the FirstItem field value if set, zero value otherwise.
func (o *ReadLinkedPoliciesRequest) GetFirstItem() int32 {
	if o == nil || o.FirstItem == nil {
		var ret int32
		return ret
	}
	return *o.FirstItem
}

// GetFirstItemOk returns a tuple with the FirstItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLinkedPoliciesRequest) GetFirstItemOk() (*int32, bool) {
	if o == nil || o.FirstItem == nil {
		return nil, false
	}
	return o.FirstItem, true
}

// HasFirstItem returns a boolean if a field has been set.
func (o *ReadLinkedPoliciesRequest) HasFirstItem() bool {
	if o != nil && o.FirstItem != nil {
		return true
	}

	return false
}

// SetFirstItem gets a reference to the given int32 and assigns it to the FirstItem field.
func (o *ReadLinkedPoliciesRequest) SetFirstItem(v int32) {
	o.FirstItem = &v
}

// GetResultsPerPage returns the ResultsPerPage field value if set, zero value otherwise.
func (o *ReadLinkedPoliciesRequest) GetResultsPerPage() int32 {
	if o == nil || o.ResultsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ResultsPerPage
}

// GetResultsPerPageOk returns a tuple with the ResultsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLinkedPoliciesRequest) GetResultsPerPageOk() (*int32, bool) {
	if o == nil || o.ResultsPerPage == nil {
		return nil, false
	}
	return o.ResultsPerPage, true
}

// HasResultsPerPage returns a boolean if a field has been set.
func (o *ReadLinkedPoliciesRequest) HasResultsPerPage() bool {
	if o != nil && o.ResultsPerPage != nil {
		return true
	}

	return false
}

// SetResultsPerPage gets a reference to the given int32 and assigns it to the ResultsPerPage field.
func (o *ReadLinkedPoliciesRequest) SetResultsPerPage(v int32) {
	o.ResultsPerPage = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ReadLinkedPoliciesRequest) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLinkedPoliciesRequest) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ReadLinkedPoliciesRequest) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ReadLinkedPoliciesRequest) SetUserName(v string) {
	o.UserName = &v
}

func (o ReadLinkedPoliciesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	if o.FirstItem != nil {
		toSerialize["FirstItem"] = o.FirstItem
	}
	if o.ResultsPerPage != nil {
		toSerialize["ResultsPerPage"] = o.ResultsPerPage
	}
	if o.UserName != nil {
		toSerialize["UserName"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableReadLinkedPoliciesRequest struct {
	value *ReadLinkedPoliciesRequest
	isSet bool
}

func (v NullableReadLinkedPoliciesRequest) Get() *ReadLinkedPoliciesRequest {
	return v.value
}

func (v *NullableReadLinkedPoliciesRequest) Set(val *ReadLinkedPoliciesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadLinkedPoliciesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadLinkedPoliciesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadLinkedPoliciesRequest(val *ReadLinkedPoliciesRequest) *NullableReadLinkedPoliciesRequest {
	return &NullableReadLinkedPoliciesRequest{value: val, isSet: true}
}

func (v NullableReadLinkedPoliciesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadLinkedPoliciesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}