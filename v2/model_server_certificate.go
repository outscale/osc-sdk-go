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

// ServerCertificate Information about the server certificate.
type ServerCertificate struct {
	// The date at which the server certificate expires.
	ExpirationDate *string `json:"ExpirationDate,omitempty"`
	// The ID of the server certificate.
	Id *string `json:"Id,omitempty"`
	// The name of the server certificate.
	Name *string `json:"Name,omitempty"`
	// The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).
	Orn *string `json:"Orn,omitempty"`
	// The path to the server certificate.
	Path *string `json:"Path,omitempty"`
	// The date at which the server certificate has been uploaded.
	UploadDate *string `json:"UploadDate,omitempty"`
}

// NewServerCertificate instantiates a new ServerCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerCertificate() *ServerCertificate {
	this := ServerCertificate{}
	return &this
}

// NewServerCertificateWithDefaults instantiates a new ServerCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerCertificateWithDefaults() *ServerCertificate {
	this := ServerCertificate{}
	return &this
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ServerCertificate) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCertificate) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ServerCertificate) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *ServerCertificate) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerCertificate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCertificate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerCertificate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerCertificate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerCertificate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCertificate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerCertificate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerCertificate) SetName(v string) {
	o.Name = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ServerCertificate) GetOrn() string {
	if o == nil || o.Orn == nil {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCertificate) GetOrnOk() (*string, bool) {
	if o == nil || o.Orn == nil {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ServerCertificate) HasOrn() bool {
	if o != nil && o.Orn != nil {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ServerCertificate) SetOrn(v string) {
	o.Orn = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ServerCertificate) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCertificate) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ServerCertificate) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ServerCertificate) SetPath(v string) {
	o.Path = &v
}

// GetUploadDate returns the UploadDate field value if set, zero value otherwise.
func (o *ServerCertificate) GetUploadDate() string {
	if o == nil || o.UploadDate == nil {
		var ret string
		return ret
	}
	return *o.UploadDate
}

// GetUploadDateOk returns a tuple with the UploadDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCertificate) GetUploadDateOk() (*string, bool) {
	if o == nil || o.UploadDate == nil {
		return nil, false
	}
	return o.UploadDate, true
}

// HasUploadDate returns a boolean if a field has been set.
func (o *ServerCertificate) HasUploadDate() bool {
	if o != nil && o.UploadDate != nil {
		return true
	}

	return false
}

// SetUploadDate gets a reference to the given string and assigns it to the UploadDate field.
func (o *ServerCertificate) SetUploadDate(v string) {
	o.UploadDate = &v
}

func (o ServerCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationDate != nil {
		toSerialize["ExpirationDate"] = o.ExpirationDate
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Orn != nil {
		toSerialize["Orn"] = o.Orn
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.UploadDate != nil {
		toSerialize["UploadDate"] = o.UploadDate
	}
	return json.Marshal(toSerialize)
}

type NullableServerCertificate struct {
	value *ServerCertificate
	isSet bool
}

func (v NullableServerCertificate) Get() *ServerCertificate {
	return v.value
}

func (v *NullableServerCertificate) Set(val *ServerCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerCertificate(val *ServerCertificate) *NullableServerCertificate {
	return &NullableServerCertificate{value: val, isSet: true}
}

func (v NullableServerCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
