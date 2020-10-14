# QuotaTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaType** | Pointer to **string** | The resource ID if it is a resource-specific quota, &#x60;global&#x60; if it is not. | [optional] 
**Quotas** | Pointer to [**[]Quota**](Quota.md) | One or more quotas associated with the user. | [optional] 

## Methods

### NewQuotaTypes

`func NewQuotaTypes() *QuotaTypes`

NewQuotaTypes instantiates a new QuotaTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaTypesWithDefaults

`func NewQuotaTypesWithDefaults() *QuotaTypes`

NewQuotaTypesWithDefaults instantiates a new QuotaTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotaType

`func (o *QuotaTypes) GetQuotaType() string`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *QuotaTypes) GetQuotaTypeOk() (*string, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaType

`func (o *QuotaTypes) SetQuotaType(v string)`

SetQuotaType sets QuotaType field to given value.

### HasQuotaType

`func (o *QuotaTypes) HasQuotaType() bool`

HasQuotaType returns a boolean if a field has been set.

### GetQuotas

`func (o *QuotaTypes) GetQuotas() []Quota`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *QuotaTypes) GetQuotasOk() (*[]Quota, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *QuotaTypes) SetQuotas(v []Quota)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *QuotaTypes) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


