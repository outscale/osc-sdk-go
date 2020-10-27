# OsuExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskImageFormat** | **string** | The format of the export disk (&#x60;qcow2&#x60; \\| &#x60;raw&#x60;). | 
**OsuApiKey** | Pointer to [**OsuApiKey**](OsuApiKey.md) |  | [optional] 
**OsuBucket** | **string** | The name of the OSU bucket where you want to export the object. | 
**OsuManifestUrl** | Pointer to **string** | The URL of the manifest file. | [optional] 
**OsuPrefix** | Pointer to **string** | The prefix for the key of the OSU object. | [optional] 

## Methods

### NewOsuExport

`func NewOsuExport(diskImageFormat string, osuBucket string, ) *OsuExport`

NewOsuExport instantiates a new OsuExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsuExportWithDefaults

`func NewOsuExportWithDefaults() *OsuExport`

NewOsuExportWithDefaults instantiates a new OsuExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskImageFormat

`func (o *OsuExport) GetDiskImageFormat() string`

GetDiskImageFormat returns the DiskImageFormat field if non-nil, zero value otherwise.

### GetDiskImageFormatOk

`func (o *OsuExport) GetDiskImageFormatOk() (*string, bool)`

GetDiskImageFormatOk returns a tuple with the DiskImageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskImageFormat

`func (o *OsuExport) SetDiskImageFormat(v string)`

SetDiskImageFormat sets DiskImageFormat field to given value.


### GetOsuApiKey

`func (o *OsuExport) GetOsuApiKey() OsuApiKey`

GetOsuApiKey returns the OsuApiKey field if non-nil, zero value otherwise.

### GetOsuApiKeyOk

`func (o *OsuExport) GetOsuApiKeyOk() (*OsuApiKey, bool)`

GetOsuApiKeyOk returns a tuple with the OsuApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuApiKey

`func (o *OsuExport) SetOsuApiKey(v OsuApiKey)`

SetOsuApiKey sets OsuApiKey field to given value.

### HasOsuApiKey

`func (o *OsuExport) HasOsuApiKey() bool`

HasOsuApiKey returns a boolean if a field has been set.

### GetOsuBucket

`func (o *OsuExport) GetOsuBucket() string`

GetOsuBucket returns the OsuBucket field if non-nil, zero value otherwise.

### GetOsuBucketOk

`func (o *OsuExport) GetOsuBucketOk() (*string, bool)`

GetOsuBucketOk returns a tuple with the OsuBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuBucket

`func (o *OsuExport) SetOsuBucket(v string)`

SetOsuBucket sets OsuBucket field to given value.


### GetOsuManifestUrl

`func (o *OsuExport) GetOsuManifestUrl() string`

GetOsuManifestUrl returns the OsuManifestUrl field if non-nil, zero value otherwise.

### GetOsuManifestUrlOk

`func (o *OsuExport) GetOsuManifestUrlOk() (*string, bool)`

GetOsuManifestUrlOk returns a tuple with the OsuManifestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuManifestUrl

`func (o *OsuExport) SetOsuManifestUrl(v string)`

SetOsuManifestUrl sets OsuManifestUrl field to given value.

### HasOsuManifestUrl

`func (o *OsuExport) HasOsuManifestUrl() bool`

HasOsuManifestUrl returns a boolean if a field has been set.

### GetOsuPrefix

`func (o *OsuExport) GetOsuPrefix() string`

GetOsuPrefix returns the OsuPrefix field if non-nil, zero value otherwise.

### GetOsuPrefixOk

`func (o *OsuExport) GetOsuPrefixOk() (*string, bool)`

GetOsuPrefixOk returns a tuple with the OsuPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuPrefix

`func (o *OsuExport) SetOsuPrefix(v string)`

SetOsuPrefix sets OsuPrefix field to given value.

### HasOsuPrefix

`func (o *OsuExport) HasOsuPrefix() bool`

HasOsuPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


