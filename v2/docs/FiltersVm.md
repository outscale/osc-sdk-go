# FiltersVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architectures** | Pointer to **[]string** | The architectures of the VMs (&#x60;i386&#x60; \\| &#x60;x86_64&#x60;). | [optional] 
**BlockDeviceMappingDeleteOnVmDeletion** | Pointer to **bool** | Whether the BSU volumes are deleted when terminating the VMs. | [optional] 
**BlockDeviceMappingDeviceNames** | Pointer to **[]string** | The device names for the BSU volumes (in the format &#x60;/dev/sdX&#x60;, &#x60;/dev/sdXX&#x60;, &#x60;/dev/xvdX&#x60;, or &#x60;/dev/xvdXX&#x60;). | [optional] 
**BlockDeviceMappingLinkDates** | Pointer to **[]string** | The link dates for the BSU volumes mapped to the VMs (for example, &#x60;2016-01-23T18:45:30.000Z&#x60;). | [optional] 
**BlockDeviceMappingStates** | Pointer to **[]string** | The states for the BSU volumes (&#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60; \\| &#x60;detached&#x60;). | [optional] 
**BlockDeviceMappingVolumeIds** | Pointer to **[]string** | The volume IDs of the BSU volumes. | [optional] 
**ClientTokens** | Pointer to **[]string** | The idempotency tokens provided when launching the VMs. | [optional] 
**CreationDates** | Pointer to **[]string** | The dates when the VMs were launched. | [optional] 
**ImageIds** | Pointer to **[]string** | The IDs of the OMIs used to launch the VMs. | [optional] 
**IsSourceDestChecked** | Pointer to **bool** | Whether the source/destination checking is enabled (true) or disabled (false). | [optional] 
**KeypairNames** | Pointer to **[]string** | The names of the keypairs used when launching the VMs. | [optional] 
**LaunchNumbers** | Pointer to **[]int32** | The numbers for the VMs when launching a group of several VMs (for example, &#x60;0&#x60;, &#x60;1&#x60;, &#x60;2&#x60;, and so on). | [optional] 
**Lifecycles** | Pointer to **[]string** | Whether the VMs are Spot Instances (spot). | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets in which the VMs are running. | [optional] 
**NicAccountIds** | Pointer to **[]string** | The IDs of the NICs. | [optional] 
**NicDescriptions** | Pointer to **[]string** | The descriptions of the NICs. | [optional] 
**NicIsSourceDestChecked** | Pointer to **bool** | Whether the source/destination checking is enabled (true) or disabled (false). | [optional] 
**NicLinkNicDeleteOnVmDeletion** | Pointer to **bool** | Whether the NICs are deleted when the VMs they are attached to are deleted. | [optional] 
**NicLinkNicDeviceNumbers** | Pointer to **[]int32** | The device numbers the NICs are attached to. | [optional] 
**NicLinkNicLinkNicDates** | Pointer to **[]string** | The dates and time when the NICs were attached to the VMs. | [optional] 
**NicLinkNicLinkNicIds** | Pointer to **[]string** | The IDs of the NIC attachments. | [optional] 
**NicLinkNicStates** | Pointer to **[]string** | The states of the attachments. | [optional] 
**NicLinkNicVmAccountIds** | Pointer to **[]string** | The account IDs of the owners of the VMs the NICs are attached to. | [optional] 
**NicLinkNicVmIds** | Pointer to **[]string** | The IDs of the VMs the NICs are attached to. | [optional] 
**NicLinkPublicIpAccountIds** | Pointer to **[]string** | The account IDs of the owners of the public IPs associated with the NICs. | [optional] 
**NicLinkPublicIpLinkPublicIpIds** | Pointer to **[]string** | The association IDs returned when the public IPs were associated with the NICs. | [optional] 
**NicLinkPublicIpPublicIpIds** | Pointer to **[]string** | The allocation IDs returned when the public IPs were allocated to their accounts. | [optional] 
**NicLinkPublicIpPublicIps** | Pointer to **[]string** | The public IPs associated with the NICs. | [optional] 
**NicMacAddresses** | Pointer to **[]string** | The Media Access Control (MAC) addresses of the NICs. | [optional] 
**NicNetIds** | Pointer to **[]string** | The IDs of the Nets where the NICs are located. | [optional] 
**NicNicIds** | Pointer to **[]string** | The IDs of the NICs. | [optional] 
**NicPrivateIpsLinkPublicIpAccountIds** | Pointer to **[]string** | The account IDs of the owner of the public IPs associated with the private IPs. | [optional] 
**NicPrivateIpsLinkPublicIpIds** | Pointer to **[]string** | The public IPs associated with the private IPs. | [optional] 
**NicPrivateIpsPrimaryIp** | Pointer to **bool** | Whether the private IPs are the primary IPs associated with the NICs. | [optional] 
**NicPrivateIpsPrivateIps** | Pointer to **[]string** | The private IPs of the NICs. | [optional] 
**NicSecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups associated with the NICs. | [optional] 
**NicSecurityGroupNames** | Pointer to **[]string** | The names of the security groups associated with the NICs. | [optional] 
**NicStates** | Pointer to **[]string** | The states of the NICs (&#x60;available&#x60; \\| &#x60;in-use&#x60;). | [optional] 
**NicSubnetIds** | Pointer to **[]string** | The IDs of the Subnets for the NICs. | [optional] 
**NicSubregionNames** | Pointer to **[]string** | The Subregions where the NICs are located. | [optional] 
**Platforms** | Pointer to **[]string** | The platforms. Use windows if you have Windows VMs. Otherwise, leave this filter blank. | [optional] 
**PrivateIps** | Pointer to **[]string** | The private IPs of the VMs. | [optional] 
**ProductCodes** | Pointer to **[]string** | The product codes associated with the OMI used to create the VMs. | [optional] 
**PublicIps** | Pointer to **[]string** | The public IPs of the VMs. | [optional] 
**ReservationIds** | Pointer to **[]string** | The IDs of the reservation of the VMs, created every time you launch VMs. These reservation IDs can be associated with several VMs when you lauch a group of VMs using the same launch request. | [optional] 
**RootDeviceNames** | Pointer to **[]string** | The names of the root devices for the VMs (for example, &#x60;/dev/sda1&#x60;) | [optional] 
**RootDeviceTypes** | Pointer to **[]string** | The root devices types used by the VMs (always &#x60;ebs&#x60;) | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups for the VMs (only in the public Cloud). | [optional] 
**SecurityGroupNames** | Pointer to **[]string** | The names of the security groups for the VMs (only in the public Cloud). | [optional] 
**StateReasonCodes** | Pointer to **[]int32** | The reason codes for the state changes. | [optional] 
**StateReasonMessages** | Pointer to **[]string** | The messages describing the state changes. | [optional] 
**StateReasons** | Pointer to **[]string** | The reasons explaining the current states of the VMs. This filter is like the &#x60;StateReasonCodes&#x60; one. | [optional] 
**SubnetIds** | Pointer to **[]string** | The IDs of the Subnets for the VMs. | [optional] 
**SubregionNames** | Pointer to **[]string** | The names of the Subregions of the VMs. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the VMs. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the VMs. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the VMs, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**Tenancies** | Pointer to **[]string** | The tenancies of the VMs (&#x60;dedicated&#x60; \\| &#x60;default&#x60; \\| &#x60;host&#x60;). | [optional] 
**VmIds** | Pointer to **[]string** | One or more IDs of VMs. | [optional] 
**VmSecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups for the VMs. | [optional] 
**VmSecurityGroupNames** | Pointer to **[]string** | The names of the security group for the VMs. | [optional] 
**VmStateCodes** | Pointer to **[]int32** | The state codes of the VMs: &#x60;-1&#x60; (quarantine), &#x60;0&#x60; (pending), &#x60;16&#x60; (running), &#x60;32&#x60; (shutting-down), &#x60;48&#x60; (terminated), &#x60;64&#x60; (stopping), and &#x60;80&#x60; (stopped). | [optional] 
**VmStateNames** | Pointer to **[]string** | The state names of the VMs (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 
**VmTypes** | Pointer to **[]string** | The VM types (for example, t2.micro). For more information, see [VM Types](https://docs.outscale.com/en/userguide/VM-Types.html). | [optional] 

## Methods

### NewFiltersVm

`func NewFiltersVm() *FiltersVm`

NewFiltersVm instantiates a new FiltersVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersVmWithDefaults

`func NewFiltersVmWithDefaults() *FiltersVm`

NewFiltersVmWithDefaults instantiates a new FiltersVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectures

`func (o *FiltersVm) GetArchitectures() []string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *FiltersVm) GetArchitecturesOk() (*[]string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *FiltersVm) SetArchitectures(v []string)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *FiltersVm) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersVm) GetBlockDeviceMappingDeleteOnVmDeletion() bool`

GetBlockDeviceMappingDeleteOnVmDeletion returns the BlockDeviceMappingDeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetBlockDeviceMappingDeleteOnVmDeletionOk

`func (o *FiltersVm) GetBlockDeviceMappingDeleteOnVmDeletionOk() (*bool, bool)`

GetBlockDeviceMappingDeleteOnVmDeletionOk returns a tuple with the BlockDeviceMappingDeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersVm) SetBlockDeviceMappingDeleteOnVmDeletion(v bool)`

SetBlockDeviceMappingDeleteOnVmDeletion sets BlockDeviceMappingDeleteOnVmDeletion field to given value.

### HasBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersVm) HasBlockDeviceMappingDeleteOnVmDeletion() bool`

HasBlockDeviceMappingDeleteOnVmDeletion returns a boolean if a field has been set.

### GetBlockDeviceMappingDeviceNames

`func (o *FiltersVm) GetBlockDeviceMappingDeviceNames() []string`

GetBlockDeviceMappingDeviceNames returns the BlockDeviceMappingDeviceNames field if non-nil, zero value otherwise.

### GetBlockDeviceMappingDeviceNamesOk

`func (o *FiltersVm) GetBlockDeviceMappingDeviceNamesOk() (*[]string, bool)`

GetBlockDeviceMappingDeviceNamesOk returns a tuple with the BlockDeviceMappingDeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingDeviceNames

`func (o *FiltersVm) SetBlockDeviceMappingDeviceNames(v []string)`

SetBlockDeviceMappingDeviceNames sets BlockDeviceMappingDeviceNames field to given value.

### HasBlockDeviceMappingDeviceNames

`func (o *FiltersVm) HasBlockDeviceMappingDeviceNames() bool`

HasBlockDeviceMappingDeviceNames returns a boolean if a field has been set.

### GetBlockDeviceMappingLinkDates

`func (o *FiltersVm) GetBlockDeviceMappingLinkDates() []string`

GetBlockDeviceMappingLinkDates returns the BlockDeviceMappingLinkDates field if non-nil, zero value otherwise.

### GetBlockDeviceMappingLinkDatesOk

`func (o *FiltersVm) GetBlockDeviceMappingLinkDatesOk() (*[]string, bool)`

GetBlockDeviceMappingLinkDatesOk returns a tuple with the BlockDeviceMappingLinkDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingLinkDates

`func (o *FiltersVm) SetBlockDeviceMappingLinkDates(v []string)`

SetBlockDeviceMappingLinkDates sets BlockDeviceMappingLinkDates field to given value.

### HasBlockDeviceMappingLinkDates

`func (o *FiltersVm) HasBlockDeviceMappingLinkDates() bool`

HasBlockDeviceMappingLinkDates returns a boolean if a field has been set.

### GetBlockDeviceMappingStates

`func (o *FiltersVm) GetBlockDeviceMappingStates() []string`

GetBlockDeviceMappingStates returns the BlockDeviceMappingStates field if non-nil, zero value otherwise.

### GetBlockDeviceMappingStatesOk

`func (o *FiltersVm) GetBlockDeviceMappingStatesOk() (*[]string, bool)`

GetBlockDeviceMappingStatesOk returns a tuple with the BlockDeviceMappingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingStates

`func (o *FiltersVm) SetBlockDeviceMappingStates(v []string)`

SetBlockDeviceMappingStates sets BlockDeviceMappingStates field to given value.

### HasBlockDeviceMappingStates

`func (o *FiltersVm) HasBlockDeviceMappingStates() bool`

HasBlockDeviceMappingStates returns a boolean if a field has been set.

### GetBlockDeviceMappingVolumeIds

`func (o *FiltersVm) GetBlockDeviceMappingVolumeIds() []string`

GetBlockDeviceMappingVolumeIds returns the BlockDeviceMappingVolumeIds field if non-nil, zero value otherwise.

### GetBlockDeviceMappingVolumeIdsOk

`func (o *FiltersVm) GetBlockDeviceMappingVolumeIdsOk() (*[]string, bool)`

GetBlockDeviceMappingVolumeIdsOk returns a tuple with the BlockDeviceMappingVolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingVolumeIds

`func (o *FiltersVm) SetBlockDeviceMappingVolumeIds(v []string)`

SetBlockDeviceMappingVolumeIds sets BlockDeviceMappingVolumeIds field to given value.

### HasBlockDeviceMappingVolumeIds

`func (o *FiltersVm) HasBlockDeviceMappingVolumeIds() bool`

HasBlockDeviceMappingVolumeIds returns a boolean if a field has been set.

### GetClientTokens

`func (o *FiltersVm) GetClientTokens() []string`

GetClientTokens returns the ClientTokens field if non-nil, zero value otherwise.

### GetClientTokensOk

`func (o *FiltersVm) GetClientTokensOk() (*[]string, bool)`

GetClientTokensOk returns a tuple with the ClientTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTokens

`func (o *FiltersVm) SetClientTokens(v []string)`

SetClientTokens sets ClientTokens field to given value.

### HasClientTokens

`func (o *FiltersVm) HasClientTokens() bool`

HasClientTokens returns a boolean if a field has been set.

### GetCreationDates

`func (o *FiltersVm) GetCreationDates() []string`

GetCreationDates returns the CreationDates field if non-nil, zero value otherwise.

### GetCreationDatesOk

`func (o *FiltersVm) GetCreationDatesOk() (*[]string, bool)`

GetCreationDatesOk returns a tuple with the CreationDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDates

`func (o *FiltersVm) SetCreationDates(v []string)`

SetCreationDates sets CreationDates field to given value.

### HasCreationDates

`func (o *FiltersVm) HasCreationDates() bool`

HasCreationDates returns a boolean if a field has been set.

### GetImageIds

`func (o *FiltersVm) GetImageIds() []string`

GetImageIds returns the ImageIds field if non-nil, zero value otherwise.

### GetImageIdsOk

`func (o *FiltersVm) GetImageIdsOk() (*[]string, bool)`

GetImageIdsOk returns a tuple with the ImageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIds

`func (o *FiltersVm) SetImageIds(v []string)`

SetImageIds sets ImageIds field to given value.

### HasImageIds

`func (o *FiltersVm) HasImageIds() bool`

HasImageIds returns a boolean if a field has been set.

### GetIsSourceDestChecked

`func (o *FiltersVm) GetIsSourceDestChecked() bool`

GetIsSourceDestChecked returns the IsSourceDestChecked field if non-nil, zero value otherwise.

### GetIsSourceDestCheckedOk

`func (o *FiltersVm) GetIsSourceDestCheckedOk() (*bool, bool)`

GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceDestChecked

`func (o *FiltersVm) SetIsSourceDestChecked(v bool)`

SetIsSourceDestChecked sets IsSourceDestChecked field to given value.

### HasIsSourceDestChecked

`func (o *FiltersVm) HasIsSourceDestChecked() bool`

HasIsSourceDestChecked returns a boolean if a field has been set.

### GetKeypairNames

`func (o *FiltersVm) GetKeypairNames() []string`

GetKeypairNames returns the KeypairNames field if non-nil, zero value otherwise.

### GetKeypairNamesOk

`func (o *FiltersVm) GetKeypairNamesOk() (*[]string, bool)`

GetKeypairNamesOk returns a tuple with the KeypairNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairNames

`func (o *FiltersVm) SetKeypairNames(v []string)`

SetKeypairNames sets KeypairNames field to given value.

### HasKeypairNames

`func (o *FiltersVm) HasKeypairNames() bool`

HasKeypairNames returns a boolean if a field has been set.

### GetLaunchNumbers

`func (o *FiltersVm) GetLaunchNumbers() []int32`

GetLaunchNumbers returns the LaunchNumbers field if non-nil, zero value otherwise.

### GetLaunchNumbersOk

`func (o *FiltersVm) GetLaunchNumbersOk() (*[]int32, bool)`

GetLaunchNumbersOk returns a tuple with the LaunchNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchNumbers

`func (o *FiltersVm) SetLaunchNumbers(v []int32)`

SetLaunchNumbers sets LaunchNumbers field to given value.

### HasLaunchNumbers

`func (o *FiltersVm) HasLaunchNumbers() bool`

HasLaunchNumbers returns a boolean if a field has been set.

### GetLifecycles

`func (o *FiltersVm) GetLifecycles() []string`

GetLifecycles returns the Lifecycles field if non-nil, zero value otherwise.

### GetLifecyclesOk

`func (o *FiltersVm) GetLifecyclesOk() (*[]string, bool)`

GetLifecyclesOk returns a tuple with the Lifecycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycles

`func (o *FiltersVm) SetLifecycles(v []string)`

SetLifecycles sets Lifecycles field to given value.

### HasLifecycles

`func (o *FiltersVm) HasLifecycles() bool`

HasLifecycles returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersVm) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersVm) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersVm) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersVm) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetNicAccountIds

`func (o *FiltersVm) GetNicAccountIds() []string`

GetNicAccountIds returns the NicAccountIds field if non-nil, zero value otherwise.

### GetNicAccountIdsOk

`func (o *FiltersVm) GetNicAccountIdsOk() (*[]string, bool)`

GetNicAccountIdsOk returns a tuple with the NicAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicAccountIds

`func (o *FiltersVm) SetNicAccountIds(v []string)`

SetNicAccountIds sets NicAccountIds field to given value.

### HasNicAccountIds

`func (o *FiltersVm) HasNicAccountIds() bool`

HasNicAccountIds returns a boolean if a field has been set.

### GetNicDescriptions

`func (o *FiltersVm) GetNicDescriptions() []string`

GetNicDescriptions returns the NicDescriptions field if non-nil, zero value otherwise.

### GetNicDescriptionsOk

`func (o *FiltersVm) GetNicDescriptionsOk() (*[]string, bool)`

GetNicDescriptionsOk returns a tuple with the NicDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicDescriptions

`func (o *FiltersVm) SetNicDescriptions(v []string)`

SetNicDescriptions sets NicDescriptions field to given value.

### HasNicDescriptions

`func (o *FiltersVm) HasNicDescriptions() bool`

HasNicDescriptions returns a boolean if a field has been set.

### GetNicIsSourceDestChecked

`func (o *FiltersVm) GetNicIsSourceDestChecked() bool`

GetNicIsSourceDestChecked returns the NicIsSourceDestChecked field if non-nil, zero value otherwise.

### GetNicIsSourceDestCheckedOk

`func (o *FiltersVm) GetNicIsSourceDestCheckedOk() (*bool, bool)`

GetNicIsSourceDestCheckedOk returns a tuple with the NicIsSourceDestChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicIsSourceDestChecked

`func (o *FiltersVm) SetNicIsSourceDestChecked(v bool)`

SetNicIsSourceDestChecked sets NicIsSourceDestChecked field to given value.

### HasNicIsSourceDestChecked

`func (o *FiltersVm) HasNicIsSourceDestChecked() bool`

HasNicIsSourceDestChecked returns a boolean if a field has been set.

### GetNicLinkNicDeleteOnVmDeletion

`func (o *FiltersVm) GetNicLinkNicDeleteOnVmDeletion() bool`

GetNicLinkNicDeleteOnVmDeletion returns the NicLinkNicDeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetNicLinkNicDeleteOnVmDeletionOk

`func (o *FiltersVm) GetNicLinkNicDeleteOnVmDeletionOk() (*bool, bool)`

GetNicLinkNicDeleteOnVmDeletionOk returns a tuple with the NicLinkNicDeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkNicDeleteOnVmDeletion

`func (o *FiltersVm) SetNicLinkNicDeleteOnVmDeletion(v bool)`

SetNicLinkNicDeleteOnVmDeletion sets NicLinkNicDeleteOnVmDeletion field to given value.

### HasNicLinkNicDeleteOnVmDeletion

`func (o *FiltersVm) HasNicLinkNicDeleteOnVmDeletion() bool`

HasNicLinkNicDeleteOnVmDeletion returns a boolean if a field has been set.

### GetNicLinkNicDeviceNumbers

`func (o *FiltersVm) GetNicLinkNicDeviceNumbers() []int32`

GetNicLinkNicDeviceNumbers returns the NicLinkNicDeviceNumbers field if non-nil, zero value otherwise.

### GetNicLinkNicDeviceNumbersOk

`func (o *FiltersVm) GetNicLinkNicDeviceNumbersOk() (*[]int32, bool)`

GetNicLinkNicDeviceNumbersOk returns a tuple with the NicLinkNicDeviceNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkNicDeviceNumbers

`func (o *FiltersVm) SetNicLinkNicDeviceNumbers(v []int32)`

SetNicLinkNicDeviceNumbers sets NicLinkNicDeviceNumbers field to given value.

### HasNicLinkNicDeviceNumbers

`func (o *FiltersVm) HasNicLinkNicDeviceNumbers() bool`

HasNicLinkNicDeviceNumbers returns a boolean if a field has been set.

### GetNicLinkNicLinkNicDates

`func (o *FiltersVm) GetNicLinkNicLinkNicDates() []string`

GetNicLinkNicLinkNicDates returns the NicLinkNicLinkNicDates field if non-nil, zero value otherwise.

### GetNicLinkNicLinkNicDatesOk

`func (o *FiltersVm) GetNicLinkNicLinkNicDatesOk() (*[]string, bool)`

GetNicLinkNicLinkNicDatesOk returns a tuple with the NicLinkNicLinkNicDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkNicLinkNicDates

`func (o *FiltersVm) SetNicLinkNicLinkNicDates(v []string)`

SetNicLinkNicLinkNicDates sets NicLinkNicLinkNicDates field to given value.

### HasNicLinkNicLinkNicDates

`func (o *FiltersVm) HasNicLinkNicLinkNicDates() bool`

HasNicLinkNicLinkNicDates returns a boolean if a field has been set.

### GetNicLinkNicLinkNicIds

`func (o *FiltersVm) GetNicLinkNicLinkNicIds() []string`

GetNicLinkNicLinkNicIds returns the NicLinkNicLinkNicIds field if non-nil, zero value otherwise.

### GetNicLinkNicLinkNicIdsOk

`func (o *FiltersVm) GetNicLinkNicLinkNicIdsOk() (*[]string, bool)`

GetNicLinkNicLinkNicIdsOk returns a tuple with the NicLinkNicLinkNicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkNicLinkNicIds

`func (o *FiltersVm) SetNicLinkNicLinkNicIds(v []string)`

SetNicLinkNicLinkNicIds sets NicLinkNicLinkNicIds field to given value.

### HasNicLinkNicLinkNicIds

`func (o *FiltersVm) HasNicLinkNicLinkNicIds() bool`

HasNicLinkNicLinkNicIds returns a boolean if a field has been set.

### GetNicLinkNicStates

`func (o *FiltersVm) GetNicLinkNicStates() []string`

GetNicLinkNicStates returns the NicLinkNicStates field if non-nil, zero value otherwise.

### GetNicLinkNicStatesOk

`func (o *FiltersVm) GetNicLinkNicStatesOk() (*[]string, bool)`

GetNicLinkNicStatesOk returns a tuple with the NicLinkNicStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkNicStates

`func (o *FiltersVm) SetNicLinkNicStates(v []string)`

SetNicLinkNicStates sets NicLinkNicStates field to given value.

### HasNicLinkNicStates

`func (o *FiltersVm) HasNicLinkNicStates() bool`

HasNicLinkNicStates returns a boolean if a field has been set.

### GetNicLinkNicVmAccountIds

`func (o *FiltersVm) GetNicLinkNicVmAccountIds() []string`

GetNicLinkNicVmAccountIds returns the NicLinkNicVmAccountIds field if non-nil, zero value otherwise.

### GetNicLinkNicVmAccountIdsOk

`func (o *FiltersVm) GetNicLinkNicVmAccountIdsOk() (*[]string, bool)`

GetNicLinkNicVmAccountIdsOk returns a tuple with the NicLinkNicVmAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkNicVmAccountIds

`func (o *FiltersVm) SetNicLinkNicVmAccountIds(v []string)`

SetNicLinkNicVmAccountIds sets NicLinkNicVmAccountIds field to given value.

### HasNicLinkNicVmAccountIds

`func (o *FiltersVm) HasNicLinkNicVmAccountIds() bool`

HasNicLinkNicVmAccountIds returns a boolean if a field has been set.

### GetNicLinkNicVmIds

`func (o *FiltersVm) GetNicLinkNicVmIds() []string`

GetNicLinkNicVmIds returns the NicLinkNicVmIds field if non-nil, zero value otherwise.

### GetNicLinkNicVmIdsOk

`func (o *FiltersVm) GetNicLinkNicVmIdsOk() (*[]string, bool)`

GetNicLinkNicVmIdsOk returns a tuple with the NicLinkNicVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkNicVmIds

`func (o *FiltersVm) SetNicLinkNicVmIds(v []string)`

SetNicLinkNicVmIds sets NicLinkNicVmIds field to given value.

### HasNicLinkNicVmIds

`func (o *FiltersVm) HasNicLinkNicVmIds() bool`

HasNicLinkNicVmIds returns a boolean if a field has been set.

### GetNicLinkPublicIpAccountIds

`func (o *FiltersVm) GetNicLinkPublicIpAccountIds() []string`

GetNicLinkPublicIpAccountIds returns the NicLinkPublicIpAccountIds field if non-nil, zero value otherwise.

### GetNicLinkPublicIpAccountIdsOk

`func (o *FiltersVm) GetNicLinkPublicIpAccountIdsOk() (*[]string, bool)`

GetNicLinkPublicIpAccountIdsOk returns a tuple with the NicLinkPublicIpAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkPublicIpAccountIds

`func (o *FiltersVm) SetNicLinkPublicIpAccountIds(v []string)`

SetNicLinkPublicIpAccountIds sets NicLinkPublicIpAccountIds field to given value.

### HasNicLinkPublicIpAccountIds

`func (o *FiltersVm) HasNicLinkPublicIpAccountIds() bool`

HasNicLinkPublicIpAccountIds returns a boolean if a field has been set.

### GetNicLinkPublicIpLinkPublicIpIds

`func (o *FiltersVm) GetNicLinkPublicIpLinkPublicIpIds() []string`

GetNicLinkPublicIpLinkPublicIpIds returns the NicLinkPublicIpLinkPublicIpIds field if non-nil, zero value otherwise.

### GetNicLinkPublicIpLinkPublicIpIdsOk

`func (o *FiltersVm) GetNicLinkPublicIpLinkPublicIpIdsOk() (*[]string, bool)`

GetNicLinkPublicIpLinkPublicIpIdsOk returns a tuple with the NicLinkPublicIpLinkPublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkPublicIpLinkPublicIpIds

`func (o *FiltersVm) SetNicLinkPublicIpLinkPublicIpIds(v []string)`

SetNicLinkPublicIpLinkPublicIpIds sets NicLinkPublicIpLinkPublicIpIds field to given value.

### HasNicLinkPublicIpLinkPublicIpIds

`func (o *FiltersVm) HasNicLinkPublicIpLinkPublicIpIds() bool`

HasNicLinkPublicIpLinkPublicIpIds returns a boolean if a field has been set.

### GetNicLinkPublicIpPublicIpIds

`func (o *FiltersVm) GetNicLinkPublicIpPublicIpIds() []string`

GetNicLinkPublicIpPublicIpIds returns the NicLinkPublicIpPublicIpIds field if non-nil, zero value otherwise.

### GetNicLinkPublicIpPublicIpIdsOk

`func (o *FiltersVm) GetNicLinkPublicIpPublicIpIdsOk() (*[]string, bool)`

GetNicLinkPublicIpPublicIpIdsOk returns a tuple with the NicLinkPublicIpPublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkPublicIpPublicIpIds

`func (o *FiltersVm) SetNicLinkPublicIpPublicIpIds(v []string)`

SetNicLinkPublicIpPublicIpIds sets NicLinkPublicIpPublicIpIds field to given value.

### HasNicLinkPublicIpPublicIpIds

`func (o *FiltersVm) HasNicLinkPublicIpPublicIpIds() bool`

HasNicLinkPublicIpPublicIpIds returns a boolean if a field has been set.

### GetNicLinkPublicIpPublicIps

`func (o *FiltersVm) GetNicLinkPublicIpPublicIps() []string`

GetNicLinkPublicIpPublicIps returns the NicLinkPublicIpPublicIps field if non-nil, zero value otherwise.

### GetNicLinkPublicIpPublicIpsOk

`func (o *FiltersVm) GetNicLinkPublicIpPublicIpsOk() (*[]string, bool)`

GetNicLinkPublicIpPublicIpsOk returns a tuple with the NicLinkPublicIpPublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicLinkPublicIpPublicIps

`func (o *FiltersVm) SetNicLinkPublicIpPublicIps(v []string)`

SetNicLinkPublicIpPublicIps sets NicLinkPublicIpPublicIps field to given value.

### HasNicLinkPublicIpPublicIps

`func (o *FiltersVm) HasNicLinkPublicIpPublicIps() bool`

HasNicLinkPublicIpPublicIps returns a boolean if a field has been set.

### GetNicMacAddresses

`func (o *FiltersVm) GetNicMacAddresses() []string`

GetNicMacAddresses returns the NicMacAddresses field if non-nil, zero value otherwise.

### GetNicMacAddressesOk

`func (o *FiltersVm) GetNicMacAddressesOk() (*[]string, bool)`

GetNicMacAddressesOk returns a tuple with the NicMacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicMacAddresses

`func (o *FiltersVm) SetNicMacAddresses(v []string)`

SetNicMacAddresses sets NicMacAddresses field to given value.

### HasNicMacAddresses

`func (o *FiltersVm) HasNicMacAddresses() bool`

HasNicMacAddresses returns a boolean if a field has been set.

### GetNicNetIds

`func (o *FiltersVm) GetNicNetIds() []string`

GetNicNetIds returns the NicNetIds field if non-nil, zero value otherwise.

### GetNicNetIdsOk

`func (o *FiltersVm) GetNicNetIdsOk() (*[]string, bool)`

GetNicNetIdsOk returns a tuple with the NicNetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicNetIds

`func (o *FiltersVm) SetNicNetIds(v []string)`

SetNicNetIds sets NicNetIds field to given value.

### HasNicNetIds

`func (o *FiltersVm) HasNicNetIds() bool`

HasNicNetIds returns a boolean if a field has been set.

### GetNicNicIds

`func (o *FiltersVm) GetNicNicIds() []string`

GetNicNicIds returns the NicNicIds field if non-nil, zero value otherwise.

### GetNicNicIdsOk

`func (o *FiltersVm) GetNicNicIdsOk() (*[]string, bool)`

GetNicNicIdsOk returns a tuple with the NicNicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicNicIds

`func (o *FiltersVm) SetNicNicIds(v []string)`

SetNicNicIds sets NicNicIds field to given value.

### HasNicNicIds

`func (o *FiltersVm) HasNicNicIds() bool`

HasNicNicIds returns a boolean if a field has been set.

### GetNicPrivateIpsLinkPublicIpAccountIds

`func (o *FiltersVm) GetNicPrivateIpsLinkPublicIpAccountIds() []string`

GetNicPrivateIpsLinkPublicIpAccountIds returns the NicPrivateIpsLinkPublicIpAccountIds field if non-nil, zero value otherwise.

### GetNicPrivateIpsLinkPublicIpAccountIdsOk

`func (o *FiltersVm) GetNicPrivateIpsLinkPublicIpAccountIdsOk() (*[]string, bool)`

GetNicPrivateIpsLinkPublicIpAccountIdsOk returns a tuple with the NicPrivateIpsLinkPublicIpAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicPrivateIpsLinkPublicIpAccountIds

`func (o *FiltersVm) SetNicPrivateIpsLinkPublicIpAccountIds(v []string)`

SetNicPrivateIpsLinkPublicIpAccountIds sets NicPrivateIpsLinkPublicIpAccountIds field to given value.

### HasNicPrivateIpsLinkPublicIpAccountIds

`func (o *FiltersVm) HasNicPrivateIpsLinkPublicIpAccountIds() bool`

HasNicPrivateIpsLinkPublicIpAccountIds returns a boolean if a field has been set.

### GetNicPrivateIpsLinkPublicIpIds

`func (o *FiltersVm) GetNicPrivateIpsLinkPublicIpIds() []string`

GetNicPrivateIpsLinkPublicIpIds returns the NicPrivateIpsLinkPublicIpIds field if non-nil, zero value otherwise.

### GetNicPrivateIpsLinkPublicIpIdsOk

`func (o *FiltersVm) GetNicPrivateIpsLinkPublicIpIdsOk() (*[]string, bool)`

GetNicPrivateIpsLinkPublicIpIdsOk returns a tuple with the NicPrivateIpsLinkPublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicPrivateIpsLinkPublicIpIds

`func (o *FiltersVm) SetNicPrivateIpsLinkPublicIpIds(v []string)`

SetNicPrivateIpsLinkPublicIpIds sets NicPrivateIpsLinkPublicIpIds field to given value.

### HasNicPrivateIpsLinkPublicIpIds

`func (o *FiltersVm) HasNicPrivateIpsLinkPublicIpIds() bool`

HasNicPrivateIpsLinkPublicIpIds returns a boolean if a field has been set.

### GetNicPrivateIpsPrimaryIp

`func (o *FiltersVm) GetNicPrivateIpsPrimaryIp() bool`

GetNicPrivateIpsPrimaryIp returns the NicPrivateIpsPrimaryIp field if non-nil, zero value otherwise.

### GetNicPrivateIpsPrimaryIpOk

`func (o *FiltersVm) GetNicPrivateIpsPrimaryIpOk() (*bool, bool)`

GetNicPrivateIpsPrimaryIpOk returns a tuple with the NicPrivateIpsPrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicPrivateIpsPrimaryIp

`func (o *FiltersVm) SetNicPrivateIpsPrimaryIp(v bool)`

SetNicPrivateIpsPrimaryIp sets NicPrivateIpsPrimaryIp field to given value.

### HasNicPrivateIpsPrimaryIp

`func (o *FiltersVm) HasNicPrivateIpsPrimaryIp() bool`

HasNicPrivateIpsPrimaryIp returns a boolean if a field has been set.

### GetNicPrivateIpsPrivateIps

`func (o *FiltersVm) GetNicPrivateIpsPrivateIps() []string`

GetNicPrivateIpsPrivateIps returns the NicPrivateIpsPrivateIps field if non-nil, zero value otherwise.

### GetNicPrivateIpsPrivateIpsOk

`func (o *FiltersVm) GetNicPrivateIpsPrivateIpsOk() (*[]string, bool)`

GetNicPrivateIpsPrivateIpsOk returns a tuple with the NicPrivateIpsPrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicPrivateIpsPrivateIps

`func (o *FiltersVm) SetNicPrivateIpsPrivateIps(v []string)`

SetNicPrivateIpsPrivateIps sets NicPrivateIpsPrivateIps field to given value.

### HasNicPrivateIpsPrivateIps

`func (o *FiltersVm) HasNicPrivateIpsPrivateIps() bool`

HasNicPrivateIpsPrivateIps returns a boolean if a field has been set.

### GetNicSecurityGroupIds

`func (o *FiltersVm) GetNicSecurityGroupIds() []string`

GetNicSecurityGroupIds returns the NicSecurityGroupIds field if non-nil, zero value otherwise.

### GetNicSecurityGroupIdsOk

`func (o *FiltersVm) GetNicSecurityGroupIdsOk() (*[]string, bool)`

GetNicSecurityGroupIdsOk returns a tuple with the NicSecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicSecurityGroupIds

`func (o *FiltersVm) SetNicSecurityGroupIds(v []string)`

SetNicSecurityGroupIds sets NicSecurityGroupIds field to given value.

### HasNicSecurityGroupIds

`func (o *FiltersVm) HasNicSecurityGroupIds() bool`

HasNicSecurityGroupIds returns a boolean if a field has been set.

### GetNicSecurityGroupNames

`func (o *FiltersVm) GetNicSecurityGroupNames() []string`

GetNicSecurityGroupNames returns the NicSecurityGroupNames field if non-nil, zero value otherwise.

### GetNicSecurityGroupNamesOk

`func (o *FiltersVm) GetNicSecurityGroupNamesOk() (*[]string, bool)`

GetNicSecurityGroupNamesOk returns a tuple with the NicSecurityGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicSecurityGroupNames

`func (o *FiltersVm) SetNicSecurityGroupNames(v []string)`

SetNicSecurityGroupNames sets NicSecurityGroupNames field to given value.

### HasNicSecurityGroupNames

`func (o *FiltersVm) HasNicSecurityGroupNames() bool`

HasNicSecurityGroupNames returns a boolean if a field has been set.

### GetNicStates

`func (o *FiltersVm) GetNicStates() []string`

GetNicStates returns the NicStates field if non-nil, zero value otherwise.

### GetNicStatesOk

`func (o *FiltersVm) GetNicStatesOk() (*[]string, bool)`

GetNicStatesOk returns a tuple with the NicStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicStates

`func (o *FiltersVm) SetNicStates(v []string)`

SetNicStates sets NicStates field to given value.

### HasNicStates

`func (o *FiltersVm) HasNicStates() bool`

HasNicStates returns a boolean if a field has been set.

### GetNicSubnetIds

`func (o *FiltersVm) GetNicSubnetIds() []string`

GetNicSubnetIds returns the NicSubnetIds field if non-nil, zero value otherwise.

### GetNicSubnetIdsOk

`func (o *FiltersVm) GetNicSubnetIdsOk() (*[]string, bool)`

GetNicSubnetIdsOk returns a tuple with the NicSubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicSubnetIds

`func (o *FiltersVm) SetNicSubnetIds(v []string)`

SetNicSubnetIds sets NicSubnetIds field to given value.

### HasNicSubnetIds

`func (o *FiltersVm) HasNicSubnetIds() bool`

HasNicSubnetIds returns a boolean if a field has been set.

### GetNicSubregionNames

`func (o *FiltersVm) GetNicSubregionNames() []string`

GetNicSubregionNames returns the NicSubregionNames field if non-nil, zero value otherwise.

### GetNicSubregionNamesOk

`func (o *FiltersVm) GetNicSubregionNamesOk() (*[]string, bool)`

GetNicSubregionNamesOk returns a tuple with the NicSubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicSubregionNames

`func (o *FiltersVm) SetNicSubregionNames(v []string)`

SetNicSubregionNames sets NicSubregionNames field to given value.

### HasNicSubregionNames

`func (o *FiltersVm) HasNicSubregionNames() bool`

HasNicSubregionNames returns a boolean if a field has been set.

### GetPlatforms

`func (o *FiltersVm) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *FiltersVm) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *FiltersVm) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *FiltersVm) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetPrivateIps

`func (o *FiltersVm) GetPrivateIps() []string`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *FiltersVm) GetPrivateIpsOk() (*[]string, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *FiltersVm) SetPrivateIps(v []string)`

SetPrivateIps sets PrivateIps field to given value.

### HasPrivateIps

`func (o *FiltersVm) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### GetProductCodes

`func (o *FiltersVm) GetProductCodes() []string`

GetProductCodes returns the ProductCodes field if non-nil, zero value otherwise.

### GetProductCodesOk

`func (o *FiltersVm) GetProductCodesOk() (*[]string, bool)`

GetProductCodesOk returns a tuple with the ProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCodes

`func (o *FiltersVm) SetProductCodes(v []string)`

SetProductCodes sets ProductCodes field to given value.

### HasProductCodes

`func (o *FiltersVm) HasProductCodes() bool`

HasProductCodes returns a boolean if a field has been set.

### GetPublicIps

`func (o *FiltersVm) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *FiltersVm) GetPublicIpsOk() (*[]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *FiltersVm) SetPublicIps(v []string)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *FiltersVm) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### GetReservationIds

`func (o *FiltersVm) GetReservationIds() []string`

GetReservationIds returns the ReservationIds field if non-nil, zero value otherwise.

### GetReservationIdsOk

`func (o *FiltersVm) GetReservationIdsOk() (*[]string, bool)`

GetReservationIdsOk returns a tuple with the ReservationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationIds

`func (o *FiltersVm) SetReservationIds(v []string)`

SetReservationIds sets ReservationIds field to given value.

### HasReservationIds

`func (o *FiltersVm) HasReservationIds() bool`

HasReservationIds returns a boolean if a field has been set.

### GetRootDeviceNames

`func (o *FiltersVm) GetRootDeviceNames() []string`

GetRootDeviceNames returns the RootDeviceNames field if non-nil, zero value otherwise.

### GetRootDeviceNamesOk

`func (o *FiltersVm) GetRootDeviceNamesOk() (*[]string, bool)`

GetRootDeviceNamesOk returns a tuple with the RootDeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceNames

`func (o *FiltersVm) SetRootDeviceNames(v []string)`

SetRootDeviceNames sets RootDeviceNames field to given value.

### HasRootDeviceNames

`func (o *FiltersVm) HasRootDeviceNames() bool`

HasRootDeviceNames returns a boolean if a field has been set.

### GetRootDeviceTypes

`func (o *FiltersVm) GetRootDeviceTypes() []string`

GetRootDeviceTypes returns the RootDeviceTypes field if non-nil, zero value otherwise.

### GetRootDeviceTypesOk

`func (o *FiltersVm) GetRootDeviceTypesOk() (*[]string, bool)`

GetRootDeviceTypesOk returns a tuple with the RootDeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceTypes

`func (o *FiltersVm) SetRootDeviceTypes(v []string)`

SetRootDeviceTypes sets RootDeviceTypes field to given value.

### HasRootDeviceTypes

`func (o *FiltersVm) HasRootDeviceTypes() bool`

HasRootDeviceTypes returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *FiltersVm) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *FiltersVm) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *FiltersVm) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *FiltersVm) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSecurityGroupNames

