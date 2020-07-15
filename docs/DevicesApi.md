# \DevicesApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeviceMountpoints**](DevicesApi.md#DeleteDeviceMountpoints) | **Delete** /api/1.0/device/mountpoints/{id}/ | Delete device mountpoints
[**DeleteDeviceRackDeviceId**](DevicesApi.md#DeleteDeviceRackDeviceId) | **Delete** /api/1.0/device/rack/{device_id}/ | Remove/Delete a device from a rack
[**DeleteDeviceUrl**](DevicesApi.md#DeleteDeviceUrl) | **Delete** /api/1.0/device/url/{id}/ | Delete a device url
[**DeleteDevicesAttachment**](DevicesApi.md#DeleteDevicesAttachment) | **Delete** /api/1.0/devices/id/{device-id}/attachments/{id}/ | Delete a device attachment by Device Id and Attachment Id.
[**DeleteDevicesId**](DevicesApi.md#DeleteDevicesId) | **Delete** /api/1.0/devices/{id}/ | Delete Device
[**GetAllDevicesAttachments**](DevicesApi.md#GetAllDevicesAttachments) | **Get** /api/1.0/devices/id/{device-id}/attachments/ | Get all attachments for a Device by Device Id.
[**GetDeviceMountpoints**](DevicesApi.md#GetDeviceMountpoints) | **Get** /api/1.0/device/mountpoints/ | Get device mountpoints - introduced in version 10.5.0
[**GetDeviceUrl**](DevicesApi.md#GetDeviceUrl) | **Get** /api/1.0/device/url/ | Get device URLs - introduced in version 7.0.0
[**GetDevices**](DevicesApi.md#GetDevices) | **Get** /api/1.0/devices/ | Retrieve basic information about all devices.
[**GetDevicesAll**](DevicesApi.md#GetDevicesAll) | **Get** /api/1.0/devices/all/ | Retrieve detailed information about all devices.
[**GetDevicesAsset**](DevicesApi.md#GetDevicesAsset) | **Get** /api/1.0/devices/asset/{device-asset}/ | Get Device by Device Asset Number
[**GetDevicesCustomerId**](DevicesApi.md#GetDevicesCustomerId) | **Get** /api/1.0/devices/customer_id/{customer-id}/ | Get Devices by Customer Id
[**GetDevicesId**](DevicesApi.md#GetDevicesId) | **Get** /api/1.0/devices/id/{device-id}/ | Get Device by Device Id
[**GetDevicesImpactlist**](DevicesApi.md#GetDevicesImpactlist) | **Get** /api/1.0/device/impactlist/{device-id}/ | This API endpoint retrieves the impact list of a device by ID.
[**GetDevicesName**](DevicesApi.md#GetDevicesName) | **Get** /api/1.0/devices/name/{device-name}/ | Get Device by Device Name
[**GetDevicesSerial**](DevicesApi.md#GetDevicesSerial) | **Get** /api/1.0/devices/serial/{device-serial}/ | Get Device by Device Serial Number
[**GetaDevicesAttachment**](DevicesApi.md#GetaDevicesAttachment) | **Get** /api/1.0/devices/id/{device-id}/attachments/{attachment_id}/ | Get a specific device attachment by Device Id and Attachment Id.
[**PostCloudInstanceDevice**](DevicesApi.md#PostCloudInstanceDevice) | **Post** /api/1.0/device/cloud_instance/ | Create/Update Device Cloud Instance Information
[**PostDevice**](DevicesApi.md#PostDevice) | **Post** /api/1.0/device/ | Create/Update Device by Name
[**PostDeviceMountpoints**](DevicesApi.md#PostDeviceMountpoints) | **Post** /api/1.0/device/mountpoints/ | Create/Update device mountpoints - introduced in version 10.5.0
[**PostDeviceRack**](DevicesApi.md#PostDeviceRack) | **Post** /api/1.0/device/rack/ | Add/Update a device in a rack.
[**PostDeviceUrl**](DevicesApi.md#PostDeviceUrl) | **Post** /api/1.0/device/url/ | Use this API to associate a url (e.g. an http, https, or telnet url) with a device.
[**PostMultiNodeDevice**](DevicesApi.md#PostMultiNodeDevice) | **Post** /api/1.0/multinodes/ | Create/Update Multi-Node Device by Name
[**PostMultiSerialDevice**](DevicesApi.md#PostMultiSerialDevice) | **Post** /api/1.0/multiserial_device/ | Create/Update Multi-Serial Device by Name
[**PutCustomField**](DevicesApi.md#PutCustomField) | **Put** /api/1.0/device/custom_field/ | Create/Update Custom Fields for a Device
[**PutDevice**](DevicesApi.md#PutDevice) | **Put** /api/1.0/device/ | Update an existing device by name, serial, ID or asset number.
[**PutDeviceUrl**](DevicesApi.md#PutDeviceUrl) | **Put** /api/1.0/device/url/ | Update Device URL
[**UploadDevicesAttachments**](DevicesApi.md#UploadDevicesAttachments) | **Post** /api/1.0/devices/id/{device-id}/attachments/ | Upload and attach a file to a Device by Device Id.


# **DeleteDeviceMountpoints**
> interface{} DeleteDeviceMountpoints(ctx, id)
Delete device mountpoints

Delete Device Mountpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| id of the mountpoint to delete | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeviceRackDeviceId**
> interface{} DeleteDeviceRackDeviceId(ctx, deviceId)
Remove/Delete a device from a rack

This API call will remove/delete a device from a rack (Introduced in v6.3.2)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of the device to be removed from the rack | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeviceUrl**
> interface{} DeleteDeviceUrl(ctx, id)
Delete a device url

This API call will delete a device url (Introduced in v7.0.0)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDevicesAttachment**
> interface{} DeleteDevicesAttachment(ctx, deviceId, id)
Delete a device attachment by Device Id and Attachment Id.

Delete a file attachment for a specific device using device Id and the attachment Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **id** | **int32**| ID of the attachment to delete. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDevicesId**
> interface{} DeleteDevicesId(ctx, id)
Delete Device

This API is used to delete a device with the device id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Device id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllDevicesAttachments**
> interface{} GetAllDevicesAttachments(ctx, deviceId)
Get all attachments for a Device by Device Id.

Retrieve all the file attachments for a specific device using device id.<br> Returns a zip file of all the attachments associated with the device with the specified ID in URL.<br> Note that this implementation of Swagger does not support file download - use another mechanism such as cULR.<br> <br>cURL example:<br> curl -X GET \\<br> URL/api/1.0/devices/id/3/attachments \\<br> -H 'Authorization: Basic YWRtaW46YWRtIW5kNDI='

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceMountpoints**
> interface{} GetDeviceMountpoints(ctx, optional)
Get device mountpoints - introduced in version 10.5.0

Get Device Mountpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiGetDeviceMountpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDeviceMountpointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.String**| id of the device | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceUrl**
> interface{} GetDeviceUrl(ctx, optional)
Get device URLs - introduced in version 7.0.0

Get Device URLs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiGetDeviceUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDeviceUrlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device** | **optional.String**| name of the device | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevices**
> interface{} GetDevices(ctx, optional)
Retrieve basic information about all devices.

Get All Devices With Brief Output

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiGetDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDevicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| filter by device type (physical, virtual, blade, other, cluster, or unknown) | 
 **deviceSubType** | **optional.String**| filter by device sub type (Added in v14.7.2) | 
 **deviceSubTypeId** | **optional.String**| filter by device sub type id (Added in v14.7.2) | 
 **serviceLevel** | **optional.String**| filter by service level name | 
 **inService** | **optional.String**| filter by whether in service or not. Use yes or no. | 
 **customer** | **optional.String**| filter by customer name | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 
 **bladeHostName** | **optional.String**| filter by blade host name | 
 **virtualHostName** | **optional.String**| filter by virtual host name | 
 **buildingId** | **optional.String**| filter by building ID (Added in v5.9.0) | 
 **building** | **optional.String**| filter by building name | 
 **roomId** | **optional.String**| filter by room ID (Added in v5.9.0) | 
 **room** | **optional.String**| filter by room name. Only works if room ID is not present (Added in v5.9.0) | 
 **rackId** | **optional.String**| filter by rack ID (Added in v5.9.0) | 
 **rack** | **optional.String**| filter by rack name. Only works if rack ID is not present (Added in v5.9.0) | 
 **serialNo** | **optional.String**| filter by serial # (Added in v6.0.0) | 
 **serialNoContains** | **optional.String**| filter by partial serial match (Added in 9.7.1) | 
 **objectCategory** | **optional.String**| filter by object category | 
 **objectCategoryId** | **optional.String**| filter by object category ID | 
 **assetNo** | **optional.String**| filter by asset # (Added in v6.0.0) | 
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **tagsAnd** | **optional.String**| filter by all the tags, separated by comma. (This is an AND filter and all tags have to match for filter, this was added in v6.3.1) | 
 **uuid** | **optional.String**| filter by uuid (exact match) (Added in v6.3.2) | 
 **isItSwitch** | **optional.String**| filter by whether switch or not. Use yes or no. (Added in v6.3.2) | 
 **isItVirtualHost** | **optional.String**| filter by whether virtual host or not. Use yes or no. (Added in v6.3.2) | 
 **isItBladeHost** | **optional.String**| filter by whether blade host or not. Use yes or no. (Added in v6.3.2) | 
 **hardware** | **optional.String**| filter by name of hardware model, comma separated for multiple hardware models (or filter). (Added in v6.3.2) | 
 **hardwareIds** | **optional.String**| filter by ID of hardware models, comma separated | 
 **os** | **optional.String**| filter by OS name (added in v8.3.0) | 
 **virtualSubtype** | **optional.String**| filter by virtual subtype (added in v8.3.2) | 
 **lastUpdatedLt** | **optional.String**| last updated less than date YYYY-MM-DD format | 
 **lastUpdatedGt** | **optional.String**| last updated greater than date YYYY-MM-DD format | 
 **firstAddedLt** | **optional.String**| first added less than date YYYY-MM-DD format | 
 **firstAddedGt** | **optional.String**| first added greater date YYYY-MM-DD format | 
 **customFieldsAnd** | **optional.String**| filter by custom fields, and filter, format of key1:value1,key2:value2 | 
 **customFieldsOr** | **optional.String**| filter by custom fields, or filter, format of key1:value1,key2:value2 | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicesAll**
> DevicesAll GetDevicesAll(ctx, optional)
Retrieve detailed information about all devices.

Get All Devices With Detailed Output (added in v6.3.4)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiGetDevicesAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDevicesAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeCols** | **optional.String**| do not return all columns just the ones specified. For example, ?include_cols&#x3D;name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager | 
 **limit** | **optional.String**| return this number of devices | 
 **offset** | **optional.String**| start with this device (e.g. limit&#x3D;100&amp;offset&#x3D;50 means start with the 50th device and return the next 100 devices) | 
 **blankasnull** | **optional.String**| start with this device (e.g. limit&#x3D;100&amp;offset&#x3D;50 means start with the 50th device and return the next 100 devices) | 

### Return type

[**DevicesAll**](devicesAll.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicesAsset**
> DevicesCustomerId GetDevicesAsset(ctx, deviceAsset, optional)
Get Device by Device Asset Number

Retrieve detailed information about a specific device by device asset number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceAsset** | [**string**](.md)|  | 
 **optional** | ***DevicesApiGetDevicesAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDevicesAssetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCols** | **optional.String**| do not return all columns just the ones specified. For example, ?include_cols&#x3D;name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager | 

### Return type

[**DevicesCustomerId**](devicesCustomerId.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicesCustomerId**
> DevicesCustomerId GetDevicesCustomerId(ctx, customerId, optional)
Get Devices by Customer Id

Retrieve all devices associated with a specific customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int32**|  | 
 **optional** | ***DevicesApiGetDevicesCustomerIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDevicesCustomerIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCols** | **optional.String**| do not return all columns just the ones specified. For example, ?include_cols&#x3D;name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager | 

### Return type

[**DevicesCustomerId**](devicesCustomerId.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicesId**
> DevicesAll GetDevicesId(ctx, deviceId, optional)
Get Device by Device Id

Retrieve detailed information about a specific device using device id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***DevicesApiGetDevicesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDevicesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **follow** | **optional.String**| use yes if you want to see virtuals in a virtual host, modules in a blade chassis and devices in a clustered device (Added in v5.7.4) | 

### Return type

[**DevicesAll**](devicesAll.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicesImpactlist**
> interface{} GetDevicesImpactlist(ctx, deviceId)
This API endpoint retrieves the impact list of a device by ID.

Get Device Impact List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicesName**
> DevicesCustomerId GetDevicesName(ctx, deviceName, optional)
Get Device by Device Name

Retrieve detailed information about a specific device by device name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceName** | [**string**](.md)|  | 
 **optional** | ***DevicesApiGetDevicesNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDevicesNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCols** | **optional.String**| do not return all columns just the ones specified. For example, ?include_cols&#x3D;name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager | 
 **follow** | **optional.String**| use yes if you want to see virtuals in a virtual host, modules in a blade chassis and devices in a clusetered device | 

### Return type

[**DevicesCustomerId**](devicesCustomerId.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicesSerial**
> DevicesCustomerId GetDevicesSerial(ctx, deviceSerial, optional)
Get Device by Device Serial Number

Retrieve detailed information about a specific device by device serial number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceSerial** | [**string**](.md)|  | 
 **optional** | ***DevicesApiGetDevicesSerialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiGetDevicesSerialOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCols** | **optional.String**| do not return all columns just the ones specified. For example, ?include_cols&#x3D;name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager | 

### Return type

[**DevicesCustomerId**](devicesCustomerId.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetaDevicesAttachment**
> interface{} GetaDevicesAttachment(ctx, deviceId, attachmentId)
Get a specific device attachment by Device Id and Attachment Id.

Retrieves a specific file attachment for a specific device using device id and attachment id.<br> Note that this implementation of Swagger does not support file download - use another mechanism such as cULR.<br> <br>cURL example:<br> curl -X GET \\<br> URL/api/1.0/devices/id/3/attachments/17 \\<br> -H 'Authorization: Basic YWRtaW46YWRtIW5kNDI=' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **attachmentId** | **int32**| ID of the attachment to get. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCloudInstanceDevice**
> interface{} PostCloudInstanceDevice(ctx, optional)
Create/Update Device Cloud Instance Information

Update cloud instance information for devices. Requires either device or device_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiPostCloudInstanceDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPostCloudInstanceDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device** | **optional.String**| Device name | 
 **deviceId** | **optional.Int32**| Device ID | 
 **instanceId** | **optional.Int32**| The cloud instance ID of the device | 
 **vendor** | **optional.String**| The cloud vendor | 
 **status** | **optional.String**| Instance status (ie, running, stopped) | 
 **location** | **optional.String**| Location/region of instance deployment | 
 **notes** | **optional.String**| Any additional notes | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDevice**
> interface{} PostDevice(ctx, optional)
Create/Update Device by Name

This API is used to create a new device with the name supplied as the required argument or update an existing device that has the name, serial # or uuid of the required argument. For information purposes, this is also the API that is used by the Device42 auto-discovery tool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiPostDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPostDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Device name. If device with name already exists - the existing device is updated. If an existing device based on serial # or uuid are found - name is ignored. In that case existing name is not changed and you can use new_name if you wish to change the name. | 
 **newName** | **optional.String**| As mentioned above, this can be used to change the name of an existing device | 
 **serialNo** | **optional.String**| the serial # of the device. If a device with matching serial # is found - that device is updated. Serial # with less than 3 characters in length are ignored. Also, certain generic serial #s like 123456789 or ‘not specified’ are ignored. | 
 **uuid** | **optional.String**| The uuid of the device. If a device with a matching uuid is found - that device is updated. | 
 **assetNo** | **optional.String**| the asset # of the device. | 
 **manufacturer** | **optional.String**| the hardware manufacturer for the device. | 
 **hardware** | **optional.String**| the name of the hardware model for the device. Use in conjunction with the manufacturer argument. However, an update will only occur if no hardware model is previously assigned to the device. | 
 **newHardware** | **optional.String**| If you want to change the hardware model for a device. If the device is rack mounted, it will unmount the device before changing hardware model and attempt to mount it back to same rack location (Changed in v6.3.2) | 
 **isItSwitch** | **optional.String**| ‘yes’ indicates if device is a network switch | 
 **isItVirtualHost** | **optional.String**| ‘yes’ indicates if device is a virtual host | 
 **isItBladeHost** | **optional.String**| ‘yes’ indicates if device is a blade host | 
 **inService** | **optional.String**| ‘yes’ indicates if device is in service, ‘no’ indicates not in service | 
 **type_** | **optional.String**| is the type for device. Valid values are ‘unknown’, ‘physical’, ‘virtual’, ‘blade’, ‘cluster’, or ‘other’. | 
 **serviceLevel** | **optional.String**| the service level name for the device. The service level must be pre-defined in the device42 appliance before it can be referenced in an api call. | 
 **virtualHost** | **optional.String**| Is the name for the host of the virtual machine. Two conditions must be met for this value to successfully update: 1. The device must be a virtual machine. 2. The virtual host must already exist in the device42 application and must be already marked as a virtual host. | 
 **bladeHost** | **optional.String**| the name of the host for the blade machine. Two conditions must be met for this value to successfully update: 1. The device must be a blade type. 2. The blade host must already exist and must already be marked as a blade host. | 
 **slotNo** | **optional.Int32**| slot # for blade device. | 
 **storageRoomId** | **optional.Int32**| ID of the room to assign device to storage room. Added in v5.5.0 | 
 **storageRoom** | **optional.String**| name of the room to assign device to, only used if the room name is unique. Added in v5.5.0 | 
 **os** | **optional.String**| the name of the operating system (os, osver, and osverno all required if updating any of the three). | 
 **osver** | **optional.String**| the version of the operating system (os, osver, and osverno all required if updating any of the three) | 
 **osarch** | **optional.String**| The architecture of the operating system (32 or 64) | 
 **osverno** | **optional.Int32**| the version # or build # of the operating system (os, osver, and osverno all required if updating any of the three) | 
 **memory** | **optional.String**| the total memory(RAM) in MB. Just numbers. | 
 **cpucount** | **optional.Int32**| total # of CPUs. (If updating, cpupower and cpucore required) | 
 **cpupower** | **optional.String**| CPU speed in MHz, just numbers. (If updating, cpucount and cpucore required) | 
 **cpucore** | **optional.String**| number of cores/CPU. Integer. (If updating, cpupower and cpucount required) | 
 **hddcount** | **optional.Int32**| total # of HDD (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddsize** | **optional.String**| HDD Size in GB (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddraid** | **optional.String**| Raid. Possible values: software or hardware. none to clear. (none added in v9.1.0) (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddraidType** | **optional.String**| Raid Type. (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **macaddress** | **optional.String**| the mac address value. Use multiple POST/PUT to add multiple mac addresses. | 
 **devicesInCluster** | **optional.String**| comma separated device names for devices in cluster. Only valid for device type cluster. | 
 **appcomps** | **optional.String**| comma separated application component names on this device. | 
 **customer** | **optional.String**| name of the customer. Customer record must already exist. | 
 **contractId** | **optional.Int32**| ID for the contract. Available via GET /api/api/1.0/contracts/ – DEPRECATED in v550. | 
 **contract** | **optional.String**| Name of the contract. Used only if contract name is unique. – DEPRECATED in v550. | 
 **aliases** | **optional.String**| optional. Comma separated aliases for the device. Must not exist, ignored otherwise. | 
 **subtype** | **optional.String**| Only for device type other. Must exist internally first. | 
 **bladeHostClear** | **optional.String**| clear blade host for device | 
 **virtualSubtype** | **optional.String**| Only for device type virtual. Default is internal VM. | [default to internal VM]
 **notes** | **optional.String**|  | 
 **tags** | **optional.String**| comma separated tags (Added in v5.9.2) | 
 **virtualHostClear** | **optional.String**| yes to clear virtual host for a VM (added in v9.1.0) | 
 **tagsRemove** | **optional.String**| comma separated tags to remove (added in v9.1.0) | 
 **aliasesRemove** | **optional.String**| comma separated aliases to remove | 
 **devicesInClusterRemove** | **optional.String**| comma separated device names for removing devices in cluster. Only valid for device type cluster. (added in v9.2.0) | 
 **objectCategory** | **optional.String**| If multitenancy is on, a category can be assigned to control access to this object, e.g. Prod_East specifies that the users with access to the Prod_East category will be able to access this device. | 
 **newObjectCategory** | **optional.String**| Used to change category on a device if it is already set. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeviceMountpoints**
> interface{} PostDeviceMountpoints(ctx, mountpoint, device, optional)
Create/Update device mountpoints - introduced in version 10.5.0

Create/Update Device Mountpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mountpoint** | **string**| path of the mountpoint | 
  **device** | **string**| device mountpoint is assigned to | 
 **optional** | ***DevicesApiPostDeviceMountpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPostDeviceMountpointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filesytem** | **optional.String**| mountpoint filesystem | 
 **fstype** | **optional.String**| mountpoint filesystem type | 
 **capacity** | **optional.String**| capacity of mountpoint in MB | 
 **freeCapacity** | **optional.String**| free capacity of mountpoint in MB | 
 **label** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeviceRack**
> interface{} PostDeviceRack(ctx, startAt, optional)
Add/Update a device in a rack.

This API will add a new or existing device to a rack or will update an existing device that is already in the rack. (Introduced in v321)<br> Required parameters: <ul><li>device <b>OR</b> device_id <b>OR</b> serial_no <b>OR</b> asset_no</li> <li>rack_id <b>OR</b> rack (with building, room, or row to identify unique rack) <li>start_at</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startAt** | **string**| This is the starting U location for the device in the rack. Starting with v535, you can use “auto” as value to automatically mount the device in next available slot. | 
 **optional** | ***DevicesApiPostDeviceRackOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPostDeviceRackOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **device** | **optional.String**| name of the new or existing device | 
 **deviceId** | **optional.Int32**| Device ID of existing device | 
 **serialNo** | **optional.String**| serial # of the existing device | 
 **assetNo** | **optional.String**| asset # of the existing device | 
 **hwModel** | **optional.String**| If the hw_model doesn’t exist or doesn’t have a type, we will add it as type “regular (rack mountable)” (changed in v6.6.0) | 
 **size** | **optional.String**| size of the hardware model, only for new hardware model or if hardware model doesn’t have size. required for new hardware model (added in v6.6.0) | 
 **manufacturer** | **optional.String**| manufacturer of the hardware model. Only for new hardware model being added(added in v6.6.0) | 
 **rackId** | **optional.String**| required if building name, room name or rack name are not provided. This is the id of the rack. It can be obtained from Tools &gt; Import &gt; Import Racked Devices. | 
 **building** | **optional.String**| building is building name, room is room name, rack is rack name and row is optional. This is used if rack_id is not provided and a unique rack is found with that combination. This could be just rack for rack name, if the rack name is unique. Otherwise add row, room or building to identify a unique rack. | 
 **room** | **optional.String**|  | 
 **rack** | **optional.String**|  | 
 **row** | **optional.String**|  | 
 **where** | **optional.String**| location in rack, one of ‘above’, ‘below’, ‘left’, ‘right’, ‘mounted’. Note: If mounted a size must be provided or available from the hardware model. | 
 **xPos** | **optional.String**| A number between 0 and 2520 representing the position within the u slot. 0 is flush left. 2520 is flush right. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeviceUrl**
> interface{} PostDeviceUrl(ctx, type_, device, optional)
Use this API to associate a url (e.g. an http, https, or telnet url) with a device.

Add Url to Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| the url type (e.g. http, https, telnet). This must be an existing url type in device42. | 
  **device** | **string**| name of the device to which this url belongs (required parameter) | 
 **optional** | ***DevicesApiPostDeviceUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPostDeviceUrlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **host** | **optional.String**| fqdn portion of the url. e.g. for http://www.device42.com/awesome, host is www.device42.com | 
 **port** | **optional.String**| port number if any | 
 **urlSuffix** | **optional.String**| url suffix if any. e.g. it is “awesome” based on example given above. | 
 **notes** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMultiNodeDevice**
> interface{} PostMultiNodeDevice(ctx, optional)
Create/Update Multi-Node Device by Name

This API is used to create a new device with the name supplied as the required argument or update an existing device that with that name. This is similar to /api/api/1.0/devices/ POST call - but you can add duplicate serial #s and UUID for high density servers that share the same serial # and/or UUID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiPostMultiNodeDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPostMultiNodeDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Device name. If device with name already exists - the existing device is updated. If an existing device based on serial # or uuid are found - name is ignored. In that case existing name is not changed and you can use new_name if you wish to change the name. | 
 **newName** | **optional.String**| As mentioned above, this can be used to change the name of an existing device | 
 **serialNo** | **optional.String**| the serial # of the device. If a device with matching serial # is found - that device is updated. Serial # with less than 3 characters in length are ignored. Also, certain generic serial #s like 123456789 or ‘not specified’ are ignored. | 
 **uuid** | **optional.Int32**| the uuid of the device. If a device with matching uuid is found - that device is updated. | 
 **assetNo** | **optional.String**| the asset # of the device. | 
 **manufacturer** | **optional.String**| the hardware manufacturer for the device. | 
 **hardware** | **optional.String**| the name of the hardware model for the device. Use in conjunction with the manufacturer argument. However, an update will only occur if no hardware model is previously assigned to the device. | 
 **newHardware** | **optional.String**| If you want to change the hardware model for a device. If the device is rack mounted, it will unmount the device before changing hardware model and attempt to mount it back to same rack location (Changed in v6.3.2) | 
 **isItSwitch** | **optional.String**| ‘yes’ indicates if device is a network switch | 
 **isItVirtualHost** | **optional.String**| ‘yes’ indicates if device is a virtual host | 
 **isItBladeHost** | **optional.String**| ‘yes’ indicates if device is a blade host | 
 **inService** | **optional.String**| ‘yes’ indicates if device is in service, ‘no’ indicates not in service | 
 **type_** | **optional.String**| the type of the device | 
 **serviceLevel** | **optional.String**| the service level name for the device. The service level must be pre-defined in the device42 appliance before it can be referenced in an api call. | 
 **virtualHost** | **optional.String**| Is the name for the host of the virtual machine. Two conditions must be met for this value to successfully update: 1. The device must be a virtual machine. 2. The virtual host must already exist in the device42 application and must be already marked as a virtual host. | 
 **bladeHost** | **optional.String**| the name of the host for the blade machine. Two conditions must be met for this value to successfully update: 1. The device must be a blade type. 2. The blade host must already exist and must already be marked as a blade host. | 
 **slotNo** | **optional.Int32**| slot # for blade device. | 
 **storageRoomId** | **optional.Int32**| ID of the room to assign device to storage room. Added in v5.5.0 | 
 **storageRoom** | **optional.String**| name of the room to assign device to, only used if the room name is unique. Added in v5.5.0 | 
 **os** | **optional.String**| the name of the operating system (os, osver, and osverno all required if updating any of the three). | 
 **osver** | **optional.String**| the version of the operating system (os, osver, and osverno all required if updating any of the three) | 
 **osarch** | **optional.String**| The architecture of the operating system (32 or 64) | 
 **osverno** | **optional.Int32**| the version # or build # of the operating system (os, osver, and osverno all required if updating any of the three) | 
 **memory** | **optional.String**| the total memory(RAM) in MB. Just numbers. | 
 **cpucount** | **optional.Int32**| total # of CPUs. (If updating, cpupower and cpucore required) | 
 **cpupower** | **optional.String**| CPU speed in MHz, just numbers. (If updating, cpucount and cpucore required) | 
 **cpucore** | **optional.String**| number of cores/CPU. Integer. (If updating, cpupower and cpucount required) | 
 **hddcount** | **optional.Int32**| total # of HDD (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddsize** | **optional.String**| HDD Size in GB (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddraid** | **optional.String**| Raid. Possible values: software or hardware. none to clear. (none added in v9.1.0) (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddraidType** | **optional.String**| Raid Type. (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **macaddress** | **optional.String**| the mac address value. Use multiple POST/PUT to add multiple mac addresses. | 
 **devicesInCluster** | **optional.String**| comma separated device names for devices in cluster. Only valid for device type cluster. | 
 **appcomps** | **optional.String**| comma separated application component names on this device. | 
 **customer** | **optional.String**| name of the customer. Customer record must already exist. | 
 **contractId** | **optional.Int32**| ID for the contract. Available via GET /api/api/1.0/contracts/ – DEPRECATED in v550. | 
 **contract** | **optional.String**| Name of the contract. Used only if contract name is unique. – DEPRECATED in v550. | 
 **aliases** | **optional.String**| optional. Comma separated aliases for the device. Must not exist, ignored otherwise. | 
 **subtype** | **optional.String**| Only for device type other. Must exist internally first. | 
 **bladeHostClear** | **optional.String**| clear blade host for device | 
 **virtualSubtype** | **optional.String**| Only for device type virtual. Default is internal VM. | [default to internal VM]
 **notes** | **optional.String**|  | 
 **tags** | **optional.String**| comma separated tags (Added in v5.9.2) | 
 **virtualHostClear** | **optional.String**| yes to clear virtual host for a VM (added in v9.1.0) | 
 **tagsRemove** | **optional.String**| comma separated tags to remove (added in v9.1.0) | 
 **aliasesRemove** | **optional.String**| comma separated aliases to remove | 
 **devicesInClusterRemove** | **optional.String**| comma separated device names for removing devices in cluster. Only valid for device type cluster. (added in v9.2.0) | 
 **objectCategory** | **optional.String**| If multitenancy is on, a category can be assigned to control access to this object, e.g. Prod_East specifies that the users with access to the Prod_East category will be able to access this device. | 
 **newObjectCategory** | **optional.String**| Used to change category on a device if it is already set. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMultiSerialDevice**
> interface{} PostMultiSerialDevice(ctx, optional)
Create/Update Multi-Serial Device by Name

This API is used to create a new device with the name supplied as the required argument or update an existing device that with that name. This is similar to /api/api/1.0/devices/ POST call - but you can add duplicate serial numberss and UUID for high density servers that share the same serial # and/or UUID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiPostMultiSerialDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPostMultiSerialDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Device name. If device with name already exists - the existing device is updated. If an existing device based on serial # or uuid are found - name is ignored. In that case existing name is not changed and you can use new_name if you wish to change the name. | 
 **newName** | **optional.String**| As mentioned above, this can be used to change the name of an existing device | 
 **serialNo** | **optional.String**| the serial # of the device. If a device with matching serial # is found - that device is updated. Serial # with less than 3 characters in length are ignored. Also, certain generic serial #s like 123456789 or ‘not specified’ are ignored. | 
 **uuid** | **optional.String**| the uuid of the device. If a device with matching uuid is found - that device is updated. | 
 **assetNo** | **optional.String**| the asset # of the device. | 
 **manufacturer** | **optional.String**| the hardware manufacturer for the device. | 
 **hardware** | **optional.String**| the name of the hardware model for the device. Use in conjunction with the manufacturer argument. However, an update will only occur if no hardware model is previously assigned to the device. | 
 **newHardware** | **optional.String**| If you want to change the hardware model for a device. If the device is rack mounted, it will unmount the device before changing hardware model and attempt to mount it back to same rack location (Changed in v6.3.2) | 
 **isItSwitch** | **optional.String**| ‘yes’ indicates if device is a network switch | 
 **isItVirtualHost** | **optional.String**| ‘yes’ indicates if device is a virtual host | 
 **isItBladeHost** | **optional.String**| ‘yes’ indicates if device is a blade host | 
 **inService** | **optional.String**| ‘yes’ indicates if device is in service, ‘no’ indicates not in service | 
 **type_** | **optional.String**| is the type for device. Valid values are ‘physical’, ‘virtual’, ‘blade’, ‘cluster’, or ‘other’. | 
 **serviceLevel** | **optional.String**| the service level name for the device. The service level must be pre-defined in the device42 appliance before it can be referenced in an api call. | 
 **virtualHost** | **optional.String**| Is the name for the host of the virtual machine. Two conditions must be met for this value to successfully update: 1. The device must be a virtual machine. 2. The virtual host must already exist in the device42 application and must be already marked as a virtual host. | 
 **bladeHost** | **optional.String**| the name of the host for the blade machine. Two conditions must be met for this value to successfully update: 1. The device must be a blade type. 2. The blade host must already exist and must already be marked as a blade host. | 
 **slotNo** | **optional.Int32**| slot # for blade device. | 
 **storageRoomId** | **optional.Int32**| ID of the room to assign device to storage room. Added in v5.5.0 | 
 **storageRoom** | **optional.String**| name of the room to assign device to, only used if the room name is unique. Added in v5.5.0 | 
 **os** | **optional.String**| the name of the operating system (os, osver, and osverno all required if updating any of the three). | 
 **osver** | **optional.String**| the version of the operating system (os, osver, and osverno all required if updating any of the three) | 
 **osarch** | **optional.String**| The architecture of the operating system (32 or 64) | 
 **osverno** | **optional.Int32**| the version # or build # of the operating system (os, osver, and osverno all required if updating any of the three) | 
 **memory** | **optional.String**| the total memory(RAM) in MB. Just numbers. | 
 **cpucount** | **optional.Int32**| total # of CPUs. (If updating, cpupower and cpucore required) | 
 **cpupower** | **optional.String**| CPU speed in MHz, just numbers. (If updating, cpucount and cpucore required) | 
 **cpucore** | **optional.String**| number of cores/CPU. Integer. (If updating, cpupower and cpucount required) | 
 **hddcount** | **optional.Int32**| total # of HDD (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddsize** | **optional.String**| HDD Size in GB (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddraid** | **optional.String**| Raid. Possible values: software or hardware. none to clear. (none added in v9.1.0) (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddraidType** | **optional.String**| Raid Type. (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **macaddress** | **optional.String**| the mac address value. Use multiple POST/PUT to add multiple mac addresses. | 
 **devicesInCluster** | **optional.String**| comma separated device names for devices in cluster. Only valid for device type cluster. | 
 **appcomps** | **optional.String**| comma separated application component names on this device. | 
 **customer** | **optional.String**| name of the customer. Customer record must already exist. | 
 **contractId** | **optional.Int32**| ID for the contract. Available via GET /api/api/1.0/contracts/ – DEPRECATED in v550. | 
 **contract** | **optional.String**| Name of the contract. Used only if contract name is unique. – DEPRECATED in v550. | 
 **aliases** | **optional.String**| optional. Comma separated aliases for the device. Must not exist, ignored otherwise. | 
 **subtype** | **optional.String**| Only for device type other. Must exist internally first. | 
 **bladeHostClear** | **optional.String**| clear blade host for device | 
 **virtualSubtype** | **optional.String**| Only for device type virtual. Default is internal VM. | [default to internal VM]
 **notes** | **optional.String**|  | 
 **tags** | **optional.String**| comma separated tags (Added in v5.9.2) | 
 **virtualHostClear** | **optional.String**| yes to clear virtual host for a VM (added in v9.1.0) | 
 **tagsRemove** | **optional.String**| comma separated tags to remove (added in v9.1.0) | 
 **aliasesRemove** | **optional.String**| comma separated aliases to remove | 
 **devicesInClusterRemove** | **optional.String**| comma separated device names for removing devices in cluster. Only valid for device type cluster. (added in v9.2.0) | 
 **objectCategory** | **optional.String**| If multitenancy is on, a category can be assigned to control access to this object, e.g. Prod_East specifies that the users with access to the Prod_East category will be able to access this device. | 
 **newObjectCategory** | **optional.String**| Used to change category on a device if it is already set. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCustomField**
> interface{} PutCustomField(ctx, key, optional)
Create/Update Custom Fields for a Device

One of the following parameters is required: name, id, asset, or serial

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***DevicesApiPutCustomFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPutCustomFieldOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the device, or | 
 **id** | **optional.Int32**| ID of the device | 
 **asset** | **optional.String**| if there is more than 1 device with the same asset #, this will return “device not found”, or | 
 **serial** | **optional.String**| if there is more than 1 device with the same serial #, this will return “device not found” | 
 **type_** | **optional.String**| this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD | 
 **relatedFieldName** | **optional.String**| Required if type &#x3D; related_field. The existing field to relate this custom field to. Can be: appcomp (for application components), asset, building, certificate, circuit, cusotmer, device, endusers, hardware (for device hardware models), ip_address, natip, netport (for ports), os, part, partmodel, password, pdu (for power units), pdu_model (for power unit models), ports, purchase, purchaselineitem (for a line item on a purchase order), rack, room, software, vlan (for subnets), switch_vlan (for vlans), vrfgroup | 
 **addToPicklist** | **optional.String**| Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored. | 
 **value** | **optional.String**| This will set the value of the custom field. | 
 **clearValue** | **optional.String**| yes to clear existing value for that field | 
 **notes** | **optional.String**| Any additional notes. | 
 **clearNotes** | **optional.String**| Yes to clear any existing notes. | 
 **bulkFields** | **optional.String**| comma separated key value pairs, with key and value separated by colon. e.g.key1:value1, key2:value2 (added in v6.4.1) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDevice**
> interface{} PutDevice(ctx, optional)
Update an existing device by name, serial, ID or asset number.

Update Device by Name, ID, Serial or Asset. Requires one of the following parameters: name, serial, asset, device_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiPutDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPutDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Updates the device found with that name. Otherwise returns “device not found\&quot; | 
 **serial** | **optional.String**| If the serial number is provided, it must be a unique serial number or a “device not found” error will result. If it is a unique serial number, then the device that has that serial number will be updated. | 
 **asset** | **optional.String**| If the asset number is provided, it must be a unique asset number or a “device not found” error will result. If it is a unique asset number, then the device that has that asset number will be updated. | 
 **deviceId** | **optional.String**| update the device by Device42 ID | 
 **newName** | **optional.String**| As mentioned above, this can be used to change the name of an existing device | 
 **serialNo** | **optional.String**| the serial # of the device. If a device with matching serial # is found - that device is updated. Serial # with less than 3 characters in length are ignored. Also, certain generic serial #s like 123456789 or ‘not specified’ are ignored. | 
 **uuid** | **optional.String**| the uuid of the device. If a device with matching uuid is found - that device is updated. | 
 **assetNo** | **optional.String**| the asset # of the device. | 
 **manufacturer** | **optional.String**| the hardware manufacturer for the device. | 
 **hardware** | **optional.String**| the name of the hardware model for the device. Use in conjunction with the manufacturer argument. However, an update will only occur if no hardware model is previously assigned to the device. | 
 **newHardware** | **optional.String**| If you want to change the hardware model for a device. If the device is rack mounted, it will unmount the device before changing hardware model and attempt to mount it back to same rack location (Changed in v6.3.2) | 
 **isItSwitch** | **optional.String**| ‘yes’ indicates if device is a network switch | 
 **isItVirtualHost** | **optional.String**| ‘yes’ indicates if device is a virtual host | 
 **isItBladeHost** | **optional.String**| ‘yes’ indicates if device is a blade host | 
 **inService** | **optional.String**| ‘yes’ indicates if device is in service, ‘no’ indicates not in service | 
 **type_** | **optional.String**| type of the device | 
 **serviceLevel** | **optional.String**| the service level name for the device. The service level must be pre-defined in the device42 appliance before it can be referenced in an api call. | 
 **virtualHost** | **optional.String**| Is the name for the host of the virtual machine. Two conditions must be met for this value to successfully update: 1. The device must be a virtual machine. 2. The virtual host must already exist in the device42 application and must be already marked as a virtual host. | 
 **bladeHost** | **optional.String**| the name of the host for the blade machine. Two conditions must be met for this value to successfully update: 1. The device must be a blade type. 2. The blade host must already exist and must already be marked as a blade host. | 
 **slotNo** | **optional.Int32**| slot # for blade device. | 
 **storageRoomId** | **optional.Int32**| ID of the room to assign device to storage room. Added in v5.5.0 | 
 **storageRoom** | **optional.String**| name of the room to assign device to, only used if the room name is unique. Added in v5.5.0 | 
 **os** | **optional.String**| the name of the operating system (os, osver, and osverno all required if updating any of the three). | 
 **osver** | **optional.String**| the version of the operating system (os, osver, and osverno all required if updating any of the three) | 
 **osarch** | **optional.String**| The architecture of the operating system (32 or 64) | 
 **osverno** | **optional.Int32**| the version # or build # of the operating system (os, osver, and osverno all required if updating any of the three) | 
 **memory** | **optional.String**| the total memory(RAM) in MB. Just numbers. | 
 **cpucount** | **optional.Int32**| total # of CPUs. (If updating, cpupower and cpucore required) | 
 **cpupower** | **optional.String**| CPU speed in MHz, just numbers. (If updating, cpucount and cpucore required) | 
 **cpucore** | **optional.String**| number of cores/CPU. Integer. (If updating, cpupower and cpucount required) | 
 **hddcount** | **optional.Int32**| total # of HDD (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddsize** | **optional.String**| HDD Size in GB (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddraid** | **optional.String**| Raid. Possible values: software or hardware. none to clear. (none added in v9.1.0) (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **hddraidType** | **optional.String**| Raid Type. (hddcount, hddsize, hddraid, and hddraid_type all required when updating any of the 4) | 
 **macaddress** | **optional.String**| the mac address value. Use multiple POST/PUT to add multiple mac addresses. | 
 **devicesInCluster** | **optional.String**| comma separated device names for devices in cluster. Only valid for device type cluster. | 
 **appcomps** | **optional.String**| comma separated application component names on this device. | 
 **customer** | **optional.String**| name of the customer. Customer record must already exist. | 
 **contractId** | **optional.Int32**| ID for the contract. Available via GET /api/api/1.0/contracts/ – DEPRECATED in v550. | 
 **contract** | **optional.String**| Name of the contract. Used only if contract name is unique. – DEPRECATED in v550. | 
 **aliases** | **optional.String**| optional. Comma separated aliases for the device. Must not exist, ignored otherwise. | 
 **subtype** | **optional.String**| Only for device type other. Must exist internally first. | 
 **bladeHostClear** | **optional.String**| clear blade host for device | 
 **virtualSubtype** | **optional.String**| Only for device type virtual. Default is internal VM. | [default to internal VM]
 **notes** | **optional.String**|  | 
 **tags** | **optional.String**| comma separated tags (Added in v5.9.2) | 
 **virtualHostClear** | **optional.String**| yes to clear virtual host for a VM (added in v9.1.0) | 
 **tagsRemove** | **optional.String**| comma separated tags to remove (added in v9.1.0) | 
 **aliasesRemove** | **optional.String**| comma separated aliases to remove | 
 **devicesInClusterRemove** | **optional.String**| comma separated device names for removing devices in cluster. Only valid for device type cluster. (added in v9.2.0) | 
 **objectCategory** | **optional.String**| If multitenancy is on, a category can be assigned to control access to this object, e.g. Prod_East specifies that the users with access to the Prod_East category will be able to access this device. | 
 **newObjectCategory** | **optional.String**| Used to change category on a device if it is already set. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDeviceUrl**
> interface{} PutDeviceUrl(ctx, optional)
Update Device URL

Use this API to update and existing url (e.g. an http, https, or telnet url) with a device. Type and device are required if no ID is provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiPutDeviceUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiPutDeviceUrlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| ID of the URL | 
 **type_** | **optional.String**| the url type (e.g. http, https, telnet). This must be an existing url type in device42. (required if no ID) | 
 **device** | **optional.String**| name of the device to which this url belongs (required if no ID) | 
 **host** | **optional.String**| fqdn portion of the url. e.g. for https://example.com/awesome, host is example.com | 
 **port** | **optional.String**| port number if any | 
 **urlSuffix** | **optional.String**| url suffix if any. e.g. it is “awesome” based on example given above. | 
 **notes** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadDevicesAttachments**
> interface{} UploadDevicesAttachments(ctx, deviceId, name, fileAttachmentDetails)
Upload and attach a file to a Device by Device Id.

Uploads a file attachment for a specific device using device id and a selected file.<br> Note that this implementation of Swagger does not support file upload - use another mechanism such as cULR.<br> <br>cURL example:<br> curl -X POST \\<br> URL/api/1.0/devices/id/3/attachments/17 \\<br> -H 'Accept-Encoding: gzip, deflate' \\<br> -H 'Authorization: Basic YWRtaW46YWRtIW5kNDI='<br> -H 'Content-Type: multipart/form-data; boundary=--------------------------216746989913126372322897' \\<br> -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \\<br> -F 'some test pdf.pdf=@/C:/path/to/attachment/Some test pdf.pdf' \\<br> -F 'file_attachment_details={\"some test pdf.pdf\":\"hello world! I'\\''m a pdf description!\"}'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **name** | **string**| The name of the file attachment you select for upload.&lt;br&gt; This filename is also required for the file_attachment_details parameter. | 
  **fileAttachmentDetails** | **string**| The name of the file to upload and a description.&lt;br&gt;Example: {\&quot;TEST_Attachment-3.txt\&quot;: \&quot;test attachment-3\&quot;}  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

