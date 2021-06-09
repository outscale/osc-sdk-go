# Phase1Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpdTimeoutAction** | **string** | The action to carry out after a Dead Peer Detection (DPD) timeout occurs. | [optional] 
**DpdTimeoutSeconds** | **int32** | The maximum waiting time for a Dead Peer Detection (DPD) response before considering the peer as dead, in seconds. | [optional] 
**IkeVersions** | **[]string** | The Internet Key Exchange (IKE) versions allowed for the VPN tunnel. | [optional] 
**Phase1DhGroupNumbers** | **[]int32** | The Diffie-Hellman (DH) group numbers allowed for the VPN tunnel for phase 1. | [optional] 
**Phase1EncryptionAlgorithms** | **[]string** | The encryption algorithms allowed for the VPN tunnel for phase 1. | [optional] 
**Phase1IntegrityAlgorithms** | **[]string** | The integrity algorithms allowed for the VPN tunnel for phase 1. | [optional] 
**Phase1LifetimeSeconds** | **int32** | The lifetime for phase 1 of the IKE negotiation process, in seconds. | [optional] 
**ReplayWindowSize** | **int32** | The number of packets in an IKE replay window. | [optional] 
**StartupAction** | **string** | The action to carry out when establishing tunnels for a VPN connection. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


