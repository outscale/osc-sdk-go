# DeleteKeypairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**KeypairId** | Pointer to **string** | The ID of the keypair you want to delete. | [optional] 
**KeypairName** | Pointer to **string** | The name of the keypair you want to delete. | [optional] 

## Methods

### NewDeleteKeypairRequest

`func NewDeleteKeypairRequest() *DeleteKeypairRequest`

NewDeleteKeypairRequest instantiates a new DeleteKeypairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteKeypairRequestWithDefaults

`func NewDeleteKeypairRequestWithDefaults() *DeleteKeypairRequest`

NewDeleteKeypairRequestWithDefaults instantiates a new DeleteKeypairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteKeypairRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteKeypairRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteKeypairRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteKeypairRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetKeypairId

`func (o *DeleteKeypairRequest) GetKeypairId() string`

GetKeypairId returns the KeypairId field if non-nil, zero value otherwise.

### GetKeypairIdOk

`func (o *DeleteKeypairRequest) GetKeypairIdOk() (*string, bool)`

GetKeypairIdOk returns a tuple with the KeypairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairId

`func (o *DeleteKeypairRequest) SetKeypairId(v string)`

SetKeypairId sets KeypairId field to given value.

### HasKeypairId

`func (o *DeleteKeypairRequest) HasKeypairId() bool`

HasKeypairId returns a boolean if a field has been set.

### GetKeypairName

`func (o *DeleteKeypairRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *DeleteKeypairRequest) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *DeleteKeypairRequest) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *DeleteKeypairRequest) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


