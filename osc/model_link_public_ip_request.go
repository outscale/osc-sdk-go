/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// LinkPublicIpRequest struct for LinkPublicIpRequest
type LinkPublicIpRequest struct {
	// If true, allows the public IP to be associated with the VM or NIC that you specify even if it is already associated with another VM or NIC. If false, prevents the EIP from being associated with the VM or NIC that you specify if it is already associated with another VM or NIC. (By default, true in the public Cloud, false in a Net.)
	AllowRelink bool `json:"AllowRelink,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// (Net only) The ID of the NIC. This parameter is required if the VM has more than one NIC attached. Otherwise, you need to specify the `VmId` parameter instead. You cannot specify both parameters at the same time.
	NicId string `json:"NicId,omitempty"`
	// (Net only) The primary or secondary private IP of the specified NIC. By default, the primary private IP.
	PrivateIp string `json:"PrivateIp,omitempty"`
	// The public IP. This parameter is required unless you use the `PublicIpId` parameter.
	PublicIp string `json:"PublicIp,omitempty"`
	// The allocation ID of the public IP. This parameter is required unless you use the `PublicIp` parameter.
	PublicIpId string `json:"PublicIpId,omitempty"`
	// The ID of the VM.<br /> - In the public Cloud, this parameter is required.<br /> - In a Net, this parameter is required if the VM has only one NIC. Otherwise, you need to specify the `NicId` parameter instead. You cannot specify both parameters at the same time.
	VmId string `json:"VmId,omitempty"`
}
