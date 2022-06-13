/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkInternetServiceRequest struct for LinkInternetServiceRequest
type LinkInternetServiceRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Internet service you want to attach.
	InternetServiceId string `json:"InternetServiceId"`
	// The ID of the Net to which you want to attach the Internet service.
	NetId string `json:"NetId"`
}

// NewLinkInternetServiceRequest instantiates a new LinkInternetServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkInternetServiceRequest(internetServiceId string, netId string) *LinkInternetServiceRequest {
	this := LinkInternetServiceRequest{}
	this.InternetServiceId = internetServiceId
	this.NetId = netId
	return &this
}

// NewLinkInternetServiceRequestWithDefaults instantiates a new LinkInternetServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkInternetServiceRequestWithDefaults() *LinkInternetServiceRequest {
	this := LinkInternetServiceRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkInternetServiceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkInternetServiceRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkInternetServiceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkInternetServiceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetInternetServiceId returns the InternetServiceId field value
func (o *LinkInternetServiceRequest) GetInternetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InternetServiceId
}

// GetInternetServiceIdOk returns a tuple with the InternetServiceId field value
// and a boolean to check if the value has been set.
func (o *LinkInternetServiceRequest) GetInternetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternetServiceId, true
}

// SetInternetServiceId sets field value
func (o *LinkInternetServiceRequest) SetInternetServiceId(v string) {
	o.InternetServiceId = v
}

// GetNetId returns the NetId field value
func (o *LinkInternetServiceRequest) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *LinkInternetServiceRequest) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *LinkInternetServiceRequest) SetNetId(v string) {
	o.NetId = v
}

func (o LinkInternetServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["InternetServiceId"] = o.InternetServiceId
	}
	if true {
		toSerialize["NetId"] = o.NetId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkInternetServiceRequest struct {
	value *LinkInternetServiceRequest
	isSet bool
}

func (v NullableLinkInternetServiceRequest) Get() *LinkInternetServiceRequest {
	return v.value
}

func (v *NullableLinkInternetServiceRequest) Set(val *LinkInternetServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkInternetServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkInternetServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkInternetServiceRequest(val *LinkInternetServiceRequest) *NullableLinkInternetServiceRequest {
	return &NullableLinkInternetServiceRequest{value: val, isSet: true}
}

func (v NullableLinkInternetServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkInternetServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
