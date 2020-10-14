# FiltersVpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsns** | Pointer to **[]int32** | The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections. | [optional] 
**ClientGatewayIds** | Pointer to **[]string** | The IDs of the client gateways. | [optional] 
**ConnectionTypes** | Pointer to **[]string** | The types of the VPN connections (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**RouteDestinationIpRanges** | Pointer to **[]string** | The destination IP ranges. | [optional] 
**States** | Pointer to **[]string** | The states of the VPN connections (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**StaticRoutesOnly** | Pointer to **bool** | If &#x60;false&#x60;, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If &#x60;true&#x60;, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the VPN connections. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the VPN connections. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the VPN connections, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 
**VirtualGatewayIds** | Pointer to **[]string** | The IDs of the virtual gateways. | [optional] 
**VpnConnectionIds** | Pointer to **[]string** | The IDs of the VPN connections. | [optional] 

## Methods

### NewFiltersVpnConnection

`func NewFiltersVpnConnection() *FiltersVpnConnection`

NewFiltersVpnConnection instantiates a new FiltersVpnConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersVpnConnectionWithDefaults

`func NewFiltersVpnConnectionWithDefaults() *FiltersVpnConnection`

NewFiltersVpnConnectionWithDefaults instantiates a new FiltersVpnConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpAsns

`func (o *FiltersVpnConnection) GetBgpAsns() []int32`

GetBgpAsns returns the BgpAsns field if non-nil, zero value otherwise.

### GetBgpAsnsOk

`func (o *FiltersVpnConnection) GetBgpAsnsOk() (*[]int32, bool)`

GetBgpAsnsOk returns a tuple with the BgpAsns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAsns

`func (o *FiltersVpnConnection) SetBgpAsns(v []int32)`

SetBgpAsns sets BgpAsns field to given value.

### HasBgpAsns

`func (o *FiltersVpnConnection) HasBgpAsns() bool`

HasBgpAsns returns a boolean if a field has been set.

### GetClientGatewayIds

`func (o *FiltersVpnConnection) GetClientGatewayIds() []string`

GetClientGatewayIds returns the ClientGatewayIds field if non-nil, zero value otherwise.

### GetClientGatewayIdsOk

`func (o *FiltersVpnConnection) GetClientGatewayIdsOk() (*[]string, bool)`

GetClientGatewayIdsOk returns a tuple with the ClientGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGatewayIds

`func (o *FiltersVpnConnection) SetClientGatewayIds(v []string)`

SetClientGatewayIds sets ClientGatewayIds field to given value.

### HasClientGatewayIds

`func (o *FiltersVpnConnection) HasClientGatewayIds() bool`

HasClientGatewayIds returns a boolean if a field has been set.

### GetConnectionTypes

`func (o *FiltersVpnConnection) GetConnectionTypes() []string`

GetConnectionTypes returns the ConnectionTypes field if non-nil, zero value otherwise.

### GetConnectionTypesOk

`func (o *FiltersVpnConnection) GetConnectionTypesOk() (*[]string, bool)`

GetConnectionTypesOk returns a tuple with the ConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTypes

`func (o *FiltersVpnConnection) SetConnectionTypes(v []string)`

SetConnectionTypes sets ConnectionTypes field to given value.

### HasConnectionTypes

`func (o *FiltersVpnConnection) HasConnectionTypes() bool`

HasConnectionTypes returns a boolean if a field has been set.

### GetRouteDestinationIpRanges

`func (o *FiltersVpnConnection) GetRouteDestinationIpRanges() []string`

GetRouteDestinationIpRanges returns the RouteDestinationIpRanges field if non-nil, zero value otherwise.

### GetRouteDestinationIpRangesOk

`func (o *FiltersVpnConnection) GetRouteDestinationIpRangesOk() (*[]string, bool)`

GetRouteDestinationIpRangesOk returns a tuple with the RouteDestinationIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDestinationIpRanges

`func (o *FiltersVpnConnection) SetRouteDestinationIpRanges(v []string)`

SetRouteDestinationIpRanges sets RouteDestinationIpRanges field to given value.

### HasRouteDestinationIpRanges

`func (o *FiltersVpnConnection) HasRouteDestinationIpRanges() bool`

HasRouteDestinationIpRanges returns a boolean if a field has been set.

### GetStates

`func (o *FiltersVpnConnection) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersVpnConnection) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersVpnConnection) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersVpnConnection) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetStaticRoutesOnly

`func (o *FiltersVpnConnection) GetStaticRoutesOnly() bool`

GetStaticRoutesOnly returns the StaticRoutesOnly field if non-nil, zero value otherwise.

### GetStaticRoutesOnlyOk

`func (o *FiltersVpnConnection) GetStaticRoutesOnlyOk() (*bool, bool)`

GetStaticRoutesOnlyOk returns a tuple with the StaticRoutesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutesOnly

`func (o *FiltersVpnConnection) SetStaticRoutesOnly(v bool)`

SetStaticRoutesOnly sets StaticRoutesOnly field to given value.

### HasStaticRoutesOnly

`func (o *FiltersVpnConnection) HasStaticRoutesOnly() bool`

HasStaticRoutesOnly returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersVpnConnection) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersVpnConnection) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersVpnConnection) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersVpnConnection) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersVpnConnection) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersVpnConnection) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersVpnConnection) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersVpnConnection) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersVpnConnection) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersVpnConnection) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersVpnConnection) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersVpnConnection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVirtualGatewayIds

`func (o *FiltersVpnConnection) GetVirtualGatewayIds() []string`

GetVirtualGatewayIds returns the VirtualGatewayIds field if non-nil, zero value otherwise.

### GetVirtualGatewayIdsOk

`func (o *FiltersVpnConnection) GetVirtualGatewayIdsOk() (*[]string, bool)`

GetVirtualGatewayIdsOk returns a tuple with the VirtualGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayIds

`func (o *FiltersVpnConnection) SetVirtualGatewayIds(v []string)`

SetVirtualGatewayIds sets VirtualGatewayIds field to given value.

### HasVirtualGatewayIds

`func (o *FiltersVpnConnection) HasVirtualGatewayIds() bool`

HasVirtualGatewayIds returns a boolean if a field has been set.

### GetVpnConnectionIds

`func (o *FiltersVpnConnection) GetVpnConnectionIds() []string`

GetVpnConnectionIds returns the VpnConnectionIds field if non-nil, zero value otherwise.

### GetVpnConnectionIdsOk

`func (o *FiltersVpnConnection) GetVpnConnectionIdsOk() (*[]string, bool)`

GetVpnConnectionIdsOk returns a tuple with the VpnConnectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionIds

`func (o *FiltersVpnConnection) SetVpnConnectionIds(v []string)`

SetVpnConnectionIds sets VpnConnectionIds field to given value.

### HasVpnConnectionIds

`func (o *FiltersVpnConnection) HasVpnConnectionIds() bool`

HasVpnConnectionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


