# FiltersApiLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryAccessKeys** | Pointer to **[]string** | The access keys used for the logged calls. | [optional] 
**QueryApiNames** | Pointer to **[]string** | The names of the APIs of the logged calls (always &#x60;oapi&#x60; for the OUTSCALE API). | [optional] 
**QueryCallNames** | Pointer to **[]string** | The names of the logged calls. | [optional] 
**QueryDateAfter** | Pointer to **time.Time** | The date and time, or the date, after which you want to retrieve logged calls, in ISO 8601 format (for example, &#x60;2020-06-14T00:00:00.000Z&#x60; or &#x60;2020-06-14&#x60;). By default, this date is set to 48 hours before the &#x60;QueryDateBefore&#x60; parameter value. | [optional] 
**QueryDateBefore** | Pointer to **time.Time** | The date and time, or the date, before which you want to retrieve logged calls, in ISO 8601 format (for example, &#x60;2020-06-30T00:00:00.000Z&#x60; or &#x60;2020-06-14&#x60;). By default, this date is set to now, or 48 hours after the &#x60;QueryDateAfter&#x60; parameter value. | [optional] 
**QueryIpAddresses** | Pointer to **[]string** | The IPs used for the logged calls. | [optional] 
**QueryUserAgents** | Pointer to **[]string** | The user agents of the HTTP requests of the logged calls. | [optional] 
**RequestIds** | Pointer to **[]string** | The request IDs provided in the responses of the logged calls. | [optional] 
**ResponseStatusCodes** | Pointer to **[]int32** | The HTTP status codes of the logged calls. | [optional] 

## Methods

### NewFiltersApiLog

`func NewFiltersApiLog() *FiltersApiLog`

NewFiltersApiLog instantiates a new FiltersApiLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersApiLogWithDefaults

`func NewFiltersApiLogWithDefaults() *FiltersApiLog`

NewFiltersApiLogWithDefaults instantiates a new FiltersApiLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryAccessKeys

`func (o *FiltersApiLog) GetQueryAccessKeys() []string`

GetQueryAccessKeys returns the QueryAccessKeys field if non-nil, zero value otherwise.

### GetQueryAccessKeysOk

`func (o *FiltersApiLog) GetQueryAccessKeysOk() (*[]string, bool)`

GetQueryAccessKeysOk returns a tuple with the QueryAccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAccessKeys

`func (o *FiltersApiLog) SetQueryAccessKeys(v []string)`

SetQueryAccessKeys sets QueryAccessKeys field to given value.

### HasQueryAccessKeys

`func (o *FiltersApiLog) HasQueryAccessKeys() bool`

HasQueryAccessKeys returns a boolean if a field has been set.

### GetQueryApiNames

`func (o *FiltersApiLog) GetQueryApiNames() []string`

GetQueryApiNames returns the QueryApiNames field if non-nil, zero value otherwise.

### GetQueryApiNamesOk

`func (o *FiltersApiLog) GetQueryApiNamesOk() (*[]string, bool)`

GetQueryApiNamesOk returns a tuple with the QueryApiNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryApiNames

`func (o *FiltersApiLog) SetQueryApiNames(v []string)`

SetQueryApiNames sets QueryApiNames field to given value.

### HasQueryApiNames

`func (o *FiltersApiLog) HasQueryApiNames() bool`

HasQueryApiNames returns a boolean if a field has been set.

### GetQueryCallNames

`func (o *FiltersApiLog) GetQueryCallNames() []string`

GetQueryCallNames returns the QueryCallNames field if non-nil, zero value otherwise.

### GetQueryCallNamesOk

`func (o *FiltersApiLog) GetQueryCallNamesOk() (*[]string, bool)`

GetQueryCallNamesOk returns a tuple with the QueryCallNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCallNames

`func (o *FiltersApiLog) SetQueryCallNames(v []string)`

SetQueryCallNames sets QueryCallNames field to given value.

### HasQueryCallNames

`func (o *FiltersApiLog) HasQueryCallNames() bool`

HasQueryCallNames returns a boolean if a field has been set.

### GetQueryDateAfter

`func (o *FiltersApiLog) GetQueryDateAfter() time.Time`

GetQueryDateAfter returns the QueryDateAfter field if non-nil, zero value otherwise.

### GetQueryDateAfterOk

`func (o *FiltersApiLog) GetQueryDateAfterOk() (*time.Time, bool)`

GetQueryDateAfterOk returns a tuple with the QueryDateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryDateAfter

`func (o *FiltersApiLog) SetQueryDateAfter(v time.Time)`

SetQueryDateAfter sets QueryDateAfter field to given value.

### HasQueryDateAfter

`func (o *FiltersApiLog) HasQueryDateAfter() bool`

HasQueryDateAfter returns a boolean if a field has been set.

### GetQueryDateBefore

`func (o *FiltersApiLog) GetQueryDateBefore() time.Time`

GetQueryDateBefore returns the QueryDateBefore field if non-nil, zero value otherwise.

### GetQueryDateBeforeOk

`func (o *FiltersApiLog) GetQueryDateBeforeOk() (*time.Time, bool)`

GetQueryDateBeforeOk returns a tuple with the QueryDateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryDateBefore

`func (o *FiltersApiLog) SetQueryDateBefore(v time.Time)`

SetQueryDateBefore sets QueryDateBefore field to given value.

### HasQueryDateBefore

`func (o *FiltersApiLog) HasQueryDateBefore() bool`

HasQueryDateBefore returns a boolean if a field has been set.

### GetQueryIpAddresses

`func (o *FiltersApiLog) GetQueryIpAddresses() []string`

GetQueryIpAddresses returns the QueryIpAddresses field if non-nil, zero value otherwise.

### GetQueryIpAddressesOk

`func (o *FiltersApiLog) GetQueryIpAddressesOk() (*[]string, bool)`

GetQueryIpAddressesOk returns a tuple with the QueryIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIpAddresses

`func (o *FiltersApiLog) SetQueryIpAddresses(v []string)`

SetQueryIpAddresses sets QueryIpAddresses field to given value.

### HasQueryIpAddresses

`func (o *FiltersApiLog) HasQueryIpAddresses() bool`

HasQueryIpAddresses returns a boolean if a field has been set.

### GetQueryUserAgents

`func (o *FiltersApiLog) GetQueryUserAgents() []string`

GetQueryUserAgents returns the QueryUserAgents field if non-nil, zero value otherwise.

### GetQueryUserAgentsOk

`func (o *FiltersApiLog) GetQueryUserAgentsOk() (*[]string, bool)`

GetQueryUserAgentsOk returns a tuple with the QueryUserAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryUserAgents

`func (o *FiltersApiLog) SetQueryUserAgents(v []string)`

SetQueryUserAgents sets QueryUserAgents field to given value.

### HasQueryUserAgents

`func (o *FiltersApiLog) HasQueryUserAgents() bool`

HasQueryUserAgents returns a boolean if a field has been set.

### GetRequestIds

`func (o *FiltersApiLog) GetRequestIds() []string`

GetRequestIds returns the RequestIds field if non-nil, zero value otherwise.

### GetRequestIdsOk

`func (o *FiltersApiLog) GetRequestIdsOk() (*[]string, bool)`

GetRequestIdsOk returns a tuple with the RequestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIds

`func (o *FiltersApiLog) SetRequestIds(v []string)`

SetRequestIds sets RequestIds field to given value.

### HasRequestIds

`func (o *FiltersApiLog) HasRequestIds() bool`

HasRequestIds returns a boolean if a field has been set.

### GetResponseStatusCodes

`func (o *FiltersApiLog) GetResponseStatusCodes() []int32`

GetResponseStatusCodes returns the ResponseStatusCodes field if non-nil, zero value otherwise.

### GetResponseStatusCodesOk

`func (o *FiltersApiLog) GetResponseStatusCodesOk() (*[]int32, bool)`

GetResponseStatusCodesOk returns a tuple with the ResponseStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCodes

`func (o *FiltersApiLog) SetResponseStatusCodes(v []int32)`

SetResponseStatusCodes sets ResponseStatusCodes field to given value.

### HasResponseStatusCodes

`func (o *FiltersApiLog) HasResponseStatusCodes() bool`

HasResponseStatusCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


