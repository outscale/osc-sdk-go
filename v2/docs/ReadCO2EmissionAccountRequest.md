# ReadCO2EmissionAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromMonth** | **string** | The beginning of the time period, in ISO 8601 date format (for example, &#x60;2020-06-01&#x60;). This value must correspond to the first day of the month and is included in the time period. | 
**Overall** | Pointer to **bool** | If false, returns only the CO2 emission of the specific account that sends the request. If true, returns either the overall CO2 emission of your paying account and all linked accounts (if the account that sends this request is a paying account) or returns nothing (if the account that sends this request is a linked account). | [optional] [default to false]
**ToMonth** | **string** | The end of the time period, in ISO 8601 date format (for example, &#x60;2020-06-14&#x60;). This value must correspond to the first day of the month and is excluded from the time period. It must be set to a later date than &#x60;FromMonth&#x60;. | 

## Methods

### NewReadCO2EmissionAccountRequest

`func NewReadCO2EmissionAccountRequest(fromMonth string, toMonth string, ) *ReadCO2EmissionAccountRequest`

NewReadCO2EmissionAccountRequest instantiates a new ReadCO2EmissionAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadCO2EmissionAccountRequestWithDefaults

`func NewReadCO2EmissionAccountRequestWithDefaults() *ReadCO2EmissionAccountRequest`

NewReadCO2EmissionAccountRequestWithDefaults instantiates a new ReadCO2EmissionAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromMonth

`func (o *ReadCO2EmissionAccountRequest) GetFromMonth() string`

GetFromMonth returns the FromMonth field if non-nil, zero value otherwise.

### GetFromMonthOk

`func (o *ReadCO2EmissionAccountRequest) GetFromMonthOk() (*string, bool)`

GetFromMonthOk returns a tuple with the FromMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMonth

`func (o *ReadCO2EmissionAccountRequest) SetFromMonth(v string)`

SetFromMonth sets FromMonth field to given value.


### GetOverall

`func (o *ReadCO2EmissionAccountRequest) GetOverall() bool`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *ReadCO2EmissionAccountRequest) GetOverallOk() (*bool, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *ReadCO2EmissionAccountRequest) SetOverall(v bool)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *ReadCO2EmissionAccountRequest) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetToMonth

`func (o *ReadCO2EmissionAccountRequest) GetToMonth() string`

GetToMonth returns the ToMonth field if non-nil, zero value otherwise.

### GetToMonthOk

`func (o *ReadCO2EmissionAccountRequest) GetToMonthOk() (*string, bool)`

GetToMonthOk returns a tuple with the ToMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToMonth

`func (o *ReadCO2EmissionAccountRequest) SetToMonth(v string)`

SetToMonth sets ToMonth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


