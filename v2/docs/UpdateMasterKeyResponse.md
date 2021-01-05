# UpdateMasterKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterKey** | Pointer to [**MasterKey**](MasterKey.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewUpdateMasterKeyResponse

`func NewUpdateMasterKeyResponse() *UpdateMasterKeyResponse`

NewUpdateMasterKeyResponse instantiates a new UpdateMasterKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMasterKeyResponseWithDefaults

`func NewUpdateMasterKeyResponseWithDefaults() *UpdateMasterKeyResponse`

NewUpdateMasterKeyResponseWithDefaults instantiates a new UpdateMasterKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterKey

`func (o *UpdateMasterKeyResponse) GetMasterKey() MasterKey`

GetMasterKey returns the MasterKey field if non-nil, zero value otherwise.

### GetMasterKeyOk

`func (o *UpdateMasterKeyResponse) GetMasterKeyOk() (*MasterKey, bool)`

GetMasterKeyOk returns a tuple with the MasterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKey

`func (o *UpdateMasterKeyResponse) SetMasterKey(v MasterKey)`

SetMasterKey sets MasterKey field to given value.

### HasMasterKey

`func (o *UpdateMasterKeyResponse) HasMasterKey() bool`

HasMasterKey returns a boolean if a field has been set.

### GetResponseContext

`func (o *UpdateMasterKeyResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UpdateMasterKeyResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UpdateMasterKeyResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UpdateMasterKeyResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


