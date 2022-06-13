/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// CreateAccountRequest struct for CreateAccountRequest
type CreateAccountRequest struct {
	// One or more additional email addresses for the account. These addresses are used for notifications only. If you already have a list of additional emails registered, you cannot add to it, only replace it. To remove all registered additional emails, specify an empty list.
	AdditionalEmails []string `json:"AdditionalEmails,omitempty"`
	// The city of the account owner.
	City string `json:"City"`
	// The name of the company for the account.
	CompanyName string `json:"CompanyName"`
	// The country of the account owner.
	Country string `json:"Country"`
	// The ID of the customer. It must be 8 digits.
	CustomerId string `json:"CustomerId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The main email address for the account. This address is used for your credentials and notifications.
	Email string `json:"Email"`
	// The first name of the account owner.
	FirstName string `json:"FirstName"`
	// The job title of the account owner.
	JobTitle string `json:"JobTitle,omitempty"`
	// The last name of the account owner.
	LastName string `json:"LastName"`
	// The mobile phone number of the account owner.
	MobileNumber string `json:"MobileNumber,omitempty"`
	// The landline phone number of the account owner.
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	// The state/province of the account.
	StateProvince string `json:"StateProvince,omitempty"`
	// The value added tax (VAT) number for the account.
	VatNumber string `json:"VatNumber,omitempty"`
	// The ZIP code of the city.
	ZipCode string `json:"ZipCode"`
}
