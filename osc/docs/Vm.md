# Vm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | The architecture of the VM (&#x60;i386&#x60; \\| &#x60;x86_64&#x60;). | [optional] 
**BlockDeviceMappings** | [**[]BlockDeviceMappingCreated**](BlockDeviceMappingCreated.md) | The block device mapping of the VM. | [optional] 
**BsuOptimized** | **bool** | This parameter is not available. It is present in our API for the sake of historical compatibility with AWS. | [optional] 
**ClientToken** | **string** | The idempotency token provided when launching the VM. | [optional] 
**CreationDate** | **string** | The date and time at which the VM was created. | [optional] 
**DeletionProtection** | **bool** | If true, you cannot delete the VM unless you change this parameter back to false. | [optional] 
**Hypervisor** | **string** | The hypervisor type of the VMs (&#x60;ovm&#x60; \\| &#x60;xen&#x60;). | [optional] 
**ImageId** | **string** | The ID of the OMI used to create the VM. | [optional] 
**IsSourceDestChecked** | **bool** | (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**KeypairName** | **string** | The name of the keypair used when launching the VM. | [optional] 
**LaunchNumber** | **int32** | The number for the VM when launching a group of several VMs (for example, &#x60;0&#x60;, &#x60;1&#x60;, &#x60;2&#x60;, and so on). | [optional] 
**NestedVirtualization** | **bool** | If true, nested virtualization is enabled. If false, it is disabled. | [optional] 
**NetId** | **string** | The ID of the Net in which the VM is running. | [optional] 
**Nics** | [**[]NicLight**](NicLight.md) | (Net only) The network interface cards (NICs) the VMs are attached to. | [optional] 
**OsFamily** | **string** | Indicates the operating system (OS) of the VM. | [optional] 
**Performance** | **string** | The performance of the VM (&#x60;medium&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). | [optional] 
**Placement** | [**Placement**](Placement.md) |  | [optional] 
**PrivateDnsName** | **string** | The name of the private DNS. | [optional] 
**PrivateIp** | **string** | The primary private IP of the VM. | [optional] 
**ProductCodes** | **[]string** | The product code associated with the OMI used to create the VM (&#x60;0001&#x60; Linux/Unix \\| &#x60;0002&#x60; Windows \\| &#x60;0004&#x60; Linux/Oracle \\| &#x60;0005&#x60; Windows 10). | [optional] 
**PublicDnsName** | **string** | The name of the public DNS. | [optional] 
**PublicIp** | **string** | The public IP of the VM. | [optional] 
**ReservationId** | **string** | The reservation ID of the VM. | [optional] 
**RootDeviceName** | **string** | The name of the root device for the VM (for example, &#x60;/dev/vda1&#x60;). | [optional] 
**RootDeviceType** | **string** | The type of root device used by the VM (always &#x60;bsu&#x60;). | [optional] 
**SecurityGroups** | [**[]SecurityGroupLight**](SecurityGroupLight.md) | One or more security groups associated with the VM. | [optional] 
**State** | **string** | The state of the VM (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 
**StateReason** | **string** | The reason explaining the current state of the VM. | [optional] 
**SubnetId** | **string** | The ID of the Subnet for the VM. | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the VM. | [optional] 
**UserData** | **string** | The Base64-encoded MIME user data. | [optional] 
**VmId** | **string** | The ID of the VM. | [optional] 
**VmInitiatedShutdownBehavior** | **string** | The VM behavior when you stop it. If set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is deleted. | [optional] 
**VmType** | **string** | The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


