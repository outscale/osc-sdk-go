/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateAccountRequest struct for UpdateAccountRequest
type UpdateAccountRequest struct {
	// One or more additional email addresses for the account. These addresses are used for notifications only. If you already have a list of additional emails registered, you cannot add to it, only replace it. To remove all registered additional emails, specify an empty list.
	AdditionalEmails *[]string `json:"AdditionalEmails,omitempty"`
	// The new city of the account owner.
	City *string `json:"City,omitempty"`
	// The new name of the company for the account.
	CompanyName *string `json:"CompanyName,omitempty"`
	// The new country of the account owner.
	Country *string `json:"Country,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The main email address for the account. This address is used for your credentials and notifications.
	Email *string `json:"Email,omitempty"`
	// The new first name of the account owner.
	FirstName *string `json:"FirstName,omitempty"`
	// The new job title of the account owner.
	JobTitle *string `json:"JobTitle,omitempty"`
	// The new last name of the account owner.
	LastName *string `json:"LastName,omitempty"`
	// The new mobile phone number of the account owner.
	MobileNumber *string `json:"MobileNumber,omitempty"`
	// The new landline phone number of the account owner.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The new state/province of the account owner.
	StateProvince *string `json:"StateProvince,omitempty"`
	// The new value added tax (VAT) number for the account.
	VatNumber *string `json:"VatNumber,omitempty"`
	// The new ZIP code of the city.
	ZipCode *string `json:"ZipCode,omitempty"`
}

// NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountRequest() *UpdateAccountRequest {
	this := UpdateAccountRequest{}
	return &this
}

// NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest {
	this := UpdateAccountRequest{}
	return &this
}

// GetAdditionalEmails returns the AdditionalEmails field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetAdditionalEmails() []string {
	if o == nil || o.AdditionalEmails == nil {
		var ret []string
		return ret
	}
	return *o.AdditionalEmails
}

// GetAdditionalEmailsOk returns a tuple with the AdditionalEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetAdditionalEmailsOk() (*[]string, bool) {
	if o == nil || o.AdditionalEmails == nil {
		return nil, false
	}
	return o.AdditionalEmails, true
}

// HasAdditionalEmails returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasAdditionalEmails() bool {
	if o != nil && o.AdditionalEmails != nil {
		return true
	}

	return false
}

// SetAdditionalEmails gets a reference to the given []string and assigns it to the AdditionalEmails field.
func (o *UpdateAccountRequest) SetAdditionalEmails(v []string) {
	o.AdditionalEmails = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *UpdateAccountRequest) SetCity(v string) {
	o.City = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UpdateAccountRequest) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *UpdateAccountRequest) SetCountry(v string) {
	o.Country = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateAccountRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateAccountRequest) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UpdateAccountRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetJobTitle() string {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetJobTitleOk() (*string, bool) {
	if o == nil || o.JobTitle == nil {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasJobTitle() bool {
	if o != nil && o.JobTitle != nil {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *UpdateAccountRequest) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UpdateAccountRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetMobileNumber() string {
	if o == nil || o.MobileNumber == nil {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetMobileNumberOk() (*string, bool) {
	if o == nil || o.MobileNumber == nil {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasMobileNumber() bool {
	if o != nil && o.MobileNumber != nil {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *UpdateAccountRequest) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UpdateAccountRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetStateProvince() string {
	if o == nil || o.StateProvince == nil {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetStateProvinceOk() (*string, bool) {
	if o == nil || o.StateProvince == nil {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasStateProvince() bool {
	if o != nil && o.StateProvince != nil {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *UpdateAccountRequest) SetStateProvince(v string) {
	o.StateProvince = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetVatNumber() string {
	if o == nil || o.VatNumber == nil {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetVatNumberOk() (*string, bool) {
	if o == nil || o.VatNumber == nil {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasVatNumber() bool {
	if o != nil && o.VatNumber != nil {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *UpdateAccountRequest) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *UpdateAccountRequest) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o UpdateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalEmails != nil {
		toSerialize["AdditionalEmails"] = o.AdditionalEmails
	}
	if o.City != nil {
		toSerialize["City"] = o.City
	}
	if o.CompanyName != nil {
		toSerialize["CompanyName"] = o.CompanyName
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["FirstName"] = o.FirstName
	}
	if o.JobTitle != nil {
		toSerialize["JobTitle"] = o.JobTitle
	}
	if o.LastName != nil {
		toSerialize["LastName"] = o.LastName
	}
	if o.MobileNumber != nil {
		toSerialize["MobileNumber"] = o.MobileNumber
	}
	if o.PhoneNumber != nil {
		toSerialize["PhoneNumber"] = o.PhoneNumber
	}
	if o.StateProvince != nil {
		toSerialize["StateProvince"] = o.StateProvince
	}
	if o.VatNumber != nil {
		toSerialize["VatNumber"] = o.VatNumber
	}
	if o.ZipCode != nil {
		toSerialize["ZipCode"] = o.ZipCode
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAccountRequest struct {
	value *UpdateAccountRequest
	isSet bool
}

func (v NullableUpdateAccountRequest) Get() *UpdateAccountRequest {
	return v.value
}

func (v *NullableUpdateAccountRequest) Set(val *UpdateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountRequest(val *UpdateAccountRequest) *NullableUpdateAccountRequest {
	return &NullableUpdateAccountRequest{value: val, isSet: true}
}

func (v NullableUpdateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
