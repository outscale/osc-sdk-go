# DeleteRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpRange** | **string** | The exact IP range for the route. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**RouteTableId** | **string** | The ID of the route table from which you want to delete a route. | 

## Methods

### NewDeleteRouteRequest

`func NewDeleteRouteRequest(destinationIpRange string, routeTableId string, ) *DeleteRouteRequest`

NewDeleteRouteRequest instantiates a new DeleteRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRouteRequestWithDefaults

`func NewDeleteRouteRequestWithDefaults() *DeleteRouteRequest`

NewDeleteRouteRequestWithDefaults instantiates a new DeleteRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIpRange

`func (o *DeleteRouteRequest) GetDestinationIpRange() string`

GetDestinationIpRange returns the DestinationIpRange field if non-nil, zero value otherwise.

### GetDestinationIpRangeOk

`func (o *DeleteRouteRequest) GetDestinationIpRangeOk() (*string, bool)`

GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpRange

`func (o *DeleteRouteRequest) SetDestinationIpRange(v string)`

SetDestinationIpRange sets DestinationIpRange field to given value.


### GetDryRun

`func (o *DeleteRouteRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteRouteRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteRouteRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteRouteRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetRouteTableId

`func (o *DeleteRouteRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *DeleteRouteRequest) GetRouteTableIdOk() (*string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableId

`func (o *DeleteRouteRequest) SetRouteTableId(v string)`

SetRouteTableId sets RouteTableId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


