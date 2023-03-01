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

// LoadBalancer Information about the load balancer.
type LoadBalancer struct {
	AccessLog AccessLog `json:"AccessLog,omitempty"`
	// The stickiness policies defined for the load balancer.
	ApplicationStickyCookiePolicies []ApplicationStickyCookiePolicy `json:"ApplicationStickyCookiePolicies,omitempty"`
	// One or more public IPs of back-end VMs.
	BackendIps []string `json:"BackendIps,omitempty"`
	// One or more IDs of back-end VMs for the load balancer.
	BackendVmIds []string `json:"BackendVmIds,omitempty"`
	// The DNS name of the load balancer.
	DnsName     string      `json:"DnsName,omitempty"`
	HealthCheck HealthCheck `json:"HealthCheck,omitempty"`
	// The listeners for the load balancer.
	Listeners []Listener `json:"Listeners,omitempty"`
	// The name of the load balancer.
	LoadBalancerName string `json:"LoadBalancerName,omitempty"`
	// The policies defined for the load balancer.
	LoadBalancerStickyCookiePolicies []LoadBalancerStickyCookiePolicy `json:"LoadBalancerStickyCookiePolicies,omitempty"`
	// The type of load balancer. Valid only for load balancers in a Net.<br /> If `LoadBalancerType` is `internet-facing`, the load balancer has a public DNS name that resolves to a public IP.<br /> If `LoadBalancerType` is `internal`, the load balancer has a public DNS name that resolves to a private IP.
	LoadBalancerType string `json:"LoadBalancerType,omitempty"`
	// The ID of the Net for the load balancer.
	NetId string `json:"NetId,omitempty"`
	// (internet-facing only) The public IP associated with the load balancer.
	PublicIp string `json:"PublicIp,omitempty"`
	// Whether secure cookies are enabled for the load balancer.
	SecuredCookies bool `json:"SecuredCookies,omitempty"`
	// One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.
	SecurityGroups      []string            `json:"SecurityGroups,omitempty"`
	SourceSecurityGroup SourceSecurityGroup `json:"SourceSecurityGroup,omitempty"`
	// The ID of the Subnet in which the load balancer was created.
	Subnets []string `json:"Subnets,omitempty"`
	// The ID of the Subregion in which the load balancer was created.
	SubregionNames []string `json:"SubregionNames,omitempty"`
	// One or more tags associated with the load balancer.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
