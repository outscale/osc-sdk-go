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

// Catalog Information about our catalog of prices.
type Catalog struct {
	// One or more catalog entries.
	Entries *[]CatalogEntry `json:"Entries,omitempty"`
}

// NewCatalog instantiates a new Catalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalog() *Catalog {
	this := Catalog{}
	return &this
}

// NewCatalogWithDefaults instantiates a new Catalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogWithDefaults() *Catalog {
	this := Catalog{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *Catalog) GetEntries() []CatalogEntry {
	if o == nil || o.Entries == nil {
		var ret []CatalogEntry
		return ret
	}
	return *o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Catalog) GetEntriesOk() (*[]CatalogEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *Catalog) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []CatalogEntry and assigns it to the Entries field.
func (o *Catalog) SetEntries(v []CatalogEntry) {
	o.Entries = &v
}

func (o Catalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entries != nil {
		toSerialize["Entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableCatalog struct {
	value *Catalog
	isSet bool
}

func (v NullableCatalog) Get() *Catalog {
	return v.value
}

func (v *NullableCatalog) Set(val *Catalog) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalog(val *Catalog) *NullableCatalog {
	return &NullableCatalog{value: val, isSet: true}
}

func (v NullableCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
