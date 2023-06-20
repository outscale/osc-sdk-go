# AccessKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | The ID of the access key. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time (UTC) of creation of the access key. | [optional] 
**ExpirationDate** | Pointer to **time.Time** | The date (UTC) at which the access key expires. | [optional] 
**LastModificationDate** | Pointer to **time.Time** | The date and time (UTC) of the last modification of the access key. | [optional] 
**State** | Pointer to **string** | The state of the access key (&#x60;ACTIVE&#x60; if the key is valid for API calls, or &#x60;INACTIVE&#x60; if not). | [optional] 

## Methods

### NewAccessKey

`func NewAccessKey() *AccessKey`

NewAccessKey instantiates a new AccessKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyWithDefaults

`func NewAccessKeyWithDefaults() *AccessKey`

NewAccessKeyWithDefaults instantiates a new AccessKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AccessKey) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AccessKey) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AccessKey) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AccessKey) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetCreationDate

`func (o *AccessKey) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AccessKey) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AccessKey) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AccessKey) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AccessKey) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AccessKey) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AccessKey) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AccessKey) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLastModificationDate

`func (o *AccessKey) GetLastModificationDate() time.Time`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *AccessKey) GetLastModificationDateOk() (*time.Time, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *AccessKey) SetLastModificationDate(v time.Time)`

SetLastModificationDate sets LastModificationDate field to given value.

### HasLastModificationDate

`func (o *AccessKey) HasLastModificationDate() bool`

HasLastModificationDate returns a boolean if a field has been set.

### GetState

`func (o *AccessKey) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccessKey) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccessKey) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AccessKey) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


