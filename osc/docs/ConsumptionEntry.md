# ConsumptionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the resource (for example, &#x60;network&#x60;). | [optional] 
**FromDate** | Pointer to **string** | The beginning of the time period. | [optional] 
**Operation** | Pointer to **string** | The API call that triggered the resource consumption (for example, &#x60;RunInstances&#x60; or &#x60;CreateVolume&#x60;). | [optional] 
**Service** | Pointer to **string** | The service of the API call (&#x60;TinaOS-FCU&#x60;, &#x60;TinaOS-LBU&#x60;, &#x60;TinaOS-OSU&#x60; or &#x60;TinaOS-DirectLink&#x60;). | [optional] 
**Title** | Pointer to **string** | A description of the consumed resource. | [optional] 
**ToDate** | Pointer to **string** | The end of the time period. | [optional] 
**Type** | Pointer to **string** | The type of resource, depending on the API call. | [optional] 
**Value** | Pointer to **string** | The consumed amount for the resource. The unit depends on the resource type. For more information, see the &#x60;Title&#x60; element. | [optional] 

## Methods

### NewConsumptionEntry

`func NewConsumptionEntry() *ConsumptionEntry`

NewConsumptionEntry instantiates a new ConsumptionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumptionEntryWithDefaults

`func NewConsumptionEntryWithDefaults() *ConsumptionEntry`

NewConsumptionEntryWithDefaults instantiates a new ConsumptionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ConsumptionEntry) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConsumptionEntry) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConsumptionEntry) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ConsumptionEntry) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFromDate

`func (o *ConsumptionEntry) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *ConsumptionEntry) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *ConsumptionEntry) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *ConsumptionEntry) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetOperation

`func (o *ConsumptionEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ConsumptionEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ConsumptionEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ConsumptionEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetService

`func (o *ConsumptionEntry) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ConsumptionEntry) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ConsumptionEntry) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ConsumptionEntry) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTitle

`func (o *ConsumptionEntry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConsumptionEntry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConsumptionEntry) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConsumptionEntry) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetToDate

`func (o *ConsumptionEntry) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *ConsumptionEntry) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *ConsumptionEntry) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *ConsumptionEntry) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetType

`func (o *ConsumptionEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumptionEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumptionEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConsumptionEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ConsumptionEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConsumptionEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConsumptionEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConsumptionEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


