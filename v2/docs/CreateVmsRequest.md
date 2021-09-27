# CreateVmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDeviceMappings** | Pointer to [**[]BlockDeviceMappingVmCreation**](BlockDeviceMappingVmCreation.md) | One or more block device mappings. | [optional] 
**BootOnCreation** | Pointer to **bool** | By default or if true, the VM is started on creation. If false, the VM is stopped on creation. | [optional] 
**BsuOptimized** | Pointer to **bool** | If true, the VM is created with optimized BSU I/O. | [optional] 
**ClientToken** | Pointer to **string** | A unique identifier which enables you to manage the idempotency. | [optional] 
**DeletionProtection** | Pointer to **bool** | If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | **string** | The ID of the OMI used to create the VM. You can find the list of OMIs by calling the [ReadImages](#readimages) method. | 
**KeypairName** | Pointer to **string** | The name of the keypair. | [optional] 
**MaxVmsCount** | Pointer to **int32** | The maximum number of VMs you want to create. If all the VMs cannot be created, the largest possible number of VMs above MinVmsCount is created. | [optional] 
**MinVmsCount** | Pointer to **int32** | The minimum number of VMs you want to create. If this number of VMs cannot be created, no VMs are created. | [optional] 
**Nics** | Pointer to [**[]NicForVmCreation**](NicForVmCreation.md) | One or more NICs. If you specify this parameter, you must define one NIC as the primary network interface of the VM with &#x60;0&#x60; as its device number. | [optional] 
**Performance** | Pointer to **string** | The performance of the VM (&#x60;medium&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). | [optional] [default to "high"]
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**PrivateIps** | Pointer to **[]string** | One or more private IP addresses of the VM. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security group for the VMs. | [optional] 
**SecurityGroups** | Pointer to **[]string** | One or more names of security groups for the VMs. | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet in which you want to create the VM. | [optional] 
**UserData** | Pointer to **string** | Data or script used to add a specific configuration to the VM. It must be Base64-encoded and is limited to 500 kibibytes (KiB). | [optional] 
**VmInitiatedShutdownBehavior** | Pointer to **string** | The VM behavior when you stop it. By default or if set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is terminated. | [optional] 
**VmType** | Pointer to **string** | The type of VM (&#x60;t2.small&#x60; by default).&lt;br /&gt; For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). | [optional] 

## Methods

### NewCreateVmsRequest

`func NewCreateVmsRequest(imageId string, ) *CreateVmsRequest`

NewCreateVmsRequest instantiates a new CreateVmsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVmsRequestWithDefaults

`func NewCreateVmsRequestWithDefaults() *CreateVmsRequest`

NewCreateVmsRequestWithDefaults instantiates a new CreateVmsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockDeviceMappings

`func (o *CreateVmsRequest) GetBlockDeviceMappings() []BlockDeviceMappingVmCreation`

GetBlockDeviceMappings returns the BlockDeviceMappings field if non-nil, zero value otherwise.

### GetBlockDeviceMappingsOk

`func (o *CreateVmsRequest) GetBlockDeviceMappingsOk() (*[]BlockDeviceMappingVmCreation, bool)`

GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappings

`func (o *CreateVmsRequest) SetBlockDeviceMappings(v []BlockDeviceMappingVmCreation)`

SetBlockDeviceMappings sets BlockDeviceMappings field to given value.

### HasBlockDeviceMappings

`func (o *CreateVmsRequest) HasBlockDeviceMappings() bool`

HasBlockDeviceMappings returns a boolean if a field has been set.

### GetBootOnCreation

`func (o *CreateVmsRequest) GetBootOnCreation() bool`

GetBootOnCreation returns the BootOnCreation field if non-nil, zero value otherwise.

### GetBootOnCreationOk

`func (o *CreateVmsRequest) GetBootOnCreationOk() (*bool, bool)`

GetBootOnCreationOk returns a tuple with the BootOnCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOnCreation

`func (o *CreateVmsRequest) SetBootOnCreation(v bool)`

SetBootOnCreation sets BootOnCreation field to given value.

### HasBootOnCreation

`func (o *CreateVmsRequest) HasBootOnCreation() bool`

HasBootOnCreation returns a boolean if a field has been set.

### GetBsuOptimized

`func (o *CreateVmsRequest) GetBsuOptimized() bool`

GetBsuOptimized returns the BsuOptimized field if non-nil, zero value otherwise.

### GetBsuOptimizedOk

`func (o *CreateVmsRequest) GetBsuOptimizedOk() (*bool, bool)`

GetBsuOptimizedOk returns a tuple with the BsuOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsuOptimized

`func (o *CreateVmsRequest) SetBsuOptimized(v bool)`

SetBsuOptimized sets BsuOptimized field to given value.

### HasBsuOptimized

`func (o *CreateVmsRequest) HasBsuOptimized() bool`

HasBsuOptimized returns a boolean if a field has been set.

### GetClientToken

`func (o *CreateVmsRequest) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *CreateVmsRequest) GetClientTokenOk() (*string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *CreateVmsRequest) SetClientToken(v string)`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *CreateVmsRequest) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### GetDeletionProtection

`func (o *CreateVmsRequest) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *CreateVmsRequest) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *CreateVmsRequest) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.

### HasDeletionProtection

`func (o *CreateVmsRequest) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateVmsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVmsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateVmsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateVmsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetImageId

`func (o *CreateVmsRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateVmsRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateVmsRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetKeypairName

`func (o *CreateVmsRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *CreateVmsRequest) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *CreateVmsRequest) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *CreateVmsRequest) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetMaxVmsCount

