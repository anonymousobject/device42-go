
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
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type PasswordsApiService service

/* 
PasswordsApiService Delete Password
This API is used to delete the password with the password id supplied as the required argument. Note: You will only be able to delete the password if the supplied username has the correct permissions.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id password id

@return interface{}
*/
func (a *PasswordsApiService) DeletePassword(ctx context.Context, id int32) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/passwords/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
PasswordsApiService Retrieve accounts (usernames) and passwords.
Get Usernames and Passwords
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PasswordsApiGetPasswordOpts - Optional Parameters:
     * @param "Category" (optional.String) -  name of the category
     * @param "Label" (optional.String) - 
     * @param "Username" (optional.String) -  Retrieve all the passwords for a specified username. ?username&#x3D;
     * @param "Device" (optional.String) -  Device name
     * @param "Appcomp" (optional.String) -  The application component that depends on this service
     * @param "Id" (optional.String) -  The ID of the software, required if not using NAME
     * @param "PlainText" (optional.String) -  Decrypt the password and return the plain text version. ?plain_text&#x3D;yes will decrypt and display the password.

@return interface{}
*/

type PasswordsApiGetPasswordOpts struct { 
	Category optional.String
	Label optional.String
	Username optional.String
	Device optional.String
	Appcomp optional.String
	Id optional.String
	PlainText optional.String
}

