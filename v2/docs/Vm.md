# Vm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | The architecture of the VM (&#x60;i386&#x60; \\| &#x60;x86_64&#x60;). | [optional] 
**BlockDeviceMappings** | Pointer to [**[]BlockDeviceMappingCreated**](BlockDeviceMappingCreated.md) | The block device mapping of the VM. | [optional] 
**BsuOptimized** | Pointer to **bool** | This parameter is not available. It is present in our API for the sake of historical compatibility with AWS. | [optional] 
**ClientToken** | Pointer to **string** | The idempotency token provided when launching the VM. | [optional] 
**CreationDate** | Pointer to **string** | The date and time of creation of the VM. | [optional] 
**DeletionProtection** | Pointer to **bool** | If true, you cannot delete the VM unless you change this parameter back to false. | [optional] 
**Hypervisor** | Pointer to **string** | The hypervisor type of the VMs (&#x60;ovm&#x60; \\| &#x60;xen&#x60;). | [optional] 
**ImageId** | Pointer to **string** | The ID of the OMI used to create the VM. | [optional] 
**IsSourceDestChecked** | Pointer to **bool** | (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**KeypairName** | Pointer to **string** | The name of the keypair used when launching the VM. | [optional] 
**LaunchNumber** | Pointer to **int32** | The number for the VM when launching a group of several VMs (for example, &#x60;0&#x60;, &#x60;1&#x60;, &#x60;2&#x60;, and so on). | [optional] 
**NestedVirtualization** | Pointer to **bool** | If true, nested virtualization is enabled. If false, it is disabled. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net in which the VM is running. | [optional] 
**Nics** | Pointer to [**[]NicLight**](NicLight.md) | (Net only) The network interface cards (NICs) the VMs are attached to. | [optional] 
**OsFamily** | Pointer to **string** | Indicates the operating system (OS) of the VM. | [optional] 
**Performance** | Pointer to **string** | The performance of the VM (&#x60;medium&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). | [optional] 
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**PrivateDnsName** | Pointer to **string** | The name of the private DNS. | [optional] 
**PrivateIp** | Pointer to **string** | The primary private IP of the VM. | [optional] 
**ProductCodes** | Pointer to **[]string** | The product codes associated with the OMI used to create the VM. | [optional] 
**PublicDnsName** | Pointer to **string** | The name of the public DNS. | [optional] 
**PublicIp** | Pointer to **string** | The public IP of the VM. | [optional] 
**ReservationId** | Pointer to **string** | The reservation ID of the VM. | [optional] 
**RootDeviceName** | Pointer to **string** | The name of the root device for the VM (for example, &#x60;/dev/sda1&#x60;). | [optional] 
**RootDeviceType** | Pointer to **string** | The type of root device used by the VM (always &#x60;bsu&#x60;). | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupLight**](SecurityGroupLight.md) | One or more security groups associated with the VM. | [optional] 
**State** | Pointer to **string** | The state of the VM (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 
**StateReason** | Pointer to **string** | The reason explaining the current state of the VM. | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet for the VM. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the VM. | [optional] 
**UserData** | Pointer to **string** | The Base64-encoded MIME user data. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 
**VmInitiatedShutdownBehavior** | Pointer to **string** | The VM behavior when you stop it. If set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is deleted. | [optional] 
**VmType** | Pointer to **string** | The type of VM. For more information, see [VM Types](https://docs.outscale.com/en/userguide/VM-Types.html). | [optional] 

## Methods

### NewVm

`func NewVm() *Vm`

NewVm instantiates a new Vm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmWithDefaults

`func NewVmWithDefaults() *Vm`

NewVmWithDefaults instantiates a new Vm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *Vm) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Vm) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Vm) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *Vm) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBlockDeviceMappings

`func (o *Vm) GetBlockDeviceMappings() []BlockDeviceMappingCreated`

GetBlockDeviceMappings returns the BlockDeviceMappings field if non-nil, zero value otherwise.

### GetBlockDeviceMappingsOk

`func (o *Vm) GetBlockDeviceMappingsOk() (*[]BlockDeviceMappingCreated, bool)`

GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappings

`func (o *Vm) SetBlockDeviceMappings(v []BlockDeviceMappingCreated)`

SetBlockDeviceMappings sets BlockDeviceMappings field to given value.

### HasBlockDeviceMappings

`func (o *Vm) HasBlockDeviceMappings() bool`

HasBlockDeviceMappings returns a boolean if a field has been set.

### GetBsuOptimized

`func (o *Vm) GetBsuOptimized() bool`

GetBsuOptimized returns the BsuOptimized field if non-nil, zero value otherwise.

### GetBsuOptimizedOk

`func (o *Vm) GetBsuOptimizedOk() (*bool, bool)`

GetBsuOptimizedOk returns a tuple with the BsuOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsuOptimized

`func (o *Vm) SetBsuOptimized(v bool)`

SetBsuOptimized sets BsuOptimized field to given value.

### HasBsuOptimized

`func (o *Vm) HasBsuOptimized() bool`

HasBsuOptimized returns a boolean if a field has been set.

### GetClientToken

`func (o *Vm) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *Vm) GetClientTokenOk() (*string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *Vm) SetClientToken(v string)`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *Vm) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### GetCreationDate

`func (o *Vm) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Vm) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Vm) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Vm) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDeletionProtection

`func (o *Vm) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *Vm) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *Vm) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.

### HasDeletionProtection

`func (o *Vm) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.

### GetHypervisor

`func (o *Vm) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *Vm) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *Vm) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *Vm) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetImageId

