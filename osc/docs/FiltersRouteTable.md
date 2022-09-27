# FiltersRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkRouteTableIds** | **[]string** | The IDs of the route tables involved in the associations. | [optional] 
**LinkRouteTableLinkRouteTableIds** | **[]string** | The IDs of the associations between the route tables and the Subnets. | [optional] 
**LinkRouteTableMain** | **bool** | If true, the route tables are the main ones for their Nets. | [optional] 
**LinkSubnetIds** | **[]string** | The IDs of the Subnets involved in the associations. | [optional] 
**NetIds** | **[]string** | The IDs of the Nets for the route tables. | [optional] 
**RouteCreationMethods** | **[]string** | The methods used to create a route. | [optional] 
**RouteDestinationIpRanges** | **[]string** | The IP ranges specified in routes in the tables. | [optional] 
**RouteDestinationServiceIds** | **[]string** | The service IDs specified in routes in the tables. | [optional] 
**RouteGatewayIds** | **[]string** | The IDs of the gateways specified in routes in the tables. | [optional] 
**RouteNatServiceIds** | **[]string** | The IDs of the NAT services specified in routes in the tables. | [optional] 
**RouteNetPeeringIds** | **[]string** | The IDs of the Net peering connections specified in routes in the tables. | [optional] 
**RouteStates** | **[]string** | The states of routes in the route tables (always &#x60;active&#x60;). | [optional] 
**RouteTableIds** | **[]string** | The IDs of the route tables. | [optional] 
**RouteVmIds** | **[]string** | The IDs of the VMs specified in routes in the tables. | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the route tables. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the route tables. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the route tables, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


