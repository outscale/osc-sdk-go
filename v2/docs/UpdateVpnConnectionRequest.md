# UpdateVpnConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGatewayId** | Pointer to **string** | The ID of the client gateway. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VirtualGatewayId** | Pointer to **string** | The ID of the virtual gateway. | [optional] 
**VpnConnectionId** | **string** | The ID of the VPN connection you want to modify. | 
**VpnOptions** | Pointer to [**VpnOptionsToUpdate**](VpnOptionsToUpdate.md) |  | [optional] 

## Methods

### NewUpdateVpnConnectionRequest

`func NewUpdateVpnConnectionRequest(vpnConnectionId string, ) *UpdateVpnConnectionRequest`

NewUpdateVpnConnectionRequest instantiates a new UpdateVpnConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVpnConnectionRequestWithDefaults

`func NewUpdateVpnConnectionRequestWithDefaults() *UpdateVpnConnectionRequest`

NewUpdateVpnConnectionRequestWithDefaults instantiates a new UpdateVpnConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientGatewayId

`func (o *UpdateVpnConnectionRequest) GetClientGatewayId() string`

GetClientGatewayId returns the ClientGatewayId field if non-nil, zero value otherwise.

### GetClientGatewayIdOk

`func (o *UpdateVpnConnectionRequest) GetClientGatewayIdOk() (*string, bool)`

GetClientGatewayIdOk returns a tuple with the ClientGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGatewayId

`func (o *UpdateVpnConnectionRequest) SetClientGatewayId(v string)`

SetClientGatewayId sets ClientGatewayId field to given value.

### HasClientGatewayId

`func (o *UpdateVpnConnectionRequest) HasClientGatewayId() bool`

HasClientGatewayId returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateVpnConnectionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateVpnConnectionRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateVpnConnectionRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateVpnConnectionRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVirtualGatewayId

`func (o *UpdateVpnConnectionRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *UpdateVpnConnectionRequest) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *UpdateVpnConnectionRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.

### HasVirtualGatewayId

`func (o *UpdateVpnConnectionRequest) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### GetVpnConnectionId

`func (o *UpdateVpnConnectionRequest) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *UpdateVpnConnectionRequest) GetVpnConnectionIdOk() (*string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionId

`func (o *UpdateVpnConnectionRequest) SetVpnConnectionId(v string)`

SetVpnConnectionId sets VpnConnectionId field to given value.


### GetVpnOptions

`func (o *UpdateVpnConnectionRequest) GetVpnOptions() VpnOptionsToUpdate`

GetVpnOptions returns the VpnOptions field if non-nil, zero value otherwise.

### GetVpnOptionsOk

`func (o *UpdateVpnConnectionRequest) GetVpnOptionsOk() (*VpnOptionsToUpdate, bool)`

GetVpnOptionsOk returns a tuple with the VpnOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnOptions

`func (o *UpdateVpnConnectionRequest) SetVpnOptions(v VpnOptionsToUpdate)`

SetVpnOptions sets VpnOptions field to given value.

### HasVpnOptions

`func (o *UpdateVpnConnectionRequest) HasVpnOptions() bool`

HasVpnOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


