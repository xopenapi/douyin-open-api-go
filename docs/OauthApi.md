# \OauthApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Connect**](OauthApi.md#Connect) | **Post** /connect | 获取授权码(code)



## Connect

> ApiResponse Connect(ctx, body)

获取授权码(code)

获取授权码(code)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ConnectReq**](ConnectReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

