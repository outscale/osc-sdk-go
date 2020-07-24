# ConsumptionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The category of the resource (for example, &#x60;network&#x60;). | [optional] 
**FromDate** | **string** | The beginning of the time period. | [optional] 
**Operation** | **string** | The API call that triggered the resource consumption (for example, &#x60;RunInstances&#x60; or &#x60;CreateVolume&#x60;). | [optional] 
**Service** | **string** | The service of the API call (&#x60;TinaOS-FCU&#x60;, &#x60;TinaOS-LBU&#x60;, &#x60;TinaOS-OSU&#x60; or &#x60;TinaOS-DirectLink&#x60;). | [optional] 
**Title** | **string** | A description of the consumed resource. | [optional] 
**ToDate** | **string** | The end of the time period. | [optional] 
**Type** | **string** | The type of resource, depending on the API call. | [optional] 
**Value** | **string** | The consumed amount for the resource. The unit depends on the resource type. For more information, see the &#x60;Title&#x60; element. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


