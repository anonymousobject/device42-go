# \CablesApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCable**](CablesApi.md#DeleteCable) | **Delete** /api/1.0/cables/{id}/ | Delete Cable
[**GetCables**](CablesApi.md#GetCables) | **Get** /api/1.0/cables/ | Retrieve information about all cables.
[**PostCables**](CablesApi.md#PostCables) | **Post** /api/1.0/cables/ | Create or update Cables


# **DeleteCable**
> interface{} DeleteCable(ctx, id)
Delete Cable

This API is used to delete the cable with the ID supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Cable ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCables**
> interface{} GetCables(ctx, optional)
Retrieve information about all cables.

Get All Cables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CablesApiGetCablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CablesApiGetCablesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cableId** | **optional.Int32**| filter by cable_id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCables**
> interface{} PostCables(ctx, cableId, optional)
Create or update Cables

Create/Update Cable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cableId** | **string**| Cable ID/Name | 
 **optional** | ***CablesApiPostCablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CablesApiPostCablesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iD** | **optional.String**| Device42 ID of cable | 
 **vendor** | **optional.String**| Cable vendor | 
 **notes** | **optional.String**| Any additional notes | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 
 **roomId** | **optional.String**| Room ID | 
 **room** | **optional.String**| Room name | 
 **originType** | **optional.String**| Type of origin point. | 
 **originId** | **optional.String**| ID of the origin point | 
 **originConnectorType** | **optional.String**| Connector Type (User Definable) | 
 **originCableType** | **optional.String**| Cable Type (User definable) | 
 **originCableColor** | **optional.String**| Origin Cable Color | 
 **originOpticType** | **optional.String**| Optic Type (Definable, ie multimode) | 
 **originBackPatchPanel** | **optional.String**|  | 
 **endPointType** | **optional.String**| Type of end point. | 
 **endPointId** | **optional.String**| ID of the end point | 
 **endConnectorType** | **optional.String**| Connector Type (User Definable) | 
 **endCableType** | **optional.String**| Endpoint Cable Type (User Definable) | 
 **endPointCableColor** | **optional.String**| Endpoint Cable Color | 
 **endPointOpticType** | **optional.String**| Optic Type (Definable, ie multimode) | 
 **endPointBackPachPanel** | **optional.String**|  | 
 **endPointMultiple** | **optional.String**| yes to allow multiple endpoints | 
 **cableLength** | **optional.String**| Length of Cable | 
 **cableLengthUnits** | **optional.String**| Units for Cable Length (“m” or “ft”) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

