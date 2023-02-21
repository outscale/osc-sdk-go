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

// CreateLoadBalancerRequest struct for CreateLoadBalancerRequest
type CreateLoadBalancerRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// One or more listeners to create.
	Listeners []ListenerForCreation `json:"Listeners"`
	// The unique name of the load balancer (32 alphanumeric or hyphen characters maximum, but cannot start or end with a hyphen).
	LoadBalancerName string `json:"LoadBalancerName"`
	// The type of load balancer: `internet-facing` or `internal`. Use this parameter only for load balancers in a Net.
	LoadBalancerType string `json:"LoadBalancerType,omitempty"`
	// (internet-facing only) The public IP you want to associate with the load balancer. If not specified, a public IP owned by 3DS OUTSCALE is associated.
	PublicIp string `json:"PublicIp,omitempty"`
	// (Net only) One or more IDs of security groups you want to assign to the load balancer. If not specified, the default security group of the Net is assigned to the load balancer.
	SecurityGroups []string `json:"SecurityGroups,omitempty"`
	// (Net only) The ID of the Subnet in which you want to create the load balancer. Regardless of this Subnet, the load balancer can distribute traffic to all Subnets. This parameter is required in a Net.
	Subnets []string `json:"Subnets,omitempty"`
	// (public Cloud only) The Subregion in which you want to create the load balancer. Regardless of this Subregion, the load balancer can distribute traffic to all Subregions. This parameter is required in the public Cloud.
	SubregionNames []string `json:"SubregionNames,omitempty"`
	// One or more tags assigned to the load balancer.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
