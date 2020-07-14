# \OperatingSystemsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeviceOs**](OperatingSystemsApi.md#DeleteDeviceOs) | **Delete** /api/1.0/device_os/{device_os_id}/ | Delete Operating system by OS ID
[**DeleteOperatingSystems**](OperatingSystemsApi.md#DeleteOperatingSystems) | **Delete** /api/1.0/operatingsystems/{id}/ | Delete Operating System
[**GetDeviceOs**](OperatingSystemsApi.md#GetDeviceOs) | **Get** /api/1.0/device_os/ | Get operating systems by devices
[**GetOperatingSystems**](OperatingSystemsApi.md#GetOperatingSystems) | **Get** /api/1.0/operatingsystems/ | This call will get information about operating systems.
[**PostDeviceOs**](OperatingSystemsApi.md#PostDeviceOs) | **Post** /api/1.0/device_os/ | This call will create or update operating systems and assign them to a device.
[**PostOperatingSystems**](OperatingSystemsApi.md#PostOperatingSystems) | **Post** /api/1.0/operatingsystems/ | This call will create/update information about operating systems.


# **DeleteDeviceOs**
> interface{} DeleteDeviceOs(ctx, deviceOsId)
Delete Operating system by OS ID

This API is used to delete the operating system with the operating system id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceOsId** | **int32**| ID of specific operating system | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOperatingSystems**
> interface{} DeleteOperatingSystems(ctx, id)
Delete Operating System

This API is used to delete the operating system with the operating system id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| opearting system id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceOs**
> interface{} GetDeviceOs(ctx, optional)
Get operating systems by devices

This call will get information about operating systems and the devices they’re discovered on.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperatingSystemsApiGetDeviceOsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatingSystemsApiGetDeviceOsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **os** | **optional.String**| filter by OS name (added in v8.3.0) | 
 **osId** | **optional.String**| Operating system ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOperatingSystems**
> interface{} GetOperatingSystems(ctx, optional)
This call will get information about operating systems.

Get all operating systems

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperatingSystemsApiGetOperatingSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatingSystemsApiGetOperatingSystemsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licensedCount** | **optional.String**| Number of licensed instances of operating system | 
 **notLicensedCount** | **optional.String**| Number of discovered instances of operating system not set to licensed | 
 **totalCount** | **optional.String**| Count of IPs returned (use with offset as max results are limited to 1000) | 
 **category** | **optional.String**| Filter by OS category (ie: Linux, Windows) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeviceOs**
> interface{} PostDeviceOs(ctx, optional)
This call will create or update operating systems and assign them to a device.

Create/Update operating systems on devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperatingSystemsApiPostDeviceOsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatingSystemsApiPostDeviceOsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceOsId** | **optional.String**| ID of specific operating system.&lt;br&gt;&lt;br&gt;Use this parameter to change an existing device OS - or -  use the device_id parameter AND the os parameter.&lt;br&gt;&lt;br&gt;Do not use this parameter to create a new device OS. | 
 **deviceId** | **optional.String**| ID of the device the OS is assigned to.&lt;br&gt;&lt;br&gt;Use this parameter AND the os parameter to create a new device OS. | 
 **os** | **optional.String**| Operating system name.&lt;br&gt;&lt;br&gt;Use this parameter to create or change a device OS. See the device parameters above. | 
 **osver** | **optional.String**| Operating system version name | 
 **osverno** | **optional.String**| Operating system version number | 
 **licenseKey** | **optional.String**| OS license key | 
 **countInLicensing** | **optional.String**| Whether or not to count OS in licensing | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOperatingSystems**
> interface{} PostOperatingSystems(ctx, optional)
This call will create/update information about operating systems.

Create/update OS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperatingSystemsApiPostOperatingSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatingSystemsApiPostOperatingSystemsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| name of the OS | 
 **manufacturer** | **optional.String**| name of the hardware/software manufacturer. | 
 **notes** | **optional.String**| Any additional notes | 
 **category** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no, Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. Use for initial insert. | 
 **licensedCount** | **optional.String**| Number of licensed instances of operating system | 
 **licensingModel** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

