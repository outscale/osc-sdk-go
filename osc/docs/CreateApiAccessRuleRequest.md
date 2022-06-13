# CreateApiAccessRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaIds** | **[]string** |  One or more IDs of Client Certificate Authorities (CAs). | [optional] 
**Cns** | **[]string** | One or more Client Certificate Common Names (CNs). If this parameter is specified, you must also specify the &#x60;CaIds&#x60; parameter. | [optional] 
**Description** | **string** | A description for the API access rule. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**IpRanges** | **[]string** | One or more IP ranges, in CIDR notation (for example, &#x60;192.0.2.0/16&#x60;). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


