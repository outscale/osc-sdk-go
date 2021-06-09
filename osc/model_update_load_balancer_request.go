/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
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
	// (Net only) One or more IDs of security groups you want to assign to the load balancer. You need to specify the already assigned security groups that you want to keep along with the new ones you are assigning. If the list is empty, the default security group of the Net is assigned to the load balancer.
	SecurityGroups []string `json:"SecurityGroups,omitempty"`
	// The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat). If this parameter is specified, you must also specify the `LoadBalancerPort` parameter.
	ServerCertificateId string `json:"ServerCertificateId,omitempty"`
}
