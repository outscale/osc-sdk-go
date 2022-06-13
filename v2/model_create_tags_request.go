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

// CreateTagsRequest struct for CreateTagsRequest
type CreateTagsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more resource IDs.
	ResourceIds []string `json:"ResourceIds"`
	// One or more tags to add to the specified resources.
	Tags []ResourceTag `json:"Tags"`
}

// NewCreateTagsRequest instantiates a new CreateTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTagsRequest(resourceIds []string, tags []ResourceTag) *CreateTagsRequest {
	this := CreateTagsRequest{}
	this.ResourceIds = resourceIds
	this.Tags = tags
	return &this
}

// NewCreateTagsRequestWithDefaults instantiates a new CreateTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTagsRequestWithDefaults() *CreateTagsRequest {
	this := CreateTagsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateTagsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTagsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateTagsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateTagsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetResourceIds returns the ResourceIds field value
func (o *CreateTagsRequest) GetResourceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value
// and a boolean to check if the value has been set.
func (o *CreateTagsRequest) GetResourceIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceIds, true
}

// SetResourceIds sets field value
func (o *CreateTagsRequest) SetResourceIds(v []string) {
	o.ResourceIds = v
}

// GetTags returns the Tags field value
func (o *CreateTagsRequest) GetTags() []ResourceTag {
	if o == nil {
		var ret []ResourceTag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreateTagsRequest) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *CreateTagsRequest) SetTags(v []ResourceTag) {
	o.Tags = v
}

func (o CreateTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["ResourceIds"] = o.ResourceIds
	}
	if true {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTagsRequest struct {
	value *CreateTagsRequest
	isSet bool
}

func (v NullableCreateTagsRequest) Get() *CreateTagsRequest {
	return v.value
}

func (v *NullableCreateTagsRequest) Set(val *CreateTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTagsRequest(val *CreateTagsRequest) *NullableCreateTagsRequest {
	return &NullableCreateTagsRequest{value: val, isSet: true}
}

func (v NullableCreateTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
