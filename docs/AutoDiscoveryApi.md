# \AutoDiscoveryApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoDiscoveryCertificate**](AutoDiscoveryApi.md#GetAutoDiscoveryCertificate) | **Get** /api/1.0/auto_discovery/certificate/ | Retrieves all auto-discovery jobs for dns.
[**GetAutoDiscoveryCloudaccount**](AutoDiscoveryApi.md#GetAutoDiscoveryCloudaccount) | **Get** /api/1.0/auto_discovery/cloudaccount/ | Retrieves all cloud auto-discovery jobs.
[**GetAutoDiscoveryDns**](AutoDiscoveryApi.md#GetAutoDiscoveryDns) | **Get** /api/1.0/auto_discovery/dns/ | Retrieves all auto-discovery jobs for dns.
[**GetAutoDiscoveryIpmi**](AutoDiscoveryApi.md#GetAutoDiscoveryIpmi) | **Get** /api/1.0/auto_discovery/ipmi/ | Retrieves all IPMI/Redfish auto-discovery jobs.
[**GetAutoDiscoveryNetworks**](AutoDiscoveryApi.md#GetAutoDiscoveryNetworks) | **Get** /api/1.0/auto_discovery/networks/ | Retrieves all auto-discovery jobs for networking devices.
[**GetAutoDiscoveryPingsweep**](AutoDiscoveryApi.md#GetAutoDiscoveryPingsweep) | **Get** /api/1.0/auto_discovery/pingsweep/ | Retrieves all auto-discovery pingsweep jobs.
[**GetAutoDiscoverySnmpDisc**](AutoDiscoveryApi.md#GetAutoDiscoverySnmpDisc) | **Get** /api/1.0/auto_discovery/snmp_disc/ | Retrieves all SNMP auto-discovery jobs.
[**GetAutoDiscoveryUcs**](AutoDiscoveryApi.md#GetAutoDiscoveryUcs) | **Get** /api/1.0/auto_discovery/ucs/ | Retrieves all UCS/ACI/Load Balancer auto-discovery jobs.
[**GetAutoDiscoverynmap**](AutoDiscoveryApi.md#GetAutoDiscoverynmap) | **Get** /api/1.0/auto_discovery/nmap/ | Retrieves all auto-discovery Nmap jobs.
[**GetAutoDiscoveryvServer**](AutoDiscoveryApi.md#GetAutoDiscoveryvServer) | **Get** /api/1.0/auto_discovery/vserver/ | Retrieves all vServer auto-discovery jobs.
[**PostAutoDiscoveryBladeDisc**](AutoDiscoveryApi.md#PostAutoDiscoveryBladeDisc) | **Post** /api/1.0/auto_discovery/blade_disc/ | Add or update blade autodiscovery job
[**PostAutoDiscoveryCertificate**](AutoDiscoveryApi.md#PostAutoDiscoveryCertificate) | **Post** /api/1.0/auto_discovery/certificate/ | Add Certificate Auto Discovery
[**PostAutoDiscoveryCloudaccount**](AutoDiscoveryApi.md#PostAutoDiscoveryCloudaccount) | **Post** /api/1.0/auto_discovery/cloudaccount/ | Add or update cloud autodiscovery job
[**PostAutoDiscoveryDns**](AutoDiscoveryApi.md#PostAutoDiscoveryDns) | **Post** /api/1.0/auto_discovery/dns/ | Add DNS Auto Discovery
[**PostAutoDiscoveryIpmi**](AutoDiscoveryApi.md#PostAutoDiscoveryIpmi) | **Post** /api/1.0/auto_discovery/ipmi/ | Add IPMI/Redfish Auto Discovery
[**PostAutoDiscoveryNetworks**](AutoDiscoveryApi.md#PostAutoDiscoveryNetworks) | **Post** /api/1.0/auto_discovery/networks/ | 
[**PostAutoDiscoveryPingsweep**](AutoDiscoveryApi.md#PostAutoDiscoveryPingsweep) | **Post** /api/1.0/auto_discovery/pingsweep/ | 
[**PostAutoDiscoveryPowerDisc**](AutoDiscoveryApi.md#PostAutoDiscoveryPowerDisc) | **Post** /api/1.0/auto_discovery/power_disc/ | Add or update power autodiscovery job
[**PostAutoDiscoverySnmpDisc**](AutoDiscoveryApi.md#PostAutoDiscoverySnmpDisc) | **Post** /api/1.0/auto_discovery/snmp_disc/ | Add or update other SNMP autodiscovery job
[**PostAutoDiscoveryUcs**](AutoDiscoveryApi.md#PostAutoDiscoveryUcs) | **Post** /api/1.0/auto_discovery/ucs/ | Add or update UCS/ACI/Load Balancer Auto Discovery
[**PostAutoDiscoveryVserver**](AutoDiscoveryApi.md#PostAutoDiscoveryVserver) | **Post** /api/1.0/auto_discovery/vserver/ | Add or update vServer autodiscovery job
[**PostAutoDiscoverynmap**](AutoDiscoveryApi.md#PostAutoDiscoverynmap) | **Post** /api/1.0/auto_discovery/nmap/ | Add Nmap auto-discovery job
[**PutAutoDiscoveryBladeDisc**](AutoDiscoveryApi.md#PutAutoDiscoveryBladeDisc) | **Put** /api/1.0/auto_discovery/blade_disc/ | Execute blade autodiscovery job
[**PutAutoDiscoveryCertificate**](AutoDiscoveryApi.md#PutAutoDiscoveryCertificate) | **Put** /api/1.0/auto_discovery/certificate/ | Execute Certificate job
[**PutAutoDiscoveryCloudaccount**](AutoDiscoveryApi.md#PutAutoDiscoveryCloudaccount) | **Put** /api/1.0/auto_discovery/cloudaccount/ | Execute Cloud job
[**PutAutoDiscoveryDns**](AutoDiscoveryApi.md#PutAutoDiscoveryDns) | **Put** /api/1.0/auto_discovery/dns/ | Execute DNS job
[**PutAutoDiscoveryIpmi**](AutoDiscoveryApi.md#PutAutoDiscoveryIpmi) | **Put** /api/1.0/auto_discovery/ipmi/ | Execute snmp autodiscovery job
[**PutAutoDiscoveryNetworks**](AutoDiscoveryApi.md#PutAutoDiscoveryNetworks) | **Put** /api/1.0/auto_discovery/networks/ | Execute network job
[**PutAutoDiscoveryPingsweep**](AutoDiscoveryApi.md#PutAutoDiscoveryPingsweep) | **Put** /api/1.0/auto_discovery/pingsweep/ | Execute pingsweep job
[**PutAutoDiscoveryPowerDisc**](AutoDiscoveryApi.md#PutAutoDiscoveryPowerDisc) | **Put** /api/1.0/auto_discovery/power_disc/ | Execute power autodiscovery job
[**PutAutoDiscoverySnmpDisc**](AutoDiscoveryApi.md#PutAutoDiscoverySnmpDisc) | **Put** /api/1.0/auto_discovery/snmp_disc/ | Execute snmp autodiscovery job
[**PutAutoDiscoveryUcs**](AutoDiscoveryApi.md#PutAutoDiscoveryUcs) | **Put** /api/1.0/auto_discovery/ucs/ | Execute UCS/ACI/Load Balancer autodiscovery job
[**PutAutoDiscoveryVserver**](AutoDiscoveryApi.md#PutAutoDiscoveryVserver) | **Put** /api/1.0/auto_discovery/vserver/ | Execute vserver job
[**PutAutoDiscoverynmap**](AutoDiscoveryApi.md#PutAutoDiscoverynmap) | **Put** /api/1.0/auto_discovery/nmap/ | Execute Nmap job


# **GetAutoDiscoveryCertificate**
> interface{} GetAutoDiscoveryCertificate(ctx, )
Retrieves all auto-discovery jobs for dns.

Get Certificate Auto Discovery Jobs

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

# **GetAutoDiscoveryCloudaccount**
> interface{} GetAutoDiscoveryCloudaccount(ctx, )
Retrieves all cloud auto-discovery jobs.

Get all cloud discovery jobs

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

# **GetAutoDiscoveryDns**
> interface{} GetAutoDiscoveryDns(ctx, )
Retrieves all auto-discovery jobs for dns.

Get DNS Auto Discovery Jobs

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

# **GetAutoDiscoveryIpmi**
> interface{} GetAutoDiscoveryIpmi(ctx, )
Retrieves all IPMI/Redfish auto-discovery jobs.

Get all IPMI/Redfish auto-discovery jobs

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

# **GetAutoDiscoveryNetworks**
> interface{} GetAutoDiscoveryNetworks(ctx, )
Retrieves all auto-discovery jobs for networking devices.

Get Network Jobs

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

# **GetAutoDiscoveryPingsweep**
> interface{} GetAutoDiscoveryPingsweep(ctx, )
Retrieves all auto-discovery pingsweep jobs.

Get Pingsweep Jobs

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

# **GetAutoDiscoverySnmpDisc**
> interface{} GetAutoDiscoverySnmpDisc(ctx, )
Retrieves all SNMP auto-discovery jobs.

Get all SNMP discovery jobs

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

# **GetAutoDiscoveryUcs**
> interface{} GetAutoDiscoveryUcs(ctx, )
Retrieves all UCS/ACI/Load Balancer auto-discovery jobs.

Get all UCS/ACI/Load Balancer auto-discovery jobs

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

# **GetAutoDiscoverynmap**
> interface{} GetAutoDiscoverynmap(ctx, )
Retrieves all auto-discovery Nmap jobs.

Get Nmap Jobs

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

# **GetAutoDiscoveryvServer**
> interface{} GetAutoDiscoveryvServer(ctx, )
Retrieves all vServer auto-discovery jobs.

Get all vServer Jobs

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

# **PostAutoDiscoveryBladeDisc**
> interface{} PostAutoDiscoveryBladeDisc(ctx, name, optional)
Add or update blade autodiscovery job

Add or update blade autodiscovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryBladeDiscOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryBladeDiscOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipaddress** | **optional.String**| IP address. Required if new | 
 **endIpAddress** | **optional.String**| End IP address | 
 **snmpString** | **optional.String**| required, if new | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 
 **stripDomainName** | **optional.String**|  | 
 **clearExistingSchedule** | **optional.String**|  | 
 **groups** | **optional.String**| name of one or more groups separated by commas | 
 **snmpPort** | **optional.Int32**| snmp port (integer only) (added in v10.4.0) | 
 **snmpVersion** | **optional.String**|  | [default to v2c]
 **snmpv3User** | **optional.String**| name of snmp v3 user (added in v10.4.0) | 
 **snmpv3AuthMode** | **optional.String**|  | 
 **snmpv3AuthProtocol** | **optional.String**|  | 
 **snmpv3AuthPassword** | **optional.String**| password (added in v10.4.0) | 
 **snmpv3PrivacyProtocol** | **optional.String**|  | 
 **snmpv3PrivacyProtocolPassword** | **optional.String**| password (added in v10.4.0) | 
 **snmpv3Context** | **optional.String**|  | 
 **debugLevel** | **optional.String**|  | 
 **hostnamePrecedence** | **optional.String**|  | 
 **serviceLevel** | **optional.String**| Must already exist | 
 **nameToUseForNewlyDiscoveredModule** | **optional.String**| name to use | 
 **toggleInServiceOnModulePowerState** | **optional.String**| yes or no | 
 **moduleNotFound** | **optional.String**| action to take on module not found. One of: Remove Host Association, Change Service Level, Delete Module (default is no action) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryCertificate**
> interface{} PostAutoDiscoveryCertificate(ctx, name, startIpAddress, endIpAddress, optional)
Add Certificate Auto Discovery

Add Certificate Auto Discovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
  **startIpAddress** | **string**| beginning of IP address range | 
  **endIpAddress** | **string**| End IP address | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryCertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **server** | **optional.String**| IP or FQDN of target server | 
 **excludeServers** | **optional.String**| comma separated liste of addresses to exclude | 
 **ports** | **optional.String**| default is 443, comma separated | 
 **followChain** | **optional.String**|  | [default to no]
 **debugLevel** | **optional.String**|  | [default to no]
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryCloudaccount**
> PostAutoDiscoveryCloudaccount(ctx, name, cloudType, optional)
Add or update cloud autodiscovery job

Add or update a cloud autodiscovery job.<br> <p><b>Required parameters by cloud vendor or job type.</b><br>(API parameter = UI equivalent)</i><p>ALL DISCOVERY JOBS<br>- name = Name<br>- cloud_type = Cloud Type <p>AMAZON AWS<br>- accountid = Access Key ID<br>- secret_key OR secret_key_id = Secret Key<br>- regions = Regions <p>MICROSOFT AZURE<br>- auth_type = Authentication Type<br>- accountid = Client ID (Service Principal) or Username (User Credentials)<br>- secret_key OR secret_key_id = Client Secret (Service Principal) or Password (User Credentials)<br>- subscriptionid = Subscription ID<br>- tenant = Tenant ID <p>LINODE<br>- api_key OR api_key_id = API Key <p>DIGITALOCEAN<br>- token_key = Token Key <p>OPENSTACK<br>- ip = URL<br>- username = User<br>- secret_key OR secret_key_id = Password<br>- tenant = Project Name <p>AMAZON API<br>- ip = URL<br>- region = Region<br>- secret_key OR secret_key_id = Secret Key<br> - accountid = API Key <p>GOOGLE CLOUD<br>- tenant = Project ID<br>- secret_key OR secret_key_id = Credentials JSON<br>- google_regions = Zones <p>ALIBABA CLOUD<br>- accountid = Access Key ID<br>- secret_key OR secret_key_id = Access Key Secret <p>ORACLE CLOUD<br>- accountid = User ID<br>- api_key OR api_key_id = Fingerprint<br>- secret_key OR secret_key_id = Key File<br>- tenant = Tenant ID<br>- oracle_regions = Zones <p>STANDALONE KUBERNETES<br>- ip = URL<br>- auth_type = Authentication Type<br>- accountid = Bearer Token or Basic Credentials (depending on Authentication Type)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the discovery job; required for all jobs. | 
  **cloudType** | **string**| Required for all jobs. | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryCloudaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryCloudaccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secretKeyId** | **optional.String**| Amazon AWS, MS Azure, OpenStack, Amazon API, Google Cloud, Alibaba Cloud, Oracle Cloud | 
 **vendor** | **optional.String**| The cloud vendor | 
 **kubernetesDiscovery** | **optional.String**| Available for Amazon AWS, Microsoft Azure, and standalone Kubernetes. &lt;br&gt;Yes / No | 
 **remoteCollectorId** | **optional.String**| ID number of the remote collector (RC) to use for the job. | 
 **customerId** | **optional.String**| Customer ID number for discovered devices. | 
 **deviceNameFormat** | **optional.String**| Name format for discovered cloud instances (only for Alibaba Cloud, Amazon API, Amazon AWS, Google Cloud, MS Azure, Oracle Cloud). | 
 **containerNotFound** | **optional.String**| For Kubernetes discovery.&lt;br&gt;1 &#x3D; changes status, 2 &#x3D; delete container | 
 **authType** | **optional.String**| Type of authentication credentials for MS Azure and Standalone Kubernetes. | 
 **accountid** | **optional.String**| Required for Alibaba Cloud, Amazon AWS, Kubernetes, MS Azure, Oracle Cloud. | 
 **vrfgroup** | **optional.String**| name of vrf group for discovered subnets (added in v10.4.0) | 
 **notes** | **optional.String**| Any additional notes | 
 **removeUnfoundInstances** | **optional.String**|  | 
 **stripDomainName** | **optional.String**|  | 
 **addSuffix** | **optional.String**|  | 
 **matchNameOnlyForVirtuals** | **optional.String**|  | 
 **matchNameOnlyForHypervisor** | **optional.String**|  | 
 **groups** | **optional.String**| name of one or more groups separated by commas | 
 **objectCategory** | **optional.String**| category of discovered vservers and vms | 
 **debugLevel** | **optional.String**|  | 
 **clearExistingSchedule** | **optional.String**|  | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 
 **secretKey** | **optional.String**| Amazon AWS, MS Azure, OpenStack, Amazon API, Google Cloud, Alibaba Cloud, Oracle Cloud | 
 **regions** | **optional.String**| Comma separated region names for Amazon AWS | 
 **googleRegions** | **optional.String**| comma-separated list of Google region names | 
 **oracleRegions** | **optional.String**| comma-separated list of Oracle region names | 
 **subscriptionid** | **optional.String**| MS Azure Subscription ID | 
 **ip** | **optional.String**| Amazon API URL; OpenStack URL; Standalone Kubernetes URL | 
 **tenant** | **optional.String**| OpenStack Project name; Google Project ID; Oracle Tenant ID | 
 **username** | **optional.String**| OpenStack username - required for OpenStack | 
 **apiKey** | **optional.String**| Oracle Cloud Fingerprint; Linode API Key | 
 **apiKeyId** | **optional.String**| Oracle Cloud Fingerprint; Linode API Key | 
 **apiToken** | **optional.String**| Linode API token - Required for Linode | 
 **tokenKey** | **optional.String**| DigitalOcean Token Key - required for DigitalOcean. | 
 **region** | **optional.String**| Required for Amazon AWS API discovery. | 
 **alibabaRegions** | **optional.String**| Comma-separated list of Alibaba Cloud region names. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryDns**
> interface{} PostAutoDiscoveryDns(ctx, zonename, nameserver, optional)
Add DNS Auto Discovery

Add DNS Auto Discovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zonename** | **string**| name of the zone | 
  **nameserver** | **string**| IP/FQDN of the nameserver | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryDnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryDnsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remoteCollectorId** | **optional.String**| D42 ID of the remote collector to use for this job. | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryIpmi**
> interface{} PostAutoDiscoveryIpmi(ctx, name, ipStart, ipEnd, bmcUser, bmcPassword, hostnameToUse, optional)
Add IPMI/Redfish Auto Discovery

Add IPMI/Redfish Auto Discovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
  **ipStart** | **string**| starting IP address | 
  **ipEnd** | **string**| ending IP address, use same as start for single address | 
  **bmcUser** | **string**| username for discovery | 
  **bmcPassword** | **string**| password for discovery | 
  **hostnameToUse** | **string**|  | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryIpmiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryIpmiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **updateModelIfFound** | **optional.String**|  | 
 **runAsOperator** | **optional.String**|  | 
 **objectCategory** | **optional.String**|  | 
 **overwriteObjectCategories** | **optional.String**|  | 
 **clearExistingSchedule** | **optional.String**|  | 
 **groups** | **optional.String**| name of one or more groups separated by commas | 
 **debugLevel** | **optional.String**|  | 
 **discoveryType** | **optional.String**|  | [default to IPMI]
 **remoteCollectorId** | **optional.String**|  | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryNetworks**
> interface{} PostAutoDiscoveryNetworks(ctx, name, optional)


Add or update network job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryNetworksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipaddress** | **optional.String**| IP address. Required if new | 
 **endIpAddress** | **optional.String**| End IP address | 
 **snmpString** | **optional.String**| required, if new | 
 **snmpStrings** | **optional.String**| Can be comma separated list of community strings to use multiple community strings | 
 **snmpStringIds** | **optional.String**| Can be comma separated list of community string IDs to use multiple community strings | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 
 **stripDomainName** | **optional.String**|  | 
 **autodiscoverCdpDevices** | **optional.String**| yes to enable CDP/LLDP (added in v8.3.2) | 
 **useNameAliasPortDescr** | **optional.String**| yes to use alias for port description during discovery (added in v8.3.2) | 
 **clearExistingSchedule** | **optional.String**|  | 
 **groups** | **optional.String**| name of one or more groups separated by commas | 
 **maskBits** | **optional.Int32**| mask bits (integer only) | 
 **snmpPort** | **optional.Int32**| snmp port (integer only) (added in v10.4.0) | 
 **snmpVersion** | **optional.String**|  | [default to v2c]
 **snmpStringId** | **optional.String**| The id of the password for the community string | 
 **snmpv3AuthPasswordId** | **optional.String**| The id of the password for the auth password | 
 **snmpv3PrivacyProtocolPasswordId** | **optional.String**| The id of the password for the privacy protocol password | 
 **snmpv3User** | **optional.String**| name of snmp v3 user (added in v10.4.0) | 
 **snmpv3AuthMode** | **optional.String**|  | 
 **snmpv3AuthProtocol** | **optional.String**|  | 
 **snmpv3AuthPassword** | **optional.String**| password (added in v10.4.0) | 
 **snmpv3PrivacyProtocol** | **optional.String**|  | 
 **snmpv3PrivacyProtocolPassword** | **optional.String**| password (added in v10.4.0) | 
 **snmpv3Context** | **optional.String**|  | 
 **getAllSwitchPorts** | **optional.String**| yes or no to get all switch ports (added in v10.4.0) | 
 **deleteOlderMacAssociationAfter** | **optional.String**| number of days (added in v10.4.0) | 
 **deleteSwitchPortNotFound** | **optional.String**| yes or no to delete switch ports not found (added in v10.4.0) | 
 **vlansToIgnore** | **optional.String**| list of vlan ids to ignore separated by commas (added in v10.4.0) | 
 **portNamePrefixToIgnoreMacs** | **optional.String**|  | 
 **debugLevel** | **optional.String**|  | 
 **hostnamePrecedence** | **optional.String**|  | 
 **serviceLevel** | **optional.String**| Must already exist | 
 **skipVlanIndexing** | **optional.String**| yes or no (added in v10.4.0) | 
 **vrfgroup** | **optional.String**| name of vrf group for discovered subnets (added in v10.4.0) | 
 **objectCategory** | **optional.String**| name of subnet category for discovered subnets | 
 **discoverServices** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryPingsweep**
> interface{} PostAutoDiscoveryPingsweep(ctx, name, optional)


Add or update pingsweep job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryPingsweepOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryPingsweepOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networks** | **optional.String**| Networks to search. Required if new | 
 **type_** | **optional.String**| Choose if you want to automatically assign type to discovered IPs. | 
 **stripDomainSuffix** | **optional.String**| Strip everything after the first dot | 
 **clearExistingSchedule** | **optional.String**|  | 
 **debugLevel** | **optional.String**|  | 
 **reverseDns** | **optional.String**| Add devices by reverse DNS of discovered IP. (Recommended to leave unchecked and use other discovery methods for devices) | 
 **createNewSubnet** | **optional.String**| Create new subnet for networks not found | 
 **vrfgroup** | **optional.String**| name of vrf group for discovered subnets | 
 **subnetCategory** | **optional.String**| name of subnet category for discovered subnets | 
 **overwriteSubnetCategories** | **optional.String**| If a subnet is discovered that exists and already has a subnet category, the category will be overwritten and all child subnets of the discovered subnet will also get the new category. | 
 **remoteCollectorId** | **optional.Int32**| ID of the remote collector to use for this job | 
 **tags** | **optional.String**| comma separated list of tags | 
 **tagsRemove** | **optional.String**| comma separated list of tags to remove | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryPowerDisc**
> interface{} PostAutoDiscoveryPowerDisc(ctx, name, appliance, optional)
Add or update power autodiscovery job

Add or update power autodiscovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
  **appliance** | **string**| name of the monitoring appliance | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryPowerDiscOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryPowerDiscOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ipaddress** | **optional.String**| IP address. Required if new | 
 **endIpAddress** | **optional.String**| End IP address | 
 **snmpString** | **optional.String**| required, if new | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 
 **stripDomainName** | **optional.String**|  | 
 **clearExistingSchedule** | **optional.String**|  | 
 **vrfgroup** | **optional.String**| name of vrf group for discovered subnets (added in v10.4.0) | 
 **objectCategory** | **optional.String**| name of subnet category for discovered subnets | 
 **category** | **optional.String**| subnet category to assign to addresses of discovered units | 
 **groups** | **optional.String**| name of one or more groups separated by commas | 
 **snmpPort** | **optional.Int32**| snmp port (integer only) (added in v10.4.0) | 
 **snmpVersion** | **optional.String**|  | [default to v2c]
 **snmpStringId** | **optional.String**| The id of the password for the community string | 
 **snmpv3AuthPasswordId** | **optional.String**| The id of the password for the auth password | 
 **snmpv3PrivacyProtocolPasswordId** | **optional.String**| The id of the password for the privacy protocol password | 
 **snmpv3User** | **optional.String**| name of snmp v3 user (added in v10.4.0) | 
 **snmpv3AuthMode** | **optional.String**|  | 
 **snmpv3AuthProtocol** | **optional.String**|  | 
 **snmpv3AuthPassword** | **optional.String**| password (added in v10.4.0) | 
 **snmpv3PrivacyProtocol** | **optional.String**|  | 
 **snmpv3PrivacyProtocolPassword** | **optional.String**| password (added in v10.4.0) | 
 **snmpv3Context** | **optional.String**|  | 
 **hostnamePrecedence** | **optional.String**|  | 
 **serviceLevel** | **optional.String**| Must already exist | 
 **nameToUseForNewlyDiscoveredModule** | **optional.String**| name to use | 
 **toggleInServiceOnModulePowerState** | **optional.String**| yes or no | 
 **moduleNotFound** | **optional.String**| action to take on module not found. One of: Remove Host Association, Change Service Level, Delete Module (default is no action) | 
 **nameToUseForNewlyDiscoveredPdu** | **optional.String**| one of: Name discovered by SNMP, Name plus serial number, Name plus serial number plus IP | 
 **pollingInterval** | **optional.String**| polling interval in minutes (integer) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoverySnmpDisc**
> PostAutoDiscoverySnmpDisc(ctx, name, optional)
Add or update other SNMP autodiscovery job

Add or update other SNMP autodiscovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoverySnmpDiscOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoverySnmpDiscOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipaddress** | **optional.String**| IP address. Required if new | 
 **endIpAddress** | **optional.String**| End IP address | 
 **snmpString** | **optional.String**| required, if new | 
 **remoteCollectorId** | **optional.String**| D42 ID of the remote collector to use for this job. | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 
 **stripDomainName** | **optional.String**|  | 
 **clearExistingSchedule** | **optional.String**|  | 
 **groups** | **optional.String**| name of one or more groups separated by commas | 
 **snmpPort** | **optional.Int32**| snmp port (integer only) (added in v10.4.0) | 
 **snmpVersion** | **optional.String**|  | [default to v2c]
 **snmpStringId** | **optional.String**| The id of the password for the community string | 
 **snmpv3AuthPasswordId** | **optional.String**| The id of the password for the auth password | 
 **snmpv3PrivacyProtocolPasswordId** | **optional.String**| The id of the password for the privacy protocol password | 
 **snmpv3User** | **optional.String**| name of snmp v3 user (added in v10.4.0) | 
 **snmpv3AuthMode** | **optional.String**|  | 
 **snmpv3AuthProtocol** | **optional.String**|  | 
 **snmpv3AuthPassword** | **optional.String**| password (added in v10.4.0) | 
 **snmpv3PrivacyProtocol** | **optional.String**|  | 
 **snmpv3PrivacyProtocolPassword** | **optional.String**| password (added in v10.4.0) | 
 **snmpv3Context** | **optional.String**|  | 
 **debugLevel** | **optional.String**|  | 
 **serviceLevel** | **optional.String**| Must already exist | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryUcs**
> interface{} PostAutoDiscoveryUcs(ctx, name, platform, port, username, password, hostnameToUse, optional)
Add or update UCS/ACI/Load Balancer Auto Discovery

Add or update a UCS/ACI/Load Balancer auto-discovery job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
  **platform** | **string**| Required | 
  **port** | **string**| Default is 443 - use this default if you are not sure. | 
  **username** | **string**|  | 
  **password** | **string**|  | 
  **hostnameToUse** | **string**| Required - relevant for new devices only. If a device with same serial number already exists, name is ignored | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryUcsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryUcsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **remoteCollectorId** | **optional.String**|  | 
 **server** | **optional.String**| FQDN or IP of the server(s) or CIDR or ranges | 
 **useSsl** | **optional.String**| Default is Yes - use this default if you are not sure. | 
 **excludeServers** | **optional.String**| Servers to ignore | 
 **hostnamePrecedence** | **optional.String**| Discovered name for the device is given precedence over existing name in the system. | 
 **toggleServiceLevelOnPowerState** | **optional.String**| Toggle the service level of a device based on power state. | 
 **decommissionedServiceLevelId** | **optional.String**| Service level for decommissioned devices. | 
 **vrfGroupId** | **optional.String**| VRF Group for discovered devices. | 
 **objectCategory** | **optional.String**| Existing object category will not be overwritten. | 
 **overwriteObjectCategories** | **optional.String**|  | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 
 **clearExistingSchedule** | **optional.String**|  | 
 **debugLevel** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoveryVserver**
> interface{} PostAutoDiscoveryVserver(ctx, name, optional)
Add or update vServer autodiscovery job

Add or update vServer autodiscovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoveryVserverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoveryVserverOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | **optional.String**| IP or FQDN of target server | 
 **platform** | **optional.String**|  | 
 **endIpAddress** | **optional.String**| End IP address | 
 **useFtp** | **optional.String**| Use FTP instead of Telnet for discovery. | 
 **clearExistingSchedule** | **optional.String**|  | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 
 **passwordIds** | **optional.String**| One or more D42 password ID numbers separated by commas | 
 **groups** | **optional.String**| name of one or more groups separated by commas | 
 **objectCategory** | **optional.String**| category of discovered vservers and vms | 
 **debugLevel** | **optional.String**|  | 
 **serviceLevel** | **optional.String**| Must already exist | 
 **urlPrefix** | **optional.String**| http or https | 
 **port** | **optional.String**| specify port | 
 **urlSuffix** | **optional.String**| append suffix to discovery url | 
 **username** | **optional.String**| username to use in discovery | 
 **password** | **optional.String**| password to use in discovery | 
 **stripDomainSuffix** | **optional.String**| Strip domain suffix if discovered on VMs or hypervisor | 
 **prependVmhostname** | **optional.String**|  | 
 **ignorePoweredOff** | **optional.String**| Ignore powered off vms | 
 **discoverVms** | **optional.String**| discover VMs or strictly hypervisors | 
 **vmNameToUse** | **optional.String**| “found from vm tools” or “as named on vserver” | 
 **addMultipleVmNamesAsAlias** | **optional.String**| add any additional names found as device alias | 
 **toggleServiceLevelOnVmPowerState** | **optional.String**|  | 
 **getGuestOsInfo** | **optional.String**|  | 
 **ignoreGuestUuid** | **optional.String**|  | 
 **vmNotFound** | **optional.String**| Choose how to handle VM not found in discovery | 
 **trackVmNameChange** | **optional.String**|  | 
 **hostnamePrecedence** | **optional.String**|  | 
 **hostidlist** | **optional.String**|  | 
 **vmAddDisk** | **optional.String**|  | 
 **hostAllowDuplicateSerials** | **optional.String**|  | 
 **ignoreHostSerial** | **optional.String**|  | 
 **ignoreHostUuid** | **optional.String**|  | 
 **overwriteObjectCategories** | **optional.String**|  | 
 **remoteCollectorId** | **optional.String**|  | 
 **useServiceAccount** | **optional.String**|  | 
 **ignoreIpv6** | **optional.String**|  | 
 **ignoreVirtSubtype** | **optional.String**|  | 
 **deviceNameFormat** | **optional.String**|  | 
 **discoverParts** | **optional.String**|  | 
 **captureHostsFile** | **optional.String**|  | 
 **discoverSoftware** | **optional.String**|  | 
 **initialSoftwareType** | **optional.String**|  | 
 **discoverLinesOfCode** | **optional.String**| If enabled, please set the max timeout for the discovery to greater than 5 minutes. | 
 **discoverServices** | **optional.String**|  | 
 **discoverApplications** | **optional.String**|  | 
 **storeConfigFiles** | **optional.String**|  | 
 **discoverCloudid** | **optional.String**|  | 
 **customerId** | **optional.String**|  | 
 **vrfGroupId** | **optional.String**| ID of the VRF group | 
 **serviceLevelDeviceId** | **optional.String**|  | 
 **sudoRetry** | **optional.String**|  | 
 **alternateSudo** | **optional.String**|  | 
 **useDomainServer** | **optional.String**|  | 
 **ldapServer** | **optional.String**|  | 
 **useFqdn** | **optional.String**|  | 
 **ldapFilterType** | **optional.String**|  | 
 **ldapUnpwdId** | **optional.String**|  | 
 **alternateSudoPasswordId** | **optional.String**|  | 
 **pollingInterval** | **optional.String**| polling interval in minutes (integer) | 
 **enableResourcesMonitoring** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoDiscoverynmap**
> interface{} PostAutoDiscoverynmap(ctx, name, ipStart, ipEnd, remoteCollectorId, optional)
Add Nmap auto-discovery job

Add Nmap auto-discovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name of the job | 
  **ipStart** | **string**| starting IP address | 
  **ipEnd** | **string**| ending IP address, use same as start for single address | 
  **remoteCollectorId** | **string**| D42 ID of the remote collector to use for this job. | 
 **optional** | ***AutoDiscoveryApiPostAutoDiscoverynmapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPostAutoDiscoverynmapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **nameserver** | **optional.String**| IP/FQDA of name server | 
 **osDetection** | **optional.String**| Select yes to detect operating system. | 
 **servicesDetection** | **optional.String**| Select yes to detect services. | 
 **objectCategory** | **optional.String**| Name of subnet category for discovered subnets | 
 **newObjectCategory** | **optional.String**| Update an existing object category. | 
 **overwriteDeviceCategories** | **optional.String**| Overwrite the object categories for existing discovered devices.&lt;br&gt;&lt;br&gt;Note - this is equivalent to the &#39;Overwrite existing object categories&#39; parameter shown in the D42 Nmap UI. | 
 **vrfgroup** | **optional.String**| Name of vrf group for discovered subnets | 
 **vrfgroupId** | **optional.String**| D42 ID of group for discovered subnets | 
 **tags** | **optional.String**| Comma separated list of tags | 
 **tagsRemove** | **optional.String**| Comma separated list of tags to remove | 
 **scheduleTime** | **optional.String**| Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/). | 
 **scheduleDays** | **optional.String**| Comma separated days of week, where Monday &#x3D; 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/). | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryBladeDisc**
> interface{} PutAutoDiscoveryBladeDisc(ctx, run, optional)
Execute blade autodiscovery job

Execute blade autodiscovery job. Required parameters: name or job_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoveryBladeDiscOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoveryBladeDiscOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id | 
 **jobId** | **optional.String**| D42 ID for the job - required if no name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryCertificate**
> interface{} PutAutoDiscoveryCertificate(ctx, run, optional)
Execute Certificate job

Execute Certificate Auto Discovery Job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoveryCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoveryCertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id | 
 **jobId** | **optional.String**| D42 ID for the job - required if no name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryCloudaccount**
> interface{} PutAutoDiscoveryCloudaccount(ctx, run, optional)
Execute Cloud job

Execute Cloud job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoveryCloudaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoveryCloudaccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id | 
 **jobId** | **optional.String**| D42 ID for the job - required if no name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryDns**
> interface{} PutAutoDiscoveryDns(ctx, jobId, run)
Execute DNS job

Execute DNS Auto Discovery Job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| D42 ID for the job | 
  **run** | **string**| yes to start | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryIpmi**
> interface{} PutAutoDiscoveryIpmi(ctx, jobId, run)
Execute snmp autodiscovery job

Execute snmp autodiscovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| D42 ID for the job | 
  **run** | **string**| yes to start | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryNetworks**
> interface{} PutAutoDiscoveryNetworks(ctx, run, optional)
Execute network job

Execute network job. Required parameters: name or job_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoveryNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoveryNetworksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id | 
 **jobId** | **optional.String**| D42 ID for the job - required if no name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryPingsweep**
> interface{} PutAutoDiscoveryPingsweep(ctx, run, optional)
Execute pingsweep job

Execute pingsweep job. Required parameters: name or job_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoveryPingsweepOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoveryPingsweepOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id | 
 **jobId** | **optional.String**| D42 ID for the job - required if no name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryPowerDisc**
> interface{} PutAutoDiscoveryPowerDisc(ctx, run, optional)
Execute power autodiscovery job

Execute power autodiscovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoveryPowerDiscOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoveryPowerDiscOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id | 
 **jobId** | **optional.String**| D42 ID for the job - required if no name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoverySnmpDisc**
> interface{} PutAutoDiscoverySnmpDisc(ctx, run, optional)
Execute snmp autodiscovery job

Execute snmp autodiscovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoverySnmpDiscOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoverySnmpDiscOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id | 
 **jobId** | **optional.String**| D42 ID for the job - required if no name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryUcs**
> interface{} PutAutoDiscoveryUcs(ctx, jobId, run)
Execute UCS/ACI/Load Balancer autodiscovery job

Execute UCS/ACI/Load Balancer autodiscovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| D42 ID for the job | 
  **run** | **string**| yes to start | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoveryVserver**
> interface{} PutAutoDiscoveryVserver(ctx, run, optional)
Execute vserver job

Execute VServer job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoveryVserverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoveryVserverOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id | 
 **jobId** | **optional.String**| D42 ID for the job - required if no name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoDiscoverynmap**
> interface{} PutAutoDiscoverynmap(ctx, run, optional)
Execute Nmap job

Execute nmap job. Required parameters: name or job_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **run** | **string**| yes to start | 
 **optional** | ***AutoDiscoveryApiPutAutoDiscoverynmapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoDiscoveryApiPutAutoDiscoverynmapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the job - required if no job_id supplied | 
 **jobId** | **optional.String**| D42 ID of the job - required if no name supplied | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

