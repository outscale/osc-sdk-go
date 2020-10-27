# UnlinkPrivateIpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | **string** | The ID of the NIC. | 
**PrivateIps** | **[]string** | One or more secondary private IP addresses you want to unassign from the NIC. | 

## Methods

### NewUnlinkPrivateIpsRequest

`func NewUnlinkPrivateIpsRequest(nicId string, privateIps []string, ) *UnlinkPrivateIpsRequest`

NewUnlinkPrivateIpsRequest instantiates a new UnlinkPrivateIpsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkPrivateIpsRequestWithDefaults

`func NewUnlinkPrivateIpsRequestWithDefaults() *UnlinkPrivateIpsRequest`

NewUnlinkPrivateIpsRequestWithDefaults instantiates a new UnlinkPrivateIpsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UnlinkPrivateIpsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkPrivateIpsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkPrivateIpsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkPrivateIpsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNicId

`func (o *UnlinkPrivateIpsRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *UnlinkPrivateIpsRequest) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *UnlinkPrivateIpsRequest) SetNicId(v string)`

SetNicId sets NicId field to given value.


### GetPrivateIps

`func (o *UnlinkPrivateIpsRequest) GetPrivateIps() []string`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *UnlinkPrivateIpsRequest) GetPrivateIpsOk() (*[]string, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *UnlinkPrivateIpsRequest) SetPrivateIps(v []string)`

SetPrivateIps sets PrivateIps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


