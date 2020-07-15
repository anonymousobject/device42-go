# \IPAMApi

All URIs are relative to *https://swaggerdemo.device42.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIPAMDnsRecords**](IPAMApi.md#DeleteIPAMDnsRecords) | **Delete** /api/1.0/dns/records/{id}/ | Delete DNS Records
[**DeleteIPAMDnsZones**](IPAMApi.md#DeleteIPAMDnsZones) | **Delete** /api/1.0/dns/zones/{id}/ | Delete DNS Zones
[**DeleteIPAMIpnat**](IPAMApi.md#DeleteIPAMIpnat) | **Delete** /api/1.0/ipnat/{id} | This API is used to delete the IP NAT.
[**DeleteIPAMIps**](IPAMApi.md#DeleteIPAMIps) | **Delete** /api/1.0/ips/{id}/ | Delete IP Address
[**DeleteIPAMMacsId**](IPAMApi.md#DeleteIPAMMacsId) | **Delete** /api/1.0/macs/{id}/ | Delete Mac Address
[**DeleteIPAMSwitchports**](IPAMApi.md#DeleteIPAMSwitchports) | **Delete** /api/1.0/switchports/{id}/ | Delete Switchport
[**DeleteIPAMsubnetcategory**](IPAMApi.md#DeleteIPAMsubnetcategory) | **Delete** /api/1.0/subnet_category/{id}/ | Delete Subnet Category
[**DeleteIPAMsubnets**](IPAMApi.md#DeleteIPAMsubnets) | **Delete** /api/1.0/subnets/{subnet_id}/ | Delete Subnet
[**DeleteIPAMvlans**](IPAMApi.md#DeleteIPAMvlans) | **Delete** /api/1.0/vlans/{id}/ | Delete VLAN
[**DeleteIPAMvrfgroupId**](IPAMApi.md#DeleteIPAMvrfgroupId) | **Delete** /api/1.0/vrfgroup/{id}/ | Delete VRF Group
[**GetIPAMCustomerId**](IPAMApi.md#GetIPAMCustomerId) | **Get** /api/1.0/subnets/customer_id/{customer_id}/ | GET method retrieves all Subnets by customer ID #.
[**GetIPAMDnsRecords**](IPAMApi.md#GetIPAMDnsRecords) | **Get** /api/1.0/dns/records/ | GET method retrieves all DNS Records. (Added in v6.5.0)
[**GetIPAMIpnat**](IPAMApi.md#GetIPAMIpnat) | **Get** /api/1.0/ipnat/ | GET method retrieves all Switch Templates.
[**GetIPAMIps**](IPAMApi.md#GetIPAMIps) | **Get** /api/1.0/ips/ | Get all IPs
[**GetIPAMIpsSubnetId**](IPAMApi.md#GetIPAMIpsSubnetId) | **Get** /api/1.0/ips/subnet_id/{subnet_id}/ | GET method retrieves all IPs by Subnet ID.
[**GetIPAMMacs**](IPAMApi.md#GetIPAMMacs) | **Get** /api/1.0/macs/ | GET method retrieves all MAC Addresses.
[**GetIPAMMacsId**](IPAMApi.md#GetIPAMMacsId) | **Get** /api/1.0/macs/{id}/ | GET method retrieves MAC address by specific MAC ID.
[**GetIPAMSearch**](IPAMApi.md#GetIPAMSearch) | **Get** /api/1.0/search/ | Find a specific IP Address.
[**GetIPAMSubnetId**](IPAMApi.md#GetIPAMSubnetId) | **Get** /api/1.0/subnets/{subnet_id}/ | Get Subnet by Subnet ID
[**GetIPAMSuggestIp**](IPAMApi.md#GetIPAMSuggestIp) | **Get** /api/1.0/suggest_ip/ | Suggest next available IP Address. Can be IPv4 or IPv6.
[**GetIPAMSuggestSubnetId**](IPAMApi.md#GetIPAMSuggestSubnetId) | **Get** /api/1.0/suggest_subnet/{id}/ | Suggest next available subnet.
[**GetIPAMSwitchTemplates**](IPAMApi.md#GetIPAMSwitchTemplates) | **Get** /api/1.0/switch_templates/ | GET method retrieves all Switch Templates.
[**GetIPAMSwitchports**](IPAMApi.md#GetIPAMSwitchports) | **Get** /api/1.0/switchports | GET method retrieves all Switch Ports.
[**GetIPAMTapPorts**](IPAMApi.md#GetIPAMTapPorts) | **Get** /api/1.0/tap_ports/ | GET method retrieves all TAP Ports.
[**GetIPAMsubnetcategory**](IPAMApi.md#GetIPAMsubnetcategory) | **Get** /api/1.0/subnet_category/ | GET method retrieves all Subnet Categories.
[**GetIPAMsubnets**](IPAMApi.md#GetIPAMsubnets) | **Get** /api/1.0/subnets/ | GET method retrieves all Subnets.
[**GetIPAMvlans**](IPAMApi.md#GetIPAMvlans) | **Get** /api/1.0/vlans/ | GET method retrieves all VLANs.
[**GetIPAMvlansId**](IPAMApi.md#GetIPAMvlansId) | **Get** /api/1.0/vlans/{id}/ | Get VLAN by ID
[**GetIPAMvrfgroup**](IPAMApi.md#GetIPAMvrfgroup) | **Get** /api/1.0/vrfgroup/ | GET method retrieves all VRF Groups.
[**PostIPAMDnsRecords**](IPAMApi.md#PostIPAMDnsRecords) | **Post** /api/1.0/dns/records/ | Create DNS Records.
[**PostIPAMDnsZones**](IPAMApi.md#PostIPAMDnsZones) | **Post** /api/1.0/dns/zones/ | Create DNS Zones.
[**PostIPAMIpnat**](IPAMApi.md#PostIPAMIpnat) | **Post** /api/1.0/ipnat/ | Create IP NAT.
[**PostIPAMIps**](IPAMApi.md#PostIPAMIps) | **Post** /api/1.0/ips/ | Create / Update IP Addresses.
[**PostIPAMMacs**](IPAMApi.md#PostIPAMMacs) | **Post** /api/1.0/macs/ | Create / Update MAC Addresses.
[**PostIPAMSubnetsCreateChild**](IPAMApi.md#PostIPAMSubnetsCreateChild) | **Post** /api/1.0/subnets/create_child/ | Create child Subnet.
[**PostIPAMSwitches**](IPAMApi.md#PostIPAMSwitches) | **Post** /api/1.0/switches/ | Add Switch/Switch Ports w/ Templates.
[**PostIPAMSwitchports**](IPAMApi.md#PostIPAMSwitchports) | **Post** /api/1.0/switchports | Create / Update Switch Port by Port.
[**PostIPAMTapPorts**](IPAMApi.md#PostIPAMTapPorts) | **Post** /api/1.0/tap_ports/ | Create / Update TAP Ports
[**PostIPAMsubnetcategory**](IPAMApi.md#PostIPAMsubnetcategory) | **Post** /api/1.0/subnet_category/ | 
[**PostIPAMsubnets**](IPAMApi.md#PostIPAMsubnets) | **Post** /api/1.0/subnets/ | Create / Update Subnets.
[**PostIPAMvlans**](IPAMApi.md#PostIPAMvlans) | **Post** /api/1.0/vlans/ | Create VLANS.
[**PostIPAMvlansSmartMergeAll**](IPAMApi.md#PostIPAMvlansSmartMergeAll) | **Post** /api/1.0/vlans/smart_merge_all/ | Update VLANS by ID in url.
[**PostIPAMvrfgroup**](IPAMApi.md#PostIPAMvrfgroup) | **Post** /api/1.0/vrfgroup/ | Create/Update VRF Group.
[**PutIPAMCustomFIipAddress**](IPAMApi.md#PutIPAMCustomFIipAddress) | **Put** /api/1.0/custom_fields/ip_address/ | Create/update custom fields for subnets.
[**PutIPAMCustomFIsubnet**](IPAMApi.md#PutIPAMCustomFIsubnet) | **Put** /api/1.0/custom_fields/subnet/ | Create/update custom fields for subnets.
[**PutIPAMCustomFIswitchport**](IPAMApi.md#PutIPAMCustomFIswitchport) | **Put** /api/1.0/custom_fields/switchport/ | Create / Update Switch Port Custom Fields
[**PutIPAMCustomFIvrfgroup**](IPAMApi.md#PutIPAMCustomFIvrfgroup) | **Put** /api/1.0/custom_fields/vrfgroup/ | VRF Group Custom Fields
[**PutIPAMIpnat**](IPAMApi.md#PutIPAMIpnat) | **Put** /api/1.0/ipnat/ | Update IP NAT (Added in v7.0.0)
[**PutIPAMsubnetcategory**](IPAMApi.md#PutIPAMsubnetcategory) | **Put** /api/1.0/subnet_category/ | Create Subnet Category.
[**PutIPAMsubnets**](IPAMApi.md#PutIPAMsubnets) | **Put** /api/1.0/subnets/ | Create / Update Subnets.
[**PutIPAMvlans**](IPAMApi.md#PutIPAMvlans) | **Put** /api/1.0/vlans/{id}/ | Update VLANS by ID in url.
[**PutIPAMvrfgroup**](IPAMApi.md#PutIPAMvrfgroup) | **Put** /api/1.0/vrfgroup/ | Update VRF Group


# **DeleteIPAMDnsRecords**
> interface{} DeleteIPAMDnsRecords(ctx, id)
Delete DNS Records

This API is used to delete the DNS record with the DNS Record id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| DNS Record id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAMDnsZones**
> interface{} DeleteIPAMDnsZones(ctx, id)
Delete DNS Zones

This API is used to delete the DNS zone with the DNS zone id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| DNS zone id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAMIpnat**
> interface{} DeleteIPAMIpnat(ctx, id)
This API is used to delete the IP NAT.

Delete IP NAT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Name of Id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAMIps**
> interface{} DeleteIPAMIps(ctx, id)
Delete IP Address

This API is used to delete an IP Address with the IP Address id supplied as the required argument.

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

# **DeleteIPAMMacsId**
> interface{} DeleteIPAMMacsId(ctx, id)
Delete Mac Address

This API is used to delete the mac address with the mac address id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Mac address id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAMSwitchports**
> interface{} DeleteIPAMSwitchports(ctx, id)
Delete Switchport

This API is used to delete the switch port with the switch port id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Switch port id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAMsubnetcategory**
> interface{} DeleteIPAMsubnetcategory(ctx, id)
Delete Subnet Category

This API is used to delete the subnet category with the subnet category id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Subnet category id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAMsubnets**
> interface{} DeleteIPAMsubnets(ctx, subnetId)
Delete Subnet

This API is used to delete the subnet with the subnet id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subnetId** | **int32**| subnet id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAMvlans**
> interface{} DeleteIPAMvlans(ctx, id)
Delete VLAN

This API is used to delete the vlan with the vlan id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| vlan id (This is Device42 ID of the VLAN) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAMvrfgroupId**
> interface{} DeleteIPAMvrfgroupId(ctx, id)
Delete VRF Group

This API is used to delete the vrf group with the vrf group id supplied as the required argument.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VRF group | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMCustomerId**
> []InlineResponse2007 GetIPAMCustomerId(ctx, customerId)
GET method retrieves all Subnets by customer ID #.

Get all Subnets by Customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int32**| Customer Id | 

### Return type

[**[]InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMDnsRecords**
> interface{} GetIPAMDnsRecords(ctx, optional)
GET method retrieves all DNS Records. (Added in v6.5.0)

Get DNS Records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiGetIPAMDnsRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMDnsRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **optional.String**| name of the domain | 
 **type_** | **optional.String**| type of the record. e.g. AAAA | 
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **nameserver** | **optional.String**| Nameserver | 
 **content** | **optional.String**| Content (e.g. IP address for type A) | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 
 **tagsAnd** | **optional.String**| filter by all the tags, separated by comma. (This is an AND filter and all tags have to match for filter, this was added in v6.3.1) | 
 **dnsZone** | **optional.String**|  | 
 **ttl** | **optional.String**| TTL value. | 
 **changeDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMIpnat**
> interface{} GetIPAMIpnat(ctx, )
GET method retrieves all Switch Templates.

Get all IP NAT Records

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMIps**
> interface{} GetIPAMIps(ctx, optional)
Get all IPs

GET method retrieves all IPs. (Added in v5.9.3) By default the limit is maximum 1000 IPs per call. Use total_count and offset if over 1000 IPs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiGetIPAMIpsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMIpsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.String**| start with this device (e.g. limit&#x3D;100&amp;offset&#x3D;50 means start with the 50th device and return the next 100 devices) | 
 **limit** | **optional.String**| return this number of devices | 
 **subnetId** | **optional.String**| ID of the subnet (added in v7.2.0) | 
 **subnet** | **optional.String**| name of the subnet | 
 **device** | **optional.String**| Device name | 
 **mac** | **optional.String**| mac address | 
 **available** | **optional.String**| could be yes or no | 
 **type_** | **optional.String**|  | 
 **lastUpdatedLt** | **optional.String**| last updated less than date YYYY-MM-DD format | 
 **lastUpdatedGt** | **optional.String**| last updated greater than date YYYY-MM-DD format | 
 **ip** | **optional.String**| ip address (added in v6.3.0) | 
 **firstAddedLt** | **optional.String**| first added less than date YYYY-MM-DD format | 
 **firstAddedGt** | **optional.String**| first added greater date YYYY-MM-DD format | 
 **ipId** | **optional.String**| D42 ID of the IP | 
 **label** | **optional.String**|  | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 
 **tagsAnd** | **optional.String**| filter by all the tags, separated by comma. (This is an AND filter and all tags have to match for filter, this was added in v6.3.1) | 
 **customFieldsAnd** | **optional.String**| filter by custom fields, and filter, format of key1:value1,key2:value2 | 
 **customFieldsOr** | **optional.String**| filter by custom fields, or filter, format of key1:value1,key2:value2 | 
 **totalCount** | **optional.String**| Count of IPs returned (use with offset as max results are limited to 1000) | 
 **ips** | **optional.String**| Details for all the IPs | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMIpsSubnetId**
> interface{} GetIPAMIpsSubnetId(ctx, subnetId)
GET method retrieves all IPs by Subnet ID.

Get all IPs in a Subnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subnetId** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMMacs**
> interface{} GetIPAMMacs(ctx, optional)
GET method retrieves all MAC Addresses.

Get all MAC Addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiGetIPAMMacsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMMacsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mac** | **optional.String**| mac address | 
 **device** | **optional.String**| Device name | 
 **deviceId** | **optional.String**| Device ID | 
 **lastUpdatedLt** | **optional.String**| last updated less than date YYYY-MM-DD format | 
 **lastUpdatedGt** | **optional.String**| last updated greater than date YYYY-MM-DD format | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMMacsId**
> interface{} GetIPAMMacsId(ctx, id)
GET method retrieves MAC address by specific MAC ID.

Get MAC Address with ID

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

# **GetIPAMSearch**
> interface{} GetIPAMSearch(ctx, query, string_)
Find a specific IP Address.

Find Specific IP Address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| enter “ip” | 
  **string_** | **string**| the IP address to search for | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMSubnetId**
> IpaMsubnets GetIPAMSubnetId(ctx, subnetId)
Get Subnet by Subnet ID

GET method retrieves the subnet with the specified subnet id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subnetId** | **int32**| Subnet Id | 

### Return type

[**IpaMsubnets**](IPAMsubnets.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMSuggestIp**
> interface{} GetIPAMSuggestIp(ctx, maskBits, optional)
Suggest next available IP Address. Can be IPv4 or IPv6.

Suggest next available IP Address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maskBits** | **string**| mask bits (added in v7.2.0) | 
 **optional** | ***IPAMApiGetIPAMSuggestIpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMSuggestIpOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subnetId** | **optional.String**| ID of the subnet (added in v7.2.0) | 
 **subnet** | **optional.String**| name of the subnet | 
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **vrfGroupId** | **optional.String**| ID of the VRF group | 
 **vrfGroup** | **optional.String**| VRF group name | 
 **reserveIp** | **optional.String**| If value of yes is passed, the suggested IP is reserved. Return value also adds reserved as yes or no. (added in v7.2.0) | 
 **number** | **optional.String**| vlan number | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMSuggestSubnetId**
> interface{} GetIPAMSuggestSubnetId(ctx, id, maskBits, optional)
Suggest next available subnet.

Suggest next available Subnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the parent subnet | 
  **maskBits** | **string**| mask bits (added in v7.2.0) | 
 **optional** | ***IPAMApiGetIPAMSuggestSubnetIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMSuggestSubnetIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subnet** | **optional.String**| name of the subnet | 
 **subnetId** | **optional.String**| ID of the subnet (added in v7.2.0) | 
 **vrfGroupId** | **optional.String**| ID of the VRF group | 
 **vrfGroup** | **optional.String**| VRF group name | 
 **name** | **optional.String**| filter by name (Added in v6.0.0) | 
 **ifParentAssigned** | **optional.String**| yes or no | 
 **ifParentAllocated** | **optional.String**| yes or no | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMSwitchTemplates**
> []InlineResponse20010 GetIPAMSwitchTemplates(ctx, )
GET method retrieves all Switch Templates.

Get all Switch Templates

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMSwitchports**
> interface{} GetIPAMSwitchports(ctx, optional)
GET method retrieves all Switch Ports.

Get all Switch Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiGetIPAMSwitchportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMSwitchportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **switchId** | **optional.String**| ID of the network device port is on | 
 **switch2Id** | **optional.String**| ID of the 2nd network device port is on | 
 **lastUpdatedLt** | **optional.String**| last updated less than date YYYY-MM-DD format | 
 **lastUpdatedGt** | **optional.String**| last updated greater than date YYYY-MM-DD format | 
 **firstAddedLt** | **optional.String**| first added less than date YYYY-MM-DD format | 
 **firstAddedGt** | **optional.String**| first added greater date YYYY-MM-DD format | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 
 **tagsAnd** | **optional.String**| filter by all the tags, separated by comma. (This is an AND filter and all tags have to match for filter, this was added in v6.3.1) | 
 **includeCols** | **optional.String**| do not return all columns just the ones specified. For example, ?include_cols&#x3D;name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMTapPorts**
> interface{} GetIPAMTapPorts(ctx, optional)
GET method retrieves all TAP Ports.

Get Tap Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiGetIPAMTapPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMTapPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| id of the tap port | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMsubnetcategory**
> []InlineResponse2005 GetIPAMsubnetcategory(ctx, )
GET method retrieves all Subnet Categories.

Get all Subnet Categories

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMsubnets**
> InlineResponse2006 GetIPAMsubnets(ctx, optional)
GET method retrieves all Subnets.

Get all Subnets - Filter parameters are below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiGetIPAMsubnetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMsubnetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| name of the subnet | 
 **vrfGroupId** | **optional.String**| ID of the VRF group | 
 **vrfGroup** | **optional.String**| VRF group name | 
 **parentSubnetId** | **optional.String**| ID of the parent subnet | 
 **parentSubnet** | **optional.String**| parent subnet name | 
 **customerId** | **optional.String**| ID of the customer (Added in v6.3.0) | 
 **customer** | **optional.String**| filter by customer name | 
 **subnetId** | **optional.String**| ID of the subnet (added in v7.2.0) | 
 **maskBits** | **optional.String**| mask bits | 
 **maskBitsLt** | **optional.String**| less than mask bits (added in v7.2.0) | 
 **maskBitsGt** | **optional.String**| greater than mask bits (added in v7.2.0) | 
 **description** | **optional.String**| matching description (added in v7.2.0) | 
 **rangeBegin** | **optional.String**| Range Begin (added in v7.2.0) | 
 **rangeEnd** | **optional.String**| Range End (added in v7.2.0) | 
 **gateway** | **optional.String**| Gateway (added in v7.2.0) | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 
 **tagsAnd** | **optional.String**| filter by all the tags, separated by comma. (This is an AND filter and all tags have to match for filter, this was added in v6.3.1) | 
 **customFieldsAnd** | **optional.String**| filter by custom fields, and filter, format of key1:value1,key2:value2 | 
 **customFieldsOr** | **optional.String**| filter by custom fields, or filter, format of key1:value1,key2:value2 | 
 **serviceLevel** | **optional.String**| filter by service level name | 
 **category** | **optional.String**| name of the category | 
 **categoryId** | **optional.String**| ID of the category | 
 **vlanId** | **optional.String**| ID of the vlan | 
 **network** | **optional.String**| Optional | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMvlans**
> InlineResponse2009 GetIPAMvlans(ctx, optional)
GET method retrieves all VLANs.

Get all VLANs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiGetIPAMvlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiGetIPAMvlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanId** | **optional.String**| ID of the vlan | 
 **number** | **optional.String**| vlan number | 
 **tags** | **optional.String**| filter by tags. comma separated for multiple tags (This is an OR filter, gets all the devices for all comma separated tags) | 
 **tagsAnd** | **optional.String**| filter by all the tags, separated by comma. (This is an AND filter and all tags have to match for filter, this was added in v6.3.1) | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMvlansId**
> IpaMvlans GetIPAMvlansId(ctx, id)
Get VLAN by ID

GET method retrieves VLAN by specific VLAN ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**IpaMvlans**](IPAMvlans.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAMvrfgroup**
> []InlineResponse2004 GetIPAMvrfgroup(ctx, )
GET method retrieves all VRF Groups.

Get all VRF Groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMDnsRecords**
> interface{} PostIPAMDnsRecords(ctx, domain, type_, optional)
Create DNS Records.

Create / Update DNS Records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| name of the domain | 
  **type_** | **string**| Type of record | 
 **optional** | ***IPAMApiPostIPAMDnsRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMDnsRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nameserver** | **optional.String**| If you have overlapping domain names, this is required to differentiate between the two | 
 **name** | **optional.String**| Record value (e.g. www) use @ for blank. | 
 **content** | **optional.String**| Content (e.g. IP address for type A) | 
 **prio** | **optional.String**| Priority for MX record. | 
 **ttl** | **optional.String**| TTL Value | 
 **tags** | **optional.String**| set tags for record | 
 **tagsRemove** | **optional.String**| remove tags for record | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMDnsZones**
> interface{} PostIPAMDnsZones(ctx, name, nameserver, optional)
Create DNS Zones.

Create / Update DNS Zones

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of DNS zone | 
  **nameserver** | **string**| IP Address/hostname of name server | 
 **optional** | ***IPAMApiPostIPAMDnsZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMDnsZonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **notes** | **optional.String**| Any additional notes | 
 **vrfGroup** | **optional.String**| VRF group name | 
 **vrfGroupId** | **optional.String**| ID of the VRF group | 
 **tags** | **optional.String**| Tags for grouping zone entries | 
 **tagsRemove** | **optional.String**| remove tags from grouping zone entries | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMIpnat**
> interface{} PostIPAMIpnat(ctx, name, ipAddressFrom, ipAddressTo, optional)
Create IP NAT.

Create IP NAT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the NAT IP entry | 
  **ipAddressFrom** | **string**| The external IP address | 
  **ipAddressTo** | **string**| The internal IP address | 
 **optional** | ***IPAMApiPostIPAMIpnatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMIpnatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ipAddressFromEnd** | **optional.String**| used for mapping a range of addresses rather than a single IP | 
 **sourcePortStart** | **optional.String**| Starting port for source IP range | 
 **sourcePortEnd** | **optional.String**| used for mapping a range of ports rather than a single port | 
 **ipAddressToEnd** | **optional.String**| used for mapping a range of addresses rather than a single IP | 
 **targetPortStart** | **optional.String**| Starting port for target IP range | 
 **targetPortEnd** | **optional.String**| used for mapping a range of ports rather than a single port | 
 **twoWayRelation** | **optional.String**| true if the internal IP addressed is masked with the external IP address for outbound traffic in addition to inbound traffic. Default is false. | 
 **protocol** | **optional.String**| the transport protocol, ie TCP | 
 **notes** | **optional.String**| Any additional notes | 
 **vrfGroupFrom** | **optional.String**| Originating VRF Group Name | 
 **vrfGroupIdFrom** | **optional.String**| Originating VRF Group ID | 
 **vrfGroupTo** | **optional.String**| Destination VRF Group Name | 
 **vrfGroupIdTo** | **optional.String**| Destination VRF Group ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMIps**
> interface{} PostIPAMIps(ctx, ipaddress, optional)
Create / Update IP Addresses.

Create / Update IP Addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipaddress** | **string**| If a matching IP address is found, it will update the first matched IP address(unless you specify a vrf_group or vrf_group_id, then it matches or adds IP to that VRF group) | 
 **optional** | ***IPAMApiPostIPAMIpsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMIpsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **optional.String**| label for the interface (tag still works for backward compatibility) | 
 **subnet** | **optional.String**| name of the subnet you want to add the IP to. Must be unique. The reason it must be unique is to handle overlapping subnet ranges. The unique name enable selection of the correct subnet. Ignored if vrf_group_id or vrf_group is present in the arguments. Works only when adding a new IP. For existing IPs, use VRF group parameters. | 
 **macaddress** | **optional.String**| MAC address – can be new or existing | 
 **device** | **optional.String**| device name, can be new or existing | 
 **type_** | **optional.String**|  | 
 **notes** | **optional.String**| Any additional notes | 
 **vrfGroupId** | **optional.String**| ID of the VRF group | 
 **vrfGroup** | **optional.String**| VRF group name | 
 **available** | **optional.String**|  | 
 **clearAll** | **optional.String**| If yes - then IP is marked as available and device and mac address associations are cleared. Also notes and label fields are cleared. Added in v5.7.2 | 
 **tags** | **optional.String**| Update IP address tags (note, different than the antiquated tag endpoint. See label parameter above) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMMacs**
> interface{} PostIPAMMacs(ctx, optional)
Create / Update MAC Addresses.

Create / Update MAC Addresses. Deprecated since v12.0. Please use /switchports/ for more detailed information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiPostIPAMMacsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMMacsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **macaddress** | **optional.String**| MAC address – can be new or existing | 
 **portName** | **optional.String**| Interface name. (Please note: This is NOT the switchport name.) | 
 **override** | **optional.String**| smart – will detect if the port_name passed exist or not, if not – it is added to the current port name. Helpful, if you want to track all the port names for that mac address (e.g. eth0 &amp; bond0).&lt;br&gt;yes – change the port name. This is default behavior even if you don’t pass this parameter&lt;br&gt;no – will not change the port name | 
 **vlanId** | **optional.String**| ID of the vlan | 
 **device** | **optional.String**| name of the device | 
 **portId** | **optional.String**| Use this parameter or a combination of port and switch to specify the port. | 
 **port** | **optional.String**| Refers to the switchport name (not the interface name) - Use with parameter switch | 
 **switch_** | **optional.String**| Refers to the device name of the switch | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMSubnetsCreateChild**
> InlineResponse2008 PostIPAMSubnetsCreateChild(ctx, maskBits, optional)
Create child Subnet.

Create Child Subnet. Required parameters: <ul><li>mask_bits</li> <li>parent_subnet_id <b>OR</b> vrf_group <b>OR</b> vrf_group_id</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maskBits** | **string**| e.g. 24 | 
 **optional** | ***IPAMApiPostIPAMSubnetsCreateChildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMSubnetsCreateChildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentSubnetId** | **optional.String**| ID of the parent subnet. Can be obtained via /api/api/1.0/subnets/ or Tools &gt; Export &gt; Subnet. Required if vrf_group and vrf_group_id are not present. | 
 **vrfGroup** | **optional.String**| Name of the VRF group. Required if parent_subnet_id or vrf_group_id are not present. | 
 **vrfGroupId** | **optional.String**| ID of the VRF group. Required if parent_subnet_id or vrf_group are not present. | 
 **parentMaskBits** | **optional.String**| only if searching within a VRF and you want to restrict to certain parents with particular mask bits (added in v9.0.0) | 
 **ipv6** | **optional.String**| Required if creating an ipv6 subnet | 
 **network** | **optional.String**|  | 
 **subnetId** | **optional.String**|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMSwitches**
> interface{} PostIPAMSwitches(ctx, switchTemplateId, optional)
Add Switch/Switch Ports w/ Templates.

Create or update using Switch Templates. Required parameters: <ul><li>device <b>OR</b> device_id</li> <li>switch_template_id</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **switchTemplateId** | **string**| GET all Switch Templates | 
 **optional** | ***IPAMApiPostIPAMSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMSwitchesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **device** | **optional.String**| Name of new or existing device. Existing device must be a network switch. If stacked switches, must be of type &#39;cluster&#39; | 
 **deviceId** | **optional.String**| ID of existing device. Existing device must be a network switch. IF stacked switches, must be of type &#39;cluster&#39; | 
 **devices** | **optional.String**| comma separated names of new devices. | 
 **deviceIds** | **optional.String**| comma separated values of existing devices. | 
 **assets** | **optional.String**| comma separated names of new assets. | 
 **assetIds** | **optional.String**| comma separated values of existing assets. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMSwitchports**
> interface{} PostIPAMSwitchports(ctx, optional)
Create / Update Switch Port by Port.

Create / Update Switch Ports. <b>port</b> or <b>hwaddress</b> are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiPostIPAMSwitchportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMSwitchportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **port** | **optional.String**| child is optional assuming you have a parent created. Required if no hwaddress&lt;br&gt;Note: Do not use a port alias for the port name. | 
 **hwaddress** | **optional.String**| mac or wwn. Required if no port. | 
 **newPort** | **optional.String**| rename exisiting port | 
 **switch_** | **optional.String**| d42 will look for existing port and switch combination and if it exists, will update that port. | 
 **description** | **optional.String**|  | 
 **type_** | **optional.String**| Must exist in device42 (add via UI) | 
 **vlanId** | **optional.String**| ID of the vlan | 
 **up** | **optional.String**| &#x3D; yes for up. &#x3D; no for down. | 
 **upAdmin** | **optional.String**| Whether port is administratively up or down. &#x3D; yes for up. &#x3D; no for down. | 
 **count** | **optional.String**| Whether to include the port in total count or not. | 
 **remotePortId** | **optional.String**| ID of the remote connected switch port. | 
 **remoteDevice** | **optional.String**| Name of the switch for remote connected switch port. | 
 **remotePort** | **optional.String**| Name of the port for remote connected switch port. | 
 **module** | **optional.String**| name of the blade that port belongs to. Blade device must be part of the switch. (added in v5.8.1) | 
 **device2** | **optional.String**| name of the device2 that port belongs to. (added in v5.8.1) | 
 **device** | **optional.String**| name of the direcly connected device (Added in v5.8.2) (used to connect remote port for legacy support) | 
 **label** | **optional.String**|  | 
 **tags** | **optional.String**| add or update tags to a switchport | 
 **tagsRemove** | **optional.String**| remove tags from a switchport | 
 **mtu** | **optional.String**| add value for mtu | 
 **name** | **optional.String**| add name of port | 
 **speed** | **optional.String**| update port speed | 
 **remotePortClear** | **optional.String**| if set to yes, will clear the remote port | 
 **parentPort** | **optional.String**|  | 
 **parentPortDevice** | **optional.String**|  | 
 **slavePorts** | **optional.String**| comma separated port names | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMTapPorts**
> interface{} PostIPAMTapPorts(ctx, optional)
Create / Update TAP Ports

Create / Update TAP Ports. Required parameters: <ul><li>id <b>OR</b> name</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiPostIPAMTapPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMTapPortsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Tap Port ID | 
 **name** | **optional.String**|  | 
 **label** | **optional.String**|  | 
 **monitorDirection** | **optional.String**| Direction - to, both (optional, only for PortCap: Monitor) | 
 **monitorPort1Id** | **optional.String**|  | 
 **monitorPort2Id** | **optional.String**|  | 
 **morrorPort** | **optional.String**|  | 
 **netportId** | **optional.String**|  | 
 **patchPanelPortId** | **optional.String**|  | 
 **portCapability** | **optional.String**| Port Capability - Device, Monitor | 
 **portToFromId** | **optional.String**|  | 
 **portType** | **optional.String**| Port Type name - RJ45, RJ11, Fiber SC, Fiber FC | 
 **portTypeId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMsubnetcategory**
> PostIPAMsubnetcategory(ctx, name, optional)


Create Subnet Category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the Subnet Category you want to add edit or create | 
 **optional** | ***IPAMApiPostIPAMsubnetcategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMsubnetcategoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **optional.String**|  | 
 **isItDefault** | **optional.String**| Defaults to no. Only one category can be yes. If yes, this is the subnet category to use if none specified in an API or auto-discovery update. | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMsubnets**
> interface{} PostIPAMsubnets(ctx, network, maskBits, optional)
Create / Update Subnets.

Create / Update Subnets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Required for creation, cannot be modified after subnet creation. | 
  **maskBits** | **string**| Cannot be modified after subnet creation | 
 **optional** | ***IPAMApiPostIPAMsubnetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMsubnetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vrfGroupId** | **optional.String**| ID of the VRF group | 
 **vrfGroup** | **optional.String**| VRF group name | 
 **name** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **gateway** | **optional.String**| Gateway (added in v7.2.0) | 
 **rangeBegin** | **optional.String**| Range Begin (added in v7.2.0) | 
 **rangeEnd** | **optional.String**| Range End (added in v7.2.0) | 
 **parentVlanId** | **optional.String**|  | 
 **parentSubnetId** | **optional.String**| Change the parent subnet of the subnet. Note: must be valid parent. | 
 **customerId** | **optional.String**|  | 
 **notes** | **optional.String**| Any additional notes | 
 **serviceLevel** | **optional.String**| Must already exist | 
 **assigned** | **optional.String**| ‘yes’ if assigned. ‘no’ (default) if unassigned. | 
 **allocated** | **optional.String**| ‘yes’ if allocated. ‘no’ (default) if unallocated. | 
 **autoAddIps** | **optional.String**| If ‘yes’, addresses within subnet will be automatically added to Device42. (Only available in POST) | 
 **category** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no, Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. Use for initial insert. | 
 **newCategory** | **optional.String**| Use new_category to update an existing category. | 
 **categoryId** | **optional.String**| ID of the category - use for initial insert. | 
 **newCategoryId** | **optional.String**| Use new_category_id to update an existing category ID. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMvlans**
> interface{} PostIPAMvlans(ctx, number, optional)
Create VLANS.

Create VLANS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **number** | **string**| VLAN | 
 **optional** | ***IPAMApiPostIPAMvlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMvlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| If left blank, name will be created as VLANxxxx, e.g. VLAN# 342 will be named VLAN0342 | 
 **description** | **optional.String**|  | 
 **switchId** | **optional.String**| Comma separated values for switch_ids | 
 **notes** | **optional.String**| Any additional notes | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMvlansSmartMergeAll**
> interface{} PostIPAMvlansSmartMergeAll(ctx, optional)
Update VLANS by ID in url.

Smart Merge VLANS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiPostIPAMvlansSmartMergeAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMvlansSmartMergeAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **number** | **optional.String**| vlan number | 
 **name** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **switchId** | **optional.String**| Comma separated values for switch_ids | 
 **notes** | **optional.String**| Any additional notes | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIPAMvrfgroup**
> interface{} PostIPAMvrfgroup(ctx, name, optional)
Create/Update VRF Group.

Create/Update VRF Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the VRF Group you want to add edit or create | 
 **optional** | ***IPAMApiPostIPAMvrfgroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPostIPAMvrfgroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **optional.String**|  | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 
 **buildings** | **optional.String**| list of building names for the VRF Group | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIPAMCustomFIipAddress**
> interface{} PutIPAMCustomFIipAddress(ctx, ipAddress, subnetId, key, optional)
Create/update custom fields for subnets.

Subnet Custom Fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**| IP address | 
  **subnetId** | **string**| Subnet ID | 
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***IPAMApiPutIPAMCustomFIipAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMCustomFIipAddressOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **optional.String**| this is the custom field type. If left blank, default is text. Additional options: number,date(text formatted yyyy-mm-dd), related_field(with related_field_name), boolean or url | 
 **relatedFieldName** | **optional.String**| The existing field to relate this custom field to. Can be: appcomp (for application components), asset, building, certificate, circuit, cusotmer, device, endusers, hardware (for device hardware models), ip_address, natip, netport (for ports), os, part, partmodel, password, pdu (for power units), pdu_model (for power unit models), ports, purchase, purchaselineitem (for a line item on a purchase order), rack, room, software, vlan (for subnets), switch_vlan (for vlans), vrfgroup | 
 **addToPicklist** | **optional.String**| Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored. | 
 **value** | **optional.String**| Value of the custom field. | 
 **clearValue** | **optional.String**| yes to clear existing value for that field | 
 **notes** | **optional.String**| Any additional notes. | 
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

# **PutIPAMCustomFIsubnet**
> interface{} PutIPAMCustomFIsubnet(ctx, network, maskBits, vrfGroup, key, optional)
Create/update custom fields for subnets.

Subnet Custom Fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Network IP | 
  **maskBits** | **string**|  | 
  **vrfGroup** | **string**| VRF group name | 
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***IPAMApiPutIPAMCustomFIsubnetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMCustomFIsubnetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **type_** | **optional.String**| this is the custom field type. If left blank, default is text. Additional options: number,date(text formatted yyyy-mm-dd), related_field(with related_field_name), boolean or url | 
 **relatedFieldName** | **optional.String**| The existing field to relate this custom field to. Can be: appcomp (for application components), asset, building, certificate, circuit, cusotmer, device, endusers, hardware (for device hardware models), ip_address, natip, netport (for ports), os, part, partmodel, password, pdu (for power units), pdu_model (for power unit models), ports, purchase, purchaselineitem (for a line item on a purchase order), rack, room, software, vlan (for subnets), switch_vlan (for vlans), vrfgroup | 
 **addToPicklist** | **optional.String**| Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored. | 
 **value** | **optional.String**| Value of the custom field. | 
 **clearValue** | **optional.String**| Yes to clear the existing value for that field. | 
 **notes** | **optional.String**| Any additional notes. | 
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

# **PutIPAMCustomFIswitchport**
> interface{} PutIPAMCustomFIswitchport(ctx, key, optional)
Create / Update Switch Port Custom Fields

Required parameters: <ul><li>id <b>OR</b> port <b>AND</b> device_name</li> <li>key</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***IPAMApiPutIPAMCustomFIswitchportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMCustomFIswitchportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**| ID of the switchport | 
 **port** | **optional.String**| required if ID is not used | 
 **deviceName** | **optional.String**| required if switchport ID not used | 
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

# **PutIPAMCustomFIvrfgroup**
> interface{} PutIPAMCustomFIvrfgroup(ctx, key, optional)
VRF Group Custom Fields

Create/update custom fields for VRF groups. Required parameters: <ul><li>name <b>OR</b> ID</li> <li>key</li>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Can be new or existing. This is the custom field name. | 
 **optional** | ***IPAMApiPutIPAMCustomFIvrfgroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMCustomFIvrfgroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Required if ID not present | 
 **id** | **optional.String**| Required if name not present | 
 **value** | **optional.String**| Value of the custom field | 
 **clearValue** | **optional.String**| Yes to clear the existing value for that field. | 
 **notes** | **optional.String**| Any additional notes. | 
 **clearNotes** | **optional.String**| Yes to clear any existing notes. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIPAMIpnat**
> interface{} PutIPAMIpnat(ctx, id, optional)
Update IP NAT (Added in v7.0.0)

Update IP NAT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the NAT entry - Required if not in the URL | 
 **optional** | ***IPAMApiPutIPAMIpnatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMIpnatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Name of the NAT IP entry | 
 **ipAddressFrom** | **optional.String**| The external IP address | 
 **ipAddressTo** | **optional.String**| The internal IP address | 
 **twoWayRelation** | **optional.String**| true if the internal IP addressed is masked with the external IP address for outbound traffic in addition to inbound traffic. Default is false. | 
 **notes** | **optional.String**| Any additional notes | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIPAMsubnetcategory**
> interface{} PutIPAMsubnetcategory(ctx, name, optional)
Create Subnet Category.

Update a Specific Subnet Category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the Subnet Category | 
 **optional** | ***IPAMApiPutIPAMsubnetcategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMsubnetcategoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **optional.String**|  | 
 **isItDefault** | **optional.String**| yes or no. Defaults to no. Only one category can be yes. If yes, this is the subnet category to use if none specified in an API or auto-discovery update. | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIPAMsubnets**
> interface{} PutIPAMsubnets(ctx, network, maskBits, optional)
Create / Update Subnets.

Update a Specific Subnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **network** | **string**| Required for creation, cannot be modified after subnet creation. | 
  **maskBits** | **string**| Cannot be modified after subnet creation | 
 **optional** | ***IPAMApiPutIPAMsubnetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMsubnetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vrfGroupId** | **optional.String**| ID of the VRF group | 
 **vrfGroup** | **optional.String**| VRF group name | 
 **name** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **gateway** | **optional.String**| Gateway (added in v7.2.0) | 
 **rangeBegin** | **optional.String**| Range Begin (added in v7.2.0) | 
 **rangeEnd** | **optional.String**| Range End (added in v7.2.0) | 
 **parentVlanId** | **optional.String**|  | 
 **parentSubnetId** | **optional.String**| Change the parent subnet of the subnet. Note: must be valid parent. | 
 **customerId** | **optional.String**|  | 
 **notes** | **optional.String**| Any additional notes | 
 **serviceLevel** | **optional.String**| Must already exist | 
 **assigned** | **optional.String**| ‘yes’ if assigned. ‘no’ (default) if unassigned. | 
 **allocated** | **optional.String**| ‘yes’ if allocated. ‘no’ (default) if unallocated. | 
 **autoAddIps** | **optional.String**| If ‘yes’, addresses within subnet will be automatically added to Device42. (Only available in POST) | 
 **category** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no, Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. Use for initial insert. | 
 **categoryId** | **optional.String**| ID of the category - use for initial insert. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIPAMvlans**
> interface{} PutIPAMvlans(ctx, id, optional)
Update VLANS by ID in url.

Update VLANS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***IPAMApiPutIPAMvlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMvlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **number** | **optional.String**| vlan number | 
 **name** | **optional.String**| If left blank, name will be created as VLANxxxx, e.g. VLAN# 342 will be named VLAN0342 | 
 **description** | **optional.String**|  | 
 **switchId** | **optional.String**| Comma separated values for switch_ids | 
 **notes** | **optional.String**| Any additional notes | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIPAMvrfgroup**
> interface{} PutIPAMvrfgroup(ctx, optional)
Update VRF Group

Update a Specific VRF Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAMApiPutIPAMvrfgroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAMApiPutIPAMvrfgroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Name of the VRF Group | 
 **description** | **optional.String**|  | 
 **groups** | **optional.String**| If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted. | 
 **buildings** | **optional.String**| list of building names for the VRF Group | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

