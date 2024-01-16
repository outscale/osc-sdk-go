# ReadVolumesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) | Information about one or more volumes. | [optional] 

## Methods

### NewReadVolumesResponse

`func NewReadVolumesResponse() *ReadVolumesResponse`

NewReadVolumesResponse instantiates a new ReadVolumesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVolumesResponseWithDefaults

`func NewReadVolumesResponseWithDefaults() *ReadVolumesResponse`

NewReadVolumesResponseWithDefaults instantiates a new ReadVolumesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ReadVolumesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadVolumesResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadVolumesResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadVolumesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadVolumesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVolumesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVolumesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVolumesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVolumes

`func (o *ReadVolumesResponse) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ReadVolumesResponse) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ReadVolumesResponse) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ReadVolumesResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


