# \ServiceLevelsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceLevel**](ServiceLevelsApi.md#GetServiceLevel) | **Get** /api/1.0/service_level/ | Get all Service Levels (devices)
[**PostServiceLevel**](ServiceLevelsApi.md#PostServiceLevel) | **Post** /api/1.0/service_level/ | Create Service Level.


# **GetServiceLevel**
> []InlineResponse20011 GetServiceLevel(ctx, )
Get all Service Levels (devices)

Get all Service Levels

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceLevel**
> interface{} PostServiceLevel(ctx, name)
Create Service Level.

Create Service Level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the Service Level | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

