/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateNetRequest struct for CreateNetRequest
type CreateNetRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The IP range for the Net, in CIDR notation (for example, `10.0.0.0/16`).
	IpRange string `json:"IpRange"`
	// The tenancy options for the VMs (`default` if a VM created in a Net can be launched with any tenancy, `dedicated` if it can be launched with dedicated tenancy VMs running on single-tenant hardware).
	Tenancy *string `json:"Tenancy,omitempty"`
}

// NewCreateNetRequest instantiates a new CreateNetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetRequest(ipRange string) *CreateNetRequest {
	this := CreateNetRequest{}
	this.IpRange = ipRange
	return &this
}

// NewCreateNetRequestWithDefaults instantiates a new CreateNetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetRequestWithDefaults() *CreateNetRequest {
	this := CreateNetRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateNetRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateNetRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateNetRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetIpRange returns the IpRange field value
func (o *CreateNetRequest) GetIpRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value
// and a boolean to check if the value has been set.
func (o *CreateNetRequest) GetIpRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpRange, true
}

// SetIpRange sets field value
func (o *CreateNetRequest) SetIpRange(v string) {
	o.IpRange = v
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise.
func (o *CreateNetRequest) GetTenancy() string {
	if o == nil || o.Tenancy == nil {
		var ret string
		return ret
	}
	return *o.Tenancy
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetRequest) GetTenancyOk() (*string, bool) {
	if o == nil || o.Tenancy == nil {
		return nil, false
	}
	return o.Tenancy, true
}

// HasTenancy returns a boolean if a field has been set.
func (o *CreateNetRequest) HasTenancy() bool {
	if o != nil && o.Tenancy != nil {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given string and assigns it to the Tenancy field.
func (o *CreateNetRequest) SetTenancy(v string) {
	o.Tenancy = &v
}

func (o CreateNetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["IpRange"] = o.IpRange
	}
	if o.Tenancy != nil {
		toSerialize["Tenancy"] = o.Tenancy
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetRequest struct {
	value *CreateNetRequest
	isSet bool
}

func (v NullableCreateNetRequest) Get() *CreateNetRequest {
	return v.value
}

func (v *NullableCreateNetRequest) Set(val *CreateNetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetRequest(val *CreateNetRequest) *NullableCreateNetRequest {
	return &NullableCreateNetRequest{value: val, isSet: true}
}

func (v NullableCreateNetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
