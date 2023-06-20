# ReadUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) | A list of EIM users. | [optional] 

## Methods

### NewReadUsersResponse

`func NewReadUsersResponse() *ReadUsersResponse`

NewReadUsersResponse instantiates a new ReadUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUsersResponseWithDefaults

`func NewReadUsersResponseWithDefaults() *ReadUsersResponse`

NewReadUsersResponseWithDefaults instantiates a new ReadUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadUsersResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadUsersResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadUsersResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadUsersResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetUsers

`func (o *ReadUsersResponse) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ReadUsersResponse) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ReadUsersResponse) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ReadUsersResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


