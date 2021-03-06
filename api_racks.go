
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

type RacksApiService service

/* 
RacksApiService This API is used to delete the rack with the rack id supplied as the required argument.
Delete Rack
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param iD rack id

@return interface{}
*/
func (a *RacksApiService) DeleteRacksID(ctx context.Context, iD int32) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/racks/{ID}/"
	localVarPath = strings.Replace(localVarPath, "{"+"ID"+"}", fmt.Sprintf("%v", iD), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
RacksApiService Get All Racks
This API will retrieve basic information about all racks.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RacksApiGetRacksOpts - Optional Parameters:
     * @param "Name" (optional.String) -  filter by name (Added in v6.0.0)
     * @param "BuildingId" (optional.String) -  filter by building ID (Added in v5.9.0)
     * @param "Building" (optional.String) -  filter by building name
     * @param "RoomId" (optional.String) -  filter by room ID (Added in v5.9.0)
     * @param "Room" (optional.String) -  filter by room name. Only works if room ID is not present (Added in v5.9.0)
     * @param "Size" (optional.Int32) -  filter by rack size in U
     * @param "Row" (optional.String) -  filter by row name
     * @param "AssetNo" (optional.String) -  filter by asset number
     * @param "Manufacturer" (optional.String) -  filter by manufacturer

@return interface{}
*/

type RacksApiGetRacksOpts struct { 
	Name optional.String
	BuildingId optional.String
	Building optional.String
	RoomId optional.String
	Room optional.String
	Size optional.Int32
	Row optional.String
	AssetNo optional.String
	Manufacturer optional.String
}

func (a *RacksApiService) GetRacks(ctx context.Context, localVarOptionals *RacksApiGetRacksOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/racks/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BuildingId.IsSet() {
		localVarQueryParams.Add("building_id", parameterToString(localVarOptionals.BuildingId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Building.IsSet() {
		localVarQueryParams.Add("building", parameterToString(localVarOptionals.Building.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoomId.IsSet() {
		localVarQueryParams.Add("room_id", parameterToString(localVarOptionals.RoomId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Room.IsSet() {
		localVarQueryParams.Add("room", parameterToString(localVarOptionals.Room.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Row.IsSet() {
		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AssetNo.IsSet() {
		localVarQueryParams.Add("asset_no", parameterToString(localVarOptionals.AssetNo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Manufacturer.IsSet() {
		localVarQueryParams.Add("manufacturer", parameterToString(localVarOptionals.Manufacturer.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
RacksApiService Retrieve detailed information about a specific rack including all racked devices, assets and PDUs.
Get a Specific Rack
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param iD The ID of the rack to retrieve

@return interface{}
*/
func (a *RacksApiService) GetRacksID(ctx context.Context, iD int32) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/racks/{ID}/"
	localVarPath = strings.Replace(localVarPath, "{"+"ID"+"}", fmt.Sprintf("%v", iD), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
RacksApiService Create a rack.
Create / Update Racks. Creating a new rack requires both &lt;ul&gt;&lt;li&gt;name&lt;/li&gt;&lt;li&gt;size&lt;/li&gt;&lt;/ul&gt;&lt;p&gt; However, if updating a rack, use &lt;/p&gt; &lt;ul&gt;&lt;li&gt;rack_id&lt;/li&gt;&lt;/ul&gt; &lt;p&gt;or all of&lt;/p&gt; &lt;ul&gt;&lt;li&gt;name&lt;/li&gt;&lt;li&gt;room&lt;/li&gt; &lt;li&gt;building&lt;/li&gt;&lt;/ul&gt;&lt;p&gt; If using room/building name, first combination of room name or room and building name will be used.&lt;/p&gt;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name Rack name - must be unique within a room
 * @param size In UI
 * @param optional nil or *RacksApiPostRacksOpts - Optional Parameters:
     * @param "RackId" (optional.Int32) -  Required to update a rack using ID. This has highest priority to update a rack.
     * @param "Room" (optional.String) -  Name of room - Required if changing a rack without rack_id.
     * @param "Building" (optional.String) -  Name of building - Used when there are non-unique room names.
     * @param "NewName" (optional.String) -  Use to change name of object.
     * @param "RoomId" (optional.String) -  Room ID if Room name is not unique
     * @param "NumberingStartFromBottom" (optional.String) -  default is yes, no to change, otherwise ignored.
     * @param "FirstNumber" (optional.String) -  default 0, add to change.
     * @param "Row" (optional.String) -  this row field is for the name of the rows, and not related to the grid positioning of the rack
     * @param "Manufacturer" (optional.String) -  name of the hardware/software manufacturer.
     * @param "Notes" (optional.String) -  Any additional notes
     * @param "StartRow" (optional.String) -  Starting row for rack, for grid positioning
     * @param "StartCol" (optional.String) -  Starting column for the rack, for grid positioning
     * @param "RowSize" (optional.String) -  How many rows long the rack is
     * @param "ColSize" (optional.String) -  how many racks wide the rack is
     * @param "Orientation" (optional.String) -  orientation of the rack in room layout view. Possible values: right, left, up or down.
     * @param "Groups" (optional.String) -  If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted.

@return interface{}
*/

type RacksApiPostRacksOpts struct { 
	RackId optional.Int32
	Room optional.String
	Building optional.String
	NewName optional.String
	RoomId optional.String
	NumberingStartFromBottom optional.String
	FirstNumber optional.String
	Row optional.String
	Manufacturer optional.String
	Notes optional.String
	StartRow optional.String
	StartCol optional.String
	RowSize optional.String
	ColSize optional.String
	Orientation optional.String
	Groups optional.String
}

func (a *RacksApiService) PostRacks(ctx context.Context, name string, size int32, localVarOptionals *RacksApiPostRacksOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/racks/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("size", parameterToString(size, ""))
	if localVarOptionals != nil && localVarOptionals.RackId.IsSet() {
		localVarFormParams.Add("rack_id", parameterToString(localVarOptionals.RackId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Room.IsSet() {
		localVarFormParams.Add("room", parameterToString(localVarOptionals.Room.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Building.IsSet() {
		localVarFormParams.Add("building", parameterToString(localVarOptionals.Building.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NewName.IsSet() {
		localVarFormParams.Add("new_name", parameterToString(localVarOptionals.NewName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoomId.IsSet() {
		localVarFormParams.Add("room_id", parameterToString(localVarOptionals.RoomId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NumberingStartFromBottom.IsSet() {
		localVarFormParams.Add("numbering_start_from_bottom", parameterToString(localVarOptionals.NumberingStartFromBottom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FirstNumber.IsSet() {
		localVarFormParams.Add("first_number", parameterToString(localVarOptionals.FirstNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Row.IsSet() {
		localVarFormParams.Add("row", parameterToString(localVarOptionals.Row.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Manufacturer.IsSet() {
		localVarFormParams.Add("manufacturer", parameterToString(localVarOptionals.Manufacturer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Notes.IsSet() {
		localVarFormParams.Add("notes", parameterToString(localVarOptionals.Notes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartRow.IsSet() {
		localVarFormParams.Add("start_row", parameterToString(localVarOptionals.StartRow.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartCol.IsSet() {
		localVarFormParams.Add("start_col", parameterToString(localVarOptionals.StartCol.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RowSize.IsSet() {
		localVarFormParams.Add("row_size", parameterToString(localVarOptionals.RowSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ColSize.IsSet() {
		localVarFormParams.Add("col_size", parameterToString(localVarOptionals.ColSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Orientation.IsSet() {
		localVarFormParams.Add("orientation", parameterToString(localVarOptionals.Orientation.Value(), ""))
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

/* 
RacksApiService Create/update custom fields for racks.
Create or update custom fields for racks. \&quot;ID\&quot; or \&quot;name\&quot; of rack is needed even when value is not being changed.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param key Can be new or existing. This is the custom field name.
 * @param optional nil or *RacksApiPutCustomFieldsRackOpts - Optional Parameters:
     * @param "Name" (optional.String) -  name of room
     * @param "Id" (optional.String) -  Rack ID or UI &gt; Tools &gt; Export &gt; Rack
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

type RacksApiPutCustomFieldsRackOpts struct { 
	Name optional.String
	Id optional.String
	Type_ optional.String
	RelatedFieldName optional.String
	AddToPicklist optional.String
	Value optional.String
	ClearValue optional.String
	Notes optional.String
	ClearNotes optional.String
	BulkFields optional.String
}

func (a *RacksApiService) PutCustomFieldsRack(ctx context.Context, key string, localVarOptionals *RacksApiPutCustomFieldsRackOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1.0/custom_fields/rack/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarFormParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarFormParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
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
