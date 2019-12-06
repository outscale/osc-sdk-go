# FiltersTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | **[]string** | The keys of the tags that are assigned to the resources. You can use this filter alongside the &#x60;Values&#x60; filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter. | [optional] 
**ResourceIds** | **[]string** | The IDs of the resources with which the tags are associated. | [optional] 
**ResourceTypes** | **[]string** | The resource type (&#x60;instance&#x60; \\| &#x60;image&#x60; \\| &#x60;volume&#x60; \\| &#x60;snapshot&#x60; \\| &#x60;public-ip&#x60; \\| &#x60;security-group&#x60; \\| &#x60;route-table&#x60; \\| &#x60;network-interface&#x60; \\| &#x60;vpc&#x60; \\| &#x60;subnet&#x60; \\| &#x60;network-link&#x60; \\| &#x60;vpc-endpoint&#x60; \\| &#x60;nat-gateway&#x60; \\| &#x60;internet-gateway&#x60; \\| &#x60;customer-gateway&#x60; \\| &#x60;vpn-gateway&#x60; \\| &#x60;vpn-connection&#x60; \\| &#x60;dhcp-options&#x60; \\| &#x60;task&#x60;). | [optional] 
**Values** | **[]string** | The values of the tags that are assigned to the resources. You can use this filter alongside the &#x60;TagKeys&#x60; filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


