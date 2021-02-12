# CreateNetAccessPointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | **string** | The ID of the Net. | 
**RouteTableIds** | Pointer to **[]string** | One or more IDs of route tables to use for the connection. | [optional] 
**ServiceName** | **string** | The name of the service (in the format &#x60;com.outscale.region.service&#x60;). | 

## Methods

### NewCreateNetAccessPointRequest

`func NewCreateNetAccessPointRequest(netId string, serviceName string, ) *CreateNetAccessPointRequest`

NewCreateNetAccessPointRequest instantiates a new CreateNetAccessPointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetAccessPointRequestWithDefaults

`func NewCreateNetAccessPointRequestWithDefaults() *CreateNetAccessPointRequest`

NewCreateNetAccessPointRequestWithDefaults instantiates a new CreateNetAccessPointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateNetAccessPointRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateNetAccessPointRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateNetAccessPointRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateNetAccessPointRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetId

`func (o *CreateNetAccessPointRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *CreateNetAccessPointRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *CreateNetAccessPointRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.


### GetRouteTableIds

`func (o *CreateNetAccessPointRequest) GetRouteTableIds() []string`

GetRouteTableIds returns the RouteTableIds field if non-nil, zero value otherwise.

### GetRouteTableIdsOk

`func (o *CreateNetAccessPointRequest) GetRouteTableIdsOk() (*[]string, bool)`

GetRouteTableIdsOk returns a tuple with the RouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableIds

`func (o *CreateNetAccessPointRequest) SetRouteTableIds(v []string)`

SetRouteTableIds sets RouteTableIds field to given value.

### HasRouteTableIds

`func (o *CreateNetAccessPointRequest) HasRouteTableIds() bool`

HasRouteTableIds returns a boolean if a field has been set.

### GetServiceName

`func (o *CreateNetAccessPointRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CreateNetAccessPointRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CreateNetAccessPointRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


