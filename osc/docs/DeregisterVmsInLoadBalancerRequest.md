# DeregisterVmsInLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendVmIds** | **[]string** | One or more IDs of back-end VMs. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer. | 

## Methods

### NewDeregisterVmsInLoadBalancerRequest

`func NewDeregisterVmsInLoadBalancerRequest(backendVmIds []string, loadBalancerName string, ) *DeregisterVmsInLoadBalancerRequest`

NewDeregisterVmsInLoadBalancerRequest instantiates a new DeregisterVmsInLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeregisterVmsInLoadBalancerRequestWithDefaults

`func NewDeregisterVmsInLoadBalancerRequestWithDefaults() *DeregisterVmsInLoadBalancerRequest`

NewDeregisterVmsInLoadBalancerRequestWithDefaults instantiates a new DeregisterVmsInLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendVmIds

`func (o *DeregisterVmsInLoadBalancerRequest) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *DeregisterVmsInLoadBalancerRequest) GetBackendVmIdsOk() (*[]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVmIds

`func (o *DeregisterVmsInLoadBalancerRequest) SetBackendVmIds(v []string)`

SetBackendVmIds sets BackendVmIds field to given value.


### GetDryRun

`func (o *DeregisterVmsInLoadBalancerRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeregisterVmsInLoadBalancerRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeregisterVmsInLoadBalancerRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeregisterVmsInLoadBalancerRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *DeregisterVmsInLoadBalancerRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *DeregisterVmsInLoadBalancerRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *DeregisterVmsInLoadBalancerRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


