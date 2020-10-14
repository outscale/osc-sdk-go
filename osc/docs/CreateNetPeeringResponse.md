# CreateNetPeeringResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetPeering** | Pointer to [**NetPeering**](NetPeering.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateNetPeeringResponse

`func NewCreateNetPeeringResponse() *CreateNetPeeringResponse`

NewCreateNetPeeringResponse instantiates a new CreateNetPeeringResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetPeeringResponseWithDefaults

`func NewCreateNetPeeringResponseWithDefaults() *CreateNetPeeringResponse`

NewCreateNetPeeringResponseWithDefaults instantiates a new CreateNetPeeringResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetPeering

`func (o *CreateNetPeeringResponse) GetNetPeering() NetPeering`

GetNetPeering returns the NetPeering field if non-nil, zero value otherwise.

### GetNetPeeringOk

`func (o *CreateNetPeeringResponse) GetNetPeeringOk() (*NetPeering, bool)`

GetNetPeeringOk returns a tuple with the NetPeering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPeering

`func (o *CreateNetPeeringResponse) SetNetPeering(v NetPeering)`

SetNetPeering sets NetPeering field to given value.

### HasNetPeering

`func (o *CreateNetPeeringResponse) HasNetPeering() bool`

HasNetPeering returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateNetPeeringResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateNetPeeringResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateNetPeeringResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateNetPeeringResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


