# CreateServerCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The PEM-encoded X509 certificate. | 
**Chain** | **string** | The PEM-encoded intermediate certification authorities. | [optional] 
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Name** | **string** | A unique name for the certificate. Constraints: 1-128 alphanumeric characters, pluses (+), equals (&#x3D;), commas (,), periods (.), at signs (@), minuses (-), or underscores (_). | 
**Path** | **string** | The path to the server certificate, set to a slash (/) if not specified. | [optional] 
**PrivateKey** | **string** | The PEM-encoded private key matching the certificate. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


