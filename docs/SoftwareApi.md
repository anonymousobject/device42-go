# \SoftwareApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSoftwareComponent**](SoftwareApi.md#DeleteSoftwareComponent) | **Delete** /api/1.0/software/{id}/ | Delete Software Component
[**DeleteSoftwareDetail**](SoftwareApi.md#DeleteSoftwareDetail) | **Delete** /api/1.0/software_details/{id}/ | Delete Software detail
[**DeleteSoftwareLicenseKeys**](SoftwareApi.md#DeleteSoftwareLicenseKeys) | **Delete** /api/1.0/software/license_keys/{id}/ | Delete Software License Keys
[**GetSoftwareComponentDetails**](SoftwareApi.md#GetSoftwareComponentDetails) | **Get** /api/1.0/software/ | Get Software Component details
[**GetSoftwareDetails**](SoftwareApi.md#GetSoftwareDetails) | **Get** /api/1.0/software_details/ | 
[**GetSoftwareLicenseKeys**](SoftwareApi.md#GetSoftwareLicenseKeys) | **Get** /api/1.0/software/license_keys/ | Get Software License Keys
[**PostUpdateServicePorts**](SoftwareApi.md#PostUpdateServicePorts) | **Post** /api/1.0/software_details/ | 
[**PostUpdateSoftwareComponents**](SoftwareApi.md#PostUpdateSoftwareComponents) | **Post** /api/1.0/software/ | Create or update software component.
[**PostUpdateSoftwareLicenses**](SoftwareApi.md#PostUpdateSoftwareLicenses) | **Post** /api/1.0/software/license_keys/ | 


# **DeleteSoftwareComponent**
> interface{} DeleteSoftwareComponent(ctx, id)
Delete Software Component

This API is used to delete the software detail with the software detail id supplied as the required argument. Note: You will only be able to delete the software if the supplied username has the correct permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| software detail id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSoftwareDetail**
> interface{} DeleteSoftwareDetail(ctx, id)
Delete Software detail

This API is used to delete the software detail with the software detail id supplied as the required argument. Note: You will only be able to delete the software if the supplied username has the correct permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| software detail id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSoftwareLicenseKeys**
> interface{} DeleteSoftwareLicenseKeys(ctx, id)
Delete Software License Keys

This API is used to delete the software license key with the software license key id supplied as the required argument. Note: You will only be able to delete the software key if the supplied username has the correct permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| software detail id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoftwareComponentDetails**
> interface{} GetSoftwareComponentDetails(ctx, optional)
Get Software Component details

You can filter software details by following parameters in the query string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareApiGetSoftwareComponentDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiGetSoftwareComponentDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **category** | **optional.String**| name of the category | 
 **vendor** | **optional.String**| Software vendor | 
 **licensingModel** | **optional.String**|  | 
 **softwareType** | **optional.String**| Filter by software type (managed, unmanaged, prohibited or ignored) | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 
 **tagsAnd** | **optional.String**| filter by all the tags, separated by comma. (This is an AND filter and all tags have to match for filter, this was added in v6.3.1) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoftwareDetails**
> interface{} GetSoftwareDetails(ctx, optional)


Get Software details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareApiGetSoftwareDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiGetSoftwareDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.String**| filter by id of device | 
 **device** | **optional.String**| Device name | 
 **softwareId** | **optional.String**| filter by id of the software | 
 **softwareDetailId** | **optional.String**| filter by id of the software | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoftwareLicenseKeys**
> interface{} GetSoftwareLicenseKeys(ctx, optional)
Get Software License Keys

You can filter software license keys by following parameters in the query string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareApiGetSoftwareLicenseKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiGetSoftwareLicenseKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **softwareId** | **optional.String**| filter by id of the software | 
 **softwareName** | **optional.String**| filter by name of the software component | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdateServicePorts**
> interface{} PostUpdateServicePorts(ctx, software, device, optional)


Create / Update software details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **software** | **string**| the name of the software | 
  **device** | **string**| The name of the device where this software is installed | 
 **optional** | ***SoftwareApiPostUpdateServicePortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiPostUpdateServicePortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version number of the software | 
 **user** | **optional.String**| The user assigned to this software | 
 **vendor** | **optional.String**| The vendor that created the server, linked to Organization | 
 **installDate** | **optional.String**| The date that the software was installed | 
 **countInLicensing** | **optional.String**| Whether or not to count OS in licensing | 
 **licenseKey** | **optional.String**| OS license key | 
 **licenseKeyCount** | **optional.String**| The number of licenses this software key supports | 
 **licenseUseCount** | **optional.String**| the number of licenses that are in use for this software instance | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdateSoftwareComponents**
> interface{} PostUpdateSoftwareComponents(ctx, licensingModel, softwareType, optional)
Create or update software component.

Create / Update Software Components. Required parameters: <ul><li>name <b>OR</b> id</li> <li>licensing_model</li> <li>software_type</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **licensingModel** | **string**| Custom models can be made via UI. Existing values include: Free, Trial, Individual - User/Perpetual, Individual - User/Subscription, Named User / Perpetual, Volume - User/Perpetual, Concurrent - User/Perpetual, Individual - Device/Perpetual, Individual - Device/Subscription, Node Locked / Perpetual, Volume - Device/Perpetual, OEM, CAL / Per Seat Device, CAL / Per Seat User, CAL / Per Server, CAL / Per Processor, CAL / Per Mailbox | [default to Individual - Device/Perpetual]
  **softwareType** | **string**| Default software type is Managed. | [default to Managed]
 **optional** | ***SoftwareApiPostUpdateSoftwareComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiPostUpdateSoftwareComponentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| The name of the software (new or existing) | 
 **id** | **optional.String**| The ID of the software, required if not using NAME | 
 **category** | **optional.String**| Filter by user-defined software categories (new or existing) | 
 **vendor** | **optional.String**| Software Vendor (new or existing) | 
 **tags** | **optional.String**|  | 
 **tagsRemove** | **optional.String**| remove tags from component | 
 **trackLicensedCountByKeys** | **optional.String**| whether or not to track software by discovered count | 
 **notes** | **optional.String**| Any additional notes | 
 **aliases** | **optional.String**| any software aliases | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdateSoftwareLicenses**
> interface{} PostUpdateSoftwareLicenses(ctx, optional)


Create / Update Software Licenses. Required parameters: <ul><li>id <b>OR</b> software_id <b>OR</b> software_name</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareApiPostUpdateSoftwareLicensesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareApiPostUpdateSoftwareLicensesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| The ID of the software_license_key object (use if updating) | 
 **softwareId** | **optional.String**| The id of the software component | 
 **softwareName** | **optional.String**| software component name | 
 **key** | **optional.String**| software license key | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

