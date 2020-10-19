# \ImApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseImMessageSend**](ImApi.md#EnterpriseImMessageSend) | **Post** /enterprise/im/message/send | (企业号)发送私信给用户



## EnterpriseImMessageSend

> EnterpriseImMessageSendRsp EnterpriseImMessageSend(ctx, openId, accessToken, body)

(企业号)发送私信给用户

(企业号)发送私信给用户

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**body** | [**EnterpriseImMessageSendReq**](EnterpriseImMessageSendReq.md)|  | 

### Return type

[**EnterpriseImMessageSendRsp**](EnterpriseImMessageSendRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

