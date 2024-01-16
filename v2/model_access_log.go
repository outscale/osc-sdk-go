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

// AccessLog Information about access logs.
type AccessLog struct {
	// If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the `OsuBucketName` parameter is required.
	IsEnabled *bool `json:"IsEnabled,omitempty"`
	// The name of the OOS bucket for the access logs.
	OsuBucketName *string `json:"OsuBucketName,omitempty"`
	// The path to the folder of the access logs in your OOS bucket (by default, the `root` level of your bucket).
	OsuBucketPrefix *string `json:"OsuBucketPrefix,omitempty"`
	// The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either `5` or `60` (by default, `60`).
	PublicationInterval *int32 `json:"PublicationInterval,omitempty"`
}

// NewAccessLog instantiates a new AccessLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessLog() *AccessLog {
	this := AccessLog{}
	return &this
}

// NewAccessLogWithDefaults instantiates a new AccessLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessLogWithDefaults() *AccessLog {
	this := AccessLog{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *AccessLog) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessLog) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *AccessLog) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *AccessLog) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetOsuBucketName returns the OsuBucketName field value if set, zero value otherwise.
func (o *AccessLog) GetOsuBucketName() string {
	if o == nil || o.OsuBucketName == nil {
		var ret string
		return ret
	}
	return *o.OsuBucketName
}

// GetOsuBucketNameOk returns a tuple with the OsuBucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessLog) GetOsuBucketNameOk() (*string, bool) {
	if o == nil || o.OsuBucketName == nil {
		return nil, false
	}
	return o.OsuBucketName, true
}

// HasOsuBucketName returns a boolean if a field has been set.
func (o *AccessLog) HasOsuBucketName() bool {
	if o != nil && o.OsuBucketName != nil {
		return true
	}

	return false
}

// SetOsuBucketName gets a reference to the given string and assigns it to the OsuBucketName field.
func (o *AccessLog) SetOsuBucketName(v string) {
	o.OsuBucketName = &v
}

// GetOsuBucketPrefix returns the OsuBucketPrefix field value if set, zero value otherwise.
func (o *AccessLog) GetOsuBucketPrefix() string {
	if o == nil || o.OsuBucketPrefix == nil {
		var ret string
		return ret
	}
	return *o.OsuBucketPrefix
}

// GetOsuBucketPrefixOk returns a tuple with the OsuBucketPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessLog) GetOsuBucketPrefixOk() (*string, bool) {
	if o == nil || o.OsuBucketPrefix == nil {
		return nil, false
	}
	return o.OsuBucketPrefix, true
}

// HasOsuBucketPrefix returns a boolean if a field has been set.
func (o *AccessLog) HasOsuBucketPrefix() bool {
	if o != nil && o.OsuBucketPrefix != nil {
		return true
	}

	return false
}

// SetOsuBucketPrefix gets a reference to the given string and assigns it to the OsuBucketPrefix field.
func (o *AccessLog) SetOsuBucketPrefix(v string) {
	o.OsuBucketPrefix = &v
}

// GetPublicationInterval returns the PublicationInterval field value if set, zero value otherwise.
func (o *AccessLog) GetPublicationInterval() int32 {
	if o == nil || o.PublicationInterval == nil {
		var ret int32
		return ret
	}
	return *o.PublicationInterval
}

// GetPublicationIntervalOk returns a tuple with the PublicationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessLog) GetPublicationIntervalOk() (*int32, bool) {
	if o == nil || o.PublicationInterval == nil {
		return nil, false
	}
	return o.PublicationInterval, true
}

// HasPublicationInterval returns a boolean if a field has been set.
func (o *AccessLog) HasPublicationInterval() bool {
	if o != nil && o.PublicationInterval != nil {
		return true
	}

	return false
}

// SetPublicationInterval gets a reference to the given int32 and assigns it to the PublicationInterval field.
func (o *AccessLog) SetPublicationInterval(v int32) {
	o.PublicationInterval = &v
}

func (o AccessLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled != nil {
		toSerialize["IsEnabled"] = o.IsEnabled
	}
	if o.OsuBucketName != nil {
		toSerialize["OsuBucketName"] = o.OsuBucketName
	}
	if o.OsuBucketPrefix != nil {
		toSerialize["OsuBucketPrefix"] = o.OsuBucketPrefix
	}
	if o.PublicationInterval != nil {
		toSerialize["PublicationInterval"] = o.PublicationInterval
	}
	return json.Marshal(toSerialize)
}

type NullableAccessLog struct {
	value *AccessLog
	isSet bool
}

func (v NullableAccessLog) Get() *AccessLog {
	return v.value
}

func (v *NullableAccessLog) Set(val *AccessLog) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessLog) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessLog(val *AccessLog) *NullableAccessLog {
	return &NullableAccessLog{value: val, isSet: true}
}

func (v NullableAccessLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
