# \ApplicationComponentsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppcomps**](ApplicationComponentsApi.md#DeleteAppcomps) | **Delete** /api/1.0/appcomps/{appcomp_id}/ | Delete Application Component
[**GetAppcomps**](ApplicationComponentsApi.md#GetAppcomps) | **Get** /api/1.0/appcomps/ | 
[**GetAppcompsAppcompId**](ApplicationComponentsApi.md#GetAppcompsAppcompId) | **Get** /api/1.0/appcomps/{appcomp_id}/ | 
[**PostAppcomps**](ApplicationComponentsApi.md#PostAppcomps) | **Post** /api/1.0/appcomps/ | 
[**PutCustomFieldAppcomp**](ApplicationComponentsApi.md#PutCustomFieldAppcomp) | **Put** /api/1.0/custom_fields/appcomp/ | API to create/update custom fields for application components.


# **DeleteAppcomps**
> interface{} DeleteAppcomps(ctx, appcompId)
Delete Application Component

This API is used to delete the application component with the application component id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appcompId** | **int32**| IP Address id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppcomps**
> interface{} GetAppcomps(ctx, optional)


Get Application Components

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationComponentsApiGetAppcompsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationComponentsApiGetAppcompsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.String**| filter by id of device | 
 **device** | **optional.String**| Device name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppcompsAppcompId**
> interface{} GetAppcompsAppcompId(ctx, appcompId)


Get Application Components by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appcompId** | [**string**](.md)| Appcomp id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAppcomps**
> interface{} PostAppcomps(ctx, name, device, optional)


Create / Update Application Components

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| unique name for component | 
  **device** | **string**| device (name) that this component is dependent on | 
 **optional** | ***ApplicationComponentsApiPostAppcompsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationComponentsApiPostAppcompsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupOwner** | **optional.String**| Name of group that is responsible for this component - must match group name exactly. | 
 **what** | **optional.String**| Description of business impact due to loss of component. | 
 **dependsOn** | **optional.String**| Names of app components this component depends on, separated by commas - must match component names exactly. | 
 **dependents** | **optional.String**| Names of app components that depend on this component separated by commas - must match component names exactly. | 
 **groupsAffected** | **optional.String**| Names of affected groups separated by commas - must match group names exactly. | 
 **deviceReason** | **optional.String**| string for the device reason on this appcomp (added in v6.6.0) | 
 **dependsOnReasons** | **optional.String**| list of string pairs for dependent appcomps on this appcomp e.g. &#x3D;&gt; depend_appcomp_name1:reason1, depend_appcomp_name2:reason2, depend_appcomp_nameN:reason3 (added in v6.6.0) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomFieldAppcomp**
> interface{} PutCustomFieldAppcomp(ctx, key, optional)
API to create/update custom fields for application components.

Required parameters: <ul><li>id <b>OR</b> name</li> <li>key</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***ApplicationComponentsApiPutCustomFieldAppcompOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationComponentsApiPutCustomFieldAppcompOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**| ID of application component | 
 **name** | **optional.String**| name of application component | 
 **type_** | **optional.String**| this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD | 
 **relatedFieldName** | **optional.String**| Required if type &#x3D; related field. | 
 **addToPicklist** | **optional.String**| Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored. | 
 **value** | **optional.String**| This will set the value of the custom field for the specific object. | 
 **clearValue** | **optional.String**| yes to clear existing value for that field | 
 **showOnChart** | **optional.String**| Show the field on impact charts | 
 **notes** | **optional.String**| Any additional notes | 
 **clearNotes** | **optional.String**| Yes to clear any existing notes. | 
 **bulkFields** | **optional.String**| comma separated key value pairs, with key and value separated by colon. e.g.key1:value1, key2:value2 | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

