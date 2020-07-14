# \ReportsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReports**](ReportsApi.md#GetReports) | **Get** /api/1.0/reports/{id}/ | 


# **GetReports**
> interface{} GetReports(ctx, id, optional)


Run a report and retrieve results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ReportsApiGetReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **csvFormat** | **optional.String**| If this parameter is ‘xls’ or is not present, an Excel file(.xlsx) will be produced. If this parameter is ‘tab’, a tab-delimited file(.txt) will be created. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

