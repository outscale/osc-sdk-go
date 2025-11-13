# CO2FactorDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Factor** | Pointer to **string** | The emission source (for example, &#x60;hardware&#x60;). | [optional] 
**Value** | Pointer to **float64** | The total CO2 emissions for the factor. | [optional] 

## Methods

### NewCO2FactorDistribution

`func NewCO2FactorDistribution() *CO2FactorDistribution`

NewCO2FactorDistribution instantiates a new CO2FactorDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCO2FactorDistributionWithDefaults

`func NewCO2FactorDistributionWithDefaults() *CO2FactorDistribution`

NewCO2FactorDistributionWithDefaults instantiates a new CO2FactorDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactor

`func (o *CO2FactorDistribution) GetFactor() string`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *CO2FactorDistribution) GetFactorOk() (*string, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *CO2FactorDistribution) SetFactor(v string)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *CO2FactorDistribution) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetValue

`func (o *CO2FactorDistribution) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CO2FactorDistribution) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CO2FactorDistribution) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *CO2FactorDistribution) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


