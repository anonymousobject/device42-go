# \DeviceNameProfileApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api10DeviceNameProfilePost**](DeviceNameProfileApi.md#Api10DeviceNameProfilePost) | **Post** /api/1.0/device_name_profile/ | This call creates or updates a device name profile.
[**DeletedeviceProfileNames**](DeviceNameProfileApi.md#DeletedeviceProfileNames) | **Delete** /api/1.0/device_name_profile/{id}/ | This call deletes a device name profile.
[**GetdeviceNameProfile**](DeviceNameProfileApi.md#GetdeviceNameProfile) | **Get** /api/1.0/device_name_profile/ | This call will get all device name profiles.


# **Api10DeviceNameProfilePost**
> interface{} Api10DeviceNameProfilePost(ctx, name, numberLength, startNumber, lastUsed, optional)
This call creates or updates a device name profile.

Create/update device name profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Used to easily identify different profiles for generating device names. | 
  **numberLength** | **int32**| Number length between 2 and 7. For numbers up to 9999, it would be 4. Default is 4. | 
  **startNumber** | **int32**| Starting number, where you want device numbers to start from. Default is 1. | 
  **lastUsed** | **int32**| DO NOT CHANGE, unless you know what it is used for :). Default is 0. | 
 **optional** | ***DeviceNameProfileApiApi10DeviceNameProfilePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceNameProfileApiApi10DeviceNameProfilePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **prefix** | **optional.String**| Prefix for the generated name number; e.g., SecretDC-Server in SecretDC-Server0042. | 
 **suffix** | **optional.String**| Suffix for the generated name number; e.g., .domain.pvt in SecretDC-Server0042.domain.pvt. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletedeviceProfileNames**
> interface{} DeletedeviceProfileNames(ctx, id)
This call deletes a device name profile.

This API is used to delete a device name profile with the device name profile id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| device name profile id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetdeviceNameProfile**
> interface{} GetdeviceNameProfile(ctx, optional)
This call will get all device name profiles.

Get all device name profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceNameProfileApiGetdeviceNameProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceNameProfileApiGetdeviceNameProfileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Used to easily identify different profiles for generating device names. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

