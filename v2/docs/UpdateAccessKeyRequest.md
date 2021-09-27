# UpdateAccessKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | The ID of the access key. | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ExpirationDate** | Pointer to **string** | The date and time at which you want the access key to expire, in ISO 8601 format (for example, &#x60;2017-06-14&#x60; or &#x60;2017-06-14T00:00:00Z&#x60;). If not specified, the access key is set to not expire. | [optional] 
**State** | **string** | The new state for the access key (&#x60;ACTIVE&#x60; \\| &#x60;INACTIVE&#x60;). When set to &#x60;ACTIVE&#x60;, the access key is enabled and can be used to send requests. When set to &#x60;INACTIVE&#x60;, the access key is disabled. | 

## Methods

### NewUpdateAccessKeyRequest

`func NewUpdateAccessKeyRequest(accessKeyId string, state string, ) *UpdateAccessKeyRequest`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


