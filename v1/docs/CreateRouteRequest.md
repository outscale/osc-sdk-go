# CreateRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpRange** | **string** | The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24). | 
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**GatewayId** | **string** | The ID of an Internet service or virtual gateway attached to your Net. | [optional] 
**NatServiceId** | **string** | The ID of a NAT service. | [optional] 
**NetPeeringId** | **string** | The ID of a Net peering connection. | [optional] 
**NicId** | **string** | The ID of a NIC. | [optional] 
**RouteTableId** | **string** | The ID of the route table for which you want to create a route. | 
**VmId** | **string** | The ID of a NAT VM in your Net (attached to exactly one NIC). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


