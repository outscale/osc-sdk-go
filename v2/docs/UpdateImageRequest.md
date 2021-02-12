# UpdateImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | **string** | The ID of the OMI you want to modify. | 
**PermissionsToLaunch** | [**PermissionsOnResourceCreation**](PermissionsOnResourceCreation.md) |  | 

## Methods

### NewUpdateImageRequest

`func NewUpdateImageRequest(imageId string, permissionsToLaunch PermissionsOnResourceCreation, ) *UpdateImageRequest`

NewUpdateImageRequest instantiates a new UpdateImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateImageRequestWithDefaults

`func NewUpdateImageRequestWithDefaults() *UpdateImageRequest`

NewUpdateImageRequestWithDefaults instantiates a new UpdateImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


