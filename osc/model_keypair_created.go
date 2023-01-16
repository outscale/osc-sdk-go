/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// KeypairCreated Information about the created keypair.
type KeypairCreated struct {
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	KeypairFingerprint string `json:"KeypairFingerprint,omitempty"`
	// The name of the keypair.
	KeypairName string `json:"KeypairName,omitempty"`
	// The private key. When saving the private key in a .rsa file, replace the `\\n` escape sequences with line breaks.
	PrivateKey string `json:"PrivateKey,omitempty"`
}
