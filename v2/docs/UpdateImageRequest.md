# UpdateImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A new description for the image. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | **string** | The ID of the OMI you want to modify. | 
**PermissionsToLaunch** | Pointer to [**PermissionsOnResourceCreation**](PermissionsOnResourceCreation.md) |  | [optional] 
**ProductCodes** | Pointer to **[]string** | The product codes associated with the OMI. Any previously set value is deleted. Make sure to specify all product codes you want to associate with the OMI. | [optional] 

## Methods

### NewUpdateImageRequest

`func NewUpdateImageRequest(imageId string, ) *UpdateImageRequest`

NewUpdateImageRequest instantiates a new UpdateImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateImageRequestWithDefaults

`func NewUpdateImageRequestWithDefaults() *UpdateImageRequest`

NewUpdateImageRequestWithDefaults instantiates a new UpdateImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateImageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateImageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateImageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateImageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateImageRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateImageRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateImageRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateImageRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetImageId

`func (o *UpdateImageRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *UpdateImageRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *UpdateImageRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetPermissionsToLaunch

`func (o *UpdateImageRequest) GetPermissionsToLaunch() PermissionsOnResourceCreation`

GetPermissionsToLaunch returns the PermissionsToLaunch field if non-nil, zero value otherwise.

### GetPermissionsToLaunchOk

`func (o *UpdateImageRequest) GetPermissionsToLaunchOk() (*PermissionsOnResourceCreation, bool)`

GetPermissionsToLaunchOk returns a tuple with the PermissionsToLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsToLaunch

`func (o *UpdateImageRequest) SetPermissionsToLaunch(v PermissionsOnResourceCreation)`

SetPermissionsToLaunch sets PermissionsToLaunch field to given value.

### HasPermissionsToLaunch

`func (o *UpdateImageRequest) HasPermissionsToLaunch() bool`

HasPermissionsToLaunch returns a boolean if a field has been set.

### GetProductCodes

`func (o *UpdateImageRequest) GetProductCodes() []string`

GetProductCodes returns the ProductCodes field if non-nil, zero value otherwise.

### GetProductCodesOk

`func (o *UpdateImageRequest) GetProductCodesOk() (*[]string, bool)`

GetProductCodesOk returns a tuple with the ProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCodes

`func (o *UpdateImageRequest) SetProductCodes(v []string)`

SetProductCodes sets ProductCodes field to given value.

### HasProductCodes

`func (o *UpdateImageRequest) HasProductCodes() bool`

HasProductCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


