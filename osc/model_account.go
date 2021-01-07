/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// Account Information about the account.
type Account struct {
	// The ID of the account.
	AccountId string `json:"AccountId,omitempty"`
	// The city of the account owner.
	City string `json:"City,omitempty"`
	// The name of the company for the account.
	CompanyName string `json:"CompanyName,omitempty"`
	// The country of the account owner.
	Country string `json:"Country,omitempty"`
	// The ID of the customer.
	CustomerId string `json:"CustomerId,omitempty"`
	// The email address for the account.
	Email string `json:"Email,omitempty"`
	// The first name of the account owner.
	FirstName string `json:"FirstName,omitempty"`
	// The job title of the account owner.
	JobTitle string `json:"JobTitle,omitempty"`
	// The last name of the account owner.
	LastName string `json:"LastName,omitempty"`
	// The mobile phone number of the account owner.
	MobileNumber string `json:"MobileNumber,omitempty"`
	// The landline phone number of the account owner.
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	// The state/province of the account.
	StateProvince string `json:"StateProvince,omitempty"`
	// The value added tax (VAT) number for the account.
	VatNumber string `json:"VatNumber,omitempty"`
	// The ZIP code of the city.
	ZipCode string `json:"ZipCode,omitempty"`
}
