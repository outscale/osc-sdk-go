# DeleteVirtualGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**VirtualGatewayId** | **string** | The ID of the virtual gateway you want to delete. | 

## Methods

### NewDeleteVirtualGatewayRequest

`func NewDeleteVirtualGatewayRequest(virtualGatewayId string, ) *DeleteVirtualGatewayRequest`

NewDeleteVirtualGatewayRequest instantiates a new DeleteVirtualGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVirtualGatewayRequestWithDefaults

`func NewDeleteVirtualGatewayRequestWithDefaults() *DeleteVirtualGatewayRequest`

NewDeleteVirtualGatewayRequestWithDefaults instantiates a new DeleteVirtualGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteVirtualGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteVirtualGatewayRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteVirtualGatewayRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteVirtualGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVirtualGatewayId

`func (o *DeleteVirtualGatewayRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *DeleteVirtualGatewayRequest) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *DeleteVirtualGatewayRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


