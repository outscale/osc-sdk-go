# Listener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendPort** | Pointer to **int32** | The port on which the back-end VM is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | [optional] 
**BackendProtocol** | Pointer to **string** | The protocol for routing traffic to back-end VMs (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60;). | [optional] 
**LoadBalancerPort** | Pointer to **int32** | The port on which the load balancer is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | [optional] 
**LoadBalancerProtocol** | Pointer to **string** | The routing protocol (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60;). | [optional] 
**PolicyNames** | Pointer to **[]string** | The names of the policies. If there are no policies enabled, the list is empty. | [optional] 
**ServerCertificateId** | Pointer to **string** | The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers &gt; OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns). | [optional] 

## Methods

### NewListener

`func NewListener() *Listener`

NewListener instantiates a new Listener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerWithDefaults

`func NewListenerWithDefaults() *Listener`

NewListenerWithDefaults instantiates a new Listener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendPort

`func (o *Listener) GetBackendPort() int32`

GetBackendPort returns the BackendPort field if non-nil, zero value otherwise.

### GetBackendPortOk

`func (o *Listener) GetBackendPortOk() (*int32, bool)`

GetBackendPortOk returns a tuple with the BackendPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendPort

`func (o *Listener) SetBackendPort(v int32)`

SetBackendPort sets BackendPort field to given value.

### HasBackendPort

`func (o *Listener) HasBackendPort() bool`

HasBackendPort returns a boolean if a field has been set.

### GetBackendProtocol

`func (o *Listener) GetBackendProtocol() string`

GetBackendProtocol returns the BackendProtocol field if non-nil, zero value otherwise.

### GetBackendProtocolOk

`func (o *Listener) GetBackendProtocolOk() (*string, bool)`

GetBackendProtocolOk returns a tuple with the BackendProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendProtocol

`func (o *Listener) SetBackendProtocol(v string)`

SetBackendProtocol sets BackendProtocol field to given value.

### HasBackendProtocol

`func (o *Listener) HasBackendProtocol() bool`

HasBackendProtocol returns a boolean if a field has been set.

### GetLoadBalancerPort

`func (o *Listener) GetLoadBalancerPort() int32`

GetLoadBalancerPort returns the LoadBalancerPort field if non-nil, zero value otherwise.

### GetLoadBalancerPortOk

`func (o *Listener) GetLoadBalancerPortOk() (*int32, bool)`

GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPort

`func (o *Listener) SetLoadBalancerPort(v int32)`

SetLoadBalancerPort sets LoadBalancerPort field to given value.

### HasLoadBalancerPort

`func (o *Listener) HasLoadBalancerPort() bool`

HasLoadBalancerPort returns a boolean if a field has been set.

### GetLoadBalancerProtocol

`func (o *Listener) GetLoadBalancerProtocol() string`

GetLoadBalancerProtocol returns the LoadBalancerProtocol field if non-nil, zero value otherwise.

### GetLoadBalancerProtocolOk

`func (o *Listener) GetLoadBalancerProtocolOk() (*string, bool)`

GetLoadBalancerProtocolOk returns a tuple with the LoadBalancerProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProtocol

`func (o *Listener) SetLoadBalancerProtocol(v string)`

SetLoadBalancerProtocol sets LoadBalancerProtocol field to given value.

### HasLoadBalancerProtocol

`func (o *Listener) HasLoadBalancerProtocol() bool`

HasLoadBalancerProtocol returns a boolean if a field has been set.

### GetPolicyNames

`func (o *Listener) GetPolicyNames() []string`

GetPolicyNames returns the PolicyNames field if non-nil, zero value otherwise.

### GetPolicyNamesOk

`func (o *Listener) GetPolicyNamesOk() (*[]string, bool)`

GetPolicyNamesOk returns a tuple with the PolicyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNames

`func (o *Listener) SetPolicyNames(v []string)`

SetPolicyNames sets PolicyNames field to given value.

### HasPolicyNames

`func (o *Listener) HasPolicyNames() bool`

HasPolicyNames returns a boolean if a field has been set.

### GetServerCertificateId

`func (o *Listener) GetServerCertificateId() string`

GetServerCertificateId returns the ServerCertificateId field if non-nil, zero value otherwise.

### GetServerCertificateIdOk

`func (o *Listener) GetServerCertificateIdOk() (*string, bool)`

GetServerCertificateIdOk returns a tuple with the ServerCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificateId

`func (o *Listener) SetServerCertificateId(v string)`

SetServerCertificateId sets ServerCertificateId field to given value.

### HasServerCertificateId

`func (o *Listener) HasServerCertificateId() bool`

HasServerCertificateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


