# CreateVpnConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGatewayId** | **string** | The ID of the client gateway. | 
**ConnectionType** | **string** | The type of VPN connection (only &#x60;ipsec.1&#x60; is supported). | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**StaticRoutesOnly** | Pointer to **bool** | By default or if false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute). | [optional] 
**VirtualGatewayId** | **string** | The ID of the virtual gateway. | 

## Methods

### NewCreateVpnConnectionRequest

`func NewCreateVpnConnectionRequest(clientGatewayId string, connectionType string, virtualGatewayId string, ) *CreateVpnConnectionRequest`

NewCreateVpnConnectionRequest instantiates a new CreateVpnConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVpnConnectionRequestWithDefaults

`func NewCreateVpnConnectionRequestWithDefaults() *CreateVpnConnectionRequest`

NewCreateVpnConnectionRequestWithDefaults instantiates a new CreateVpnConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientGatewayId

`func (o *CreateVpnConnectionRequest) GetClientGatewayId() string`

GetClientGatewayId returns the ClientGatewayId field if non-nil, zero value otherwise.

### GetClientGatewayIdOk

`func (o *CreateVpnConnectionRequest) GetClientGatewayIdOk() (*string, bool)`

GetClientGatewayIdOk returns a tuple with the ClientGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGatewayId

`func (o *CreateVpnConnectionRequest) SetClientGatewayId(v string)`

SetClientGatewayId sets ClientGatewayId field to given value.


### GetConnectionType

`func (o *CreateVpnConnectionRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateVpnConnectionRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateVpnConnectionRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetDryRun

`func (o *CreateVpnConnectionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVpnConnectionRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateVpnConnectionRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateVpnConnectionRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetStaticRoutesOnly

`func (o *CreateVpnConnectionRequest) GetStaticRoutesOnly() bool`

GetStaticRoutesOnly returns the StaticRoutesOnly field if non-nil, zero value otherwise.

### GetStaticRoutesOnlyOk

`func (o *CreateVpnConnectionRequest) GetStaticRoutesOnlyOk() (*bool, bool)`

GetStaticRoutesOnlyOk returns a tuple with the StaticRoutesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutesOnly

`func (o *CreateVpnConnectionRequest) SetStaticRoutesOnly(v bool)`

SetStaticRoutesOnly sets StaticRoutesOnly field to given value.

### HasStaticRoutesOnly

`func (o *CreateVpnConnectionRequest) HasStaticRoutesOnly() bool`

HasStaticRoutesOnly returns a boolean if a field has been set.

### GetVirtualGatewayId

`func (o *CreateVpnConnectionRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *CreateVpnConnectionRequest) GetVirtualGatewayIdOk() (*string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGatewayId

`func (o *CreateVpnConnectionRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId sets VirtualGatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


