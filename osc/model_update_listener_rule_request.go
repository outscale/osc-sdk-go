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

// UpdateListenerRuleRequest struct for UpdateListenerRuleRequest
type UpdateListenerRuleRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].
	HostPattern *string `json:"HostPattern,omitempty"`
	// The name of the listener rule.
	ListenerRuleName string `json:"ListenerRuleName"`
	// A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].
	PathPattern *string `json:"PathPattern,omitempty"`
}
