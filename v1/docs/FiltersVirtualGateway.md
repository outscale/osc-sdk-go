# FiltersVirtualGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionTypes** | **[]string** | The types of the virtual gateways (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**LinkNetIds** | **[]string** | The IDs of the Nets the virtual gateways are attached to. | [optional] 
**LinkStates** | **[]string** | The current states of the attachments between the virtual gateways and the Nets (&#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60; \\| &#x60;detached&#x60;). | [optional] 
**States** | **[]string** | The states of the virtual gateways (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the virtual gateways. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the virtual gateways. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the virtual gateways, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 
**VirtualGatewayIds** | **[]string** | The IDs of the virtual gateways. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


