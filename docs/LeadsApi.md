# \LeadsApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseLeadsTagList**](LeadsApi.md#EnterpriseLeadsTagList) | **Get** /enterprise/leads/tag/list | 获取标签列表
[**EnterpriseLeadsTagUserList**](LeadsApi.md#EnterpriseLeadsTagUserList) | **Get** /enterprise/leads/tag/user/list | 获取打标签的用户列表



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

