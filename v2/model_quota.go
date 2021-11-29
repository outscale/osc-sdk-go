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

// Quota Information about the quota.
type Quota struct {
	// The account ID of the owner of the quotas.
	AccountId *string `json:"AccountId,omitempty"`
	// The description of the quota.
	Description *string `json:"Description,omitempty"`
	// The maximum value of the quota for the OUTSCALE user account (if there is no limit, `0`).
	MaxValue *int32 `json:"MaxValue,omitempty"`
	// The unique name of the quota.
	Name *string `json:"Name,omitempty"`
	// The group name of the quota.
	QuotaCollection *string `json:"QuotaCollection,omitempty"`
	// The description of the quota.
	ShortDescription *string `json:"ShortDescription,omitempty"`
	// The limit value currently used by the OUTSCALE user account.
	UsedValue *int32 `json:"UsedValue,omitempty"`
}

// NewQuota instantiates a new Quota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuota() *Quota {
	this := Quota{}
	return &this
}

// NewQuotaWithDefaults instantiates a new Quota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaWithDefaults() *Quota {
	this := Quota{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Quota) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Quota) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Quota) SetAccountId(v string) {
	o.AccountId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Quota) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Quota) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Quota) SetDescription(v string) {
	o.Description = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *Quota) GetMaxValue() int32 {
	if o == nil || o.MaxValue == nil {
		var ret int32
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetMaxValueOk() (*int32, bool) {
	if o == nil || o.MaxValue == nil {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *Quota) HasMaxValue() bool {
	if o != nil && o.MaxValue != nil {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int32 and assigns it to the MaxValue field.
func (o *Quota) SetMaxValue(v int32) {
	o.MaxValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Quota) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Quota) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Quota) SetName(v string) {
	o.Name = &v
}

// GetQuotaCollection returns the QuotaCollection field value if set, zero value otherwise.
func (o *Quota) GetQuotaCollection() string {
	if o == nil || o.QuotaCollection == nil {
		var ret string
		return ret
	}
	return *o.QuotaCollection
}

// GetQuotaCollectionOk returns a tuple with the QuotaCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetQuotaCollectionOk() (*string, bool) {
	if o == nil || o.QuotaCollection == nil {
		return nil, false
	}
	return o.QuotaCollection, true
}

// HasQuotaCollection returns a boolean if a field has been set.
func (o *Quota) HasQuotaCollection() bool {
	if o != nil && o.QuotaCollection != nil {
		return true
	}

	return false
}

// SetQuotaCollection gets a reference to the given string and assigns it to the QuotaCollection field.
func (o *Quota) SetQuotaCollection(v string) {
	o.QuotaCollection = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *Quota) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *Quota) HasShortDescription() bool {
	if o != nil && o.ShortDescription != nil {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *Quota) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetUsedValue returns the UsedValue field value if set, zero value otherwise.
func (o *Quota) GetUsedValue() int32 {
	if o == nil || o.UsedValue == nil {
		var ret int32
		return ret
	}
	return *o.UsedValue
}

// GetUsedValueOk returns a tuple with the UsedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetUsedValueOk() (*int32, bool) {
	if o == nil || o.UsedValue == nil {
		return nil, false
	}
	return o.UsedValue, true
}

// HasUsedValue returns a boolean if a field has been set.
func (o *Quota) HasUsedValue() bool {
	if o != nil && o.UsedValue != nil {
		return true
	}

	return false
}

// SetUsedValue gets a reference to the given int32 and assigns it to the UsedValue field.
func (o *Quota) SetUsedValue(v int32) {
	o.UsedValue = &v
}

func (o Quota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.MaxValue != nil {
		toSerialize["MaxValue"] = o.MaxValue
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.QuotaCollection != nil {
		toSerialize["QuotaCollection"] = o.QuotaCollection
	}
	if o.ShortDescription != nil {
		toSerialize["ShortDescription"] = o.ShortDescription
	}
	if o.UsedValue != nil {
		toSerialize["UsedValue"] = o.UsedValue
	}
	return json.Marshal(toSerialize)
}

type NullableQuota struct {
	value *Quota
	isSet bool
}

func (v NullableQuota) Get() *Quota {
	return v.value
}

func (v *NullableQuota) Set(val *Quota) {
	v.value = val
	v.isSet = true
}

func (v NullableQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuota(val *Quota) *NullableQuota {
	return &NullableQuota{value: val, isSet: true}
}

func (v NullableQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
