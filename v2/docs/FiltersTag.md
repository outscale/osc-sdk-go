# FiltersTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** | The keys of the tags that are assigned to the resources. You can use this filter alongside the &#x60;Values&#x60; filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter. | [optional] 
**ResourceIds** | Pointer to **[]string** | The IDs of the resources with which the tags are associated. | [optional] 
**ResourceTypes** | Pointer to **[]string** | The resource type (&#x60;vm&#x60; \\| &#x60;image&#x60; \\| &#x60;volume&#x60; \\| &#x60;snapshot&#x60; \\| &#x60;public-ip&#x60; \\| &#x60;security-group&#x60; \\| &#x60;route-table&#x60; \\| &#x60;nic&#x60; \\| &#x60;net&#x60; \\| &#x60;subnet&#x60; \\| &#x60;net-peering&#x60; \\| &#x60;net-access-point&#x60; \\| &#x60;nat-service&#x60; \\| &#x60;internet-service&#x60; \\| &#x60;client-gateway&#x60; \\| &#x60;virtual-gateway&#x60; \\| &#x60;vpn-connection&#x60; \\| &#x60;dhcp-options&#x60; \\| &#x60;task&#x60;). | [optional] 
**Values** | Pointer to **[]string** | The values of the tags that are assigned to the resources. You can use this filter alongside the &#x60;TagKeys&#x60; filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter. | [optional] 

## Methods

### NewFiltersTag

`func NewFiltersTag() *FiltersTag`

NewFiltersTag instantiates a new FiltersTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersTagWithDefaults

`func NewFiltersTagWithDefaults() *FiltersTag`

NewFiltersTagWithDefaults instantiates a new FiltersTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *FiltersTag) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *FiltersTag) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *FiltersTag) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *FiltersTag) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetResourceIds

`func (o *FiltersTag) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *FiltersTag) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *FiltersTag) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *FiltersTag) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### GetResourceTypes

`func (o *FiltersTag) GetResourceTypes() []string`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *FiltersTag) GetResourceTypesOk() (*[]string, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *FiltersTag) SetResourceTypes(v []string)`

SetResourceTypes sets ResourceTypes field to given value.

### HasResourceTypes

`func (o *FiltersTag) HasResourceTypes() bool`

HasResourceTypes returns a boolean if a field has been set.

### GetValues

`func (o *FiltersTag) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FiltersTag) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FiltersTag) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *FiltersTag) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


