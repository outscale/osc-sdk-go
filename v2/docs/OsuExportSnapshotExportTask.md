# OsuExportSnapshotExportTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskImageFormat** | **string** | The format of the export disk (&#x60;qcow2&#x60; \\| &#x60;raw&#x60;). | 
**OsuBucket** | **string** | The name of the OOS bucket the snapshot is exported to. | 
**OsuPrefix** | Pointer to **string** | The prefix for the key of the OOS object corresponding to the snapshot. | [optional] 

## Methods

### NewOsuExportSnapshotExportTask

`func NewOsuExportSnapshotExportTask(diskImageFormat string, osuBucket string, ) *OsuExportSnapshotExportTask`

NewOsuExportSnapshotExportTask instantiates a new OsuExportSnapshotExportTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsuExportSnapshotExportTaskWithDefaults

`func NewOsuExportSnapshotExportTaskWithDefaults() *OsuExportSnapshotExportTask`

NewOsuExportSnapshotExportTaskWithDefaults instantiates a new OsuExportSnapshotExportTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskImageFormat

`func (o *OsuExportSnapshotExportTask) GetDiskImageFormat() string`

GetDiskImageFormat returns the DiskImageFormat field if non-nil, zero value otherwise.

### GetDiskImageFormatOk

`func (o *OsuExportSnapshotExportTask) GetDiskImageFormatOk() (*string, bool)`

GetDiskImageFormatOk returns a tuple with the DiskImageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskImageFormat

`func (o *OsuExportSnapshotExportTask) SetDiskImageFormat(v string)`

SetDiskImageFormat sets DiskImageFormat field to given value.


### GetOsuBucket

`func (o *OsuExportSnapshotExportTask) GetOsuBucket() string`

GetOsuBucket returns the OsuBucket field if non-nil, zero value otherwise.

### GetOsuBucketOk

`func (o *OsuExportSnapshotExportTask) GetOsuBucketOk() (*string, bool)`

GetOsuBucketOk returns a tuple with the OsuBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuBucket

`func (o *OsuExportSnapshotExportTask) SetOsuBucket(v string)`

SetOsuBucket sets OsuBucket field to given value.


### GetOsuPrefix

`func (o *OsuExportSnapshotExportTask) GetOsuPrefix() string`

GetOsuPrefix returns the OsuPrefix field if non-nil, zero value otherwise.

### GetOsuPrefixOk

`func (o *OsuExportSnapshotExportTask) GetOsuPrefixOk() (*string, bool)`

GetOsuPrefixOk returns a tuple with the OsuPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuPrefix

`func (o *OsuExportSnapshotExportTask) SetOsuPrefix(v string)`

SetOsuPrefix sets OsuPrefix field to given value.

### HasOsuPrefix

`func (o *OsuExportSnapshotExportTask) HasOsuPrefix() bool`

HasOsuPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


