# ReadAdminPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**VmId** | **string** | The ID of the VM. | 

## Methods

### NewReadAdminPasswordRequest

`func NewReadAdminPasswordRequest(vmId string, ) *ReadAdminPasswordRequest`

NewReadAdminPasswordRequest instantiates a new ReadAdminPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadAdminPasswordRequestWithDefaults

`func NewReadAdminPasswordRequestWithDefaults() *ReadAdminPasswordRequest`

NewReadAdminPasswordRequestWithDefaults instantiates a new ReadAdminPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadAdminPasswordRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadAdminPasswordRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadAdminPasswordRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadAdminPasswordRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVmId

`func (o *ReadAdminPasswordRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *ReadAdminPasswordRequest) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *ReadAdminPasswordRequest) SetVmId(v string)`

SetVmId sets VmId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


