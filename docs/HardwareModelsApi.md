# \HardwareModelsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHardwares**](HardwareModelsApi.md#DeleteHardwares) | **Delete** /api/1.0/hardwares/{id}/ | Delete Hardware Model
[**GetHardwares**](HardwareModelsApi.md#GetHardwares) | **Get** /api/1.0/hardwares/ | This call will get information about hardware models.
[**PostHardwares**](HardwareModelsApi.md#PostHardwares) | **Post** /api/1.0/hardwares/ | This call will create/update information about hardware models.


# **DeleteHardwares**
> interface{} DeleteHardwares(ctx, id)
Delete Hardware Model

This API is used to delete the hardware model with the hardware model id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| hardware model id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHardwares**
> interface{} GetHardwares(ctx, optional)
This call will get information about hardware models.

Get all hardware models

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwareModelsApiGetHardwaresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwareModelsApiGetHardwaresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **type_** | **optional.String**| could be physical, blade, or other | 
 **deviceSubType** | **optional.String**| filter by device sub type (Added in v14.7.2) | 
 **size** | **optional.String**| filter by exact size | 
 **depth** | **optional.String**| could be half or full | 
 **partNo** | **optional.String**| filter by part # | 
 **watts** | **optional.String**| filter by exact watts | 
 **manufacturer** | **optional.String**| name of the hardware manufacturer. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostHardwares**
> interface{} PostHardwares(ctx, name, optional)
This call will create/update information about hardware models.

Create/update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| if similar hardware name already exists, first matching entry is updated | 
 **optional** | ***HardwareModelsApiPostHardwaresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwareModelsApiPostHardwaresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newName** | **optional.String**| Use to change name of object. | 
 **type_** | **optional.Int32**| 1 &#x3D; Regular, 2 &#x3D; Blade, 3 &#x3D; Other | 
 **deviceSubtype** | **optional.String**| Subtype of \&quot;other\&quot; type devices | 
 **size** | **optional.String**| Size in U for hardware type regular | 
 **widthRatio** | **optional.String**| Default&#x3D;1. Can be ½, 1/3,… 1/10, etc. | 
 **depth** | **optional.String**| half by default, full to override | 
 **bladeSize** | **optional.Int32**| 1&#x3D;Full Height 2&#x3D;Half Height 3&#x3D;Double Half Height 4&#x3D;Double Full Height 5&#x3D;Quarter Height | 
 **partNo** | **optional.String**|  | 
 **watts** | **optional.String**| per power supply | 
 **specUrl** | **optional.String**| Specification url for the hardware model. | 
 **manufacturer** | **optional.String**| name of the hardware/software manufacturer. | 
 **frontImageId** | **optional.String**|  | 
 **frontImage** | **optional.String**| name of the image file (Added in v5.8.2) | 
 **backImageId** | **optional.String**| back image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now. | 
 **backImage** | **optional.String**| name of the back image file. Use instead of back_image_id. | 
 **notes** | **optional.String**| Any additional notes | 
 **maxBladesPerRow** | **optional.String**|  | 
 **slotNumbering** | **optional.String**|  | 
 **modulePos** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

