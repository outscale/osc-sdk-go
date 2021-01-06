# GenerateDataKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**EncryptionContext** | **map[string]string** | Information about the encryption context. | [optional] 
**GeneratePlaintext** | **bool** | If &#x60;true&#x60;, the data key is generated in plaintext and ciphertext. If &#x60;false&#x60;, the data key is generated only in ciphertext. | [optional] 
**MasterKeyId** | **string** | The ID of the master key used to generate the data key. | 
**Size** | **int32** | The size of the data key you want to generate, in bytes (between 1 and 1024). | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


