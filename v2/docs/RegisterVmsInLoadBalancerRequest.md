# RegisterVmsInLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendVmIds** | **[]string** | One or more IDs of back-end VMs.&lt;br /&gt; Specifying the same ID several times has no effect as each back-end VM has equal weight. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer. | 

## Methods

### NewRegisterVmsInLoadBalancerRequest

`func NewRegisterVmsInLoadBalancerRequest(backendVmIds []string, loadBalancerName string, ) *RegisterVmsInLoadBalancerRequest`

NewRegisterVmsInLoadBalancerRequest instantiates a new RegisterVmsInLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterVmsInLoadBalancerRequestWithDefaults

`func NewRegisterVmsInLoadBalancerRequestWithDefaults() *RegisterVmsInLoadBalancerRequest`

NewRegisterVmsInLoadBalancerRequestWithDefaults instantiates a new RegisterVmsInLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendVmIds

`func (o *RegisterVmsInLoadBalancerRequest) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *RegisterVmsInLoadBalancerRequest) GetBackendVmIdsOk() (*[]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVmIds

`func (o *RegisterVmsInLoadBalancerRequest) SetBackendVmIds(v []string)`

SetBackendVmIds sets BackendVmIds field to given value.


### GetDryRun

`func (o *RegisterVmsInLoadBalancerRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *RegisterVmsInLoadBalancerRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *RegisterVmsInLoadBalancerRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *RegisterVmsInLoadBalancerRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *RegisterVmsInLoadBalancerRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *RegisterVmsInLoadBalancerRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *RegisterVmsInLoadBalancerRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


