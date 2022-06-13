# ServerCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **string** | The date at which the server certificate expires. | [optional] 
**Id** | Pointer to **string** | The ID of the server certificate. | [optional] 
**Name** | Pointer to **string** | The name of the server certificate. | [optional] 
**Orn** | Pointer to **string** | The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers &gt; Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns). | [optional] 
**Path** | Pointer to **string** | The path to the server certificate. | [optional] 
**UploadDate** | Pointer to **string** | The date at which the server certificate has been uploaded. | [optional] 

## Methods

### NewServerCertificate

`func NewServerCertificate() *ServerCertificate`

NewServerCertificate instantiates a new ServerCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCertificateWithDefaults

`func NewServerCertificateWithDefaults() *ServerCertificate`

NewServerCertificateWithDefaults instantiates a new ServerCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *ServerCertificate) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ServerCertificate) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ServerCertificate) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ServerCertificate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetId

`func (o *ServerCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServerCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrn

`func (o *ServerCertificate) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ServerCertificate) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ServerCertificate) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *ServerCertificate) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetPath

`func (o *ServerCertificate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ServerCertificate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ServerCertificate) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ServerCertificate) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUploadDate

`func (o *ServerCertificate) GetUploadDate() string`

GetUploadDate returns the UploadDate field if non-nil, zero value otherwise.

### GetUploadDateOk

`func (o *ServerCertificate) GetUploadDateOk() (*string, bool)`

GetUploadDateOk returns a tuple with the UploadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadDate

`func (o *ServerCertificate) SetUploadDate(v string)`

SetUploadDate sets UploadDate field to given value.

### HasUploadDate

`func (o *ServerCertificate) HasUploadDate() bool`

HasUploadDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


