# \OauthApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Connect**](OauthApi.md#Connect) | **Post** /platform/oauth/connect | 获取授权码(code)



## Connect

> ApiResponse Connect(ctx, clientKey, responseType, scope, optionalScope, redirectUri, state)

获取授权码(code)

获取授权码(code)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientKey** | **string**| 应用唯一标识 | 
**responseType** | **string**| 填写code | 
**scope** | **string**| 应用授权作用域,多个授权作用域以英文逗号（,）分隔 | 
**optionalScope** | **string**| 应用授权可选作用域,多个授权作用域以英文逗号（,）分隔，每一个授权作用域后需要加上一个是否默认勾选的参数，1为默认勾选，0为默认不勾选 | 
**redirectUri** | **string**| 授权成功后的回调地址，必须以http/https开头。域名必须对应申请应用时填写的域名，如不清楚请联系应用申请人。 | 
**state** | **string**| 用于保持请求和回调的状态 | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

