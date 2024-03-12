# CatalogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the catalog entry (for example, &#x60;network&#x60;). | [optional] 
**Flags** | Pointer to **string** | When returned and equal to &#x60;PER_MONTH&#x60;, the price of the catalog entry is calculated on a monthly basis. | [optional] 
**Operation** | Pointer to **string** | The API call associated with the catalog entry (for example, &#x60;CreateVms&#x60; or &#x60;RunInstances&#x60;). | [optional] 
**Service** | Pointer to **string** | The service associated with the catalog entry (&#x60;TinaOS-FCU&#x60;, &#x60;TinaOS-LBU&#x60;, &#x60;TinaOS-DirectLink&#x60;, or &#x60;TinaOS-OOS&#x60;). | [optional] 
**SubregionName** | Pointer to **string** | The Subregion associated with the catalog entry. | [optional] 
**Title** | Pointer to **string** | The description of the catalog entry. | [optional] 
**Type** | Pointer to **string** | The type of resource associated with the catalog entry. | [optional] 
**UnitPrice** | Pointer to **float32** | The unit price of the catalog entry, in the currency of the Region&#39;s catalog. | [optional] 

## Methods

### NewCatalogEntry

`func NewCatalogEntry() *CatalogEntry`

NewCatalogEntry instantiates a new CatalogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogEntryWithDefaults

`func NewCatalogEntryWithDefaults() *CatalogEntry`

NewCatalogEntryWithDefaults instantiates a new CatalogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CatalogEntry) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CatalogEntry) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CatalogEntry) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CatalogEntry) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFlags

`func (o *CatalogEntry) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *CatalogEntry) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *CatalogEntry) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *CatalogEntry) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetOperation

`func (o *CatalogEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CatalogEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CatalogEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *CatalogEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetService

`func (o *CatalogEntry) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CatalogEntry) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CatalogEntry) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CatalogEntry) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSubregionName

`func (o *CatalogEntry) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *CatalogEntry) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *CatalogEntry) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *CatalogEntry) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### GetTitle

`func (o *CatalogEntry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogEntry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogEntry) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogEntry) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CatalogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnitPrice

`func (o *CatalogEntry) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CatalogEntry) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CatalogEntry) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *CatalogEntry) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


