# DeleteExportTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ExportTaskId** | **string** | The ID of the export task to delete. | 

## Methods

### NewDeleteExportTaskRequest

`func NewDeleteExportTaskRequest(exportTaskId string, ) *DeleteExportTaskRequest`

NewDeleteExportTaskRequest instantiates a new DeleteExportTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteExportTaskRequestWithDefaults

`func NewDeleteExportTaskRequestWithDefaults() *DeleteExportTaskRequest`

NewDeleteExportTaskRequestWithDefaults instantiates a new DeleteExportTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteExportTaskRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteExportTaskRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteExportTaskRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteExportTaskRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetExportTaskId

`func (o *DeleteExportTaskRequest) GetExportTaskId() string`

GetExportTaskId returns the ExportTaskId field if non-nil, zero value otherwise.

### GetExportTaskIdOk

`func (o *DeleteExportTaskRequest) GetExportTaskIdOk() (*string, bool)`

GetExportTaskIdOk returns a tuple with the ExportTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTaskId

`func (o *DeleteExportTaskRequest) SetExportTaskId(v string)`

SetExportTaskId sets ExportTaskId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


