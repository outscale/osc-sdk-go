# UnlinkPublicIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LinkPublicIpId** | Pointer to **string** | The ID representing the association of the EIP with the VM or the NIC. This parameter is required unless you use the &#x60;PublicIp&#x60; parameter. | [optional] 
**PublicIp** | Pointer to **string** | The EIP. This parameter is required unless you use the &#x60;LinkPublicIpId&#x60; parameter. | [optional] 

## Methods

### NewUnlinkPublicIpRequest

`func NewUnlinkPublicIpRequest() *UnlinkPublicIpRequest`

NewUnlinkPublicIpRequest instantiates a new UnlinkPublicIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkPublicIpRequestWithDefaults

`func NewUnlinkPublicIpRequestWithDefaults() *UnlinkPublicIpRequest`

NewUnlinkPublicIpRequestWithDefaults instantiates a new UnlinkPublicIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UnlinkPublicIpRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkPublicIpRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkPublicIpRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkPublicIpRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLinkPublicIpId

`func (o *UnlinkPublicIpRequest) GetLinkPublicIpId() string`

GetLinkPublicIpId returns the LinkPublicIpId field if non-nil, zero value otherwise.

### GetLinkPublicIpIdOk

`func (o *UnlinkPublicIpRequest) GetLinkPublicIpIdOk() (*string, bool)`

GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIpId

`func (o *UnlinkPublicIpRequest) SetLinkPublicIpId(v string)`

SetLinkPublicIpId sets LinkPublicIpId field to given value.

### HasLinkPublicIpId

`func (o *UnlinkPublicIpRequest) HasLinkPublicIpId() bool`

HasLinkPublicIpId returns a boolean if a field has been set.

### GetPublicIp

`func (o *UnlinkPublicIpRequest) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *UnlinkPublicIpRequest) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *UnlinkPublicIpRequest) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *UnlinkPublicIpRequest) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


