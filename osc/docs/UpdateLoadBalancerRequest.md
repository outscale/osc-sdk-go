# UpdateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLog** | [**AccessLog**](AccessLog.md) |  | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**HealthCheck** | [**HealthCheck**](HealthCheck.md) |  | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer. | 
**LoadBalancerPort** | **int32** | The port on which the load balancer is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). This parameter is required if you want to update the server certificate. | [optional] 
**PolicyNames** | **[]string** | The name of the policy you want to enable for the listener. | [optional] 
**SecurityGroups** | **[]string** | (Net only) One or more IDs of security groups you want to assign to the load balancer. You need to specify the already assigned security groups that you want to keep along with the new ones you are assigning. If the list is empty, the default security group of the Net is assigned to the load balancer. | [optional] 
**ServerCertificateId** | **string** | The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers &gt; Outscale Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat). If this parameter is specified, you must also specify the &#x60;LoadBalancerPort&#x60; parameter. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


