# StateComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateCode** | Pointer to **string** | The code of the change of state. | [optional] 
**StateMessage** | Pointer to **string** | A message explaining the change of state. | [optional] 

## Methods

### NewStateComment

`func NewStateComment() *StateComment`

NewStateComment instantiates a new StateComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateCommentWithDefaults

`func NewStateCommentWithDefaults() *StateComment`

NewStateCommentWithDefaults instantiates a new StateComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateCode

`func (o *StateComment) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *StateComment) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *StateComment) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *StateComment) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetStateMessage

`func (o *StateComment) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *StateComment) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *StateComment) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *StateComment) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


