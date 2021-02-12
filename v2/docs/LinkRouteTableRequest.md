# LinkRouteTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**RouteTableId** | **string** | The ID of the route table. | 
**SubnetId** | **string** | The ID of the Subnet. | 

## Methods

### NewLinkRouteTableRequest

`func NewLinkRouteTableRequest(routeTableId string, subnetId string, ) *LinkRouteTableRequest`

NewLinkRouteTableRequest instantiates a new LinkRouteTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRouteTableRequestWithDefaults

`func NewLinkRouteTableRequestWithDefaults() *LinkRouteTableRequest`

NewLinkRouteTableRequestWithDefaults instantiates a new LinkRouteTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *LinkRouteTableRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkRouteTableRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkRouteTableRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkRouteTableRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetRouteTableId

`func (o *LinkRouteTableRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *LinkRouteTableRequest) GetRouteTableIdOk() (*string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableId

`func (o *LinkRouteTableRequest) SetRouteTableId(v string)`

SetRouteTableId sets RouteTableId field to given value.


### GetSubnetId

`func (o *LinkRouteTableRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LinkRouteTableRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LinkRouteTableRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


