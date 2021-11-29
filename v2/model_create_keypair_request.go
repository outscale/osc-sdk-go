/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.16
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateKeypairRequest struct for CreateKeypairRequest
type CreateKeypairRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// A unique name for the keypair, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).
	KeypairName string `json:"KeypairName"`
	// The public key. It must be Base64-encoded.
	PublicKey *string `json:"PublicKey,omitempty"`
}

// NewCreateKeypairRequest instantiates a new CreateKeypairRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeypairRequest(keypairName string) *CreateKeypairRequest {
	this := CreateKeypairRequest{}
	this.KeypairName = keypairName
	return &this
}

// NewCreateKeypairRequestWithDefaults instantiates a new CreateKeypairRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeypairRequestWithDefaults() *CreateKeypairRequest {
	this := CreateKeypairRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateKeypairRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKeypairRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateKeypairRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateKeypairRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetKeypairName returns the KeypairName field value
func (o *CreateKeypairRequest) GetKeypairName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value
// and a boolean to check if the value has been set.
func (o *CreateKeypairRequest) GetKeypairNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeypairName, true
}

// SetKeypairName sets field value
func (o *CreateKeypairRequest) SetKeypairName(v string) {
	o.KeypairName = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *CreateKeypairRequest) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKeypairRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CreateKeypairRequest) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *CreateKeypairRequest) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o CreateKeypairRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["KeypairName"] = o.KeypairName
	}
	if o.PublicKey != nil {
		toSerialize["PublicKey"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableCreateKeypairRequest struct {
	value *CreateKeypairRequest
	isSet bool
}

func (v NullableCreateKeypairRequest) Get() *CreateKeypairRequest {
	return v.value
}

func (v *NullableCreateKeypairRequest) Set(val *CreateKeypairRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeypairRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeypairRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeypairRequest(val *CreateKeypairRequest) *NullableCreateKeypairRequest {
	return &NullableCreateKeypairRequest{value: val, isSet: true}
}

func (v NullableCreateKeypairRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeypairRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
