# ReadConsumptionAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**FromDate** | **string** | The beginning of the time period, in ISO 8601 date format (for example, &#x60;2020-06-14&#x60;). The date-time format is also accepted, but only with a time set to midnight (for example, &#x60;2020-06-14T00:00:00.000Z&#x60;). | 
**Overall** | Pointer to **bool** | By default or if false, returns only the consumption of the specific account that sends this request. If true, returns either the overall consumption of your paying account and all linked accounts (if the account that sends this request is a paying account) or returns nothing (if the account that sends this request is a linked account). | [optional] [default to false]
**ShowPrice** | Pointer to **bool** | If true, the response also includes the unit price of the consumed resource (&#x60;UnitPrice&#x60;) and the total price of the consumed resource during the specified time period (&#x60;Price&#x60;), in the currency of the Region&#39;s catalog. | [optional] 
**ToDate** | **string** | The end of the time period, in ISO 8601 date format (for example, &#x60;2020-06-30&#x60;). The date-time format is also accepted, but only with a time set to midnight (for example, &#x60;2020-06-30T00:00:00.000Z&#x60;). | 

## Methods

### NewReadConsumptionAccountRequest

`func NewReadConsumptionAccountRequest(fromDate string, toDate string, ) *ReadConsumptionAccountRequest`

NewReadConsumptionAccountRequest instantiates a new ReadConsumptionAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadConsumptionAccountRequestWithDefaults

`func NewReadConsumptionAccountRequestWithDefaults() *ReadConsumptionAccountRequest`

NewReadConsumptionAccountRequestWithDefaults instantiates a new ReadConsumptionAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadConsumptionAccountRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadConsumptionAccountRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadConsumptionAccountRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadConsumptionAccountRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFromDate

`func (o *ReadConsumptionAccountRequest) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *ReadConsumptionAccountRequest) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *ReadConsumptionAccountRequest) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.


### GetOverall

`func (o *ReadConsumptionAccountRequest) GetOverall() bool`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *ReadConsumptionAccountRequest) GetOverallOk() (*bool, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *ReadConsumptionAccountRequest) SetOverall(v bool)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *ReadConsumptionAccountRequest) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetShowPrice

`func (o *ReadConsumptionAccountRequest) GetShowPrice() bool`

GetShowPrice returns the ShowPrice field if non-nil, zero value otherwise.

### GetShowPriceOk

`func (o *ReadConsumptionAccountRequest) GetShowPriceOk() (*bool, bool)`

GetShowPriceOk returns a tuple with the ShowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPrice

`func (o *ReadConsumptionAccountRequest) SetShowPrice(v bool)`

SetShowPrice sets ShowPrice field to given value.

### HasShowPrice

`func (o *ReadConsumptionAccountRequest) HasShowPrice() bool`

HasShowPrice returns a boolean if a field has been set.

### GetToDate

`func (o *ReadConsumptionAccountRequest) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *ReadConsumptionAccountRequest) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *ReadConsumptionAccountRequest) SetToDate(v string)`

SetToDate sets ToDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


