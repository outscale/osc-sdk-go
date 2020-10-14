# FiltersQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]string** | The group names of the quotas. | [optional] 
**QuotaNames** | Pointer to **[]string** | The names of the quotas. | [optional] 
**QuotaTypes** | Pointer to **[]string** | The resource IDs if these are resource-specific quotas, &#x60;global&#x60; if they are not. | [optional] 
**ShortDescriptions** | Pointer to **[]string** | The description of the quotas. | [optional] 

## Methods

### NewFiltersQuota

`func NewFiltersQuota() *FiltersQuota`

NewFiltersQuota instantiates a new FiltersQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersQuotaWithDefaults

`func NewFiltersQuotaWithDefaults() *FiltersQuota`

NewFiltersQuotaWithDefaults instantiates a new FiltersQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *FiltersQuota) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *FiltersQuota) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *FiltersQuota) SetCollections(v []string)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *FiltersQuota) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetQuotaNames

`func (o *FiltersQuota) GetQuotaNames() []string`

GetQuotaNames returns the QuotaNames field if non-nil, zero value otherwise.

### GetQuotaNamesOk

`func (o *FiltersQuota) GetQuotaNamesOk() (*[]string, bool)`

GetQuotaNamesOk returns a tuple with the QuotaNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaNames

`func (o *FiltersQuota) SetQuotaNames(v []string)`

SetQuotaNames sets QuotaNames field to given value.

### HasQuotaNames

`func (o *FiltersQuota) HasQuotaNames() bool`

HasQuotaNames returns a boolean if a field has been set.

### GetQuotaTypes

`func (o *FiltersQuota) GetQuotaTypes() []string`

GetQuotaTypes returns the QuotaTypes field if non-nil, zero value otherwise.

### GetQuotaTypesOk

`func (o *FiltersQuota) GetQuotaTypesOk() (*[]string, bool)`

GetQuotaTypesOk returns a tuple with the QuotaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaTypes

`func (o *FiltersQuota) SetQuotaTypes(v []string)`

SetQuotaTypes sets QuotaTypes field to given value.

### HasQuotaTypes

`func (o *FiltersQuota) HasQuotaTypes() bool`

HasQuotaTypes returns a boolean if a field has been set.

### GetShortDescriptions

`func (o *FiltersQuota) GetShortDescriptions() []string`

GetShortDescriptions returns the ShortDescriptions field if non-nil, zero value otherwise.

### GetShortDescriptionsOk

`func (o *FiltersQuota) GetShortDescriptionsOk() (*[]string, bool)`

GetShortDescriptionsOk returns a tuple with the ShortDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescriptions

`func (o *FiltersQuota) SetShortDescriptions(v []string)`

SetShortDescriptions sets ShortDescriptions field to given value.

### HasShortDescriptions

`func (o *FiltersQuota) HasShortDescriptions() bool`

HasShortDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


