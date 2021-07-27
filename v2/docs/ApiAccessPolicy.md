# ApiAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAccessKeyExpirationSeconds** | Pointer to **int64** | The maximum possible lifetime for your access keys, in seconds. If &#x60;0&#x60;, your access keys can have unlimited lifetimes. | [optional] 
**RequireTrustedEnv** | Pointer to **bool** | If true, a trusted session is activated, allowing you to bypass Certificate Authorities (CAs) enforcement. For more information, see the &#x60;ApiKeyAuth&#x60; authentication scheme in the [Authentication](#authentication) section. | [optional] 

## Methods

### NewApiAccessPolicy

`func NewApiAccessPolicy() *ApiAccessPolicy`

NewApiAccessPolicy instantiates a new ApiAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccessPolicyWithDefaults

`func NewApiAccessPolicyWithDefaults() *ApiAccessPolicy`

NewApiAccessPolicyWithDefaults instantiates a new ApiAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAccessKeyExpirationSeconds

`func (o *ApiAccessPolicy) GetMaxAccessKeyExpirationSeconds() int64`

GetMaxAccessKeyExpirationSeconds returns the MaxAccessKeyExpirationSeconds field if non-nil, zero value otherwise.

### GetMaxAccessKeyExpirationSecondsOk

`func (o *ApiAccessPolicy) GetMaxAccessKeyExpirationSecondsOk() (*int64, bool)`

GetMaxAccessKeyExpirationSecondsOk returns a tuple with the MaxAccessKeyExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccessKeyExpirationSeconds

`func (o *ApiAccessPolicy) SetMaxAccessKeyExpirationSeconds(v int64)`

SetMaxAccessKeyExpirationSeconds sets MaxAccessKeyExpirationSeconds field to given value.

### HasMaxAccessKeyExpirationSeconds

`func (o *ApiAccessPolicy) HasMaxAccessKeyExpirationSeconds() bool`

HasMaxAccessKeyExpirationSeconds returns a boolean if a field has been set.

### GetRequireTrustedEnv

`func (o *ApiAccessPolicy) GetRequireTrustedEnv() bool`

GetRequireTrustedEnv returns the RequireTrustedEnv field if non-nil, zero value otherwise.

### GetRequireTrustedEnvOk

`func (o *ApiAccessPolicy) GetRequireTrustedEnvOk() (*bool, bool)`

GetRequireTrustedEnvOk returns a tuple with the RequireTrustedEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTrustedEnv

`func (o *ApiAccessPolicy) SetRequireTrustedEnv(v bool)`

SetRequireTrustedEnv sets RequireTrustedEnv field to given value.

### HasRequireTrustedEnv

`func (o *ApiAccessPolicy) HasRequireTrustedEnv() bool`

HasRequireTrustedEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


