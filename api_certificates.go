
/*
 * Device42 API
 *
 * <b><h2>API Overview</h2></b><br>Restful API are supported in Device42 as one of the primary methods of entering, editing and retrieving data.<br><br> The API documentation has samples in curl for now and we will be adding other programming languages in future. You can refer to our github repositories at https://github.com/device42 for some sample code in python.<br><br> The RESTful API's enforce the role-based security that is created with the Device42 app. If you want a user to have access via the API, but not via the UI - deselect 'Staff Status' for that user from UI Tools > Admins & Permissions > Administrators.<br><br>If you'd like to see your changes on our UI, please visit http://swaggerdemo.device42.com and use credentials 'guest:device42_rocks!'. Please be aware that every 30 minutes our database resets and all data entered with POST/PUT calls from this page will be lost.<br><br>If you have any questions/suggestions, please send us a note to support at device42.com below. <br><h2><a id = 'Sample_API_Code'><b>Sample API Code</b></a></h2> <b>Sample Code with API Calls in the Python and .Net Programming Languages.</b><br>The following programs written in the Python programming language are available at <a href= 'https://github.com/device42/Device42-AutoDiscovery-Scripts'>Device42 sample programs</a><br><ul><li>api-sample.py: Runs against a single Windows system and uploads info to device42 appliance.</li><li>ad-sample.py: Can run against Active directory computers, servers or a given list and upload discovered systems’ info to device42 appliance.</li><li>d42_api_linux_upload_sample_script.py: Runs on a single *nix based system and uploads info to device42 appliance.</li><li>sample-script-facter-facts-to-d42: Runs on puppet master and uploads nodes info from facter fact files to device42 appliance.</li><li>d42_api_solaris_sample_script.py: Runs on an individual solaris system and uploads info to device42 appliance.</li></ul>The csv2d42apis.py sample program shows how to create a CSV file of data to import into device42. It reads a CSV file, matches columns to arguments for APIs and sends data to device42 via POST or PUT. This program can be found at: <a href= 'https://github.com/device42/API_Helpers'> Device42 sample CSV importer</a>.<br><br>The Auto Discovery Client source code can be found at Device42 <a href='https://docs.device42.com/auto-discovery/'>Auto Discovery Client</a>. This is a .net program that uses the device42 Restful API’s to load the discovered data.<h2><a id = 'Response_Status_Messages'><b>Response & Status Messages</b></a></h2><b>Responses</b><br>The response to most POST calls will have the following format: <br>{'msg':['PartModel added/updated', 17, 'RAM Acme 123456', true, true], 'code':0}<br><br>Using the example above:<br>'17' = ID of the object<br>'RAM Acme 123456' = representation of the object<br> 'true, true' = object added and is new.<br>'true, false' = object updated.<br>'false, false' = no changes.<br><br><b>Status Codes</b><br><u>Code 0</u><br>Success! e.g. = {'msg': 'device added or updated', <b>code:0</b>}<br><br><u>'HTTP/1.1 200 OK'</u><br>Success! All other responses are errors and will display an error code e.g.: 400, 401, 403, 405, 410, 500, 503. These are explained in each call.<br><br>Please let us know if you notice something odd with a response and we will fix it! :) <h2><a id = 'Get_Limits_Offsets'><b>API Get Limits and Offsets</b></a></h2>In Device42’s global settings we have the option to enforce API GET limits and this setting is recommended for better performance. To set this value from Device42, go to Tools>Settings>Global Settings and the click Edit in the top right. Toward the bottom of the screen you will see the API section. First check the “Enforce API GET Limits” checkbox, then enter a value for API GET Limit. A good limit to start with is 500 or 1000, depending on the performance you experience.<br><br>After setting this limit all GET calls below can have the results augmented via “offset” and “limit” parameters to tell Device42 which value to start returning results from, and how many results to return at maximum.<br>Examples:<br><br>'/?limit=50' returns the first 50 results.<br><br> '/?offset=100' will start at the 101st result. <br><br>'/?offset=42&limit=42' returns 42 results, starting at the 43rd.<br><br> Note that offset is used for paging - that is, the offset is only applied when the total number of objects returned exceeds the limit that is returned.
 *
 * API version: 2.0
 * Contact: support@device42.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type CertificatesApiService service

/* 
CertificatesApiService Get certificates
This API is used to return the certificates you have added or discovered in Device42.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CertificatesApiGetCertificatesOpts - Optional Parameters:
     * @param "CertificateId" (optional.String) -  ID of the certificate

@return interface{}
*/

