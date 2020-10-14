# NetToVirtualGatewayLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetId** | Pointer to **string** | The ID of the Net to which the virtual gateway is attached. | [optional] 
**State** | Pointer to **string** | The state of the attachment (&#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60; \\| &#x60;detached&#x60;). | [optional] 

## Methods

### NewNetToVirtualGatewayLink

`func NewNetToVirtualGatewayLink() *NetToVirtualGatewayLink`

NewNetToVirtualGatewayLink instantiates a new NetToVirtualGatewayLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetToVirtualGatewayLinkWithDefaults

`func NewNetToVirtualGatewayLinkWithDefaults() *NetToVirtualGatewayLink`

NewNetToVirtualGatewayLinkWithDefaults instantiates a new NetToVirtualGatewayLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetId

`func (o *NetToVirtualGatewayLink) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *NetToVirtualGatewayLink) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *NetToVirtualGatewayLink) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *NetToVirtualGatewayLink) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetState

`func (o *NetToVirtualGatewayLink) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetToVirtualGatewayLink) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetToVirtualGatewayLink) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NetToVirtualGatewayLink) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


