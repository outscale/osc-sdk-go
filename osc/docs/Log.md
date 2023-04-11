# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID of the logged call. | [optional] 
**CallDuration** | **int32** | The duration of the logged call, in microseconds. | [optional] 
**QueryAccessKey** | **string** | The access key used for the logged call. | [optional] 
**QueryApiName** | **string** | The name of the API used by the logged call (always &#x60;oapi&#x60; for the OUTSCALE API). | [optional] 
**QueryApiVersion** | **string** | The version of the API used by the logged call. | [optional] 
**QueryCallName** | **string** | The name of the logged call. | [optional] 
**QueryDate** | [**time.Time**](time.Time.md) | The date and time of the logged call, in ISO 8601 date-time format. | [optional] 
**QueryHeaderRaw** | **string** | The raw header of the HTTP request of the logged call. | [optional] 
**QueryHeaderSize** | **int32** | The size of the raw header of the HTTP request of the logged call, in bytes. | [optional] 
**QueryIpAddress** | **string** | The IP used for the logged call. | [optional] 
**QueryPayloadRaw** | **string** | The raw payload of the HTTP request of the logged call. | [optional] 
**QueryPayloadSize** | **int32** | The size of the raw payload of the HTTP request of the logged call, in bytes. | [optional] 
**QueryUserAgent** | **string** | The user agent of the HTTP request of the logged call. | [optional] 
**RequestId** | **string** | The request ID provided in the response of the logged call. | [optional] 
**ResponseSize** | **int32** | The size of the response of the logged call, in bytes. | [optional] 
**ResponseStatusCode** | **int32** | The HTTP status code of the response of the logged call. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


