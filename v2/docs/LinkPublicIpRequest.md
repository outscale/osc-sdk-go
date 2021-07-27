# LinkPublicIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRelink** | Pointer to **bool** | If true, allows the EIP to be associated with the VM or NIC that you specify even if it is already associated with another VM or NIC. If false, prevents the EIP from being associated with the VM or NIC that you specify if it is already associated with another VM or NIC. (By default, true in the public Cloud, false in a Net.) | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | Pointer to **string** | (Net only) The ID of the NIC. This parameter is required if the VM has more than one NIC attached. Otherwise, you need to specify the &#x60;VmId&#x60; parameter instead. You cannot specify both parameters at the same time. | [optional] 
**PrivateIp** | Pointer to **string** | (Net only) The primary or secondary private IP address of the specified NIC. By default, the primary private IP address. | [optional] 
**PublicIp** | Pointer to **string** | The EIP. This parameter is required unless you use the &#x60;PublicIpId&#x60; parameter. | [optional] 
**PublicIpId** | Pointer to **string** | The allocation ID of the EIP. This parameter is required unless you use the &#x60;PublicIp&#x60; parameter. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM.&lt;br /&gt; - In the public Cloud, this parameter is required.&lt;br /&gt; - In a Net, this parameter is required if the VM has only one NIC. Otherwise, you need to specify the &#x60;NicId&#x60; parameter instead. You cannot specify both parameters at the same time. | [optional] 

## Methods

### NewLinkPublicIpRequest

`func NewLinkPublicIpRequest() *LinkPublicIpRequest`

NewLinkPublicIpRequest instantiates a new LinkPublicIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPublicIpRequestWithDefaults

`func NewLinkPublicIpRequestWithDefaults() *LinkPublicIpRequest`

NewLinkPublicIpRequestWithDefaults instantiates a new LinkPublicIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRelink

`func (o *LinkPublicIpRequest) GetAllowRelink() bool`

GetAllowRelink returns the AllowRelink field if non-nil, zero value otherwise.

### GetAllowRelinkOk

`func (o *LinkPublicIpRequest) GetAllowRelinkOk() (*bool, bool)`

GetAllowRelinkOk returns a tuple with the AllowRelink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRelink

`func (o *LinkPublicIpRequest) SetAllowRelink(v bool)`

SetAllowRelink sets AllowRelink field to given value.

### HasAllowRelink

`func (o *LinkPublicIpRequest) HasAllowRelink() bool`

HasAllowRelink returns a boolean if a field has been set.

### GetDryRun

`func (o *LinkPublicIpRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkPublicIpRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkPublicIpRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkPublicIpRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNicId

`func (o *LinkPublicIpRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *LinkPublicIpRequest) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *LinkPublicIpRequest) SetNicId(v string)`

SetNicId sets NicId field to given value.

### HasNicId

`func (o *LinkPublicIpRequest) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### GetPrivateIp

`func (o *LinkPublicIpRequest) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *LinkPublicIpRequest) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *LinkPublicIpRequest) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *LinkPublicIpRequest) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *LinkPublicIpRequest) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *LinkPublicIpRequest) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *LinkPublicIpRequest) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *LinkPublicIpRequest) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicIpId

`func (o *LinkPublicIpRequest) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *LinkPublicIpRequest) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *LinkPublicIpRequest) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.

### HasPublicIpId

`func (o *LinkPublicIpRequest) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### GetVmId

`func (o *LinkPublicIpRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *LinkPublicIpRequest) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *LinkPublicIpRequest) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *LinkPublicIpRequest) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


