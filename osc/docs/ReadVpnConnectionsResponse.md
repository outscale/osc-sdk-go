# ReadVpnConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VpnConnections** | Pointer to [**[]VpnConnection**](VpnConnection.md) | Information about one or more VPN connections. | [optional] 

## Methods

### NewReadVpnConnectionsResponse

`func NewReadVpnConnectionsResponse() *ReadVpnConnectionsResponse`

NewReadVpnConnectionsResponse instantiates a new ReadVpnConnectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVpnConnectionsResponseWithDefaults

`func NewReadVpnConnectionsResponseWithDefaults() *ReadVpnConnectionsResponse`

NewReadVpnConnectionsResponseWithDefaults instantiates a new ReadVpnConnectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadVpnConnectionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVpnConnectionsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVpnConnectionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVpnConnectionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVpnConnections

`func (o *ReadVpnConnectionsResponse) GetVpnConnections() []VpnConnection`

GetVpnConnections returns the VpnConnections field if non-nil, zero value otherwise.

### GetVpnConnectionsOk

`func (o *ReadVpnConnectionsResponse) GetVpnConnectionsOk() (*[]VpnConnection, bool)`

GetVpnConnectionsOk returns a tuple with the VpnConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnections

`func (o *ReadVpnConnectionsResponse) SetVpnConnections(v []VpnConnection)`

SetVpnConnections sets VpnConnections field to given value.

### HasVpnConnections

`func (o *ReadVpnConnectionsResponse) HasVpnConnections() bool`

HasVpnConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


