# FiltersNetPeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccepterNetAccountIds** | Pointer to **[]string** | The account IDs of the owners of the peer Nets. | [optional] 
**AccepterNetIpRanges** | Pointer to **[]string** | The IP ranges of the peer Nets, in CIDR notation (for example, 10.0.0.0/24). | [optional] 
**AccepterNetNetIds** | Pointer to **[]string** | The IDs of the peer Nets. | [optional] 
**NetPeeringIds** | Pointer to **[]string** | The IDs of the Net peering connections. | [optional] 
**SourceNetAccountIds** | Pointer to **[]string** | The account IDs of the owners of the peer Nets. | [optional] 
**SourceNetIpRanges** | Pointer to **[]string** | The IP ranges of the peer Nets. | [optional] 
**SourceNetNetIds** | Pointer to **[]string** | The IDs of the peer Nets. | [optional] 
**StateMessages** | Pointer to **[]string** | Additional information about the states of the Net peering connections. | [optional] 
**StateNames** | Pointer to **[]string** | The states of the Net peering connections (&#x60;pending-acceptance&#x60; \\| &#x60;active&#x60; \\| &#x60;rejected&#x60; \\| &#x60;failed&#x60; \\| &#x60;expired&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the Net peering connections. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the Net peering connections. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the Net peering connections, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 

## Methods

### NewFiltersNetPeering

`func NewFiltersNetPeering() *FiltersNetPeering`

NewFiltersNetPeering instantiates a new FiltersNetPeering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersNetPeeringWithDefaults

`func NewFiltersNetPeeringWithDefaults() *FiltersNetPeering`

NewFiltersNetPeeringWithDefaults instantiates a new FiltersNetPeering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepterNetAccountIds

`func (o *FiltersNetPeering) GetAccepterNetAccountIds() []string`

GetAccepterNetAccountIds returns the AccepterNetAccountIds field if non-nil, zero value otherwise.

### GetAccepterNetAccountIdsOk

`func (o *FiltersNetPeering) GetAccepterNetAccountIdsOk() (*[]string, bool)`

GetAccepterNetAccountIdsOk returns a tuple with the AccepterNetAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterNetAccountIds

`func (o *FiltersNetPeering) SetAccepterNetAccountIds(v []string)`

SetAccepterNetAccountIds sets AccepterNetAccountIds field to given value.

### HasAccepterNetAccountIds

`func (o *FiltersNetPeering) HasAccepterNetAccountIds() bool`

HasAccepterNetAccountIds returns a boolean if a field has been set.

### GetAccepterNetIpRanges

`func (o *FiltersNetPeering) GetAccepterNetIpRanges() []string`

GetAccepterNetIpRanges returns the AccepterNetIpRanges field if non-nil, zero value otherwise.

### GetAccepterNetIpRangesOk

`func (o *FiltersNetPeering) GetAccepterNetIpRangesOk() (*[]string, bool)`

GetAccepterNetIpRangesOk returns a tuple with the AccepterNetIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterNetIpRanges

`func (o *FiltersNetPeering) SetAccepterNetIpRanges(v []string)`

SetAccepterNetIpRanges sets AccepterNetIpRanges field to given value.

### HasAccepterNetIpRanges

`func (o *FiltersNetPeering) HasAccepterNetIpRanges() bool`

HasAccepterNetIpRanges returns a boolean if a field has been set.

### GetAccepterNetNetIds

`func (o *FiltersNetPeering) GetAccepterNetNetIds() []string`

GetAccepterNetNetIds returns the AccepterNetNetIds field if non-nil, zero value otherwise.

### GetAccepterNetNetIdsOk

`func (o *FiltersNetPeering) GetAccepterNetNetIdsOk() (*[]string, bool)`

GetAccepterNetNetIdsOk returns a tuple with the AccepterNetNetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterNetNetIds

`func (o *FiltersNetPeering) SetAccepterNetNetIds(v []string)`

SetAccepterNetNetIds sets AccepterNetNetIds field to given value.

### HasAccepterNetNetIds

`func (o *FiltersNetPeering) HasAccepterNetNetIds() bool`

HasAccepterNetNetIds returns a boolean if a field has been set.

### GetNetPeeringIds

`func (o *FiltersNetPeering) GetNetPeeringIds() []string`

GetNetPeeringIds returns the NetPeeringIds field if non-nil, zero value otherwise.

### GetNetPeeringIdsOk

`func (o *FiltersNetPeering) GetNetPeeringIdsOk() (*[]string, bool)`

GetNetPeeringIdsOk returns a tuple with the NetPeeringIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPeeringIds

`func (o *FiltersNetPeering) SetNetPeeringIds(v []string)`

SetNetPeeringIds sets NetPeeringIds field to given value.

### HasNetPeeringIds

`func (o *FiltersNetPeering) HasNetPeeringIds() bool`

HasNetPeeringIds returns a boolean if a field has been set.

### GetSourceNetAccountIds

`func (o *FiltersNetPeering) GetSourceNetAccountIds() []string`

GetSourceNetAccountIds returns the SourceNetAccountIds field if non-nil, zero value otherwise.

### GetSourceNetAccountIdsOk

`func (o *FiltersNetPeering) GetSourceNetAccountIdsOk() (*[]string, bool)`

GetSourceNetAccountIdsOk returns a tuple with the SourceNetAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetAccountIds

`func (o *FiltersNetPeering) SetSourceNetAccountIds(v []string)`

SetSourceNetAccountIds sets SourceNetAccountIds field to given value.

### HasSourceNetAccountIds

`func (o *FiltersNetPeering) HasSourceNetAccountIds() bool`

HasSourceNetAccountIds returns a boolean if a field has been set.

### GetSourceNetIpRanges

`func (o *FiltersNetPeering) GetSourceNetIpRanges() []string`

GetSourceNetIpRanges returns the SourceNetIpRanges field if non-nil, zero value otherwise.

### GetSourceNetIpRangesOk

`func (o *FiltersNetPeering) GetSourceNetIpRangesOk() (*[]string, bool)`

GetSourceNetIpRangesOk returns a tuple with the SourceNetIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetIpRanges

`func (o *FiltersNetPeering) SetSourceNetIpRanges(v []string)`

SetSourceNetIpRanges sets SourceNetIpRanges field to given value.

### HasSourceNetIpRanges

`func (o *FiltersNetPeering) HasSourceNetIpRanges() bool`

HasSourceNetIpRanges returns a boolean if a field has been set.

### GetSourceNetNetIds

`func (o *FiltersNetPeering) GetSourceNetNetIds() []string`

GetSourceNetNetIds returns the SourceNetNetIds field if non-nil, zero value otherwise.

### GetSourceNetNetIdsOk

`func (o *FiltersNetPeering) GetSourceNetNetIdsOk() (*[]string, bool)`

GetSourceNetNetIdsOk returns a tuple with the SourceNetNetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetNetIds

`func (o *FiltersNetPeering) SetSourceNetNetIds(v []string)`

SetSourceNetNetIds sets SourceNetNetIds field to given value.

### HasSourceNetNetIds

`func (o *FiltersNetPeering) HasSourceNetNetIds() bool`

HasSourceNetNetIds returns a boolean if a field has been set.

### GetStateMessages

`func (o *FiltersNetPeering) GetStateMessages() []string`

GetStateMessages returns the StateMessages field if non-nil, zero value otherwise.

### GetStateMessagesOk

`func (o *FiltersNetPeering) GetStateMessagesOk() (*[]string, bool)`

GetStateMessagesOk returns a tuple with the StateMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessages

`func (o *FiltersNetPeering) SetStateMessages(v []string)`

SetStateMessages sets StateMessages field to given value.

### HasStateMessages

`func (o *FiltersNetPeering) HasStateMessages() bool`

HasStateMessages returns a boolean if a field has been set.

### GetStateNames

`func (o *FiltersNetPeering) GetStateNames() []string`

GetStateNames returns the StateNames field if non-nil, zero value otherwise.

### GetStateNamesOk

`func (o *FiltersNetPeering) GetStateNamesOk() (*[]string, bool)`

GetStateNamesOk returns a tuple with the StateNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNames

`func (o *FiltersNetPeering) SetStateNames(v []string)`

SetStateNames sets StateNames field to given value.

### HasStateNames

`func (o *FiltersNetPeering) HasStateNames() bool`

HasStateNames returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersNetPeering) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersNetPeering) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersNetPeering) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersNetPeering) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersNetPeering) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersNetPeering) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersNetPeering) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersNetPeering) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersNetPeering) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersNetPeering) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersNetPeering) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersNetPeering) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


