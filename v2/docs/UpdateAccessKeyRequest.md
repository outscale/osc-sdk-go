# UpdateAccessKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | The ID of the access key. | 
**ClearExpirationDate** | Pointer to **bool** | If true, the current expiration date is deleted and the access key is set to not expire. | [optional] 
**ClearTag** | Pointer to **bool** | If true, the current tag of the access key is deleted. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ExpirationDate** | Pointer to **string** | The date and time, or the date, at which you want the access key to expire, in ISO 8601 format (for example, &#x60;2020-06-14T00:00:00.000Z&#x60; or &#x60;2020-06-14&#x60;). If not specified, the access key is set to not expire. If the &#x60;ClearExpirationDate&#x60; parameter is set to true, the expiration date is ignored. | [optional] 
**State** | Pointer to **string** | The new state for the access key (&#x60;ACTIVE&#x60; \\| &#x60;INACTIVE&#x60;). When set to &#x60;ACTIVE&#x60;, the access key is enabled and can be used to send requests. When set to &#x60;INACTIVE&#x60;, the access key is disabled. | [optional] 
**Tag** | Pointer to **string** | A new tag to add to the access key. If the access key already had a tag, this replaces it. If the &#x60;ClearTag&#x60; parameter is set to true, the tag is ignored. | [optional] 
**UserName** | Pointer to **string** | The name of the EIM user that the access key you want to modify is associated with. If you do not specify a user name, this action modifies the access key of the user who sends the request (which can be the root user). | [optional] 

## Methods

### NewUpdateAccessKeyRequest

`func NewUpdateAccessKeyRequest(accessKeyId string, ) *UpdateAccessKeyRequest`

NewUpdateAccessKeyRequest instantiates a new UpdateAccessKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccessKeyRequestWithDefaults

`func NewUpdateAccessKeyRequestWithDefaults() *UpdateAccessKeyRequest`

NewUpdateAccessKeyRequestWithDefaults instantiates a new UpdateAccessKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *UpdateAccessKeyRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *UpdateAccessKeyRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *UpdateAccessKeyRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetClearExpirationDate

`func (o *UpdateAccessKeyRequest) GetClearExpirationDate() bool`

GetClearExpirationDate returns the ClearExpirationDate field if non-nil, zero value otherwise.

### GetClearExpirationDateOk

`func (o *UpdateAccessKeyRequest) GetClearExpirationDateOk() (*bool, bool)`

GetClearExpirationDateOk returns a tuple with the ClearExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearExpirationDate

`func (o *UpdateAccessKeyRequest) SetClearExpirationDate(v bool)`

SetClearExpirationDate sets ClearExpirationDate field to given value.

### HasClearExpirationDate

`func (o *UpdateAccessKeyRequest) HasClearExpirationDate() bool`

HasClearExpirationDate returns a boolean if a field has been set.

### GetClearTag

`func (o *UpdateAccessKeyRequest) GetClearTag() bool`

GetClearTag returns the ClearTag field if non-nil, zero value otherwise.

### GetClearTagOk

`func (o *UpdateAccessKeyRequest) GetClearTagOk() (*bool, bool)`

GetClearTagOk returns a tuple with the ClearTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearTag

`func (o *UpdateAccessKeyRequest) SetClearTag(v bool)`

SetClearTag sets ClearTag field to given value.

### HasClearTag

`func (o *UpdateAccessKeyRequest) HasClearTag() bool`

HasClearTag returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateAccessKeyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateAccessKeyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateAccessKeyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateAccessKeyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetExpirationDate

`func (o *UpdateAccessKeyRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *UpdateAccessKeyRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *UpdateAccessKeyRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *UpdateAccessKeyRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetState

`func (o *UpdateAccessKeyRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateAccessKeyRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateAccessKeyRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateAccessKeyRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *UpdateAccessKeyRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UpdateAccessKeyRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UpdateAccessKeyRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *UpdateAccessKeyRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUserName

`func (o *UpdateAccessKeyRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UpdateAccessKeyRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UpdateAccessKeyRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UpdateAccessKeyRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


