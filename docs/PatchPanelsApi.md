# \PatchPanelsApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPatchPanelID**](PatchPanelsApi.md#GetPatchPanelID) | **Get** /api/1.0/patch_panel/{ID}/ | Get Patch Panel details by ID
[**GetPatchPanelModels**](PatchPanelsApi.md#GetPatchPanelModels) | **Get** /api/1.0/patch_panel_models/ | Retrieves all patch panel models.
[**GetPatchPanelModuleModels**](PatchPanelsApi.md#GetPatchPanelModuleModels) | **Get** /api/1.0/patch_panel_module_models/ | Retrieves all patch panel module models.
[**GetPatchPanelPortsPatchPanelId**](PatchPanelsApi.md#GetPatchPanelPortsPatchPanelId) | **Get** /api/1.0/patch_panel_ports/{patch_panel_id}/ | Retrieves patch panel ports for specified patch panel.
[**PostPatchPanelModels**](PatchPanelsApi.md#PostPatchPanelModels) | **Post** /api/1.0/patch_panel_models/ | Create Patch Panel Model.
[**PostPatchPanelModuleModels**](PatchPanelsApi.md#PostPatchPanelModuleModels) | **Post** /api/1.0/patch_panel_module_models/ | Create/Update Patch Panel Module Model.
[**PostPatchPanelPortsPatchPanelId**](PatchPanelsApi.md#PostPatchPanelPortsPatchPanelId) | **Post** /api/1.0/patch_panel_ports/ | Update Patch Panel Ports


# **GetPatchPanelID**
> PatchPanel GetPatchPanelID(ctx, iD)
Get Patch Panel details by ID

Retrieve detailed information about a specific Patch Panel by Patch Panel ID. This also includes end point connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **int32**| The ID of the Patch Panel to retrieve | 

### Return type

[**PatchPanel**](Patch_panel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPatchPanelModels**
> []InlineResponse2002 GetPatchPanelModels(ctx, optional)
Retrieves all patch panel models.

Get all Patch Panel Models

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PatchPanelsApiGetPatchPanelModelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchPanelsApiGetPatchPanelModelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| filter by name | 

### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPatchPanelModuleModels**
> []InlineResponse2003 GetPatchPanelModuleModels(ctx, )
Retrieves all patch panel module models.

Get all Patch Panel Module Models

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPatchPanelPortsPatchPanelId**
> []InlineResponse200 GetPatchPanelPortsPatchPanelId(ctx, patchPanelId)
Retrieves patch panel ports for specified patch panel.

Get all Patch Panel Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patchPanelId** | **int32**| Patch panel id | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPatchPanelModels**
> interface{} PostPatchPanelModels(ctx, optional)
Create Patch Panel Model.

Create/Update Patch Panel Model. Required parameters: <ul><li>patch_panel_model_id <b>OR</b> name</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PatchPanelsApiPostPatchPanelModelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchPanelsApiPostPatchPanelModelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchPanelModelId** | **optional.String**| Patch Panel Model ID or UI Tools &gt; Export &gt; Patch Panel Model | 
 **name** | **optional.String**| Name of the patch panel model you want to edit or create. | 
 **manufacturer** | **optional.String**| name of the hardware/software manufacturer. | 
 **type_** | **optional.String**|  | [default to singluar]
 **portType** | **optional.String**| name of the port type. created if it doesn&#39;t exist already | 
 **pairedPorts** | **optional.String**|  | 
 **imgfileId** | **optional.String**| image file id. You can see these from Tools &gt; Import &gt; Hardware Import for now. | 
 **imgfile** | **optional.String**| name of the image file (Added in v5.8.2). Use instead of imgfile_id | 
 **numberOfPorts** | **optional.String**| number of ports. required for creating a new patch panel type singular. Ignored for patch panel type modular | 
 **numberOfPortsInRow** | **optional.String**| number of ports in a row. required for creating a new patch panel type singular. Ignored for patch panel type modular | 
 **modulePosition** | **optional.String**| for Modular Patch Panel Models. Possible values are horizontal or vertical (Added in v5.8.2) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPatchPanelModuleModels**
> interface{} PostPatchPanelModuleModels(ctx, optional)
Create/Update Patch Panel Module Model.

Create/Update Patch Panel Module Model. Required parameters: <ul><li>patch_panel_module_model_id <b>OR</b> name</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PatchPanelsApiPostPatchPanelModuleModelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchPanelsApiPostPatchPanelModuleModelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchPanelModuleModelId** | **optional.String**|  | 
 **name** | **optional.String**| Name of the patch panel module model you want to edit or create. | 
 **portType** | **optional.String**| name of the port type. created if it doesn&#39;t exist already | 
 **pairedPorts** | **optional.String**|  | 
 **numberOfPorts** | **optional.String**| number of ports. required for creating a new patch panel type singular. Ignored for patch panel type modular | 
 **numberOfPortsInRow** | **optional.String**| number of ports in a row. required for creating a new patch panel type singular. Ignored for patch panel type modular | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPatchPanelPortsPatchPanelId**
> []InlineResponse2001 PostPatchPanelPortsPatchPanelId(ctx, number, optional)
Update Patch Panel Ports

Update Patch Panel Ports. Requires the following parameters: <ul><li>patch_panel_id <b>OR</b> patch_panel</li> <li>number</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **number** | **string**| Port number for the patch panel | 
 **optional** | ***PatchPanelsApiPostPatchPanelPortsPatchPanelIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchPanelsApiPostPatchPanelPortsPatchPanelIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchPanelId** | **optional.String**| Patch panel id | 
 **patchPanel** | **optional.String**| Must be a unique asset name for this to work (Added in v5.8.2) | 
 **moduleSlot** | **optional.String**| change patch panel port for a certain module slot # (Added in v5.8.2) | 
 **macId** | **optional.String**| mac_id references port_id value. mac_id available via GET at /api/api/1.0/macs/ or Tools &gt; Export &gt; MAC Address * Please see note below | 
 **macAddress** | **optional.String**| mac_address or hwaddress | 
 **switchportId** | **optional.String**| ID for the switch port. Available via GET at /api/api/1.0/switchports/ or Tools &gt; Export &gt; Switch Port | 
 **switch_** | **optional.String**| Must be used with switchport if switchport_id is not used. | 
 **switchport** | **optional.String**| Must be used with switch if switchport_id is not used. | 
 **patchPanelPortId** | **optional.String**| Or use front_patch_panel and front_port combination | 
 **frontPatchPanelId** | **optional.String**| Name of the front patch panel - if a patch panel port is connected in front (Added in v5.8.2) | 
 **frontPatchPanel** | **optional.String**| Name of the front patch panel - if a patch panel port is connected in front (Added in v5.8.2) | 
 **frontPort** | **optional.String**| Number of the port on the front patch panel - if a patch panel port is connected in front (Added in v5.8.2) | 
 **label** | **optional.String**|  | 
 **objLabel1** | **optional.String**| object label 1 | 
 **objLabel2** | **optional.String**| object label 2 | 
 **backConnectionId** | **optional.String**| ID for the back connection port. Available via GET at /api/api/1.0/patch_panel_ports// or Tools &gt; Export &gt; Patch Panel Port | 
 **backSwitchportId** | **optional.String**| Used if back connection type is switch. | 
 **backSwitch** | **optional.String**| If back connection type is switch, use switch and switchport names in combination. | 
 **backSwitchport** | **optional.String**| Use if back connection type is switch - use in combination with back_switch. | 
 **backPatchPanelId** | **optional.String**| ID of the back patch panel - if a patch panel port is connected in back (Added in v5.8.2) | 
 **backPatchPanel** | **optional.String**| Name of the back patch panel - if a patch panel port is connected in back (Added in v5.8.2) | 
 **backPort** | **optional.String**| Number of the port on the back patch panel - if a patch panel port is connected in back (Added in v5.8.2) | 
 **clearFront** | **optional.String**| “yes” will clear front connection for port | 
 **clearBack** | **optional.String**| “yes” will clear back connection on port | 
 **cableType** | **optional.String**| named value of the cable type. Must already exist. | 

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

