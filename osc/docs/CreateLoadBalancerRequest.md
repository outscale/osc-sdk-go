# CreateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Listeners** | [**[]ListenerForCreation**](ListenerForCreation.md) | One or more listeners to create. | 
**LoadBalancerName** | **string** | The unique name of the load balancer (32 alphanumeric or hyphen characters maximum, but cannot start or end with a hyphen). | 
**LoadBalancerType** | **string** | The type of load balancer: &#x60;internet-facing&#x60; or &#x60;internal&#x60;. Use this parameter only for load balancers in a Net. | [optional] 
**SecurityGroups** | **[]string** | (Net only) One or more IDs of security groups you want to assign to the load balancer. If not specified, the default security group of the Net is assigned to the load balancer. | [optional] 
**Subnets** | **[]string** | (Net only) The ID of the Subnet in which you want to create the load balancer. Regardless of this Subnet, the load balancer can distribute traffic to all Subnets. This parameter is required in a Net. | [optional] 
**SubregionNames** | **[]string** | (public Cloud only) The Subregion in which you want to create the load balancer. Regardless of this Subregion, the load balancer can distribute traffic to all Subregions. This parameter is required in the public Cloud. | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags assigned to the load balancer. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


