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

type Pdus struct {
	Building *interface{} `json:"building,omitempty"`
	BuildingId *interface{} `json:"building_id,omitempty"`
	Category *interface{} `json:"category,omitempty"`
	DeviceMapping *interface{} `json:"device mapping,omitempty"`
	Name *interface{} `json:"name,omitempty"`
	Notes *interface{} `json:"notes,omitempty"`
	PduId *interface{} `json:"pdu_id,omitempty"`
	PduModel *PdusPduModel `json:"pdu_model,omitempty"`
	Rack *interface{} `json:"rack,omitempty"`
	RackId *interface{} `json:"rack_id,omitempty"`
	Room *interface{} `json:"room,omitempty"`
	RoomId *interface{} `json:"room_id,omitempty"`
	Type_ *interface{} `json:"type,omitempty"`
}
