# DeleteVpnConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VpnConnectionId** | **string** | The ID of the VPN connection you want to delete. | 

## Methods

### NewDeleteVpnConnectionRequest

`func NewDeleteVpnConnectionRequest(vpnConnectionId string, ) *DeleteVpnConnectionRequest`

NewDeleteVpnConnectionRequest instantiates a new DeleteVpnConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVpnConnectionRequestWithDefaults

`func NewDeleteVpnConnectionRequestWithDefaults() *DeleteVpnConnectionRequest`

NewDeleteVpnConnectionRequestWithDefaults instantiates a new DeleteVpnConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteVpnConnectionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteVpnConnectionRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteVpnConnectionRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteVpnConnectionRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVpnConnectionId

`func (o *DeleteVpnConnectionRequest) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *DeleteVpnConnectionRequest) GetVpnConnectionIdOk() (*string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionId

`func (o *DeleteVpnConnectionRequest) SetVpnConnectionId(v string)`

SetVpnConnectionId sets VpnConnectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


