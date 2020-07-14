# \ServicesApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIgnoredServiceByID**](ServicesApi.md#DeleteIgnoredServiceByID) | **Delete** /api/2.0/ignored_service/{id}/ | Delete Ignored Service By ID
[**DeleteScheduledTasksByID**](ServicesApi.md#DeleteScheduledTasksByID) | **Delete** /api/2.0/scheduled_tasks/{id} | Delete Scheduled Task by ID
[**DeleteServiceDetails**](ServicesApi.md#DeleteServiceDetails) | **Delete** /api/1.0/service_detail/{id}/ | This API is used to delete the service detail with the service detail id supplied as the required argument. Note: You will only be able to delete the service if the supplied username has the correct permissions.
[**DeleteServiceInstancesByID**](ServicesApi.md#DeleteServiceInstancesByID) | **Delete** /api/2.0/service_instances/{id} | Delete Service Instances By ID
[**DeleteServicePorts**](ServicesApi.md#DeleteServicePorts) | **Delete** /api/1.0/service_ports/{id}/ | This API is used to delete the service port with the service port id supplied as the required argument. Note: You will only be able to delete the service port if the supplied username has the correct permissions.
[**DeleteServices**](ServicesApi.md#DeleteServices) | **Delete** /api/1.0/services/{id}/ | This API is used to delete the service with the service id supplied as the required argument. Note: You will only be able to delete the service if the supplied username has the correct permissions.
[**DeleteServicesByID**](ServicesApi.md#DeleteServicesByID) | **Delete** /api/2.0/services/{id} | Delete Services By ID
[**GetIgnoredService**](ServicesApi.md#GetIgnoredService) | **Get** /api/2.0/ignored_service/ | Get Ignored Service
[**GetListenerConnectionStats**](ServicesApi.md#GetListenerConnectionStats) | **Get** /api/1.0/listener_connection_stats/ | 
[**GetListenerConnectionStatsByID**](ServicesApi.md#GetListenerConnectionStatsByID) | **Get** /api/2.0/listener_connection_stats/{id} | Retrieve listener connection statistics by service port ID
[**GetListenerConnectionStats_0**](ServicesApi.md#GetListenerConnectionStats_0) | **Get** /api/2.0/listener_connection_stats/ | Retrieve all listener connection statistics
[**GetNetworkShares**](ServicesApi.md#GetNetworkShares) | **Get** /api/1.0/network_shares/ | 
[**GetScheduledTasks**](ServicesApi.md#GetScheduledTasks) | **Get** /api/2.0/scheduled_tasks/ | Retreive information about all Scheduled Tasks
[**GetScheduledTasksByID**](ServicesApi.md#GetScheduledTasksByID) | **Get** /api/2.0/scheduled_tasks/{id} | Retrieve Scheduled Task by Service Schedule ID
[**GetServiceClientConnections**](ServicesApi.md#GetServiceClientConnections) | **Get** /api/2.0/service_client_connections/{id} | Get Service Client Connection information by Service Detail ID
[**GetServiceDetails**](ServicesApi.md#GetServiceDetails) | **Get** /api/1.0/service_details/ | 
[**GetServiceInstances**](ServicesApi.md#GetServiceInstances) | **Get** /api/2.0/service_instances/ | Get Service Instances
[**GetServiceInstancesByID**](ServicesApi.md#GetServiceInstancesByID) | **Get** /api/2.0/service_instances/{id} | Retrieve Service Instance information by Service Instance ID
[**GetServiceListenerPorts**](ServicesApi.md#GetServiceListenerPorts) | **Get** /api/2.0/service_listener_ports/ | Get Service Listener Ports
[**GetServiceListenerPortsByID**](ServicesApi.md#GetServiceListenerPortsByID) | **Get** /api/2.0/service_listener_ports/{id}/ | Retrieve Service Listener port information by Service Port ID
[**GetServicePorts**](ServicesApi.md#GetServicePorts) | **Get** /api/1.0/service_ports/ | 
[**GetServices**](ServicesApi.md#GetServices) | **Get** /api/1.0/services/ | 
[**GetServices2**](ServicesApi.md#GetServices2) | **Get** /api/2.0/services/ | Get list of all Services
[**PostIgnoredService**](ServicesApi.md#PostIgnoredService) | **Post** /api/2.0/ignored_service/ | Create / Update Ignored Service
[**PostScheduledTasks**](ServicesApi.md#PostScheduledTasks) | **Post** /api/2.0/scheduled_tasks/ | Create / Update Scheduled Tasks
[**PostServiceDetails**](ServicesApi.md#PostServiceDetails) | **Post** /api/1.0/service_details/ | 
[**PostServiceInstances**](ServicesApi.md#PostServiceInstances) | **Post** /api/2.0/service_instances/ | Create / Update service instances
[**PostServiceListenerPorts**](ServicesApi.md#PostServiceListenerPorts) | **Post** /api/2.0/service_listener_ports/ | Create / Update service ports
[**PostServicePorts**](ServicesApi.md#PostServicePorts) | **Post** /api/1.0/service_ports/ | 
[**PostServices**](ServicesApi.md#PostServices) | **Post** /api/1.0/services/ | 
[**PostServices2**](ServicesApi.md#PostServices2) | **Post** /api/2.0/services/ | Create / Update Services
[**PutServiceListenerPorts**](ServicesApi.md#PutServiceListenerPorts) | **Put** /api/2.0/service_listener_ports/{id}/ | Create / Update service ports


# **DeleteIgnoredServiceByID**
> interface{} DeleteIgnoredServiceByID(ctx, id)
Delete Ignored Service By ID

This API is used to delete the ignored service with the ignored service id supplied as the required argument. Note: You will only be able to delete the service if the supplied username has the correct permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScheduledTasksByID**
> interface{} DeleteScheduledTasksByID(ctx, id)
Delete Scheduled Task by ID

Used to delete the scheduled task with the service schedule ID as the required argument. Note: You will only be able to delete the scheduled task if the supplied username has the correct permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service schedule id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceDetails**
> interface{} DeleteServiceDetails(ctx, id)
This API is used to delete the service detail with the service detail id supplied as the required argument. Note: You will only be able to delete the service if the supplied username has the correct permissions.

Delete Service detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| service detail id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceInstancesByID**
> interface{} DeleteServiceInstancesByID(ctx, id)
Delete Service Instances By ID

This API is used to delete the service detail with the service detail id supplied as the required argument. Note - You will only be able to delete the service if the supplied username has the correct permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServicePorts**
> interface{} DeleteServicePorts(ctx, id)
This API is used to delete the service port with the service port id supplied as the required argument. Note: You will only be able to delete the service port if the supplied username has the correct permissions.

Delete Service port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| service port id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServices**
> interface{} DeleteServices(ctx, id)
This API is used to delete the service with the service id supplied as the required argument. Note: You will only be able to delete the service if the supplied username has the correct permissions.

Delete Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| service id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServicesByID**
> interface{} DeleteServicesByID(ctx, id)
Delete Services By ID

This API is used to delete the service with the service id supplied as the required argument. Note: You will only be able to delete the service if the supplied username has the correct permissions. (Added in v6.3.4)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIgnoredService**
> interface{} GetIgnoredService(ctx, optional)
Get Ignored Service

Retrieve list of Ignored Services; filter ignored service details by following parameters in query string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiGetIgnoredServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetIgnoredServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ignoredId** | **optional.Int32**| service id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListenerConnectionStats**
> interface{} GetListenerConnectionStats(ctx, port, optional)


Get Listener Connection Statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **port** | **string**| child is optional assuming you have a parent created (see examples) | 
 **optional** | ***ServicesApiGetListenerConnectionStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetListenerConnectionStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listenerDeviceName** | **optional.String**| Name of device with listening services | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListenerConnectionStatsByID**
> ListenerConnectionStatsId GetListenerConnectionStatsByID(ctx, id)
Retrieve listener connection statistics by service port ID

Get Listener Connection Statistics By ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service port id | 

### Return type

[**ListenerConnectionStatsId**](Listener_connection_stats_id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListenerConnectionStats_0**
> ListenerConnectionStats GetListenerConnectionStats_0(ctx, optional)
Retrieve all listener connection statistics

Get Listener Connection Statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiGetListenerConnectionStats_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetListenerConnectionStats_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **port** | **optional.String**| child is optional assuming you have a parent created (see examples) | 
 **deviceName** | **optional.String**| the name of the device | 
 **listenerDeviceName** | **optional.String**| Name of device with listening services | 
 **discoveredService** | **optional.String**| the name of the discovered service listening on this IP/port | 
 **mappedService** | **optional.String**| the name of the mapped service listening on this IP/port | 
 **serviceInstanceId** | **optional.Int32**| filter by id of the service instance in use | 

### Return type

[**ListenerConnectionStats**](Listener_connection_stats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkShares**
> interface{} GetNetworkShares(ctx, )


Get Network Shares

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

# **GetScheduledTasks**
> interface{} GetScheduledTasks(ctx, optional)
Retreive information about all Scheduled Tasks

Get Scheduled Tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiGetScheduledTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetScheduledTasksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int32**| filter results by Service schedule ID | 
 **deviceId** | **optional.Int32**| device id | 
 **device** | **optional.String**| device | 
 **userId** | **optional.String**| user id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduledTasksByID**
> ServiceSchedule GetScheduledTasksByID(ctx, id)
Retrieve Scheduled Task by Service Schedule ID

Get scheduled task By ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service schedule id | 

### Return type

[**ServiceSchedule**](Service_schedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceClientConnections**
> ServiceClientConnections GetServiceClientConnections(ctx, id)
Get Service Client Connection information by Service Detail ID

Get Service Client Connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service id | 

### Return type

[**ServiceClientConnections**](Service_client_connections.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceDetails**
> interface{} GetServiceDetails(ctx, deviceId, optional)


Get Service details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| Device ID | 
 **optional** | ***ServicesApiGetServiceDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServiceDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **device** | **optional.String**| Device name | 
 **serviceId** | **optional.String**| filter by id of the service | 
 **serviceDetailId** | **optional.String**| filter by id of the service in use | 
 **userId** | **optional.String**| filter by id of the user | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInstances**
> interface{} GetServiceInstances(ctx, optional)
Get Service Instances

You can filter service details by following parameters in the query string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiGetServiceInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServiceInstancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.Int32**| filter by id of the device | 
 **device** | **optional.String**| filter by name of the device | 
 **serviceId** | **optional.Int32**| filter by id of the service | 
 **serviceDetailId** | **optional.Int32**| filter by id of the service in use | 
 **userId** | **optional.String**| filter by id of the user | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInstancesByID**
> ServiceInstance GetServiceInstancesByID(ctx, id)
Retrieve Service Instance information by Service Instance ID

Get Service Instances By ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service instance id | 

### Return type

[**ServiceInstance**](Service_instance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceListenerPorts**
> ServiceListenerPorts GetServiceListenerPorts(ctx, optional)
Get Service Listener Ports

You can filter service ports by following parameters in the query string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiGetServiceListenerPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServiceListenerPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **port** | **optional.String**| port | 
 **deviceName** | **optional.String**| the name of the device | 
 **listenerDeviceName** | **optional.String**| Name of device with listening services | 
 **discoveredService** | **optional.String**| the name of the discovered service listening on this IP/port | 
 **mappedService** | **optional.String**| the name of the mapped service listening on this IP/port | 
 **serviceInstanceId** | **optional.Int32**| filter by id of the service instance in use | 

### Return type

[**ServiceListenerPorts**](Service_listener_ports.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceListenerPortsByID**
> ServiceListenerPort GetServiceListenerPortsByID(ctx, id)
Retrieve Service Listener port information by Service Port ID

Get Service Listener Ports By ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service port id | 

### Return type

[**ServiceListenerPort**](Service_listener_port.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServicePorts**
> interface{} GetServicePorts(ctx, optional)


Get Service ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiGetServicePortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServicePortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| filter by id of the service port | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServices**
> interface{} GetServices(ctx, serviceId, optional)


Get Services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| D42 ID of the service | 
 **optional** | ***ServicesApiGetServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **category** | **optional.String**| name of the category | 
 **vendor** | **optional.String**| The cloud vendor | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServices2**
> interface{} GetServices2(ctx, optional)
Get list of all Services

You can filter services by following parameters in the query string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiGetServices2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiGetServices2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int32**| D42 ID of the service | 
 **displayname** | **optional.String**| name of the service | 
 **category** | **optional.String**| name of the category | 
 **vendor** | **optional.String**| vendor name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIgnoredService**
> interface{} PostIgnoredService(ctx, ignoredLevel, ignoredText, active)
Create / Update Ignored Service

Post Ignored Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ignoredLevel** | **int32**| ignored level - 1 for Service Name, 2 for Service Instance Args | 
  **ignoredText** | **string**| ignored text | 
  **active** | **bool**| active flag | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostScheduledTasks**
> interface{} PostScheduledTasks(ctx, caption, optional)
Create / Update Scheduled Tasks

Post Scheduled Tasks. 3 parameters are required: <ul><li>schedule_name <b>OR</b> service_name <b>OR</b> name</li> <li>device_id <b>OR</b> device</li> <li>caption</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **caption** | **string**| caption | 
 **optional** | ***ServicesApiPostScheduledTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiPostScheduledTasksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleId** | **optional.Int32**| schedule id | 
 **serviceDetailId** | **optional.Int32**| filter by id of the service in use | 
 **scheduleName** | **optional.String**| schedule name | 
 **serviceName** | **optional.String**| The executable name of the service | 
 **name** | **optional.String**| name | 
 **device** | **optional.String**| The device that this service runs on | 
 **deviceId** | **optional.Int32**| The ID of the device that this service runs on | 
 **userId** | **optional.String**| end user id | 
 **schedDescription** | **optional.String**| schedule description | 
 **arguments** | **optional.String**| arguments | 
 **status** | **optional.String**| status | 
 **installDate** | **optional.String**| (note capital D for schedules) only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **atLogon** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **atStartup** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **eventBased** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **idleTime** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **otherTrigger** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **otherType** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **minutes** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **hours** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **days** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **weeks** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **dayOfMonth** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **monthOfYear** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 
 **dayOfWeek** | **optional.String**| only for schedule based services where startmode &#x3D; &#39;Scheduled&#39; | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceDetails**
> interface{} PostServiceDetails(ctx, serviceName, serviceDisplayName, optional)


Create / Update service details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceName** | **string**| The executable name of the service | 
  **serviceDisplayName** | **string**| The user freindly display name of the service | 
 **optional** | ***ServicesApiPostServiceDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiPostServiceDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vendor** | **optional.String**| The cloud vendor | 
 **description** | **optional.String**|  | 
 **serviceType** | **optional.String**| could be ignored or tracked. Default is tracked. | 
 **category** | **optional.String**| name of the category | 
 **startmode** | **optional.String**| The start mode of this service - valid values are ‘Automatic’, ‘Manual’, ‘Disabled’ and ‘Unknown’ | 
 **state** | **optional.String**| The current running state of this service. Valid values are ‘Running’, ‘Started’, ‘Paused’, ‘Stopped’ and ‘Unknown’ | 
 **device** | **optional.String**| Device name | 
 **user** | **optional.String**| enduser name | 
 **appcomp** | **optional.String**| The application component that depends on this service | 
 **installDate** | **optional.String**| The date that the software was installed | 
 **status** | **optional.String**| Instance status (ie, running, stopped) | 
 **atLogon** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **atStartup** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **eventBased** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **idleTime** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **otherTrigger** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **otherType** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **minutes** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **hours** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **days** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **weeks** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **dayOfMonth** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **monthOfYear** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 
 **dayOfWeek** | **optional.String**| only for schedule based services where startmode &#x3D; ‘Scheduled’ | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceInstances**
> interface{} PostServiceInstances(ctx, optional)
Create / Update service instances

Post Service Instances. Required parameters are either: <ul><li>service_name <b>OR</b> service_display_name <b>OR</b> service_id</li> <li> device <b>OR</b> device_id</li><p>------ OR ------<li>service_detail_id</li></ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiPostServiceInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiPostServiceInstancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceDetailId** | **optional.Int32**| service detail id | 
 **device** | **optional.String**| The device that this service runs on | 
 **deviceId** | **optional.Int32**| id of the device | 
 **listeningIp** | **optional.String**| the local IP address that listening on this port | 
 **port** | **optional.Int32**| the listening port on this device | 
 **loadbalancer** | **optional.String**| loadbalancer | 
 **serviceName** | **optional.String**| The executable name of the service | 
 **serviceDisplayName** | **optional.String**| The user freindly display name of the service | 
 **serviceId** | **optional.Int32**| service id | 
 **userId** | **optional.String**| id of the user | 
 **removeAppcompIds** | **optional.Int32**| comma separated list of application components IDs to remove | 
 **clearAppcomps** | **optional.String**| Set to &#39;yes&#39; to clear associated application components | 
 **appcomps** | **optional.String**| comma separated list of application components to associated with service instance | 
 **servicePath** | **optional.String**| service path | 
 **installDate** | **optional.String**| install date | 
 **state** | **optional.String**| The current running state of this service. Valid values are &#39;Running&#39;, &#39;Started&#39;, &#39;Paused&#39;, &#39;Stopped&#39; and &#39;Unknown&#39; | 
 **ignoreClientConnections** | **optional.String**|  | 
 **startmode** | **optional.String**| The start mode of this service - valid values are &#39;Automatic&#39;, &#39;Manual&#39;, &#39;Disabled&#39; and &#39;Unknown&#39; | 
 **pinned** | **optional.String**| pinned | 
 **topologyStatus** | **optional.String**| topology status | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceListenerPorts**
> interface{} PostServiceListenerPorts(ctx, optional)
Create / Update service ports

Post Service Listener Ports. Required parameters: <ul> <li>id</li></ul> <b>---------- OR ----------</b> <ul> <li>port</li> <li>device_name <b>OR</b> device_id</li> <li>listening_ip <b>OR</b> remote_ips</li> </ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiPostServiceListenerPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiPostServiceListenerPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int32**| service port id | 
 **port** | **optional.Int32**| the listening port on this device | 
 **deviceName** | **optional.String**| the name of the device | 
 **deviceId** | **optional.Int32**| device id | 
 **listeningIp** | **optional.String**| the name of the device | 
 **remoteIps** | **optional.String**| the comma separated list of remote IPs that are connected to this listening IP/port | 
 **discoveredServiceId** | **optional.Int32**| device service id | 
 **discoveredService** | **optional.String**| the name of the discovered service listening on this IP/port | 
 **discoveredServiceDisplayName** | **optional.String**| discovered service display name | 
 **discoveredProcess** | **optional.String**| the process name that has a handle to the port | 
 **discoveredProcessDisplayName** | **optional.String**| discovered process display name | 
 **mappedService** | **optional.String**| the name of the mapped service listening on this IP/port | 
 **mappedServiceDisplayName** | **optional.String**| mapped service display name | 
 **protocol** | **optional.String**| the transport protocol, ie TCP | 
 **remotePortTime** | **optional.String**| remote port time | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServicePorts**
> interface{} PostServicePorts(ctx, port, optional)


Create / Update service ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **port** | **string**| child is optional assuming you have a parent created (see examples) | 
 **optional** | ***ServicesApiPostServicePortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiPostServicePortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceName** | **optional.String**| required if switchport ID not used | 
 **listeningIp** | **optional.String**| the name of the device | 
 **description** | **optional.String**| matching description (added in v7.2.0) | 
 **discoveredService** | **optional.String**| the name of the discovered service listening on this IP/port | 
 **discoveredProcess** | **optional.String**| the process name that has a handle to the port | 
 **mappedService** | **optional.String**| the name of the mapped service listening on this IP/port | 
 **remoteIps** | **optional.String**| the comma separated list of remote IPs that are connected to this listening IP/port | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServices**
> interface{} PostServices(ctx, displayName, optional)


Create / Update Services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **displayName** | **string**| if not provided, the name is used as display name | 
 **optional** | ***ServicesApiPostServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiPostServicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **category** | **optional.String**| name of the category | 
 **serviceType** | **optional.String**| could be ignored or tracked. Default is tracked. | 
 **vendor** | **optional.String**| The cloud vendor | 
 **description** | **optional.String**|  | 
 **notes** | **optional.String**| Any additional notes | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServices2**
> interface{} PostServices2(ctx, optional)
Create / Update Services

Post a Service. Required parameters: <ul> <li>name <b>OR</b> display_name</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServicesApiPostServices2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiPostServices2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| name for service | 
 **displayName** | **optional.String**| if not provided, the name is used as display name | 
 **vendor** | **optional.String**| Names of vendor | 
 **description** | **optional.String**| description of service | 
 **notes** | **optional.String**| notes | 
 **category** | **optional.String**| Name of category | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutServiceListenerPorts**
> interface{} PutServiceListenerPorts(ctx, id, optional)
Create / Update service ports

Currently, in order to update a service port, id (path), device_name/device_id, and listening_ip/remote_ips are needed even if they are staying the same - we use those fields to identify the port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| service port id | 
 **optional** | ***ServicesApiPutServiceListenerPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiPutServiceListenerPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **port** | **optional.Int32**| the listening port on this device | 
 **deviceName** | **optional.String**| the name of the device | 
 **deviceId** | **optional.Int32**| device id | 
 **listeningIp** | **optional.String**| the name of the device | 
 **remoteIps** | **optional.String**| the comma separated list of remote IPs that are connected to this listening IP/port | 
 **discoveredServiceId** | **optional.Int32**| device service id | 
 **discoveredService** | **optional.String**| the name of the discovered service listening on this IP/port | 
 **discoveredServiceDisplayName** | **optional.String**| discovered service display name | 
 **discoveredProcess** | **optional.String**| the process name that has a handle to the port | 
 **discoveredProcessDisplayName** | **optional.String**| discovered process display name | 
 **mappedService** | **optional.String**| the name of the mapped service listening on this IP/port | 
 **mappedServiceDisplayName** | **optional.String**| mapped service display name | 
 **protocol** | **optional.String**| the transport protocol, ie TCP | 
 **remotePortTime** | **optional.String**| remote port time | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

