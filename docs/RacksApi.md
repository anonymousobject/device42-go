# \RacksApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRacksID**](RacksApi.md#DeleteRacksID) | **Delete** /api/1.0/racks/{ID}/ | This API is used to delete the rack with the rack id supplied as the required argument.
[**GetRacks**](RacksApi.md#GetRacks) | **Get** /api/1.0/racks/ | Get All Racks
[**GetRacksID**](RacksApi.md#GetRacksID) | **Get** /api/1.0/racks/{ID}/ | Retrieve detailed information about a specific rack including all racked devices, assets and PDUs.
[**PostRacks**](RacksApi.md#PostRacks) | **Post** /api/1.0/racks/ | Create a rack.
[**PutCustomFieldsRack**](RacksApi.md#PutCustomFieldsRack) | **Put** /api/1.0/custom_fields/rack/ | Create/update custom fields for racks.


# **DeleteRacksID**
> interface{} DeleteRacksID(ctx, iD)
This API is used to delete the rack with the rack id supplied as the required argument.

Delete Rack

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| rack id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRacks**
> interface{} GetRacks(ctx, optional)
Get All Racks

This API will retrieve basic information about all racks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RacksApiGetRacksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RacksApiGetRacksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **buildingId** | **optional.String**| filter by building ID (Added in v5.9.0) | 
 **building** | **optional.String**| filter by building name | 
 **roomId** | **optional.String**| filter by room ID (Added in v5.9.0) | 
 **room** | **optional.String**| filter by room name. Only works if room ID is not present (Added in v5.9.0) | 
 **size** | **optional.Int32**| filter by rack size in U | 
 **row** | **optional.String**| filter by row name | 
 **assetNo** | **optional.String**| filter by asset number | 
 **manufacturer** | **optional.String**| filter by manufacturer | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRacksID**
> interface{} GetRacksID(ctx, iD)
Retrieve detailed information about a specific rack including all racked devices, assets and PDUs.

Get a Specific Rack

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| The ID of the rack to retrieve | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRacks**
> interface{} PostRacks(ctx, name, size, optional)
Create a rack.

Create / Update Racks. Creating a new rack requires both <ul><li>name</li><li>size</li></ul><p> However, if updating a rack, use </p> <ul><li>rack_id</li></ul> <p>or all of</p> <ul><li>name</li><li>room</li> <li>building</li></ul><p> If using room/building name, first combination of room name or room and building name will be used.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Rack name - must be unique within a room | 
  **size** | **int32**| In UI | 
 **optional** | ***RacksApiPostRacksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RacksApiPostRacksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rackId** | **optional.Int32**| Required to update a rack using ID. This has highest priority to update a rack. | 
 **room** | **optional.String**| Name of room - Required if changing a rack without rack_id. | 
 **building** | **optional.String**| Name of building - Used when there are non-unique room names. | 
 **newName** | **optional.String**| Use to change name of object. | 
 **roomId** | **optional.String**| Room ID if Room name is not unique | 
 **numberingStartFromBottom** | **optional.String**| default is yes, no to change, otherwise ignored. | 
 **firstNumber** | **optional.String**| default 0, add to change. | 
 **row** | **optional.String**| this row field is for the name of the rows, and not related to the grid positioning of the rack | 
 **manufacturer** | **optional.String**| name of the hardware/software manufacturer. | 
 **notes** | **optional.String**| Any additional notes | 
 **startRow** | **optional.String**| Starting row for rack, for grid positioning | 
 **startCol** | **optional.String**| Starting column for the rack, for grid positioning | 
 **rowSize** | **optional.String**| How many rows long the rack is | 
 **colSize** | **optional.String**| how many racks wide the rack is | 
 **orientation** | **optional.String**| orientation of the rack in room layout view. Possible values: right, left, up or down. | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomFieldsRack**
> interface{} PutCustomFieldsRack(ctx, key, optional)
Create/update custom fields for racks.

Create or update custom fields for racks. \"ID\" or \"name\" of rack is needed even when value is not being changed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***RacksApiPutCustomFieldsRackOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RacksApiPutCustomFieldsRackOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of room | 
 **id** | **optional.String**| Rack ID or UI &gt; Tools &gt; Export &gt; Rack | 
 **type_** | **optional.String**| this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD | 
 **relatedFieldName** | **optional.String**| Required if type &#x3D; related field. | 
 **addToPicklist** | **optional.String**| Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored. | 
 **value** | **optional.String**| This will set the value of the custom field for the specific object. | 
 **clearValue** | **optional.String**| yes to clear existing value for that field | 
 **notes** | **optional.String**| Any additional notes | 
 **clearNotes** | **optional.String**| Yes to clear any existing notes. | 
 **bulkFields** | **optional.String**| comma separated key value pairs, with key and value separated by colon. e.g.key1:value1, key2:value2 | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

