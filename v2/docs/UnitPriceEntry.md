# UnitPriceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of your account for the &#x60;UnitPrice&#x60; parameter, in the ISO-4217 format (for example, &#x60;EUR&#x60;). | [optional] 
**Operation** | Pointer to **string** | The operation associated with the catalog entry (for example, &#x60;RunInstances-OD&#x60; or &#x60;CreateVolume&#x60;). | [optional] 
**Service** | Pointer to **string** | The service associated with the catalog entry (for example, &#x60;TinaOS-FCU&#x60; or &#x60;TinaOS-OOS&#x60;). | [optional] 
**Type** | Pointer to **string** | The type associated with the catalog entry (for example, &#x60;BSU:VolumeIOPS:io1&#x60; or &#x60;BoxUsage:tinav6.c6r16p3&#x60;). | [optional] 
**Unit** | Pointer to **string** | The unit associated with the catalog entry (for example, &#x60;PER_MONTH&#x60; or &#x60;PER_COUNT&#x60;). | [optional] 
**UnitPrice** | Pointer to **float64** | The unit price of the catalog entry in the currency of your account, in the ISO-4217 format (for example, &#x60;EUR&#x60;). | [optional] 

## Methods

### NewUnitPriceEntry

`func NewUnitPriceEntry() *UnitPriceEntry`

NewUnitPriceEntry instantiates a new UnitPriceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitPriceEntryWithDefaults

`func NewUnitPriceEntryWithDefaults() *UnitPriceEntry`

NewUnitPriceEntryWithDefaults instantiates a new UnitPriceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnitPriceEntry) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnitPriceEntry) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnitPriceEntry) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnitPriceEntry) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOperation

`func (o *UnitPriceEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *UnitPriceEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *UnitPriceEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *UnitPriceEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetService

`func (o *UnitPriceEntry) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *UnitPriceEntry) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *UnitPriceEntry) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *UnitPriceEntry) HasService() bool`

HasService returns a boolean if a field has been set.

### GetType

`func (o *UnitPriceEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnitPriceEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnitPriceEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UnitPriceEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *UnitPriceEntry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UnitPriceEntry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UnitPriceEntry) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UnitPriceEntry) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPrice

`func (o *UnitPriceEntry) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *UnitPriceEntry) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *UnitPriceEntry) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *UnitPriceEntry) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


