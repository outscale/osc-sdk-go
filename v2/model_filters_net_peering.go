/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
	"time"
)

// FiltersNetPeering One or more filters.
type FiltersNetPeering struct {
	// The account IDs of the owners of the peer Nets.
	AccepterNetAccountIds *[]string `json:"AccepterNetAccountIds,omitempty"`
	// The IP ranges of the peer Nets, in CIDR notation (for example, `10.0.0.0/24`).
	AccepterNetIpRanges *[]string `json:"AccepterNetIpRanges,omitempty"`
	// The IDs of the peer Nets.
	AccepterNetNetIds *[]string `json:"AccepterNetNetIds,omitempty"`
	// The dates and times at which the Net peerings expire, in ISO 8601 date-time format (for example, `2020-06-14T00:00:00.000Z`).
	ExpirationDates *[]time.Time `json:"ExpirationDates,omitempty"`
	// The IDs of the Net peerings.
	NetPeeringIds *[]string `json:"NetPeeringIds,omitempty"`
	// The account IDs of the owners of the peer Nets.
	SourceNetAccountIds *[]string `json:"SourceNetAccountIds,omitempty"`
	// The IP ranges of the peer Nets.
	SourceNetIpRanges *[]string `json:"SourceNetIpRanges,omitempty"`
	// The IDs of the peer Nets.
	SourceNetNetIds *[]string `json:"SourceNetNetIds,omitempty"`
	// Additional information about the states of the Net peerings.
	StateMessages *[]string `json:"StateMessages,omitempty"`
	// The states of the Net peerings (`pending-acceptance` \\| `active` \\| `rejected` \\| `failed` \\| `expired` \\| `deleted`).
	StateNames *[]string `json:"StateNames,omitempty"`
	// The keys of the tags associated with the Net peerings.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Net peerings.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Net peerings, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersNetPeering instantiates a new FiltersNetPeering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersNetPeering() *FiltersNetPeering {
	this := FiltersNetPeering{}
	return &this
}

// NewFiltersNetPeeringWithDefaults instantiates a new FiltersNetPeering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersNetPeeringWithDefaults() *FiltersNetPeering {
	this := FiltersNetPeering{}
	return &this
}

// GetAccepterNetAccountIds returns the AccepterNetAccountIds field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetAccepterNetAccountIds() []string {
	if o == nil || o.AccepterNetAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccepterNetAccountIds
}

// GetAccepterNetAccountIdsOk returns a tuple with the AccepterNetAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetAccepterNetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccepterNetAccountIds == nil {
		return nil, false
	}
	return o.AccepterNetAccountIds, true
}

// HasAccepterNetAccountIds returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasAccepterNetAccountIds() bool {
	if o != nil && o.AccepterNetAccountIds != nil {
		return true
	}

	return false
}

// SetAccepterNetAccountIds gets a reference to the given []string and assigns it to the AccepterNetAccountIds field.
func (o *FiltersNetPeering) SetAccepterNetAccountIds(v []string) {
	o.AccepterNetAccountIds = &v
}

// GetAccepterNetIpRanges returns the AccepterNetIpRanges field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetAccepterNetIpRanges() []string {
	if o == nil || o.AccepterNetIpRanges == nil {
		var ret []string
		return ret
	}
	return *o.AccepterNetIpRanges
}

// GetAccepterNetIpRangesOk returns a tuple with the AccepterNetIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetAccepterNetIpRangesOk() (*[]string, bool) {
	if o == nil || o.AccepterNetIpRanges == nil {
		return nil, false
	}
	return o.AccepterNetIpRanges, true
}

// HasAccepterNetIpRanges returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasAccepterNetIpRanges() bool {
	if o != nil && o.AccepterNetIpRanges != nil {
		return true
	}

	return false
}

// SetAccepterNetIpRanges gets a reference to the given []string and assigns it to the AccepterNetIpRanges field.
func (o *FiltersNetPeering) SetAccepterNetIpRanges(v []string) {
	o.AccepterNetIpRanges = &v
}

// GetAccepterNetNetIds returns the AccepterNetNetIds field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetAccepterNetNetIds() []string {
	if o == nil || o.AccepterNetNetIds == nil {
		var ret []string
		return ret
	}
	return *o.AccepterNetNetIds
}

