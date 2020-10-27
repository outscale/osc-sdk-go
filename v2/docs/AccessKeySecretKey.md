# AccessKeySecretKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | The ID of the secret access key. | [optional] 
**CreationDate** | Pointer to **string** | The date and time of creation of the secret access key. | [optional] 
**ExpirationDate** | Pointer to **string** | The date at which the access key expires. | [optional] 
**LastModificationDate** | Pointer to **string** | The date and time of the last modification of the secret access key. | [optional] 
**SecretKey** | Pointer to **string** | The secret access key that enables you to send requests. | [optional] 
**State** | Pointer to **string** | The state of the secret access key (&#x60;ACTIVE&#x60; if the key is valid for API calls, or &#x60;INACTIVE&#x60; if not). | [optional] 

## Methods

### NewAccessKeySecretKey

`func NewAccessKeySecretKey() *AccessKeySecretKey`

NewAccessKeySecretKey instantiates a new AccessKeySecretKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeySecretKeyWithDefaults

`func NewAccessKeySecretKeyWithDefaults() *AccessKeySecretKey`

NewAccessKeySecretKeyWithDefaults instantiates a new AccessKeySecretKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AccessKeySecretKey) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AccessKeySecretKey) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AccessKeySecretKey) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AccessKeySecretKey) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetCreationDate

`func (o *AccessKeySecretKey) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AccessKeySecretKey) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AccessKeySecretKey) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AccessKeySecretKey) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AccessKeySecretKey) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AccessKeySecretKey) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AccessKeySecretKey) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AccessKeySecretKey) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLastModificationDate

`func (o *AccessKeySecretKey) GetLastModificationDate() string`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *AccessKeySecretKey) GetLastModificationDateOk() (*string, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *AccessKeySecretKey) SetLastModificationDate(v string)`

SetLastModificationDate sets LastModificationDate field to given value.

### HasLastModificationDate

`func (o *AccessKeySecretKey) HasLastModificationDate() bool`

HasLastModificationDate returns a boolean if a field has been set.

### GetSecretKey

`func (o *AccessKeySecretKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AccessKeySecretKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AccessKeySecretKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AccessKeySecretKey) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetState

`func (o *AccessKeySecretKey) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccessKeySecretKey) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccessKeySecretKey) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AccessKeySecretKey) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


