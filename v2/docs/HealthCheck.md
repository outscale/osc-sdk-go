# HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckInterval** | **int32** | The number of seconds between two pings (between &#x60;5&#x60; and &#x60;600&#x60; both included). | 
**HealthyThreshold** | **int32** | The number of consecutive successful pings before considering the VM as healthy (between &#x60;2&#x60; and &#x60;10&#x60; both included). | 
**Path** | Pointer to **string** | The path for HTTP or HTTPS requests. | [optional] 
**Port** | **int32** | The port number (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**Protocol** | **string** | The protocol for the URL of the VM (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | 
**Timeout** | **int32** | The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between &#x60;2&#x60; and &#x60;60&#x60; both included). | 
**UnhealthyThreshold** | **int32** | The number of consecutive failed pings before considering the VM as unhealthy (between &#x60;2&#x60; and &#x60;10&#x60; both included). | 

## Methods

### NewHealthCheck

`func NewHealthCheck(checkInterval int32, healthyThreshold int32, port int32, protocol string, timeout int32, unhealthyThreshold int32, ) *HealthCheck`

NewHealthCheck instantiates a new HealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckWithDefaults

`func NewHealthCheckWithDefaults() *HealthCheck`

NewHealthCheckWithDefaults instantiates a new HealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckInterval

`func (o *HealthCheck) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *HealthCheck) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *HealthCheck) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.


### GetHealthyThreshold

`func (o *HealthCheck) GetHealthyThreshold() int32`

GetHealthyThreshold returns the HealthyThreshold field if non-nil, zero value otherwise.

### GetHealthyThresholdOk

`func (o *HealthCheck) GetHealthyThresholdOk() (*int32, bool)`

GetHealthyThresholdOk returns a tuple with the HealthyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyThreshold

`func (o *HealthCheck) SetHealthyThreshold(v int32)`

SetHealthyThreshold sets HealthyThreshold field to given value.


### GetPath

`func (o *HealthCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HealthCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HealthCheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HealthCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *HealthCheck) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HealthCheck) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HealthCheck) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *HealthCheck) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HealthCheck) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HealthCheck) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetTimeout

`func (o *HealthCheck) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HealthCheck) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HealthCheck) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetUnhealthyThreshold

`func (o *HealthCheck) GetUnhealthyThreshold() int32`

GetUnhealthyThreshold returns the UnhealthyThreshold field if non-nil, zero value otherwise.

### GetUnhealthyThresholdOk

`func (o *HealthCheck) GetUnhealthyThresholdOk() (*int32, bool)`

GetUnhealthyThresholdOk returns a tuple with the UnhealthyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyThreshold

`func (o *HealthCheck) SetUnhealthyThreshold(v int32)`

SetUnhealthyThreshold sets UnhealthyThreshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


