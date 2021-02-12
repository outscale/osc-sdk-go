# UpdateRoutePropagationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Enable** | **bool** | If true, a virtual gateway can propagate routes to a specified route table of a Net. If false, the propagation is disabled. | 
**RouteTableId** | **string** | The ID of the route table. | 
**VirtualGatewayId** | **string** | The ID of the virtual gateway. | 

## Methods

### NewUpdateRoutePropagationRequest

`func NewUpdateRoutePropagationRequest(enable bool, routeTableId string, virtualGatewayId string, ) *UpdateRoutePropagationRequest`

NewUpdateRoutePropagationRequest instantiates a new UpdateRoutePropagationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoutePropagationRequestWithDefaults

`func NewUpdateRoutePropagationRequestWithDefaults() *UpdateRoutePropagationRequest`

NewUpdateRoutePropagationRequestWithDefaults instantiates a new UpdateRoutePropagationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateRoutePropagationRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateRoutePropagationRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateRoutePropagationRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateRoutePropagationRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetEnable

`func (o *UpdateRoutePropagationRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UpdateRoutePropagationRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UpdateRoutePropagationRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetRouteTableId

`func (o *UpdateRoutePropagationRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *UpdateRoutePropagationRequest) GetRouteTableIdOk() (*string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableId

`func (o *UpdateRoutePropagationRequest) SetRouteTableId(v string)`

SetRouteTableId sets RouteTableId field to given value.


### GetVirtualGatewayId

`func (o *UpdateRoutePropagationRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *UpdateRoutePropagationRequest) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *UpdateRoutePropagationRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


