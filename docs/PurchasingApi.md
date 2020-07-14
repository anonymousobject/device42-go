# \PurchasingApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePurchases**](PurchasingApi.md#DeletePurchases) | **Delete** /api/1.0/purchases/{id}/ | Delete Purchase
[**GetPurchases**](PurchasingApi.md#GetPurchases) | **Get** /api/1.0/purchases/ | GET method retrieves all Purchases.
[**PostPurchases**](PurchasingApi.md#PostPurchases) | **Post** /api/1.0/purchases/ | Create / Update Purchases
[**PutCustomFieldPurchases**](PurchasingApi.md#PutCustomFieldPurchases) | **Put** /api/1.0/custom_fields/purchases/ | Create/update custom fields for purchases.


# **DeletePurchases**
> interface{} DeletePurchases(ctx, id)
Delete Purchase

This API is used to delete the purchase order with the purchase order id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| purchase id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurchases**
> Purchases GetPurchases(ctx, optional)
GET method retrieves all Purchases.

Get All Purchases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PurchasingApiGetPurchasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchasingApiGetPurchasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseId** | **optional.String**| Device42 Purchase ID | 
 **orderNo** | **optional.String**| Order number | 
 **vendor** | **optional.String**| Vendor name | 
 **costCenter** | **optional.String**| Cost Center | 
 **building** | **optional.String**| Associated building | 
 **completed** | **optional.String**| Line Item is completed | 

### Return type

[**Purchases**](Purchases.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPurchases**
> interface{} PostPurchases(ctx, optional)
Create / Update Purchases

Create / Update Purchases. Required parameters: <ul><li>order_no <b>OR</b> purchase_id</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PurchasingApiPostPurchasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchasingApiPostPurchasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderNo** | **optional.String**| order number / name for the purchase. Can be new or existing. | 
 **purchaseId** | **optional.String**| Can be used instead of order_no to update existing purchase | 
 **lineType** | **optional.String**| required for any new line being added for both device or contract. | 
 **lineNo** | **optional.String**| required for existing line items, use existing line # to change existing line item. | 
 **vendor** | **optional.String**| The cloud vendor | 
 **cost** | **optional.String**|  | 
 **poDate** | **optional.String**|  | 
 **costCenter** | **optional.String**|  | 
 **completed** | **optional.String**|  | 
 **notes** | **optional.String**| Any additional notes | 
 **lineName** | **optional.String**|  | 
 **lineQuantity** | **optional.String**| can be calculated automatically from # of objects associated | 
 **lineCost** | **optional.String**| cost for single object / item. | 
 **lineCostCenter** | **optional.String**|  | 
 **lineCustomer** | **optional.String**|  | 
 **lineItemType** | **optional.String**| Default is device. | 
 **lineAssetIds** | **optional.String**| Comma separated asset_id. Only applicable if line_item_type is asset. | 
 **lineStartDate** | **optional.String**| Date in YYYY-MM-DD format | 
 **lineEndDate** | **optional.String**| Date in YYYY-MM-DD format | 
 **lineFrequency** | **optional.String**|  | 
 **lineRenewDate** | **optional.String**| Date in YYYY-MM-DD format | 
 **lineCancelPolicy** | **optional.String**|  | 
 **lineContractType** | **optional.String**|  | 
 **lineServiceType** | **optional.String**| new service type will be created, if it doesn’t exist (added in v9.0.2) | 
 **lineContractId** | **optional.String**| (added in v9.0.2) | 
 **lineNotes** | **optional.String**|  | 
 **lineCompleted** | **optional.String**|  | 
 **lineDevices** | **optional.String**| Comma separated device names. Only applicable if line_item_type is device. Will create new devices if device with name specific here does not exist. | 
 **lineDeviceSerialNos** | **optional.String**| Comma separated serial numbers. Only applicable if line_item_type is device. Will only work on existing serial numbers. | 
 **lineDeviceAssetNos** | **optional.String**| Comma separated asset numbers. Only applicable if line_item_type is device. Will only work on existing asset numbers. | 
 **lineCircuits** | **optional.String**| circuit ID name | 
 **lineCircuitIds** | **optional.String**| D42 ID for the circuit(s) | 
 **lineBuildingIds** | **optional.String**| D42 ID for the building(s) | 
 **lineCertificateIds** | **optional.String**| D42 ID for the certificate(s) | 
 **lineSoftwareIds** | **optional.String**| D42 ID for the software | 
 **lineRoomIds** | **optional.String**| D42 ID for the room(s) | 
 **lineRackIds** | **optional.String**| D42 ID for the rack(s) | 
 **linePartIds** | **optional.String**| D42 ID for the part(s) | 
 **lineDeviceOsIds** | **optional.String**| D42 ID for the Device OS | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomFieldPurchases**
> interface{} PutCustomFieldPurchases(ctx, orderNo, key, optional)
Create/update custom fields for purchases.

Custom Fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderNo** | **string**| order number / name for the purchase. Can be new or existing. | 
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***PurchasingApiPutCustomFieldPurchasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchasingApiPutCustomFieldPurchasesOpts struct

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

