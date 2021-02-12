# UnlinkVirtualGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | **string** | The ID of the Net from which you want to detach the virtual gateway. | 
**VirtualGatewayId** | **string** | The ID of the virtual gateway. | 

## Methods

### NewUnlinkVirtualGatewayRequest

`func NewUnlinkVirtualGatewayRequest(netId string, virtualGatewayId string, ) *UnlinkVirtualGatewayRequest`

NewUnlinkVirtualGatewayRequest instantiates a new UnlinkVirtualGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkVirtualGatewayRequestWithDefaults

`func NewUnlinkVirtualGatewayRequestWithDefaults() *UnlinkVirtualGatewayRequest`

NewUnlinkVirtualGatewayRequestWithDefaults instantiates a new UnlinkVirtualGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UnlinkVirtualGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkVirtualGatewayRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkVirtualGatewayRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkVirtualGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetId

`func (o *UnlinkVirtualGatewayRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *UnlinkVirtualGatewayRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *UnlinkVirtualGatewayRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.


### GetVirtualGatewayId

`func (o *UnlinkVirtualGatewayRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *UnlinkVirtualGatewayRequest) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *UnlinkVirtualGatewayRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


