# ListenerForCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendPort** | **int32** | The port on which the back-end VM is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**BackendProtocol** | Pointer to **string** | The protocol for routing traffic to back-end VMs (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | [optional] 
**LoadBalancerPort** | **int32** | The port on which the load balancer is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**LoadBalancerProtocol** | **string** | The routing protocol (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | 
**ServerCertificateId** | Pointer to **string** | The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers &gt; Outscale Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat). | [optional] 

## Methods

### NewListenerForCreation

`func NewListenerForCreation(backendPort int32, loadBalancerPort int32, loadBalancerProtocol string, ) *ListenerForCreation`

NewListenerForCreation instantiates a new ListenerForCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerForCreationWithDefaults

`func NewListenerForCreationWithDefaults() *ListenerForCreation`

NewListenerForCreationWithDefaults instantiates a new ListenerForCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendPort

`func (o *ListenerForCreation) GetBackendPort() int32`

GetBackendPort returns the BackendPort field if non-nil, zero value otherwise.

### GetBackendPortOk

`func (o *ListenerForCreation) GetBackendPortOk() (*int32, bool)`

GetBackendPortOk returns a tuple with the BackendPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendPort

`func (o *ListenerForCreation) SetBackendPort(v int32)`

SetBackendPort sets BackendPort field to given value.


### GetBackendProtocol

`func (o *ListenerForCreation) GetBackendProtocol() string`

GetBackendProtocol returns the BackendProtocol field if non-nil, zero value otherwise.

### GetBackendProtocolOk

`func (o *ListenerForCreation) GetBackendProtocolOk() (*string, bool)`

GetBackendProtocolOk returns a tuple with the BackendProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendProtocol

`func (o *ListenerForCreation) SetBackendProtocol(v string)`

SetBackendProtocol sets BackendProtocol field to given value.

### HasBackendProtocol

`func (o *ListenerForCreation) HasBackendProtocol() bool`

HasBackendProtocol returns a boolean if a field has been set.

### GetLoadBalancerPort

`func (o *ListenerForCreation) GetLoadBalancerPort() int32`

GetLoadBalancerPort returns the LoadBalancerPort field if non-nil, zero value otherwise.

### GetLoadBalancerPortOk

`func (o *ListenerForCreation) GetLoadBalancerPortOk() (*int32, bool)`

GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPort

`func (o *ListenerForCreation) SetLoadBalancerPort(v int32)`

SetLoadBalancerPort sets LoadBalancerPort field to given value.


### GetLoadBalancerProtocol

`func (o *ListenerForCreation) GetLoadBalancerProtocol() string`

GetLoadBalancerProtocol returns the LoadBalancerProtocol field if non-nil, zero value otherwise.

### GetLoadBalancerProtocolOk

`func (o *ListenerForCreation) GetLoadBalancerProtocolOk() (*string, bool)`

GetLoadBalancerProtocolOk returns a tuple with the LoadBalancerProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProtocol

`func (o *ListenerForCreation) SetLoadBalancerProtocol(v string)`

SetLoadBalancerProtocol sets LoadBalancerProtocol field to given value.


### GetServerCertificateId

`func (o *ListenerForCreation) GetServerCertificateId() string`

GetServerCertificateId returns the ServerCertificateId field if non-nil, zero value otherwise.

### GetServerCertificateIdOk

`func (o *ListenerForCreation) GetServerCertificateIdOk() (*string, bool)`

GetServerCertificateIdOk returns a tuple with the ServerCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificateId

`func (o *ListenerForCreation) SetServerCertificateId(v string)`

SetServerCertificateId sets ServerCertificateId field to given value.

### HasServerCertificateId

`func (o *ListenerForCreation) HasServerCertificateId() bool`

HasServerCertificateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