// GetAccepterNetNetIdsOk returns a tuple with the AccepterNetNetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetAccepterNetNetIdsOk() (*[]string, bool) {
	if o == nil || o.AccepterNetNetIds == nil {
		return nil, false
	}
	return o.AccepterNetNetIds, true
}

// HasAccepterNetNetIds returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasAccepterNetNetIds() bool {
	if o != nil && o.AccepterNetNetIds != nil {
		return true
	}

	return false
}

// SetAccepterNetNetIds gets a reference to the given []string and assigns it to the AccepterNetNetIds field.
func (o *FiltersNetPeering) SetAccepterNetNetIds(v []string) {
	o.AccepterNetNetIds = &v
}

// GetExpirationDates returns the ExpirationDates field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetExpirationDates() []time.Time {
	if o == nil || o.ExpirationDates == nil {
		var ret []time.Time
		return ret
	}
	return *o.ExpirationDates
}

// GetExpirationDatesOk returns a tuple with the ExpirationDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetExpirationDatesOk() (*[]time.Time, bool) {
	if o == nil || o.ExpirationDates == nil {
		return nil, false
	}
	return o.ExpirationDates, true
}

// HasExpirationDates returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasExpirationDates() bool {
	if o != nil && o.ExpirationDates != nil {
		return true
	}

	return false
}

// SetExpirationDates gets a reference to the given []time.Time and assigns it to the ExpirationDates field.
func (o *FiltersNetPeering) SetExpirationDates(v []time.Time) {
	o.ExpirationDates = &v
}

// GetNetPeeringIds returns the NetPeeringIds field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetNetPeeringIds() []string {
	if o == nil || o.NetPeeringIds == nil {
		var ret []string
		return ret
	}
	return *o.NetPeeringIds
}

// GetNetPeeringIdsOk returns a tuple with the NetPeeringIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetNetPeeringIdsOk() (*[]string, bool) {
	if o == nil || o.NetPeeringIds == nil {
		return nil, false
	}
	return o.NetPeeringIds, true
}

// HasNetPeeringIds returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasNetPeeringIds() bool {
	if o != nil && o.NetPeeringIds != nil {
		return true
	}

	return false
}

// SetNetPeeringIds gets a reference to the given []string and assigns it to the NetPeeringIds field.
func (o *FiltersNetPeering) SetNetPeeringIds(v []string) {
	o.NetPeeringIds = &v
}

// GetSourceNetAccountIds returns the SourceNetAccountIds field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetSourceNetAccountIds() []string {
	if o == nil || o.SourceNetAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.SourceNetAccountIds
}

// GetSourceNetAccountIdsOk returns a tuple with the SourceNetAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetSourceNetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.SourceNetAccountIds == nil {
		return nil, false
	}
	return o.SourceNetAccountIds, true
}

// HasSourceNetAccountIds returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasSourceNetAccountIds() bool {
	if o != nil && o.SourceNetAccountIds != nil {
		return true
	}

	return false
}

// SetSourceNetAccountIds gets a reference to the given []string and assigns it to the SourceNetAccountIds field.
func (o *FiltersNetPeering) SetSourceNetAccountIds(v []string) {
	o.SourceNetAccountIds = &v
}

// GetSourceNetIpRanges returns the SourceNetIpRanges field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetSourceNetIpRanges() []string {
	if o == nil || o.SourceNetIpRanges == nil {
		var ret []string
		return ret
	}
	return *o.SourceNetIpRanges
}

// GetSourceNetIpRangesOk returns a tuple with the SourceNetIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetSourceNetIpRangesOk() (*[]string, bool) {
	if o == nil || o.SourceNetIpRanges == nil {
		return nil, false
	}
	return o.SourceNetIpRanges, true
}

// HasSourceNetIpRanges returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasSourceNetIpRanges() bool {
	if o != nil && o.SourceNetIpRanges != nil {
		return true
	}

	return false
}

// SetSourceNetIpRanges gets a reference to the given []string and assigns it to the SourceNetIpRanges field.
func (o *FiltersNetPeering) SetSourceNetIpRanges(v []string) {
	o.SourceNetIpRanges = &v
}

// GetSourceNetNetIds returns the SourceNetNetIds field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetSourceNetNetIds() []string {
	if o == nil || o.SourceNetNetIds == nil {
		var ret []string
		return ret
	}
	return *o.SourceNetNetIds
}

