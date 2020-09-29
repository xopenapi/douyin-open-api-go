# \RankApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoveryEntRankItem**](RankApi.md#DiscoveryEntRankItem) | **Get** /discovery/ent/rank/item | 获取抖音电影榜、抖音电视剧榜、抖音综艺榜
[**DiscoveryEntRankVersion**](RankApi.md#DiscoveryEntRankVersion) | **Get** /discovery/ent/rank/version | 获取抖音影视综榜单版本



## DiscoveryEntRankItem

> DiscoveryEntRankItemRsp DiscoveryEntRankItem(ctx, accessToken, type_, optional)

获取抖音电影榜、抖音电视剧榜、抖音综艺榜

获取抖音电影榜、抖音电视剧榜、抖音综艺榜

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**type_** | **int32**| 榜单类型： 1 - 电影 2 - 电视剧 3 - 综艺 | 
 **optional** | ***DiscoveryEntRankItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiscoveryEntRankItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.Int32**| 榜单版本：空值默认为本周榜单 | 

### Return type

[**DiscoveryEntRankItemRsp**](DiscoveryEntRankItemRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryEntRankVersion

> DiscoveryEntRankVersionRsp DiscoveryEntRankVersion(ctx, accessToken, count, optional)

获取抖音影视综榜单版本

获取抖音影视综榜单版本

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
 **optional** | ***DiscoveryEntRankVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiscoveryEntRankVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 
 **type_** | **optional.Int32**| 榜单类型： 1 - 电影 2 - 电视剧 3 - 综艺 | 

### Return type

[**DiscoveryEntRankVersionRsp**](DiscoveryEntRankVersionRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

