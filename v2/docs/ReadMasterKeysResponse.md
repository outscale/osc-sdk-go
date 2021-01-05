# ReadMasterKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterKeys** | Pointer to [**[]MasterKey**](MasterKey.md) | Information about one or more master keys. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadMasterKeysResponse

`func NewReadMasterKeysResponse() *ReadMasterKeysResponse`

NewReadMasterKeysResponse instantiates a new ReadMasterKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadMasterKeysResponseWithDefaults

`func NewReadMasterKeysResponseWithDefaults() *ReadMasterKeysResponse`

NewReadMasterKeysResponseWithDefaults instantiates a new ReadMasterKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterKeys

`func (o *ReadMasterKeysResponse) GetMasterKeys() []MasterKey`

GetMasterKeys returns the MasterKeys field if non-nil, zero value otherwise.

### GetMasterKeysOk

`func (o *ReadMasterKeysResponse) GetMasterKeysOk() (*[]MasterKey, bool)`

GetMasterKeysOk returns a tuple with the MasterKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeys

`func (o *ReadMasterKeysResponse) SetMasterKeys(v []MasterKey)`

SetMasterKeys sets MasterKeys field to given value.

### HasMasterKeys

`func (o *ReadMasterKeysResponse) HasMasterKeys() bool`

HasMasterKeys returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadMasterKeysResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadMasterKeysResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadMasterKeysResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadMasterKeysResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


