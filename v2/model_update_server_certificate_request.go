/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateServerCertificateRequest struct for UpdateServerCertificateRequest
type UpdateServerCertificateRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The name of the server certificate you want to modify.
	Name string `json:"Name"`
	// A new name for the server certificate.
	NewName *string `json:"NewName,omitempty"`
	// A new path for the server certificate.
	NewPath *string `json:"NewPath,omitempty"`
}

// NewUpdateServerCertificateRequest instantiates a new UpdateServerCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServerCertificateRequest(name string) *UpdateServerCertificateRequest {
	this := UpdateServerCertificateRequest{}
	this.Name = name
	return &this
}

// NewUpdateServerCertificateRequestWithDefaults instantiates a new UpdateServerCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServerCertificateRequestWithDefaults() *UpdateServerCertificateRequest {
	this := UpdateServerCertificateRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateServerCertificateRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateServerCertificateRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateServerCertificateRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetName returns the Name field value
func (o *UpdateServerCertificateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateServerCertificateRequest) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateServerCertificateRequest) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateRequest) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateServerCertificateRequest) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateServerCertificateRequest) SetNewName(v string) {
	o.NewName = &v
}

// GetNewPath returns the NewPath field value if set, zero value otherwise.
func (o *UpdateServerCertificateRequest) GetNewPath() string {
	if o == nil || o.NewPath == nil {
		var ret string
		return ret
	}
	return *o.NewPath
}

// GetNewPathOk returns a tuple with the NewPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateRequest) GetNewPathOk() (*string, bool) {
	if o == nil || o.NewPath == nil {
		return nil, false
	}
	return o.NewPath, true
}

// HasNewPath returns a boolean if a field has been set.
func (o *UpdateServerCertificateRequest) HasNewPath() bool {
	if o != nil && o.NewPath != nil {
		return true
	}

	return false
}

// SetNewPath gets a reference to the given string and assigns it to the NewPath field.
func (o *UpdateServerCertificateRequest) SetNewPath(v string) {
	o.NewPath = &v
}

func (o UpdateServerCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["NewName"] = o.NewName
	}
	if o.NewPath != nil {
		toSerialize["NewPath"] = o.NewPath
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateServerCertificateRequest struct {
	value *UpdateServerCertificateRequest
	isSet bool
}

func (v NullableUpdateServerCertificateRequest) Get() *UpdateServerCertificateRequest {
	return v.value
}

func (v *NullableUpdateServerCertificateRequest) Set(val *UpdateServerCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServerCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServerCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServerCertificateRequest(val *UpdateServerCertificateRequest) *NullableUpdateServerCertificateRequest {
	return &NullableUpdateServerCertificateRequest{value: val, isSet: true}
}

func (v NullableUpdateServerCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServerCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
