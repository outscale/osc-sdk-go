# CreateVmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDeviceMappings** | [**[]BlockDeviceMappingVmCreation**](BlockDeviceMappingVmCreation.md) | One or more block device mappings. | [optional] 
**BootOnCreation** | **bool** | By default or if &#x60;true&#x60;, the VM is started on creation. If &#x60;false&#x60;, the VM is stopped on creation. | [optional] 
**BsuOptimized** | **bool** | If &#x60;true&#x60;, the VM is created with optimized BSU I/O. | [optional] 
**ClientToken** | **string** | A unique identifier which enables you to manage the idempotency. | [optional] 
**DeletionProtection** | **bool** | If &#x60;true&#x60;, you cannot terminate the VM using Cockpit, the CLI or the API. If &#x60;false&#x60;, you can. | [optional] 
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | **string** | The ID of the OMI used to create the VM. You can find the list of OMIs by calling the [ReadImages](#readimages) method. | 
**KeypairName** | **string** | The name of the keypair. | [optional] 
**MaxVmsCount** | **int32** | The maximum number of VMs you want to create. If all the VMs cannot be created, the largest possible number of VMs above MinVmsCount is created. | [optional] 
**MinVmsCount** | **int32** | The minimum number of VMs you want to create. If this number of VMs cannot be created, no VMs are created. | [optional] 
**Nics** | [**[]NicForVmCreation**](NicForVmCreation.md) | One or more NICs. If you specify this parameter, you must define one NIC as the primary network interface of the VM with &#x60;0&#x60; as its device number. | [optional] 
**Performance** | **string** | The performance of the VM (&#x60;standard&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). | [optional] [default to PERFORMANCE_HIGH]
**Placement** | [**Placement**](Placement.md) |  | [optional] 
**PrivateIps** | **[]string** | One or more private IP addresses of the VM. | [optional] 
**SecurityGroupIds** | **[]string** | One or more IDs of security group for the VMs. | [optional] 
**SecurityGroups** | **[]string** | One or more names of security groups for the VMs. | [optional] 
**SubnetId** | **string** | The ID of the Subnet in which you want to create the VM. | [optional] 
**UserData** | **string** | Data or script used to add a specific configuration to the VM. It must be base64-encoded. | [optional] 
**VmInitiatedShutdownBehavior** | **string** | The VM behavior when you stop it. By default or if set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is terminated. | [optional] 
**VmType** | **string** | The type of VM (&#x60;tinav2.c1r2&#x60; by default).&lt;br /&gt; For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


