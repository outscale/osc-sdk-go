/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FiltersApiLog One or more filters.
type FiltersApiLog struct {
	// The access keys used for the logged calls.
	QueryAccessKeys []string `json:"QueryAccessKeys,omitempty"`
	// The names of the APIs of the logged calls (always `oapi` for the OUTSCALE API).
	QueryApiNames []string `json:"QueryApiNames,omitempty"`
	// The names of the logged calls.
	QueryCallNames []string `json:"QueryCallNames,omitempty"`
	// The date after which you want to retrieve logged calls, in ISO 8601 format (for example, `2020-06-14`). By default, this date is set to 48 hours before the `QueryDateBefore` parameter value.
	QueryDateAfter string `json:"QueryDateAfter,omitempty"`
	// The date before which you want to retrieve logged calls, in ISO 8601 format (for example, `2020-06-30`). By default, this date is set to now, or 48 hours after the `QueryDateAfter` parameter value.
	QueryDateBefore string `json:"QueryDateBefore,omitempty"`
	// The IPs used for the logged calls.
	QueryIpAddresses []string `json:"QueryIpAddresses,omitempty"`
	// The user agents of the HTTP requests of the logged calls.
	QueryUserAgents []string `json:"QueryUserAgents,omitempty"`
	// The request IDs provided in the responses of the logged calls.
	RequestIds []string `json:"RequestIds,omitempty"`
	// The HTTP status codes of the logged calls.
	ResponseStatusCodes []int32 `json:"ResponseStatusCodes,omitempty"`
}
