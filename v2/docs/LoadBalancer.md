# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLog** | Pointer to [**AccessLog**](AccessLog.md) |  | [optional] 
**ApplicationStickyCookiePolicies** | Pointer to [**[]ApplicationStickyCookiePolicy**](ApplicationStickyCookiePolicy.md) | The stickiness policies defined for the load balancer. | [optional] 
**BackendIps** | Pointer to **[]string** | One or more public IPs of back-end VMs. | [optional] 
**BackendVmIds** | Pointer to **[]string** | One or more IDs of back-end VMs for the load balancer. | [optional] 
**DnsName** | Pointer to **string** | The DNS name of the load balancer. | [optional] 
**HealthCheck** | Pointer to [**HealthCheck**](HealthCheck.md) |  | [optional] 
**Listeners** | Pointer to [**[]Listener**](Listener.md) | The listeners for the load balancer. | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer. | [optional] 
**LoadBalancerStickyCookiePolicies** | Pointer to [**[]LoadBalancerStickyCookiePolicy**](LoadBalancerStickyCookiePolicy.md) | The policies defined for the load balancer. | [optional] 
**LoadBalancerType** | Pointer to **string** | The type of load balancer. Valid only for load balancers in a Net.&lt;br /&gt; If &#x60;LoadBalancerType&#x60; is &#x60;internet-facing&#x60;, the load balancer has a public DNS name that resolves to a public IP.&lt;br /&gt; If &#x60;LoadBalancerType&#x60; is &#x60;internal&#x60;, the load balancer has a public DNS name that resolves to a private IP. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the load balancer. | [optional] 
**PublicIp** | Pointer to **string** | (internet-facing only) The public IP associated with the load balancer. | [optional] 
**SecuredCookies** | Pointer to **bool** | Whether secure cookies are enabled for the load balancer. | [optional] 
**SecurityGroups** | Pointer to **[]string** | One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net. | [optional] 
**SourceSecurityGroup** | Pointer to [**SourceSecurityGroup**](SourceSecurityGroup.md) |  | [optional] 
**Subnets** | Pointer to **[]string** | The ID of the Subnet in which the load balancer was created. | [optional] 
**SubregionNames** | Pointer to **[]string** | The ID of the Subregion in which the load balancer was created. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the load balancer. | [optional] 

## Methods

### NewLoadBalancer

`func NewLoadBalancer() *LoadBalancer`

NewLoadBalancer instantiates a new LoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerWithDefaults

`func NewLoadBalancerWithDefaults() *LoadBalancer`

NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLog

`func (o *LoadBalancer) GetAccessLog() AccessLog`

GetAccessLog returns the AccessLog field if non-nil, zero value otherwise.

### GetAccessLogOk

`func (o *LoadBalancer) GetAccessLogOk() (*AccessLog, bool)`

GetAccessLogOk returns a tuple with the AccessLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLog

`func (o *LoadBalancer) SetAccessLog(v AccessLog)`

SetAccessLog sets AccessLog field to given value.

### HasAccessLog

`func (o *LoadBalancer) HasAccessLog() bool`

HasAccessLog returns a boolean if a field has been set.

### GetApplicationStickyCookiePolicies

`func (o *LoadBalancer) GetApplicationStickyCookiePolicies() []ApplicationStickyCookiePolicy`

GetApplicationStickyCookiePolicies returns the ApplicationStickyCookiePolicies field if non-nil, zero value otherwise.

### GetApplicationStickyCookiePoliciesOk

`func (o *LoadBalancer) GetApplicationStickyCookiePoliciesOk() (*[]ApplicationStickyCookiePolicy, bool)`

GetApplicationStickyCookiePoliciesOk returns a tuple with the ApplicationStickyCookiePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationStickyCookiePolicies

`func (o *LoadBalancer) SetApplicationStickyCookiePolicies(v []ApplicationStickyCookiePolicy)`

SetApplicationStickyCookiePolicies sets ApplicationStickyCookiePolicies field to given value.

### HasApplicationStickyCookiePolicies

`func (o *LoadBalancer) HasApplicationStickyCookiePolicies() bool`

HasApplicationStickyCookiePolicies returns a boolean if a field has been set.

### GetBackendIps

`func (o *LoadBalancer) GetBackendIps() []string`

GetBackendIps returns the BackendIps field if non-nil, zero value otherwise.

### GetBackendIpsOk

`func (o *LoadBalancer) GetBackendIpsOk() (*[]string, bool)`

GetBackendIpsOk returns a tuple with the BackendIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendIps

`func (o *LoadBalancer) SetBackendIps(v []string)`

SetBackendIps sets BackendIps field to given value.

### HasBackendIps

`func (o *LoadBalancer) HasBackendIps() bool`

HasBackendIps returns a boolean if a field has been set.

### GetBackendVmIds

