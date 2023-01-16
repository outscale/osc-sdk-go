/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateAccessKeyRequest struct for CreateAccessKeyRequest
type CreateAccessKeyRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The date and time at which you want the access key to expire, in ISO 8601 format (for example, `2017-06-14` or `2017-06-14T00:00:00Z`). To remove an existing expiration date, use the method without specifying this parameter.
	ExpirationDate *string `json:"ExpirationDate,omitempty"`
}

// NewCreateAccessKeyRequest instantiates a new CreateAccessKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccessKeyRequest() *CreateAccessKeyRequest {
	this := CreateAccessKeyRequest{}
	return &this
}

// NewCreateAccessKeyRequestWithDefaults instantiates a new CreateAccessKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccessKeyRequestWithDefaults() *CreateAccessKeyRequest {
	this := CreateAccessKeyRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateAccessKeyRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessKeyRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateAccessKeyRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateAccessKeyRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CreateAccessKeyRequest) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessKeyRequest) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CreateAccessKeyRequest) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *CreateAccessKeyRequest) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

func (o CreateAccessKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.ExpirationDate != nil {
		toSerialize["ExpirationDate"] = o.ExpirationDate
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAccessKeyRequest struct {
	value *CreateAccessKeyRequest
	isSet bool
}

func (v NullableCreateAccessKeyRequest) Get() *CreateAccessKeyRequest {
	return v.value
}

func (v *NullableCreateAccessKeyRequest) Set(val *CreateAccessKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccessKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccessKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccessKeyRequest(val *CreateAccessKeyRequest) *NullableCreateAccessKeyRequest {
	return &NullableCreateAccessKeyRequest{value: val, isSet: true}
}

func (v NullableCreateAccessKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccessKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
