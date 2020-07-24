# FiltersVpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsns** | **[]int32** | The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections. | [optional] 
**ClientGatewayIds** | **[]string** | The IDs of the client gateways. | [optional] 
**ConnectionTypes** | **[]string** | The types of the VPN connections (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**RouteDestinationIpRanges** | **[]string** | The destination IP ranges. | [optional] 
**States** | **[]string** | The states of the VPN connections (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**StaticRoutesOnly** | **bool** | If &#x60;false&#x60;, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If &#x60;true&#x60;, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute). | [optional] 
**VirtualGatewayIds** | **[]string** | The IDs of the virtual gateways. | [optional] 
**VpnConnectionIds** | **[]string** | The IDs of the VPN connections. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


