# ReadConsumptionAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumptionEntries** | Pointer to [**[]ConsumptionEntry**](ConsumptionEntry.md) | Information about the resources consumed during the specified time period. | [optional] 
**Currency** | Pointer to **NullableString** | The currency of your account for the &#x60;UnitPrice&#x60; and &#x60;Price&#x60; parameters, in the ISO-4217 format (for example, &#x60;EUR&#x60;). | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadConsumptionAccountResponse

`func NewReadConsumptionAccountResponse() *ReadConsumptionAccountResponse`

NewReadConsumptionAccountResponse instantiates a new ReadConsumptionAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadConsumptionAccountResponseWithDefaults

`func NewReadConsumptionAccountResponseWithDefaults() *ReadConsumptionAccountResponse`

NewReadConsumptionAccountResponseWithDefaults instantiates a new ReadConsumptionAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumptionEntries

`func (o *ReadConsumptionAccountResponse) GetConsumptionEntries() []ConsumptionEntry`

GetConsumptionEntries returns the ConsumptionEntries field if non-nil, zero value otherwise.

### GetConsumptionEntriesOk

`func (o *ReadConsumptionAccountResponse) GetConsumptionEntriesOk() (*[]ConsumptionEntry, bool)`

GetConsumptionEntriesOk returns a tuple with the ConsumptionEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionEntries

`func (o *ReadConsumptionAccountResponse) SetConsumptionEntries(v []ConsumptionEntry)`

SetConsumptionEntries sets ConsumptionEntries field to given value.

### HasConsumptionEntries

`func (o *ReadConsumptionAccountResponse) HasConsumptionEntries() bool`

HasConsumptionEntries returns a boolean if a field has been set.

### GetCurrency

`func (o *ReadConsumptionAccountResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReadConsumptionAccountResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReadConsumptionAccountResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ReadConsumptionAccountResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ReadConsumptionAccountResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ReadConsumptionAccountResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetResponseContext

`func (o *ReadConsumptionAccountResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadConsumptionAccountResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadConsumptionAccountResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadConsumptionAccountResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


