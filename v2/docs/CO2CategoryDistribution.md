# CO2CategoryDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the resource (for example, &#x60;storage&#x60;). | [optional] 
**Value** | Pointer to **float64** | The total CO2 emissions for the category. | [optional] 

## Methods

### NewCO2CategoryDistribution

`func NewCO2CategoryDistribution() *CO2CategoryDistribution`

NewCO2CategoryDistribution instantiates a new CO2CategoryDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCO2CategoryDistributionWithDefaults

`func NewCO2CategoryDistributionWithDefaults() *CO2CategoryDistribution`

NewCO2CategoryDistributionWithDefaults instantiates a new CO2CategoryDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CO2CategoryDistribution) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CO2CategoryDistribution) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CO2CategoryDistribution) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CO2CategoryDistribution) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetValue

`func (o *CO2CategoryDistribution) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CO2CategoryDistribution) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CO2CategoryDistribution) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *CO2CategoryDistribution) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


