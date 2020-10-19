# \MediaApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseMediaDelete**](MediaApi.md#EnterpriseMediaDelete) | **Post** /enterprise/media/delete | 删除永久素材
[**EnterpriseMediaList**](MediaApi.md#EnterpriseMediaList) | **Get** /enterprise/media/list | 获取永久素材列表
[**EnterpriseMediaTempUpload**](MediaApi.md#EnterpriseMediaTempUpload) | **Post** /enterprise/media/temp/upload | 上传临时素材
[**EnterpriseMediaUpload**](MediaApi.md#EnterpriseMediaUpload) | **Post** /enterprise/media/upload | 上传素材



## EnterpriseMediaDelete

> EnterpriseMediaDeleteRsp EnterpriseMediaDelete(ctx, accessToken, openId, body)

删除永久素材

删除永久素材

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**body** | [**EnterpriseMediaDeleteReq**](EnterpriseMediaDeleteReq.md)|  | 

### Return type

[**EnterpriseMediaDeleteRsp**](EnterpriseMediaDeleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseMediaList

> EnterpriseMediaListRsp EnterpriseMediaList(ctx, accessToken, openId, count, optional)

获取永久素材列表

获取永久素材列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**| 每页数量 | 
 **optional** | ***EnterpriseMediaListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseMediaListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据 | 

### Return type

[**EnterpriseMediaListRsp**](EnterpriseMediaListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseMediaTempUpload

> EnterpriseMediaTempUploadRsp EnterpriseMediaTempUpload(ctx, accessToken, openId, media)

上传临时素材

上传临时素材

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**media** | [**[]string**](string.md)|  | 

### Return type

[**EnterpriseMediaTempUploadRsp**](EnterpriseMediaTempUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseMediaUpload

> EnterpriseMediaUploadRsp EnterpriseMediaUpload(ctx, accessToken, openId, media)

上传素材

上传素材

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**media** | [**[]string**](string.md)|  | 

### Return type

[**EnterpriseMediaUploadRsp**](EnterpriseMediaUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

