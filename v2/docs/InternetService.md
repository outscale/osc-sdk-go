# InternetService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternetServiceId** | Pointer to **string** | The ID of the Internet service. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net attached to the Internet service. | [optional] 
**State** | Pointer to **string** | The state of the attachment of the Internet service to the Net (always &#x60;available&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the Internet service. | [optional] 

## Methods

### NewInternetService

`func NewInternetService() *InternetService`

NewInternetService instantiates a new InternetService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetServiceWithDefaults

`func NewInternetServiceWithDefaults() *InternetService`

NewInternetServiceWithDefaults instantiates a new InternetService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternetServiceId

`func (o *InternetService) GetInternetServiceId() string`

GetInternetServiceId returns the InternetServiceId field if non-nil, zero value otherwise.

### GetInternetServiceIdOk

`func (o *InternetService) GetInternetServiceIdOk() (*string, bool)`

GetInternetServiceIdOk returns a tuple with the InternetServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetServiceId

`func (o *InternetService) SetInternetServiceId(v string)`

SetInternetServiceId sets InternetServiceId field to given value.

### HasInternetServiceId

`func (o *InternetService) HasInternetServiceId() bool`

HasInternetServiceId returns a boolean if a field has been set.

### GetNetId

`func (o *InternetService) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *InternetService) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *InternetService) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *InternetService) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetState

`func (o *InternetService) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InternetService) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InternetService) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InternetService) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *InternetService) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InternetService) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InternetService) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InternetService) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


