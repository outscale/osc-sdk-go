# ReadNetAccessPointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAccessPoints** | Pointer to [**[]NetAccessPoint**](NetAccessPoint.md) | One or more Net access points. | [optional] 
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadNetAccessPointsResponse

`func NewReadNetAccessPointsResponse() *ReadNetAccessPointsResponse`

NewReadNetAccessPointsResponse instantiates a new ReadNetAccessPointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadNetAccessPointsResponseWithDefaults

`func NewReadNetAccessPointsResponseWithDefaults() *ReadNetAccessPointsResponse`

NewReadNetAccessPointsResponseWithDefaults instantiates a new ReadNetAccessPointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetAccessPoints

`func (o *ReadNetAccessPointsResponse) GetNetAccessPoints() []NetAccessPoint`

GetNetAccessPoints returns the NetAccessPoints field if non-nil, zero value otherwise.

### GetNetAccessPointsOk

`func (o *ReadNetAccessPointsResponse) GetNetAccessPointsOk() (*[]NetAccessPoint, bool)`

GetNetAccessPointsOk returns a tuple with the NetAccessPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAccessPoints

`func (o *ReadNetAccessPointsResponse) SetNetAccessPoints(v []NetAccessPoint)`

SetNetAccessPoints sets NetAccessPoints field to given value.

### HasNetAccessPoints

`func (o *ReadNetAccessPointsResponse) HasNetAccessPoints() bool`

HasNetAccessPoints returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ReadNetAccessPointsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadNetAccessPointsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadNetAccessPointsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadNetAccessPointsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadNetAccessPointsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNetAccessPointsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadNetAccessPointsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadNetAccessPointsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


