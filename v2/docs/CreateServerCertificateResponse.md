# CreateServerCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**ServerCertificate** | Pointer to [**ServerCertificate**](ServerCertificate.md) |  | [optional] 

## Methods

### NewCreateServerCertificateResponse

`func NewCreateServerCertificateResponse() *CreateServerCertificateResponse`

NewCreateServerCertificateResponse instantiates a new CreateServerCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerCertificateResponseWithDefaults

`func NewCreateServerCertificateResponseWithDefaults() *CreateServerCertificateResponse`

NewCreateServerCertificateResponseWithDefaults instantiates a new CreateServerCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateServerCertificateResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateServerCertificateResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateServerCertificateResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateServerCertificateResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetServerCertificate

`func (o *CreateServerCertificateResponse) GetServerCertificate() ServerCertificate`

GetServerCertificate returns the ServerCertificate field if non-nil, zero value otherwise.

### GetServerCertificateOk

`func (o *CreateServerCertificateResponse) GetServerCertificateOk() (*ServerCertificate, bool)`

GetServerCertificateOk returns a tuple with the ServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificate

`func (o *CreateServerCertificateResponse) SetServerCertificate(v ServerCertificate)`

SetServerCertificate sets ServerCertificate field to given value.

### HasServerCertificate

`func (o *CreateServerCertificateResponse) HasServerCertificate() bool`

HasServerCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


