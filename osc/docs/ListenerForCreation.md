# ListenerForCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendPort** | **int32** | The port on which the back-end VM is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**BackendProtocol** | **string** | The protocol for routing traffic to back-end VMs (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60;). | [optional] 
**LoadBalancerPort** | **int32** | The port on which the load balancer is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**LoadBalancerProtocol** | **string** | The routing protocol (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60;). | 
**ServerCertificateId** | **string** | The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers &gt; OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


