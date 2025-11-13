# CO2EmissionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The ID of the account associated with the consumption. | [optional] 
**CategoryDistribution** | Pointer to [**[]CO2CategoryDistribution**](CO2CategoryDistribution.md) | The allocation of the &#x60;Value&#x60; among categories. | [optional] 
**FactorDistribution** | Pointer to [**[]CO2FactorDistribution**](CO2FactorDistribution.md) | The allocation of the &#x60;Value&#x60; among factors. | [optional] 
**Month** | Pointer to **string** | The month associated with the CO2 emission entry. | [optional] 
**PayingAccountId** | Pointer to **string** | The ID of the paying account related to the &#x60;AccountId&#x60; parameter. | [optional] 
**Value** | Pointer to **float64** | The total CO2 emissions for the &#x60;Month&#x60; and &#x60;AccountId&#x60; specified. This value corresponds to the sum of all entries in &#x60;CategoryDistribution&#x60; and &#x60;FactorDistributionEntry&#x60;. | [optional] 

## Methods

### NewCO2EmissionEntry

`func NewCO2EmissionEntry() *CO2EmissionEntry`

NewCO2EmissionEntry instantiates a new CO2EmissionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCO2EmissionEntryWithDefaults

`func NewCO2EmissionEntryWithDefaults() *CO2EmissionEntry`

NewCO2EmissionEntryWithDefaults instantiates a new CO2EmissionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CO2EmissionEntry) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CO2EmissionEntry) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CO2EmissionEntry) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CO2EmissionEntry) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCategoryDistribution

`func (o *CO2EmissionEntry) GetCategoryDistribution() []CO2CategoryDistribution`

GetCategoryDistribution returns the CategoryDistribution field if non-nil, zero value otherwise.

### GetCategoryDistributionOk

`func (o *CO2EmissionEntry) GetCategoryDistributionOk() (*[]CO2CategoryDistribution, bool)`

GetCategoryDistributionOk returns a tuple with the CategoryDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryDistribution

`func (o *CO2EmissionEntry) SetCategoryDistribution(v []CO2CategoryDistribution)`

SetCategoryDistribution sets CategoryDistribution field to given value.

### HasCategoryDistribution

`func (o *CO2EmissionEntry) HasCategoryDistribution() bool`

HasCategoryDistribution returns a boolean if a field has been set.

### GetFactorDistribution

`func (o *CO2EmissionEntry) GetFactorDistribution() []CO2FactorDistribution`

GetFactorDistribution returns the FactorDistribution field if non-nil, zero value otherwise.

### GetFactorDistributionOk

`func (o *CO2EmissionEntry) GetFactorDistributionOk() (*[]CO2FactorDistribution, bool)`

GetFactorDistributionOk returns a tuple with the FactorDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorDistribution

`func (o *CO2EmissionEntry) SetFactorDistribution(v []CO2FactorDistribution)`

SetFactorDistribution sets FactorDistribution field to given value.

### HasFactorDistribution

`func (o *CO2EmissionEntry) HasFactorDistribution() bool`

HasFactorDistribution returns a boolean if a field has been set.

### GetMonth

`func (o *CO2EmissionEntry) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *CO2EmissionEntry) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *CO2EmissionEntry) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *CO2EmissionEntry) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetPayingAccountId

`func (o *CO2EmissionEntry) GetPayingAccountId() string`

GetPayingAccountId returns the PayingAccountId field if non-nil, zero value otherwise.

### GetPayingAccountIdOk

`func (o *CO2EmissionEntry) GetPayingAccountIdOk() (*string, bool)`

GetPayingAccountIdOk returns a tuple with the PayingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayingAccountId

`func (o *CO2EmissionEntry) SetPayingAccountId(v string)`

SetPayingAccountId sets PayingAccountId field to given value.

### HasPayingAccountId

`func (o *CO2EmissionEntry) HasPayingAccountId() bool`

HasPayingAccountId returns a boolean if a field has been set.

### GetValue

`func (o *CO2EmissionEntry) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CO2EmissionEntry) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CO2EmissionEntry) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *CO2EmissionEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


