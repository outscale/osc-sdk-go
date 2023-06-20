# DeleteAccessKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | The ID of the access key you want to delete. | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**UserName** | Pointer to **string** | The name of the EIM user the access key you want to delete is associated with. By default, the user who sends the request (which can be the root account). | [optional] 

## Methods

### NewDeleteAccessKeyRequest

`func NewDeleteAccessKeyRequest(accessKeyId string, ) *DeleteAccessKeyRequest`

NewDeleteAccessKeyRequest instantiates a new DeleteAccessKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAccessKeyRequestWithDefaults

`func NewDeleteAccessKeyRequestWithDefaults() *DeleteAccessKeyRequest`

NewDeleteAccessKeyRequestWithDefaults instantiates a new DeleteAccessKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *DeleteAccessKeyRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *DeleteAccessKeyRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *DeleteAccessKeyRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetDryRun

`func (o *DeleteAccessKeyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteAccessKeyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteAccessKeyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteAccessKeyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetUserName

`func (o *DeleteAccessKeyRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DeleteAccessKeyRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DeleteAccessKeyRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DeleteAccessKeyRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


