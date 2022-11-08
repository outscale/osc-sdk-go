# FiltersRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkRouteTableIds** | Pointer to **[]string** | The IDs of the route tables involved in the associations. | [optional] 
**LinkRouteTableLinkRouteTableIds** | Pointer to **[]string** | The IDs of the associations between the route tables and the Subnets. | [optional] 
**LinkRouteTableMain** | Pointer to **bool** | If true, the route tables are the main ones for their Nets. | [optional] 
**LinkSubnetIds** | Pointer to **[]string** | The IDs of the Subnets involved in the associations. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets for the route tables. | [optional] 
**RouteCreationMethods** | Pointer to **[]string** | The methods used to create a route. | [optional] 
**RouteDestinationIpRanges** | Pointer to **[]string** | The IP ranges specified in routes in the tables. | [optional] 
**RouteDestinationServiceIds** | Pointer to **[]string** | The service IDs specified in routes in the tables. | [optional] 
**RouteGatewayIds** | Pointer to **[]string** | The IDs of the gateways specified in routes in the tables. | [optional] 
**RouteNatServiceIds** | Pointer to **[]string** | The IDs of the NAT services specified in routes in the tables. | [optional] 
**RouteNetPeeringIds** | Pointer to **[]string** | The IDs of the Net peerings specified in routes in the tables. | [optional] 
**RouteStates** | Pointer to **[]string** | The states of routes in the route tables (always &#x60;active&#x60;). | [optional] 
**RouteTableIds** | Pointer to **[]string** | The IDs of the route tables. | [optional] 
**RouteVmIds** | Pointer to **[]string** | The IDs of the VMs specified in routes in the tables. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the route tables. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the route tables. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the route tables, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

## Methods

### NewFiltersRouteTable

`func NewFiltersRouteTable() *FiltersRouteTable`

NewFiltersRouteTable instantiates a new FiltersRouteTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersRouteTableWithDefaults

`func NewFiltersRouteTableWithDefaults() *FiltersRouteTable`

NewFiltersRouteTableWithDefaults instantiates a new FiltersRouteTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkRouteTableIds

`func (o *FiltersRouteTable) GetLinkRouteTableIds() []string`

GetLinkRouteTableIds returns the LinkRouteTableIds field if non-nil, zero value otherwise.

### GetLinkRouteTableIdsOk

`func (o *FiltersRouteTable) GetLinkRouteTableIdsOk() (*[]string, bool)`

GetLinkRouteTableIdsOk returns a tuple with the LinkRouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRouteTableIds

`func (o *FiltersRouteTable) SetLinkRouteTableIds(v []string)`

SetLinkRouteTableIds sets LinkRouteTableIds field to given value.

### HasLinkRouteTableIds

`func (o *FiltersRouteTable) HasLinkRouteTableIds() bool`

HasLinkRouteTableIds returns a boolean if a field has been set.

### GetLinkRouteTableLinkRouteTableIds

`func (o *FiltersRouteTable) GetLinkRouteTableLinkRouteTableIds() []string`

GetLinkRouteTableLinkRouteTableIds returns the LinkRouteTableLinkRouteTableIds field if non-nil, zero value otherwise.

### GetLinkRouteTableLinkRouteTableIdsOk

`func (o *FiltersRouteTable) GetLinkRouteTableLinkRouteTableIdsOk() (*[]string, bool)`

GetLinkRouteTableLinkRouteTableIdsOk returns a tuple with the LinkRouteTableLinkRouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRouteTableLinkRouteTableIds

`func (o *FiltersRouteTable) SetLinkRouteTableLinkRouteTableIds(v []string)`

SetLinkRouteTableLinkRouteTableIds sets LinkRouteTableLinkRouteTableIds field to given value.

### HasLinkRouteTableLinkRouteTableIds

`func (o *FiltersRouteTable) HasLinkRouteTableLinkRouteTableIds() bool`

HasLinkRouteTableLinkRouteTableIds returns a boolean if a field has been set.

### GetLinkRouteTableMain

`func (o *FiltersRouteTable) GetLinkRouteTableMain() bool`

GetLinkRouteTableMain returns the LinkRouteTableMain field if non-nil, zero value otherwise.

### GetLinkRouteTableMainOk

`func (o *FiltersRouteTable) GetLinkRouteTableMainOk() (*bool, bool)`

GetLinkRouteTableMainOk returns a tuple with the LinkRouteTableMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRouteTableMain

`func (o *FiltersRouteTable) SetLinkRouteTableMain(v bool)`

SetLinkRouteTableMain sets LinkRouteTableMain field to given value.

### HasLinkRouteTableMain

`func (o *FiltersRouteTable) HasLinkRouteTableMain() bool`

HasLinkRouteTableMain returns a boolean if a field has been set.

### GetLinkSubnetIds

`func (o *FiltersRouteTable) GetLinkSubnetIds() []string`

GetLinkSubnetIds returns the LinkSubnetIds field if non-nil, zero value otherwise.

### GetLinkSubnetIdsOk

`func (o *FiltersRouteTable) GetLinkSubnetIdsOk() (*[]string, bool)`

GetLinkSubnetIdsOk returns a tuple with the LinkSubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSubnetIds

`func (o *FiltersRouteTable) SetLinkSubnetIds(v []string)`

SetLinkSubnetIds sets LinkSubnetIds field to given value.

### HasLinkSubnetIds

`func (o *FiltersRouteTable) HasLinkSubnetIds() bool`

HasLinkSubnetIds returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersRouteTable) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersRouteTable) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersRouteTable) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersRouteTable) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetRouteCreationMethods

