# \TelcoCircuitsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCircuit**](TelcoCircuitsApi.md#DeleteCircuit) | **Delete** /api/1.0/circuits/{id}/ | Delete Circuit
[**GetCircuits**](TelcoCircuitsApi.md#GetCircuits) | **Get** /api/1.0/circuits/ | Get all Circuits
[**PostUpdateCircuits**](TelcoCircuitsApi.md#PostUpdateCircuits) | **Post** /api/1.0/circuits/ | Create / Update Circuits
[**PutCustomFields**](TelcoCircuitsApi.md#PutCustomFields) | **Put** /api/1.0/custom_fields/circuit/ | Create/update custom fields for circuits. (Added in v5.9.3)


# **DeleteCircuit**
> interface{} DeleteCircuit(ctx, id)
Delete Circuit

This API is used to delete the circuit with the circuit id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| circuit id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCircuits**
> interface{} GetCircuits(ctx, optional)
Get all Circuits

Get all Circuits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TelcoCircuitsApiGetCircuitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelcoCircuitsApiGetCircuitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iD** | **optional.String**| ID for the cricuit (D42 ID) | 
 **circuitId** | **optional.String**| circuit id | 
 **type_** | **optional.String**|  | 
 **orderDate** | **optional.String**| Order number | 
 **provisionDate** | **optional.String**| Text field. | 
 **turnOnDate** | **optional.String**|  | 
 **notes** | **optional.String**| Any additional notes | 
 **bandwidth** | **optional.String**| bandwidth in kbps (mutiply factor of 1024) | 
 **vendor** | **optional.String**| The cloud vendor | 
 **customer** | **optional.String**| filter by customer name | 
 **dramcId** | **optional.String**| ID of the DMARC | 
 **dmarc** | **optional.String**| DMARC name | 
 **originType** | **optional.String**| Type of origin point. | 
 **originId** | **optional.String**| ID of the origin point | 
 **originDevice** | **optional.String**| if origin type is device, this is name of the device. Absent otherwise | 
 **originCircuitId** | **optional.String**| if origin type is circuit, this is circuit_id of the circuit. Absent otherwise | 
 **originSwitchport** | **optional.String**| if origin type is switchport, this is name of the switch port. Absent otherwise | 
 **originSwitch** | **optional.String**| if origin type is switchport, this is name of the switch. Absent otherwise | 
 **originPatchPanelPort** | **optional.String**| if origin type is patch_panel_port, this is name of the patch panel port. Absent otherwise | 
 **originPatchPanel** | **optional.String**| if origin type is patch_panel_port, this is name of the patch panel. Absent otherwise | 
 **originPatchPanelId** | **optional.String**| if origin type is patch_panel_port, this is ID of the patch panel. Absent otherwise | 
 **originVendor** | **optional.String**| if origin type is vendor, this is name of the vendor. Absent otherwise | 
 **endPointType** | **optional.String**| Type of end point. | 
 **endPointId** | **optional.String**| ID of the end point | 
 **endPointDevice** | **optional.String**| if end_point type is device, this is name of the device. Absent otherwise | 
 **endPointCircuitId** | **optional.String**| if end_point type is circuit, this is circuit_id of the circuit. Absent otherwise | 
 **endPointSwitchport** | **optional.String**| if end_point type is switchport, this is name of the switch port. Absent otherwise | 
 **endPointSwitch** | **optional.String**| if end_point type is switchport, this is name of the switch. Absent otherwise | 
 **endPointPatchPanelPort** | **optional.String**| if end_point type is patch_panel_port, this is name of the patch panel port. Absent otherwise | 
 **endPointPatchPanel** | **optional.String**| if end_point type is patch_panel_port, this is name of the patch panel. Absent otherwise | 
 **endPointPatchPanelId** | **optional.String**| if end_point type is patch_panel_port, this is ID of the patch panel. Absent otherwise | 
 **endPointVendor** | **optional.String**| if end_point type is vendor, this is name of the vendor. Absent otherwise | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdateCircuits**
> interface{} PostUpdateCircuits(ctx, optional)
Create / Update Circuits

Create / Update Circuits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TelcoCircuitsApiPostUpdateCircuitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelcoCircuitsApiPostUpdateCircuitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iD** | **optional.String**| ID for the specific circuit (D42 ID) | 
 **circuitId** | **optional.String**| general circuit ID | 
 **type_** | **optional.String**| Required for new circuits | 
 **orderDate** | **optional.String**|  | 
 **provisionDate** | **optional.String**|  | 
 **turnOnDate** | **optional.String**|  | 
 **notes** | **optional.String**| Any additional notes | 
 **bandwidth** | **optional.String**| bandwidth in kbps (multiply factor of 1024) | 
 **vendor** | **optional.String**| The cloud vendor | 
 **customer** | **optional.String**| customer name | 
 **dmarcId** | **optional.String**| ID of the DMARC | 
 **dmarc** | **optional.String**| name of the DMARC | 
 **originType** | **optional.String**| Type of origin point. | 
 **originId** | **optional.String**| ID of the origin point | 
 **originDevice** | **optional.String**| if origin type is device, this is name of the device. Absent otherwise | 
 **originCircuitId** | **optional.String**| if origin type is circuit, this is circuit_id of the circuit. Absent otherwise | 
 **originSwitchport** | **optional.String**| if origin type is switchport, this is name of the switch port. Absent otherwise | 
 **originSwitch** | **optional.String**| if origin type is switchport, this is name of the switch. Absent otherwise | 
 **originPatchPanelPort** | **optional.String**| if origin type is patch_panel_port, this is name of the patch panel port. Absent otherwise | 
 **originPatchPanel** | **optional.String**| if origin type is patch_panel_port, this is name of the patch panel. Absent otherwise | 
 **originPatchPanelId** | **optional.String**| if origin type is patch_panel_port, this is ID of the patch panel. Absent otherwise | 
 **originVendor** | **optional.String**| if origin type is vendor, this is name of the vendor. Absent otherwise | 
 **endPointType** | **optional.String**| Type of end point. | 
 **endPointId** | **optional.String**| ID of the end point | 
 **endPointDevice** | **optional.String**| if end_point type is device, this is name of the device. Absent otherwise | 
 **endPointCircuitId** | **optional.String**| if end_point type is circuit, this is circuit_id of the circuit. Absent otherwise | 
 **endPointSwitchport** | **optional.String**| if end_point type is switchport, this is name of the switch port. Absent otherwise | 
 **endPointSwitch** | **optional.String**| if end_point type is switchport, this is name of the switch. Absent otherwise | 
 **endPointPatchPanelPort** | **optional.String**| if end_point type is patch_panel_port, this is name of the patch panel port. Absent otherwise | 
 **endPointPatchPanel** | **optional.String**| if end_point type is patch_panel_port, this is name of the patch panel. Absent otherwise | 
 **endPointPatchPanelId** | **optional.String**| if end_point type is patch_panel_port, this is ID of the patch panel. Absent otherwise | 
 **endPointVendor** | **optional.String**| if end_point type is vendor, this is name of the vendor. Absent otherwise | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomFields**
> interface{} PutCustomFields(ctx, key, optional)
Create/update custom fields for circuits. (Added in v5.9.3)

Custom Fields for circuits. Either ID or circuit_id is required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***TelcoCircuitsApiPutCustomFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelcoCircuitsApiPutCustomFieldsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iD** | **optional.String**| ID for the cricuit (D42 ID) | 
 **circuitId** | **optional.String**| circuit_id for the circuit | 
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

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