type CertificatesApiGetCertificatesOpts struct { 
	CertificateId optional.String
}

func (a *CertificatesApiService) GetCertificates(ctx context.Context, localVarOptionals *CertificatesApiGetCertificatesOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/certificates/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.CertificateId.IsSet() {
		localVarQueryParams.Add("certificate_id", parameterToString(localVarOptionals.CertificateId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
CertificatesApiService Create/Update certificates
This API is used to update or add new certificates in bulk. Note: If you enter the DNS for the certificate, that is the only required field. Otherwise issued_to, valid_from, valid_to, and subject are required.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CertificatesApiPostUpdateCertificatesOpts - Optional Parameters:
     * @param "Dns" (optional.String) -  dns address of the site certificate is issued for. If entered, is the only required field
     * @param "IssuedTo" (optional.String) -  Entity certificate is issued to, required if no dns value
     * @param "IssuedBy" (optional.String) -  The entity issuing the certificate
     * @param "ValidFrom" (optional.String) -  The start date of the certificate, required if no dns value. YYYY-MM-DD format
     * @param "ValidTo" (optional.String) -  The expiration date of the certificate, required if no dns value. YYYY-MM-DD format
     * @param "Subject" (optional.String) -  The person, or entity identified. Required if no dns value
     * @param "Version" (optional.String) -  The version of the encoded certificate
     * @param "SerialNumber" (optional.String) -  Used to uniquely identify the certificate
     * @param "SignatureAlgorithm" (optional.String) -  The algorithm used to create the signature.
     * @param "SignatureHash" (optional.String) -  The actual signature to verify that it came from the issuer
     * @param "DigitalSignatureUsage" (optional.String) -  True when the subject public key is used for verifying digital signatures
     * @param "ContentCommitmentUsage" (optional.String) - 
     * @param "KeyEnciphermentUsage" (optional.String) -  True when the subject public key is used for enciphering private or secret keys
     * @param "DataEnciphermentUsage" (optional.String) -  True when the subject public key is used for directly enciphering raw user data without the use of an intermediate symmetric cipher.
     * @param "KeyAgreementUsage" (optional.String) -  True when the subject public key is used for key agreement.
     * @param "KeyCertSignUsage" (optional.String) -  True when the subject public key is used for verifying signatures on public key certificates.
     * @param "CrlSignUsage" (optional.String) -  True when the subject public key is used for verifying signatures on certificate revocation lists.
     * @param "EncipherOnlyUsage" (optional.String) -  When the encipherOnly bit is asserted and the keyAgreement bit is also set, the subject public key may be used only for enciphering data while performing key agreement.
     * @param "DecipherOnlyUsage" (optional.String) -  When the decipherOnly bit is asserted and the keyAgreement bit is also set, the subject public key may be used only for deciphering data while performing key agreement.
     * @param "ExtendedKeyUsage" (optional.String) -  Indicates the purpose of the public key contained in the certificate. It contains a list of OIDs, each of which indicates an allowed use.
     * @param "Vendor" (optional.String) -  The name of the vendor that provided this certificate
     * @param "EndPointType" (optional.String) - 
     * @param "EndPointId" (optional.String) - 
     * @param "Groups" (optional.String) -  Use only if Multitenancy is on. List of groups separated by commas. Each group has name followed by colon followed by yes or no for change_permission, eg. CorpUS:yes, CorpNY:no.

@return interface{}
*/

type CertificatesApiPostUpdateCertificatesOpts struct { 
	Dns optional.String
	IssuedTo optional.String
	IssuedBy optional.String
	ValidFrom optional.String
	ValidTo optional.String
	Subject optional.String
	Version optional.String
	SerialNumber optional.String
	SignatureAlgorithm optional.String
	SignatureHash optional.String
	DigitalSignatureUsage optional.String
	ContentCommitmentUsage optional.String
	KeyEnciphermentUsage optional.String
	DataEnciphermentUsage optional.String
	KeyAgreementUsage optional.String
	KeyCertSignUsage optional.String
	CrlSignUsage optional.String
	EncipherOnlyUsage optional.String
	DecipherOnlyUsage optional.String
	ExtendedKeyUsage optional.String
	Vendor optional.String
	EndPointType optional.String
	EndPointId optional.String
	Groups optional.String
}

func (a *CertificatesApiService) PostUpdateCertificates(ctx context.Context, localVarOptionals *CertificatesApiPostUpdateCertificatesOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/certificates/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Dns.IsSet() {
		localVarFormParams.Add("dns", parameterToString(localVarOptionals.Dns.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IssuedTo.IsSet() {
		localVarFormParams.Add("issued_to", parameterToString(localVarOptionals.IssuedTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IssuedBy.IsSet() {
		localVarFormParams.Add("issued_by", parameterToString(localVarOptionals.IssuedBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ValidFrom.IsSet() {
		localVarFormParams.Add("valid_from", parameterToString(localVarOptionals.ValidFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ValidTo.IsSet() {
		localVarFormParams.Add("valid_to", parameterToString(localVarOptionals.ValidTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Subject.IsSet() {
		localVarFormParams.Add("subject", parameterToString(localVarOptionals.Subject.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Version.IsSet() {
		localVarFormParams.Add("version", parameterToString(localVarOptionals.Version.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SerialNumber.IsSet() {
		localVarFormParams.Add("serial_number", parameterToString(localVarOptionals.SerialNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SignatureAlgorithm.IsSet() {
		localVarFormParams.Add("signature_algorithm", parameterToString(localVarOptionals.SignatureAlgorithm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SignatureHash.IsSet() {
		localVarFormParams.Add("signature_hash", parameterToString(localVarOptionals.SignatureHash.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DigitalSignatureUsage.IsSet() {
		localVarFormParams.Add("digital_signature_usage", parameterToString(localVarOptionals.DigitalSignatureUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ContentCommitmentUsage.IsSet() {
		localVarFormParams.Add("content_commitment_usage", parameterToString(localVarOptionals.ContentCommitmentUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.KeyEnciphermentUsage.IsSet() {
		localVarFormParams.Add("key_encipherment_usage", parameterToString(localVarOptionals.KeyEnciphermentUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DataEnciphermentUsage.IsSet() {
		localVarFormParams.Add("data_encipherment_usage", parameterToString(localVarOptionals.DataEnciphermentUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.KeyAgreementUsage.IsSet() {
		localVarFormParams.Add("key_agreement_usage", parameterToString(localVarOptionals.KeyAgreementUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.KeyCertSignUsage.IsSet() {
		localVarFormParams.Add("key_cert_sign_usage", parameterToString(localVarOptionals.KeyCertSignUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CrlSignUsage.IsSet() {
		localVarFormParams.Add("crl_sign_usage", parameterToString(localVarOptionals.CrlSignUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EncipherOnlyUsage.IsSet() {
		localVarFormParams.Add("encipher_only_usage", parameterToString(localVarOptionals.EncipherOnlyUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DecipherOnlyUsage.IsSet() {
		localVarFormParams.Add("decipher_only_usage", parameterToString(localVarOptionals.DecipherOnlyUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExtendedKeyUsage.IsSet() {
		localVarFormParams.Add("extended_key_usage", parameterToString(localVarOptionals.ExtendedKeyUsage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Vendor.IsSet() {
		localVarFormParams.Add("vendor", parameterToString(localVarOptionals.Vendor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndPointType.IsSet() {
		localVarFormParams.Add("end_point_type", parameterToString(localVarOptionals.EndPointType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndPointId.IsSet() {
		localVarFormParams.Add("end_point_id", parameterToString(localVarOptionals.EndPointId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Groups.IsSet() {
		localVarFormParams.Add("groups", parameterToString(localVarOptionals.Groups.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}