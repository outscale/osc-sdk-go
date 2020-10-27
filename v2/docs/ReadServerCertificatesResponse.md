# ReadServerCertificatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**ServerCertificates** | Pointer to [**[]ServerCertificate**](ServerCertificate.md) | Information about one or more server certificates. | [optional] 

## Methods

### NewReadServerCertificatesResponse

`func NewReadServerCertificatesResponse() *ReadServerCertificatesResponse`

NewReadServerCertificatesResponse instantiates a new ReadServerCertificatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadServerCertificatesResponseWithDefaults

`func NewReadServerCertificatesResponseWithDefaults() *ReadServerCertificatesResponse`

NewReadServerCertificatesResponseWithDefaults instantiates a new ReadServerCertificatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadServerCertificatesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadServerCertificatesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadServerCertificatesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadServerCertificatesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetServerCertificates

`func (o *ReadServerCertificatesResponse) GetServerCertificates() []ServerCertificate`

GetServerCertificates returns the ServerCertificates field if non-nil, zero value otherwise.

### GetServerCertificatesOk

`func (o *ReadServerCertificatesResponse) GetServerCertificatesOk() (*[]ServerCertificate, bool)`

GetServerCertificatesOk returns a tuple with the ServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificates

`func (o *ReadServerCertificatesResponse) SetServerCertificates(v []ServerCertificate)`

SetServerCertificates sets ServerCertificates field to given value.

### HasServerCertificates

`func (o *ReadServerCertificatesResponse) HasServerCertificates() bool`

HasServerCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


