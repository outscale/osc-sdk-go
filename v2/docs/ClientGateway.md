# ClientGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsn** | Pointer to **int32** | The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet. | [optional] 
**ClientGatewayId** | Pointer to **string** | The ID of the client gateway. | [optional] 
**ConnectionType** | Pointer to **string** | The type of communication tunnel used by the client gateway (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**PublicIp** | Pointer to **string** | The public IPv4 address of the client gateway (must be a fixed address into a NATed network). | [optional] 
**State** | Pointer to **string** | The state of the client gateway (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the client gateway. | [optional] 

## Methods

### NewClientGateway

`func NewClientGateway() *ClientGateway`

NewClientGateway instantiates a new ClientGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGatewayWithDefaults

`func NewClientGatewayWithDefaults() *ClientGateway`

NewClientGatewayWithDefaults instantiates a new ClientGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpAsn

`func (o *ClientGateway) GetBgpAsn() int32`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *ClientGateway) GetBgpAsnOk() (*int32, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAsn

`func (o *ClientGateway) SetBgpAsn(v int32)`

SetBgpAsn sets BgpAsn field to given value.

### HasBgpAsn

`func (o *ClientGateway) HasBgpAsn() bool`

HasBgpAsn returns a boolean if a field has been set.

### GetClientGatewayId

`func (o *ClientGateway) GetClientGatewayId() string`

GetClientGatewayId returns the ClientGatewayId field if non-nil, zero value otherwise.

### GetClientGatewayIdOk

`func (o *ClientGateway) GetClientGatewayIdOk() (*string, bool)`

GetClientGatewayIdOk returns a tuple with the ClientGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGatewayId

`func (o *ClientGateway) SetClientGatewayId(v string)`

SetClientGatewayId sets ClientGatewayId field to given value.

### HasClientGatewayId

`func (o *ClientGateway) HasClientGatewayId() bool`

HasClientGatewayId returns a boolean if a field has been set.

### GetConnectionType

`func (o *ClientGateway) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ClientGateway) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ClientGateway) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *ClientGateway) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetPublicIp

`func (o *ClientGateway) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ClientGateway) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ClientGateway) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ClientGateway) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetState

`func (o *ClientGateway) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClientGateway) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClientGateway) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ClientGateway) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *ClientGateway) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClientGateway) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClientGateway) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClientGateway) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


