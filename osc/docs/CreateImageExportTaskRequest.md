# CreateImageExportTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | **string** | The ID of the OMI to export. | 
**OsuExport** | [**OsuExport**](OsuExport.md) |  | 

## Methods

### NewCreateImageExportTaskRequest

`func NewCreateImageExportTaskRequest(imageId string, osuExport OsuExport, ) *CreateImageExportTaskRequest`

NewCreateImageExportTaskRequest instantiates a new CreateImageExportTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageExportTaskRequestWithDefaults

`func NewCreateImageExportTaskRequestWithDefaults() *CreateImageExportTaskRequest`

NewCreateImageExportTaskRequestWithDefaults instantiates a new CreateImageExportTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateImageExportTaskRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateImageExportTaskRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateImageExportTaskRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateImageExportTaskRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetImageId

`func (o *CreateImageExportTaskRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateImageExportTaskRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateImageExportTaskRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetOsuExport

`func (o *CreateImageExportTaskRequest) GetOsuExport() OsuExport`

GetOsuExport returns the OsuExport field if non-nil, zero value otherwise.

### GetOsuExportOk

`func (o *CreateImageExportTaskRequest) GetOsuExportOk() (*OsuExport, bool)`

GetOsuExportOk returns a tuple with the OsuExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuExport

`func (o *CreateImageExportTaskRequest) SetOsuExport(v OsuExport)`

SetOsuExport sets OsuExport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


