# \PartsManagementApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePartmodels**](PartsManagementApi.md#DeletePartmodels) | **Delete** /api/1.0/partmodels/{id}/ | Delete Part Model
[**DeleteParts**](PartsManagementApi.md#DeleteParts) | **Delete** /api/1.0/parts/{id}/ | Delete Part
[**GetPartmodels**](PartsManagementApi.md#GetPartmodels) | **Get** /api/1.0/partmodels/ | Get all Part Models - introduced in version 5.7.2
[**GetParts**](PartsManagementApi.md#GetParts) | **Get** /api/1.0/parts/ | Get all Parts - introduced in version 5.7.2
[**PostPartmodels**](PartsManagementApi.md#PostPartmodels) | **Post** /api/1.0/partmodels/ | Create / Update Part Models - introduced in version 5.7.2
[**PostParts**](PartsManagementApi.md#PostParts) | **Post** /api/1.0/parts/ | Create / Update Parts - introduced in version 5.7.2
[**PutCustomFieldPart**](PartsManagementApi.md#PutCustomFieldPart) | **Put** /api/1.0/custom_fields/part/ | Create/update existing custom fields for parts.
[**PutCustomFieldPartmodel**](PartsManagementApi.md#PutCustomFieldPartmodel) | **Put** /api/1.0/custom_fields/partmodel/ | Create/update custom fields for Part Models.


# **DeletePartmodels**
> interface{} DeletePartmodels(ctx, id)
Delete Part Model

This API is used to delete the part model with the part model id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| part model id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteParts**
> interface{} DeleteParts(ctx, id)
Delete Part

This API is used to delete the part with the part id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| part id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPartmodels**
> Partmodels GetPartmodels(ctx, optional)
Get all Part Models - introduced in version 5.7.2

Get all Part Models

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PartsManagementApiGetPartmodelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartsManagementApiGetPartmodelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 

### Return type

[**Partmodels**](Partmodels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParts**
> interface{} GetParts(ctx, optional)
Get all Parts - introduced in version 5.7.2

Get all Parts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PartsManagementApiGetPartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartsManagementApiGetPartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| type of the partmodel, cpu, mem, hdd for CPU, memory, and Harddisk. For others it must match the type name for part model | 
 **device** | **optional.String**| name of the device where part is checked out to | 
 **deviceId** | **optional.String**| id of the device where part is checked out to | 
 **deviceSerial** | **optional.String**| serial number of the device, where part is checked out to | 
 **room** | **optional.String**| name of the room where part is checked out to | 
 **roomId** | **optional.String**| id of the room where part is checked out to | 
 **partId** | **optional.String**| id of the part (added in v6.3.3) | 
 **partmodelId** | **optional.String**| id of the part model (added in v6.3.3) | 
 **serialNo** | **optional.String**| serial number of the part | 
 **lastUpdatedLt** | **optional.String**| last updated less than date YYYY-MM-DD format | 
 **lastUpdatedGt** | **optional.String**| last updated greater than date YYYY-MM-DD format | 
 **rackId** | **optional.String**| id of the rack where part is checked out to | 
 **rack** | **optional.String**| name of the rack where part is checked out to | 
 **assetNo** | **optional.String**| filter by asset # (Added in v6.0.0) | 
 **customFieldsAnd** | **optional.String**| filter by custom fields, and filter, format of key1:value1,key2:value2 | 
 **customFieldsOr** | **optional.String**| filter by custom fields, or filter, format of key1:value1,key2:value2 | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPartmodels**
> interface{} PostPartmodels(ctx, optional)
Create / Update Part Models - introduced in version 5.7.2

Create / Update Part Models

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PartsManagementApiPostPartmodelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartsManagementApiPostPartmodelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Part type - new or existing. Must be hdd or harddisk to update HDD model parameters (hddsize, hddtype, etc) | 
 **name** | **optional.String**| name of part model - new or existing | 
 **description** | **optional.String**|  | 
 **partmodelId** | **optional.String**| use for updating existing part model | 
 **modelno** | **optional.String**| Model # of the part model | 
 **partno** | **optional.String**| Part # of the part model | 
 **manufacturer** | **optional.String**| name of the hardware/software manufacturer. | 
 **mediaType** | **optional.String**| Type of media | 
 **connectorType** | **optional.String**| Type of connector, ie rj45 | 
 **length** | **optional.String**| Cable length | 
 **connectivity** | **optional.String**| New or existing (not used for CPU, RAM, HDD) | 
 **totalCount** | **optional.String**|  | 
 **location** | **optional.String**| Location/region of instance deployment | 
 **notes** | **optional.String**| Any additional notes | 
 **cores** | **optional.String**| number of cores | 
 **cpuspeed** | **optional.String**| enter in MHZ, e.g.: 3.5 GHZ use 3500 | 
 **threads** | **optional.String**| number of threads | 
 **ramsize** | **optional.String**| enter in MB, e.g.: 8 GB enter 8192 | 
 **ramtype** | **optional.String**| e.g.: DDR3 | 
 **ramspeed** | **optional.String**| e.g.: 1600 | 
 **hddsize** | **optional.String**| enter in GB, e.g.: 250 GB enter 250 | 
 **hddtype** | **optional.String**| new or existing | 
 **hddrpm** | **optional.String**| new or existing | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostParts**
> interface{} PostParts(ctx, optional)
Create / Update Parts - introduced in version 5.7.2

Create / Update Parts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PartsManagementApiPostPartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartsManagementApiPostPartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Part type - new or existing. Must be hdd or harddisk to update HDD model parameters (hddsize, hddtype, etc) | 
 **name** | **optional.String**| name of part model - new or existing | 
 **count** | **optional.Int32**| number of parts | 
 **partId** | **optional.String**| Use for updating existing part | 
 **partmodelId** | **optional.String**|  | 
 **serialNo** | **optional.String**| Use for updating existing part. Caution: will update first matching serial if multiple parts with same serial exist. Use part_id or partmodel_id to uniquely identify. | 
 **firmware** | **optional.String**|  | 
 **assignment** | **optional.String**| room, device, rma - required if assigning device | 
 **room** | **optional.String**| Room name - required if assigned to room | 
 **device** | **optional.String**| Room name - required if assigned to device | 
 **dateChanged** | **optional.String**| Update the Date Changed field, using format YYYY-MM-DD HH:MM:SS | 
 **description** | **optional.String**|  | 
 **cores** | **optional.String**| number of cores | 
 **cpuspeed** | **optional.String**| enter in MHZ, e.g.: 3.5 GHZ use 3500 | 
 **threads** | **optional.String**| number of threads | 
 **ramsize** | **optional.String**| enter in MB, e.g.: 8 GB enter 8192 | 
 **ramtype** | **optional.String**| e.g.: DDR3 | 
 **ramspeed** | **optional.String**| e.g.: 1600 | 
 **hddsize** | **optional.String**| enter in GB, e.g.: 250 GB enter 250 | 
 **hddtype** | **optional.String**| new or existing | 
 **hddrpm** | **optional.String**| new or existing | 
 **raidType** | **optional.String**| type of RAID | 
 **raidGroup** | **optional.String**| RAID group name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomFieldPart**
> interface{} PutCustomFieldPart(ctx, id, key, optional)
Create/update existing custom fields for parts.

Create/updated custom fields for parts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of part. | 
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***PartsManagementApiPutCustomFieldPartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartsManagementApiPutCustomFieldPartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

# **PutCustomFieldPartmodel**
> interface{} PutCustomFieldPartmodel(ctx, id, key, optional)
Create/update custom fields for Part Models.

Custom Fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of part. | 
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***PartsManagementApiPutCustomFieldPartmodelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartsManagementApiPutCustomFieldPartmodelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

