# InlinePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The policy document, corresponding to a JSON string that contains the policy. For more information, see [EIM Reference Information](https://docs.outscale.com/en/userguide/EIM-Reference-Information.html) and [EIM Policy Generator](https://docs.outscale.com/en/userguide/EIM-Policy-Generator.html). | [optional] 
**Name** | Pointer to **string** | The name of the policy. | [optional] 

## Methods

### NewInlinePolicy

`func NewInlinePolicy() *InlinePolicy`

NewInlinePolicy instantiates a new InlinePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlinePolicyWithDefaults

`func NewInlinePolicyWithDefaults() *InlinePolicy`

NewInlinePolicyWithDefaults instantiates a new InlinePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *InlinePolicy) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlinePolicy) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlinePolicy) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlinePolicy) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetName

`func (o *InlinePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlinePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlinePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlinePolicy) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


