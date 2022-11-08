# NetPeeringState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Additional information about the state of the Net peering. | [optional] 
**Name** | Pointer to **string** | The state of the Net peering (&#x60;pending-acceptance&#x60; \\| &#x60;active&#x60; \\| &#x60;rejected&#x60; \\| &#x60;failed&#x60; \\| &#x60;expired&#x60; \\| &#x60;deleted&#x60;). | [optional] 

## Methods

### NewNetPeeringState

`func NewNetPeeringState() *NetPeeringState`

NewNetPeeringState instantiates a new NetPeeringState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetPeeringStateWithDefaults

`func NewNetPeeringStateWithDefaults() *NetPeeringState`

NewNetPeeringStateWithDefaults instantiates a new NetPeeringState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *NetPeeringState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NetPeeringState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NetPeeringState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NetPeeringState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *NetPeeringState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetPeeringState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetPeeringState) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetPeeringState) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


