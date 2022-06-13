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

// FiltersVmsState One or more filters.
type FiltersVmsState struct {
	// The code for the scheduled event (`system-reboot` \\| `system-maintenance`).
	MaintenanceEventCodes *[]string `json:"MaintenanceEventCodes,omitempty"`
	// The description of the scheduled event.
	MaintenanceEventDescriptions *[]string `json:"MaintenanceEventDescriptions,omitempty"`
	// The latest time the event can end.
	MaintenanceEventsNotAfter *[]string `json:"MaintenanceEventsNotAfter,omitempty"`
	// The earliest time the event can start.
	MaintenanceEventsNotBefore *[]string `json:"MaintenanceEventsNotBefore,omitempty"`
	// The names of the Subregions of the VMs.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
	// One or more IDs of VMs.
	VmIds *[]string `json:"VmIds,omitempty"`
	// The states of the VMs (`pending` \\| `running` \\| `stopping` \\| `stopped` \\| `shutting-down` \\| `terminated` \\| `quarantine`).
	VmStates *[]string `json:"VmStates,omitempty"`
}

// NewFiltersVmsState instantiates a new FiltersVmsState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersVmsState() *FiltersVmsState {
	this := FiltersVmsState{}
	return &this
}

// NewFiltersVmsStateWithDefaults instantiates a new FiltersVmsState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersVmsStateWithDefaults() *FiltersVmsState {
	this := FiltersVmsState{}
	return &this
}

// GetMaintenanceEventCodes returns the MaintenanceEventCodes field value if set, zero value otherwise.
func (o *FiltersVmsState) GetMaintenanceEventCodes() []string {
	if o == nil || o.MaintenanceEventCodes == nil {
		var ret []string
		return ret
	}
	return *o.MaintenanceEventCodes
}

// GetMaintenanceEventCodesOk returns a tuple with the MaintenanceEventCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetMaintenanceEventCodesOk() (*[]string, bool) {
	if o == nil || o.MaintenanceEventCodes == nil {
		return nil, false
	}
	return o.MaintenanceEventCodes, true
}

// HasMaintenanceEventCodes returns a boolean if a field has been set.
func (o *FiltersVmsState) HasMaintenanceEventCodes() bool {
	if o != nil && o.MaintenanceEventCodes != nil {
		return true
	}

	return false
}

// SetMaintenanceEventCodes gets a reference to the given []string and assigns it to the MaintenanceEventCodes field.
func (o *FiltersVmsState) SetMaintenanceEventCodes(v []string) {
	o.MaintenanceEventCodes = &v
}

// GetMaintenanceEventDescriptions returns the MaintenanceEventDescriptions field value if set, zero value otherwise.
func (o *FiltersVmsState) GetMaintenanceEventDescriptions() []string {
	if o == nil || o.MaintenanceEventDescriptions == nil {
		var ret []string
		return ret
	}
	return *o.MaintenanceEventDescriptions
}

// GetMaintenanceEventDescriptionsOk returns a tuple with the MaintenanceEventDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetMaintenanceEventDescriptionsOk() (*[]string, bool) {
	if o == nil || o.MaintenanceEventDescriptions == nil {
		return nil, false
	}
	return o.MaintenanceEventDescriptions, true
}

// HasMaintenanceEventDescriptions returns a boolean if a field has been set.
func (o *FiltersVmsState) HasMaintenanceEventDescriptions() bool {
	if o != nil && o.MaintenanceEventDescriptions != nil {
		return true
	}

	return false
}

// SetMaintenanceEventDescriptions gets a reference to the given []string and assigns it to the MaintenanceEventDescriptions field.
func (o *FiltersVmsState) SetMaintenanceEventDescriptions(v []string) {
	o.MaintenanceEventDescriptions = &v
}

// GetMaintenanceEventsNotAfter returns the MaintenanceEventsNotAfter field value if set, zero value otherwise.
func (o *FiltersVmsState) GetMaintenanceEventsNotAfter() []string {
	if o == nil || o.MaintenanceEventsNotAfter == nil {
		var ret []string
		return ret
	}
	return *o.MaintenanceEventsNotAfter
}

// GetMaintenanceEventsNotAfterOk returns a tuple with the MaintenanceEventsNotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetMaintenanceEventsNotAfterOk() (*[]string, bool) {
	if o == nil || o.MaintenanceEventsNotAfter == nil {
		return nil, false
	}
	return o.MaintenanceEventsNotAfter, true
}

// HasMaintenanceEventsNotAfter returns a boolean if a field has been set.
func (o *FiltersVmsState) HasMaintenanceEventsNotAfter() bool {
	if o != nil && o.MaintenanceEventsNotAfter != nil {
		return true
	}

	return false
}

