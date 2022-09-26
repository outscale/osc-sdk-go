/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// Listener Information about the listener.
type Listener struct {
	// The port on which the back-end VM is listening (between `1` and `65535`, both included).
	BackendPort *int32 `json:"BackendPort,omitempty"`
	// The protocol for routing traffic to back-end VMs (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL`).
	BackendProtocol *string `json:"BackendProtocol,omitempty"`
	// The port on which the load balancer is listening (between `1` and `65535`, both included).
	LoadBalancerPort *int32 `json:"LoadBalancerPort,omitempty"`
	// The routing protocol (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL`).
	LoadBalancerProtocol *string `json:"LoadBalancerProtocol,omitempty"`
	// The names of the policies. If there are no policies enabled, the list is empty.
	PolicyNames *[]string `json:"PolicyNames,omitempty"`
	// The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).
	ServerCertificateId *string `json:"ServerCertificateId,omitempty"`
}

// NewListener instantiates a new Listener object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListener() *Listener {
	this := Listener{}
	return &this
}

// NewListenerWithDefaults instantiates a new Listener object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenerWithDefaults() *Listener {
	this := Listener{}
	return &this
}

// GetBackendPort returns the BackendPort field value if set, zero value otherwise.
func (o *Listener) GetBackendPort() int32 {
	if o == nil || o.BackendPort == nil {
		var ret int32
		return ret
	}
	return *o.BackendPort
}

// GetBackendPortOk returns a tuple with the BackendPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetBackendPortOk() (*int32, bool) {
	if o == nil || o.BackendPort == nil {
		return nil, false
	}
	return o.BackendPort, true
}

// HasBackendPort returns a boolean if a field has been set.
func (o *Listener) HasBackendPort() bool {
	if o != nil && o.BackendPort != nil {
		return true
	}

	return false
}

// SetBackendPort gets a reference to the given int32 and assigns it to the BackendPort field.
func (o *Listener) SetBackendPort(v int32) {
	o.BackendPort = &v
}

// GetBackendProtocol returns the BackendProtocol field value if set, zero value otherwise.
func (o *Listener) GetBackendProtocol() string {
	if o == nil || o.BackendProtocol == nil {
		var ret string
		return ret
	}
	return *o.BackendProtocol
}

// GetBackendProtocolOk returns a tuple with the BackendProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetBackendProtocolOk() (*string, bool) {
	if o == nil || o.BackendProtocol == nil {
		return nil, false
	}
	return o.BackendProtocol, true
}

// HasBackendProtocol returns a boolean if a field has been set.
func (o *Listener) HasBackendProtocol() bool {
	if o != nil && o.BackendProtocol != nil {
		return true
	}

	return false
}

// SetBackendProtocol gets a reference to the given string and assigns it to the BackendProtocol field.
func (o *Listener) SetBackendProtocol(v string) {
	o.BackendProtocol = &v
}

// GetLoadBalancerPort returns the LoadBalancerPort field value if set, zero value otherwise.
func (o *Listener) GetLoadBalancerPort() int32 {
	if o == nil || o.LoadBalancerPort == nil {
		var ret int32
		return ret
	}
	return *o.LoadBalancerPort
}

// GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetLoadBalancerPortOk() (*int32, bool) {
	if o == nil || o.LoadBalancerPort == nil {
		return nil, false
	}
	return o.LoadBalancerPort, true
}

// HasLoadBalancerPort returns a boolean if a field has been set.
func (o *Listener) HasLoadBalancerPort() bool {
	if o != nil && o.LoadBalancerPort != nil {
		return true
	}

	return false
}

// SetLoadBalancerPort gets a reference to the given int32 and assigns it to the LoadBalancerPort field.
func (o *Listener) SetLoadBalancerPort(v int32) {
	o.LoadBalancerPort = &v
}

