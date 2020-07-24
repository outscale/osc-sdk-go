# Listener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendPort** | **int32** | The port on which the back-end VM is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | [optional] 
**BackendProtocol** | **string** | The protocol for routing traffic to back-end VMs (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | [optional] 
**LoadBalancerPort** | **int32** | The port on which the load balancer is listening (between 1 and &#x60;65535&#x60;, both included). | [optional] 
**LoadBalancerProtocol** | **string** | The routing protocol (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | [optional] 
**PolicyNames** | **[]string** | The names of the policies. If there are no policies enabled, the list is empty. | [optional] 
**ServerCertificateId** | **string** | The ID of the server certificate. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


