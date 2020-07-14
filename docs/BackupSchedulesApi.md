# \BackupSchedulesApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackupSchedules**](BackupSchedulesApi.md#GetBackupSchedules) | **Get** :4343/api/1.0/backup_schedules/ | Get Backup Schedules
[**PostBackupSchedules**](BackupSchedulesApi.md#PostBackupSchedules) | **Post** :4343/api/1.0/backup_schedules/ | Set Backup Schedules via API


# **GetBackupSchedules**
> interface{} GetBackupSchedules(ctx, )
Get Backup Schedules

Get Backup Schedules

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

# **PostBackupSchedules**
> interface{} PostBackupSchedules(ctx, name, method, scheduleTime, optional)
Set Backup Schedules via API

This will create a scheduled backup nightly at 11:00 PM, using the NFS server settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of backup job | 
  **method** | **string**| Backup method: 1 for mail,2 for sftp,3 for nfs or 4 for amazon s3 | 
  **scheduleTime** | **string**| time to perform backup (0:00 to 23:59) | 
 **optional** | ***BackupSchedulesApiPostBackupSchedulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupSchedulesApiPostBackupSchedulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