`func (o *FiltersVm) GetSecurityGroupNames() []string`

GetSecurityGroupNames returns the SecurityGroupNames field if non-nil, zero value otherwise.

### GetSecurityGroupNamesOk

`func (o *FiltersVm) GetSecurityGroupNamesOk() (*[]string, bool)`

GetSecurityGroupNamesOk returns a tuple with the SecurityGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupNames

`func (o *FiltersVm) SetSecurityGroupNames(v []string)`

SetSecurityGroupNames sets SecurityGroupNames field to given value.

### HasSecurityGroupNames

`func (o *FiltersVm) HasSecurityGroupNames() bool`

HasSecurityGroupNames returns a boolean if a field has been set.

### GetStateReasonCodes

`func (o *FiltersVm) GetStateReasonCodes() []int32`

GetStateReasonCodes returns the StateReasonCodes field if non-nil, zero value otherwise.

### GetStateReasonCodesOk

`func (o *FiltersVm) GetStateReasonCodesOk() (*[]int32, bool)`

GetStateReasonCodesOk returns a tuple with the StateReasonCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReasonCodes

`func (o *FiltersVm) SetStateReasonCodes(v []int32)`

SetStateReasonCodes sets StateReasonCodes field to given value.

### HasStateReasonCodes

`func (o *FiltersVm) HasStateReasonCodes() bool`

