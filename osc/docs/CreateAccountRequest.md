# CreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEmails** | **[]string** | One or more additional email addresses for the account. These addresses are used for notifications only. If you already have a list of additional emails registered, you cannot add to it, only replace it. To remove all registered additional emails, specify an empty list. | [optional] 
**City** | **string** | The city of the account owner. | 
**CompanyName** | **string** | The name of the company for the account. | 
**Country** | **string** | The country of the account owner. | 
**CustomerId** | **string** | The ID of the customer. It must be 8 digits. | 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Email** | **string** | The main email address for the account. This address is used for your credentials and notifications. | 
**FirstName** | **string** | The first name of the account owner. | 
**JobTitle** | **string** | The job title of the account owner. | [optional] 
**LastName** | **string** | The last name of the account owner. | 
**MobileNumber** | **string** | The mobile phone number of the account owner. | [optional] 
**PhoneNumber** | **string** | The landline phone number of the account owner. | [optional] 
**StateProvince** | **string** | The state/province of the account. | [optional] 
**VatNumber** | **string** | The value added tax (VAT) number for the account. | [optional] 
**ZipCode** | **string** | The ZIP code of the city. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


