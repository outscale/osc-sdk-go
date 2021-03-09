# FiltersVpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsns** | **[]int32** | The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections. | [optional] 
**ClientGatewayIds** | **[]string** | The IDs of the client gateways. | [optional] 
**ConnectionTypes** | **[]string** | The types of the VPN connections (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**RouteDestinationIpRanges** | **[]string** | The destination IP ranges. | [optional] 
**States** | **[]string** | The states of the VPN connections (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**StaticRoutesOnly** | **bool** | If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute). | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the VPN connections. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the VPN connections. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the VPN connections, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**VirtualGatewayIds** | **[]string** | The IDs of the virtual gateways. | [optional] 
**VpnConnectionIds** | **[]string** | The IDs of the VPN connections. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


