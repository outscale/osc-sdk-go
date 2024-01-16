# NetPeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccepterNet** | Pointer to [**AccepterNet**](AccepterNet.md) |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | The date and time at which the Net peerings expire. | [optional] 
**NetPeeringId** | Pointer to **string** | The ID of the Net peering. | [optional] 
**SourceNet** | Pointer to [**SourceNet**](SourceNet.md) |  | [optional] 
**State** | Pointer to [**NetPeeringState**](NetPeeringState.md) |  | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the Net peering. | [optional] 

## Methods

### NewNetPeering

`func NewNetPeering() *NetPeering`

NewNetPeering instantiates a new NetPeering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetPeeringWithDefaults

`func NewNetPeeringWithDefaults() *NetPeering`

NewNetPeeringWithDefaults instantiates a new NetPeering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepterNet

`func (o *NetPeering) GetAccepterNet() AccepterNet`

GetAccepterNet returns the AccepterNet field if non-nil, zero value otherwise.

### GetAccepterNetOk

`func (o *NetPeering) GetAccepterNetOk() (*AccepterNet, bool)`

GetAccepterNetOk returns a tuple with the AccepterNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterNet

`func (o *NetPeering) SetAccepterNet(v AccepterNet)`

SetAccepterNet sets AccepterNet field to given value.

### HasAccepterNet

`func (o *NetPeering) HasAccepterNet() bool`

HasAccepterNet returns a boolean if a field has been set.

### GetExpirationDate

`func (o *NetPeering) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *NetPeering) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *NetPeering) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *NetPeering) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *NetPeering) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *NetPeering) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetNetPeeringId

`func (o *NetPeering) GetNetPeeringId() string`

GetNetPeeringId returns the NetPeeringId field if non-nil, zero value otherwise.

### GetNetPeeringIdOk

`func (o *NetPeering) GetNetPeeringIdOk() (*string, bool)`

GetNetPeeringIdOk returns a tuple with the NetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPeeringId

`func (o *NetPeering) SetNetPeeringId(v string)`

SetNetPeeringId sets NetPeeringId field to given value.

### HasNetPeeringId

`func (o *NetPeering) HasNetPeeringId() bool`

HasNetPeeringId returns a boolean if a field has been set.

### GetSourceNet

`func (o *NetPeering) GetSourceNet() SourceNet`

GetSourceNet returns the SourceNet field if non-nil, zero value otherwise.

### GetSourceNetOk

`func (o *NetPeering) GetSourceNetOk() (*SourceNet, bool)`

GetSourceNetOk returns a tuple with the SourceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNet

`func (o *NetPeering) SetSourceNet(v SourceNet)`

SetSourceNet sets SourceNet field to given value.

### HasSourceNet

`func (o *NetPeering) HasSourceNet() bool`

HasSourceNet returns a boolean if a field has been set.

### GetState

`func (o *NetPeering) GetState() NetPeeringState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetPeering) GetStateOk() (*NetPeeringState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetPeering) SetState(v NetPeeringState)`

SetState sets State field to given value.

### HasState

`func (o *NetPeering) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *NetPeering) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetPeering) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetPeering) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NetPeering) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


