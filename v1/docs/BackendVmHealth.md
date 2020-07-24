# BackendVmHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the state of the back-end VM. | [optional] 
**State** | **string** | The state of the back-end VM (&#x60;InService&#x60; \\| &#x60;OutOfService&#x60; \\| &#x60;Unknown&#x60;). | [optional] 
**StateReason** | **string** | Information about the cause of &#x60;OutOfService&#x60; VMs.&lt;br /&gt; Specifically, whether the cause is Elastic Load Balancing or the VM (&#x60;ELB&#x60; \\| &#x60;Instance&#x60; \\| &#x60;N/A&#x60;). | [optional] 
**VmId** | **string** | The ID of the back-end VM. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


