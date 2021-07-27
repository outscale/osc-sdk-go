# UpdateApiAccessPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**MaxAccessKeyExpirationSeconds** | **int64** | The maximum possible lifetime for your access keys, in seconds (between &#x60;0&#x60; and &#x60;3153600000&#x60;, both included). By default or if set to &#x60;O&#x60;, your access keys can have unlimited lifetimes. Otherwise, all your access keys must have an expiration date. This value must be greater than the remaining lifetime of each access key of your account. | 
**RequireTrustedEnv** | **bool** | If true, a trusted session is activated, provided that you specify the &#x60;MaxAccessKeyExpirationSeconds&#x60; parameter with a value greater than &#x60;0&#x60;. | 

## Methods

### NewUpdateApiAccessPolicyRequest

`func NewUpdateApiAccessPolicyRequest(maxAccessKeyExpirationSeconds int64, requireTrustedEnv bool, ) *UpdateApiAccessPolicyRequest`

NewUpdateApiAccessPolicyRequest instantiates a new UpdateApiAccessPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApiAccessPolicyRequestWithDefaults

`func NewUpdateApiAccessPolicyRequestWithDefaults() *UpdateApiAccessPolicyRequest`

NewUpdateApiAccessPolicyRequestWithDefaults instantiates a new UpdateApiAccessPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateApiAccessPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateApiAccessPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateApiAccessPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateApiAccessPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMaxAccessKeyExpirationSeconds

`func (o *UpdateApiAccessPolicyRequest) GetMaxAccessKeyExpirationSeconds() int64`

GetMaxAccessKeyExpirationSeconds returns the MaxAccessKeyExpirationSeconds field if non-nil, zero value otherwise.

### GetMaxAccessKeyExpirationSecondsOk

`func (o *UpdateApiAccessPolicyRequest) GetMaxAccessKeyExpirationSecondsOk() (*int64, bool)`

GetMaxAccessKeyExpirationSecondsOk returns a tuple with the MaxAccessKeyExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccessKeyExpirationSeconds

`func (o *UpdateApiAccessPolicyRequest) SetMaxAccessKeyExpirationSeconds(v int64)`

SetMaxAccessKeyExpirationSeconds sets MaxAccessKeyExpirationSeconds field to given value.


### GetRequireTrustedEnv

`func (o *UpdateApiAccessPolicyRequest) GetRequireTrustedEnv() bool`

GetRequireTrustedEnv returns the RequireTrustedEnv field if non-nil, zero value otherwise.

### GetRequireTrustedEnvOk

`func (o *UpdateApiAccessPolicyRequest) GetRequireTrustedEnvOk() (*bool, bool)`

GetRequireTrustedEnvOk returns a tuple with the RequireTrustedEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTrustedEnv

`func (o *UpdateApiAccessPolicyRequest) SetRequireTrustedEnv(v bool)`

SetRequireTrustedEnv sets RequireTrustedEnv field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


