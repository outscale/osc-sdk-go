# DeletePublicIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PublicIp** | Pointer to **string** | The public IP. In the public Cloud, this parameter is required. | [optional] 
**PublicIpId** | Pointer to **string** | The ID representing the association of the public IP with the VM or the NIC. In a Net, this parameter is required. | [optional] 

## Methods

### NewDeletePublicIpRequest

`func NewDeletePublicIpRequest() *DeletePublicIpRequest`

NewDeletePublicIpRequest instantiates a new DeletePublicIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePublicIpRequestWithDefaults

`func NewDeletePublicIpRequestWithDefaults() *DeletePublicIpRequest`

NewDeletePublicIpRequestWithDefaults instantiates a new DeletePublicIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeletePublicIpRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeletePublicIpRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeletePublicIpRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeletePublicIpRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPublicIp

`func (o *DeletePublicIpRequest) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *DeletePublicIpRequest) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *DeletePublicIpRequest) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *DeletePublicIpRequest) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicIpId

`func (o *DeletePublicIpRequest) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *DeletePublicIpRequest) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *DeletePublicIpRequest) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.

### HasPublicIpId

`func (o *DeletePublicIpRequest) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


