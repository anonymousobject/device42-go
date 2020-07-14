# \PasswordsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePassword**](PasswordsApi.md#DeletePassword) | **Delete** /api/1.0/passwords/{id}/ | Delete Password
[**GetPassword**](PasswordsApi.md#GetPassword) | **Get** /api/1.0/passwords/ | Retrieve accounts (usernames) and passwords.
[**PostCustomFields**](PasswordsApi.md#PostCustomFields) | **Post** /api/1.0/custom_fields/password/ | 
[**PostUpdatePasswords**](PasswordsApi.md#PostUpdatePasswords) | **Post** /api/1.0/passwords/ | 


# **DeletePassword**
> interface{} DeletePassword(ctx, id)
Delete Password

This API is used to delete the password with the password id supplied as the required argument. Note: You will only be able to delete the password if the supplied username has the correct permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| password id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPassword**
> interface{} GetPassword(ctx, optional)
Retrieve accounts (usernames) and passwords.

Get Usernames and Passwords

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PasswordsApiGetPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PasswordsApiGetPasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **optional.String**| name of the category | 
 **label** | **optional.String**|  | 
 **username** | **optional.String**| Retrieve all the passwords for a specified username. ?username&#x3D; | 
 **device** | **optional.String**| Device name | 
 **appcomp** | **optional.String**| The application component that depends on this service | 
 **id** | **optional.String**| The ID of the software, required if not using NAME | 
 **plainText** | **optional.String**| Decrypt the password and return the plain text version. ?plain_text&#x3D;yes will decrypt and display the password. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomFields**
> interface{} PostCustomFields(ctx, username, key, optional)


Custom Fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Name of the password | 
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***PasswordsApiPostCustomFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PasswordsApiPostCustomFieldsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD | 
 **relatedFieldName** | **optional.String**| Required if type &#x3D; related field. | 
 **addToPicklist** | **optional.String**| Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored. | 
 **value** | **optional.String**| This will set the value of the custom field for the specific object. | 
 **clearValue** | **optional.String**| yes to clear existing value for that field | 
 **notes** | **optional.String**| Any additional notes | 
 **clearNotes** | **optional.String**| Yes to clear any existing notes. | 
 **bulkFields** | **optional.String**| comma separated key value pairs, with key and value separated by colon. e.g.key1:value1, key2:value2 | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdatePasswords**
> interface{} PostUpdatePasswords(ctx, optional)


Create / Update Passwords. Use id if updating existing password - else, username and password are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PasswordsApiPostUpdatePasswordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PasswordsApiPostUpdatePasswordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Required to modify existing password | 
 **username** | **optional.String**|  | 
 **password** | **optional.String**| The password. | 
 **label** | **optional.String**|  | 
 **category** | **optional.String**| A password category. If it doesn&#39;t exist a new one will be created. | 
 **devices** | **optional.String**| A comma separated list of device names. | 
 **appcomps** | **optional.String**| A comma separated list of application component names. | 
 **daysBeforeExpiry** | **optional.String**| number of days before password is set as expired | 
 **notes** | **optional.String**| Any additional notes | 
 **viewUsers** | **optional.String**| A comma separated list of users that have permission to view this password. | 
 **viewGroups** | **optional.String**| A comma separated list of user groups that have permission to view this password. | 
 **viewEditUsers** | **optional.String**| A comma separated list of users that have permission to view and edit this password. | 
 **viewUsersRemove** | **optional.String**| A comma separated list of users to remove view permissions. | 
 **useOnlyUsersRemove** | **optional.String**| A comma separated list of users to remove use only permissions. | 
 **viewEditUsersRemove** | **optional.String**| A comma separated list of users to remove view and edit permissions. | 
 **viewGroupsRemove** | **optional.String**| A comma separated list of user groups to remove use permissions. | 
 **useOnlyGroupsRemove** | **optional.String**| A comma separated list of user groups to remove use permissions. | 
 **viewEditGroupsRemove** | **optional.String**| A comma separated list of user groups to remove view and edit permissions. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

