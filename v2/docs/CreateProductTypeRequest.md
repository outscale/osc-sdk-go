# CreateProductTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the product type. | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the product type. | [optional] 

## Methods

### NewCreateProductTypeRequest

`func NewCreateProductTypeRequest(description string, ) *CreateProductTypeRequest`

NewCreateProductTypeRequest instantiates a new CreateProductTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductTypeRequestWithDefaults

`func NewCreateProductTypeRequestWithDefaults() *CreateProductTypeRequest`

NewCreateProductTypeRequestWithDefaults instantiates a new CreateProductTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateProductTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDryRun

`func (o *CreateProductTypeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateProductTypeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateProductTypeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateProductTypeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVendor

`func (o *CreateProductTypeRequest) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CreateProductTypeRequest) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CreateProductTypeRequest) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CreateProductTypeRequest) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


