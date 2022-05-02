# UpdateVmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDeviceMappings** | [**[]BlockDeviceMappingVmUpdate**](BlockDeviceMappingVmUpdate.md) | One or more block device mappings of the VM. | [optional] 
**BsuOptimized** | **bool** | This parameter is not available. It is present in our API for the sake of historical compatibility with AWS. | [optional] 
**DeletionProtection** | **bool** | If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**IsSourceDestChecked** | **bool** | (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**KeypairName** | **string** | The name of the keypair.&lt;br /&gt; To complete the replacement, manually replace the old public key with the new public key in the ~/.ssh/authorized_keys file located in the VM. Restart the VM to apply the change. | [optional] 
**Performance** | **string** | The performance of the VM (&#x60;medium&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). | [optional] 
**SecurityGroupIds** | **[]string** | One or more IDs of security groups for the VM. | [optional] 
**UserData** | **string** | The Base64-encoded MIME user data, limited to 500 kibibytes (KiB). | [optional] 
**VmId** | **string** | The ID of the VM. | 
**VmInitiatedShutdownBehavior** | **string** | The VM behavior when you stop it. If set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is terminated. | [optional] 
**VmType** | **string** | The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


