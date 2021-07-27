# CreateLoadBalancerPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CookieExpirationPeriod** | **int32** | The lifetime of the cookie, in seconds. If not specified, the default value of this parameter is 1, which means that the sticky session lasts for the duration of the browser session. | [optional] 
**CookieName** | **string** | The name of the application cookie used for stickiness. This parameter is required if you create a stickiness policy based on an application-generated cookie. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer for which you want to create a policy. | 
**PolicyName** | **string** | The name of the policy. This name must be unique and consist of alphanumeric characters and dashes (-). | 
**PolicyType** | **string** | The type of stickiness policy you want to create: &#x60;app&#x60; or &#x60;load_balancer&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


