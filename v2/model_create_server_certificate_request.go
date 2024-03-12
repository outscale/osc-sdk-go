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
)

// CreateServerCertificateRequest struct for CreateServerCertificateRequest
type CreateServerCertificateRequest struct {
	// The PEM-encoded X509 certificate.<br />With OSC CLI, use the following syntax to make sure your CA file is correctly parsed: `--CaPem=&quot;$(cat FILENAME)&quot;`.
	Body string `json:"Body"`
	// The PEM-encoded intermediate certification authorities.<br />With OSC CLI, use the following syntax to make sure your CA file is correctly parsed: `--CaPem=&quot;$(cat FILENAME)&quot;`.
	Chain *string `json:"Chain,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// A unique name for the certificate. Constraints: 1-128 alphanumeric characters, pluses (+), equals (=), commas (,), periods (.), at signs (@), minuses (-), or underscores (_).
	Name string `json:"Name"`
	// The path to the server certificate, set to a slash (/) if not specified.
	Path *string `json:"Path,omitempty"`
	// The PEM-encoded private key matching the certificate.<br />With OSC CLI, use the following syntax to make sure your CA file is correctly parsed: `--CaPem=&quot;$(cat FILENAME)&quot;`.
	PrivateKey string `json:"PrivateKey"`
}

// NewCreateServerCertificateRequest instantiates a new CreateServerCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServerCertificateRequest(body string, name string, privateKey string) *CreateServerCertificateRequest {
	this := CreateServerCertificateRequest{}
	this.Body = body
	this.Name = name
	this.PrivateKey = privateKey
	return &this
}

// NewCreateServerCertificateRequestWithDefaults instantiates a new CreateServerCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServerCertificateRequestWithDefaults() *CreateServerCertificateRequest {
	this := CreateServerCertificateRequest{}
	return &this
}

// GetBody returns the Body field value
func (o *CreateServerCertificateRequest) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *CreateServerCertificateRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *CreateServerCertificateRequest) SetBody(v string) {
	o.Body = v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *CreateServerCertificateRequest) GetChain() string {
	if o == nil || o.Chain == nil {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerCertificateRequest) GetChainOk() (*string, bool) {
	if o == nil || o.Chain == nil {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *CreateServerCertificateRequest) HasChain() bool {
	if o != nil && o.Chain != nil {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *CreateServerCertificateRequest) SetChain(v string) {
	o.Chain = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateServerCertificateRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerCertificateRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateServerCertificateRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateServerCertificateRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetName returns the Name field value
func (o *CreateServerCertificateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateServerCertificateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateServerCertificateRequest) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CreateServerCertificateRequest) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerCertificateRequest) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CreateServerCertificateRequest) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CreateServerCertificateRequest) SetPath(v string) {
	o.Path = &v
}

// GetPrivateKey returns the PrivateKey field value
func (o *CreateServerCertificateRequest) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *CreateServerCertificateRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *CreateServerCertificateRequest) SetPrivateKey(v string) {
	o.PrivateKey = v
}

func (o CreateServerCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Body"] = o.Body
	}
	if o.Chain != nil {
		toSerialize["Chain"] = o.Chain
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if true {
		toSerialize["PrivateKey"] = o.PrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullableCreateServerCertificateRequest struct {
	value *CreateServerCertificateRequest
	isSet bool
}

func (v NullableCreateServerCertificateRequest) Get() *CreateServerCertificateRequest {
	return v.value
}

func (v *NullableCreateServerCertificateRequest) Set(val *CreateServerCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServerCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServerCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServerCertificateRequest(val *CreateServerCertificateRequest) *NullableCreateServerCertificateRequest {
	return &NullableCreateServerCertificateRequest{value: val, isSet: true}
}

func (v NullableCreateServerCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServerCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
