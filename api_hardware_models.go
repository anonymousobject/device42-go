
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

type HardwareModelsApiService service

/* 
HardwareModelsApiService Delete Hardware Model
This API is used to delete the hardware model with the hardware model id supplied as the required argument.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id hardware model id

@return interface{}
*/
func (a *HardwareModelsApiService) DeleteHardwares(ctx context.Context, id int32) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/hardwares/{id}/"
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
HardwareModelsApiService This call will get information about hardware models.
Get all hardware models
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *HardwareModelsApiGetHardwaresOpts - Optional Parameters:
     * @param "Name" (optional.String) -  filter by name (Added in v6.0.0)
     * @param "Type_" (optional.String) -  could be physical, blade, or other
     * @param "DeviceSubType" (optional.String) -  filter by device sub type (Added in v14.7.2)
     * @param "Size" (optional.String) -  filter by exact size
     * @param "Depth" (optional.String) -  could be half or full
     * @param "PartNo" (optional.String) -  filter by part #
     * @param "Watts" (optional.String) -  filter by exact watts
     * @param "Manufacturer" (optional.String) -  name of the hardware manufacturer.

@return interface{}
*/

type HardwareModelsApiGetHardwaresOpts struct { 
	Name optional.String
	Type_ optional.String
	DeviceSubType optional.String
	Size optional.String
	Depth optional.String
	PartNo optional.String
	Watts optional.String
	Manufacturer optional.String
}

func (a *HardwareModelsApiService) GetHardwares(ctx context.Context, localVarOptionals *HardwareModelsApiGetHardwaresOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/hardwares/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DeviceSubType.IsSet() {
		localVarQueryParams.Add("device_sub_type", parameterToString(localVarOptionals.DeviceSubType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Depth.IsSet() {
		localVarQueryParams.Add("depth", parameterToString(localVarOptionals.Depth.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PartNo.IsSet() {
		localVarQueryParams.Add("part_no", parameterToString(localVarOptionals.PartNo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Watts.IsSet() {
		localVarQueryParams.Add("watts", parameterToString(localVarOptionals.Watts.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Manufacturer.IsSet() {
		localVarQueryParams.Add("manufacturer", parameterToString(localVarOptionals.Manufacturer.Value(), ""))
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
HardwareModelsApiService This call will create/update information about hardware models.
Create/update
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name if similar hardware name already exists, first matching entry is updated
 * @param optional nil or *HardwareModelsApiPostHardwaresOpts - Optional Parameters:
     * @param "NewName" (optional.String) -  Use to change name of object.
     * @param "Type_" (optional.Int32) -  1 &#x3D; Regular, 2 &#x3D; Blade, 3 &#x3D; Other
     * @param "DeviceSubtype" (optional.String) -  Subtype of \&quot;other\&quot; type devices
     * @param "Size" (optional.String) -  Size in U for hardware type regular
     * @param "WidthRatio" (optional.String) -  Default&#x3D;1. Can be ½, 1/3,… 1/10, etc.
     * @param "Depth" (optional.String) -  half by default, full to override
     * @param "BladeSize" (optional.Int32) -  1&#x3D;Full Height 2&#x3D;Half Height 3&#x3D;Double Half Height 4&#x3D;Double Full Height 5&#x3D;Quarter Height
     * @param "PartNo" (optional.String) - 
     * @param "Watts" (optional.String) -  per power supply
     * @param "SpecUrl" (optional.String) -  Specification url for the hardware model.
     * @param "Manufacturer" (optional.String) -  name of the hardware/software manufacturer.
     * @param "FrontImageId" (optional.String) - 
     * @param "FrontImage" (optional.String) -  name of the image file (Added in v5.8.2)
     * @param "BackImageId" (optional.String) -  back image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now.
     * @param "BackImage" (optional.String) -  name of the back image file. Use instead of back_image_id.
     * @param "Notes" (optional.String) -  Any additional notes
     * @param "MaxBladesPerRow" (optional.String) - 
     * @param "SlotNumbering" (optional.String) - 
     * @param "ModulePos" (optional.String) - 

@return interface{}
*/

type HardwareModelsApiPostHardwaresOpts struct { 
	NewName optional.String
	Type_ optional.Int32
	DeviceSubtype optional.String
	Size optional.String
	WidthRatio optional.String
	Depth optional.String
	BladeSize optional.Int32
	PartNo optional.String
	Watts optional.String
	SpecUrl optional.String
	Manufacturer optional.String
	FrontImageId optional.String
	FrontImage optional.String
	BackImageId optional.String
	BackImage optional.String
	Notes optional.String
	MaxBladesPerRow optional.String
	SlotNumbering optional.String
	ModulePos optional.String
}

func (a *HardwareModelsApiService) PostHardwares(ctx context.Context, name string, localVarOptionals *HardwareModelsApiPostHardwaresOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/hardwares/"

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
	localVarFormParams.Add("name", parameterToString(name, ""))
	if localVarOptionals != nil && localVarOptionals.NewName.IsSet() {
		localVarFormParams.Add("new_name", parameterToString(localVarOptionals.NewName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarFormParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DeviceSubtype.IsSet() {
		localVarFormParams.Add("device_subtype", parameterToString(localVarOptionals.DeviceSubtype.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarFormParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WidthRatio.IsSet() {
		localVarFormParams.Add("width_ratio", parameterToString(localVarOptionals.WidthRatio.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Depth.IsSet() {
		localVarFormParams.Add("depth", parameterToString(localVarOptionals.Depth.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BladeSize.IsSet() {
		localVarFormParams.Add("blade_size", parameterToString(localVarOptionals.BladeSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PartNo.IsSet() {
		localVarFormParams.Add("part_no", parameterToString(localVarOptionals.PartNo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Watts.IsSet() {
		localVarFormParams.Add("watts", parameterToString(localVarOptionals.Watts.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SpecUrl.IsSet() {
		localVarFormParams.Add("spec_url", parameterToString(localVarOptionals.SpecUrl.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Manufacturer.IsSet() {
		localVarFormParams.Add("manufacturer", parameterToString(localVarOptionals.Manufacturer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FrontImageId.IsSet() {
		localVarFormParams.Add("front_image_id", parameterToString(localVarOptionals.FrontImageId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FrontImage.IsSet() {
		localVarFormParams.Add("front_image", parameterToString(localVarOptionals.FrontImage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BackImageId.IsSet() {
		localVarFormParams.Add("back_image_id", parameterToString(localVarOptionals.BackImageId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BackImage.IsSet() {
		localVarFormParams.Add("back_image", parameterToString(localVarOptionals.BackImage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Notes.IsSet() {
		localVarFormParams.Add("notes", parameterToString(localVarOptionals.Notes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaxBladesPerRow.IsSet() {
		localVarFormParams.Add("max_blades_per_row", parameterToString(localVarOptionals.MaxBladesPerRow.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SlotNumbering.IsSet() {
		localVarFormParams.Add("slot_numbering", parameterToString(localVarOptionals.SlotNumbering.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ModulePos.IsSet() {
		localVarFormParams.Add("module_pos", parameterToString(localVarOptionals.ModulePos.Value(), ""))
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
