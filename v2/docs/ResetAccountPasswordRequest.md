# ResetAccountPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Password** | **string** | The new password for the account. | 
**Token** | **string** | The token you received at the email address provided for the account. | 

## Methods

### NewResetAccountPasswordRequest

`func NewResetAccountPasswordRequest(password string, token string, ) *ResetAccountPasswordRequest`

NewResetAccountPasswordRequest instantiates a new ResetAccountPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetAccountPasswordRequestWithDefaults

`func NewResetAccountPasswordRequestWithDefaults() *ResetAccountPasswordRequest`

NewResetAccountPasswordRequestWithDefaults instantiates a new ResetAccountPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ResetAccountPasswordRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ResetAccountPasswordRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ResetAccountPasswordRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ResetAccountPasswordRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPassword

`func (o *ResetAccountPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ResetAccountPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ResetAccountPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *ResetAccountPasswordRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResetAccountPasswordRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResetAccountPasswordRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


