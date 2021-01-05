# DeleteCaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaId** | **string** | The ID of the CA you want to delete. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewDeleteCaRequest

`func NewDeleteCaRequest(caId string, ) *DeleteCaRequest`

NewDeleteCaRequest instantiates a new DeleteCaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCaRequestWithDefaults

`func NewDeleteCaRequestWithDefaults() *DeleteCaRequest`

NewDeleteCaRequestWithDefaults instantiates a new DeleteCaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaId

`func (o *DeleteCaRequest) GetCaId() string`

GetCaId returns the CaId field if non-nil, zero value otherwise.

### GetCaIdOk

`func (o *DeleteCaRequest) GetCaIdOk() (*string, bool)`

GetCaIdOk returns a tuple with the CaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaId

`func (o *DeleteCaRequest) SetCaId(v string)`

SetCaId sets CaId field to given value.


### GetDryRun

`func (o *DeleteCaRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteCaRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteCaRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteCaRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


