# VpnOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase1Options** | Pointer to [**Phase1Options**](Phase1Options.md) |  | [optional] 
**Phase2Options** | Pointer to [**Phase2Options**](Phase2Options.md) |  | [optional] 
**TunnelInsideIpRange** | Pointer to **string** | The range of inside IP addresses for the tunnel. This must be a /30 CIDR block from the 169.254.254.0/24 range. | [optional] 

## Methods

### NewVpnOptions

`func NewVpnOptions() *VpnOptions`

NewVpnOptions instantiates a new VpnOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnOptionsWithDefaults

`func NewVpnOptionsWithDefaults() *VpnOptions`

NewVpnOptionsWithDefaults instantiates a new VpnOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase1Options

`func (o *VpnOptions) GetPhase1Options() Phase1Options`

GetPhase1Options returns the Phase1Options field if non-nil, zero value otherwise.

### GetPhase1OptionsOk

`func (o *VpnOptions) GetPhase1OptionsOk() (*Phase1Options, bool)`

GetPhase1OptionsOk returns a tuple with the Phase1Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Options

`func (o *VpnOptions) SetPhase1Options(v Phase1Options)`

SetPhase1Options sets Phase1Options field to given value.

### HasPhase1Options

`func (o *VpnOptions) HasPhase1Options() bool`

HasPhase1Options returns a boolean if a field has been set.

### GetPhase2Options

`func (o *VpnOptions) GetPhase2Options() Phase2Options`

GetPhase2Options returns the Phase2Options field if non-nil, zero value otherwise.

### GetPhase2OptionsOk

`func (o *VpnOptions) GetPhase2OptionsOk() (*Phase2Options, bool)`

GetPhase2OptionsOk returns a tuple with the Phase2Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Options

`func (o *VpnOptions) SetPhase2Options(v Phase2Options)`

SetPhase2Options sets Phase2Options field to given value.

### HasPhase2Options

`func (o *VpnOptions) HasPhase2Options() bool`

HasPhase2Options returns a boolean if a field has been set.

### GetTunnelInsideIpRange

`func (o *VpnOptions) GetTunnelInsideIpRange() string`

GetTunnelInsideIpRange returns the TunnelInsideIpRange field if non-nil, zero value otherwise.

### GetTunnelInsideIpRangeOk

`func (o *VpnOptions) GetTunnelInsideIpRangeOk() (*string, bool)`

GetTunnelInsideIpRangeOk returns a tuple with the TunnelInsideIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInsideIpRange

`func (o *VpnOptions) SetTunnelInsideIpRange(v string)`

SetTunnelInsideIpRange sets TunnelInsideIpRange field to given value.

### HasTunnelInsideIpRange

`func (o *VpnOptions) HasTunnelInsideIpRange() bool`

HasTunnelInsideIpRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


