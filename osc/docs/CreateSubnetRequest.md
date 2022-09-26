# CreateSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**IpRange** | **string** | The IP range in the Subnet, in CIDR notation (for example, &#x60;10.0.0.0/16&#x60;).&lt;br /&gt; The IP range of the Subnet can be either the same as the Net one if you create only a single Subnet in this Net, or a subset of the Net one. In case of several Subnets in a Net, their IP ranges must not overlap. The smallest Subnet you can create uses a /29 netmask (eight IPs). For more information, see [About VPCs](https://docs.outscale.com/en/userguide/About-VPCs.html). | 
**NetId** | **string** | The ID of the Net for which you want to create a Subnet. | 
**SubregionName** | **string** | The name of the Subregion in which you want to create the Subnet. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


