# CreateVmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDeviceMappings** | [**[]BlockDeviceMappingVmCreation**](BlockDeviceMappingVmCreation.md) | One or more block device mappings. | [optional] 
**BootOnCreation** | **bool** | By default or if true, the VM is started on creation. If false, the VM is stopped on creation. | [optional] [default to true]
**BsuOptimized** | **bool** | This parameter is not available. It is present in our API for the sake of historical compatibility with AWS. | [optional] 
**ClientToken** | **string** | A unique identifier which enables you to manage the idempotency. | [optional] 
**DeletionProtection** | **bool** | If true, you cannot delete the VM unless you change this parameter back to false. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | **string** | The ID of the OMI used to create the VM. You can find the list of OMIs by calling the [ReadImages](#readimages) method. | 
**KeypairName** | **string** | The name of the keypair. | [optional] 
**MaxVmsCount** | **int32** | The maximum number of VMs you want to create. If all the VMs cannot be created, the largest possible number of VMs above MinVmsCount is created. | [optional] 
**MinVmsCount** | **int32** | The minimum number of VMs you want to create. If this number of VMs cannot be created, no VMs are created. | [optional] 
**NestedVirtualization** | **bool** | (dedicated tenancy only) If true, nested virtualization is enabled. If false, it is disabled. | [optional] [default to false]
**Nics** | [**[]NicForVmCreation**](NicForVmCreation.md) | One or more NICs. If you specify this parameter, you must not specify the &#x60;SubnetId&#x60; and &#x60;SubregionName&#x60; parameters. You also must define one NIC as the primary network interface of the VM with &#x60;0&#x60; as its device number. | [optional] 
**Performance** | **string** | The performance of the VM (&#x60;medium&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). By default, &#x60;high&#x60;. This parameter is ignored if you specify a performance flag directly in the &#x60;VmType&#x60; parameter. | [optional] [default to PERFORMANCE_HIGH]
**Placement** | [**Placement**](Placement.md) |  | [optional] 
**PrivateIps** | **[]string** | One or more private IPs of the VM. | [optional] 
**SecurityGroupIds** | **[]string** | One or more IDs of security group for the VMs. | [optional] 
**SecurityGroups** | **[]string** | One or more names of security groups for the VMs. | [optional] 
**SubnetId** | **string** | The ID of the Subnet in which you want to create the VM. If you specify this parameter, you must not specify the &#x60;Nics&#x60; parameter. | [optional] 
**UserData** | **string** | Data or script used to add a specific configuration to the VM. It must be Base64-encoded and is limited to 500 kibibytes (KiB). | [optional] 
**VmInitiatedShutdownBehavior** | **string** | The VM behavior when you stop it. By default or if set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is terminated. | [optional] [default to stop]
**VmType** | **string** | The type of VM. You can specify a TINA type (in the &#x60;tinavW.cXrYpZ&#x60; or &#x60;tinavW.cXrY&#x60; format), or an AWS type (for example, &#x60;t2.small&#x60;, which is the default value).&lt;br /&gt; If you specify an AWS type, it is converted in the background to its corresponding TINA type, but the AWS type is still returned. If the specified or converted TINA type includes a performance flag, this performance flag is applied regardless of the value you may have provided in the &#x60;Performance&#x60; parameter. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


