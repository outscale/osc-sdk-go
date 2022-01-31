/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// UpdateApiAccessPolicyRequest struct for UpdateApiAccessPolicyRequest
type UpdateApiAccessPolicyRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The maximum possible lifetime for your access keys, in seconds (between `0` and `3153600000`, both included). By default or if set to `O`, your access keys can have unlimited lifetimes. Otherwise, all your access keys must have an expiration date. This value must be greater than the remaining lifetime of each access key of your account.
	MaxAccessKeyExpirationSeconds int64 `json:"MaxAccessKeyExpirationSeconds"`
	// If true, a trusted session is activated, provided that you specify the `MaxAccessKeyExpirationSeconds` parameter with a value greater than `0`.
	RequireTrustedEnv bool `json:"RequireTrustedEnv"`
}
