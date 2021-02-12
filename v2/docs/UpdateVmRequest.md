# UpdateVmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDeviceMappings** | Pointer to [**[]BlockDeviceMappingVmUpdate**](BlockDeviceMappingVmUpdate.md) | One or more block device mappings of the VM. | [optional] 
**BsuOptimized** | Pointer to **bool** | If true, the VM is optimized for BSU I/O. | [optional] 
**DeletionProtection** | Pointer to **bool** | If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**IsSourceDestChecked** | Pointer to **bool** | (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**KeypairName** | Pointer to **string** | The name of the keypair.&lt;br /&gt; To complete the replacement, manually replace the old public key with the new public key in the ~/.ssh/authorized_keys file located in the VM. Restart the VM to apply the change. | [optional] 
**Performance** | Pointer to **string** | The performance of the VM (&#x60;medium&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security groups for the VM. | [optional] 
**UserData** | Pointer to **string** | The Base64-encoded MIME user data. | [optional] 
**VmId** | **string** | The ID of the VM. | 
**VmInitiatedShutdownBehavior** | Pointer to **string** | The VM behavior when you stop it. By default or if set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is terminated. | [optional] 
**VmType** | Pointer to **string** | The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). | [optional] 

## Methods

### NewUpdateVmRequest

`func NewUpdateVmRequest(vmId string, ) *UpdateVmRequest`

NewUpdateVmRequest instantiates a new UpdateVmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVmRequestWithDefaults

`func NewUpdateVmRequestWithDefaults() *UpdateVmRequest`

NewUpdateVmRequestWithDefaults instantiates a new UpdateVmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockDeviceMappings

`func (o *UpdateVmRequest) GetBlockDeviceMappings() []BlockDeviceMappingVmUpdate`

GetBlockDeviceMappings returns the BlockDeviceMappings field if non-nil, zero value otherwise.

### GetBlockDeviceMappingsOk

`func (o *UpdateVmRequest) GetBlockDeviceMappingsOk() (*[]BlockDeviceMappingVmUpdate, bool)`

GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappings

`func (o *UpdateVmRequest) SetBlockDeviceMappings(v []BlockDeviceMappingVmUpdate)`

SetBlockDeviceMappings sets BlockDeviceMappings field to given value.

### HasBlockDeviceMappings

`func (o *UpdateVmRequest) HasBlockDeviceMappings() bool`

HasBlockDeviceMappings returns a boolean if a field has been set.

### GetBsuOptimized

`func (o *UpdateVmRequest) GetBsuOptimized() bool`

GetBsuOptimized returns the BsuOptimized field if non-nil, zero value otherwise.

### GetBsuOptimizedOk

`func (o *UpdateVmRequest) GetBsuOptimizedOk() (*bool, bool)`

GetBsuOptimizedOk returns a tuple with the BsuOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsuOptimized

`func (o *UpdateVmRequest) SetBsuOptimized(v bool)`

SetBsuOptimized sets BsuOptimized field to given value.

### HasBsuOptimized

`func (o *UpdateVmRequest) HasBsuOptimized() bool`

HasBsuOptimized returns a boolean if a field has been set.

### GetDeletionProtection

`func (o *UpdateVmRequest) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *UpdateVmRequest) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *UpdateVmRequest) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.

### HasDeletionProtection

`func (o *UpdateVmRequest) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateVmRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateVmRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateVmRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateVmRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetIsSourceDestChecked

`func (o *UpdateVmRequest) GetIsSourceDestChecked() bool`

GetIsSourceDestChecked returns the IsSourceDestChecked field if non-nil, zero value otherwise.

### GetIsSourceDestCheckedOk

`func (o *UpdateVmRequest) GetIsSourceDestCheckedOk() (*bool, bool)`

GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceDestChecked

`func (o *UpdateVmRequest) SetIsSourceDestChecked(v bool)`

SetIsSourceDestChecked sets IsSourceDestChecked field to given value.

### HasIsSourceDestChecked

`func (o *UpdateVmRequest) HasIsSourceDestChecked() bool`

HasIsSourceDestChecked returns a boolean if a field has been set.

### GetKeypairName

`func (o *UpdateVmRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *UpdateVmRequest) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *UpdateVmRequest) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *UpdateVmRequest) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetPerformance

`func (o *UpdateVmRequest) GetPerformance() string`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *UpdateVmRequest) GetPerformanceOk() (*string, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *UpdateVmRequest) SetPerformance(v string)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *UpdateVmRequest) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *UpdateVmRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *UpdateVmRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *UpdateVmRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *UpdateVmRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetUserData

`func (o *UpdateVmRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *UpdateVmRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *UpdateVmRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *UpdateVmRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetVmId

`func (o *UpdateVmRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *UpdateVmRequest) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *UpdateVmRequest) SetVmId(v string)`

SetVmId sets VmId field to given value.


### GetVmInitiatedShutdownBehavior

`func (o *UpdateVmRequest) GetVmInitiatedShutdownBehavior() string`

GetVmInitiatedShutdownBehavior returns the VmInitiatedShutdownBehavior field if non-nil, zero value otherwise.

### GetVmInitiatedShutdownBehaviorOk

`func (o *UpdateVmRequest) GetVmInitiatedShutdownBehaviorOk() (*string, bool)`

GetVmInitiatedShutdownBehaviorOk returns a tuple with the VmInitiatedShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInitiatedShutdownBehavior

`func (o *UpdateVmRequest) SetVmInitiatedShutdownBehavior(v string)`

SetVmInitiatedShutdownBehavior sets VmInitiatedShutdownBehavior field to given value.

### HasVmInitiatedShutdownBehavior

`func (o *UpdateVmRequest) HasVmInitiatedShutdownBehavior() bool`

HasVmInitiatedShutdownBehavior returns a boolean if a field has been set.

### GetVmType

`func (o *UpdateVmRequest) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *UpdateVmRequest) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *UpdateVmRequest) SetVmType(v string)`

SetVmType sets VmType field to given value.

### HasVmType

`func (o *UpdateVmRequest) HasVmType() bool`

HasVmType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


