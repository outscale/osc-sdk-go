# CreateVpnConnectionRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpRange** | **string** | The network prefix of the route, in CIDR notation (for example, 10.12.0.0/16). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**VpnConnectionId** | **string** | The ID of the target VPN connection of the static route. | 

## Methods

### NewCreateVpnConnectionRouteRequest

`func NewCreateVpnConnectionRouteRequest(destinationIpRange string, vpnConnectionId string, ) *CreateVpnConnectionRouteRequest`

NewCreateVpnConnectionRouteRequest instantiates a new CreateVpnConnectionRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVpnConnectionRouteRequestWithDefaults

`func NewCreateVpnConnectionRouteRequestWithDefaults() *CreateVpnConnectionRouteRequest`

NewCreateVpnConnectionRouteRequestWithDefaults instantiates a new CreateVpnConnectionRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIpRange

`func (o *CreateVpnConnectionRouteRequest) GetDestinationIpRange() string`

GetDestinationIpRange returns the DestinationIpRange field if non-nil, zero value otherwise.

### GetDestinationIpRangeOk

`func (o *CreateVpnConnectionRouteRequest) GetDestinationIpRangeOk() (*string, bool)`

GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpRange

`func (o *CreateVpnConnectionRouteRequest) SetDestinationIpRange(v string)`

SetDestinationIpRange sets DestinationIpRange field to given value.


### GetDryRun

`func (o *CreateVpnConnectionRouteRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVpnConnectionRouteRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateVpnConnectionRouteRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateVpnConnectionRouteRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVpnConnectionId

`func (o *CreateVpnConnectionRouteRequest) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *CreateVpnConnectionRouteRequest) GetVpnConnectionIdOk() (*string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionId

`func (o *CreateVpnConnectionRouteRequest) SetVpnConnectionId(v string)`

SetVpnConnectionId sets VpnConnectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