HasStateReasonCodes returns a boolean if a field has been set.

### GetStateReasonMessages

`func (o *FiltersVm) GetStateReasonMessages() []string`

GetStateReasonMessages returns the StateReasonMessages field if non-nil, zero value otherwise.

### GetStateReasonMessagesOk

`func (o *FiltersVm) GetStateReasonMessagesOk() (*[]string, bool)`

GetStateReasonMessagesOk returns a tuple with the StateReasonMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReasonMessages

`func (o *FiltersVm) SetStateReasonMessages(v []string)`

SetStateReasonMessages sets StateReasonMessages field to given value.

### HasStateReasonMessages

`func (o *FiltersVm) HasStateReasonMessages() bool`

HasStateReasonMessages returns a boolean if a field has been set.

### GetStateReasons

`func (o *FiltersVm) GetStateReasons() []string`

GetStateReasons returns the StateReasons field if non-nil, zero value otherwise.

### GetStateReasonsOk

`func (o *FiltersVm) GetStateReasonsOk() (*[]string, bool)`

GetStateReasonsOk returns a tuple with the StateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReasons

`func (o *FiltersVm) SetStateReasons(v []string)`

SetStateReasons sets StateReasons field to given value.

### HasStateReasons

`func (o *FiltersVm) HasStateReasons() bool`

