# ReadAccessKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeys** | Pointer to [**[]AccessKey**](AccessKey.md) | A list of access keys. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadAccessKeysResponse

`func NewReadAccessKeysResponse() *ReadAccessKeysResponse`

NewReadAccessKeysResponse instantiates a new ReadAccessKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadAccessKeysResponseWithDefaults

`func NewReadAccessKeysResponseWithDefaults() *ReadAccessKeysResponse`

NewReadAccessKeysResponseWithDefaults instantiates a new ReadAccessKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeys

`func (o *ReadAccessKeysResponse) GetAccessKeys() []AccessKey`

GetAccessKeys returns the AccessKeys field if non-nil, zero value otherwise.

### GetAccessKeysOk

`func (o *ReadAccessKeysResponse) GetAccessKeysOk() (*[]AccessKey, bool)`

GetAccessKeysOk returns a tuple with the AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeys

`func (o *ReadAccessKeysResponse) SetAccessKeys(v []AccessKey)`

SetAccessKeys sets AccessKeys field to given value.

### HasAccessKeys

`func (o *ReadAccessKeysResponse) HasAccessKeys() bool`

HasAccessKeys returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadAccessKeysResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadAccessKeysResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadAccessKeysResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadAccessKeysResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


