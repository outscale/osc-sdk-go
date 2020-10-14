# CreateClientGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsn** | **int32** | An Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet. | 
**ConnectionType** | **string** | The communication protocol used to establish tunnel with your client gateway (only &#x60;ipsec.1&#x60; is supported). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**PublicIp** | **string** | The public fixed IPv4 address of your client gateway. | 

## Methods

### NewCreateClientGatewayRequest

`func NewCreateClientGatewayRequest(bgpAsn int32, connectionType string, publicIp string, ) *CreateClientGatewayRequest`

NewCreateClientGatewayRequest instantiates a new CreateClientGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientGatewayRequestWithDefaults

`func NewCreateClientGatewayRequestWithDefaults() *CreateClientGatewayRequest`

NewCreateClientGatewayRequestWithDefaults instantiates a new CreateClientGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpAsn

`func (o *CreateClientGatewayRequest) GetBgpAsn() int32`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *CreateClientGatewayRequest) GetBgpAsnOk() (*int32, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAsn

`func (o *CreateClientGatewayRequest) SetBgpAsn(v int32)`

SetBgpAsn sets BgpAsn field to given value.


### GetConnectionType

`func (o *CreateClientGatewayRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateClientGatewayRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateClientGatewayRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetDryRun

`func (o *CreateClientGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateClientGatewayRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateClientGatewayRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateClientGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPublicIp

`func (o *CreateClientGatewayRequest) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CreateClientGatewayRequest) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CreateClientGatewayRequest) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


