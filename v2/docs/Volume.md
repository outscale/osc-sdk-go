# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **string** | The date and time of creation of the volume. | [optional] 
**Iops** | Pointer to **int32** | The number of I/O operations per second (IOPS):&lt;br /&gt; - For &#x60;io1&#x60; volumes, the number of provisioned IOPS&lt;br /&gt; - For &#x60;gp2&#x60; volumes, the baseline performance of the volume | [optional] 
**LinkedVolumes** | Pointer to [**[]LinkedVolume**](LinkedVolume.md) | Information about your volume attachment. | [optional] 
**Size** | Pointer to **int32** | The size of the volume, in gibibytes (GiB). | [optional] 
**SnapshotId** | Pointer to **string** | The snapshot from which the volume was created. | [optional] 
**State** | Pointer to **string** | The state of the volume (&#x60;creating&#x60; \\| &#x60;available&#x60; \\| &#x60;in-use&#x60; \\| &#x60;updating&#x60; \\| &#x60;deleting&#x60; \\| &#x60;error&#x60;). | [optional] 
**SubregionName** | Pointer to **string** | The Subregion in which the volume was created. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the volume. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume. | [optional] 
**VolumeType** | Pointer to **string** | The type of the volume (&#x60;standard&#x60; \\| &#x60;gp2&#x60; \\| &#x60;io1&#x60;). | [optional] 

## Methods

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *Volume) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Volume) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Volume) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Volume) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetIops

`func (o *Volume) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *Volume) GetIopsOk() (*int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *Volume) SetIops(v int32)`

SetIops sets Iops field to given value.

### HasIops

`func (o *Volume) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetLinkedVolumes

`func (o *Volume) GetLinkedVolumes() []LinkedVolume`

GetLinkedVolumes returns the LinkedVolumes field if non-nil, zero value otherwise.

### GetLinkedVolumesOk

`func (o *Volume) GetLinkedVolumesOk() (*[]LinkedVolume, bool)`

GetLinkedVolumesOk returns a tuple with the LinkedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedVolumes

`func (o *Volume) SetLinkedVolumes(v []LinkedVolume)`

SetLinkedVolumes sets LinkedVolumes field to given value.

### HasLinkedVolumes

`func (o *Volume) HasLinkedVolumes() bool`

HasLinkedVolumes returns a boolean if a field has been set.

### GetSize

`func (o *Volume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Volume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Volume) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Volume) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapshotId

`func (o *Volume) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *Volume) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *Volume) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *Volume) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetState

`func (o *Volume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Volume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Volume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Volume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubregionName

`func (o *Volume) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *Volume) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *Volume) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *Volume) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### GetTags

`func (o *Volume) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Volume) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Volume) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Volume) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumeId

`func (o *Volume) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *Volume) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *Volume) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *Volume) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetVolumeType

`func (o *Volume) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *Volume) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *Volume) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *Volume) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


