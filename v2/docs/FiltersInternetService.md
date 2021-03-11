# FiltersInternetService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternetServiceIds** | Pointer to **[]string** | The IDs of the Internet services. | [optional] 
**LinkNetIds** | Pointer to **[]string** | The IDs of the Nets the Internet services are attached to. | [optional] 
**LinkStates** | Pointer to **[]string** | The current states of the attachments between the Internet services and the Nets (only &#x60;available&#x60;, if the Internet gateway is attached to a VPC). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the Internet services. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the Internet services. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the Internet services, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

## Methods

### NewFiltersInternetService

`func NewFiltersInternetService() *FiltersInternetService`

NewFiltersInternetService instantiates a new FiltersInternetService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersInternetServiceWithDefaults

`func NewFiltersInternetServiceWithDefaults() *FiltersInternetService`

NewFiltersInternetServiceWithDefaults instantiates a new FiltersInternetService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternetServiceIds

`func (o *FiltersInternetService) GetInternetServiceIds() []string`

GetInternetServiceIds returns the InternetServiceIds field if non-nil, zero value otherwise.

### GetInternetServiceIdsOk

`func (o *FiltersInternetService) GetInternetServiceIdsOk() (*[]string, bool)`

GetInternetServiceIdsOk returns a tuple with the InternetServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetServiceIds

`func (o *FiltersInternetService) SetInternetServiceIds(v []string)`

SetInternetServiceIds sets InternetServiceIds field to given value.

### HasInternetServiceIds

`func (o *FiltersInternetService) HasInternetServiceIds() bool`

HasInternetServiceIds returns a boolean if a field has been set.

### GetLinkNetIds

`func (o *FiltersInternetService) GetLinkNetIds() []string`

GetLinkNetIds returns the LinkNetIds field if non-nil, zero value otherwise.

### GetLinkNetIdsOk

`func (o *FiltersInternetService) GetLinkNetIdsOk() (*[]string, bool)`

GetLinkNetIdsOk returns a tuple with the LinkNetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNetIds

`func (o *FiltersInternetService) SetLinkNetIds(v []string)`

SetLinkNetIds sets LinkNetIds field to given value.

### HasLinkNetIds

`func (o *FiltersInternetService) HasLinkNetIds() bool`

HasLinkNetIds returns a boolean if a field has been set.

### GetLinkStates

`func (o *FiltersInternetService) GetLinkStates() []string`

GetLinkStates returns the LinkStates field if non-nil, zero value otherwise.

### GetLinkStatesOk

`func (o *FiltersInternetService) GetLinkStatesOk() (*[]string, bool)`

GetLinkStatesOk returns a tuple with the LinkStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStates

`func (o *FiltersInternetService) SetLinkStates(v []string)`

SetLinkStates sets LinkStates field to given value.

### HasLinkStates

`func (o *FiltersInternetService) HasLinkStates() bool`

HasLinkStates returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersInternetService) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersInternetService) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersInternetService) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersInternetService) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersInternetService) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersInternetService) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersInternetService) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersInternetService) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersInternetService) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersInternetService) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersInternetService) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersInternetService) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


