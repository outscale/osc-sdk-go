# CreateClientGatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGateway** | Pointer to [**ClientGateway**](ClientGateway.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateClientGatewayResponse

`func NewCreateClientGatewayResponse() *CreateClientGatewayResponse`

NewCreateClientGatewayResponse instantiates a new CreateClientGatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientGatewayResponseWithDefaults

`func NewCreateClientGatewayResponseWithDefaults() *CreateClientGatewayResponse`

NewCreateClientGatewayResponseWithDefaults instantiates a new CreateClientGatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientGateway

`func (o *CreateClientGatewayResponse) GetClientGateway() ClientGateway`

GetClientGateway returns the ClientGateway field if non-nil, zero value otherwise.

### GetClientGatewayOk

`func (o *CreateClientGatewayResponse) GetClientGatewayOk() (*ClientGateway, bool)`

GetClientGatewayOk returns a tuple with the ClientGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGateway

`func (o *CreateClientGatewayResponse) SetClientGateway(v ClientGateway)`

SetClientGateway sets ClientGateway field to given value.

### HasClientGateway

`func (o *CreateClientGatewayResponse) HasClientGateway() bool`

HasClientGateway returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateClientGatewayResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateClientGatewayResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateClientGatewayResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateClientGatewayResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


