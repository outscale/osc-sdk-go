# FiltersClientGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsns** | **[]int32** | The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections. | [optional] 
**ClientGatewayIds** | **[]string** | The IDs of the client gateways. | [optional] 
**ConnectionTypes** | **[]string** | The types of communication tunnels used by the client gateways (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**PublicIps** | **[]string** | The public IPv4 addresses of the client gateways. | [optional] 
**States** | **[]string** | The states of the client gateways (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the client gateways. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the client gateways. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the client gateways, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


