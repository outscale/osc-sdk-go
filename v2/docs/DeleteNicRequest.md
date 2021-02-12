# DeleteNicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | **string** | The ID of the NIC you want to delete. | 

## Methods

### NewDeleteNicRequest

`func NewDeleteNicRequest(nicId string, ) *DeleteNicRequest`

NewDeleteNicRequest instantiates a new DeleteNicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteNicRequestWithDefaults

`func NewDeleteNicRequestWithDefaults() *DeleteNicRequest`

NewDeleteNicRequestWithDefaults instantiates a new DeleteNicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteNicRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteNicRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteNicRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteNicRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNicId

`func (o *DeleteNicRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *DeleteNicRequest) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *DeleteNicRequest) SetNicId(v string)`

SetNicId sets NicId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


