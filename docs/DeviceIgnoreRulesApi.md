# \DeviceIgnoreRulesApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20DeviceIgnoreRulesPost**](DeviceIgnoreRulesApi.md#Api20DeviceIgnoreRulesPost) | **Post** /api/2.0/device_ignore_rules/ | This call creates or updates a device ignore rule.
[**DeletedeviceIgnoreRules**](DeviceIgnoreRulesApi.md#DeletedeviceIgnoreRules) | **Delete** /api/2.0/device_ignore_rules/{id}/ | This call deletes a device ignore rule.
[**GetdeviceIgnoreRules**](DeviceIgnoreRulesApi.md#GetdeviceIgnoreRules) | **Get** /api/2.0/device_ignore_rules/ | This call will get all device ignore rules.


# **Api20DeviceIgnoreRulesPost**
> interface{} Api20DeviceIgnoreRulesPost(ctx, text, optionType, optional)
This call creates or updates a device ignore rule.

Create/update device ignore rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **text** | **string**| Text for the ignore rule - in UI, &#39;Ignored text contains&#39;. | 
  **optionType** | **string**| Type of ignore rule for the device - in UI, &#39;Ignore device based on&#39;.&lt;br&gt; 1 &#x3D; Device name&lt;br&gt; 2 &#x3D; OS name&lt;br&gt; 3 &#x3D; Mac address prefix&lt;br&gt; 4 &#x3D; Hardware model | 
 **optional** | ***DeviceIgnoreRulesApiApi20DeviceIgnoreRulesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceIgnoreRulesApiApi20DeviceIgnoreRulesPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.String**| Required to update an existing ignore rule - the D42 ID of the rule. | 
 **lastDiscovery** | **optional.String**| The last time this rule was applied to ignore a device during discovery (YYYY-MM-DD HH:MM). | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletedeviceIgnoreRules**
> interface{} DeletedeviceIgnoreRules(ctx, id)
This call deletes a device ignore rule.

Used to delete a device ignore rule with the D42 rule ID as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| D42 ID of the ignore rule to delete. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetdeviceIgnoreRules**
> interface{} GetdeviceIgnoreRules(ctx, )
This call will get all device ignore rules.

Get all device ignore rules.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

