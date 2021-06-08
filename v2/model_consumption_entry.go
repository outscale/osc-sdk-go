/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ConsumptionEntry Information about the resources consumed during the specified time period.
type ConsumptionEntry struct {
	// The ID of your TINA account.
	AccountId *string `json:"AccountId,omitempty"`
	// The category of the resource (for example, `network`).
	Category *string `json:"Category,omitempty"`
	// The beginning of the time period.
	FromDate *string `json:"FromDate,omitempty"`
	// The API call that triggered the resource consumption (for example, `RunInstances` or `CreateVolume`).
	Operation *string `json:"Operation,omitempty"`
	// The ID of the TINA account which is billed for your consumption. It can be different from your account in the `AccountId` parameter.
	PayingAccountId *string `json:"PayingAccountId,omitempty"`
	// The service of the API call (`TinaOS-FCU`, `TinaOS-LBU`, `TinaOS-DirectLink`, `TinaOS-OOS`, or `TinaOS-OSU`).
	Service *string `json:"Service,omitempty"`
	// The name of the Subregion.
	SubregionName *string `json:"SubregionName,omitempty"`
	// A description of the consumed resource.
	Title *string `json:"Title,omitempty"`
	// The end of the time period.
	ToDate *string `json:"ToDate,omitempty"`
	// The type of resource, depending on the API call.
	Type *string `json:"Type,omitempty"`
	// The consumed amount for the resource. The unit depends on the resource type. For more information, see the `Title` element.
	Value *float64 `json:"Value,omitempty"`
}

// NewConsumptionEntry instantiates a new ConsumptionEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumptionEntry() *ConsumptionEntry {
	this := ConsumptionEntry{}
	return &this
}

// NewConsumptionEntryWithDefaults instantiates a new ConsumptionEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumptionEntryWithDefaults() *ConsumptionEntry {
	this := ConsumptionEntry{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ConsumptionEntry) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ConsumptionEntry) SetCategory(v string) {
	o.Category = &v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetFromDate() string {
	if o == nil || o.FromDate == nil {
		var ret string
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetFromDateOk() (*string, bool) {
	if o == nil || o.FromDate == nil {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasFromDate() bool {
	if o != nil && o.FromDate != nil {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given string and assigns it to the FromDate field.
func (o *ConsumptionEntry) SetFromDate(v string) {
	o.FromDate = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ConsumptionEntry) SetOperation(v string) {
	o.Operation = &v
}

// GetPayingAccountId returns the PayingAccountId field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetPayingAccountId() string {
	if o == nil || o.PayingAccountId == nil {
		var ret string
		return ret
	}
	return *o.PayingAccountId
}

// GetPayingAccountIdOk returns a tuple with the PayingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetPayingAccountIdOk() (*string, bool) {
	if o == nil || o.PayingAccountId == nil {
		return nil, false
	}
	return o.PayingAccountId, true
}

// HasPayingAccountId returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasPayingAccountId() bool {
	if o != nil && o.PayingAccountId != nil {
		return true
	}

	return false
}

// SetPayingAccountId gets a reference to the given string and assigns it to the PayingAccountId field.
func (o *ConsumptionEntry) SetPayingAccountId(v string) {
	o.PayingAccountId = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ConsumptionEntry) SetService(v string) {
	o.Service = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetSubregionNameOk() (*string, bool) {
	if o == nil || o.SubregionName == nil {
		return nil, false
	}
	return o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *ConsumptionEntry) SetSubregionName(v string) {
	o.SubregionName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ConsumptionEntry) SetTitle(v string) {
	o.Title = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetToDate() string {
	if o == nil || o.ToDate == nil {
		var ret string
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetToDateOk() (*string, bool) {
	if o == nil || o.ToDate == nil {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasToDate() bool {
	if o != nil && o.ToDate != nil {
		return true
	}

	return false
}

// SetToDate gets a reference to the given string and assigns it to the ToDate field.
func (o *ConsumptionEntry) SetToDate(v string) {
	o.ToDate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConsumptionEntry) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConsumptionEntry) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionEntry) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConsumptionEntry) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *ConsumptionEntry) SetValue(v float64) {
	o.Value = &v
}

func (o ConsumptionEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.FromDate != nil {
		toSerialize["FromDate"] = o.FromDate
	}
	if o.Operation != nil {
		toSerialize["Operation"] = o.Operation
	}
	if o.PayingAccountId != nil {
		toSerialize["PayingAccountId"] = o.PayingAccountId
	}
	if o.Service != nil {
		toSerialize["Service"] = o.Service
	}
	if o.SubregionName != nil {
		toSerialize["SubregionName"] = o.SubregionName
	}
	if o.Title != nil {
		toSerialize["Title"] = o.Title
	}
	if o.ToDate != nil {
		toSerialize["ToDate"] = o.ToDate
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableConsumptionEntry struct {
	value *ConsumptionEntry
	isSet bool
}

func (v NullableConsumptionEntry) Get() *ConsumptionEntry {
	return v.value
}

func (v *NullableConsumptionEntry) Set(val *ConsumptionEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumptionEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumptionEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumptionEntry(val *ConsumptionEntry) *NullableConsumptionEntry {
	return &NullableConsumptionEntry{value: val, isSet: true}
}

func (v NullableConsumptionEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumptionEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}