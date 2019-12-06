# FiltersRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkRouteTableLinkRouteTableIds** | **[]string** | The IDs of the associations between the route tables and the Subnets. | [optional] 
**LinkSubnetIds** | **[]string** | The IDs of the Subnets involved in the associations. | [optional] 
**NetIds** | **[]string** | The IDs of the Nets for the route tables. | [optional] 
**RouteCreationMethods** | **[]string** | The methods used to create a route. | [optional] 
**RouteDestinationIpRanges** | **[]string** | The IP ranges specified in routes in the tables. | [optional] 
**RouteGatewayIds** | **[]string** | The IDs of the gateways specified in routes in the tables. | [optional] 
**RouteNatServiceIds** | **[]string** | The IDs of the NAT services specified in routes in the tables. | [optional] 
**RouteNetPeeringIds** | **[]string** | The IDs of the Net peering connections specified in routes in the tables. | [optional] 
**RouteStates** | **[]string** | The states of routes in the route tables (&#x60;active&#x60; \\| &#x60;blackhole&#x60;). The &#x60;blackhole&#x60; state indicates that the target of the route is not available. | [optional] 
**RouteTableIds** | **[]string** | The IDs of the route tables. | [optional] 
**RouteVmIds** | **[]string** | The IDs of the VMs specified in routes in the tables. | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the route tables. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the route tables. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the route tables, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


