# \HistoryApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistory**](HistoryApi.md#GetHistory) | **Get** /api/1.0/history/ | Retrieves all history for the past one week. (Note, as of 10.3.0 Device42 now has a more robust audit log available at /api/api/1.0/auditlogs/
[**GetHistoryNumberOfWeeks**](HistoryApi.md#GetHistoryNumberOfWeeks) | **Get** /api/1.0/history/&lt;number_of_weeks&gt;/ | Retrieves history for specified number of weeks in the past.


# **GetHistory**
> []GetHistory GetHistory(ctx, )
Retrieves all history for the past one week. (Note, as of 10.3.0 Device42 now has a more robust audit log available at /api/api/1.0/auditlogs/

Get History

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GetHistory**](getHistory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryNumberOfWeeks**
> []GetHistory GetHistoryNumberOfWeeks(ctx, )
Retrieves history for specified number of weeks in the past.

Get History by # of Weeks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GetHistory**](getHistory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

