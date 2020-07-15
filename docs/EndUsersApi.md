# \EndUsersApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEndUsers**](EndUsersApi.md#DeleteEndUsers) | **Delete** /api/1.0/endusers/{id}/ | Delete End User
[**GetEndusers**](EndUsersApi.md#GetEndusers) | **Get** /api/1.0/endusers/ | Get all End Users
[**PostEndUsers**](EndUsersApi.md#PostEndUsers) | **Post** /api/1.0/endusers/ | Create / Update End Users


# **DeleteEndUsers**
> interface{} DeleteEndUsers(ctx, id)
Delete End User

This API is used to delete the end user with the end user id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| end user id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndusers**
> interface{} GetEndusers(ctx, )
Get all End Users

Get End Users

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

# **PostEndUsers**
> interface{} PostEndUsers(ctx, name, optional)
Create / Update End Users

Create / Update End Users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| enduser name | 
 **optional** | ***EndUsersApiPostEndUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndUsersApiPostEndUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **optional.String**| Text field. | 
 **contact** | **optional.String**|  | 
 **location** | **optional.String**|  | 
 **notes** | **optional.String**| Any additional notes | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

