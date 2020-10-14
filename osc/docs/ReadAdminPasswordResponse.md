# ReadAdminPasswordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPassword** | Pointer to **string** | The password of the VM. After the first boot, returns an empty string. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 

## Methods

### NewReadAdminPasswordResponse

`func NewReadAdminPasswordResponse() *ReadAdminPasswordResponse`

NewReadAdminPasswordResponse instantiates a new ReadAdminPasswordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadAdminPasswordResponseWithDefaults

`func NewReadAdminPasswordResponseWithDefaults() *ReadAdminPasswordResponse`

NewReadAdminPasswordResponseWithDefaults instantiates a new ReadAdminPasswordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPassword

`func (o *ReadAdminPasswordResponse) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *ReadAdminPasswordResponse) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *ReadAdminPasswordResponse) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *ReadAdminPasswordResponse) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadAdminPasswordResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadAdminPasswordResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadAdminPasswordResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadAdminPasswordResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmId

`func (o *ReadAdminPasswordResponse) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *ReadAdminPasswordResponse) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *ReadAdminPasswordResponse) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *ReadAdminPasswordResponse) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


