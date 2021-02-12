# DeleteDirectLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinkId** | **string** | The ID of the DirectLink you want to delete. | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewDeleteDirectLinkRequest

`func NewDeleteDirectLinkRequest(directLinkId string, ) *DeleteDirectLinkRequest`

NewDeleteDirectLinkRequest instantiates a new DeleteDirectLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDirectLinkRequestWithDefaults

`func NewDeleteDirectLinkRequestWithDefaults() *DeleteDirectLinkRequest`

NewDeleteDirectLinkRequestWithDefaults instantiates a new DeleteDirectLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectLinkId

`func (o *DeleteDirectLinkRequest) GetDirectLinkId() string`

GetDirectLinkId returns the DirectLinkId field if non-nil, zero value otherwise.

### GetDirectLinkIdOk

`func (o *DeleteDirectLinkRequest) GetDirectLinkIdOk() (*string, bool)`

GetDirectLinkIdOk returns a tuple with the DirectLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkId

`func (o *DeleteDirectLinkRequest) SetDirectLinkId(v string)`

SetDirectLinkId sets DirectLinkId field to given value.


### GetDryRun

`func (o *DeleteDirectLinkRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteDirectLinkRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteDirectLinkRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteDirectLinkRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


