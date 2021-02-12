# LinkPrivateIpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRelink** | Pointer to **bool** | If true, allows an IP address that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | **string** | The ID of the NIC. | 
**PrivateIps** | Pointer to **[]string** | The secondary private IP address or addresses you want to assign to the NIC within the IP address range of the Subnet. | [optional] 
**SecondaryPrivateIpCount** | Pointer to **int32** | The number of secondary private IP addresses to assign to the NIC. | [optional] 

## Methods

### NewLinkPrivateIpsRequest

`func NewLinkPrivateIpsRequest(nicId string, ) *LinkPrivateIpsRequest`

NewLinkPrivateIpsRequest instantiates a new LinkPrivateIpsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPrivateIpsRequestWithDefaults

`func NewLinkPrivateIpsRequestWithDefaults() *LinkPrivateIpsRequest`

NewLinkPrivateIpsRequestWithDefaults instantiates a new LinkPrivateIpsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRelink

`func (o *LinkPrivateIpsRequest) GetAllowRelink() bool`

GetAllowRelink returns the AllowRelink field if non-nil, zero value otherwise.

### GetAllowRelinkOk

`func (o *LinkPrivateIpsRequest) GetAllowRelinkOk() (*bool, bool)`

GetAllowRelinkOk returns a tuple with the AllowRelink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRelink

`func (o *LinkPrivateIpsRequest) SetAllowRelink(v bool)`

SetAllowRelink sets AllowRelink field to given value.

### HasAllowRelink

`func (o *LinkPrivateIpsRequest) HasAllowRelink() bool`

HasAllowRelink returns a boolean if a field has been set.

### GetDryRun

`func (o *LinkPrivateIpsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkPrivateIpsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkPrivateIpsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkPrivateIpsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNicId

`func (o *LinkPrivateIpsRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *LinkPrivateIpsRequest) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *LinkPrivateIpsRequest) SetNicId(v string)`

SetNicId sets NicId field to given value.


### GetPrivateIps

`func (o *LinkPrivateIpsRequest) GetPrivateIps() []string`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *LinkPrivateIpsRequest) GetPrivateIpsOk() (*[]string, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *LinkPrivateIpsRequest) SetPrivateIps(v []string)`

SetPrivateIps sets PrivateIps field to given value.

### HasPrivateIps

`func (o *LinkPrivateIpsRequest) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### GetSecondaryPrivateIpCount

`func (o *LinkPrivateIpsRequest) GetSecondaryPrivateIpCount() int32`

GetSecondaryPrivateIpCount returns the SecondaryPrivateIpCount field if non-nil, zero value otherwise.

### GetSecondaryPrivateIpCountOk

`func (o *LinkPrivateIpsRequest) GetSecondaryPrivateIpCountOk() (*int32, bool)`

GetSecondaryPrivateIpCountOk returns a tuple with the SecondaryPrivateIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPrivateIpCount

`func (o *LinkPrivateIpsRequest) SetSecondaryPrivateIpCount(v int32)`

SetSecondaryPrivateIpCount sets SecondaryPrivateIpCount field to given value.

### HasSecondaryPrivateIpCount

`func (o *LinkPrivateIpsRequest) HasSecondaryPrivateIpCount() bool`

HasSecondaryPrivateIpCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


