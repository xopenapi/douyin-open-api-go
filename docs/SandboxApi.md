# \SandboxApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxWebhookEventSend**](SandboxApi.md#SandboxWebhookEventSend) | **Post** /sandbox/webhook/event/send | 模拟webhook事件



## SandboxWebhookEventSend

> SandboxWebhookEventEendRsp SandboxWebhookEventSend(ctx, accessToken, body)

模拟webhook事件

模拟webhook事件

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 
**body** | [**SandboxWebhookEventEendReq**](SandboxWebhookEventEendReq.md)|  | 

### Return type

[**SandboxWebhookEventEendRsp**](SandboxWebhookEventEendRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

