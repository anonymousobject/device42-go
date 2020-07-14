# \ObjectCategoriesApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetObjectCategories**](ObjectCategoriesApi.md#GetObjectCategories) | **Get** /api/1.0/object_categories/ | This call will get information about object categories.
[**PostObjectCategories**](ObjectCategoriesApi.md#PostObjectCategories) | **Post** /api/1.0/object_categories/ | This call will create/update information about object categories.


# **GetObjectCategories**
> interface{} GetObjectCategories(ctx, optional)
This call will get information about object categories.

Get all Object Categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ObjectCategoriesApiGetObjectCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectCategoriesApiGetObjectCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostObjectCategories**
> interface{} PostObjectCategories(ctx, name, optional)
This call will create/update information about object categories.

Create/update Object Categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of object category to create/update | 
 **optional** | ***ObjectCategoriesApiPostObjectCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectCategoriesApiPostObjectCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **optional.String**| Description of object category | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

