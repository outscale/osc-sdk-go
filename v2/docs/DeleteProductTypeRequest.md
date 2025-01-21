# DeleteProductTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Force** | Pointer to **bool** | If true, forces the deletion of the product type associated with one or more OMIs. | [optional] 
**ProductTypeId** | **string** | The ID of the product type you want to delete. | 

## Methods

### NewDeleteProductTypeRequest

`func NewDeleteProductTypeRequest(productTypeId string, ) *DeleteProductTypeRequest`

NewDeleteProductTypeRequest instantiates a new DeleteProductTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteProductTypeRequestWithDefaults

`func NewDeleteProductTypeRequestWithDefaults() *DeleteProductTypeRequest`

NewDeleteProductTypeRequestWithDefaults instantiates a new DeleteProductTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteProductTypeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteProductTypeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteProductTypeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteProductTypeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetForce

`func (o *DeleteProductTypeRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *DeleteProductTypeRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *DeleteProductTypeRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *DeleteProductTypeRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetProductTypeId

`func (o *DeleteProductTypeRequest) GetProductTypeId() string`

GetProductTypeId returns the ProductTypeId field if non-nil, zero value otherwise.

### GetProductTypeIdOk

`func (o *DeleteProductTypeRequest) GetProductTypeIdOk() (*string, bool)`

GetProductTypeIdOk returns a tuple with the ProductTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeId

`func (o *DeleteProductTypeRequest) SetProductTypeId(v string)`

SetProductTypeId sets ProductTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


