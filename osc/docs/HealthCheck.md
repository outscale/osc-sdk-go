# HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckInterval** | **int32** | The number of seconds between two pings (between &#x60;5&#x60; and &#x60;600&#x60; both included). | 
**HealthyThreshold** | **int32** | The number of consecutive successful pings before considering the VM as healthy (between &#x60;2&#x60; and &#x60;10&#x60; both included). | 
**Path** | **string** | The path for HTTP or HTTPS requests. | 
**Port** | **int32** | The port number (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**Protocol** | **string** | The protocol for the URL of the VM (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | 
**Timeout** | **int32** | The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between &#x60;2&#x60; and &#x60;60&#x60; both included). | 
**UnhealthyThreshold** | **int32** | The number of consecutive failed pings before considering the VM as unhealthy (between &#x60;2&#x60; and &#x60;10&#x60; both included). | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


