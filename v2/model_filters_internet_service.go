/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersInternetService One or more filters.
type FiltersInternetService struct {
	// The IDs of the Internet services.
	InternetServiceIds *[]string `json:"InternetServiceIds,omitempty"`
	// The IDs of the Nets the Internet services are attached to.
	LinkNetIds *[]string `json:"LinkNetIds,omitempty"`
	// The current states of the attachments between the Internet services and the Nets (only `available`, if the Internet gateway is attached to a VPC).
	LinkStates *[]string `json:"LinkStates,omitempty"`
	// The keys of the tags associated with the Internet services.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Internet services.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Internet services, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersInternetService instantiates a new FiltersInternetService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersInternetService() *FiltersInternetService {
	this := FiltersInternetService{}
	return &this
}

// NewFiltersInternetServiceWithDefaults instantiates a new FiltersInternetService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersInternetServiceWithDefaults() *FiltersInternetService {
	this := FiltersInternetService{}
	return &this
}

// GetInternetServiceIds returns the InternetServiceIds field value if set, zero value otherwise.
func (o *FiltersInternetService) GetInternetServiceIds() []string {
	if o == nil || o.InternetServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.InternetServiceIds
}

// GetInternetServiceIdsOk returns a tuple with the InternetServiceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersInternetService) GetInternetServiceIdsOk() (*[]string, bool) {
	if o == nil || o.InternetServiceIds == nil {
		return nil, false
	}
	return o.InternetServiceIds, true
}

// HasInternetServiceIds returns a boolean if a field has been set.
func (o *FiltersInternetService) HasInternetServiceIds() bool {
	if o != nil && o.InternetServiceIds != nil {
		return true
	}

	return false
}

// SetInternetServiceIds gets a reference to the given []string and assigns it to the InternetServiceIds field.
func (o *FiltersInternetService) SetInternetServiceIds(v []string) {
	o.InternetServiceIds = &v
}

// GetLinkNetIds returns the LinkNetIds field value if set, zero value otherwise.
func (o *FiltersInternetService) GetLinkNetIds() []string {
	if o == nil || o.LinkNetIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkNetIds
}

// GetLinkNetIdsOk returns a tuple with the LinkNetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersInternetService) GetLinkNetIdsOk() (*[]string, bool) {
	if o == nil || o.LinkNetIds == nil {
		return nil, false
	}
	return o.LinkNetIds, true
}

// HasLinkNetIds returns a boolean if a field has been set.
func (o *FiltersInternetService) HasLinkNetIds() bool {
	if o != nil && o.LinkNetIds != nil {
		return true
	}

	return false
}

// SetLinkNetIds gets a reference to the given []string and assigns it to the LinkNetIds field.
func (o *FiltersInternetService) SetLinkNetIds(v []string) {
	o.LinkNetIds = &v
}

// GetLinkStates returns the LinkStates field value if set, zero value otherwise.
func (o *FiltersInternetService) GetLinkStates() []string {
	if o == nil || o.LinkStates == nil {
		var ret []string
		return ret
	}
	return *o.LinkStates
}

// GetLinkStatesOk returns a tuple with the LinkStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersInternetService) GetLinkStatesOk() (*[]string, bool) {
	if o == nil || o.LinkStates == nil {
		return nil, false
	}
	return o.LinkStates, true
}

// HasLinkStates returns a boolean if a field has been set.
func (o *FiltersInternetService) HasLinkStates() bool {
	if o != nil && o.LinkStates != nil {
		return true
	}

	return false
}

// SetLinkStates gets a reference to the given []string and assigns it to the LinkStates field.
func (o *FiltersInternetService) SetLinkStates(v []string) {
	o.LinkStates = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersInternetService) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersInternetService) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersInternetService) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersInternetService) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersInternetService) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersInternetService) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersInternetService) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersInternetService) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersInternetService) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersInternetService) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersInternetService) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersInternetService) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersInternetService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InternetServiceIds != nil {
		toSerialize["InternetServiceIds"] = o.InternetServiceIds
	}
	if o.LinkNetIds != nil {
		toSerialize["LinkNetIds"] = o.LinkNetIds
	}
	if o.LinkStates != nil {
		toSerialize["LinkStates"] = o.LinkStates
	}
	if o.TagKeys != nil {
		toSerialize["TagKeys"] = o.TagKeys
	}
	if o.TagValues != nil {
		toSerialize["TagValues"] = o.TagValues
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersInternetService struct {
	value *FiltersInternetService
	isSet bool
}

func (v NullableFiltersInternetService) Get() *FiltersInternetService {
	return v.value
}

func (v *NullableFiltersInternetService) Set(val *FiltersInternetService) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersInternetService) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersInternetService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersInternetService(val *FiltersInternetService) *NullableFiltersInternetService {
	return &NullableFiltersInternetService{value: val, isSet: true}
}

func (v NullableFiltersInternetService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersInternetService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
