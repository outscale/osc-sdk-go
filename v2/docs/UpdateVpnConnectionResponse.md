# UpdateVpnConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VpnConnection** | Pointer to [**VpnConnection**](VpnConnection.md) |  | [optional] 

## Methods

### NewUpdateVpnConnectionResponse

`func NewUpdateVpnConnectionResponse() *UpdateVpnConnectionResponse`

NewUpdateVpnConnectionResponse instantiates a new UpdateVpnConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVpnConnectionResponseWithDefaults

`func NewUpdateVpnConnectionResponseWithDefaults() *UpdateVpnConnectionResponse`

NewUpdateVpnConnectionResponseWithDefaults instantiates a new UpdateVpnConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *UpdateVpnConnectionResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UpdateVpnConnectionResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UpdateVpnConnectionResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UpdateVpnConnectionResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVpnConnection

`func (o *UpdateVpnConnectionResponse) GetVpnConnection() VpnConnection`

GetVpnConnection returns the VpnConnection field if non-nil, zero value otherwise.

### GetVpnConnectionOk

`func (o *UpdateVpnConnectionResponse) GetVpnConnectionOk() (*VpnConnection, bool)`

GetVpnConnectionOk returns a tuple with the VpnConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnection

`func (o *UpdateVpnConnectionResponse) SetVpnConnection(v VpnConnection)`

SetVpnConnection sets VpnConnection field to given value.

### HasVpnConnection

`func (o *UpdateVpnConnectionResponse) HasVpnConnection() bool`

HasVpnConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


