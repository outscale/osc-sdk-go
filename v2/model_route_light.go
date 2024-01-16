/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.5
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// RouteLight Information about the route.
type RouteLight struct {
	// The IP range used for the destination match, in CIDR notation (for example, `10.0.0.0/24`).
	DestinationIpRange *string `json:"DestinationIpRange,omitempty"`
	// The type of route (always `static`).
	RouteType *string `json:"RouteType,omitempty"`
	// The current state of the static route (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
}

// NewRouteLight instantiates a new RouteLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteLight() *RouteLight {
	this := RouteLight{}
	return &this
}

// NewRouteLightWithDefaults instantiates a new RouteLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteLightWithDefaults() *RouteLight {
	this := RouteLight{}
	return &this
}

// GetDestinationIpRange returns the DestinationIpRange field value if set, zero value otherwise.
func (o *RouteLight) GetDestinationIpRange() string {
	if o == nil || o.DestinationIpRange == nil {
		var ret string
		return ret
	}
	return *o.DestinationIpRange
}

// GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetDestinationIpRangeOk() (*string, bool) {
	if o == nil || o.DestinationIpRange == nil {
		return nil, false
	}
	return o.DestinationIpRange, true
}

// HasDestinationIpRange returns a boolean if a field has been set.
func (o *RouteLight) HasDestinationIpRange() bool {
	if o != nil && o.DestinationIpRange != nil {
		return true
	}

	return false
}

// SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.
func (o *RouteLight) SetDestinationIpRange(v string) {
	o.DestinationIpRange = &v
}

// GetRouteType returns the RouteType field value if set, zero value otherwise.
func (o *RouteLight) GetRouteType() string {
	if o == nil || o.RouteType == nil {
		var ret string
		return ret
	}
	return *o.RouteType
}

// GetRouteTypeOk returns a tuple with the RouteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetRouteTypeOk() (*string, bool) {
	if o == nil || o.RouteType == nil {
		return nil, false
	}
	return o.RouteType, true
}

// HasRouteType returns a boolean if a field has been set.
func (o *RouteLight) HasRouteType() bool {
	if o != nil && o.RouteType != nil {
		return true
	}

	return false
}

// SetRouteType gets a reference to the given string and assigns it to the RouteType field.
func (o *RouteLight) SetRouteType(v string) {
	o.RouteType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RouteLight) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RouteLight) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RouteLight) SetState(v string) {
	o.State = &v
}

func (o RouteLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationIpRange != nil {
		toSerialize["DestinationIpRange"] = o.DestinationIpRange
	}
	if o.RouteType != nil {
		toSerialize["RouteType"] = o.RouteType
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableRouteLight struct {
	value *RouteLight
	isSet bool
}

func (v NullableRouteLight) Get() *RouteLight {
	return v.value
}

func (v *NullableRouteLight) Set(val *RouteLight) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteLight) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteLight(val *RouteLight) *NullableRouteLight {
	return &NullableRouteLight{value: val, isSet: true}
}

func (v NullableRouteLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