// GetLoadBalancerProtocol returns the LoadBalancerProtocol field value if set, zero value otherwise.
func (o *Listener) GetLoadBalancerProtocol() string {
	if o == nil || o.LoadBalancerProtocol == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerProtocol
}

// GetLoadBalancerProtocolOk returns a tuple with the LoadBalancerProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetLoadBalancerProtocolOk() (*string, bool) {
	if o == nil || o.LoadBalancerProtocol == nil {
		return nil, false
	}
	return o.LoadBalancerProtocol, true
}

// HasLoadBalancerProtocol returns a boolean if a field has been set.
func (o *Listener) HasLoadBalancerProtocol() bool {
	if o != nil && o.LoadBalancerProtocol != nil {
		return true
	}

	return false
}

// SetLoadBalancerProtocol gets a reference to the given string and assigns it to the LoadBalancerProtocol field.
func (o *Listener) SetLoadBalancerProtocol(v string) {
	o.LoadBalancerProtocol = &v
}

// GetPolicyNames returns the PolicyNames field value if set, zero value otherwise.
func (o *Listener) GetPolicyNames() []string {
	if o == nil || o.PolicyNames == nil {
		var ret []string
		return ret
	}
	return *o.PolicyNames
}

// GetPolicyNamesOk returns a tuple with the PolicyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetPolicyNamesOk() (*[]string, bool) {
	if o == nil || o.PolicyNames == nil {
		return nil, false
	}
	return o.PolicyNames, true
}

// HasPolicyNames returns a boolean if a field has been set.
func (o *Listener) HasPolicyNames() bool {
	if o != nil && o.PolicyNames != nil {
		return true
	}

	return false
}

// SetPolicyNames gets a reference to the given []string and assigns it to the PolicyNames field.
func (o *Listener) SetPolicyNames(v []string) {
	o.PolicyNames = &v
}

// GetServerCertificateId returns the ServerCertificateId field value if set, zero value otherwise.
func (o *Listener) GetServerCertificateId() string {
	if o == nil || o.ServerCertificateId == nil {
		var ret string
		return ret
	}
	return *o.ServerCertificateId
}

// GetServerCertificateIdOk returns a tuple with the ServerCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetServerCertificateIdOk() (*string, bool) {
	if o == nil || o.ServerCertificateId == nil {
		return nil, false
	}
	return o.ServerCertificateId, true
}

// HasServerCertificateId returns a boolean if a field has been set.
func (o *Listener) HasServerCertificateId() bool {
	if o != nil && o.ServerCertificateId != nil {
		return true
	}

	return false
}

// SetServerCertificateId gets a reference to the given string and assigns it to the ServerCertificateId field.
func (o *Listener) SetServerCertificateId(v string) {
	o.ServerCertificateId = &v
}

func (o Listener) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackendPort != nil {
		toSerialize["BackendPort"] = o.BackendPort
	}
	if o.BackendProtocol != nil {
		toSerialize["BackendProtocol"] = o.BackendProtocol
	}
	if o.LoadBalancerPort != nil {
		toSerialize["LoadBalancerPort"] = o.LoadBalancerPort
	}
	if o.LoadBalancerProtocol != nil {
		toSerialize["LoadBalancerProtocol"] = o.LoadBalancerProtocol
	}
	if o.PolicyNames != nil {
		toSerialize["PolicyNames"] = o.PolicyNames
	}
	if o.ServerCertificateId != nil {
		toSerialize["ServerCertificateId"] = o.ServerCertificateId
	}
	return json.Marshal(toSerialize)
}

type NullableListener struct {
	value *Listener
	isSet bool
}

func (v NullableListener) Get() *Listener {
	return v.value
}

func (v *NullableListener) Set(val *Listener) {
	v.value = val
	v.isSet = true
}

func (v NullableListener) IsSet() bool {
	return v.isSet
}

func (v *NullableListener) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListener(val *Listener) *NullableListener {
	return &NullableListener{value: val, isSet: true}
}

func (v NullableListener) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListener) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
