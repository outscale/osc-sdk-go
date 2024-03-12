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

// ReadVirtualGatewaysRequest struct for ReadVirtualGatewaysRequest
type ReadVirtualGatewaysRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool                  `json:"DryRun,omitempty"`
	Filters *FiltersVirtualGateway `json:"Filters,omitempty"`
	// The token to request the next page of results. Each token refers to a specific page.
	NextPageToken *string `json:"NextPageToken,omitempty"`
	// The maximum number of logs returned in a single response (between `1`and `1000`, both included). By default, `100`.
	ResultsPerPage *int32 `json:"ResultsPerPage,omitempty"`
}

// NewReadVirtualGatewaysRequest instantiates a new ReadVirtualGatewaysRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVirtualGatewaysRequest() *ReadVirtualGatewaysRequest {
	this := ReadVirtualGatewaysRequest{}
	return &this
}

// NewReadVirtualGatewaysRequestWithDefaults instantiates a new ReadVirtualGatewaysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVirtualGatewaysRequestWithDefaults() *ReadVirtualGatewaysRequest {
	this := ReadVirtualGatewaysRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadVirtualGatewaysRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVirtualGatewaysRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadVirtualGatewaysRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadVirtualGatewaysRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadVirtualGatewaysRequest) GetFilters() FiltersVirtualGateway {
	if o == nil || o.Filters == nil {
		var ret FiltersVirtualGateway
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVirtualGatewaysRequest) GetFiltersOk() (*FiltersVirtualGateway, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadVirtualGatewaysRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersVirtualGateway and assigns it to the Filters field.
func (o *ReadVirtualGatewaysRequest) SetFilters(v FiltersVirtualGateway) {
	o.Filters = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ReadVirtualGatewaysRequest) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVirtualGatewaysRequest) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ReadVirtualGatewaysRequest) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ReadVirtualGatewaysRequest) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetResultsPerPage returns the ResultsPerPage field value if set, zero value otherwise.
func (o *ReadVirtualGatewaysRequest) GetResultsPerPage() int32 {
	if o == nil || o.ResultsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ResultsPerPage
}

// GetResultsPerPageOk returns a tuple with the ResultsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVirtualGatewaysRequest) GetResultsPerPageOk() (*int32, bool) {
	if o == nil || o.ResultsPerPage == nil {
		return nil, false
	}
	return o.ResultsPerPage, true
}

// HasResultsPerPage returns a boolean if a field has been set.
func (o *ReadVirtualGatewaysRequest) HasResultsPerPage() bool {
	if o != nil && o.ResultsPerPage != nil {
		return true
	}

	return false
}

// SetResultsPerPage gets a reference to the given int32 and assigns it to the ResultsPerPage field.
func (o *ReadVirtualGatewaysRequest) SetResultsPerPage(v int32) {
	o.ResultsPerPage = &v
}

func (o ReadVirtualGatewaysRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	if o.NextPageToken != nil {
		toSerialize["NextPageToken"] = o.NextPageToken
	}
	if o.ResultsPerPage != nil {
		toSerialize["ResultsPerPage"] = o.ResultsPerPage
	}
	return json.Marshal(toSerialize)
}

type NullableReadVirtualGatewaysRequest struct {
	value *ReadVirtualGatewaysRequest
	isSet bool
}

func (v NullableReadVirtualGatewaysRequest) Get() *ReadVirtualGatewaysRequest {
	return v.value
}

func (v *NullableReadVirtualGatewaysRequest) Set(val *ReadVirtualGatewaysRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVirtualGatewaysRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVirtualGatewaysRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVirtualGatewaysRequest(val *ReadVirtualGatewaysRequest) *NullableReadVirtualGatewaysRequest {
	return &NullableReadVirtualGatewaysRequest{value: val, isSet: true}
}

func (v NullableReadVirtualGatewaysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVirtualGatewaysRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
