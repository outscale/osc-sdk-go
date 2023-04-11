# FiltersCatalogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCatalogOnly** | Pointer to **bool** | By default or if set to true, only returns the current catalog. If false, returns the current catalog and past catalogs. | [optional] 
**FromDate** | Pointer to **string** | The beginning of the time period, in ISO 8601 date format (for example, &#x60;2020-06-14&#x60;). This date cannot be older than 3 years. You must specify the parameters &#x60;FromDate&#x60; and &#x60;ToDate&#x60; together. | [optional] 
**ToDate** | Pointer to **string** | The end of the time period, in ISO 8601 date format (for example, &#x60;2020-06-30&#x60;). You must specify the parameters &#x60;FromDate&#x60; and &#x60;ToDate&#x60; together. | [optional] 

## Methods

### NewFiltersCatalogs

`func NewFiltersCatalogs() *FiltersCatalogs`

NewFiltersCatalogs instantiates a new FiltersCatalogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersCatalogsWithDefaults

`func NewFiltersCatalogsWithDefaults() *FiltersCatalogs`

NewFiltersCatalogsWithDefaults instantiates a new FiltersCatalogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentCatalogOnly

`func (o *FiltersCatalogs) GetCurrentCatalogOnly() bool`

GetCurrentCatalogOnly returns the CurrentCatalogOnly field if non-nil, zero value otherwise.

### GetCurrentCatalogOnlyOk

`func (o *FiltersCatalogs) GetCurrentCatalogOnlyOk() (*bool, bool)`

GetCurrentCatalogOnlyOk returns a tuple with the CurrentCatalogOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCatalogOnly

`func (o *FiltersCatalogs) SetCurrentCatalogOnly(v bool)`

SetCurrentCatalogOnly sets CurrentCatalogOnly field to given value.

### HasCurrentCatalogOnly

`func (o *FiltersCatalogs) HasCurrentCatalogOnly() bool`

HasCurrentCatalogOnly returns a boolean if a field has been set.

### GetFromDate

`func (o *FiltersCatalogs) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *FiltersCatalogs) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *FiltersCatalogs) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *FiltersCatalogs) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *FiltersCatalogs) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *FiltersCatalogs) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *FiltersCatalogs) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *FiltersCatalogs) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


