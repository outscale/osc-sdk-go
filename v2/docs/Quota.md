# Quota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the quotas. | [optional] 
**Description** | Pointer to **string** | The description of the quota. | [optional] 
**MaxValue** | Pointer to **int32** | The maximum value of the quota for the OUTSCALE user account (if there is no limit, &#x60;0&#x60;). | [optional] 
**Name** | Pointer to **string** | The unique name of the quota. | [optional] 
**QuotaCollection** | Pointer to **string** | The group name of the quota. | [optional] 
**ShortDescription** | Pointer to **string** | The description of the quota. | [optional] 
**UsedValue** | Pointer to **int32** | The limit value currently used by the OUTSCALE user account. | [optional] 

## Methods

### NewQuota

`func NewQuota() *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Quota) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Quota) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Quota) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Quota) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDescription

`func (o *Quota) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Quota) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Quota) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Quota) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxValue

`func (o *Quota) GetMaxValue() int32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *Quota) GetMaxValueOk() (*int32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *Quota) SetMaxValue(v int32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *Quota) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetName

`func (o *Quota) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Quota) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Quota) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Quota) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuotaCollection

`func (o *Quota) GetQuotaCollection() string`

GetQuotaCollection returns the QuotaCollection field if non-nil, zero value otherwise.

### GetQuotaCollectionOk

`func (o *Quota) GetQuotaCollectionOk() (*string, bool)`

GetQuotaCollectionOk returns a tuple with the QuotaCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaCollection

`func (o *Quota) SetQuotaCollection(v string)`

SetQuotaCollection sets QuotaCollection field to given value.

### HasQuotaCollection

`func (o *Quota) HasQuotaCollection() bool`

HasQuotaCollection returns a boolean if a field has been set.

### GetShortDescription

`func (o *Quota) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Quota) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Quota) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *Quota) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetUsedValue

`func (o *Quota) GetUsedValue() int32`

GetUsedValue returns the UsedValue field if non-nil, zero value otherwise.

### GetUsedValueOk

`func (o *Quota) GetUsedValueOk() (*int32, bool)`

GetUsedValueOk returns a tuple with the UsedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedValue

`func (o *Quota) SetUsedValue(v int32)`

SetUsedValue sets UsedValue field to given value.

### HasUsedValue

`func (o *Quota) HasUsedValue() bool`

HasUsedValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