HasStateReasons returns a boolean if a field has been set.

### GetSubnetIds

`func (o *FiltersVm) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *FiltersVm) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *FiltersVm) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *FiltersVm) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### GetSubregionNames

`func (o *FiltersVm) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersVm) GetSubregionNamesOk() (*[]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionNames

`func (o *FiltersVm) SetSubregionNames(v []string)`

SetSubregionNames sets SubregionNames field to given value.

### HasSubregionNames

`func (o *FiltersVm) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersVm) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersVm) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersVm) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersVm) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersVm) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersVm) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersVm) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersVm) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersVm) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersVm) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersVm) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersVm) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTenancies

`func (o *FiltersVm) GetTenancies() []string`

GetTenancies returns the Tenancies field if non-nil, zero value otherwise.

### GetTenanciesOk

`func (o *FiltersVm) GetTenanciesOk() (*[]string, bool)`

GetTenanciesOk returns a tuple with the Tenancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancies

`func (o *FiltersVm) SetTenancies(v []string)`

SetTenancies sets Tenancies field to given value.

### HasTenancies

`func (o *FiltersVm) HasTenancies() bool`

HasTenancies returns a boolean if a field has been set.

### GetVmIds

`func (o *FiltersVm) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *FiltersVm) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *FiltersVm) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *FiltersVm) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### GetVmSecurityGroupIds

