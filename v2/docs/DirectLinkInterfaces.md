# DirectLinkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the DirectLink interface. | [optional] 
**BgpAsn** | Pointer to **int32** | The BGP (Border Gateway Protocol) ASN (Autonomous System Number) on the customer&#39;s side of the DirectLink interface. | [optional] 
**BgpKey** | Pointer to **string** | The BGP authentication key. | [optional] 
**ClientPrivateIp** | Pointer to **string** | The IP on the customer&#39;s side of the DirectLink interface. | [optional] 
**DirectLinkId** | Pointer to **string** | The ID of the DirectLink. | [optional] 
**DirectLinkInterfaceId** | Pointer to **string** | The ID of the DirectLink interface. | [optional] 
**DirectLinkInterfaceName** | Pointer to **string** | The name of the DirectLink interface. | [optional] 
**InterfaceType** | Pointer to **string** | The type of the DirectLink interface (always &#x60;private&#x60;). | [optional] 
**Location** | Pointer to **string** | The datacenter where the DirectLink interface is located. | [optional] 
**OutscalePrivateIp** | Pointer to **string** | The IP on the OUTSCALE side of the DirectLink interface. | [optional] 
**State** | Pointer to **string** | The state of the DirectLink interface (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60; \\| &#x60;confirming&#x60; \\| &#x60;rejected&#x60; \\| &#x60;expired&#x60;). | [optional] 
**VirtualGatewayId** | Pointer to **string** | The ID of the target virtual gateway. | [optional] 
**Vlan** | Pointer to **int32** | The VLAN number associated with the DirectLink interface. | [optional] 

## Methods

### NewDirectLinkInterfaces

`func NewDirectLinkInterfaces() *DirectLinkInterfaces`

NewDirectLinkInterfaces instantiates a new DirectLinkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectLinkInterfacesWithDefaults

`func NewDirectLinkInterfacesWithDefaults() *DirectLinkInterfaces`

NewDirectLinkInterfacesWithDefaults instantiates a new DirectLinkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DirectLinkInterfaces) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DirectLinkInterfaces) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DirectLinkInterfaces) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *DirectLinkInterfaces) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBgpAsn

