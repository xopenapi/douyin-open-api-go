# \EnterpriseApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseLeadsTagCreate**](EnterpriseApi.md#EnterpriseLeadsTagCreate) | **Post** /enterprise/leads/tag/create | 创建标签
[**EnterpriseLeadsTagDelete**](EnterpriseApi.md#EnterpriseLeadsTagDelete) | **Post** /enterprise/leads/tag/delete | 删除标签
[**EnterpriseLeadsTagList**](EnterpriseApi.md#EnterpriseLeadsTagList) | **Get** /enterprise/leads/tag/list | 获取标签列表
[**EnterpriseLeadsTagUpdate**](EnterpriseApi.md#EnterpriseLeadsTagUpdate) | **Post** /enterprise/leads/tag/update | 编辑标签
[**EnterpriseLeadsTagUserList**](EnterpriseApi.md#EnterpriseLeadsTagUserList) | **Get** /enterprise/leads/tag/user/list | 获取打标签的用户列表
[**EnterpriseLeadsTagUserUpdate**](EnterpriseApi.md#EnterpriseLeadsTagUserUpdate) | **Post** /enterprise/leads/tag/user/update | 给用户设置标签
[**EnterpriseLeadsUserActionList**](EnterpriseApi.md#EnterpriseLeadsUserActionList) | **Get** /enterprise/leads/user/action/list | 获取意向用户互动记录
[**EnterpriseLeadsUserDetail**](EnterpriseApi.md#EnterpriseLeadsUserDetail) | **Get** /enterprise/leads/user/detail | 获取意向用户详情
[**EnterpriseLeadsUserList**](EnterpriseApi.md#EnterpriseLeadsUserList) | **Get** /enterprise/leads/user/list | 获取意向用户列表



## EnterpriseLeadsTagCreate

> EnterpriseLeadsTagCreateRsp EnterpriseLeadsTagCreate(ctx, accessToken, openId, optional)

创建标签

创建标签

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseLeadsTagCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseLeadsTagCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnterpriseLeadsTagCreateReq**](EnterpriseLeadsTagCreateReq.md)|  | 

### Return type

[**EnterpriseLeadsTagCreateRsp**](EnterpriseLeadsTagCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLeadsTagDelete

> EnterpriseLeadsTagDeleteRsp EnterpriseLeadsTagDelete(ctx, accessToken, openId, optional)

删除标签

删除标签

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseLeadsTagDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseLeadsTagDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnterpriseLeadsTagDeleteReq**](EnterpriseLeadsTagDeleteReq.md)|  | 

### Return type

[**EnterpriseLeadsTagDeleteRsp**](EnterpriseLeadsTagDeleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLeadsTagList

> EnterpriseLeadsTagListRsp EnterpriseLeadsTagList(ctx, accessToken, openId, count, optional)

获取标签列表

获取标签列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**|  | 
 **optional** | ***EnterpriseLeadsTagListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseLeadsTagListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**|  | 

### Return type

[**EnterpriseLeadsTagListRsp**](EnterpriseLeadsTagListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLeadsTagUpdate

> EnterpriseLeadsTagUpdateRsp EnterpriseLeadsTagUpdate(ctx, accessToken, openId, optional)

编辑标签

编辑标签

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseLeadsTagUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseLeadsTagUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnterpriseLeadsTagUpdateReq**](EnterpriseLeadsTagUpdateReq.md)|  | 

### Return type

[**EnterpriseLeadsTagUpdateRsp**](EnterpriseLeadsTagUpdateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLeadsTagUserList

> EnterpriseLeadsTagUserListRsp EnterpriseLeadsTagUserList(ctx, accessToken, openId, count, tagId, optional)

获取打标签的用户列表

获取打标签的用户列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**|  | 
**tagId** | **int64**|  | 
 **optional** | ***EnterpriseLeadsTagUserListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseLeadsTagUserListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **cursor** | **optional.Int64**|  | 

### Return type

[**EnterpriseLeadsTagUserListRsp**](EnterpriseLeadsTagUserListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLeadsTagUserUpdate

> EnterpriseLeadsTagUserUpdateRsp EnterpriseLeadsTagUserUpdate(ctx, accessToken, openId, optional)

给用户设置标签

给用户设置标签

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseLeadsTagUserUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseLeadsTagUserUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnterpriseLeadsTagUserUpdateReq**](EnterpriseLeadsTagUserUpdateReq.md)|  | 

### Return type

[**EnterpriseLeadsTagUserUpdateRsp**](EnterpriseLeadsTagUserUpdateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLeadsUserActionList

> EnterpriseLeadsUserActionListRsp EnterpriseLeadsUserActionList(ctx, accessToken, openId, count, userId, optional)

获取意向用户互动记录

获取意向用户互动记录

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**|  | 
**userId** | **string**|  | 
 **optional** | ***EnterpriseLeadsUserActionListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseLeadsUserActionListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **cursor** | **optional.Int64**|  | 
 **actionType** | **optional.Int64**|  | 

### Return type

[**EnterpriseLeadsUserActionListRsp**](EnterpriseLeadsUserActionListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLeadsUserDetail

> EnterpriseLeadsUserDetailRsp EnterpriseLeadsUserDetail(ctx, accessToken, openId, userId)

获取意向用户详情

获取意向用户详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**userId** | **string**|  | 

### Return type

[**EnterpriseLeadsUserDetailRsp**](EnterpriseLeadsUserDetailRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLeadsUserList

> EnterpriseLeadsUserListRsp EnterpriseLeadsUserList(ctx, accessToken, openId, count, startTime, endTime, actionType, optional)

获取意向用户列表

获取意向用户列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**| 分页数量。 | 
**startTime** | **int64**|  | 
**endTime** | **int64**|  | 
**actionType** | **int64**| 分类 &#x60;0&#x60; - 全部 &#x60;1&#x60; - 私信互动 &#x60;2&#x60; - 组件互动 &#x60;3&#x60; - 主页互动 | 
 **optional** | ***EnterpriseLeadsUserListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseLeadsUserListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **cursor** | **optional.Int64**| 分页游标 | 
 **leadsLevel** | **optional.Int64**| 用户状态 &#x60;-1&#x60; - 没兴趣 &#x60;0&#x60; - 了解 &#x60;1&#x60; - 有兴趣 &#x60;2&#x60; - 有意愿 &#x60;10&#x60; - 已转化 | 

### Return type

[**EnterpriseLeadsUserListRsp**](EnterpriseLeadsUserListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

