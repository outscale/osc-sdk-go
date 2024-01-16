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

// CreateKeypairRequest struct for CreateKeypairRequest
type CreateKeypairRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// A unique name for the keypair, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).
	KeypairName string `json:"KeypairName"`
	// The public key to import in your account, if you are importing an existing keypair. This value must be Base64-encoded.
	PublicKey *string `json:"PublicKey,omitempty"`
}

// NewCreateKeypairRequest instantiates a new CreateKeypairRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeypairRequest(keypairName string) *CreateKeypairRequest {
	this := CreateKeypairRequest{}
	this.KeypairName = keypairName
	return &this
}

// NewCreateKeypairRequestWithDefaults instantiates a new CreateKeypairRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeypairRequestWithDefaults() *CreateKeypairRequest {
	this := CreateKeypairRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateKeypairRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKeypairRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateKeypairRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateKeypairRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetKeypairName returns the KeypairName field value
func (o *CreateKeypairRequest) GetKeypairName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value
// and a boolean to check if the value has been set.
func (o *CreateKeypairRequest) GetKeypairNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeypairName, true
}

// SetKeypairName sets field value
func (o *CreateKeypairRequest) SetKeypairName(v string) {
	o.KeypairName = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *CreateKeypairRequest) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKeypairRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CreateKeypairRequest) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *CreateKeypairRequest) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o CreateKeypairRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["KeypairName"] = o.KeypairName
	}
	if o.PublicKey != nil {
		toSerialize["PublicKey"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableCreateKeypairRequest struct {
	value *CreateKeypairRequest
	isSet bool
}

func (v NullableCreateKeypairRequest) Get() *CreateKeypairRequest {
	return v.value
}

func (v *NullableCreateKeypairRequest) Set(val *CreateKeypairRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeypairRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeypairRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeypairRequest(val *CreateKeypairRequest) *NullableCreateKeypairRequest {
	return &NullableCreateKeypairRequest{value: val, isSet: true}
}

func (v NullableCreateKeypairRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeypairRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
