# UpdateRouteTableLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LinkRouteTableId** | **string** | The ID of the current route table link. | 
**RouteTableId** | **string** | The ID of the new route table to associate with the Subnet. | 

## Methods

### NewUpdateRouteTableLinkRequest

`func NewUpdateRouteTableLinkRequest(linkRouteTableId string, routeTableId string, ) *UpdateRouteTableLinkRequest`

NewUpdateRouteTableLinkRequest instantiates a new UpdateRouteTableLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRouteTableLinkRequestWithDefaults

`func NewUpdateRouteTableLinkRequestWithDefaults() *UpdateRouteTableLinkRequest`

NewUpdateRouteTableLinkRequestWithDefaults instantiates a new UpdateRouteTableLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateRouteTableLinkRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateRouteTableLinkRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateRouteTableLinkRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateRouteTableLinkRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLinkRouteTableId

`func (o *UpdateRouteTableLinkRequest) GetLinkRouteTableId() string`

GetLinkRouteTableId returns the LinkRouteTableId field if non-nil, zero value otherwise.

### GetLinkRouteTableIdOk

`func (o *UpdateRouteTableLinkRequest) GetLinkRouteTableIdOk() (*string, bool)`

GetLinkRouteTableIdOk returns a tuple with the LinkRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRouteTableId

`func (o *UpdateRouteTableLinkRequest) SetLinkRouteTableId(v string)`

SetLinkRouteTableId sets LinkRouteTableId field to given value.


### GetRouteTableId

`func (o *UpdateRouteTableLinkRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *UpdateRouteTableLinkRequest) GetRouteTableIdOk() (*string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableId

`func (o *UpdateRouteTableLinkRequest) SetRouteTableId(v string)`

SetRouteTableId sets RouteTableId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


