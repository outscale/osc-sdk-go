# CreateClientGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsn** | **int32** | An unsigned 32-bits Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find out the path to your client gateway through the Internet network. The integer must be within the [0;4294967295] range. By default, 65000. | 
**ConnectionType** | **string** | The communication protocol used to establish tunnel with your client gateway (only &#x60;ipsec.1&#x60; is supported). | 
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**PublicIp** | **string** | The public fixed IPv4 address of your client gateway. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


