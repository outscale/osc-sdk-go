# ReadVmsHealthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendVmIds** | Pointer to **[]string** | One or more IDs of back-end VMs. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer. | 

## Methods

### NewReadVmsHealthRequest

`func NewReadVmsHealthRequest(loadBalancerName string, ) *ReadVmsHealthRequest`

NewReadVmsHealthRequest instantiates a new ReadVmsHealthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmsHealthRequestWithDefaults

`func NewReadVmsHealthRequestWithDefaults() *ReadVmsHealthRequest`

NewReadVmsHealthRequestWithDefaults instantiates a new ReadVmsHealthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendVmIds

`func (o *ReadVmsHealthRequest) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *ReadVmsHealthRequest) GetBackendVmIdsOk() (*[]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVmIds

`func (o *ReadVmsHealthRequest) SetBackendVmIds(v []string)`

SetBackendVmIds sets BackendVmIds field to given value.

### HasBackendVmIds

`func (o *ReadVmsHealthRequest) HasBackendVmIds() bool`

HasBackendVmIds returns a boolean if a field has been set.

### GetDryRun

`func (o *ReadVmsHealthRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadVmsHealthRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadVmsHealthRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadVmsHealthRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *ReadVmsHealthRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *ReadVmsHealthRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *ReadVmsHealthRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


