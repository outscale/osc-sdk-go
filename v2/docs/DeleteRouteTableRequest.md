# DeleteRouteTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**RouteTableId** | **string** | The ID of the route table you want to delete. | 

## Methods

### NewDeleteRouteTableRequest

`func NewDeleteRouteTableRequest(routeTableId string, ) *DeleteRouteTableRequest`

NewDeleteRouteTableRequest instantiates a new DeleteRouteTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRouteTableRequestWithDefaults

`func NewDeleteRouteTableRequestWithDefaults() *DeleteRouteTableRequest`

NewDeleteRouteTableRequestWithDefaults instantiates a new DeleteRouteTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteRouteTableRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteRouteTableRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteRouteTableRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteRouteTableRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetRouteTableId

`func (o *DeleteRouteTableRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *DeleteRouteTableRequest) GetRouteTableIdOk() (*string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableId

`func (o *DeleteRouteTableRequest) SetRouteTableId(v string)`

SetRouteTableId sets RouteTableId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


