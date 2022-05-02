/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
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
	// The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).
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
