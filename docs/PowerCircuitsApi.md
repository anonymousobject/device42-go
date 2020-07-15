# \PowerCircuitsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePowerCircuit**](PowerCircuitsApi.md#DeletePowerCircuit) | **Delete** /api/1.0/powercircuits/{id}/ | Power Circuit
[**GetAllPowerCircuits**](PowerCircuitsApi.md#GetAllPowerCircuits) | **Get** /api/1.0/powercircuits/ | Get All Power Circuits
[**PostUpdatePowerCircuits**](PowerCircuitsApi.md#PostUpdatePowerCircuits) | **Post** /api/1.0/powercircuits/ | Create/Update Power Circuit


# **DeletePowerCircuit**
> interface{} DeletePowerCircuit(ctx, id)
Power Circuit

This API is used to delete the power circuit with the ID supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Power circuit ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPowerCircuits**
> interface{} GetAllPowerCircuits(ctx, optional)
Get All Power Circuits

Get All Power Circuits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PowerCircuitsApiGetAllPowerCircuitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PowerCircuitsApiGetAllPowerCircuitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **breakerpanelId** | **optional.String**| Breaker panel ID | 
 **bcpmId** | **optional.String**| Branch Circuit Power Meter ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdatePowerCircuits**
> interface{} PostUpdatePowerCircuits(ctx, breakerpanelId, number, optional)
Create/Update Power Circuit

Create/Update Power Circuit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **breakerpanelId** | **string**| Breaker panel ID | 
  **number** | **string**| Number of the circuit on the breaker panel | 
 **optional** | ***PowerCircuitsApiPostUpdatePowerCircuitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PowerCircuitsApiPostUpdatePowerCircuitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | **optional.String**| Label of the circuit on the breaker panel | 
 **connection** | **optional.String**| Description of the connection for the circuit | 
 **voltage** | **optional.String**| The rated voltage on this circuit | 
 **amps** | **optional.String**| The rated amps on this circuit | 
 **powerunitConnectionIds** | **optional.String**| A comma-separated list of PU IDs connected to this circuit | 
 **deviceConnectionIds** | **optional.String**| A comma-separated list of device IDs connected to this circuit | 
 **deviceConnectionNames** | **optional.String**| A comma-separated list of device names connected to this circuit | 
 **assetConnectionIds** | **optional.String**| A comma-separated list of asset IDs connected to this circuit | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

