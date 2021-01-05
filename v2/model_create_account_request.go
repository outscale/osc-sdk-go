/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateAccountRequest struct for CreateAccountRequest
type CreateAccountRequest struct {
	// The city of the account owner.
	City string `json:"City"`
	// The name of the company for the account.
	CompanyName string `json:"CompanyName"`
	// The country of the account owner.
	Country string `json:"Country"`
	// The ID of the customer. It must be 8 digits.
	CustomerId string `json:"CustomerId"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The email address for the account.
	Email string `json:"Email"`
	// The first name of the account owner.
	FirstName string `json:"FirstName"`
	// The job title of the account owner.
	JobTitle *string `json:"JobTitle,omitempty"`
	// The last name of the account owner.
	LastName string `json:"LastName"`
	// The mobile phone number of the account owner.
	MobileNumber *string `json:"MobileNumber,omitempty"`
	// The landline phone number of the account owner.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The state/province of the account.
	StateProvince *string `json:"StateProvince,omitempty"`
	// The value added tax (VAT) number for the account.
	VatNumber *string `json:"VatNumber,omitempty"`
	// The ZIP code of the city.
	ZipCode string `json:"ZipCode"`
}

// NewCreateAccountRequest instantiates a new CreateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountRequest(city string, companyName string, country string, customerId string, email string, firstName string, lastName string, zipCode string, ) *CreateAccountRequest {
	this := CreateAccountRequest{}
	this.City = city
	this.CompanyName = companyName
	this.Country = country
	this.CustomerId = customerId
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.ZipCode = zipCode
	return &this
}

// NewCreateAccountRequestWithDefaults instantiates a new CreateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountRequestWithDefaults() *CreateAccountRequest {
	this := CreateAccountRequest{}
	return &this
}

// GetCity returns the City field value
func (o *CreateAccountRequest) GetCity() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *CreateAccountRequest) SetCity(v string) {
	o.City = v
}

// GetCompanyName returns the CompanyName field value
func (o *CreateAccountRequest) GetCompanyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetCompanyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *CreateAccountRequest) SetCompanyName(v string) {
	o.CompanyName = v
}

// GetCountry returns the Country field value
func (o *CreateAccountRequest) GetCountry() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *CreateAccountRequest) SetCountry(v string) {
	o.Country = v
}

// GetCustomerId returns the CustomerId field value
func (o *CreateAccountRequest) GetCustomerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CreateAccountRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateAccountRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetEmail returns the Email field value
func (o *CreateAccountRequest) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateAccountRequest) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *CreateAccountRequest) GetFirstName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *CreateAccountRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetJobTitle() string {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetJobTitleOk() (*string, bool) {
	if o == nil || o.JobTitle == nil {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasJobTitle() bool {
	if o != nil && o.JobTitle != nil {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *CreateAccountRequest) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetLastName returns the LastName field value
func (o *CreateAccountRequest) GetLastName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *CreateAccountRequest) SetLastName(v string) {
	o.LastName = v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetMobileNumber() string {
	if o == nil || o.MobileNumber == nil {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetMobileNumberOk() (*string, bool) {
	if o == nil || o.MobileNumber == nil {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasMobileNumber() bool {
	if o != nil && o.MobileNumber != nil {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *CreateAccountRequest) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CreateAccountRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetStateProvince() string {
	if o == nil || o.StateProvince == nil {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetStateProvinceOk() (*string, bool) {
	if o == nil || o.StateProvince == nil {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasStateProvince() bool {
	if o != nil && o.StateProvince != nil {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *CreateAccountRequest) SetStateProvince(v string) {
	o.StateProvince = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetVatNumber() string {
	if o == nil || o.VatNumber == nil {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetVatNumberOk() (*string, bool) {
	if o == nil || o.VatNumber == nil {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasVatNumber() bool {
	if o != nil && o.VatNumber != nil {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *CreateAccountRequest) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetZipCode returns the ZipCode field value
func (o *CreateAccountRequest) GetZipCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetZipCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ZipCode, true
}

// SetZipCode sets field value
func (o *CreateAccountRequest) SetZipCode(v string) {
	o.ZipCode = v
}

func (o CreateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["City"] = o.City
	}
	if true {
		toSerialize["CompanyName"] = o.CompanyName
	}
	if true {
		toSerialize["Country"] = o.Country
	}
	if true {
		toSerialize["CustomerId"] = o.CustomerId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Email"] = o.Email
	}
	if true {
		toSerialize["FirstName"] = o.FirstName
	}
	if o.JobTitle != nil {
		toSerialize["JobTitle"] = o.JobTitle
	}
	if true {
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
	if true {
		toSerialize["ZipCode"] = o.ZipCode
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAccountRequest struct {
	value *CreateAccountRequest
	isSet bool
}

func (v NullableCreateAccountRequest) Get() *CreateAccountRequest {
	return v.value
}

func (v *NullableCreateAccountRequest) Set(val *CreateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountRequest(val *CreateAccountRequest) *NullableCreateAccountRequest {
	return &NullableCreateAccountRequest{value: val, isSet: true}
}

func (v NullableCreateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


