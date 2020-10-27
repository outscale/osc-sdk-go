# ReadClientGatewaysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGateways** | Pointer to [**[]ClientGateway**](ClientGateway.md) | Information about one or more client gateways. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadClientGatewaysResponse

`func NewReadClientGatewaysResponse() *ReadClientGatewaysResponse`

NewReadClientGatewaysResponse instantiates a new ReadClientGatewaysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadClientGatewaysResponseWithDefaults

`func NewReadClientGatewaysResponseWithDefaults() *ReadClientGatewaysResponse`

NewReadClientGatewaysResponseWithDefaults instantiates a new ReadClientGatewaysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientGateways

`func (o *ReadClientGatewaysResponse) GetClientGateways() []ClientGateway`

GetClientGateways returns the ClientGateways field if non-nil, zero value otherwise.

### GetClientGatewaysOk

`func (o *ReadClientGatewaysResponse) GetClientGatewaysOk() (*[]ClientGateway, bool)`

GetClientGatewaysOk returns a tuple with the ClientGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGateways

`func (o *ReadClientGatewaysResponse) SetClientGateways(v []ClientGateway)`

SetClientGateways sets ClientGateways field to given value.

### HasClientGateways

`func (o *ReadClientGatewaysResponse) HasClientGateways() bool`

HasClientGateways returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadClientGatewaysResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadClientGatewaysResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadClientGatewaysResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadClientGatewaysResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


