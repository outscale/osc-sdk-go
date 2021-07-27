# FiltersApiLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryAccessKeys** | **[]string** | The access keys used for the logged calls. | [optional] 
**QueryApiNames** | **[]string** | The names of the APIs of the logged calls (always &#x60;oapi&#x60; for the OUTSCALE API). | [optional] 
**QueryCallNames** | **[]string** | The names of the logged calls. | [optional] 
**QueryDateAfter** | **string** | The date after which you want to retrieve logged calls, in ISO 8601 format (for example, &#x60;2020-06-14&#x60;). By default, this date is set to 2 days ago. | [optional] 
**QueryDateBefore** | **string** | The date before which you want to retrieve logged calls, in ISO 8601 format (for example, &#x60;2020-06-30&#x60;). By default, this date is set to now. | [optional] 
**QueryIpAddresses** | **[]string** | The IP addresses used for the logged calls. | [optional] 
**QueryUserAgents** | **[]string** | The user agents of the HTTP requests of the logged calls. | [optional] 
**RequestIds** | **[]string** | The request IDs provided in the responses of the logged calls. | [optional] 
**ResponseStatusCodes** | **[]int32** | The HTTP status codes of the logged calls. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


