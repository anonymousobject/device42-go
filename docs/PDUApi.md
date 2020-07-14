# \PDUApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePdus**](PDUApi.md#DeletePdus) | **Delete** /api/1.0/pdus/{ID}/ | Delete PDU
[**DeletePdusRack**](PDUApi.md#DeletePdusRack) | **Delete** /api/1.0/pdus/rack/{id}/ | Delete PDU from a rack
[**GetPduModels**](PDUApi.md#GetPduModels) | **Get** /api/1.0/pdu_models/ | GET method retrieves all PDU Models.
[**GetPdus**](PDUApi.md#GetPdus) | **Get** /api/1.0/pdus/ | GET method retrieves all PDUs.
[**GetPdusID**](PDUApi.md#GetPdusID) | **Get** /api/1.0/pdus/{ID}/ | Get PDU by ID
[**PostPduModels**](PDUApi.md#PostPduModels) | **Post** /api/1.0/pdu_models/ | Create / Update PDU Models.
[**PostPduModelsPorts**](PDUApi.md#PostPduModelsPorts) | **Post** /api/1.0/pdu_models/ports/ | Create PDU Model Ports.
[**PostPdus**](PDUApi.md#PostPdus) | **Post** /api/1.0/pdus/ | This call will create PDUs.
[**PostPdusPorts**](PDUApi.md#PostPdusPorts) | **Post** /api/1.0/pdus/ports/ | Update PDU Ports w/ no names
[**PostPdusRack**](PDUApi.md#PostPdusRack) | **Post** /api/1.0/pdus/rack/ | This call will add / update PDUs in or around a Rack.
[**PutPdus**](PDUApi.md#PutPdus) | **Put** /api/1.0/pdus/ | Update PDUs
[**PutPdusPorts**](PDUApi.md#PutPdusPorts) | **Put** /api/1.0/pdus/ports/ | Update PDU Ports w/ names


# **DeletePdus**
> interface{} DeletePdus(ctx, iD)
Delete PDU

This API is used to delete the pdu with the pdu id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| PDU id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePdusRack**
> interface{} DeletePdusRack(ctx, id)
Delete PDU from a rack

This API is used to delete from its rack the device with the device id supplied as the required argument. (The device itself is not deleted).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| device id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPduModels**
> interface{} GetPduModels(ctx, optional)
GET method retrieves all PDU Models.

Get all PDU Models

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PDUApiGetPduModelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiGetPduModelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by model name | 
 **pduModelId** | **optional.Int32**| filter by pdu model id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdus**
> interface{} GetPdus(ctx, optional)
GET method retrieves all PDUs.

Get all PDUs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PDUApiGetPdusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiGetPdusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by name | 
 **type_** | **optional.String**| filter by type | 
 **pduModelId** | **optional.Int32**| filter by PDU model ID | 
 **pduModel** | **optional.String**| filter by PDU model name | 
 **buildingId** | **optional.Int32**| filter by building id | 
 **roomId** | **optional.Int32**| filter by room id | 
 **rackId** | **optional.Int32**| filter by rack id | 
 **serialNo** | **optional.String**| filter by PDU serial_no | 
 **deviceId** | **optional.Int32**| filter by device id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdusID**
> PduById GetPdusID(ctx, iD)
Get PDU by ID

Retrieve detailed information about a specific PDU by PDU ID. This also includes end point connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| The ID of the PDU to retrieve | 

### Return type

[**PduById**](Pdu_by_Id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPduModels**
> interface{} PostPduModels(ctx, optional)
Create / Update PDU Models.

Create / Update PDU Models

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PDUApiPostPduModelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiPostPduModelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pduModelId** | **optional.String**| ID of the PDU model you want to update | 
 **name** | **optional.String**| Name of the PDU model you want to create or update | 
 **manufacturer** | **optional.String**| name of the hardware/software manufacturer. | 
 **size** | **optional.String**| Size of the PDU in U | 
 **sequentialNumberingForPorts** | **optional.String**| Could be “yes” or “no”. Yes if ports are numbered starting from 1. | 
 **depth** | **optional.String**| Half depth by default. full to override. Optional. | 
 **imgfileId** | **optional.String**| image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now. | 
 **imgfile** | **optional.String**| name of the image file (Added in v5.8.2). Use instead of imgfile_id | 
 **backImageId** | **optional.String**| back image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now. | 
 **backImage** | **optional.String**| name of the back image file. Use instead of back_image_id. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPduModelsPorts**
> interface{} PostPduModelsPorts(ctx, count, type_, optional)
Create PDU Model Ports.

Create PDU Model Ports. Required parameters: <ul><li>pdu_model_id <b>OR</b> pdu_model_name</li> <li>count</li> <li>type</li></ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **count** | **int32**| number of ports | 
  **type_** | **string**| Type of the port. If not already existing, a new port type is created. | 
 **optional** | ***PDUApiPostPduModelsPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiPostPduModelsPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pduModelId** | **optional.String**| ID of the PDU model you want to update | 
 **pduModelName** | **optional.String**| Name of the PDU model you want to add ports to. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPdus**
> interface{} PostPdus(ctx, name, optional)
This call will create PDUs.

Create PDUs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| PDU name | 
 **optional** | ***PDUApiPostPdusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiPostPdusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pduModelId** | **optional.String**| ID of the PDU model you want to update | 
 **pduModel** | **optional.String**| Name of the PDU model. You can use this instead of the ID above. (Added in v5.8.2) | 
 **device** | **optional.String**| If you want to associate asset information with this PDU, use device type &#39;other&#39; | 
 **notes** | **optional.String**| Any additional notes | 
 **category** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no, Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. Use for initial insert. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPdusPorts**
> interface{} PostPdusPorts(ctx, portType, optional)
Update PDU Ports w/ no names

Assign a name and/or an object (see below) objects to a pdu port. It will pick the lowest port id # available (or first available port in order created).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portType** | **string**| Verbose name of the port type. Must exist already. | 
 **optional** | ***PDUApiPostPdusPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiPostPdusPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentPduId** | **optional.String**| Available from /api/api/1.0/pdus/ or Tools | 
 **parentPdu** | **optional.String**| name of the parent PDU. Must be unique name. Added in v5.8.2 | 
 **device** | **optional.String**| Name of the device the port points to. | 
 **deviceId** | **optional.String**| ID of the device the port points to | 
 **pduId** | **optional.String**| ID of the PDU the port points to | 
 **assetId** | **optional.String**| ID of the asset the port points to | 
 **name** | **optional.String**| PDU port name, typically the PDU port number when autodiscovered. | 
 **outletName** | **optional.String**| outlet name | 
 **watts** | **optional.String**| per power supply | 
 **psuLabel** | **optional.String**| typically used when device has multiple power supplies, e.g.: power supply 1, power supply 2, etc. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPdusRack**
> interface{} PostPdusRack(ctx, optional)
This call will add / update PDUs in or around a Rack.

Add / Update PDUs in Racks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PDUApiPostPdusRackOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiPostPdusRackOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pduId** | **optional.String**| ID of the PDU to be edited | 
 **pdu** | **optional.String**| name of the PDU. only works if the name is unique in the system | 
 **rackId** | **optional.String**| This is the ID of the rack. It can be obtained from Tools &gt; Import &gt; Import Racked Devices or /api/api/1.0/racks/ | 
 **building** | **optional.String**|  | 
 **where** | **optional.String**| Location in a rack. Note: If mounted a size must be provided or available from the hardware model. | 
 **startAt** | **optional.String**| Required if adding to rack. U Start location. | 
 **orientation** | **optional.String**| orientation of the PDU in rack. back for rear facing, otherwise front is default. | 
 **xPos** | **optional.Int32**| A number between 0 and 2520 representing the position within the u slot in increments of 252, which is equal to 1/10th of the width of the rack. 0 will place a device flush left, 1260 will place the left side of a device in center. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdus**
> interface{} PutPdus(ctx, name, optional)
Update PDUs

This call will update existing PDUs. PDU ID or name required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Use instead of PDU ID. Must be unique | 
 **optional** | ***PDUApiPutPdusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiPutPdusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pduId** | **optional.String**| ID of the PDU to be edited | 
 **newName** | **optional.String**| Use to change name of object. | 
 **pduModelId** | **optional.String**| ID of the PDU model you want to update | 
 **pduModel** | **optional.String**| Name of the PDU model. You can use this instead of the ID above. (Added in v5.8.2) | 
 **type_** | **optional.String**| Type of power unit (pdu, ups, ats) | 
 **device** | **optional.String**| If you want to associate asset information with this PDU, use device type &#39;other&#39; | 
 **portNumber** | **optional.String**| port number | 
 **outletName** | **optional.String**| outlet name | 
 **notes** | **optional.String**| Any additional notes | 
 **ratedPower** | **optional.String**|  | 
 **objectCategory** | **optional.String**|  | 
 **storageRoomId** | **optional.String**| ID of storage room to assign power unit to | 
 **storageRoom** | **optional.String**| Name of storage room to apply power unit to | 
 **category** | **optional.String**| If multitenancy is on, admin groups that have access to this category, will have access to the power unit. If this parameter is present with no value, all categories are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdusPorts**
> interface{} PutPdusPorts(ctx, name, optional)
Update PDU Ports w/ names

This call requires the name of an existing pdu port and enables you to add new or edit existing values for that particular PDU port. Requires parent_pdu_id or parent_pdu

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the existing PDU port | 
 **optional** | ***PDUApiPutPdusPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PDUApiPutPdusPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentPduId** | **optional.String**| Available from /api/api/1.0/pdus/ or Tools | 
 **parentPdu** | **optional.String**| name of the parent PDU. Must be unique name. Added in v5.8.2 | 
 **outletName** | **optional.String**| outlet name | 
 **device** | **optional.String**| Name of the device the port points to. | 
 **deviceId** | **optional.String**| ID of the device the port points to | 
 **pduId** | **optional.String**| ID of the PDU the port points to | 
 **assetId** | **optional.String**| ID of the asset the port points to | 
 **watts** | **optional.String**| per power supply | 
 **psuLabel** | **optional.String**| typically used when device has multiple power supplies, e.g.: power supply 1, power supply 2, etc. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

