# UndeleteMasterKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**MasterKeyId** | **string** | The ID of the master key. | 

## Methods

### NewUndeleteMasterKeyRequest

`func NewUndeleteMasterKeyRequest(masterKeyId string, ) *UndeleteMasterKeyRequest`

NewUndeleteMasterKeyRequest instantiates a new UndeleteMasterKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUndeleteMasterKeyRequestWithDefaults

`func NewUndeleteMasterKeyRequestWithDefaults() *UndeleteMasterKeyRequest`

NewUndeleteMasterKeyRequestWithDefaults instantiates a new UndeleteMasterKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UndeleteMasterKeyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UndeleteMasterKeyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UndeleteMasterKeyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UndeleteMasterKeyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMasterKeyId

`func (o *UndeleteMasterKeyRequest) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *UndeleteMasterKeyRequest) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *UndeleteMasterKeyRequest) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


