/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// UpdateAccountRequest struct for UpdateAccountRequest
type UpdateAccountRequest struct {
	// The new city of the account owner.
	City string `json:"City,omitempty"`
	// The new name of the company for the account.
	CompanyName string `json:"CompanyName,omitempty"`
	// The new country of the account owner.
	Country string `json:"Country,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The new email address for the account.
	Email string `json:"Email,omitempty"`
	// The new first name of the account owner.
	FirstName string `json:"FirstName,omitempty"`
	// The new job title of the account owner.
	JobTitle string `json:"JobTitle,omitempty"`
	// The new last name of the account owner.
	LastName string `json:"LastName,omitempty"`
	// The new mobile phone number of the account owner.
	MobileNumber string `json:"MobileNumber,omitempty"`
	// The new landline phone number of the account owner.
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	// The new state/province of the account owner.
	StateProvince string `json:"StateProvince,omitempty"`
	// The new value added tax (VAT) number for the account.
	VatNumber string `json:"VatNumber,omitempty"`
	// The new ZIP code of the city.
	ZipCode string `json:"ZipCode,omitempty"`
}