func (a *PasswordsApiService) GetPassword(ctx context.Context, localVarOptionals *PasswordsApiGetPasswordOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/passwords/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Category.IsSet() {
		localVarQueryParams.Add("category", parameterToString(localVarOptionals.Category.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Label.IsSet() {
		localVarQueryParams.Add("label", parameterToString(localVarOptionals.Label.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Username.IsSet() {
		localVarQueryParams.Add("username", parameterToString(localVarOptionals.Username.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Device.IsSet() {
		localVarQueryParams.Add("device", parameterToString(localVarOptionals.Device.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Appcomp.IsSet() {
		localVarQueryParams.Add("appcomp", parameterToString(localVarOptionals.Appcomp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlainText.IsSet() {
		localVarQueryParams.Add("plain_text", parameterToString(localVarOptionals.PlainText.Value(), ""))
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
PasswordsApiService
Custom Fields
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param username Name of the password
 * @param key Can be new or existing. This is the custom field name.
 * @param optional nil or *PasswordsApiPostCustomFieldsOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD
     * @param "RelatedFieldName" (optional.String) -  Required if type &#x3D; related field.
     * @param "AddToPicklist" (optional.String) -  Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored.
     * @param "Value" (optional.String) -  This will set the value of the custom field for the specific object.
     * @param "ClearValue" (optional.String) -  yes to clear existing value for that field
     * @param "Notes" (optional.String) -  Any additional notes
     * @param "ClearNotes" (optional.String) -  Yes to clear any existing notes.
     * @param "BulkFields" (optional.String) -  comma separated key value pairs, with key and value separated by colon. e.g.key1:value1, key2:value2

@return interface{}
*/

type PasswordsApiPostCustomFieldsOpts struct { 
	Type_ optional.String
	RelatedFieldName optional.String
	AddToPicklist optional.String
	Value optional.String
	ClearValue optional.String
	Notes optional.String
	ClearNotes optional.String
	BulkFields optional.String
}

func (a *PasswordsApiService) PostCustomFields(ctx context.Context, username string, key string, localVarOptionals *PasswordsApiPostCustomFieldsOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/custom_fields/password/"

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
	localVarFormParams.Add("username", parameterToString(username, ""))
	localVarFormParams.Add("key", parameterToString(key, ""))
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarFormParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RelatedFieldName.IsSet() {
		localVarFormParams.Add("related_field_name", parameterToString(localVarOptionals.RelatedFieldName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AddToPicklist.IsSet() {
		localVarFormParams.Add("add_to_picklist", parameterToString(localVarOptionals.AddToPicklist.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Value.IsSet() {
		localVarFormParams.Add("value", parameterToString(localVarOptionals.Value.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClearValue.IsSet() {
		localVarFormParams.Add("clear_value", parameterToString(localVarOptionals.ClearValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Notes.IsSet() {
		localVarFormParams.Add("notes", parameterToString(localVarOptionals.Notes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClearNotes.IsSet() {
		localVarFormParams.Add("clear_notes", parameterToString(localVarOptionals.ClearNotes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BulkFields.IsSet() {
		localVarFormParams.Add("bulk_fields", parameterToString(localVarOptionals.BulkFields.Value(), ""))
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
PasswordsApiService
Create / Update Passwords. Use id if updating existing password - else, username and password are required.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PasswordsApiPostUpdatePasswordsOpts - Optional Parameters:
     * @param "Id" (optional.String) -  Required to modify existing password
     * @param "Username" (optional.String) - 
     * @param "Password" (optional.String) -  The password.
     * @param "Label" (optional.String) - 
     * @param "Category" (optional.String) -  A password category. If it doesn&#39;t exist a new one will be created.
     * @param "Devices" (optional.String) -  A comma separated list of device names.
     * @param "Appcomps" (optional.String) -  A comma separated list of application component names.
     * @param "DaysBeforeExpiry" (optional.String) -  number of days before password is set as expired
     * @param "Notes" (optional.String) -  Any additional notes
     * @param "ViewUsers" (optional.String) -  A comma separated list of users that have permission to view this password.
     * @param "ViewGroups" (optional.String) -  A comma separated list of user groups that have permission to view this password.
     * @param "ViewEditUsers" (optional.String) -  A comma separated list of users that have permission to view and edit this password.
     * @param "ViewUsersRemove" (optional.String) -  A comma separated list of users to remove view permissions.
     * @param "UseOnlyUsersRemove" (optional.String) -  A comma separated list of users to remove use only permissions.
     * @param "ViewEditUsersRemove" (optional.String) -  A comma separated list of users to remove view and edit permissions.
     * @param "ViewGroupsRemove" (optional.String) -  A comma separated list of user groups to remove use permissions.
     * @param "UseOnlyGroupsRemove" (optional.String) -  A comma separated list of user groups to remove use permissions.
     * @param "ViewEditGroupsRemove" (optional.String) -  A comma separated list of user groups to remove view and edit permissions.

@return interface{}
*/

type PasswordsApiPostUpdatePasswordsOpts struct { 
	Id optional.String
	Username optional.String
	Password optional.String
	Label optional.String
	Category optional.String
	Devices optional.String
	Appcomps optional.String
	DaysBeforeExpiry optional.String
	Notes optional.String
	ViewUsers optional.String
	ViewGroups optional.String
	ViewEditUsers optional.String
	ViewUsersRemove optional.String
	UseOnlyUsersRemove optional.String
	ViewEditUsersRemove optional.String
	ViewGroupsRemove optional.String
	UseOnlyGroupsRemove optional.String
	ViewEditGroupsRemove optional.String
}

func (a *PasswordsApiService) PostUpdatePasswords(ctx context.Context, localVarOptionals *PasswordsApiPostUpdatePasswordsOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/passwords/"

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
	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarFormParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Username.IsSet() {
		localVarFormParams.Add("username", parameterToString(localVarOptionals.Username.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Password.IsSet() {
		localVarFormParams.Add("password", parameterToString(localVarOptionals.Password.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Label.IsSet() {
		localVarFormParams.Add("label", parameterToString(localVarOptionals.Label.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Category.IsSet() {
		localVarFormParams.Add("category", parameterToString(localVarOptionals.Category.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Devices.IsSet() {
		localVarFormParams.Add("devices", parameterToString(localVarOptionals.Devices.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Appcomps.IsSet() {
		localVarFormParams.Add("appcomps", parameterToString(localVarOptionals.Appcomps.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DaysBeforeExpiry.IsSet() {
		localVarFormParams.Add("days_before_expiry", parameterToString(localVarOptionals.DaysBeforeExpiry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Notes.IsSet() {
		localVarFormParams.Add("notes", parameterToString(localVarOptionals.Notes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ViewUsers.IsSet() {
		localVarFormParams.Add("view_users", parameterToString(localVarOptionals.ViewUsers.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ViewGroups.IsSet() {
		localVarFormParams.Add("view_groups", parameterToString(localVarOptionals.ViewGroups.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ViewEditUsers.IsSet() {
		localVarFormParams.Add("view_edit_users", parameterToString(localVarOptionals.ViewEditUsers.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ViewUsersRemove.IsSet() {
		localVarFormParams.Add("view_users_remove", parameterToString(localVarOptionals.ViewUsersRemove.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UseOnlyUsersRemove.IsSet() {
		localVarFormParams.Add("use_only_users_remove", parameterToString(localVarOptionals.UseOnlyUsersRemove.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ViewEditUsersRemove.IsSet() {
		localVarFormParams.Add("view_edit_users_remove", parameterToString(localVarOptionals.ViewEditUsersRemove.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ViewGroupsRemove.IsSet() {
		localVarFormParams.Add("view_groups_remove", parameterToString(localVarOptionals.ViewGroupsRemove.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UseOnlyGroupsRemove.IsSet() {
		localVarFormParams.Add("use_only_groups_remove", parameterToString(localVarOptionals.UseOnlyGroupsRemove.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ViewEditGroupsRemove.IsSet() {
		localVarFormParams.Add("view_edit_groups_remove", parameterToString(localVarOptionals.ViewEditGroupsRemove.Value(), ""))
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
