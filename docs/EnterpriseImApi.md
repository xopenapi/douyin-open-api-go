# \EnterpriseImApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseImPersonConversationCreate**](EnterpriseImApi.md#EnterpriseImPersonConversationCreate) | **Post** /enterprise/im/persona/conversation/create | 主动创建客服会话
[**EnterpriseImPersonConversationDelete**](EnterpriseImApi.md#EnterpriseImPersonConversationDelete) | **Post** /enterprise/im/persona/conversation/delete | 删除客服会话
[**EnterpriseImPersonCreate**](EnterpriseImApi.md#EnterpriseImPersonCreate) | **Post** /enterprise/im/persona/create | 客服账号
[**EnterpriseImPersonDelete**](EnterpriseImApi.md#EnterpriseImPersonDelete) | **Post** /enterprise/im/persona/delete | 删除客服账号
[**EnterpriseImPersonList**](EnterpriseImApi.md#EnterpriseImPersonList) | **Get** /enterprise/im/persona/list | 获取客服列表



## EnterpriseImPersonConversationCreate

> EnterpriseImPersonConversationCreateRsp EnterpriseImPersonConversationCreate(ctx, accessToken, openId, optional)

主动创建客服会话

主动创建客服会话

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseImPersonConversationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseImPersonConversationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnterpriseImPersonConversationCreateReq**](EnterpriseImPersonConversationCreateReq.md)|  | 

### Return type

[**EnterpriseImPersonConversationCreateRsp**](EnterpriseImPersonConversationCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseImPersonConversationDelete

> EnterpriseImPersonConversationDeleteRsp EnterpriseImPersonConversationDelete(ctx, accessToken, openId, optional)

删除客服会话

删除客服会话

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseImPersonConversationDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseImPersonConversationDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnterpriseImPersonConversationDeleteReq**](EnterpriseImPersonConversationDeleteReq.md)|  | 

### Return type

[**EnterpriseImPersonConversationDeleteRsp**](EnterpriseImPersonConversationDeleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseImPersonCreate

> EnterpriseImPersonCreateRsp EnterpriseImPersonCreate(ctx, accessToken, openId, optional)

客服账号

客服账号

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseImPersonCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseImPersonCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnterpriseImPersonCreateReq**](EnterpriseImPersonCreateReq.md)|  | 

### Return type

[**EnterpriseImPersonCreateRsp**](EnterpriseImPersonCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseImPersonDelete

> EnterpriseImPersonDeleteRsp EnterpriseImPersonDelete(ctx, accessToken, openId, optional)

删除客服账号

删除客服账号

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseImPersonDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseImPersonDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnterpriseImPersonDeleteReq**](EnterpriseImPersonDeleteReq.md)|  | 

### Return type

[**EnterpriseImPersonDeleteRsp**](EnterpriseImPersonDeleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseImPersonList

> EnterpriseImPersonListRsp EnterpriseImPersonList(ctx, accessToken, openId, cursor, count)

获取客服列表

获取客服列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**cursor** | **int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据 | 
**count** | **int64**| 每页数量 | 

### Return type

[**EnterpriseImPersonListRsp**](EnterpriseImPersonListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

