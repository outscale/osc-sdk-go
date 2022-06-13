# AccessLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the &#x60;OsuBucketName&#x60; parameter is required. | [optional] 
**OsuBucketName** | **string** | The name of the OOS bucket for the access logs. | [optional] 
**OsuBucketPrefix** | **string** | The path to the folder of the access logs in your OOS bucket (by default, the &#x60;root&#x60; level of your bucket). | [optional] 
**PublicationInterval** | **int32** | The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either &#x60;5&#x60; or &#x60;60&#x60; (by default, &#x60;60&#x60;). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


