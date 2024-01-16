# ReadNetPeeringsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetPeerings** | Pointer to [**[]NetPeering**](NetPeering.md) | Information about one or more Net peerings. | [optional] 
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadNetPeeringsResponse

`func NewReadNetPeeringsResponse() *ReadNetPeeringsResponse`

NewReadNetPeeringsResponse instantiates a new ReadNetPeeringsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadNetPeeringsResponseWithDefaults

`func NewReadNetPeeringsResponseWithDefaults() *ReadNetPeeringsResponse`

NewReadNetPeeringsResponseWithDefaults instantiates a new ReadNetPeeringsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetPeerings

`func (o *ReadNetPeeringsResponse) GetNetPeerings() []NetPeering`

GetNetPeerings returns the NetPeerings field if non-nil, zero value otherwise.

### GetNetPeeringsOk

`func (o *ReadNetPeeringsResponse) GetNetPeeringsOk() (*[]NetPeering, bool)`

GetNetPeeringsOk returns a tuple with the NetPeerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPeerings

`func (o *ReadNetPeeringsResponse) SetNetPeerings(v []NetPeering)`

SetNetPeerings sets NetPeerings field to given value.

### HasNetPeerings

`func (o *ReadNetPeeringsResponse) HasNetPeerings() bool`

HasNetPeerings returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ReadNetPeeringsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadNetPeeringsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadNetPeeringsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadNetPeeringsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadNetPeeringsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNetPeeringsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadNetPeeringsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadNetPeeringsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


