/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersTag One or more filters.
type FiltersTag struct {
	// The keys of the tags that are assigned to the resources. You can use this filter alongside the `Values` filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter.
	Keys *[]string `json:"Keys,omitempty"`
	// The IDs of the resources with which the tags are associated.
	ResourceIds *[]string `json:"ResourceIds,omitempty"`
	// The resource type (`vm` \\| `image` \\| `volume` \\| `snapshot` \\| `public-ip` \\| `security-group` \\| `route-table` \\| `nic` \\| `net` \\| `subnet` \\| `net-peering` \\| `net-access-point` \\| `nat-service` \\| `internet-service` \\| `client-gateway` \\| `virtual-gateway` \\| `vpn-connection` \\| `dhcp-options` \\| `task`).
	ResourceTypes *[]string `json:"ResourceTypes,omitempty"`
	// The values of the tags that are assigned to the resources. You can use this filter alongside the `TagKeys` filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter.
	Values *[]string `json:"Values,omitempty"`
}

// NewFiltersTag instantiates a new FiltersTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersTag() *FiltersTag {
	this := FiltersTag{}
	return &this
}

// NewFiltersTagWithDefaults instantiates a new FiltersTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersTagWithDefaults() *FiltersTag {
	this := FiltersTag{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *FiltersTag) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersTag) GetKeysOk() (*[]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *FiltersTag) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *FiltersTag) SetKeys(v []string) {
	o.Keys = &v
}

// GetResourceIds returns the ResourceIds field value if set, zero value otherwise.
func (o *FiltersTag) GetResourceIds() []string {
	if o == nil || o.ResourceIds == nil {
		var ret []string
		return ret
	}
	return *o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersTag) GetResourceIdsOk() (*[]string, bool) {
	if o == nil || o.ResourceIds == nil {
		return nil, false
	}
	return o.ResourceIds, true
}

// HasResourceIds returns a boolean if a field has been set.
func (o *FiltersTag) HasResourceIds() bool {
	if o != nil && o.ResourceIds != nil {
		return true
	}

	return false
}

// SetResourceIds gets a reference to the given []string and assigns it to the ResourceIds field.
func (o *FiltersTag) SetResourceIds(v []string) {
	o.ResourceIds = &v
}

// GetResourceTypes returns the ResourceTypes field value if set, zero value otherwise.
func (o *FiltersTag) GetResourceTypes() []string {
	if o == nil || o.ResourceTypes == nil {
		var ret []string
		return ret
	}
	return *o.ResourceTypes
}

// GetResourceTypesOk returns a tuple with the ResourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersTag) GetResourceTypesOk() (*[]string, bool) {
	if o == nil || o.ResourceTypes == nil {
		return nil, false
	}
	return o.ResourceTypes, true
}

// HasResourceTypes returns a boolean if a field has been set.
func (o *FiltersTag) HasResourceTypes() bool {
	if o != nil && o.ResourceTypes != nil {
		return true
	}

	return false
}

// SetResourceTypes gets a reference to the given []string and assigns it to the ResourceTypes field.
func (o *FiltersTag) SetResourceTypes(v []string) {
	o.ResourceTypes = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FiltersTag) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersTag) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FiltersTag) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *FiltersTag) SetValues(v []string) {
	o.Values = &v
}

func (o FiltersTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["Keys"] = o.Keys
	}
	if o.ResourceIds != nil {
		toSerialize["ResourceIds"] = o.ResourceIds
	}
	if o.ResourceTypes != nil {
		toSerialize["ResourceTypes"] = o.ResourceTypes
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersTag struct {
	value *FiltersTag
	isSet bool
}

func (v NullableFiltersTag) Get() *FiltersTag {
	return v.value
}

func (v *NullableFiltersTag) Set(val *FiltersTag) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersTag) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersTag(val *FiltersTag) *NullableFiltersTag {
	return &NullableFiltersTag{value: val, isSet: true}
}

func (v NullableFiltersTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
