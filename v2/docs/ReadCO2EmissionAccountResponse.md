# ReadCO2EmissionAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CO2EmissionEntries** | Pointer to [**[]CO2EmissionEntry**](CO2EmissionEntry.md) | The CO2 emission by month and account, for the specified request. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Unit** | Pointer to **string** | The unit of all the &#x60;Value&#x60; fields of the response, expressed in kgCO₂e. | [optional] 
**Value** | Pointer to **float64** | The total CO2 emission for the specified request. | [optional] 

## Methods

### NewReadCO2EmissionAccountResponse

`func NewReadCO2EmissionAccountResponse() *ReadCO2EmissionAccountResponse`

NewReadCO2EmissionAccountResponse instantiates a new ReadCO2EmissionAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadCO2EmissionAccountResponseWithDefaults

`func NewReadCO2EmissionAccountResponseWithDefaults() *ReadCO2EmissionAccountResponse`

NewReadCO2EmissionAccountResponseWithDefaults instantiates a new ReadCO2EmissionAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCO2EmissionEntries

`func (o *ReadCO2EmissionAccountResponse) GetCO2EmissionEntries() []CO2EmissionEntry`

GetCO2EmissionEntries returns the CO2EmissionEntries field if non-nil, zero value otherwise.

### GetCO2EmissionEntriesOk

`func (o *ReadCO2EmissionAccountResponse) GetCO2EmissionEntriesOk() (*[]CO2EmissionEntry, bool)`

GetCO2EmissionEntriesOk returns a tuple with the CO2EmissionEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCO2EmissionEntries

`func (o *ReadCO2EmissionAccountResponse) SetCO2EmissionEntries(v []CO2EmissionEntry)`

SetCO2EmissionEntries sets CO2EmissionEntries field to given value.

### HasCO2EmissionEntries

`func (o *ReadCO2EmissionAccountResponse) HasCO2EmissionEntries() bool`

HasCO2EmissionEntries returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadCO2EmissionAccountResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadCO2EmissionAccountResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadCO2EmissionAccountResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadCO2EmissionAccountResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetUnit

`func (o *ReadCO2EmissionAccountResponse) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ReadCO2EmissionAccountResponse) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ReadCO2EmissionAccountResponse) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ReadCO2EmissionAccountResponse) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *ReadCO2EmissionAccountResponse) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReadCO2EmissionAccountResponse) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReadCO2EmissionAccountResponse) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ReadCO2EmissionAccountResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


