# ReadPolicyVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyVersion** | Pointer to [**PolicyVersion**](PolicyVersion.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadPolicyVersionResponse

`func NewReadPolicyVersionResponse() *ReadPolicyVersionResponse`

NewReadPolicyVersionResponse instantiates a new ReadPolicyVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPolicyVersionResponseWithDefaults

`func NewReadPolicyVersionResponseWithDefaults() *ReadPolicyVersionResponse`

NewReadPolicyVersionResponseWithDefaults instantiates a new ReadPolicyVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyVersion

`func (o *ReadPolicyVersionResponse) GetPolicyVersion() PolicyVersion`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *ReadPolicyVersionResponse) GetPolicyVersionOk() (*PolicyVersion, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *ReadPolicyVersionResponse) SetPolicyVersion(v PolicyVersion)`

SetPolicyVersion sets PolicyVersion field to given value.

### HasPolicyVersion

`func (o *ReadPolicyVersionResponse) HasPolicyVersion() bool`

HasPolicyVersion returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadPolicyVersionResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadPolicyVersionResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadPolicyVersionResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadPolicyVersionResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