`func (o *Vm) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Vm) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *Vm) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *Vm) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetIsSourceDestChecked

`func (o *Vm) GetIsSourceDestChecked() bool`

GetIsSourceDestChecked returns the IsSourceDestChecked field if non-nil, zero value otherwise.

### GetIsSourceDestCheckedOk

`func (o *Vm) GetIsSourceDestCheckedOk() (*bool, bool)`

GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceDestChecked

`func (o *Vm) SetIsSourceDestChecked(v bool)`

SetIsSourceDestChecked sets IsSourceDestChecked field to given value.

### HasIsSourceDestChecked

`func (o *Vm) HasIsSourceDestChecked() bool`

HasIsSourceDestChecked returns a boolean if a field has been set.

### GetKeypairName

`func (o *Vm) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *Vm) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *Vm) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *Vm) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetLaunchNumber

`func (o *Vm) GetLaunchNumber() int32`

GetLaunchNumber returns the LaunchNumber field if non-nil, zero value otherwise.

### GetLaunchNumberOk

`func (o *Vm) GetLaunchNumberOk() (*int32, bool)`

GetLaunchNumberOk returns a tuple with the LaunchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchNumber

`func (o *Vm) SetLaunchNumber(v int32)`

SetLaunchNumber sets LaunchNumber field to given value.

### HasLaunchNumber

`func (o *Vm) HasLaunchNumber() bool`

HasLaunchNumber returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *Vm) GetNestedVirtualization() bool`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *Vm) GetNestedVirtualizationOk() (*bool, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *Vm) SetNestedVirtualization(v bool)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *Vm) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetNetId

`func (o *Vm) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *Vm) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *Vm) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *Vm) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetNics

`func (o *Vm) GetNics() []NicLight`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *Vm) GetNicsOk() (*[]NicLight, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *Vm) SetNics(v []NicLight)`

SetNics sets Nics field to given value.

### HasNics

`func (o *Vm) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetOsFamily

`func (o *Vm) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *Vm) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *Vm) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *Vm) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetPerformance

`func (o *Vm) GetPerformance() string`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *Vm) GetPerformanceOk() (*string, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *Vm) SetPerformance(v string)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *Vm) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetPlacement

`func (o *Vm) GetPlacement() Placement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *Vm) GetPlacementOk() (*Placement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *Vm) SetPlacement(v Placement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *Vm) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetPrivateDnsName

`func (o *Vm) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *Vm) GetPrivateDnsNameOk() (*string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsName

`func (o *Vm) SetPrivateDnsName(v string)`

SetPrivateDnsName sets PrivateDnsName field to given value.

### HasPrivateDnsName

`func (o *Vm) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### GetPrivateIp