`func (o *FiltersRouteTable) GetRouteCreationMethods() []string`

GetRouteCreationMethods returns the RouteCreationMethods field if non-nil, zero value otherwise.

### GetRouteCreationMethodsOk

`func (o *FiltersRouteTable) GetRouteCreationMethodsOk() (*[]string, bool)`

GetRouteCreationMethodsOk returns a tuple with the RouteCreationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteCreationMethods

`func (o *FiltersRouteTable) SetRouteCreationMethods(v []string)`

SetRouteCreationMethods sets RouteCreationMethods field to given value.

### HasRouteCreationMethods

`func (o *FiltersRouteTable) HasRouteCreationMethods() bool`

HasRouteCreationMethods returns a boolean if a field has been set.

### GetRouteDestinationIpRanges

`func (o *FiltersRouteTable) GetRouteDestinationIpRanges() []string`

GetRouteDestinationIpRanges returns the RouteDestinationIpRanges field if non-nil, zero value otherwise.

### GetRouteDestinationIpRangesOk

`func (o *FiltersRouteTable) GetRouteDestinationIpRangesOk() (*[]string, bool)`

GetRouteDestinationIpRangesOk returns a tuple with the RouteDestinationIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDestinationIpRanges

`func (o *FiltersRouteTable) SetRouteDestinationIpRanges(v []string)`

SetRouteDestinationIpRanges sets RouteDestinationIpRanges field to given value.

### HasRouteDestinationIpRanges

`func (o *FiltersRouteTable) HasRouteDestinationIpRanges() bool`

HasRouteDestinationIpRanges returns a boolean if a field has been set.

### GetRouteDestinationServiceIds

`func (o *FiltersRouteTable) GetRouteDestinationServiceIds() []string`

GetRouteDestinationServiceIds returns the RouteDestinationServiceIds field if non-nil, zero value otherwise.

### GetRouteDestinationServiceIdsOk

`func (o *FiltersRouteTable) GetRouteDestinationServiceIdsOk() (*[]string, bool)`

GetRouteDestinationServiceIdsOk returns a tuple with the RouteDestinationServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDestinationServiceIds

`func (o *FiltersRouteTable) SetRouteDestinationServiceIds(v []string)`

SetRouteDestinationServiceIds sets RouteDestinationServiceIds field to given value.

### HasRouteDestinationServiceIds

`func (o *FiltersRouteTable) HasRouteDestinationServiceIds() bool`

HasRouteDestinationServiceIds returns a boolean if a field has been set.

### GetRouteGatewayIds

`func (o *FiltersRouteTable) GetRouteGatewayIds() []string`

GetRouteGatewayIds returns the RouteGatewayIds field if non-nil, zero value otherwise.

### GetRouteGatewayIdsOk

`func (o *FiltersRouteTable) GetRouteGatewayIdsOk() (*[]string, bool)`

GetRouteGatewayIdsOk returns a tuple with the RouteGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteGatewayIds

