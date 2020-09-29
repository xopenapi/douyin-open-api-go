# \EnterpriseImCardApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseImCardDelete**](EnterpriseImCardApi.md#EnterpriseImCardDelete) | **Post** /enterprise/im/card/delete | 删除消息卡片
[**EnterpriseImCardList**](EnterpriseImCardApi.md#EnterpriseImCardList) | **Get** /enterprise/im/card/list | 获取消息卡片列表
[**EnterpriseImCardSave**](EnterpriseImCardApi.md#EnterpriseImCardSave) | **Post** /enterprise/im/card/save | 创建/更新消息卡片



## EnterpriseImCardDelete

> EnterpriseImCardDeleteRsp EnterpriseImCardDelete(ctx, accessToken, openId, optional)

删除消息卡片

删除消息卡片

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseImCardDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseImCardDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cardId** | **optional.String**|  | 

### Return type

[**EnterpriseImCardDeleteRsp**](EnterpriseImCardDeleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseImCardList

> EnterpriseImCardListRsp EnterpriseImCardList(ctx, accessToken, openId, cursor, count)

获取消息卡片列表

获取消息卡片列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**cursor** | **int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 
**count** | **int64**| 每页数量 | 

### Return type

[**EnterpriseImCardListRsp**](EnterpriseImCardListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseImCardSave

> EnterpriseImCardSaveRsp EnterpriseImCardSave(ctx, accessToken, openId, optional)

创建/更新消息卡片

创建/更新消息卡片

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseImCardSaveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseImCardSaveOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cardId** | **optional.String**| 卡片id，创建时不传；更新时必传。同一个用户的卡片id不可重复 | 
 **cardType** | **optional.String**| 卡片类型 &#x60;question_list&#x60; - 问题列表 | 
 **content** | **optional.String**| 卡片内容字段 json序列化成string， { \\\&quot;question_list\\\&quot; { \\\&quot;text\\\&quot; \\\&quot;有什么疑问呢\\\&quot;, \\\&quot;questions\\\&quot; [{ \\\&quot;name\\\&quot; \\\&quot;问题1\\\&quot;, \\\&quot;text\\\&quot; \\\&quot;关键词1\\\&quot; }, { \\\&quot;name\\\&quot; \\\&quot;问题2\\\&quot;, \\\&quot;text\\\&quot; \\\&quot;关键词2\\\&quot; } ] } } | 

### Return type

[**EnterpriseImCardSaveRsp**](EnterpriseImCardSaveRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

