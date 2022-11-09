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

// ReadVpnConnectionsResponse struct for ReadVpnConnectionsResponse
type ReadVpnConnectionsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more VPN connections.
	VpnConnections *[]VpnConnection `json:"VpnConnections,omitempty"`
}

// NewReadVpnConnectionsResponse instantiates a new ReadVpnConnectionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVpnConnectionsResponse() *ReadVpnConnectionsResponse {
	this := ReadVpnConnectionsResponse{}
	return &this
}

// NewReadVpnConnectionsResponseWithDefaults instantiates a new ReadVpnConnectionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVpnConnectionsResponseWithDefaults() *ReadVpnConnectionsResponse {
	this := ReadVpnConnectionsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVpnConnectionsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVpnConnectionsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVpnConnectionsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVpnConnectionsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVpnConnections returns the VpnConnections field value if set, zero value otherwise.
func (o *ReadVpnConnectionsResponse) GetVpnConnections() []VpnConnection {
	if o == nil || o.VpnConnections == nil {
		var ret []VpnConnection
		return ret
	}
	return *o.VpnConnections
}

// GetVpnConnectionsOk returns a tuple with the VpnConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVpnConnectionsResponse) GetVpnConnectionsOk() (*[]VpnConnection, bool) {
	if o == nil || o.VpnConnections == nil {
		return nil, false
	}
	return o.VpnConnections, true
}

// HasVpnConnections returns a boolean if a field has been set.
func (o *ReadVpnConnectionsResponse) HasVpnConnections() bool {
	if o != nil && o.VpnConnections != nil {
		return true
	}

	return false
}

// SetVpnConnections gets a reference to the given []VpnConnection and assigns it to the VpnConnections field.
func (o *ReadVpnConnectionsResponse) SetVpnConnections(v []VpnConnection) {
	o.VpnConnections = &v
}

func (o ReadVpnConnectionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.VpnConnections != nil {
		toSerialize["VpnConnections"] = o.VpnConnections
	}
	return json.Marshal(toSerialize)
}

type NullableReadVpnConnectionsResponse struct {
	value *ReadVpnConnectionsResponse
	isSet bool
}

func (v NullableReadVpnConnectionsResponse) Get() *ReadVpnConnectionsResponse {
	return v.value
}

func (v *NullableReadVpnConnectionsResponse) Set(val *ReadVpnConnectionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVpnConnectionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVpnConnectionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVpnConnectionsResponse(val *ReadVpnConnectionsResponse) *NullableReadVpnConnectionsResponse {
	return &NullableReadVpnConnectionsResponse{value: val, isSet: true}
}

func (v NullableReadVpnConnectionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVpnConnectionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
