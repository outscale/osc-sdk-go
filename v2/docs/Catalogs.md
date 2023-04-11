# Catalogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]CatalogEntry**](CatalogEntry.md) | One or more catalog entries. | [optional] 
**FromDate** | Pointer to **time.Time** | The beginning of the time period, in ISO 8601 date-time format. | [optional] 
**State** | Pointer to **string** | The state of the catalog (&#x60;CURRENT&#x60; \\| &#x60;OBSOLETE&#x60;). | [optional] 
**ToDate** | Pointer to **time.Time** | The end of the time period, in ISO 8601 date-time format. | [optional] 

## Methods

### NewCatalogs

`func NewCatalogs() *Catalogs`

NewCatalogs instantiates a new Catalogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogsWithDefaults

`func NewCatalogsWithDefaults() *Catalogs`

NewCatalogsWithDefaults instantiates a new Catalogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *Catalogs) GetEntries() []CatalogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *Catalogs) GetEntriesOk() (*[]CatalogEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *Catalogs) SetEntries(v []CatalogEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *Catalogs) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetFromDate

`func (o *Catalogs) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *Catalogs) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *Catalogs) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *Catalogs) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetState

`func (o *Catalogs) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Catalogs) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Catalogs) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Catalogs) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToDate

`func (o *Catalogs) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *Catalogs) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *Catalogs) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *Catalogs) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


