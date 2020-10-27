# CreateVirtualGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | The type of VPN connection supported by the virtual gateway (only &#x60;ipsec.1&#x60; is supported). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewCreateVirtualGatewayRequest

`func NewCreateVirtualGatewayRequest(connectionType string, ) *CreateVirtualGatewayRequest`

NewCreateVirtualGatewayRequest instantiates a new CreateVirtualGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVirtualGatewayRequestWithDefaults

`func NewCreateVirtualGatewayRequestWithDefaults() *CreateVirtualGatewayRequest`

NewCreateVirtualGatewayRequestWithDefaults instantiates a new CreateVirtualGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *CreateVirtualGatewayRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateVirtualGatewayRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateVirtualGatewayRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetDryRun

`func (o *CreateVirtualGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVirtualGatewayRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateVirtualGatewayRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateVirtualGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


