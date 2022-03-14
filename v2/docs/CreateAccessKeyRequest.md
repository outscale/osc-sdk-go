# CreateAccessKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ExpirationDate** | Pointer to **string** | The date and time at which you want the access key to expire, in ISO 8601 format (for example, &#x60;2017-06-14&#x60; or &#x60;2017-06-14T00:00:00Z&#x60;). To remove an existing expiration date, use the method without specifying this parameter. | [optional] 

## Methods

### NewCreateAccessKeyRequest

`func NewCreateAccessKeyRequest() *CreateAccessKeyRequest`

NewCreateAccessKeyRequest instantiates a new CreateAccessKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessKeyRequestWithDefaults

`func NewCreateAccessKeyRequestWithDefaults() *CreateAccessKeyRequest`

NewCreateAccessKeyRequestWithDefaults instantiates a new CreateAccessKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateAccessKeyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateAccessKeyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateAccessKeyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateAccessKeyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CreateAccessKeyRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CreateAccessKeyRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CreateAccessKeyRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CreateAccessKeyRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


