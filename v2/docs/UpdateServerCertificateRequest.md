# UpdateServerCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Name** | **string** | The name of the server certificate you want to modify. | 
**NewName** | Pointer to **string** | A new name for the server certificate. | [optional] 
**NewPath** | Pointer to **string** | A new path for the server certificate. | [optional] 

## Methods

### NewUpdateServerCertificateRequest

`func NewUpdateServerCertificateRequest(name string, ) *UpdateServerCertificateRequest`

NewUpdateServerCertificateRequest instantiates a new UpdateServerCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerCertificateRequestWithDefaults

`func NewUpdateServerCertificateRequestWithDefaults() *UpdateServerCertificateRequest`

NewUpdateServerCertificateRequestWithDefaults instantiates a new UpdateServerCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateServerCertificateRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateServerCertificateRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateServerCertificateRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateServerCertificateRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetName

`func (o *UpdateServerCertificateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServerCertificateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServerCertificateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateServerCertificateRequest) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateServerCertificateRequest) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateServerCertificateRequest) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateServerCertificateRequest) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNewPath

`func (o *UpdateServerCertificateRequest) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *UpdateServerCertificateRequest) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *UpdateServerCertificateRequest) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *UpdateServerCertificateRequest) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


