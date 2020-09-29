# \DevtoolApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevtoolMicappIsLegal**](DevtoolApi.md#DevtoolMicappIsLegal) | **Get** /devtool/micapp/is_legal | 提供一个接口给开发者校验小程序appid是否可挂载到短视频



## DevtoolMicappIsLegal

> DevtoolMicappIsLegalRsp DevtoolMicappIsLegal(ctx, accessToken, micappId)

提供一个接口给开发者校验小程序appid是否可挂载到短视频

提供一个接口给开发者校验小程序appid是否可挂载到短视频

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**micappId** | **string**| 输入小程序的micapp_id | 

### Return type

[**DevtoolMicappIsLegalRsp**](DevtoolMicappIsLegalRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

