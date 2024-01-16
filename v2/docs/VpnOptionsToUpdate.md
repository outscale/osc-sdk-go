# VpnOptionsToUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase2Options** | Pointer to [**Phase2OptionsToUpdate**](Phase2OptionsToUpdate.md) |  | [optional] 
**TunnelInsideIpRange** | Pointer to **string** | The range of inside IPs for the tunnel. This must be a /30 CIDR block from the 169.254.254.0/24 range. | [optional] 

## Methods

### NewVpnOptionsToUpdate

`func NewVpnOptionsToUpdate() *VpnOptionsToUpdate`

NewVpnOptionsToUpdate instantiates a new VpnOptionsToUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnOptionsToUpdateWithDefaults

`func NewVpnOptionsToUpdateWithDefaults() *VpnOptionsToUpdate`

NewVpnOptionsToUpdateWithDefaults instantiates a new VpnOptionsToUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase2Options

`func (o *VpnOptionsToUpdate) GetPhase2Options() Phase2OptionsToUpdate`

GetPhase2Options returns the Phase2Options field if non-nil, zero value otherwise.

### GetPhase2OptionsOk

`func (o *VpnOptionsToUpdate) GetPhase2OptionsOk() (*Phase2OptionsToUpdate, bool)`

GetPhase2OptionsOk returns a tuple with the Phase2Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Options

`func (o *VpnOptionsToUpdate) SetPhase2Options(v Phase2OptionsToUpdate)`

SetPhase2Options sets Phase2Options field to given value.

### HasPhase2Options

`func (o *VpnOptionsToUpdate) HasPhase2Options() bool`

HasPhase2Options returns a boolean if a field has been set.

### GetTunnelInsideIpRange

`func (o *VpnOptionsToUpdate) GetTunnelInsideIpRange() string`

GetTunnelInsideIpRange returns the TunnelInsideIpRange field if non-nil, zero value otherwise.

### GetTunnelInsideIpRangeOk

`func (o *VpnOptionsToUpdate) GetTunnelInsideIpRangeOk() (*string, bool)`

GetTunnelInsideIpRangeOk returns a tuple with the TunnelInsideIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInsideIpRange

`func (o *VpnOptionsToUpdate) SetTunnelInsideIpRange(v string)`

SetTunnelInsideIpRange sets TunnelInsideIpRange field to given value.

### HasTunnelInsideIpRange

`func (o *VpnOptionsToUpdate) HasTunnelInsideIpRange() bool`

HasTunnelInsideIpRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


