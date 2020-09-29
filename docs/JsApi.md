# \JsApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JsGetticket**](JsApi.md#JsGetticket) | **Get** /js/getticket | 获取jsapi_ticket



## JsGetticket

> JsGetticketRsp JsGetticket(ctx, accessToken)

获取jsapi_ticket

获取jsapi_ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权 | 

### Return type

[**JsGetticketRsp**](JsGetticketRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

