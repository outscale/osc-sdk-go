# UpdateNetAccessPointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddRouteTableIds** | Pointer to **[]string** | One or more IDs of route tables to associate with the specified Net access point. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NetAccessPointId** | **string** | The ID of the Net access point. | 
**RemoveRouteTableIds** | Pointer to **[]string** | One or more IDs of route tables to disassociate from the specified Net access point. | [optional] 

## Methods

### NewUpdateNetAccessPointRequest

`func NewUpdateNetAccessPointRequest(netAccessPointId string, ) *UpdateNetAccessPointRequest`

NewUpdateNetAccessPointRequest instantiates a new UpdateNetAccessPointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetAccessPointRequestWithDefaults

`func NewUpdateNetAccessPointRequestWithDefaults() *UpdateNetAccessPointRequest`

NewUpdateNetAccessPointRequestWithDefaults instantiates a new UpdateNetAccessPointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddRouteTableIds

`func (o *UpdateNetAccessPointRequest) GetAddRouteTableIds() []string`

GetAddRouteTableIds returns the AddRouteTableIds field if non-nil, zero value otherwise.

### GetAddRouteTableIdsOk

`func (o *UpdateNetAccessPointRequest) GetAddRouteTableIdsOk() (*[]string, bool)`

GetAddRouteTableIdsOk returns a tuple with the AddRouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRouteTableIds

`func (o *UpdateNetAccessPointRequest) SetAddRouteTableIds(v []string)`

SetAddRouteTableIds sets AddRouteTableIds field to given value.

### HasAddRouteTableIds

`func (o *UpdateNetAccessPointRequest) HasAddRouteTableIds() bool`

HasAddRouteTableIds returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateNetAccessPointRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateNetAccessPointRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateNetAccessPointRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateNetAccessPointRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetAccessPointId

`func (o *UpdateNetAccessPointRequest) GetNetAccessPointId() string`

GetNetAccessPointId returns the NetAccessPointId field if non-nil, zero value otherwise.

### GetNetAccessPointIdOk

`func (o *UpdateNetAccessPointRequest) GetNetAccessPointIdOk() (*string, bool)`

GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAccessPointId

`func (o *UpdateNetAccessPointRequest) SetNetAccessPointId(v string)`

SetNetAccessPointId sets NetAccessPointId field to given value.


### GetRemoveRouteTableIds

`func (o *UpdateNetAccessPointRequest) GetRemoveRouteTableIds() []string`

GetRemoveRouteTableIds returns the RemoveRouteTableIds field if non-nil, zero value otherwise.

### GetRemoveRouteTableIdsOk

`func (o *UpdateNetAccessPointRequest) GetRemoveRouteTableIdsOk() (*[]string, bool)`

GetRemoveRouteTableIdsOk returns a tuple with the RemoveRouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveRouteTableIds

`func (o *UpdateNetAccessPointRequest) SetRemoveRouteTableIds(v []string)`

SetRemoveRouteTableIds sets RemoveRouteTableIds field to given value.

### HasRemoveRouteTableIds

`func (o *UpdateNetAccessPointRequest) HasRemoveRouteTableIds() bool`

HasRemoveRouteTableIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


