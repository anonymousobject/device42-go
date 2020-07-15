# \AssetDeviceLifeCycleApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLifecycleEvent**](AssetDeviceLifeCycleApi.md#GetLifecycleEvent) | **Get** /api/1.0/lifecycle_event/ | Retrieve Life Cycle Events using filters - introduced in version 5.5.7
[**PutLifecycleEvent**](AssetDeviceLifeCycleApi.md#PutLifecycleEvent) | **Put** /api/1.0/lifecycle_event/ | Create LifeCycle Events


# **GetLifecycleEvent**
> interface{} GetLifecycleEvent(ctx, optional)
Retrieve Life Cycle Events using filters - introduced in version 5.5.7

Get Life Cycle Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetDeviceLifeCycleApiGetLifecycleEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetDeviceLifeCycleApiGetLifecycleEventOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| filter by existing event type | 
 **device** | **optional.String**| filter by device name | 
 **asset** | **optional.String**| filter by asset name | 
 **enduser** | **optional.String**| filter by end user name | 
 **dateGt** | **optional.String**| filter by date greater than (YYYY-MM-DD) | 
 **dateLt** | **optional.String**| filter by date less than (YYYY-MM-DD) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLifecycleEvent**
> interface{} PutLifecycleEvent(ctx, date, type_, optional)
Create LifeCycle Events

Use this API to create life cycle events for devices or assets. Use device, device_id, asset_no, serial_no, or asset_id to indicate the device or asset the event is to be created for.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**| in YYYY-MM-DD or&lt;br&gt;YYYY-MM-DD HH:MM format. | 
  **type_** | **string**| must be defined already in device42. | 
 **optional** | ***AssetDeviceLifeCycleApiPutLifecycleEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetDeviceLifeCycleApiPutLifecycleEventOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | **optional.String**| Name of the device that the event is for | 
 **deviceId** | **optional.String**| ID of the device that the event is for | 
 **assetNo** | **optional.String**| Asset number of the device that the event is for | 
 **serialNo** | **optional.String**| Serial number of the device that the event is for | 
 **assetId** | **optional.String**| ID of the asset that the event is for | 
 **user** | **optional.String**| enduser name | 
 **notes** | **optional.String**| Any additional notes | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

