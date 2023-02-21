# DirectLinkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsn** | **int32** | The BGP (Border Gateway Protocol) ASN (Autonomous System Number) on the customer&#39;s side of the DirectLink interface. This number must be between &#x60;64512&#x60; and &#x60;65534&#x60;. | 
**BgpKey** | Pointer to **string** | The BGP authentication key. | [optional] 
**ClientPrivateIp** | Pointer to **string** | The IP on the customer&#39;s side of the DirectLink interface. | [optional] 
**DirectLinkInterfaceName** | **string** | The name of the DirectLink interface. | 
**OutscalePrivateIp** | Pointer to **string** | The IP on the OUTSCALE side of the DirectLink interface. | [optional] 
**VirtualGatewayId** | **string** | The ID of the target virtual gateway. | 
**Vlan** | **int32** | The VLAN number associated with the DirectLink interface. | 

## Methods

### NewDirectLinkInterface

`func NewDirectLinkInterface(bgpAsn int32, directLinkInterfaceName string, virtualGatewayId string, vlan int32, ) *DirectLinkInterface`

NewDirectLinkInterface instantiates a new DirectLinkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectLinkInterfaceWithDefaults

`func NewDirectLinkInterfaceWithDefaults() *DirectLinkInterface`

NewDirectLinkInterfaceWithDefaults instantiates a new DirectLinkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpAsn

`func (o *DirectLinkInterface) GetBgpAsn() int32`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *DirectLinkInterface) GetBgpAsnOk() (*int32, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAsn

`func (o *DirectLinkInterface) SetBgpAsn(v int32)`

SetBgpAsn sets BgpAsn field to given value.


### GetBgpKey

`func (o *DirectLinkInterface) GetBgpKey() string`

GetBgpKey returns the BgpKey field if non-nil, zero value otherwise.

### GetBgpKeyOk

`func (o *DirectLinkInterface) GetBgpKeyOk() (*string, bool)`

GetBgpKeyOk returns a tuple with the BgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpKey

`func (o *DirectLinkInterface) SetBgpKey(v string)`

SetBgpKey sets BgpKey field to given value.

### HasBgpKey

`func (o *DirectLinkInterface) HasBgpKey() bool`

HasBgpKey returns a boolean if a field has been set.

### GetClientPrivateIp

`func (o *DirectLinkInterface) GetClientPrivateIp() string`

GetClientPrivateIp returns the ClientPrivateIp field if non-nil, zero value otherwise.

### GetClientPrivateIpOk

`func (o *DirectLinkInterface) GetClientPrivateIpOk() (*string, bool)`

GetClientPrivateIpOk returns a tuple with the ClientPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateIp

`func (o *DirectLinkInterface) SetClientPrivateIp(v string)`

SetClientPrivateIp sets ClientPrivateIp field to given value.

### HasClientPrivateIp

`func (o *DirectLinkInterface) HasClientPrivateIp() bool`

HasClientPrivateIp returns a boolean if a field has been set.

### GetDirectLinkInterfaceName

`func (o *DirectLinkInterface) GetDirectLinkInterfaceName() string`

GetDirectLinkInterfaceName returns the DirectLinkInterfaceName field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceNameOk

`func (o *DirectLinkInterface) GetDirectLinkInterfaceNameOk() (*string, bool)`

GetDirectLinkInterfaceNameOk returns a tuple with the DirectLinkInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkInterfaceName

`func (o *DirectLinkInterface) SetDirectLinkInterfaceName(v string)`

SetDirectLinkInterfaceName sets DirectLinkInterfaceName field to given value.


### GetOutscalePrivateIp

`func (o *DirectLinkInterface) GetOutscalePrivateIp() string`

GetOutscalePrivateIp returns the OutscalePrivateIp field if non-nil, zero value otherwise.

### GetOutscalePrivateIpOk

`func (o *DirectLinkInterface) GetOutscalePrivateIpOk() (*string, bool)`

GetOutscalePrivateIpOk returns a tuple with the OutscalePrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutscalePrivateIp

`func (o *DirectLinkInterface) SetOutscalePrivateIp(v string)`

SetOutscalePrivateIp sets OutscalePrivateIp field to given value.

### HasOutscalePrivateIp

`func (o *DirectLinkInterface) HasOutscalePrivateIp() bool`

HasOutscalePrivateIp returns a boolean if a field has been set.

### GetVirtualGatewayId

`func (o *DirectLinkInterface) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *DirectLinkInterface) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *DirectLinkInterface) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.


### GetVlan

`func (o *DirectLinkInterface) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *DirectLinkInterface) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *DirectLinkInterface) SetVlan(v int32)`

SetVlan sets Vlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


