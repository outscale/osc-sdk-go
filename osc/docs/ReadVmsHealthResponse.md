# ReadVmsHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendVmHealth** | Pointer to [**[]BackendVmHealth**](BackendVmHealth.md) | Information about the health of one or more back-end VMs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadVmsHealthResponse

`func NewReadVmsHealthResponse() *ReadVmsHealthResponse`

NewReadVmsHealthResponse instantiates a new ReadVmsHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmsHealthResponseWithDefaults

`func NewReadVmsHealthResponseWithDefaults() *ReadVmsHealthResponse`

NewReadVmsHealthResponseWithDefaults instantiates a new ReadVmsHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendVmHealth

`func (o *ReadVmsHealthResponse) GetBackendVmHealth() []BackendVmHealth`

GetBackendVmHealth returns the BackendVmHealth field if non-nil, zero value otherwise.

### GetBackendVmHealthOk

`func (o *ReadVmsHealthResponse) GetBackendVmHealthOk() (*[]BackendVmHealth, bool)`

GetBackendVmHealthOk returns a tuple with the BackendVmHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVmHealth

`func (o *ReadVmsHealthResponse) SetBackendVmHealth(v []BackendVmHealth)`

SetBackendVmHealth sets BackendVmHealth field to given value.

### HasBackendVmHealth

`func (o *ReadVmsHealthResponse) HasBackendVmHealth() bool`

HasBackendVmHealth returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadVmsHealthResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmsHealthResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVmsHealthResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVmsHealthResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


