/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
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
	// The date after which you want to retrieve logged calls, in ISO 8601 format (for example, `2020-06-14`). By default, this date is set to 2 days ago.
	QueryDateAfter string `json:"QueryDateAfter,omitempty"`
	// The date before which you want to retrieve logged calls, in ISO 8601 format (for example, `2020-06-30`). By default, this date is set to now.
	QueryDateBefore string `json:"QueryDateBefore,omitempty"`
	// The IP addresses used for the logged calls.
	QueryIpAddresses []string `json:"QueryIpAddresses,omitempty"`
	// The user agents of the HTTP requests of the logged calls.
	QueryUserAgents []string `json:"QueryUserAgents,omitempty"`
	// The request IDs provided in the responses of the logged calls.
	RequestIds []string `json:"RequestIds,omitempty"`
	// The HTTP status codes of the logged calls.
	ResponseStatusCodes []int32 `json:"ResponseStatusCodes,omitempty"`
}
