# \ImageApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DouyinImageCreate**](ImageApi.md#DouyinImageCreate) | **Post** /image/create | 发布图片
[**DouyinImageUpload**](ImageApi.md#DouyinImageUpload) | **Post** /image/upload | 上传图片到文件服务器



## DouyinImageCreate

> DouyinImageCreateRsp DouyinImageCreate(ctx, openId, accessToken, body)

发布图片

发布图片

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**body** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**DouyinImageCreateRsp**](DouyinImageCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinImageUpload

> DouyinImageUploadRsp DouyinImageUpload(ctx, openId, accessToken, image)

上传图片到文件服务器

上传图片到文件服务器

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**image** | ***os.File*****os.File**|  | 

### Return type

[**DouyinImageUploadRsp**](DouyinImageUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

