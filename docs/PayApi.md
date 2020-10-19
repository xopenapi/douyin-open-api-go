# \PayApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DouyinPayAccountQuery**](PayApi.md#DouyinPayAccountQuery) | **Get** /douyin-pay/account-query | 账户余额查询
[**DouyinPayAccountTrans**](PayApi.md#DouyinPayAccountTrans) | **Post** /douyin-pay/account-trans | 商户向用户转账
[**DouyinPayOrderQuery**](PayApi.md#DouyinPayOrderQuery) | **Get** /douyin-pay/order-query | 订单查询，可查询一个月内的订单，优先级biz_order_no&gt;pay_order_no



## DouyinPayAccountQuery

> DouyinPayAccountQueryRsp DouyinPayAccountQuery(ctx, accessToken, merchantId, liveId)

账户余额查询

账户余额查询

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 通过/oauth/access_token/获取，用户唯一标志。 | 
**merchantId** | **int64**| 商户id | 
**liveId** | **int64**| 业务id。 | 

### Return type

[**DouyinPayAccountQueryRsp**](DouyinPayAccountQueryRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinPayAccountTrans

> DouyinPayAccountTransRsp DouyinPayAccountTrans(ctx, openId, accessToken, body)

商户向用户转账

商户向用户转账

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志。 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**body** | [**DouyinPayAccountTransReq**](DouyinPayAccountTransReq.md)|  | 

### Return type

[**DouyinPayAccountTransRsp**](DouyinPayAccountTransRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinPayOrderQuery

> DouyinPayOrderQueryRsp DouyinPayOrderQuery(ctx, accessToken, merchantId, liveId, optional)

订单查询，可查询一个月内的订单，优先级biz_order_no>pay_order_no

订单查询，可查询一个月内的订单，优先级biz_order_no>pay_order_no

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 通过/oauth/access_token/获取，用户唯一标志。 | 
**merchantId** | **int64**| 商户id | 
**liveId** | **int64**| 业务id。 | 
 **optional** | ***DouyinPayOrderQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DouyinPayOrderQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bizOrderNo** | **optional.String**| 外部订单号，由调用方生成 | 
 **payOrderNo** | **optional.String**| 抖音订单号，由抖音生成 | 

### Return type

[**DouyinPayOrderQueryRsp**](DouyinPayOrderQueryRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

