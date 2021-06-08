# UpdateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLog** | Pointer to [**AccessLog**](AccessLog.md) |  | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**HealthCheck** | Pointer to [**HealthCheck**](HealthCheck.md) |  | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer. | 
**LoadBalancerPort** | Pointer to **int32** | The port on which the load balancer is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). This parameter is required if you want to update the server certificate. | [optional] 
**PolicyNames** | Pointer to **[]string** | The name of the policy you want to enable for the listener. | [optional] 
**SecurityGroups** | Pointer to **[]string** | (Net only) One or more IDs of security groups you want to assign to the load balancer. You need to specify the already assigned security groups that you want to keep along with the new ones you are assigning. If the list is empty, the default security group of the Net is assigned to the load balancer. | [optional] 
**ServerCertificateId** | Pointer to **string** | The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers &gt; Outscale Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat). If this parameter is specified, you must also specify the &#x60;LoadBalancerPort&#x60; parameter. | [optional] 

## Methods

### NewUpdateLoadBalancerRequest

`func NewUpdateLoadBalancerRequest(loadBalancerName string, ) *UpdateLoadBalancerRequest`

NewUpdateLoadBalancerRequest instantiates a new UpdateLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerRequestWithDefaults

`func NewUpdateLoadBalancerRequestWithDefaults() *UpdateLoadBalancerRequest`

NewUpdateLoadBalancerRequestWithDefaults instantiates a new UpdateLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLog

`func (o *UpdateLoadBalancerRequest) GetAccessLog() AccessLog`

GetAccessLog returns the AccessLog field if non-nil, zero value otherwise.

### GetAccessLogOk

`func (o *UpdateLoadBalancerRequest) GetAccessLogOk() (*AccessLog, bool)`

GetAccessLogOk returns a tuple with the AccessLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLog

`func (o *UpdateLoadBalancerRequest) SetAccessLog(v AccessLog)`

SetAccessLog sets AccessLog field to given value.

### HasAccessLog

`func (o *UpdateLoadBalancerRequest) HasAccessLog() bool`

HasAccessLog returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateLoadBalancerRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateLoadBalancerRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateLoadBalancerRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateLoadBalancerRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetHealthCheck

`func (o *UpdateLoadBalancerRequest) GetHealthCheck() HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *UpdateLoadBalancerRequest) GetHealthCheckOk() (*HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *UpdateLoadBalancerRequest) SetHealthCheck(v HealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *UpdateLoadBalancerRequest) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *UpdateLoadBalancerRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.


### GetLoadBalancerPort

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerPort() int32`

GetLoadBalancerPort returns the LoadBalancerPort field if non-nil, zero value otherwise.

### GetLoadBalancerPortOk

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerPortOk() (*int32, bool)`

GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPort

`func (o *UpdateLoadBalancerRequest) SetLoadBalancerPort(v int32)`

SetLoadBalancerPort sets LoadBalancerPort field to given value.

### HasLoadBalancerPort

`func (o *UpdateLoadBalancerRequest) HasLoadBalancerPort() bool`

HasLoadBalancerPort returns a boolean if a field has been set.

### GetPolicyNames

`func (o *UpdateLoadBalancerRequest) GetPolicyNames() []string`

GetPolicyNames returns the PolicyNames field if non-nil, zero value otherwise.

### GetPolicyNamesOk

`func (o *UpdateLoadBalancerRequest) GetPolicyNamesOk() (*[]string, bool)`

GetPolicyNamesOk returns a tuple with the PolicyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNames

`func (o *UpdateLoadBalancerRequest) SetPolicyNames(v []string)`

SetPolicyNames sets PolicyNames field to given value.

### HasPolicyNames

`func (o *UpdateLoadBalancerRequest) HasPolicyNames() bool`

HasPolicyNames returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *UpdateLoadBalancerRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *UpdateLoadBalancerRequest) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *UpdateLoadBalancerRequest) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *UpdateLoadBalancerRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetServerCertificateId

`func (o *UpdateLoadBalancerRequest) GetServerCertificateId() string`

GetServerCertificateId returns the ServerCertificateId field if non-nil, zero value otherwise.

### GetServerCertificateIdOk

`func (o *UpdateLoadBalancerRequest) GetServerCertificateIdOk() (*string, bool)`

GetServerCertificateIdOk returns a tuple with the ServerCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificateId

`func (o *UpdateLoadBalancerRequest) SetServerCertificateId(v string)`

SetServerCertificateId sets ServerCertificateId field to given value.

### HasServerCertificateId

`func (o *UpdateLoadBalancerRequest) HasServerCertificateId() bool`

HasServerCertificateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


