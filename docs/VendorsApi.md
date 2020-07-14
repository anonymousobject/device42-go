# \VendorsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVendors**](VendorsApi.md#DeleteVendors) | **Delete** /api/1.0/vendors/{id}/ | Delete Vendor
[**GetVendors**](VendorsApi.md#GetVendors) | **Get** /api/1.0/vendors/ | Get all Vendors
[**PostVendors**](VendorsApi.md#PostVendors) | **Post** /api/1.0/vendors/ | 


# **DeleteVendors**
> interface{} DeleteVendors(ctx, id)
Delete Vendor

This API is used to delete the vendor with the vendor id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| IP Address id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVendors**
> interface{} GetVendors(ctx, )
Get all Vendors

Get all Vendors

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

# **PostVendors**
> interface{} PostVendors(ctx, name, optional)


Create / Update Vendors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Vendor name | 
 **optional** | ***VendorsApiPostVendorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorsApiPostVendorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountNo** | **optional.String**|  | 
 **homePage** | **optional.String**| Text field. | 
 **contactInfo** | **optional.String**|  | 
 **escalation1** | **optional.String**| Text field. | 
 **escalation2** | **optional.String**| Text field. | 
 **notes** | **optional.String**| Any additional notes | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