`func (o *FiltersRouteTable) SetRouteGatewayIds(v []string)`

SetRouteGatewayIds sets RouteGatewayIds field to given value.

### HasRouteGatewayIds

`func (o *FiltersRouteTable) HasRouteGatewayIds() bool`

HasRouteGatewayIds returns a boolean if a field has been set.

### GetRouteNatServiceIds

`func (o *FiltersRouteTable) GetRouteNatServiceIds() []string`

GetRouteNatServiceIds returns the RouteNatServiceIds field if non-nil, zero value otherwise.

### GetRouteNatServiceIdsOk

`func (o *FiltersRouteTable) GetRouteNatServiceIdsOk() (*[]string, bool)`

GetRouteNatServiceIdsOk returns a tuple with the RouteNatServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteNatServiceIds

`func (o *FiltersRouteTable) SetRouteNatServiceIds(v []string)`

SetRouteNatServiceIds sets RouteNatServiceIds field to given value.

### HasRouteNatServiceIds

`func (o *FiltersRouteTable) HasRouteNatServiceIds() bool`

HasRouteNatServiceIds returns a boolean if a field has been set.

### GetRouteNetPeeringIds

`func (o *FiltersRouteTable) GetRouteNetPeeringIds() []string`

GetRouteNetPeeringIds returns the RouteNetPeeringIds field if non-nil, zero value otherwise.

### GetRouteNetPeeringIdsOk

`func (o *FiltersRouteTable) GetRouteNetPeeringIdsOk() (*[]string, bool)`

GetRouteNetPeeringIdsOk returns a tuple with the RouteNetPeeringIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteNetPeeringIds

`func (o *FiltersRouteTable) SetRouteNetPeeringIds(v []string)`

SetRouteNetPeeringIds sets RouteNetPeeringIds field to given value.

### HasRouteNetPeeringIds

`func (o *FiltersRouteTable) HasRouteNetPeeringIds() bool`

HasRouteNetPeeringIds returns a boolean if a field has been set.

### GetRouteStates

`func (o *FiltersRouteTable) GetRouteStates() []string`

GetRouteStates returns the RouteStates field if non-nil, zero value otherwise.

### GetRouteStatesOk

`func (o *FiltersRouteTable) GetRouteStatesOk() (*[]string, bool)`

GetRouteStatesOk returns a tuple with the RouteStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteStates

`func (o *FiltersRouteTable) SetRouteStates(v []string)`

SetRouteStates sets RouteStates field to given value.

### HasRouteStates

`func (o *FiltersRouteTable) HasRouteStates() bool`

HasRouteStates returns a boolean if a field has been set.

### GetRouteTableIds

`func (o *FiltersRouteTable) GetRouteTableIds() []string`

GetRouteTableIds returns the RouteTableIds field if non-nil, zero value otherwise.

### GetRouteTableIdsOk

`func (o *FiltersRouteTable) GetRouteTableIdsOk() (*[]string, bool)`

GetRouteTableIdsOk returns a tuple with the RouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableIds

`func (o *FiltersRouteTable) SetRouteTableIds(v []string)`

SetRouteTableIds sets RouteTableIds field to given value.

### HasRouteTableIds

`func (o *FiltersRouteTable) HasRouteTableIds() bool`

HasRouteTableIds returns a boolean if a field has been set.

### GetRouteVmIds

`func (o *FiltersRouteTable) GetRouteVmIds() []string`

GetRouteVmIds returns the RouteVmIds field if non-nil, zero value otherwise.

### GetRouteVmIdsOk

`func (o *FiltersRouteTable) GetRouteVmIdsOk() (*[]string, bool)`

GetRouteVmIdsOk returns a tuple with the RouteVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteVmIds

`func (o *FiltersRouteTable) SetRouteVmIds(v []string)`

SetRouteVmIds sets RouteVmIds field to given value.

### HasRouteVmIds

`func (o *FiltersRouteTable) HasRouteVmIds() bool`

HasRouteVmIds returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersRouteTable) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersRouteTable) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersRouteTable) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersRouteTable) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersRouteTable) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersRouteTable) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersRouteTable) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersRouteTable) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersRouteTable) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersRouteTable) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersRouteTable) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersRouteTable) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


