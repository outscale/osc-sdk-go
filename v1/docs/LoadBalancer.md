# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLog** | [**AccessLog**](AccessLog.md) |  | [optional] 
**ApplicationStickyCookiePolicies** | [**[]ApplicationStickyCookiePolicy**](ApplicationStickyCookiePolicy.md) | The stickiness policies defined for the load balancer. | [optional] 
**BackendVmIds** | **[]string** | One or more IDs of back-end VMs for the load balancer. | [optional] 
**DnsName** | **string** | The DNS name of the load balancer. | [optional] 
**HealthCheck** | [**HealthCheck**](HealthCheck.md) |  | [optional] 
**Listeners** | [**[]Listener**](Listener.md) | The listeners for the load balancer. | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer. | [optional] 
**LoadBalancerStickyCookiePolicies** | [**[]LoadBalancerStickyCookiePolicy**](LoadBalancerStickyCookiePolicy.md) | The policies defined for the load balancer. | [optional] 
**LoadBalancerType** | **string** | The type of load balancer. Valid only for load balancers in a Net.&lt;br /&gt; If &#x60;LoadBalancerType&#x60; is &#x60;internet-facing&#x60;, the load balancer has a public DNS name that resolves to a public IP address.&lt;br /&gt; If &#x60;LoadBalancerType&#x60; is &#x60;internal&#x60;, the load balancer has a public DNS name that resolves to a private IP address. | [optional] 
**NetId** | **string** | The ID of the Net for the load balancer. | [optional] 
**SecurityGroups** | **[]string** | One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net. | [optional] 
**SourceSecurityGroup** | [**SourceSecurityGroup**](SourceSecurityGroup.md) |  | [optional] 
**Subnets** | **[]string** | The IDs of the Subnets for the load balancer. | [optional] 
**SubregionNames** | **[]string** | One or more names of Subregions for the load balancer. | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the load balancer. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


