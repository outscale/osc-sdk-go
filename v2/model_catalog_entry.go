/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.5
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CatalogEntry Information about the catalog entry.
type CatalogEntry struct {
	// The category of the catalog entry (for example, `network`).
	Category *string `json:"Category,omitempty"`
	// When returned and equal to `PER_MONTH`, the price of the catalog entry is calculated on a monthly basis.
	Flags *string `json:"Flags,omitempty"`
	// The API call associated with the catalog entry (for example, `CreateVms` or `RunInstances`).
	Operation *string `json:"Operation,omitempty"`
	// The service associated with the catalog entry (`TinaOS-FCU`, `TinaOS-LBU`, `TinaOS-DirectLink`, or `TinaOS-OOS`).
	Service *string `json:"Service,omitempty"`
	// The Subregion associated with the catalog entry.
	SubregionName *string `json:"SubregionName,omitempty"`
	// The description of the catalog entry.
	Title *string `json:"Title,omitempty"`
	// The type of resource associated with the catalog entry.
	Type *string `json:"Type,omitempty"`
	// The unit price of the catalog entry, in the currency of the catalog of the Region where the API method was used.
	UnitPrice *float32 `json:"UnitPrice,omitempty"`
}

// NewCatalogEntry instantiates a new CatalogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogEntry() *CatalogEntry {
	this := CatalogEntry{}
	return &this
}

// NewCatalogEntryWithDefaults instantiates a new CatalogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogEntryWithDefaults() *CatalogEntry {
	this := CatalogEntry{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CatalogEntry) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntry) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CatalogEntry) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CatalogEntry) SetCategory(v string) {
	o.Category = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *CatalogEntry) GetFlags() string {
	if o == nil || o.Flags == nil {
		var ret string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntry) GetFlagsOk() (*string, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *CatalogEntry) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given string and assigns it to the Flags field.
func (o *CatalogEntry) SetFlags(v string) {
	o.Flags = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *CatalogEntry) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntry) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *CatalogEntry) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *CatalogEntry) SetOperation(v string) {
	o.Operation = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CatalogEntry) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntry) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CatalogEntry) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *CatalogEntry) SetService(v string) {
	o.Service = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *CatalogEntry) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntry) GetSubregionNameOk() (*string, bool) {
	if o == nil || o.SubregionName == nil {
		return nil, false
	}
	return o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *CatalogEntry) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *CatalogEntry) SetSubregionName(v string) {
	o.SubregionName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CatalogEntry) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntry) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CatalogEntry) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CatalogEntry) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CatalogEntry) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntry) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CatalogEntry) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CatalogEntry) SetType(v string) {
	o.Type = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *CatalogEntry) GetUnitPrice() float32 {
	if o == nil || o.UnitPrice == nil {
		var ret float32
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntry) GetUnitPriceOk() (*float32, bool) {
	if o == nil || o.UnitPrice == nil {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *CatalogEntry) HasUnitPrice() bool {
	if o != nil && o.UnitPrice != nil {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given float32 and assigns it to the UnitPrice field.
func (o *CatalogEntry) SetUnitPrice(v float32) {
	o.UnitPrice = &v
}

func (o CatalogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.Flags != nil {
		toSerialize["Flags"] = o.Flags
	}
	if o.Operation != nil {
		toSerialize["Operation"] = o.Operation
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
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.UnitPrice != nil {
		toSerialize["UnitPrice"] = o.UnitPrice
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogEntry struct {
	value *CatalogEntry
	isSet bool
}

func (v NullableCatalogEntry) Get() *CatalogEntry {
	return v.value
}

func (v *NullableCatalogEntry) Set(val *CatalogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogEntry(val *CatalogEntry) *NullableCatalogEntry {
	return &NullableCatalogEntry{value: val, isSet: true}
}

func (v NullableCatalogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
