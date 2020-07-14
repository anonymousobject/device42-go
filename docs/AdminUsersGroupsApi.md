# \AdminUsersGroupsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdmingroups**](AdminUsersGroupsApi.md#GetAdmingroups) | **Get** /api/1.0/admingroups/ | Get all Admin Groups
[**GetAdminusers**](AdminUsersGroupsApi.md#GetAdminusers) | **Get** /api/1.0/adminusers/ | Get all Admin Users
[**PostAdminUsers**](AdminUsersGroupsApi.md#PostAdminUsers) | **Post** /api/1.0/adminusers/ | Create/Update Admin Users


# **GetAdmingroups**
> interface{} GetAdmingroups(ctx, )
Get all Admin Groups

Get all Admin Groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdminusers**
> interface{} GetAdminusers(ctx, )
Get all Admin Users

Get Admin Users

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAdminUsers**
> interface{} PostAdminUsers(ctx, username, password, optional)
Create/Update Admin Users

Create/Edit Admin Users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| admin user name | 
  **password** | **string**| new user password | 
 **optional** | ***AdminUsersGroupsApiPostAdminUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminUsersGroupsApiPostAdminUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groups** | **optional.String**| Admin groups to assign user to. Note: groups with commas need to be wrapped in quotes. ie: “System Generated Read, Add, Edit and Delete” | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

