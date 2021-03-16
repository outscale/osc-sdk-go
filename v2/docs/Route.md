# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationMethod** | Pointer to **string** | The method used to create the route. | [optional] 
**DestinationIpRange** | Pointer to **string** | The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24). | [optional] 
**DestinationServiceId** | Pointer to **string** | The ID of the 3DS OUTSCALE service. | [optional] 
**GatewayId** | Pointer to **string** | The ID of the Internet service or virtual gateway attached to the Net. | [optional] 
**NatServiceId** | Pointer to **string** | The ID of a NAT service attached to the Net. | [optional] 
**NetAccessPointId** | Pointer to **string** | The ID of the Net access point. | [optional] 
**NetPeeringId** | Pointer to **string** | The ID of the Net peering connection. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC. | [optional] 
**State** | Pointer to **string** | The state of a route in the route table (&#x60;active&#x60; \\| &#x60;blackhole&#x60;). The &#x60;blackhole&#x60; state indicates that the target of the route is not available. | [optional] 
**VmAccountId** | Pointer to **string** | The account ID of the owner of the VM. | [optional] 
**VmId** | Pointer to **string** | The ID of a VM specified in a route in the table. | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationMethod

`func (o *Route) GetCreationMethod() string`

GetCreationMethod returns the CreationMethod field if non-nil, zero value otherwise.

### GetCreationMethodOk

`func (o *Route) GetCreationMethodOk() (*string, bool)`

GetCreationMethodOk returns a tuple with the CreationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMethod

`func (o *Route) SetCreationMethod(v string)`

SetCreationMethod sets CreationMethod field to given value.

### HasCreationMethod

`func (o *Route) HasCreationMethod() bool`

HasCreationMethod returns a boolean if a field has been set.

### GetDestinationIpRange

`func (o *Route) GetDestinationIpRange() string`

GetDestinationIpRange returns the DestinationIpRange field if non-nil, zero value otherwise.

### GetDestinationIpRangeOk

`func (o *Route) GetDestinationIpRangeOk() (*string, bool)`

GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpRange

`func (o *Route) SetDestinationIpRange(v string)`

SetDestinationIpRange sets DestinationIpRange field to given value.

### HasDestinationIpRange

`func (o *Route) HasDestinationIpRange() bool`

HasDestinationIpRange returns a boolean if a field has been set.

### GetDestinationServiceId

`func (o *Route) GetDestinationServiceId() string`

GetDestinationServiceId returns the DestinationServiceId field if non-nil, zero value otherwise.

### GetDestinationServiceIdOk

`func (o *Route) GetDestinationServiceIdOk() (*string, bool)`

GetDestinationServiceIdOk returns a tuple with the DestinationServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationServiceId

`func (o *Route) SetDestinationServiceId(v string)`

SetDestinationServiceId sets DestinationServiceId field to given value.

### HasDestinationServiceId

`func (o *Route) HasDestinationServiceId() bool`

HasDestinationServiceId returns a boolean if a field has been set.

### GetGatewayId

`func (o *Route) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *Route) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *Route) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *Route) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetNatServiceId

`func (o *Route) GetNatServiceId() string`

GetNatServiceId returns the NatServiceId field if non-nil, zero value otherwise.

### GetNatServiceIdOk

`func (o *Route) GetNatServiceIdOk() (*string, bool)`

GetNatServiceIdOk returns a tuple with the NatServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatServiceId

`func (o *Route) SetNatServiceId(v string)`

SetNatServiceId sets NatServiceId field to given value.

### HasNatServiceId

`func (o *Route) HasNatServiceId() bool`

HasNatServiceId returns a boolean if a field has been set.

### GetNetAccessPointId

`func (o *Route) GetNetAccessPointId() string`

GetNetAccessPointId returns the NetAccessPointId field if non-nil, zero value otherwise.

### GetNetAccessPointIdOk

`func (o *Route) GetNetAccessPointIdOk() (*string, bool)`

GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAccessPointId

`func (o *Route) SetNetAccessPointId(v string)`

SetNetAccessPointId sets NetAccessPointId field to given value.

### HasNetAccessPointId

`func (o *Route) HasNetAccessPointId() bool`

HasNetAccessPointId returns a boolean if a field has been set.

### GetNetPeeringId

`func (o *Route) GetNetPeeringId() string`

GetNetPeeringId returns the NetPeeringId field if non-nil, zero value otherwise.

### GetNetPeeringIdOk

`func (o *Route) GetNetPeeringIdOk() (*string, bool)`

GetNetPeeringIdOk returns a tuple with the NetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPeeringId

`func (o *Route) SetNetPeeringId(v string)`

SetNetPeeringId sets NetPeeringId field to given value.

### HasNetPeeringId

`func (o *Route) HasNetPeeringId() bool`

HasNetPeeringId returns a boolean if a field has been set.

### GetNicId

`func (o *Route) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *Route) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *Route) SetNicId(v string)`

SetNicId sets NicId field to given value.

### HasNicId

`func (o *Route) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### GetState

`func (o *Route) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Route) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Route) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Route) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVmAccountId

`func (o *Route) GetVmAccountId() string`

GetVmAccountId returns the VmAccountId field if non-nil, zero value otherwise.

### GetVmAccountIdOk

`func (o *Route) GetVmAccountIdOk() (*string, bool)`

GetVmAccountIdOk returns a tuple with the VmAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmAccountId

`func (o *Route) SetVmAccountId(v string)`

SetVmAccountId sets VmAccountId field to given value.

### HasVmAccountId

`func (o *Route) HasVmAccountId() bool`

HasVmAccountId returns a boolean if a field has been set.

### GetVmId

`func (o *Route) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *Route) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *Route) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *Route) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


