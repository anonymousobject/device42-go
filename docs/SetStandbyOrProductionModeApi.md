# \SetStandbyOrProductionModeApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAppliancemode**](SetStandbyOrProductionModeApi.md#PostAppliancemode) | **Post** :4343/api/1.0/appliancemode/ | Set Standby or Production Mode


# **PostAppliancemode**
> interface{} PostAppliancemode(ctx, applianceMode)
Set Standby or Production Mode

This will set the Device42 appliance to Standby or Production mode. Used in conjunction with automatic backups and restores for a warm HA solution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applianceMode** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

