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

// DeleteClientGatewayRequest struct for DeleteClientGatewayRequest
type DeleteClientGatewayRequest struct {
	// The ID of the client gateway you want to delete.
	ClientGatewayId string `json:"ClientGatewayId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewDeleteClientGatewayRequest instantiates a new DeleteClientGatewayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteClientGatewayRequest(clientGatewayId string) *DeleteClientGatewayRequest {
	this := DeleteClientGatewayRequest{}
	this.ClientGatewayId = clientGatewayId
	return &this
}

// NewDeleteClientGatewayRequestWithDefaults instantiates a new DeleteClientGatewayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteClientGatewayRequestWithDefaults() *DeleteClientGatewayRequest {
	this := DeleteClientGatewayRequest{}
	return &this
}

// GetClientGatewayId returns the ClientGatewayId field value
func (o *DeleteClientGatewayRequest) GetClientGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientGatewayId
}

// GetClientGatewayIdOk returns a tuple with the ClientGatewayId field value
// and a boolean to check if the value has been set.
func (o *DeleteClientGatewayRequest) GetClientGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientGatewayId, true
}

// SetClientGatewayId sets field value
func (o *DeleteClientGatewayRequest) SetClientGatewayId(v string) {
	o.ClientGatewayId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteClientGatewayRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteClientGatewayRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteClientGatewayRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteClientGatewayRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o DeleteClientGatewayRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClientGatewayId"] = o.ClientGatewayId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteClientGatewayRequest struct {
	value *DeleteClientGatewayRequest
	isSet bool
}

func (v NullableDeleteClientGatewayRequest) Get() *DeleteClientGatewayRequest {
	return v.value
}

func (v *NullableDeleteClientGatewayRequest) Set(val *DeleteClientGatewayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteClientGatewayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteClientGatewayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteClientGatewayRequest(val *DeleteClientGatewayRequest) *NullableDeleteClientGatewayRequest {
	return &NullableDeleteClientGatewayRequest{value: val, isSet: true}
}

func (v NullableDeleteClientGatewayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteClientGatewayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
