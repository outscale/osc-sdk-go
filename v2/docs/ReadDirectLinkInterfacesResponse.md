# ReadDirectLinkInterfacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinkInterfaces** | Pointer to [**[]DirectLinkInterfaces**](DirectLinkInterfaces.md) | Information about one or more DirectLink interfaces. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadDirectLinkInterfacesResponse

`func NewReadDirectLinkInterfacesResponse() *ReadDirectLinkInterfacesResponse`

NewReadDirectLinkInterfacesResponse instantiates a new ReadDirectLinkInterfacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadDirectLinkInterfacesResponseWithDefaults

`func NewReadDirectLinkInterfacesResponseWithDefaults() *ReadDirectLinkInterfacesResponse`

NewReadDirectLinkInterfacesResponseWithDefaults instantiates a new ReadDirectLinkInterfacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectLinkInterfaces

`func (o *ReadDirectLinkInterfacesResponse) GetDirectLinkInterfaces() []DirectLinkInterfaces`

GetDirectLinkInterfaces returns the DirectLinkInterfaces field if non-nil, zero value otherwise.

### GetDirectLinkInterfacesOk

`func (o *ReadDirectLinkInterfacesResponse) GetDirectLinkInterfacesOk() (*[]DirectLinkInterfaces, bool)`

GetDirectLinkInterfacesOk returns a tuple with the DirectLinkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkInterfaces

`func (o *ReadDirectLinkInterfacesResponse) SetDirectLinkInterfaces(v []DirectLinkInterfaces)`

SetDirectLinkInterfaces sets DirectLinkInterfaces field to given value.

### HasDirectLinkInterfaces

`func (o *ReadDirectLinkInterfacesResponse) HasDirectLinkInterfaces() bool`

HasDirectLinkInterfaces returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadDirectLinkInterfacesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadDirectLinkInterfacesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadDirectLinkInterfacesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadDirectLinkInterfacesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


