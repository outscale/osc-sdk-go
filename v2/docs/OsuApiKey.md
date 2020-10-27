# OsuApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **string** | The API key of the OSU account that enables you to access the bucket. | [optional] 
**SecretKey** | Pointer to **string** | The secret key of the OSU account that enables you to access the bucket. | [optional] 

## Methods

### NewOsuApiKey

`func NewOsuApiKey() *OsuApiKey`

NewOsuApiKey instantiates a new OsuApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsuApiKeyWithDefaults

`func NewOsuApiKeyWithDefaults() *OsuApiKey`

NewOsuApiKeyWithDefaults instantiates a new OsuApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *OsuApiKey) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *OsuApiKey) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *OsuApiKey) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *OsuApiKey) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetSecretKey

`func (o *OsuApiKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *OsuApiKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *OsuApiKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *OsuApiKey) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


