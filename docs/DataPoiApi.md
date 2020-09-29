# \DataPoiApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataExternalPoiBase**](DataPoiApi.md#DataExternalPoiBase) | **Get** /data/external/poi/base | 获取POI基础数据
[**DataExternalPoiBillboard**](DataPoiApi.md#DataExternalPoiBillboard) | **Get** /data/external/poi/billboard | POI热度榜
[**DataExternalPoiClaimList**](DataPoiApi.md#DataExternalPoiClaimList) | **Get** /data/external/poi/claim/list | POI认领列表
[**DataExternalPoiServiceBase**](DataPoiApi.md#DataExternalPoiServiceBase) | **Get** /data/external/poi/service_base | POI服务基础数据
[**DataExternalPoiServiceUser**](DataPoiApi.md#DataExternalPoiServiceUser) | **Get** /data/external/poi/service_user | POI服务成交用户数据
[**DataExternalPoiUser**](DataPoiApi.md#DataExternalPoiUser) | **Get** /data/external/poi/user | POI用户数据
[**PoiBaseQueryAmap**](DataPoiApi.md#PoiBaseQueryAmap) | **Get** /poi/base/query/amap | 通过高德POI ID获取抖音POI ID



## DataExternalPoiBase

> DataExternalPoiBaseRsp DataExternalPoiBase(ctx, accessToken, poiId, beginDate, endDate)

获取POI基础数据

获取POI基础数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**poiId** | **string**| 抖音poi_id | 
**beginDate** | **string**| 最近30天，开始日期(yyyy-MM-dd) | 
**endDate** | **string**| 最近30天，结束日期(yyyy-MM-dd) | 

### Return type

[**DataExternalPoiBaseRsp**](DataExternalPoiBaseRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalPoiBillboard

> DataExternalPoiBillboardRsp DataExternalPoiBillboard(ctx, accessToken, billboardType)

POI热度榜

POI热度榜

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**billboardType** | **int64**| 10-近30日餐饮类POI的热度TOP500；20-近30日景点类POI的热度TOP500；30-近30日住宿类POI的热度TOP500 | 

### Return type

[**DataExternalPoiBillboardRsp**](DataExternalPoiBillboardRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalPoiClaimList

> DataExternalPoiClaimListRsp DataExternalPoiClaimList(ctx, accessToken, openId, count, optional)

POI认领列表

POI认领列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**| 每页数量 | 
 **optional** | ***DataExternalPoiClaimListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DataExternalPoiClaimListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 

### Return type

[**DataExternalPoiClaimListRsp**](DataExternalPoiClaimListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalPoiServiceBase

> DataExternalPoiServiceBaseRsp DataExternalPoiServiceBase(ctx, accessToken, poiId, beginDate, endDate, optional)

POI服务基础数据

POI服务基础数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**poiId** | **string**| 抖音poi_id | 
**beginDate** | **string**| 最近30天，开始日期(yyyy-MM-dd) | 
**endDate** | **string**| 最近30天，结束日期(yyyy-MM-dd) | 
 **optional** | ***DataExternalPoiServiceBaseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DataExternalPoiServiceBaseOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **serviceType** | **optional.Int64**| 服务类型，40:民宿 | 

### Return type

[**DataExternalPoiServiceBaseRsp**](DataExternalPoiServiceBaseRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalPoiServiceUser

> DataExternalPoiServiceUserRsp DataExternalPoiServiceUser(ctx, accessToken, poiId, dateType, optional)

POI服务成交用户数据

POI服务成交用户数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**poiId** | **string**| 抖音poi_id | 
**dateType** | **int64**| 近7/15/30天 | 
 **optional** | ***DataExternalPoiServiceUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DataExternalPoiServiceUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **serviceType** | **optional.Int64**| 服务类型，40:民宿 | 

### Return type

[**DataExternalPoiServiceUserRsp**](DataExternalPoiServiceUserRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalPoiUser

> DataExternalPoiUserRsp DataExternalPoiUser(ctx, accessToken, poiId, dateType)

POI用户数据

POI用户数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**poiId** | **string**| 抖音poi_id | 
**dateType** | **int64**| 近7/15/30天 | 

### Return type

[**DataExternalPoiUserRsp**](DataExternalPoiUserRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiBaseQueryAmap

> PoiBaseQueryAmapRsp PoiBaseQueryAmap(ctx, accessToken, amapId)

通过高德POI ID获取抖音POI ID

通过高德POI ID获取抖音POI ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**amapId** | **string**| 高德POI ID | 

### Return type

[**PoiBaseQueryAmapRsp**](PoiBaseQueryAmapRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

