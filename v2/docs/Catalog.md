# Catalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]CatalogEntry**](CatalogEntry.md) | One or more catalog entries. | [optional] 

## Methods

### NewCatalog

`func NewCatalog() *Catalog`

NewCatalog instantiates a new Catalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogWithDefaults

`func NewCatalogWithDefaults() *Catalog`

NewCatalogWithDefaults instantiates a new Catalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *Catalog) GetEntries() []CatalogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *Catalog) GetEntriesOk() (*[]CatalogEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *Catalog) SetEntries(v []CatalogEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *Catalog) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


