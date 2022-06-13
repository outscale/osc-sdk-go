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

// Quota Information about the quota.
type Quota struct {
	// The account ID of the owner of the quotas.
	AccountId string `json:"AccountId,omitempty"`
	// The description of the quota.
	Description string `json:"Description,omitempty"`
	// The maximum value of the quota for the OUTSCALE user account (if there is no limit, `0`).
	MaxValue int32 `json:"MaxValue,omitempty"`
	// The unique name of the quota.
	Name string `json:"Name,omitempty"`
	// The group name of the quota.
	QuotaCollection string `json:"QuotaCollection,omitempty"`
	// The description of the quota.
	ShortDescription string `json:"ShortDescription,omitempty"`
	// The limit value currently used by the OUTSCALE user account.
	UsedValue int32 `json:"UsedValue,omitempty"`
}
