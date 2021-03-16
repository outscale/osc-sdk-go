# FiltersVmsState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceEventCodes** | **[]string** | The code for the scheduled event (&#x60;system-reboot&#x60; | &#x60;system-maintenance&#x60;). | [optional] 
**MaintenanceEventDescriptions** | **[]string** | The description of the scheduled event. | [optional] 
**MaintenanceEventsNotAfter** | **[]string** | The latest time the event can end. | [optional] 
**MaintenanceEventsNotBefore** | **[]string** | The earliest time the event can start. | [optional] 
**SubregionNames** | **[]string** | The names of the Subregions of the VMs. | [optional] 
**VmIds** | **[]string** | One or more IDs of VMs. | [optional] 
**VmStates** | **[]string** | The states of the VMs (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


