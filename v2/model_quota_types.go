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

// QuotaTypes One or more quotas.
type QuotaTypes struct {
	// The resource ID if it is a resource-specific quota, `global` if it is not.
	QuotaType *string `json:"QuotaType,omitempty"`
	// One or more quotas associated with the user.
	Quotas *[]Quota `json:"Quotas,omitempty"`
}

// NewQuotaTypes instantiates a new QuotaTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaTypes() *QuotaTypes {
	this := QuotaTypes{}
	return &this
}

// NewQuotaTypesWithDefaults instantiates a new QuotaTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaTypesWithDefaults() *QuotaTypes {
	this := QuotaTypes{}
	return &this
}

// GetQuotaType returns the QuotaType field value if set, zero value otherwise.
func (o *QuotaTypes) GetQuotaType() string {
	if o == nil || o.QuotaType == nil {
		var ret string
		return ret
	}
	return *o.QuotaType
}

// GetQuotaTypeOk returns a tuple with the QuotaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaTypes) GetQuotaTypeOk() (*string, bool) {
	if o == nil || o.QuotaType == nil {
		return nil, false
	}
	return o.QuotaType, true
}

// HasQuotaType returns a boolean if a field has been set.
func (o *QuotaTypes) HasQuotaType() bool {
	if o != nil && o.QuotaType != nil {
		return true
	}

	return false
}

// SetQuotaType gets a reference to the given string and assigns it to the QuotaType field.
func (o *QuotaTypes) SetQuotaType(v string) {
	o.QuotaType = &v
}

// GetQuotas returns the Quotas field value if set, zero value otherwise.
func (o *QuotaTypes) GetQuotas() []Quota {
	if o == nil || o.Quotas == nil {
		var ret []Quota
		return ret
	}
	return *o.Quotas
}

// GetQuotasOk returns a tuple with the Quotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaTypes) GetQuotasOk() (*[]Quota, bool) {
	if o == nil || o.Quotas == nil {
		return nil, false
	}
	return o.Quotas, true
}

// HasQuotas returns a boolean if a field has been set.
func (o *QuotaTypes) HasQuotas() bool {
	if o != nil && o.Quotas != nil {
		return true
	}

	return false
}

// SetQuotas gets a reference to the given []Quota and assigns it to the Quotas field.
func (o *QuotaTypes) SetQuotas(v []Quota) {
	o.Quotas = &v
}

func (o QuotaTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QuotaType != nil {
		toSerialize["QuotaType"] = o.QuotaType
	}
	if o.Quotas != nil {
		toSerialize["Quotas"] = o.Quotas
	}
	return json.Marshal(toSerialize)
}

type NullableQuotaTypes struct {
	value *QuotaTypes
	isSet bool
}

func (v NullableQuotaTypes) Get() *QuotaTypes {
	return v.value
}

func (v *NullableQuotaTypes) Set(val *QuotaTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaTypes(val *QuotaTypes) *NullableQuotaTypes {
	return &NullableQuotaTypes{value: val, isSet: true}
}

func (v NullableQuotaTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
