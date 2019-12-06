# ClientGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsn** | **int32** | An unsigned 32-bits Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find out the path to your client gateway through the Internet network. | [optional] 
**ClientGatewayId** | **string** | The ID of the client gateway. | [optional] 
**ConnectionType** | **string** | The type of communication tunnel used by the client gateway (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**PublicIp** | **string** | The public IPv4 address of the client gateway (must be a fixed address into a NATed network). | [optional] 
**State** | **string** | The state of the client gateway (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the client gateway. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


