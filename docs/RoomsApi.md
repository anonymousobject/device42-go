# \RoomsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAssetsRoom**](RoomsApi.md#DeleteAssetsRoom) | **Delete** /api/1.0/assets/room/{id}/ | Delete an Asset from its Room
[**DeleteDeviceRoom**](RoomsApi.md#DeleteDeviceRoom) | **Delete** /api/1.0/device/room/{ID}/ | Delete Device from its Room
[**DeleteRoomsID**](RoomsApi.md#DeleteRoomsID) | **Delete** /api/1.0/rooms/{ID}/ | Delete Room
[**GetRooms**](RoomsApi.md#GetRooms) | **Get** /api/1.0/rooms/ | Retrieve information about all rooms.
[**GetRoomsID**](RoomsApi.md#GetRoomsID) | **Get** /api/1.0/rooms/{ID}/ | Get a Specific Room
[**PostAssetsRoom**](RoomsApi.md#PostAssetsRoom) | **Post** /api/1.0/assets/room/ | Add an asset to a room.
[**PostDeviceRoom**](RoomsApi.md#PostDeviceRoom) | **Post** /api/1.0/device/room/ | Add a device to a room.
[**PostRooms**](RoomsApi.md#PostRooms) | **Post** /api/1.0/rooms/ | Create or update a room.
[**PutCustomFieldsRoom**](RoomsApi.md#PutCustomFieldsRoom) | **Put** /api/1.0/custom_fields/room/ | Create/update custom fields for rooms.
[**PutRoomsID**](RoomsApi.md#PutRoomsID) | **Put** /api/1.0/rooms/{ID}/ | Update information about an existing room.


# **DeleteAssetsRoom**
> interface{} DeleteAssetsRoom(ctx, id)
Delete an Asset from its Room

This API is used to delete the asset with the asset id supplied as the required argument from its room.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| device id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeviceRoom**
> interface{} DeleteDeviceRoom(ctx, iD)
Delete Device from its Room

This API is used to delete the device from its room with the device id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| device id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoomsID**
> interface{} DeleteRoomsID(ctx, iD)
Delete Room

This API is used to delete the room with the room id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| room id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRooms**
> interface{} GetRooms(ctx, optional)
Retrieve information about all rooms.

Get All Rooms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoomsApiGetRoomsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoomsApiGetRoomsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **buildingId** | **optional.String**| filter by building ID (Added in v5.9.0) | 
 **building** | **optional.String**| filter by building name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoomsID**
> interface{} GetRoomsID(ctx, iD)
Get a Specific Room

Retrieve detailed information about a specific rooms includes racks, devices and objects directly related to that room.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| The ID of the room to retrieve | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAssetsRoom**
> interface{} PostAssetsRoom(ctx, roomId, asset, type_, optional)
Add an asset to a room.

Add asset to room

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **string**| Room ID or UI &gt; Tools &gt; Export &gt; Room | 
  **asset** | **string**| The name of the asset | 
  **type_** | **string**|  | 
 **optional** | ***RoomsApiPostAssetsRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoomsApiPostAssetsRoomOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **wall** | **optional.String**| Choose &#39;middle&#39; if you do not want the object to be placed along one of the 4 walls. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeviceRoom**
> interface{} PostDeviceRoom(ctx, roomId, name, type_, optional)
Add a device to a room.

Add device to room

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **string**| Room ID or UI &gt; Tools &gt; Export &gt; Room | 
  **name** | **string**| The name of the device to add | 
  **type_** | **string**|  | 
 **optional** | ***RoomsApiPostDeviceRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoomsApiPostDeviceRoomOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **wall** | **optional.String**| Choose &#39;middle&#39; if you do not want the object to be placed along one of the 4 walls. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRooms**
> interface{} PostRooms(ctx, name, optional)
Create or update a room.

Create or update a Room. Required parameters: <ul><li>name</li> <li>building_id <b>OR</b> building</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of room | 
 **optional** | ***RoomsApiPostRoomsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoomsApiPostRoomsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildingId** | **optional.String**| Existing building ID | 
 **building** | **optional.String**| Existing building name | 
 **notes** | **optional.String**| Any additional notes | 
 **horizontalGridNumbering** | **optional.String**| numeric by default | 
 **horizontalGridStart** | **optional.String**|  | 
 **verticalGridNumbering** | **optional.String**| numeric by default | 
 **verticalGridStart** | **optional.String**|  | 
 **uom** | **optional.String**| unit of measurement (meters or inches) | 
 **height** | **optional.String**| room height | 
 **gridRows** | **optional.String**| number of rows in the room grid | 
 **gridCols** | **optional.String**| number of columns in the room grid | 
 **raisedFloor** | **optional.String**|  | 
 **raisedFloorHeight** | **optional.String**| height of raised floor | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 
 **reverseXaxis** | **optional.String**| &#39;yes&#39; reverses the numbering order on the x-axis | 
 **reverseYaxis** | **optional.String**| &#39;yes&#39; reverses the numbering order on the y-axis | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomFieldsRoom**
> interface{} PutCustomFieldsRoom(ctx, key, optional)
Create/update custom fields for rooms.

Create or update custom fields for rooms. \"ID\" or \"name\" of room is needed even when value is not being changed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***RoomsApiPutCustomFieldsRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoomsApiPutCustomFieldsRoomOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Name of room | 
 **id** | **optional.String**| Room ID or UI &gt; Tools &gt; Export &gt; Room | 
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

# **PutRoomsID**
> interface{} PutRoomsID(ctx, iD, optional)
Update information about an existing room.

Update a Room

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| id of the room | 
 **optional** | ***RoomsApiPutRoomsIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoomsApiPutRoomsIDOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Name of room | 
 **notes** | **optional.String**| Any additional notes | 
 **buildingId** | **optional.String**| Existing building ID | 
 **horizontalGridNumbering** | **optional.String**| numeric by default | 
 **horizontalGridStart** | **optional.String**|  | 
 **verticalGridNumbering** | **optional.String**| numeric by default | 
 **verticalGridStart** | **optional.String**|  | 
 **uom** | **optional.String**| unit of measurement (meters or inches) | 
 **height** | **optional.String**| room height | 
 **gridRows** | **optional.String**| number of rows in the room grid | 
 **gridCols** | **optional.String**| number of columns in the room grid | 
 **raisedFloor** | **optional.String**|  | 
 **raisedFloorHeight** | **optional.String**| height of raised floor | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 
 **reverseXaxis** | **optional.String**| &#39;yes&#39; reverses the numbering order on the x-axis | 
 **reverseYaxis** | **optional.String**| &#39;yes&#39; reverses the numbering order on the y-axis | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

