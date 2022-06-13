# DirectLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID of the owner of the DirectLink. | [optional] 
**Bandwidth** | **string** | The physical link bandwidth (either 1 Gbps or 10 Gbps). | [optional] 
**DirectLinkId** | **string** | The ID of the DirectLink (for example, &#x60;dxcon-xxxxxxxx&#x60;). | [optional] 
**DirectLinkName** | **string** | The name of the DirectLink. | [optional] 
**Location** | **string** | The datacenter where the DirectLink is located. | [optional] 
**RegionName** | **string** | The Region in which the DirectLink has been created. | [optional] 
**State** | **string** | The state of the DirectLink.&lt;br /&gt; * &#x60;requested&#x60;: The DirectLink is requested but the request has not been validated yet.&lt;br /&gt; * &#x60;pending&#x60;: The DirectLink request has been validated. It remains in the &#x60;pending&#x60; state until you establish the physical link.&lt;br /&gt; * &#x60;available&#x60;: The physical link is established and the connection is ready to use.&lt;br /&gt; * &#x60;deleting&#x60;: The deletion process is in progress.&lt;br /&gt; * &#x60;deleted&#x60;: The DirectLink is deleted. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


