# \AssetsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAssets**](AssetsApi.md#DeleteAssets) | **Delete** /api/1.0/assets/{asset-id}/ | This API is used to delete the asset with the asset id supplied as the required argument.
[**GetAssets**](AssetsApi.md#GetAssets) | **Get** /api/1.0/assets/ | Retrieve basic information about all assets.
[**GetAssetsAssetId**](AssetsApi.md#GetAssetsAssetId) | **Get** /api/1.0/assets/{asset-id}/ | Retrieve detailed information about a specific asset.
[**PostAssets**](AssetsApi.md#PostAssets) | **Post** /api/1.0/assets/ | Create assets.
[**PutAssets**](AssetsApi.md#PutAssets) | **Put** /api/1.0/assets/ | Modify assets.
[**PutCustomFieldsAsset**](AssetsApi.md#PutCustomFieldsAsset) | **Put** /api/1.0/custom_fields/asset/ | Create/update custom fields for assets.


# **DeleteAssets**
> interface{} DeleteAssets(ctx, assetId)
This API is used to delete the asset with the asset id supplied as the required argument.

Delete Asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetId** | **int32**| asset id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssets**
> interface{} GetAssets(ctx, optional)
Retrieve basic information about all assets.

Get All Assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsApiGetAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiGetAssetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetNo** | **optional.String**| filter by asset # (Added in v6.0.0) | 
 **serialNo** | **optional.String**| filter by serial # (Added in v6.0.0) | 
 **lastUpdatedLt** | **optional.String**| last updated less than date YYYY-MM-DD format | 
 **lastUpdatedGt** | **optional.String**| last updated greater than date YYYY-MM-DD format | 
 **firstAddedLt** | **optional.String**| first added less than date YYYY-MM-DD format | 
 **firstAddedGt** | **optional.String**| first added greater date YYYY-MM-DD format | 
 **type_** | **optional.String**| filter by asset type | 
 **assetId** | **optional.String**| comma separated values of existing assets. | 
 **serviceLevel** | **optional.String**| filter by service level name | 
 **customer** | **optional.String**| filter by customer name | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 
 **tagsAnd** | **optional.String**| filter by all the tags, separated by comma. (This is an AND filter and all tags have to match for filter, this was added in v6.3.1) | 
 **assetNoContains** | **optional.String**| search for any asset that contains matching asset # (Added in v9.2.0) | 
 **customFieldsAnd** | **optional.String**| filter by custom fields, and filter, format of key1:value1,key2:value2 | 
 **customFieldsOr** | **optional.String**| filter by custom fields, or filter, format of key1:value1,key2:value2 | 
 **relatedDeviceId** | **optional.String**| ID of the related device (added in v9.3.0) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssetsAssetId**
> Assets GetAssetsAssetId(ctx, assetId, optional)
Retrieve detailed information about a specific asset.

Get a Specific Asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetId** | **int32**| The ID of the asset to retrieve | 
 **optional** | ***AssetsApiGetAssetsAssetIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiGetAssetsAssetIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCols** | **optional.String**| do not return all columns just the ones specified. For example, ?include_cols&#x3D;name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager | 

### Return type

[**Assets**](Assets.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAssets**
> interface{} PostAssets(ctx, type_, optional)
Create assets.

Create Assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Select an existing asset type - required.&lt;br&gt; You can add a new asset type in the UI or with PUT. | 
 **optional** | ***AssetsApiPostAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiPostAssetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of asset | 
 **serviceLevel** | **optional.String**| In Service, Spare, Not in Service are pre-defined - or choose your own. | 
 **serialNo** | **optional.String**|  | 
 **assetNo** | **optional.String**|  | 
 **customer** | **optional.String**| Name of existing customer. | 
 **location** | **optional.String**| Location/region of instance deployment | 
 **notes** | **optional.String**| Any additional notes | 
 **building** | **optional.String**|  | 
 **vendor** | **optional.String**| Name of existing vendor | 
 **imgfileId** | **optional.String**| image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now. | 
 **imgfile** | **optional.String**| name of the image file (Added in v5.8.2). Use instead of imgfile_id | 
 **backImageId** | **optional.String**| back image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now. | 
 **backImage** | **optional.String**| name of the back image file. Use instead of back_image_id. | 
 **rackId** | **optional.String**| ID of existing rack to add asset to. | 
 **room** | **optional.String**| Used to identify a unique rack or to place asset in existing room. | 
 **startAt** | **optional.String**| Required if adding to rack. U Start location. | 
 **size** | **optional.Int32**| Required if adding asset to rack. in U. | 
 **orientation** | **optional.String**| Back if back facing. Otherwise ignored | 
 **where** | **optional.String**| Location in a rack. Note: If mounted a size must be provided or available from the hardware model. | 
 **xPos** | **optional.Int32**| A number between 0 and 2520 representing the position within the u slot in increments of 252, which is equal to 1/10th of the width of the rack. 0 will place a device flush left, 1260 will place the left side of a device in center. | 
 **depth** | **optional.String**| Half depth by default. full to override. Optional. | 
 **deviceId** | **optional.Int32**| ID of the related device | 
 **tags** | **optional.String**| add tags (comma separated) | 
 **tagsRemove** | **optional.String**| remove tags (comma separated) | 
 **category** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no, Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. Use for initial insert. | 
 **patchPanelModelId** | **optional.String**| Patch Panel Model ID or UI Tools &gt; Export &gt; Patch Panel Model | 
 **patchPanelModel** | **optional.String**| Name of the patch panel model (use instead of ID, Added in v5.8.2) | 
 **numberingStartFrom** | **optional.String**| This is starting # for patch panel ports. Defaults to 1 if not entered. | 
 **patchPanelModuleModelId** | **optional.String**|  | 
 **patchPanelModuleModel** | **optional.String**| Name of the patch panel module model (use instead of ID, Added in v5.8.2) | 
 **moduleHostId** | **optional.String**|  | 
 **moduleHost** | **optional.String**| Name of the Module host. Must be unique asset name for this to work. (use instead of ID, Added in v5.8.2) | 
 **slotNo** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAssets**
> interface{} PutAssets(ctx, optional)
Modify assets.

Modify Assets. Need either <b>Asset ID</b> or <b>Asset</b>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsApiPutAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiPutAssetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetId** | **optional.Int32**| Asset ID | 
 **asset** | **optional.String**| Name of asset. Required if asset_id not provided. | 
 **type_** | **optional.String**| Used to change/add the type of asset | 
 **serviceLevel** | **optional.String**| In Service, Spare, Not in Service are pre-defined - or choose your own. | 
 **serialNo** | **optional.String**|  | 
 **assetNo** | **optional.String**|  | 
 **customer** | **optional.String**| Name of existing customer. | 
 **location** | **optional.String**| Location/region of instance deployment | 
 **notes** | **optional.String**| Any additional notes | 
 **building** | **optional.String**|  | 
 **vendor** | **optional.String**| Name of existing vendor | 
 **imgfileId** | **optional.String**| image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now. | 
 **imgfile** | **optional.String**| name of the image file (Added in v5.8.2). Use instead of imgfile_id | 
 **backImageId** | **optional.String**| back image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now. | 
 **backImage** | **optional.String**| name of the back image file. Use instead of back_image_id. | 
 **rackId** | **optional.String**| ID of existing rack to add asset to. | 
 **room** | **optional.String**| Used to identify a unique rack or to place asset in existing room. | 
 **startAt** | **optional.String**| Required if adding to rack. U Start location. | 
 **size** | **optional.Int32**| Required if adding asset to rack. in U. | 
 **orientation** | **optional.String**| Back if back facing. Otherwise ignored | 
 **where** | **optional.String**| Location in a rack. Note: If mounted a size must be provided or available from the hardware model. | 
 **xPos** | **optional.Int32**| A number between 0 and 2520 representing the position within the u slot in increments of 252, which is equal to 1/10th of the width of the rack. 0 will place a device flush left, 1260 will place the left side of a device in center. | 
 **depth** | **optional.String**| Half depth by default. full to override. Optional. | 
 **deviceId** | **optional.Int32**| ID of the related device | 
 **tags** | **optional.String**| add tags (comma separated) | 
 **tagsRemove** | **optional.String**| remove tags (comma separated) | 
 **category** | **optional.String**| name of the category | 
 **patchPanelModelId** | **optional.String**| Patch Panel Model ID or UI Tools &gt; Export &gt; Patch Panel Model | 
 **patchPanelModel** | **optional.String**| Name of the patch panel model (use instead of ID, Added in v5.8.2) | 
 **numberingStartFrom** | **optional.String**| This is starting # for patch panel ports. Defaults to 1 if not entered. | 
 **patchPanelModuleModelId** | **optional.String**|  | 
 **patchPanelModuleModel** | **optional.String**| Name of the patch panel module model (use instead of ID, Added in v5.8.2) | 
 **moduleHostId** | **optional.String**|  | 
 **moduleHost** | **optional.String**| Name of the Module host. Must be unique asset name for this to work. (use instead of ID, Added in v5.8.2) | 
 **slotNo** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomFieldsAsset**
> interface{} PutCustomFieldsAsset(ctx, key, optional)
Create/update custom fields for assets.

Create or update custom fields for assets. \"ID\" or \"name\" of asset is needed even when value is not being changed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***AssetsApiPutCustomFieldsAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiPutCustomFieldsAssetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Name of asset | 
 **id** | **optional.String**| ID of asset | 
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

