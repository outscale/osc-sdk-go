# CreateAccessKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to [**AccessKeySecretKey**](AccessKeySecretKey.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateAccessKeyResponse

`func NewCreateAccessKeyResponse() *CreateAccessKeyResponse`

NewCreateAccessKeyResponse instantiates a new CreateAccessKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessKeyResponseWithDefaults

`func NewCreateAccessKeyResponseWithDefaults() *CreateAccessKeyResponse`

NewCreateAccessKeyResponseWithDefaults instantiates a new CreateAccessKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *CreateAccessKeyResponse) GetAccessKey() AccessKeySecretKey`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CreateAccessKeyResponse) GetAccessKeyOk() (*AccessKeySecretKey, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CreateAccessKeyResponse) SetAccessKey(v AccessKeySecretKey)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *CreateAccessKeyResponse) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateAccessKeyResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateAccessKeyResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateAccessKeyResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateAccessKeyResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