`func (o *LoadBalancer) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *LoadBalancer) GetBackendVmIdsOk() (*[]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVmIds

`func (o *LoadBalancer) SetBackendVmIds(v []string)`

SetBackendVmIds sets BackendVmIds field to given value.

### HasBackendVmIds

`func (o *LoadBalancer) HasBackendVmIds() bool`

HasBackendVmIds returns a boolean if a field has been set.

### GetDnsName

`func (o *LoadBalancer) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *LoadBalancer) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *LoadBalancer) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *LoadBalancer) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetHealthCheck

`func (o *LoadBalancer) GetHealthCheck() HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *LoadBalancer) GetHealthCheckOk() (*HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *LoadBalancer) SetHealthCheck(v HealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *LoadBalancer) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetListeners

`func (o *LoadBalancer) GetListeners() []Listener`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *LoadBalancer) GetListenersOk() (*[]Listener, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *LoadBalancer) SetListeners(v []Listener)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *LoadBalancer) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *LoadBalancer) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *LoadBalancer) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *LoadBalancer) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.

### HasLoadBalancerName

`func (o *LoadBalancer) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### GetLoadBalancerStickyCookiePolicies

`func (o *LoadBalancer) GetLoadBalancerStickyCookiePolicies() []LoadBalancerStickyCookiePolicy`

GetLoadBalancerStickyCookiePolicies returns the LoadBalancerStickyCookiePolicies field if non-nil, zero value otherwise.

### GetLoadBalancerStickyCookiePoliciesOk

`func (o *LoadBalancer) GetLoadBalancerStickyCookiePoliciesOk() (*[]LoadBalancerStickyCookiePolicy, bool)`

GetLoadBalancerStickyCookiePoliciesOk returns a tuple with the LoadBalancerStickyCookiePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerStickyCookiePolicies

`func (o *LoadBalancer) SetLoadBalancerStickyCookiePolicies(v []LoadBalancerStickyCookiePolicy)`

SetLoadBalancerStickyCookiePolicies sets LoadBalancerStickyCookiePolicies field to given value.

### HasLoadBalancerStickyCookiePolicies

`func (o *LoadBalancer) HasLoadBalancerStickyCookiePolicies() bool`

HasLoadBalancerStickyCookiePolicies returns a boolean if a field has been set.

### GetLoadBalancerType

`func (o *LoadBalancer) GetLoadBalancerType() string`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *LoadBalancer) GetLoadBalancerTypeOk() (*string, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *LoadBalancer) SetLoadBalancerType(v string)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *LoadBalancer) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### GetNetId

`func (o *LoadBalancer) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *LoadBalancer) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *LoadBalancer) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *LoadBalancer) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetPublicIp

`func (o *LoadBalancer) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *LoadBalancer) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *LoadBalancer) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *LoadBalancer) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetSecuredCookies

`func (o *LoadBalancer) GetSecuredCookies() bool`

GetSecuredCookies returns the SecuredCookies field if non-nil, zero value otherwise.

### GetSecuredCookiesOk

`func (o *LoadBalancer) GetSecuredCookiesOk() (*bool, bool)`

GetSecuredCookiesOk returns a tuple with the SecuredCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredCookies

`func (o *LoadBalancer) SetSecuredCookies(v bool)`

SetSecuredCookies sets SecuredCookies field to given value.

### HasSecuredCookies

`func (o *LoadBalancer) HasSecuredCookies() bool`

HasSecuredCookies returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *LoadBalancer) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *LoadBalancer) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *LoadBalancer) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *LoadBalancer) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSourceSecurityGroup

`func (o *LoadBalancer) GetSourceSecurityGroup() SourceSecurityGroup`

GetSourceSecurityGroup returns the SourceSecurityGroup field if non-nil, zero value otherwise.

### GetSourceSecurityGroupOk

`func (o *LoadBalancer) GetSourceSecurityGroupOk() (*SourceSecurityGroup, bool)`

GetSourceSecurityGroupOk returns a tuple with the SourceSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSecurityGroup

`func (o *LoadBalancer) SetSourceSecurityGroup(v SourceSecurityGroup)`

SetSourceSecurityGroup sets SourceSecurityGroup field to given value.

### HasSourceSecurityGroup

`func (o *LoadBalancer) HasSourceSecurityGroup() bool`

HasSourceSecurityGroup returns a boolean if a field has been set.

### GetSubnets

`func (o *LoadBalancer) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *LoadBalancer) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *LoadBalancer) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *LoadBalancer) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSubregionNames

`func (o *LoadBalancer) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *LoadBalancer) GetSubregionNamesOk() (*[]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionNames

`func (o *LoadBalancer) SetSubregionNames(v []string)`

SetSubregionNames sets SubregionNames field to given value.

### HasSubregionNames

`func (o *LoadBalancer) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### GetTags

`func (o *LoadBalancer) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoadBalancer) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoadBalancer) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoadBalancer) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


