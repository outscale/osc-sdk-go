/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// CreateServerCertificateRequest struct for CreateServerCertificateRequest
type CreateServerCertificateRequest struct {
	// The PEM-encoded X509 certificate.
	Body string `json:"Body"`
	// The PEM-encoded intermediate certification authorities.
	Chain string `json:"Chain,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// A unique name for the certificate. Constraints: 1-128 alphanumeric characters, pluses (+), equals (=), commas (,), periods (.), at signs (@), minuses (-), or underscores (_).
	Name string `json:"Name"`
	// The path to the server certificate, set to a slash (/) if not specified.
	Path string `json:"Path,omitempty"`
	// The PEM-encoded private key matching the certificate.
	PrivateKey string `json:"PrivateKey"`
}