`func (o *Vm) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *Vm) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *Vm) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *Vm) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProductCodes

`func (o *Vm) GetProductCodes() []string`

GetProductCodes returns the ProductCodes field if non-nil, zero value otherwise.

### GetProductCodesOk

`func (o *Vm) GetProductCodesOk() (*[]string, bool)`

GetProductCodesOk returns a tuple with the ProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCodes

`func (o *Vm) SetProductCodes(v []string)`

SetProductCodes sets ProductCodes field to given value.

### HasProductCodes

`func (o *Vm) HasProductCodes() bool`

HasProductCodes returns a boolean if a field has been set.

### GetPublicDnsName

`func (o *Vm) GetPublicDnsName() string`

GetPublicDnsName returns the PublicDnsName field if non-nil, zero value otherwise.

### GetPublicDnsNameOk

`func (o *Vm) GetPublicDnsNameOk() (*string, bool)`

GetPublicDnsNameOk returns a tuple with the PublicDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDnsName

`func (o *Vm) SetPublicDnsName(v string)`

SetPublicDnsName sets PublicDnsName field to given value.

### HasPublicDnsName

`func (o *Vm) HasPublicDnsName() bool`

HasPublicDnsName returns a boolean if a field has been set.

### GetPublicIp

`func (o *Vm) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *Vm) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *Vm) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *Vm) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetReservationId

`func (o *Vm) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *Vm) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *Vm) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.

### HasReservationId

`func (o *Vm) HasReservationId() bool`

HasReservationId returns a boolean if a field has been set.

### GetRootDeviceName

`func (o *Vm) GetRootDeviceName() string`

GetRootDeviceName returns the RootDeviceName field if non-nil, zero value otherwise.

### GetRootDeviceNameOk

`func (o *Vm) GetRootDeviceNameOk() (*string, bool)`

GetRootDeviceNameOk returns a tuple with the RootDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceName

`func (o *Vm) SetRootDeviceName(v string)`

SetRootDeviceName sets RootDeviceName field to given value.

### HasRootDeviceName

`func (o *Vm) HasRootDeviceName() bool`

HasRootDeviceName returns a boolean if a field has been set.

### GetRootDeviceType

`func (o *Vm) GetRootDeviceType() string`

GetRootDeviceType returns the RootDeviceType field if non-nil, zero value otherwise.

### GetRootDeviceTypeOk

`func (o *Vm) GetRootDeviceTypeOk() (*string, bool)`

GetRootDeviceTypeOk returns a tuple with the RootDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceType

`func (o *Vm) SetRootDeviceType(v string)`

SetRootDeviceType sets RootDeviceType field to given value.

### HasRootDeviceType

`func (o *Vm) HasRootDeviceType() bool`

HasRootDeviceType returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *Vm) GetSecurityGroups() []SecurityGroupLight`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *Vm) GetSecurityGroupsOk() (*[]SecurityGroupLight, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *Vm) SetSecurityGroups(v []SecurityGroupLight)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *Vm) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetState

`func (o *Vm) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Vm) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Vm) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Vm) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *Vm) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Vm) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *Vm) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *Vm) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetSubnetId

`func (o *Vm) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Vm) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Vm) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *Vm) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTags

`func (o *Vm) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Vm) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Vm) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Vm) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserData

`func (o *Vm) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *Vm) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *Vm) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *Vm) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetVmId

`func (o *Vm) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *Vm) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *Vm) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *Vm) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetVmInitiatedShutdownBehavior

`func (o *Vm) GetVmInitiatedShutdownBehavior() string`

GetVmInitiatedShutdownBehavior returns the VmInitiatedShutdownBehavior field if non-nil, zero value otherwise.

### GetVmInitiatedShutdownBehaviorOk

`func (o *Vm) GetVmInitiatedShutdownBehaviorOk() (*string, bool)`

GetVmInitiatedShutdownBehaviorOk returns a tuple with the VmInitiatedShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInitiatedShutdownBehavior

`func (o *Vm) SetVmInitiatedShutdownBehavior(v string)`

SetVmInitiatedShutdownBehavior sets VmInitiatedShutdownBehavior field to given value.

### HasVmInitiatedShutdownBehavior

`func (o *Vm) HasVmInitiatedShutdownBehavior() bool`

HasVmInitiatedShutdownBehavior returns a boolean if a field has been set.

### GetVmType

`func (o *Vm) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *Vm) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *Vm) SetVmType(v string)`

SetVmType sets VmType field to given value.

### HasVmType

`func (o *Vm) HasVmType() bool`

HasVmType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


