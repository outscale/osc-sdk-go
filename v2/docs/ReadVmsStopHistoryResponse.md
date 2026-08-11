# ReadVmsStopHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmsStopHistory** | Pointer to [**[]VmsStopHistory**](VmsStopHistory.md) | Information about the VM(s) stop history. | [optional] 

## Methods

### NewReadVmsStopHistoryResponse

`func NewReadVmsStopHistoryResponse() *ReadVmsStopHistoryResponse`

NewReadVmsStopHistoryResponse instantiates a new ReadVmsStopHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmsStopHistoryResponseWithDefaults

`func NewReadVmsStopHistoryResponseWithDefaults() *ReadVmsStopHistoryResponse`

NewReadVmsStopHistoryResponseWithDefaults instantiates a new ReadVmsStopHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ReadVmsStopHistoryResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadVmsStopHistoryResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadVmsStopHistoryResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadVmsStopHistoryResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadVmsStopHistoryResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmsStopHistoryResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVmsStopHistoryResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVmsStopHistoryResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmsStopHistory

`func (o *ReadVmsStopHistoryResponse) GetVmsStopHistory() []VmsStopHistory`

GetVmsStopHistory returns the VmsStopHistory field if non-nil, zero value otherwise.

### GetVmsStopHistoryOk

`func (o *ReadVmsStopHistoryResponse) GetVmsStopHistoryOk() (*[]VmsStopHistory, bool)`

GetVmsStopHistoryOk returns a tuple with the VmsStopHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmsStopHistory

`func (o *ReadVmsStopHistoryResponse) SetVmsStopHistory(v []VmsStopHistory)`

SetVmsStopHistory sets VmsStopHistory field to given value.

### HasVmsStopHistory

`func (o *ReadVmsStopHistoryResponse) HasVmsStopHistory() bool`

HasVmsStopHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


