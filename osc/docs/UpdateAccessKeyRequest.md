# UpdateAccessKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | The ID of the access key. | 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ExpirationDate** | [**OneOfDateTimedate**](oneOf&lt;DateTime,date&gt;.md) | The date and time, or the date, at which you want the access key to expire, in ISO 8601 format (for example, &#x60;2020-06-14T00:00:00.000Z&#x60; or &#x60;2020-06-14&#x60;). If not specified, the access key is set to not expire. | [optional] 
**State** | **string** | The new state for the access key (&#x60;ACTIVE&#x60; \\| &#x60;INACTIVE&#x60;). When set to &#x60;ACTIVE&#x60;, the access key is enabled and can be used to send requests. When set to &#x60;INACTIVE&#x60;, the access key is disabled. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


