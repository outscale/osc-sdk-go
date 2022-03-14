# LinkLoadBalancerBackendMachinesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendIps** | Pointer to **[]string** |  One or more public IPs of back-end VMs. | [optional] 
**BackendVmIds** | Pointer to **[]string** |  One or more IDs of back-end VMs. | [optional] 
**DryRun** | Pointer to **bool** |  If true, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | **string** |  The name of the load balancer.  | 

## Methods

### NewLinkLoadBalancerBackendMachinesRequest

`func NewLinkLoadBalancerBackendMachinesRequest(loadBalancerName string, ) *LinkLoadBalancerBackendMachinesRequest`

NewLinkLoadBalancerBackendMachinesRequest instantiates a new LinkLoadBalancerBackendMachinesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkLoadBalancerBackendMachinesRequestWithDefaults

`func NewLinkLoadBalancerBackendMachinesRequestWithDefaults() *LinkLoadBalancerBackendMachinesRequest`

NewLinkLoadBalancerBackendMachinesRequestWithDefaults instantiates a new LinkLoadBalancerBackendMachinesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendIps

`func (o *LinkLoadBalancerBackendMachinesRequest) GetBackendIps() []string`

GetBackendIps returns the BackendIps field if non-nil, zero value otherwise.

### GetBackendIpsOk

`func (o *LinkLoadBalancerBackendMachinesRequest) GetBackendIpsOk() (*[]string, bool)`

GetBackendIpsOk returns a tuple with the BackendIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendIps

`func (o *LinkLoadBalancerBackendMachinesRequest) SetBackendIps(v []string)`

SetBackendIps sets BackendIps field to given value.

### HasBackendIps

`func (o *LinkLoadBalancerBackendMachinesRequest) HasBackendIps() bool`

HasBackendIps returns a boolean if a field has been set.

### GetBackendVmIds

`func (o *LinkLoadBalancerBackendMachinesRequest) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *LinkLoadBalancerBackendMachinesRequest) GetBackendVmIdsOk() (*[]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVmIds

`func (o *LinkLoadBalancerBackendMachinesRequest) SetBackendVmIds(v []string)`

SetBackendVmIds sets BackendVmIds field to given value.

### HasBackendVmIds

`func (o *LinkLoadBalancerBackendMachinesRequest) HasBackendVmIds() bool`

HasBackendVmIds returns a boolean if a field has been set.

### GetDryRun

`func (o *LinkLoadBalancerBackendMachinesRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkLoadBalancerBackendMachinesRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkLoadBalancerBackendMachinesRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkLoadBalancerBackendMachinesRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *LinkLoadBalancerBackendMachinesRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *LinkLoadBalancerBackendMachinesRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *LinkLoadBalancerBackendMachinesRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


