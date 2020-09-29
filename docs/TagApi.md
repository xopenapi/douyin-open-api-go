# \TagApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseLeadsTagCreate**](TagApi.md#EnterpriseLeadsTagCreate) | **Post** /enterprise/leads/tag/create | 创建标签
[**EnterpriseLeadsTagDelete**](TagApi.md#EnterpriseLeadsTagDelete) | **Post** /enterprise/leads/tag/delete | 删除标签
[**EnterpriseLeadsTagUpdate**](TagApi.md#EnterpriseLeadsTagUpdate) | **Post** /enterprise/leads/tag/update | 编辑标签
[**EnterpriseLeadsTagUserUpdate**](TagApi.md#EnterpriseLeadsTagUserUpdate) | **Post** /enterprise/leads/tag/user/update | 给用户设置标签



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


 **tagName** | **optional.String**|  | 

### Return type

[**EnterpriseLeadsTagCreateRsp**](EnterpriseLeadsTagCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
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


 **tagId** | **optional.Int64**| 标签id | 

### Return type

[**EnterpriseLeadsTagDeleteRsp**](EnterpriseLeadsTagDeleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
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


 **tagId** | **optional.Int64**| 标签id | 
 **tagName** | **optional.String**| 标签名称 | 

### Return type

[**EnterpriseLeadsTagUpdateRsp**](EnterpriseLeadsTagUpdateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
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


 **bind** | **optional.Bool**| 是否绑定 | 
 **tagId** | **optional.Int64**| 标签id | 
 **userId** | **optional.String**| 用户openid | 

### Return type

[**EnterpriseLeadsTagUserUpdateRsp**](EnterpriseLeadsTagUserUpdateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

