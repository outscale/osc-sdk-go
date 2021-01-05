# UpdateMasterKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The new description for the master key, between 0 and 8192 Unicode characters. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Enabled** | Pointer to **bool** | If &#x60;true&#x60;, the state of the master key becomes &#x60;Enabled&#x60;. If &#x60;false&#x60;, the state of the master key becomes &#x60;Disabled&#x60;. | [optional] 
**MasterKeyId** | **string** | The ID of the master key. | 

## Methods

### NewUpdateMasterKeyRequest

`func NewUpdateMasterKeyRequest(masterKeyId string, ) *UpdateMasterKeyRequest`

NewUpdateMasterKeyRequest instantiates a new UpdateMasterKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMasterKeyRequestWithDefaults

`func NewUpdateMasterKeyRequestWithDefaults() *UpdateMasterKeyRequest`

NewUpdateMasterKeyRequestWithDefaults instantiates a new UpdateMasterKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateMasterKeyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMasterKeyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMasterKeyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMasterKeyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateMasterKeyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateMasterKeyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateMasterKeyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateMasterKeyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateMasterKeyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateMasterKeyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateMasterKeyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateMasterKeyRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMasterKeyId

`func (o *UpdateMasterKeyRequest) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *UpdateMasterKeyRequest) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *UpdateMasterKeyRequest) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