`func (o *FiltersVm) GetVmSecurityGroupIds() []string`

GetVmSecurityGroupIds returns the VmSecurityGroupIds field if non-nil, zero value otherwise.

### GetVmSecurityGroupIdsOk

`func (o *FiltersVm) GetVmSecurityGroupIdsOk() (*[]string, bool)`

GetVmSecurityGroupIdsOk returns a tuple with the VmSecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSecurityGroupIds

`func (o *FiltersVm) SetVmSecurityGroupIds(v []string)`

SetVmSecurityGroupIds sets VmSecurityGroupIds field to given value.

### HasVmSecurityGroupIds

`func (o *FiltersVm) HasVmSecurityGroupIds() bool`

HasVmSecurityGroupIds returns a boolean if a field has been set.

### GetVmSecurityGroupNames

`func (o *FiltersVm) GetVmSecurityGroupNames() []string`

GetVmSecurityGroupNames returns the VmSecurityGroupNames field if non-nil, zero value otherwise.

### GetVmSecurityGroupNamesOk

`func (o *FiltersVm) GetVmSecurityGroupNamesOk() (*[]string, bool)`

GetVmSecurityGroupNamesOk returns a tuple with the VmSecurityGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSecurityGroupNames

`func (o *FiltersVm) SetVmSecurityGroupNames(v []string)`

