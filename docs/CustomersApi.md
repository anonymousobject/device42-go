# \CustomersApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomers**](CustomersApi.md#DeleteCustomers) | **Delete** /api/1.0/customers/{id}/ | Delete Customer
[**GetCustomers**](CustomersApi.md#GetCustomers) | **Get** /api/1.0/customers/ | Get all Customers
[**PostCustomers**](CustomersApi.md#PostCustomers) | **Post** /api/1.0/customers/ | 
[**PostUpdateCustomerContacts**](CustomersApi.md#PostUpdateCustomerContacts) | **Post** /api/1.0/customers/contacts/ | 
[**PutCustomFields**](CustomersApi.md#PutCustomFields) | **Put** /api/1.0/custom_fields/customer/ | 


# **DeleteCustomers**
> interface{} DeleteCustomers(ctx, id)
Delete Customer

This API is used to delete the customer with the customer id supplied as the required argument.

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

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomers**
> interface{} GetCustomers(ctx, optional)
Get all Customers

Get all Customers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomersApiGetCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiGetCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeCols** | **optional.String**| do not return all columns just the ones specified. For example, ?include_cols&#x3D;name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomers**
> interface{} PostCustomers(ctx, name, optional)


Create / Update Customers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Customer name | 
 **optional** | ***CustomersApiPostCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiPostCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactInfo** | **optional.String**| Text field for contact information | 
 **notes** | **optional.String**| Any additional notes | 
 **type_** | **optional.String**|  | 
 **newName** | **optional.String**| New name for an existing customer &#39;name&#39;.&lt;br&gt;If the customer &#39;name&#39; does not exist, this creates a new customer. | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdateCustomerContacts**
> interface{} PostUpdateCustomerContacts(ctx, name, type_, customer, optional)


Create / Update Customer Contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Customer contact name | 
  **type_** | **string**| Contact type, must already exist | 
  **customer** | **string**| Customer name | 
 **optional** | ***CustomersApiPostUpdateCustomerContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiPostUpdateCustomerContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **email** | **optional.String**| Text field. | 
 **phone** | **optional.String**| Text field. | 
 **address** | **optional.String**| Text field | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomFields**
> interface{} PutCustomFields(ctx, name, key, optional)


Custom Fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the customer | 
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***CustomersApiPutCustomFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiPutCustomFieldsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD | 
 **relatedFieldName** | **optional.String**| Required if type &#x3D; related field. | 
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

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

