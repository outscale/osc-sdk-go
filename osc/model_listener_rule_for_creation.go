/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// ListenerRuleForCreation Information about the listener rule.
type ListenerRuleForCreation struct {
	// The type of action for the rule (always `forward`).
	Action string `json:"Action,omitempty"`
	// A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].
	HostNamePattern string `json:"HostNamePattern,omitempty"`
	// A human-readable name for the listener rule.
	ListenerRuleName string `json:"ListenerRuleName"`
	// A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].
	PathPattern string `json:"PathPattern,omitempty"`
	// The priority level of the listener rule, between `1` and `19999` both included. Each rule must have a unique priority level. Otherwise, an error is returned.
	Priority int32 `json:"Priority"`
}
