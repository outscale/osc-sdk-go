# UpdateSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**MapPublicIpOnLaunch** | **bool** | If &#x60;true&#x60;, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet. | 
**SubnetId** | **string** | The ID of the Subnet. | 

## Methods

### NewUpdateSubnetRequest

`func NewUpdateSubnetRequest(mapPublicIpOnLaunch bool, subnetId string, ) *UpdateSubnetRequest`

NewUpdateSubnetRequest instantiates a new UpdateSubnetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubnetRequestWithDefaults

`func NewUpdateSubnetRequestWithDefaults() *UpdateSubnetRequest`

NewUpdateSubnetRequestWithDefaults instantiates a new UpdateSubnetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateSubnetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateSubnetRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateSubnetRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateSubnetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMapPublicIpOnLaunch

`func (o *UpdateSubnetRequest) GetMapPublicIpOnLaunch() bool`

GetMapPublicIpOnLaunch returns the MapPublicIpOnLaunch field if non-nil, zero value otherwise.

### GetMapPublicIpOnLaunchOk

`func (o *UpdateSubnetRequest) GetMapPublicIpOnLaunchOk() (*bool, bool)`

GetMapPublicIpOnLaunchOk returns a tuple with the MapPublicIpOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapPublicIpOnLaunch

`func (o *UpdateSubnetRequest) SetMapPublicIpOnLaunch(v bool)`

SetMapPublicIpOnLaunch sets MapPublicIpOnLaunch field to given value.


### GetSubnetId

`func (o *UpdateSubnetRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *UpdateSubnetRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *UpdateSubnetRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


