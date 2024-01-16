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

// KeypairCreated Information about the created keypair.
type KeypairCreated struct {
	// The MD5 public key fingerprint, as specified in section 4 of RFC 4716.
	KeypairFingerprint *string `json:"KeypairFingerprint,omitempty"`
	// The name of the keypair.
	KeypairName *string `json:"KeypairName,omitempty"`
	// The type of the keypair (`ssh-rsa`, `ssh-ed25519`, `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, or `ecdsa-sha2-nistp521`).
	KeypairType *string `json:"KeypairType,omitempty"`
	// The private key, returned only if you are creating a keypair (not if you are importing). When you save this private key in a .rsa file, make sure you replace the `\\n` escape sequences with real line breaks.
	PrivateKey *string `json:"PrivateKey,omitempty"`
}

// NewKeypairCreated instantiates a new KeypairCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeypairCreated() *KeypairCreated {
	this := KeypairCreated{}
	return &this
}

// NewKeypairCreatedWithDefaults instantiates a new KeypairCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeypairCreatedWithDefaults() *KeypairCreated {
	this := KeypairCreated{}
	return &this
}

// GetKeypairFingerprint returns the KeypairFingerprint field value if set, zero value otherwise.
func (o *KeypairCreated) GetKeypairFingerprint() string {
	if o == nil || o.KeypairFingerprint == nil {
		var ret string
		return ret
	}
	return *o.KeypairFingerprint
}

// GetKeypairFingerprintOk returns a tuple with the KeypairFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeypairCreated) GetKeypairFingerprintOk() (*string, bool) {
	if o == nil || o.KeypairFingerprint == nil {
		return nil, false
	}
	return o.KeypairFingerprint, true
}

// HasKeypairFingerprint returns a boolean if a field has been set.
func (o *KeypairCreated) HasKeypairFingerprint() bool {
	if o != nil && o.KeypairFingerprint != nil {
		return true
	}

	return false
}

// SetKeypairFingerprint gets a reference to the given string and assigns it to the KeypairFingerprint field.
func (o *KeypairCreated) SetKeypairFingerprint(v string) {
	o.KeypairFingerprint = &v
}

// GetKeypairName returns the KeypairName field value if set, zero value otherwise.
func (o *KeypairCreated) GetKeypairName() string {
	if o == nil || o.KeypairName == nil {
		var ret string
		return ret
	}
	return *o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeypairCreated) GetKeypairNameOk() (*string, bool) {
	if o == nil || o.KeypairName == nil {
		return nil, false
	}
	return o.KeypairName, true
}

// HasKeypairName returns a boolean if a field has been set.
func (o *KeypairCreated) HasKeypairName() bool {
	if o != nil && o.KeypairName != nil {
		return true
	}

	return false
}

// SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.
func (o *KeypairCreated) SetKeypairName(v string) {
	o.KeypairName = &v
}

// GetKeypairType returns the KeypairType field value if set, zero value otherwise.
func (o *KeypairCreated) GetKeypairType() string {
	if o == nil || o.KeypairType == nil {
		var ret string
		return ret
	}
	return *o.KeypairType
}

// GetKeypairTypeOk returns a tuple with the KeypairType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeypairCreated) GetKeypairTypeOk() (*string, bool) {
	if o == nil || o.KeypairType == nil {
		return nil, false
	}
	return o.KeypairType, true
}

// HasKeypairType returns a boolean if a field has been set.
func (o *KeypairCreated) HasKeypairType() bool {
	if o != nil && o.KeypairType != nil {
		return true
	}

	return false
}

// SetKeypairType gets a reference to the given string and assigns it to the KeypairType field.
func (o *KeypairCreated) SetKeypairType(v string) {
	o.KeypairType = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *KeypairCreated) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeypairCreated) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *KeypairCreated) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *KeypairCreated) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

func (o KeypairCreated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeypairFingerprint != nil {
		toSerialize["KeypairFingerprint"] = o.KeypairFingerprint
	}
	if o.KeypairName != nil {
		toSerialize["KeypairName"] = o.KeypairName
	}
	if o.KeypairType != nil {
		toSerialize["KeypairType"] = o.KeypairType
	}
	if o.PrivateKey != nil {
		toSerialize["PrivateKey"] = o.PrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullableKeypairCreated struct {
	value *KeypairCreated
	isSet bool
}

func (v NullableKeypairCreated) Get() *KeypairCreated {
	return v.value
}

func (v *NullableKeypairCreated) Set(val *KeypairCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableKeypairCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableKeypairCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeypairCreated(val *KeypairCreated) *NullableKeypairCreated {
	return &NullableKeypairCreated{value: val, isSet: true}
}

func (v NullableKeypairCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeypairCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
