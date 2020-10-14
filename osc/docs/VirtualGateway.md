# VirtualGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** | The type of VPN connection supported by the virtual gateway (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**NetToVirtualGatewayLinks** | Pointer to [**[]NetToVirtualGatewayLink**](NetToVirtualGatewayLink.md) | The Net to which the virtual gateway is attached. | [optional] 
**State** | Pointer to **string** | The state of the virtual gateway (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the virtual gateway. | [optional] 
**VirtualGatewayId** | Pointer to **string** | The ID of the virtual gateway. | [optional] 

## Methods

### NewVirtualGateway

`func NewVirtualGateway() *VirtualGateway`

NewVirtualGateway instantiates a new VirtualGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualGatewayWithDefaults

`func NewVirtualGatewayWithDefaults() *VirtualGateway`

NewVirtualGatewayWithDefaults instantiates a new VirtualGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *VirtualGateway) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *VirtualGateway) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *VirtualGateway) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *VirtualGateway) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetNetToVirtualGatewayLinks

`func (o *VirtualGateway) GetNetToVirtualGatewayLinks() []NetToVirtualGatewayLink`

GetNetToVirtualGatewayLinks returns the NetToVirtualGatewayLinks field if non-nil, zero value otherwise.

### GetNetToVirtualGatewayLinksOk

`func (o *VirtualGateway) GetNetToVirtualGatewayLinksOk() (*[]NetToVirtualGatewayLink, bool)`

GetNetToVirtualGatewayLinksOk returns a tuple with the NetToVirtualGatewayLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetToVirtualGatewayLinks

`func (o *VirtualGateway) SetNetToVirtualGatewayLinks(v []NetToVirtualGatewayLink)`

SetNetToVirtualGatewayLinks sets NetToVirtualGatewayLinks field to given value.

### HasNetToVirtualGatewayLinks

`func (o *VirtualGateway) HasNetToVirtualGatewayLinks() bool`

HasNetToVirtualGatewayLinks returns a boolean if a field has been set.

### GetState

`func (o *VirtualGateway) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualGateway) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualGateway) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VirtualGateway) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *VirtualGateway) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualGateway) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualGateway) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualGateway) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVirtualGatewayId

`func (o *VirtualGateway) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *VirtualGateway) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *VirtualGateway) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.

### HasVirtualGatewayId

`func (o *VirtualGateway) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


