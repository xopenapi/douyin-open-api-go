# \SandboxApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxWebhookEventSend**](SandboxApi.md#SandboxWebhookEventSend) | **Post** /sandbox/webhook/event/send | 模拟webhook事件



## SandboxWebhookEventSend

> SandboxWebhookEventEendRsp SandboxWebhookEventSend(ctx, accessToken, optional)

模拟webhook事件

模拟webhook事件

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
 **optional** | ***SandboxWebhookEventSendOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SandboxWebhookEventSendOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventType** | **optional.String**| 需要mock的事件类型, 开放平台会通过webhook发送一条mock数据给你 | 

### Return type

[**SandboxWebhookEventEendRsp**](SandboxWebhookEventEendRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

