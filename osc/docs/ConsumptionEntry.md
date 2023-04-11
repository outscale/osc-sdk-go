# ConsumptionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of your TINA account. | [optional] 
**Category** | **string** | The category of the resource (for example, &#x60;network&#x60;). | [optional] 
**FromDate** | **string** | The beginning of the time period, in ISO 8601 date-time format. | [optional] 
**Operation** | **string** | The API call that triggered the resource consumption (for example, &#x60;RunInstances&#x60; or &#x60;CreateVolume&#x60;). | [optional] 
**PayingAccountId** | **string** | The ID of the TINA account which is billed for your consumption. It can be different from your account in the &#x60;AccountId&#x60; parameter. | [optional] 
**Service** | **string** | The service of the API call (&#x60;TinaOS-FCU&#x60;, &#x60;TinaOS-LBU&#x60;, &#x60;TinaOS-DirectLink&#x60;, &#x60;TinaOS-OOS&#x60;, or &#x60;TinaOS-OSU&#x60;). | [optional] 
**SubregionName** | **string** | The name of the Subregion. | [optional] 
**Title** | **string** | A description of the consumed resource. | [optional] 
**ToDate** | **string** | The end of the time period, in ISO 8601 date-time format. | [optional] 
**Type** | **string** | The type of resource, depending on the API call. | [optional] 
**Value** | **float64** | The consumed amount for the resource. The unit depends on the resource type. For more information, see the &#x60;Title&#x60; element. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


