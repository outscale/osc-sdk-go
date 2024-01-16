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

// BlockDeviceMappingImage One or more parameters used to automatically set up volumes when the VM is created.
type BlockDeviceMappingImage struct {
	Bsu *BsuToCreate `json:"Bsu,omitempty"`
	// The device name for the volume. For a root device, you must use `/dev/sda1`. For other volumes, you must use `/dev/sdX`, `/dev/sdXX`, `/dev/xvdX`, or `/dev/xvdXX` (where the first `X` is a letter between `b` and `z`, and the second `X` is a letter between `a` and `z`).
	DeviceName *string `json:"DeviceName,omitempty"`
	// The name of the virtual device (`ephemeralN`).
	VirtualDeviceName *string `json:"VirtualDeviceName,omitempty"`
}

// NewBlockDeviceMappingImage instantiates a new BlockDeviceMappingImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockDeviceMappingImage() *BlockDeviceMappingImage {
	this := BlockDeviceMappingImage{}
	return &this
}

// NewBlockDeviceMappingImageWithDefaults instantiates a new BlockDeviceMappingImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockDeviceMappingImageWithDefaults() *BlockDeviceMappingImage {
	this := BlockDeviceMappingImage{}
	return &this
}

// GetBsu returns the Bsu field value if set, zero value otherwise.
func (o *BlockDeviceMappingImage) GetBsu() BsuToCreate {
	if o == nil || o.Bsu == nil {
		var ret BsuToCreate
		return ret
	}
	return *o.Bsu
}

// GetBsuOk returns a tuple with the Bsu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingImage) GetBsuOk() (*BsuToCreate, bool) {
	if o == nil || o.Bsu == nil {
		return nil, false
	}
	return o.Bsu, true
}

// HasBsu returns a boolean if a field has been set.
func (o *BlockDeviceMappingImage) HasBsu() bool {
	if o != nil && o.Bsu != nil {
		return true
	}

	return false
}

// SetBsu gets a reference to the given BsuToCreate and assigns it to the Bsu field.
func (o *BlockDeviceMappingImage) SetBsu(v BsuToCreate) {
	o.Bsu = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingImage) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingImage) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingImage) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *BlockDeviceMappingImage) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetVirtualDeviceName returns the VirtualDeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingImage) GetVirtualDeviceName() string {
	if o == nil || o.VirtualDeviceName == nil {
		var ret string
		return ret
	}
	return *o.VirtualDeviceName
}

// GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingImage) GetVirtualDeviceNameOk() (*string, bool) {
	if o == nil || o.VirtualDeviceName == nil {
		return nil, false
	}
	return o.VirtualDeviceName, true
}

// HasVirtualDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingImage) HasVirtualDeviceName() bool {
	if o != nil && o.VirtualDeviceName != nil {
		return true
	}

	return false
}

// SetVirtualDeviceName gets a reference to the given string and assigns it to the VirtualDeviceName field.
func (o *BlockDeviceMappingImage) SetVirtualDeviceName(v string) {
	o.VirtualDeviceName = &v
}

func (o BlockDeviceMappingImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bsu != nil {
		toSerialize["Bsu"] = o.Bsu
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.VirtualDeviceName != nil {
		toSerialize["VirtualDeviceName"] = o.VirtualDeviceName
	}
	return json.Marshal(toSerialize)
}

type NullableBlockDeviceMappingImage struct {
	value *BlockDeviceMappingImage
	isSet bool
}

func (v NullableBlockDeviceMappingImage) Get() *BlockDeviceMappingImage {
	return v.value
}

func (v *NullableBlockDeviceMappingImage) Set(val *BlockDeviceMappingImage) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockDeviceMappingImage) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockDeviceMappingImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockDeviceMappingImage(val *BlockDeviceMappingImage) *NullableBlockDeviceMappingImage {
	return &NullableBlockDeviceMappingImage{value: val, isSet: true}
}

func (v NullableBlockDeviceMappingImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockDeviceMappingImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
