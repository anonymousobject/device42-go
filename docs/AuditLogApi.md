# \AuditLogApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditlogs**](AuditLogApi.md#GetAuditlogs) | **Get** /api/1.0/auditlogs/ | Retrieves the audit history for all changes made in Device42.


# **GetAuditlogs**
> GetAuditLogs GetAuditlogs(ctx, optional)
Retrieves the audit history for all changes made in Device42.

Get Audit Logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditLogApiGetAuditlogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditLogApiGetAuditlogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectId** | **optional.String**| Filters by object ID (ie, device ID, asset ID) | 
 **contentType** | **optional.String**| Returns changes done to a particular content type | 
 **actionTimeGt** | **optional.String**| Filters actions that have happened past the time entered (ie, greater than 2 weeks) in YYYY-MM-DDTHH:MM:ss.uuuuuu (ie 2016-10-27T13:52:01.213416) | 
 **actionTimeLt** | **optional.String**| Returns actions within the last X amount of days in YYYY-MM-DDTHH:MM:ss.uuuuuu (ie 2016-10-27T13:52:01.213416) | 

### Return type

[**GetAuditLogs**](getAuditLogs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

