# UnlinkNicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LinkNicId** | **string** | The ID of the attachment operation. | 

## Methods

### NewUnlinkNicRequest

`func NewUnlinkNicRequest(linkNicId string, ) *UnlinkNicRequest`

NewUnlinkNicRequest instantiates a new UnlinkNicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkNicRequestWithDefaults

`func NewUnlinkNicRequestWithDefaults() *UnlinkNicRequest`

NewUnlinkNicRequestWithDefaults instantiates a new UnlinkNicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UnlinkNicRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkNicRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkNicRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkNicRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLinkNicId

`func (o *UnlinkNicRequest) GetLinkNicId() string`

GetLinkNicId returns the LinkNicId field if non-nil, zero value otherwise.

### GetLinkNicIdOk

`func (o *UnlinkNicRequest) GetLinkNicIdOk() (*string, bool)`

GetLinkNicIdOk returns a tuple with the LinkNicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicId

`func (o *UnlinkNicRequest) SetLinkNicId(v string)`

SetLinkNicId sets LinkNicId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