// GetSourceNetNetIdsOk returns a tuple with the SourceNetNetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetSourceNetNetIdsOk() (*[]string, bool) {
	if o == nil || o.SourceNetNetIds == nil {
		return nil, false
	}
	return o.SourceNetNetIds, true
}

// HasSourceNetNetIds returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasSourceNetNetIds() bool {
	if o != nil && o.SourceNetNetIds != nil {
		return true
	}

	return false
}

// SetSourceNetNetIds gets a reference to the given []string and assigns it to the SourceNetNetIds field.
func (o *FiltersNetPeering) SetSourceNetNetIds(v []string) {
	o.SourceNetNetIds = &v
}

// GetStateMessages returns the StateMessages field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetStateMessages() []string {
	if o == nil || o.StateMessages == nil {
		var ret []string
		return ret
	}
	return *o.StateMessages
}

// GetStateMessagesOk returns a tuple with the StateMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetStateMessagesOk() (*[]string, bool) {
	if o == nil || o.StateMessages == nil {
		return nil, false
	}
	return o.StateMessages, true
}

// HasStateMessages returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasStateMessages() bool {
	if o != nil && o.StateMessages != nil {
		return true
	}

	return false
}

// SetStateMessages gets a reference to the given []string and assigns it to the StateMessages field.
func (o *FiltersNetPeering) SetStateMessages(v []string) {
	o.StateMessages = &v
}

// GetStateNames returns the StateNames field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetStateNames() []string {
	if o == nil || o.StateNames == nil {
		var ret []string
		return ret
	}
	return *o.StateNames
}

// GetStateNamesOk returns a tuple with the StateNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetStateNamesOk() (*[]string, bool) {
	if o == nil || o.StateNames == nil {
		return nil, false
	}
	return o.StateNames, true
}

// HasStateNames returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasStateNames() bool {
	if o != nil && o.StateNames != nil {
		return true
	}

	return false
}

// SetStateNames gets a reference to the given []string and assigns it to the StateNames field.
func (o *FiltersNetPeering) SetStateNames(v []string) {
	o.StateNames = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersNetPeering) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersNetPeering) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersNetPeering) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetPeering) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersNetPeering) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersNetPeering) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersNetPeering) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccepterNetAccountIds != nil {
		toSerialize["AccepterNetAccountIds"] = o.AccepterNetAccountIds
	}
	if o.AccepterNetIpRanges != nil {
		toSerialize["AccepterNetIpRanges"] = o.AccepterNetIpRanges
	}
	if o.AccepterNetNetIds != nil {
		toSerialize["AccepterNetNetIds"] = o.AccepterNetNetIds
	}
	if o.ExpirationDates != nil {
		toSerialize["ExpirationDates"] = o.ExpirationDates
	}
	if o.NetPeeringIds != nil {
		toSerialize["NetPeeringIds"] = o.NetPeeringIds
	}
	if o.SourceNetAccountIds != nil {
		toSerialize["SourceNetAccountIds"] = o.SourceNetAccountIds
	}
	if o.SourceNetIpRanges != nil {
		toSerialize["SourceNetIpRanges"] = o.SourceNetIpRanges
	}
	if o.SourceNetNetIds != nil {
		toSerialize["SourceNetNetIds"] = o.SourceNetNetIds
	}
	if o.StateMessages != nil {
		toSerialize["StateMessages"] = o.StateMessages
	}
	if o.StateNames != nil {
		toSerialize["StateNames"] = o.StateNames
	}
	if o.TagKeys != nil {
		toSerialize["TagKeys"] = o.TagKeys
	}
	if o.TagValues != nil {
		toSerialize["TagValues"] = o.TagValues
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersNetPeering struct {
	value *FiltersNetPeering
	isSet bool
}

func (v NullableFiltersNetPeering) Get() *FiltersNetPeering {
	return v.value
}

func (v *NullableFiltersNetPeering) Set(val *FiltersNetPeering) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersNetPeering) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersNetPeering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersNetPeering(val *FiltersNetPeering) *NullableFiltersNetPeering {
	return &NullableFiltersNetPeering{value: val, isSet: true}
}

func (v NullableFiltersNetPeering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersNetPeering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
