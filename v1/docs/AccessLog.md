# AccessLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If &#x60;true&#x60;, access logs are enabled for your load balancer. If &#x60;false&#x60;, they are not. If you set this to &#x60;true&#x60; in your request, the &#x60;OsuBucketName&#x60; parameter is required. | [optional] 
**OsuBucketName** | **string** | The name of the Object Storage Unit (OSU) bucket for the access logs. | [optional] 
**OsuBucketPrefix** | **string** | The path to the folder of the access logs in your Object Storage Unit (OSU) bucket (by default, the &#x60;root&#x60; level of your bucket). | [optional] 
**PublicationInterval** | **int32** | The time interval for the publication of access logs in the Object Storage Unit (OSU) bucket, in minutes. This value can be either 5 or 60 (by default, 60). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


