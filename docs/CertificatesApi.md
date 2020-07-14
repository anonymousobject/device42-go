# \CertificatesApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCertificates**](CertificatesApi.md#GetCertificates) | **Get** /api/1.0/certificates/ | Get certificates
[**PostUpdateCertificates**](CertificatesApi.md#PostUpdateCertificates) | **Post** /api/1.0/certificates/ | Create/Update certificates


# **GetCertificates**
> interface{} GetCertificates(ctx, optional)
Get certificates

This API is used to return the certificates you have added or discovered in Device42.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificatesApiGetCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiGetCertificatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateId** | **optional.String**| ID of the certificate | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdateCertificates**
> interface{} PostUpdateCertificates(ctx, optional)
Create/Update certificates

This API is used to update or add new certificates in bulk. Note: If you enter the DNS for the certificate, that is the only required field. Otherwise issued_to, valid_from, valid_to, and subject are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificatesApiPostUpdateCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiPostUpdateCertificatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dns** | **optional.String**| dns address of the site certificate is issued for. If entered, is the only required field | 
 **issuedTo** | **optional.String**| Entity certificate is issued to, required if no dns value | 
 **issuedBy** | **optional.String**| The entity issuing the certificate | 
 **validFrom** | **optional.String**| The start date of the certificate, required if no dns value. YYYY-MM-DD format | 
 **validTo** | **optional.String**| The expiration date of the certificate, required if no dns value. YYYY-MM-DD format | 
 **subject** | **optional.String**| The person, or entity identified. Required if no dns value | 
 **version** | **optional.String**| The version of the encoded certificate | 
 **serialNumber** | **optional.String**| Used to uniquely identify the certificate | 
 **signatureAlgorithm** | **optional.String**| The algorithm used to create the signature. | 
 **signatureHash** | **optional.String**| The actual signature to verify that it came from the issuer | 
 **digitalSignatureUsage** | **optional.String**| True when the subject public key is used for verifying digital signatures | 
 **contentCommitmentUsage** | **optional.String**|  | 
 **keyEnciphermentUsage** | **optional.String**| True when the subject public key is used for enciphering private or secret keys | 
 **dataEnciphermentUsage** | **optional.String**| True when the subject public key is used for directly enciphering raw user data without the use of an intermediate symmetric cipher. | 
 **keyAgreementUsage** | **optional.String**| True when the subject public key is used for key agreement. | 
 **keyCertSignUsage** | **optional.String**| True when the subject public key is used for verifying signatures on public key certificates. | 
 **crlSignUsage** | **optional.String**| True when the subject public key is used for verifying signatures on certificate revocation lists. | 
 **encipherOnlyUsage** | **optional.String**| When the encipherOnly bit is asserted and the keyAgreement bit is also set, the subject public key may be used only for enciphering data while performing key agreement. | 
 **decipherOnlyUsage** | **optional.String**| When the decipherOnly bit is asserted and the keyAgreement bit is also set, the subject public key may be used only for deciphering data while performing key agreement. | 
 **extendedKeyUsage** | **optional.String**| Indicates the purpose of the public key contained in the certificate. It contains a list of OIDs, each of which indicates an allowed use. | 
 **vendor** | **optional.String**| The name of the vendor that provided this certificate | 
 **endPointType** | **optional.String**|  | 
 **endPointId** | **optional.String**|  | 
 **groups** | **optional.String**| Use only if Multitenancy is on. List of groups separated by commas. Each group has name followed by colon followed by yes or no for change_permission, eg. CorpUS:yes, CorpNY:no. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

