# ReadConsumptionAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**FromDate** | **string** | The beginning of the time period, in ISO 8601 date format (for example, &#x60;2020-06-14&#x60;). The date-time format is also accepted, but only with a time set to midnight (for example, &#x60;2020-06-14T00:00:00.000Z&#x60;). | 
**Overall** | **bool** | By default or if false, returns only the consumption of the specific account that sends this request. If true, returns either the overall consumption of your paying account and all linked accounts (if the account that sends this request is a paying account) or returns nothing (if the account that sends this request is a linked account). | [optional] [default to false]
**ToDate** | **string** | The end of the time period, in ISO 8601 date format (for example, &#x60;2020-06-30&#x60;). The date-time format is also accepted, but only with a time set to midnight (for example, &#x60;2020-06-30T00:00:00.000Z&#x60;). | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


