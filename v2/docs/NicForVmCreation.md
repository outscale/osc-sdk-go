# NicForVmCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If true, the NIC is deleted when the VM is terminated. You can specify this parameter only for a new NIC. To modify this value for an existing NIC, see [UpdateNic](#updatenic). | [optional] 
**Description** | Pointer to **string** | The description of the NIC, if you are creating a NIC when creating the VM. | [optional] 
**DeviceNumber** | Pointer to **int32** | The index of the VM device for the NIC attachment (between 0 and 7, both included). This parameter is required if you create a NIC when creating the VM. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC, if you are attaching an existing NIC when creating a VM. | [optional] 
**PrivateIps** | Pointer to [**[]PrivateIpLight**](PrivateIpLight.md) | One or more private IPs to assign to the NIC, if you create a NIC when creating a VM. Only one private IP can be the primary private IP. | [optional] 
**SecondaryPrivateIpCount** | Pointer to **int32** | The number of secondary private IPs, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP in the &#x60;PrivateIps&#x60; parameter. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security groups for the NIC, if you create a NIC when creating a VM. | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet for the NIC, if you create a NIC when creating a VM. This parameter is required if you create a NIC when creating the VM. | [optional] 

## Methods

### NewNicForVmCreation

`func NewNicForVmCreation() *NicForVmCreation`

NewNicForVmCreation instantiates a new NicForVmCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicForVmCreationWithDefaults

`func NewNicForVmCreationWithDefaults() *NicForVmCreation`

NewNicForVmCreationWithDefaults instantiates a new NicForVmCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnVmDeletion

`func (o *NicForVmCreation) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *NicForVmCreation) GetDeleteOnVmDeletionOk() (*bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnVmDeletion

`func (o *NicForVmCreation) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion sets DeleteOnVmDeletion field to given value.

### HasDeleteOnVmDeletion

`func (o *NicForVmCreation) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### GetDescription

`func (o *NicForVmCreation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NicForVmCreation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NicForVmCreation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NicForVmCreation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceNumber

`func (o *NicForVmCreation) GetDeviceNumber() int32`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *NicForVmCreation) GetDeviceNumberOk() (*int32, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNumber

`func (o *NicForVmCreation) SetDeviceNumber(v int32)`

SetDeviceNumber sets DeviceNumber field to given value.

### HasDeviceNumber

`func (o *NicForVmCreation) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### GetNicId

`func (o *NicForVmCreation) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *NicForVmCreation) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *NicForVmCreation) SetNicId(v string)`

SetNicId sets NicId field to given value.

### HasNicId

`func (o *NicForVmCreation) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### GetPrivateIps

`func (o *NicForVmCreation) GetPrivateIps() []PrivateIpLight`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *NicForVmCreation) GetPrivateIpsOk() (*[]PrivateIpLight, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *NicForVmCreation) SetPrivateIps(v []PrivateIpLight)`

SetPrivateIps sets PrivateIps field to given value.

### HasPrivateIps

`func (o *NicForVmCreation) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### GetSecondaryPrivateIpCount

`func (o *NicForVmCreation) GetSecondaryPrivateIpCount() int32`

GetSecondaryPrivateIpCount returns the SecondaryPrivateIpCount field if non-nil, zero value otherwise.

### GetSecondaryPrivateIpCountOk

`func (o *NicForVmCreation) GetSecondaryPrivateIpCountOk() (*int32, bool)`

GetSecondaryPrivateIpCountOk returns a tuple with the SecondaryPrivateIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPrivateIpCount

`func (o *NicForVmCreation) SetSecondaryPrivateIpCount(v int32)`

SetSecondaryPrivateIpCount sets SecondaryPrivateIpCount field to given value.

### HasSecondaryPrivateIpCount

`func (o *NicForVmCreation) HasSecondaryPrivateIpCount() bool`

HasSecondaryPrivateIpCount returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *NicForVmCreation) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *NicForVmCreation) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *NicForVmCreation) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *NicForVmCreation) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSubnetId

`func (o *NicForVmCreation) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *NicForVmCreation) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *NicForVmCreation) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *NicForVmCreation) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


