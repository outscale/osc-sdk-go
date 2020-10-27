# SendResetPasswordEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Email** | **string** | The email address provided for the account. | 

## Methods

### NewSendResetPasswordEmailRequest

`func NewSendResetPasswordEmailRequest(email string, ) *SendResetPasswordEmailRequest`

NewSendResetPasswordEmailRequest instantiates a new SendResetPasswordEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendResetPasswordEmailRequestWithDefaults

`func NewSendResetPasswordEmailRequestWithDefaults() *SendResetPasswordEmailRequest`

NewSendResetPasswordEmailRequestWithDefaults instantiates a new SendResetPasswordEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *SendResetPasswordEmailRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SendResetPasswordEmailRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SendResetPasswordEmailRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *SendResetPasswordEmailRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetEmail

`func (o *SendResetPasswordEmailRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendResetPasswordEmailRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendResetPasswordEmailRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


