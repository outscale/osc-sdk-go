# ReadSecretAccessKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to [**AccessKeySecretKey**](AccessKeySecretKey.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadSecretAccessKeyResponse

`func NewReadSecretAccessKeyResponse() *ReadSecretAccessKeyResponse`

NewReadSecretAccessKeyResponse instantiates a new ReadSecretAccessKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSecretAccessKeyResponseWithDefaults

`func NewReadSecretAccessKeyResponseWithDefaults() *ReadSecretAccessKeyResponse`

NewReadSecretAccessKeyResponseWithDefaults instantiates a new ReadSecretAccessKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ReadSecretAccessKeyResponse) GetAccessKey() AccessKeySecretKey`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ReadSecretAccessKeyResponse) GetAccessKeyOk() (*AccessKeySecretKey, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ReadSecretAccessKeyResponse) SetAccessKey(v AccessKeySecretKey)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ReadSecretAccessKeyResponse) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadSecretAccessKeyResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSecretAccessKeyResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadSecretAccessKeyResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadSecretAccessKeyResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


