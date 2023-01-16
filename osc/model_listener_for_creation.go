/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// ListenerForCreation Information about the listener to create.
type ListenerForCreation struct {
	// The port on which the back-end VM is listening (between `1` and `65535`, both included).
	BackendPort int32 `json:"BackendPort"`
	// The protocol for routing traffic to back-end VMs (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL`).
	BackendProtocol string `json:"BackendProtocol,omitempty"`
	// The port on which the load balancer is listening (between `1` and `65535`, both included).
	LoadBalancerPort int32 `json:"LoadBalancerPort"`
	// The routing protocol (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL`).
	LoadBalancerProtocol string `json:"LoadBalancerProtocol"`
	// The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).
	ServerCertificateId string `json:"ServerCertificateId,omitempty"`
}
