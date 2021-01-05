# DeleteMasterKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysUntilDeletion** | Pointer to **int32** | The waiting period before deletion, in days (between &#x60;7&#x60; and &#x60;30&#x60;). By default, &#x60;30&#x60;. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**MasterKeyId** | **string** | The ID of the master key. | 

## Methods

### NewDeleteMasterKeyRequest

`func NewDeleteMasterKeyRequest(masterKeyId string, ) *DeleteMasterKeyRequest`

NewDeleteMasterKeyRequest instantiates a new DeleteMasterKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMasterKeyRequestWithDefaults

`func NewDeleteMasterKeyRequestWithDefaults() *DeleteMasterKeyRequest`

NewDeleteMasterKeyRequestWithDefaults instantiates a new DeleteMasterKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysUntilDeletion

`func (o *DeleteMasterKeyRequest) GetDaysUntilDeletion() int32`

GetDaysUntilDeletion returns the DaysUntilDeletion field if non-nil, zero value otherwise.

### GetDaysUntilDeletionOk

`func (o *DeleteMasterKeyRequest) GetDaysUntilDeletionOk() (*int32, bool)`

GetDaysUntilDeletionOk returns a tuple with the DaysUntilDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilDeletion

`func (o *DeleteMasterKeyRequest) SetDaysUntilDeletion(v int32)`

SetDaysUntilDeletion sets DaysUntilDeletion field to given value.

### HasDaysUntilDeletion

`func (o *DeleteMasterKeyRequest) HasDaysUntilDeletion() bool`

HasDaysUntilDeletion returns a boolean if a field has been set.

### GetDryRun

`func (o *DeleteMasterKeyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteMasterKeyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteMasterKeyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteMasterKeyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMasterKeyId

`func (o *DeleteMasterKeyRequest) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *DeleteMasterKeyRequest) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *DeleteMasterKeyRequest) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


