/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// HealthCheck Information about the health check configuration.
type HealthCheck struct {
	// The number of seconds between two pings (between `5` and `600` both included).
	CheckInterval int32 `json:"CheckInterval"`
	// The number of consecutive successful pings before considering the VM as healthy (between `2` and `10` both included).
	HealthyThreshold int32 `json:"HealthyThreshold"`
	// The path for HTTP or HTTPS requests.
	Path string `json:"Path,omitempty"`
	// The port number (between `1` and `65535`, both included).
	Port int32 `json:"Port"`
	// The protocol for the URL of the VM (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	Protocol string `json:"Protocol"`
	// The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between `2` and `60` both included).
	Timeout int32 `json:"Timeout"`
	// The number of consecutive failed pings before considering the VM as unhealthy (between `2` and `10` both included).
	UnhealthyThreshold int32 `json:"UnhealthyThreshold"`
}