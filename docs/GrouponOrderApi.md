# \GrouponOrderApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseGrouponOrderDetail**](GrouponOrderApi.md#EnterpriseGrouponOrderDetail) | **Get** /enterprise/groupon/order/detail | 团购活动订单详情
[**EnterpriseGrouponOrderList**](GrouponOrderApi.md#EnterpriseGrouponOrderList) | **Get** /enterprise/groupon/order/list | 团购活动订单列表详情
[**EnterpriseGrouponOrderOverview**](GrouponOrderApi.md#EnterpriseGrouponOrderOverview) | **Get** /enterprise/groupon/order/overview | 团购活动订单汇总信息
[**EnterpriseGrouponOrderRefundApplyList**](GrouponOrderApi.md#EnterpriseGrouponOrderRefundApplyList) | **Get** /enterprise/groupon/order/refund/apply/list | 团购活动用户申请退款订单列表
[**EnterpriseGrouponOrderRefundConfirm**](GrouponOrderApi.md#EnterpriseGrouponOrderRefundConfirm) | **Post** /enterprise/groupon/order/refund/confirm | 确认退款



## EnterpriseGrouponOrderDetail

> EnterpriseGrouponOrderDetailRsp EnterpriseGrouponOrderDetail(ctx, accessToken, openId, orderId)

团购活动订单详情

团购活动订单详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**orderId** | **string**| 订单id | 

### Return type

[**EnterpriseGrouponOrderDetailRsp**](EnterpriseGrouponOrderDetailRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGrouponOrderList

> EnterpriseGrouponOrderListRsp EnterpriseGrouponOrderList(ctx, accessToken, openId, count, startTime, endTime, optional)

团购活动订单列表详情

团购活动订单列表详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**| 每页数量 | 
**startTime** | **int64**| 订单创建开始时间，unix时间戳 | 
**endTime** | **int64**| 订单创建结束时间，unix时间戳 | 
 **optional** | ***EnterpriseGrouponOrderListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponOrderListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 
 **filterStatus** | **optional.String**| 过滤订单状态，用\&quot;,\&quot;分隔，不传默认下发所有状态 | 

### Return type

[**EnterpriseGrouponOrderListRsp**](EnterpriseGrouponOrderListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGrouponOrderOverview

> EnterpriseGrouponOrderOverviewRsp EnterpriseGrouponOrderOverview(ctx, accessToken, openId, startTime, endTime)

团购活动订单汇总信息

团购活动订单汇总信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**startTime** | **int64**| 订单创建开始时间，unix时间戳 | 
**endTime** | **int64**| 订单创建结束时间，unix时间戳 | 

### Return type

[**EnterpriseGrouponOrderOverviewRsp**](EnterpriseGrouponOrderOverviewRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGrouponOrderRefundApplyList

> EnterpriseGrouponOrderRefundApplyListRsp EnterpriseGrouponOrderRefundApplyList(ctx, accessToken, openId, count, startTime, endTime, optional)

团购活动用户申请退款订单列表

团购活动用户申请退款订单列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**| 每页数量 | 
**startTime** | **int64**| 退款申请开始时间，unix时间戳 | 
**endTime** | **int64**| 退款申请结束时间，unix时间戳 | 
 **optional** | ***EnterpriseGrouponOrderRefundApplyListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponOrderRefundApplyListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 
 **status** | **optional.Int64**| 状态筛选 1 待确认 2 已确认 | 

### Return type

[**EnterpriseGrouponOrderRefundApplyListRsp**](EnterpriseGrouponOrderRefundApplyListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGrouponOrderRefundConfirm

> EnterpriseGrouponOrderRefundConfirmRsp EnterpriseGrouponOrderRefundConfirm(ctx, accessToken, openId, optional)

确认退款

确认退款

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseGrouponOrderRefundConfirmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponOrderRefundConfirmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orderId** | **optional.String**| 团购活动订单Id | 

### Return type

[**EnterpriseGrouponOrderRefundConfirmRsp**](EnterpriseGrouponOrderRefundConfirmRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

