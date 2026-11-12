# CreateCaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaPem** | **string** | The CA in PEM format. | 
**Description** | Pointer to **string** | The description of the CA. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewCreateCaRequest

`func NewCreateCaRequest(caPem string, ) *CreateCaRequest`

NewCreateCaRequest instantiates a new CreateCaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCaRequestWithDefaults

`func NewCreateCaRequestWithDefaults() *CreateCaRequest`

NewCreateCaRequestWithDefaults instantiates a new CreateCaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaPem

`func (o *CreateCaRequest) GetCaPem() string`

GetCaPem returns the CaPem field if non-nil, zero value otherwise.

### GetCaPemOk

`func (o *CreateCaRequest) GetCaPemOk() (*string, bool)`

GetCaPemOk returns a tuple with the CaPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPem

`func (o *CreateCaRequest) SetCaPem(v string)`

SetCaPem sets CaPem field to given value.


### GetDescription

`func (o *CreateCaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateCaRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateCaRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateCaRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateCaRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


