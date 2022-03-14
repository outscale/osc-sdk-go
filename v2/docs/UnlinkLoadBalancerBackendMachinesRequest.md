# UnlinkLoadBalancerBackendMachinesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendIps** | Pointer to **[]string** |  One or more public IPs of back-end VMs. | [optional] 
**BackendVmIds** | Pointer to **[]string** |  One or more IDs of back-end VMs. | [optional] 
**DryRun** | Pointer to **bool** |  If true, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | **string** |  The name of the load balancer. | 

## Methods

### NewUnlinkLoadBalancerBackendMachinesRequest

`func NewUnlinkLoadBalancerBackendMachinesRequest(loadBalancerName string, ) *UnlinkLoadBalancerBackendMachinesRequest`

NewUnlinkLoadBalancerBackendMachinesRequest instantiates a new UnlinkLoadBalancerBackendMachinesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkLoadBalancerBackendMachinesRequestWithDefaults

`func NewUnlinkLoadBalancerBackendMachinesRequestWithDefaults() *UnlinkLoadBalancerBackendMachinesRequest`

NewUnlinkLoadBalancerBackendMachinesRequestWithDefaults instantiates a new UnlinkLoadBalancerBackendMachinesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendIps

`func (o *UnlinkLoadBalancerBackendMachinesRequest) GetBackendIps() []string`

GetBackendIps returns the BackendIps field if non-nil, zero value otherwise.

### GetBackendIpsOk

`func (o *UnlinkLoadBalancerBackendMachinesRequest) GetBackendIpsOk() (*[]string, bool)`

GetBackendIpsOk returns a tuple with the BackendIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendIps

`func (o *UnlinkLoadBalancerBackendMachinesRequest) SetBackendIps(v []string)`

SetBackendIps sets BackendIps field to given value.

### HasBackendIps

`func (o *UnlinkLoadBalancerBackendMachinesRequest) HasBackendIps() bool`

HasBackendIps returns a boolean if a field has been set.

### GetBackendVmIds

`func (o *UnlinkLoadBalancerBackendMachinesRequest) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *UnlinkLoadBalancerBackendMachinesRequest) GetBackendVmIdsOk() (*[]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVmIds

`func (o *UnlinkLoadBalancerBackendMachinesRequest) SetBackendVmIds(v []string)`

SetBackendVmIds sets BackendVmIds field to given value.

### HasBackendVmIds

`func (o *UnlinkLoadBalancerBackendMachinesRequest) HasBackendVmIds() bool`

HasBackendVmIds returns a boolean if a field has been set.

### GetDryRun

`func (o *UnlinkLoadBalancerBackendMachinesRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkLoadBalancerBackendMachinesRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkLoadBalancerBackendMachinesRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkLoadBalancerBackendMachinesRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *UnlinkLoadBalancerBackendMachinesRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *UnlinkLoadBalancerBackendMachinesRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *UnlinkLoadBalancerBackendMachinesRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