`func (o *CreateVmsRequest) GetMaxVmsCount() int32`

GetMaxVmsCount returns the MaxVmsCount field if non-nil, zero value otherwise.

### GetMaxVmsCountOk

`func (o *CreateVmsRequest) GetMaxVmsCountOk() (*int32, bool)`

GetMaxVmsCountOk returns a tuple with the MaxVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVmsCount

`func (o *CreateVmsRequest) SetMaxVmsCount(v int32)`

SetMaxVmsCount sets MaxVmsCount field to given value.

### HasMaxVmsCount

`func (o *CreateVmsRequest) HasMaxVmsCount() bool`

HasMaxVmsCount returns a boolean if a field has been set.

### GetMinVmsCount

`func (o *CreateVmsRequest) GetMinVmsCount() int32`

GetMinVmsCount returns the MinVmsCount field if non-nil, zero value otherwise.

### GetMinVmsCountOk

`func (o *CreateVmsRequest) GetMinVmsCountOk() (*int32, bool)`

GetMinVmsCountOk returns a tuple with the MinVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVmsCount

`func (o *CreateVmsRequest) SetMinVmsCount(v int32)`

SetMinVmsCount sets MinVmsCount field to given value.

### HasMinVmsCount

`func (o *CreateVmsRequest) HasMinVmsCount() bool`

HasMinVmsCount returns a boolean if a field has been set.

### GetNics

`func (o *CreateVmsRequest) GetNics() []NicForVmCreation`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *CreateVmsRequest) GetNicsOk() (*[]NicForVmCreation, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *CreateVmsRequest) SetNics(v []NicForVmCreation)`

SetNics sets Nics field to given value.

### HasNics

`func (o *CreateVmsRequest) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetPerformance

`func (o *CreateVmsRequest) GetPerformance() string`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *CreateVmsRequest) GetPerformanceOk() (*string, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *CreateVmsRequest) SetPerformance(v string)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *CreateVmsRequest) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetPlacement

`func (o *CreateVmsRequest) GetPlacement() Placement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *CreateVmsRequest) GetPlacementOk() (*Placement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *CreateVmsRequest) SetPlacement(v Placement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *CreateVmsRequest) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetPrivateIps

`func (o *CreateVmsRequest) GetPrivateIps() []string`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *CreateVmsRequest) GetPrivateIpsOk() (*[]string, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *CreateVmsRequest) SetPrivateIps(v []string)`

SetPrivateIps sets PrivateIps field to given value.

### HasPrivateIps

`func (o *CreateVmsRequest) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *CreateVmsRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *CreateVmsRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *CreateVmsRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *CreateVmsRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CreateVmsRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateVmsRequest) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateVmsRequest) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CreateVmsRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSubnetId

`func (o *CreateVmsRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateVmsRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateVmsRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *CreateVmsRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetUserData

`func (o *CreateVmsRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateVmsRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateVmsRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateVmsRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetVmInitiatedShutdownBehavior

`func (o *CreateVmsRequest) GetVmInitiatedShutdownBehavior() string`

GetVmInitiatedShutdownBehavior returns the VmInitiatedShutdownBehavior field if non-nil, zero value otherwise.

### GetVmInitiatedShutdownBehaviorOk

`func (o *CreateVmsRequest) GetVmInitiatedShutdownBehaviorOk() (*string, bool)`

GetVmInitiatedShutdownBehaviorOk returns a tuple with the VmInitiatedShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInitiatedShutdownBehavior

`func (o *CreateVmsRequest) SetVmInitiatedShutdownBehavior(v string)`

SetVmInitiatedShutdownBehavior sets VmInitiatedShutdownBehavior field to given value.

### HasVmInitiatedShutdownBehavior

`func (o *CreateVmsRequest) HasVmInitiatedShutdownBehavior() bool`

HasVmInitiatedShutdownBehavior returns a boolean if a field has been set.

### GetVmType

`func (o *CreateVmsRequest) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *CreateVmsRequest) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *CreateVmsRequest) SetVmType(v string)`

SetVmType sets VmType field to given value.

### HasVmType

`func (o *CreateVmsRequest) HasVmType() bool`

HasVmType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


