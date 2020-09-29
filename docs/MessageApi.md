# \MessageApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseImMessageSend**](MessageApi.md#EnterpriseImMessageSend) | **Post** /enterprise/im/message/send | (企业号)发送私信给用户



## EnterpriseImMessageSend

> EnterpriseImMessageSendRsp EnterpriseImMessageSend(ctx, openId, accessToken, optional)

(企业号)发送私信给用户

(企业号)发送私信给用户

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***EnterpriseImMessageSendOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseImMessageSendOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **personaId** | **optional.String**| 客服id，传则走客服会话，否则为普通会话 | 
 **toUserId** | **optional.String**| 消息的接收方openid | 
 **clientMsgId** | **optional.String**| 内部使用 | 
 **content** | **optional.String**| 消息体见下方schema | 
 **messageType** | **optional.String**| 消息内容格式 &#x60;text&#x60; - 文本消息 &#x60;image&#x60; - 图片消息 &#x60;video&#x60; - 视频消息 &#x60;card&#x60; - 卡片消息 | 

### Return type

[**EnterpriseImMessageSendRsp**](EnterpriseImMessageSendRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

