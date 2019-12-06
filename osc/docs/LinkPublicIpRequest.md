# LinkPublicIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRelink** | **bool** | - If &#x60;true&#x60;, allows the EIP to be associated with the VM or NIC that you specify even if it is already associated with another VM or NIC.&lt;br /&gt; - If &#x60;false&#x60;, prevents the EIP from being associated with the VM or NIC that you specify if it is already associated with another VM or NIC.&lt;br /&gt;&lt;br /&gt;  (By default, &#x60;true&#x60; in the public Cloud, &#x60;false&#x60; in a Net.) | [optional] 
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | **string** | (Net only) The ID of the NIC. This parameter is required if the VM has more than one NIC attached. Otherwise, you need to specify the &#x60;VmId&#x60; parameter instead. You cannot specify both parameters at the same time. | [optional] 
**PrivateIp** | **string** | (Net only) The primary or secondary private IP address of the specified NIC. By default, the primary private IP address. | [optional] 
**PublicIp** | **string** | The EIP. In the public Cloud, this parameter is required. | [optional] 
**PublicIpId** | **string** | The allocation ID of the EIP. In a Net, this parameter is required. | [optional] 
**VmId** | **string** | The ID of the VM.&lt;br /&gt; - In the public Cloud, this parameter is required.&lt;br /&gt; - In a Net, this parameter is required if the VM has only one NIC. Otherwise, you need to specify the &#x60;NicId&#x60; parameter instead. You cannot specify both parameters at the same time. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


