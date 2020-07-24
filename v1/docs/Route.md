# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationMethod** | **string** | The method used to create the route. | [optional] 
**DestinationIpRange** | **string** | The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24). | [optional] 
**DestinationServiceId** | **string** | The ID of the 3DS OUTSCALE service. | [optional] 
**GatewayId** | **string** | The ID of the Internet service or virtual gateway attached to the Net. | [optional] 
**NatServiceId** | **string** | The ID of a NAT service attached to the Net. | [optional] 
**NetAccessPointId** | **string** | The ID of the Net access point. | [optional] 
**NetPeeringId** | **string** | The ID of the Net peering connection. | [optional] 
**NicId** | **string** | The ID of the NIC. | [optional] 
**State** | **string** | The state of a route in the route table (&#x60;active&#x60; \\| &#x60;blackhole&#x60;). The &#x60;blackhole&#x60; state indicates that the target of the route is not available. | [optional] 
**VmAccountId** | **string** | The account ID of the owner of the VM. | [optional] 
**VmId** | **string** | The ID of a VM specified in a route in the table. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


