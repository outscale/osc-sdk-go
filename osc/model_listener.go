/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// Listener Information about the listener.
type Listener struct {
	// The port on which the back-end VM is listening (between `1` and `65535`, both included).
	BackendPort int32 `json:"BackendPort,omitempty"`
	// The protocol for routing traffic to back-end VMs (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL`).
	BackendProtocol string `json:"BackendProtocol,omitempty"`
	// The port on which the load balancer is listening (between 1 and `65535`, both included).
	LoadBalancerPort int32 `json:"LoadBalancerPort,omitempty"`
	// The routing protocol (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL`).
	LoadBalancerProtocol string `json:"LoadBalancerProtocol,omitempty"`
	// The names of the policies. If there are no policies enabled, the list is empty.
	PolicyNames []string `json:"PolicyNames,omitempty"`
	// The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).
	ServerCertificateId string `json:"ServerCertificateId,omitempty"`
}
