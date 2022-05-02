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

// UpdateLoadBalancerRequest struct for UpdateLoadBalancerRequest
type UpdateLoadBalancerRequest struct {
	AccessLog AccessLog `json:"AccessLog,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun      bool        `json:"DryRun,omitempty"`
	HealthCheck HealthCheck `json:"HealthCheck,omitempty"`
	// The name of the load balancer.
	LoadBalancerName string `json:"LoadBalancerName"`
	// The port on which the load balancer is listening (between `1` and `65535`, both included). This parameter is required if you want to update the server certificate.
	LoadBalancerPort int32 `json:"LoadBalancerPort,omitempty"`
	// The name of the policy you want to enable for the listener.
	PolicyNames []string `json:"PolicyNames,omitempty"`
	// (internet-facing only) The public IP you want to associate with the load balancer. The former public IP of the load balancer is then disassociated. If you specify an empty string and the former public IP belonged to you, it is disassociated and replaced by a public IP owned by 3DS OUTSCALE.
	PublicIp string `json:"PublicIp,omitempty"`
	// (Net only) One or more IDs of security groups you want to assign to the load balancer. You need to specify the already assigned security groups that you want to keep along with the new ones you are assigning. If the list is empty, the default security group of the Net is assigned to the load balancer.
	SecurityGroups []string `json:"SecurityGroups,omitempty"`
	// The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns). If this parameter is specified, you must also specify the `LoadBalancerPort` parameter.
	ServerCertificateId string `json:"ServerCertificateId,omitempty"`
}
