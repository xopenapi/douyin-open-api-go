# \StarApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StarAuthorScore**](StarApi.md#StarAuthorScore) | **Get** /star/author_score | 获取抖音星图达人指数
[**StarAuthorScoreV2**](StarApi.md#StarAuthorScoreV2) | **Get** /star/author_score_v2 | 获取抖音星图达人指数数据V2
[**StarHostList**](StarApi.md#StarHostList) | **Get** /star/hot_list | 获取抖音星图达人热榜



## StarAuthorScore

> StarAuthorScoreRsp StarAuthorScore(ctx, accessToken, openId)

获取抖音星图达人指数

获取抖音星图达人指数

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 

### Return type

[**StarAuthorScoreRsp**](StarAuthorScoreRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StarAuthorScoreV2

> StarAuthorScoreV2Rsp StarAuthorScoreV2(ctx, accessToken, uniqueId)

获取抖音星图达人指数数据V2

获取抖音星图达人指数数据V2

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**uniqueId** | **string**| 达人抖音号 | 

### Return type

[**StarAuthorScoreV2Rsp**](StarAuthorScoreV2Rsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StarHostList

> StarHostListRsp StarHostList(ctx, accessToken, hotListType)

获取抖音星图达人热榜

获取抖音星图达人热榜

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**hotListType** | **int64**| 达人热榜类型 &#x60;1&#x60; - 星图指数榜 &#x60;2&#x60; - 涨粉指数榜 &#x60;3&#x60; - 性价比指数榜 &#x60;4&#x60; - 种草指数榜 &#x60;5&#x60; - 精选指数榜 &#x60;6&#x60; - 传播指数榜 | 

### Return type

[**StarHostListRsp**](StarHostListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

