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

// NetAccessPoint Information about the Net access point.
type NetAccessPoint struct {
	// The ID of the Net access point.
	NetAccessPointId string `json:"NetAccessPointId,omitempty"`
	// The ID of the Net with which the Net access point is associated.
	NetId string `json:"NetId,omitempty"`
	// The ID of the route tables associated with the Net access point.
	RouteTableIds []string `json:"RouteTableIds,omitempty"`
	// The name of the service with which the Net access point is associated.
	ServiceName string `json:"ServiceName,omitempty"`
	// The state of the Net access point (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State string `json:"State,omitempty"`
	// One or more tags associated with the Net access point.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
