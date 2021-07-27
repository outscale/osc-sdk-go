# UpdateApiAccessPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**MaxAccessKeyExpirationSeconds** | **int64** | The maximum possible lifetime for your access keys, in seconds (between &#x60;0&#x60; and &#x60;3153600000&#x60;, both included). By default or if set to &#x60;O&#x60;, your access keys can have unlimited lifetimes. Otherwise, all your access keys must have an expiration date. This value must be greater than the remaining lifetime of each access key of your account. | 
**RequireTrustedEnv** | **bool** | If true, a trusted session is activated, provided that you specify the &#x60;MaxAccessKeyExpirationSeconds&#x60; parameter with a value greater than &#x60;0&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


