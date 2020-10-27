# CreateSnapshotExportTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**OsuExport** | [**OsuExport**](OsuExport.md) |  | 
**SnapshotId** | **string** | The ID of the snapshot to export. | 

## Methods

### NewCreateSnapshotExportTaskRequest

`func NewCreateSnapshotExportTaskRequest(osuExport OsuExport, snapshotId string, ) *CreateSnapshotExportTaskRequest`

NewCreateSnapshotExportTaskRequest instantiates a new CreateSnapshotExportTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotExportTaskRequestWithDefaults

`func NewCreateSnapshotExportTaskRequestWithDefaults() *CreateSnapshotExportTaskRequest`

NewCreateSnapshotExportTaskRequestWithDefaults instantiates a new CreateSnapshotExportTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateSnapshotExportTaskRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSnapshotExportTaskRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateSnapshotExportTaskRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateSnapshotExportTaskRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetOsuExport

`func (o *CreateSnapshotExportTaskRequest) GetOsuExport() OsuExport`

GetOsuExport returns the OsuExport field if non-nil, zero value otherwise.

### GetOsuExportOk

`func (o *CreateSnapshotExportTaskRequest) GetOsuExportOk() (*OsuExport, bool)`

GetOsuExportOk returns a tuple with the OsuExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuExport

`func (o *CreateSnapshotExportTaskRequest) SetOsuExport(v OsuExport)`

SetOsuExport sets OsuExport field to given value.


### GetSnapshotId

`func (o *CreateSnapshotExportTaskRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CreateSnapshotExportTaskRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *CreateSnapshotExportTaskRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