`func (o *DirectLinkInterfaces) GetBgpAsn() int32`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *DirectLinkInterfaces) GetBgpAsnOk() (*int32, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAsn

`func (o *DirectLinkInterfaces) SetBgpAsn(v int32)`

SetBgpAsn sets BgpAsn field to given value.

### HasBgpAsn

`func (o *DirectLinkInterfaces) HasBgpAsn() bool`

HasBgpAsn returns a boolean if a field has been set.

### GetBgpKey

`func (o *DirectLinkInterfaces) GetBgpKey() string`

GetBgpKey returns the BgpKey field if non-nil, zero value otherwise.

### GetBgpKeyOk

`func (o *DirectLinkInterfaces) GetBgpKeyOk() (*string, bool)`

GetBgpKeyOk returns a tuple with the BgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpKey

`func (o *DirectLinkInterfaces) SetBgpKey(v string)`

SetBgpKey sets BgpKey field to given value.

### HasBgpKey

`func (o *DirectLinkInterfaces) HasBgpKey() bool`

HasBgpKey returns a boolean if a field has been set.

### GetClientPrivateIp

`func (o *DirectLinkInterfaces) GetClientPrivateIp() string`

GetClientPrivateIp returns the ClientPrivateIp field if non-nil, zero value otherwise.

### GetClientPrivateIpOk

`func (o *DirectLinkInterfaces) GetClientPrivateIpOk() (*string, bool)`

GetClientPrivateIpOk returns a tuple with the ClientPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateIp

`func (o *DirectLinkInterfaces) SetClientPrivateIp(v string)`

SetClientPrivateIp sets ClientPrivateIp field to given value.

### HasClientPrivateIp

`func (o *DirectLinkInterfaces) HasClientPrivateIp() bool`

HasClientPrivateIp returns a boolean if a field has been set.

### GetDirectLinkId

`func (o *DirectLinkInterfaces) GetDirectLinkId() string`

GetDirectLinkId returns the DirectLinkId field if non-nil, zero value otherwise.

### GetDirectLinkIdOk

`func (o *DirectLinkInterfaces) GetDirectLinkIdOk() (*string, bool)`

GetDirectLinkIdOk returns a tuple with the DirectLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkId

`func (o *DirectLinkInterfaces) SetDirectLinkId(v string)`

SetDirectLinkId sets DirectLinkId field to given value.

### HasDirectLinkId

`func (o *DirectLinkInterfaces) HasDirectLinkId() bool`

HasDirectLinkId returns a boolean if a field has been set.

### GetDirectLinkInterfaceId

`func (o *DirectLinkInterfaces) GetDirectLinkInterfaceId() string`

GetDirectLinkInterfaceId returns the DirectLinkInterfaceId field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceIdOk

`func (o *DirectLinkInterfaces) GetDirectLinkInterfaceIdOk() (*string, bool)`

GetDirectLinkInterfaceIdOk returns a tuple with the DirectLinkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkInterfaceId

`func (o *DirectLinkInterfaces) SetDirectLinkInterfaceId(v string)`

SetDirectLinkInterfaceId sets DirectLinkInterfaceId field to given value.

### HasDirectLinkInterfaceId

`func (o *DirectLinkInterfaces) HasDirectLinkInterfaceId() bool`

HasDirectLinkInterfaceId returns a boolean if a field has been set.

### GetDirectLinkInterfaceName

`func (o *DirectLinkInterfaces) GetDirectLinkInterfaceName() string`

GetDirectLinkInterfaceName returns the DirectLinkInterfaceName field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceNameOk

`func (o *DirectLinkInterfaces) GetDirectLinkInterfaceNameOk() (*string, bool)`

GetDirectLinkInterfaceNameOk returns a tuple with the DirectLinkInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkInterfaceName

`func (o *DirectLinkInterfaces) SetDirectLinkInterfaceName(v string)`

SetDirectLinkInterfaceName sets DirectLinkInterfaceName field to given value.

### HasDirectLinkInterfaceName

`func (o *DirectLinkInterfaces) HasDirectLinkInterfaceName() bool`

HasDirectLinkInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *DirectLinkInterfaces) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *DirectLinkInterfaces) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *DirectLinkInterfaces) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *DirectLinkInterfaces) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetLocation

`func (o *DirectLinkInterfaces) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DirectLinkInterfaces) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DirectLinkInterfaces) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DirectLinkInterfaces) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOutscalePrivateIp

`func (o *DirectLinkInterfaces) GetOutscalePrivateIp() string`

GetOutscalePrivateIp returns the OutscalePrivateIp field if non-nil, zero value otherwise.

### GetOutscalePrivateIpOk

`func (o *DirectLinkInterfaces) GetOutscalePrivateIpOk() (*string, bool)`

GetOutscalePrivateIpOk returns a tuple with the OutscalePrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutscalePrivateIp

`func (o *DirectLinkInterfaces) SetOutscalePrivateIp(v string)`

SetOutscalePrivateIp sets OutscalePrivateIp field to given value.

### HasOutscalePrivateIp

`func (o *DirectLinkInterfaces) HasOutscalePrivateIp() bool`

HasOutscalePrivateIp returns a boolean if a field has been set.

### GetState

`func (o *DirectLinkInterfaces) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectLinkInterfaces) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectLinkInterfaces) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectLinkInterfaces) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVirtualGatewayId

`func (o *DirectLinkInterfaces) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *DirectLinkInterfaces) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *DirectLinkInterfaces) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.

### HasVirtualGatewayId

`func (o *DirectLinkInterfaces) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### GetVlan

`func (o *DirectLinkInterfaces) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *DirectLinkInterfaces) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *DirectLinkInterfaces) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *DirectLinkInterfaces) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


