# CheckAuthenticationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Login** | **string** | The email address of the account. | 
**Password** | **string** | The password of the account. | 

## Methods

### NewCheckAuthenticationRequest

`func NewCheckAuthenticationRequest(login string, password string, ) *CheckAuthenticationRequest`

NewCheckAuthenticationRequest instantiates a new CheckAuthenticationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAuthenticationRequestWithDefaults

`func NewCheckAuthenticationRequestWithDefaults() *CheckAuthenticationRequest`

NewCheckAuthenticationRequestWithDefaults instantiates a new CheckAuthenticationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CheckAuthenticationRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CheckAuthenticationRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CheckAuthenticationRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CheckAuthenticationRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLogin

`func (o *CheckAuthenticationRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CheckAuthenticationRequest) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CheckAuthenticationRequest) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *CheckAuthenticationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CheckAuthenticationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CheckAuthenticationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


