# OsuExportImageExportTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskImageFormat** | **string** | The format of the export disk (&#x60;qcow2&#x60; \\| &#x60;raw&#x60;). | 
**OsuBucket** | **string** | The name of the OOS bucket the OMI is exported to. | 
**OsuManifestUrl** | Pointer to **string** | The URL of the manifest file. | [optional] 
**OsuPrefix** | Pointer to **string** | The prefix for the key of the OOS object corresponding to the image. | [optional] 

## Methods

### NewOsuExportImageExportTask

`func NewOsuExportImageExportTask(diskImageFormat string, osuBucket string, ) *OsuExportImageExportTask`

NewOsuExportImageExportTask instantiates a new OsuExportImageExportTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsuExportImageExportTaskWithDefaults

`func NewOsuExportImageExportTaskWithDefaults() *OsuExportImageExportTask`

NewOsuExportImageExportTaskWithDefaults instantiates a new OsuExportImageExportTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskImageFormat

`func (o *OsuExportImageExportTask) GetDiskImageFormat() string`

GetDiskImageFormat returns the DiskImageFormat field if non-nil, zero value otherwise.

### GetDiskImageFormatOk

`func (o *OsuExportImageExportTask) GetDiskImageFormatOk() (*string, bool)`

GetDiskImageFormatOk returns a tuple with the DiskImageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskImageFormat

`func (o *OsuExportImageExportTask) SetDiskImageFormat(v string)`

SetDiskImageFormat sets DiskImageFormat field to given value.


### GetOsuBucket

`func (o *OsuExportImageExportTask) GetOsuBucket() string`

GetOsuBucket returns the OsuBucket field if non-nil, zero value otherwise.

### GetOsuBucketOk

`func (o *OsuExportImageExportTask) GetOsuBucketOk() (*string, bool)`

GetOsuBucketOk returns a tuple with the OsuBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuBucket

`func (o *OsuExportImageExportTask) SetOsuBucket(v string)`

SetOsuBucket sets OsuBucket field to given value.


### GetOsuManifestUrl

`func (o *OsuExportImageExportTask) GetOsuManifestUrl() string`

GetOsuManifestUrl returns the OsuManifestUrl field if non-nil, zero value otherwise.

### GetOsuManifestUrlOk

`func (o *OsuExportImageExportTask) GetOsuManifestUrlOk() (*string, bool)`

GetOsuManifestUrlOk returns a tuple with the OsuManifestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuManifestUrl

`func (o *OsuExportImageExportTask) SetOsuManifestUrl(v string)`

SetOsuManifestUrl sets OsuManifestUrl field to given value.

### HasOsuManifestUrl

`func (o *OsuExportImageExportTask) HasOsuManifestUrl() bool`

HasOsuManifestUrl returns a boolean if a field has been set.

### GetOsuPrefix

`func (o *OsuExportImageExportTask) GetOsuPrefix() string`

GetOsuPrefix returns the OsuPrefix field if non-nil, zero value otherwise.

### GetOsuPrefixOk

`func (o *OsuExportImageExportTask) GetOsuPrefixOk() (*string, bool)`

GetOsuPrefixOk returns a tuple with the OsuPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuPrefix

`func (o *OsuExportImageExportTask) SetOsuPrefix(v string)`

SetOsuPrefix sets OsuPrefix field to given value.

### HasOsuPrefix

`func (o *OsuExportImageExportTask) HasOsuPrefix() bool`

HasOsuPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