// SetMaintenanceEventsNotAfter gets a reference to the given []string and assigns it to the MaintenanceEventsNotAfter field.
func (o *FiltersVmsState) SetMaintenanceEventsNotAfter(v []string) {
	o.MaintenanceEventsNotAfter = &v
}

// GetMaintenanceEventsNotBefore returns the MaintenanceEventsNotBefore field value if set, zero value otherwise.
func (o *FiltersVmsState) GetMaintenanceEventsNotBefore() []string {
	if o == nil || o.MaintenanceEventsNotBefore == nil {
		var ret []string
		return ret
	}
	return *o.MaintenanceEventsNotBefore
}

// GetMaintenanceEventsNotBeforeOk returns a tuple with the MaintenanceEventsNotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetMaintenanceEventsNotBeforeOk() (*[]string, bool) {
	if o == nil || o.MaintenanceEventsNotBefore == nil {
		return nil, false
	}
	return o.MaintenanceEventsNotBefore, true
}

// HasMaintenanceEventsNotBefore returns a boolean if a field has been set.
func (o *FiltersVmsState) HasMaintenanceEventsNotBefore() bool {
	if o != nil && o.MaintenanceEventsNotBefore != nil {
		return true
	}

	return false
}

// SetMaintenanceEventsNotBefore gets a reference to the given []string and assigns it to the MaintenanceEventsNotBefore field.
func (o *FiltersVmsState) SetMaintenanceEventsNotBefore(v []string) {
	o.MaintenanceEventsNotBefore = &v
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *FiltersVmsState) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetSubregionNamesOk() (*[]string, bool) {
	if o == nil || o.SubregionNames == nil {
		return nil, false
	}
	return o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *FiltersVmsState) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *FiltersVmsState) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

// GetVmIds returns the VmIds field value if set, zero value otherwise.
func (o *FiltersVmsState) GetVmIds() []string {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret
	}
	return *o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetVmIdsOk() (*[]string, bool) {
	if o == nil || o.VmIds == nil {
		return nil, false
	}
	return o.VmIds, true
}

// HasVmIds returns a boolean if a field has been set.
func (o *FiltersVmsState) HasVmIds() bool {
	if o != nil && o.VmIds != nil {
		return true
	}

	return false
}

// SetVmIds gets a reference to the given []string and assigns it to the VmIds field.
func (o *FiltersVmsState) SetVmIds(v []string) {
	o.VmIds = &v
}

// GetVmStates returns the VmStates field value if set, zero value otherwise.
func (o *FiltersVmsState) GetVmStates() []string {
	if o == nil || o.VmStates == nil {
		var ret []string
		return ret
	}
	return *o.VmStates
}

// GetVmStatesOk returns a tuple with the VmStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetVmStatesOk() (*[]string, bool) {
	if o == nil || o.VmStates == nil {
		return nil, false
	}
	return o.VmStates, true
}

// HasVmStates returns a boolean if a field has been set.
func (o *FiltersVmsState) HasVmStates() bool {
	if o != nil && o.VmStates != nil {
		return true
	}

	return false
}

// SetVmStates gets a reference to the given []string and assigns it to the VmStates field.
func (o *FiltersVmsState) SetVmStates(v []string) {
	o.VmStates = &v
}

func (o FiltersVmsState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaintenanceEventCodes != nil {
		toSerialize["MaintenanceEventCodes"] = o.MaintenanceEventCodes
	}
	if o.MaintenanceEventDescriptions != nil {
		toSerialize["MaintenanceEventDescriptions"] = o.MaintenanceEventDescriptions
	}
	if o.MaintenanceEventsNotAfter != nil {
		toSerialize["MaintenanceEventsNotAfter"] = o.MaintenanceEventsNotAfter
	}
	if o.MaintenanceEventsNotBefore != nil {
		toSerialize["MaintenanceEventsNotBefore"] = o.MaintenanceEventsNotBefore
	}
	if o.SubregionNames != nil {
		toSerialize["SubregionNames"] = o.SubregionNames
	}
	if o.VmIds != nil {
		toSerialize["VmIds"] = o.VmIds
	}
	if o.VmStates != nil {
		toSerialize["VmStates"] = o.VmStates
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersVmsState struct {
	value *FiltersVmsState
	isSet bool
}

func (v NullableFiltersVmsState) Get() *FiltersVmsState {
	return v.value
}

func (v *NullableFiltersVmsState) Set(val *FiltersVmsState) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersVmsState) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersVmsState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersVmsState(val *FiltersVmsState) *NullableFiltersVmsState {
	return &NullableFiltersVmsState{value: val, isSet: true}
}

func (v NullableFiltersVmsState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersVmsState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
