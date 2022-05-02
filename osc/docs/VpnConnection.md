# VpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGatewayConfiguration** | **string** | Example configuration for the client gateway. | [optional] 
**ClientGatewayId** | **string** | The ID of the client gateway used on the client end of the connection. | [optional] 
**ConnectionType** | **string** | The type of VPN connection (always &#x60;ipsec.1&#x60;). | [optional] 
**Routes** | [**[]RouteLight**](RouteLight.md) | Information about one or more static routes associated with the VPN connection, if any. | [optional] 
**State** | **string** | The state of the VPN connection (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**StaticRoutesOnly** | **bool** | If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute). | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the VPN connection. | [optional] 
**VgwTelemetries** | [**[]VgwTelemetry**](VgwTelemetry.md) | Information about the current state of one or more of the VPN tunnels. | [optional] 
**VirtualGatewayId** | **string** | The ID of the virtual gateway used on the OUTSCALE end of the connection. | [optional] 
**VpnConnectionId** | **string** | The ID of the VPN connection. | [optional] 
**VpnOptions** | [**VpnOptions**](VpnOptions.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


