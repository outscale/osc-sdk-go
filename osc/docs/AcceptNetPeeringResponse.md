# AcceptNetPeeringResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetPeering** | Pointer to [**NetPeering**](NetPeering.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewAcceptNetPeeringResponse

`func NewAcceptNetPeeringResponse() *AcceptNetPeeringResponse`

NewAcceptNetPeeringResponse instantiates a new AcceptNetPeeringResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptNetPeeringResponseWithDefaults

`func NewAcceptNetPeeringResponseWithDefaults() *AcceptNetPeeringResponse`

NewAcceptNetPeeringResponseWithDefaults instantiates a new AcceptNetPeeringResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetPeering

`func (o *AcceptNetPeeringResponse) GetNetPeering() NetPeering`

GetNetPeering returns the NetPeering field if non-nil, zero value otherwise.

### GetNetPeeringOk

`func (o *AcceptNetPeeringResponse) GetNetPeeringOk() (*NetPeering, bool)`

GetNetPeeringOk returns a tuple with the NetPeering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPeering

`func (o *AcceptNetPeeringResponse) SetNetPeering(v NetPeering)`

SetNetPeering sets NetPeering field to given value.

### HasNetPeering

`func (o *AcceptNetPeeringResponse) HasNetPeering() bool`

HasNetPeering returns a boolean if a field has been set.

### GetResponseContext

`func (o *AcceptNetPeeringResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *AcceptNetPeeringResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *AcceptNetPeeringResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *AcceptNetPeeringResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


