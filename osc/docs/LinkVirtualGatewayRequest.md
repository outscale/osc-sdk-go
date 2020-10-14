# LinkVirtualGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | **string** | The ID of the Net to which you want to attach the virtual gateway. | 
**VirtualGatewayId** | **string** | The ID of the virtual gateway. | 

## Methods

### NewLinkVirtualGatewayRequest

`func NewLinkVirtualGatewayRequest(netId string, virtualGatewayId string, ) *LinkVirtualGatewayRequest`

NewLinkVirtualGatewayRequest instantiates a new LinkVirtualGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkVirtualGatewayRequestWithDefaults

`func NewLinkVirtualGatewayRequestWithDefaults() *LinkVirtualGatewayRequest`

NewLinkVirtualGatewayRequestWithDefaults instantiates a new LinkVirtualGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *LinkVirtualGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkVirtualGatewayRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkVirtualGatewayRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkVirtualGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetId

`func (o *LinkVirtualGatewayRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *LinkVirtualGatewayRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *LinkVirtualGatewayRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.


### GetVirtualGatewayId

`func (o *LinkVirtualGatewayRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *LinkVirtualGatewayRequest) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *LinkVirtualGatewayRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


