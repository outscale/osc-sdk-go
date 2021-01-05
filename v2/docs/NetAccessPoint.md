# NetAccessPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAccessPointId** | Pointer to **string** | The ID of the Net access point. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net with which the Net access point is associated. | [optional] 
**RouteTableIds** | Pointer to **[]string** | The ID of the route tables associated with the Net access point. | [optional] 
**ServiceName** | Pointer to **string** | The name of the service with which the Net access point is associated. | [optional] 
**State** | Pointer to **string** | The state of the Net access point (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the Net access point. | [optional] 

## Methods

### NewNetAccessPoint

`func NewNetAccessPoint() *NetAccessPoint`

NewNetAccessPoint instantiates a new NetAccessPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetAccessPointWithDefaults

`func NewNetAccessPointWithDefaults() *NetAccessPoint`

NewNetAccessPointWithDefaults instantiates a new NetAccessPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetAccessPointId

`func (o *NetAccessPoint) GetNetAccessPointId() string`

GetNetAccessPointId returns the NetAccessPointId field if non-nil, zero value otherwise.

### GetNetAccessPointIdOk

`func (o *NetAccessPoint) GetNetAccessPointIdOk() (*string, bool)`

GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAccessPointId

`func (o *NetAccessPoint) SetNetAccessPointId(v string)`

SetNetAccessPointId sets NetAccessPointId field to given value.

### HasNetAccessPointId

`func (o *NetAccessPoint) HasNetAccessPointId() bool`

HasNetAccessPointId returns a boolean if a field has been set.

### GetNetId

`func (o *NetAccessPoint) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *NetAccessPoint) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *NetAccessPoint) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *NetAccessPoint) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetRouteTableIds

`func (o *NetAccessPoint) GetRouteTableIds() []string`

GetRouteTableIds returns the RouteTableIds field if non-nil, zero value otherwise.

### GetRouteTableIdsOk

`func (o *NetAccessPoint) GetRouteTableIdsOk() (*[]string, bool)`

GetRouteTableIdsOk returns a tuple with the RouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableIds

`func (o *NetAccessPoint) SetRouteTableIds(v []string)`

SetRouteTableIds sets RouteTableIds field to given value.

### HasRouteTableIds

`func (o *NetAccessPoint) HasRouteTableIds() bool`

HasRouteTableIds returns a boolean if a field has been set.

### GetServiceName

`func (o *NetAccessPoint) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NetAccessPoint) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *NetAccessPoint) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *NetAccessPoint) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetState

`func (o *NetAccessPoint) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetAccessPoint) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetAccessPoint) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NetAccessPoint) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *NetAccessPoint) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetAccessPoint) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetAccessPoint) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NetAccessPoint) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


