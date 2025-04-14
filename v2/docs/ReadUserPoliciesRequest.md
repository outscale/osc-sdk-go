# ReadUserPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**UserName** | **string** | The name of the user. | 

## Methods

### NewReadUserPoliciesRequest

`func NewReadUserPoliciesRequest(userName string, ) *ReadUserPoliciesRequest`

NewReadUserPoliciesRequest instantiates a new ReadUserPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserPoliciesRequestWithDefaults

`func NewReadUserPoliciesRequestWithDefaults() *ReadUserPoliciesRequest`

NewReadUserPoliciesRequestWithDefaults instantiates a new ReadUserPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadUserPoliciesRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadUserPoliciesRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadUserPoliciesRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadUserPoliciesRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetUserName

`func (o *ReadUserPoliciesRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ReadUserPoliciesRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ReadUserPoliciesRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


