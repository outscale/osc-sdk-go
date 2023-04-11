# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | **string** | The account alias of the owner of the OMI. | [optional] 
**AccountId** | **string** | The account ID of the owner of the OMI. | [optional] 
**Architecture** | **string** | The architecture of the OMI (by default, &#x60;i386&#x60;). | [optional] 
**BlockDeviceMappings** | [**[]BlockDeviceMappingImage**](BlockDeviceMappingImage.md) | One or more block device mappings. | [optional] 
**CreationDate** | [**time.Time**](time.Time.md) | The date and time of creation of the OMI, in ISO 8601 date-time format. | [optional] 
**Description** | **string** | The description of the OMI. | [optional] 
**FileLocation** | **string** | The location of the bucket where the OMI files are stored. | [optional] 
**ImageId** | **string** | The ID of the OMI. | [optional] 
**ImageName** | **string** | The name of the OMI. | [optional] 
**ImageType** | **string** | The type of the OMI. | [optional] 
**PermissionsToLaunch** | [**PermissionsOnResource**](PermissionsOnResource.md) |  | [optional] 
**ProductCodes** | **[]string** | The product codes associated with the OMI. | [optional] 
**RootDeviceName** | **string** | The name of the root device. | [optional] 
**RootDeviceType** | **string** | The type of root device used by the OMI (always &#x60;bsu&#x60;). | [optional] 
**State** | **string** | The state of the OMI (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;failed&#x60;). | [optional] 
**StateComment** | [**StateComment**](StateComment.md) |  | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the OMI. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