SetVmSecurityGroupNames sets VmSecurityGroupNames field to given value.

### HasVmSecurityGroupNames

`func (o *FiltersVm) HasVmSecurityGroupNames() bool`

HasVmSecurityGroupNames returns a boolean if a field has been set.

### GetVmStateCodes

`func (o *FiltersVm) GetVmStateCodes() []int32`

GetVmStateCodes returns the VmStateCodes field if non-nil, zero value otherwise.

### GetVmStateCodesOk

`func (o *FiltersVm) GetVmStateCodesOk() (*[]int32, bool)`

GetVmStateCodesOk returns a tuple with the VmStateCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStateCodes

`func (o *FiltersVm) SetVmStateCodes(v []int32)`

SetVmStateCodes sets VmStateCodes field to given value.

### HasVmStateCodes

`func (o *FiltersVm) HasVmStateCodes() bool`

HasVmStateCodes returns a boolean if a field has been set.

### GetVmStateNames

`func (o *FiltersVm) GetVmStateNames() []string`

GetVmStateNames returns the VmStateNames field if non-nil, zero value otherwise.

### GetVmStateNamesOk

`func (o *FiltersVm) GetVmStateNamesOk() (*[]string, bool)`

GetVmStateNamesOk returns a tuple with the VmStateNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStateNames

`func (o *FiltersVm) SetVmStateNames(v []string)`

SetVmStateNames sets VmStateNames field to given value.

### HasVmStateNames

`func (o *FiltersVm) HasVmStateNames() bool`

HasVmStateNames returns a boolean if a field has been set.

### GetVmTypes

`func (o *FiltersVm) GetVmTypes() []string`

GetVmTypes returns the VmTypes field if non-nil, zero value otherwise.

### GetVmTypesOk

`func (o *FiltersVm) GetVmTypesOk() (*[]string, bool)`

GetVmTypesOk returns a tuple with the VmTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTypes

`func (o *FiltersVm) SetVmTypes(v []string)`

SetVmTypes sets VmTypes field to given value.

### HasVmTypes

`func (o *FiltersVm) HasVmTypes() bool`

HasVmTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


