# ReadVirtualGatewaysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VirtualGateways** | Pointer to [**[]VirtualGateway**](VirtualGateway.md) | Information about one or more virtual gateways. | [optional] 

## Methods

### NewReadVirtualGatewaysResponse

`func NewReadVirtualGatewaysResponse() *ReadVirtualGatewaysResponse`

NewReadVirtualGatewaysResponse instantiates a new ReadVirtualGatewaysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVirtualGatewaysResponseWithDefaults

`func NewReadVirtualGatewaysResponseWithDefaults() *ReadVirtualGatewaysResponse`

NewReadVirtualGatewaysResponseWithDefaults instantiates a new ReadVirtualGatewaysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadVirtualGatewaysResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVirtualGatewaysResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVirtualGatewaysResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVirtualGatewaysResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVirtualGateways

`func (o *ReadVirtualGatewaysResponse) GetVirtualGateways() []VirtualGateway`

GetVirtualGateways returns the VirtualGateways field if non-nil, zero value otherwise.

### GetVirtualGatewaysOk

`func (o *ReadVirtualGatewaysResponse) GetVirtualGatewaysOk() (*[]VirtualGateway, bool)`

GetVirtualGatewaysOk returns a tuple with the VirtualGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGateways

`func (o *ReadVirtualGatewaysResponse) SetVirtualGateways(v []VirtualGateway)`

SetVirtualGateways sets VirtualGateways field to given value.

### HasVirtualGateways

`func (o *ReadVirtualGatewaysResponse) HasVirtualGateways() bool`

HasVirtualGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


