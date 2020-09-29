# \OauthApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessToken**](OauthApi.md#AccessToken) | **Get** /oauth/access_token | 获取access_token
[**Authorize**](OauthApi.md#Authorize) | **Get** /oauth/authorize | 头条获取授权码(code)
[**AuthorizeV2**](OauthApi.md#AuthorizeV2) | **Get** /oauth/authorize/v2 | 抖音静默获取授权码(code)
[**ClientToken**](OauthApi.md#ClientToken) | **Get** /oauth/client_token | 生成client_token
[**Connect**](OauthApi.md#Connect) | **Get** /platform/oauth/connect | 抖音获取授权码(code)
[**RefreshToken**](OauthApi.md#RefreshToken) | **Get** /oauth/refresh_token | 刷新access_token
[**RenewRefreshToken**](OauthApi.md#RenewRefreshToken) | **Get** /oauth/renew_refresh_token | 刷新refresh_token



## AccessToken

> AccessTokenRsp AccessToken(ctx, clientKey, clientSecret, code, grantType)

获取access_token

获取access_token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientKey** | **string**| 应用唯一标识 | 
**clientSecret** | **string**| 应用唯一标识对应的密钥 | 
**code** | **string**| 授权码 | 
**grantType** | **string**| 写死\&quot;authorization_code\&quot;即可 | 

### Return type

[**AccessTokenRsp**](AccessTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Authorize

> AuthorizeRsp Authorize(ctx, clientKey, responseType, scope, redirectUri, optional)

头条获取授权码(code)

头条获取授权码(code)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientKey** | **string**| 应用唯一标识 | 
**responseType** | **string**| 填写code | 
**scope** | **string**| 应用授权作用域,多个授权作用域以英文逗号（,）分隔 | 
**redirectUri** | **string**| 授权成功后的回调地址，必须以http/https开头。域名必须对应申请应用时填写的域名，如不清楚请联系应用申请人。 | 
 **optional** | ***AuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **state** | **optional.String**| 用于保持请求和回调的状态 | 

### Return type

[**AuthorizeRsp**](AuthorizeRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizeV2

> AuthorizeV2Rsp AuthorizeV2(ctx, clientKey, responseType, scope, redirectUri, optional)

抖音静默获取授权码(code)

抖音静默获取授权码(code)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientKey** | **string**| 应用唯一标识 | 
**responseType** | **string**| 填写code | 
**scope** | **string**| 填login_id | 
**redirectUri** | **string**| 授权成功后的回调地址，必须以http/https开头。域名要跟申请应用时填写的授权回调域一致。用于调用https://open.douyin.com/oauth/access_token/换token。 | 
 **optional** | ***AuthorizeV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizeV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **state** | **optional.String**| 用于保持请求和回调状态，授权请求后会原样返回给接入方,如果是App则不用传该参数 | 

### Return type

[**AuthorizeV2Rsp**](AuthorizeV2Rsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientToken

> ClientTokenRsp ClientToken(ctx, clientKey, clientSecret, grantType)

生成client_token

生成client_token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientKey** | **string**| 应用唯一标识 | 
**clientSecret** | **string**| 应用唯一标识对应的密钥 | 
**grantType** | **string**| 传client_credential | 

### Return type

[**ClientTokenRsp**](ClientTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Connect

> ConnectRsp Connect(ctx, clientKey, responseType, scope, redirectUri, optional)

抖音获取授权码(code)

抖音获取授权码(code)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientKey** | **string**| 应用唯一标识 | 
**responseType** | **string**| 填写code | 
**scope** | **string**| 应用授权作用域,多个授权作用域以英文逗号（,）分隔 | 
**redirectUri** | **string**| 授权成功后的回调地址，必须以http/https开头。域名必须对应申请应用时填写的域名，如不清楚请联系应用申请人。 | 
 **optional** | ***ConnectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **optionalScope** | **optional.String**| 应用授权可选作用域,多个授权作用域以英文逗号（,）分隔，每一个授权作用域后需要加上一个是否默认勾选的参数，1为默认勾选，0为默认不勾选 | 
 **state** | **optional.String**| 用于保持请求和回调的状态 | 

### Return type

[**ConnectRsp**](ConnectRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshToken

> RefreshTokenRsp RefreshToken(ctx, clientKey, refreshToken, grantType)

刷新access_token

刷新access_token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientKey** | **string**| 应用唯一标识 | 
**refreshToken** | **string**| 填写通过access_token获取到的refresh_token参数 | 
**grantType** | **string**| 填refresh_token | 

### Return type

[**RefreshTokenRsp**](RefreshTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewRefreshToken

> RenewRefreshTokenRsp RenewRefreshToken(ctx, clientKey, refreshToken)

刷新refresh_token

刷新refresh_token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientKey** | **string**| 应用唯一标识 | 
**refreshToken** | **string**| 填写通过access_token获取到的refresh_token参数 | 

### Return type

[**RenewRefreshTokenRsp**](RenewRefreshTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

