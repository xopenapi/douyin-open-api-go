# \PoiOrderApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoiExtHotelOrderCancel**](PoiOrderApi.md#PoiExtHotelOrderCancel) | **Post** /poi/ext/hotel/order/cancel | 取消订单(该接口由接入方实现)
[**PoiExtHotelOrderCommit**](PoiOrderApi.md#PoiExtHotelOrderCommit) | **Post** /poi/ext/hotel/order/commit | 下单接口(该接口由接入方实现)
[**PoiExtHotelOrderStatus**](PoiOrderApi.md#PoiExtHotelOrderStatus) | **Post** /poi/ext/hotel/order/status | 支付状态通知(该接口由接入方实现)
[**PoiOrderStatus**](PoiOrderApi.md#PoiOrderStatus) | **Post** /poi/order/status | 订单状态同步接口
[**PoiOrderSync**](PoiOrderApi.md#PoiOrderSync) | **Post** /poi/order/sync | 订单同步



## PoiExtHotelOrderCancel

> PoiExtHotelOrderCancelRsp PoiExtHotelOrderCancel(ctx, body)

取消订单(该接口由接入方实现)

取消订单(该接口由接入方实现)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**PoiExtHotelOrderCancelReq**](PoiExtHotelOrderCancelReq.md)|  | 

### Return type

[**PoiExtHotelOrderCancelRsp**](PoiExtHotelOrderCancelRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiExtHotelOrderCommit

> PoiExtHotelOrderCommitRsp PoiExtHotelOrderCommit(ctx, body)

下单接口(该接口由接入方实现)

下单接口(该接口由接入方实现)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**PoiExtHotelOrderCommitReq**](PoiExtHotelOrderCommitReq.md)|  | 

### Return type

[**PoiExtHotelOrderCommitRsp**](PoiExtHotelOrderCommitRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiExtHotelOrderStatus

> PoiExtHotelOrderStatusRsp PoiExtHotelOrderStatus(ctx, body)

支付状态通知(该接口由接入方实现)

支付状态通知(该接口由接入方实现)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**PoiExtHotelOrderStatusReq**](PoiExtHotelOrderStatusReq.md)|  | 

### Return type

[**PoiExtHotelOrderStatusRsp**](PoiExtHotelOrderStatusRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiOrderStatus

> PoiOrderStatusRsp PoiOrderStatus(ctx, accessToken, body)

订单状态同步接口

订单状态同步接口

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 通过/oauth/access_token/获取，用户唯一标志。 | 
**body** | [**PoiOrderStatusReq**](PoiOrderStatusReq.md)|  | 

### Return type

[**PoiOrderStatusRsp**](PoiOrderStatusRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiOrderSync

> PoiOrderSyncRsp PoiOrderSync(ctx, accessToken, optional)

订单同步

订单同步

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***PoiOrderSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiOrderSyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PoiOrderSyncReq**](PoiOrderSyncReq.md)|  | 

### Return type

[**PoiOrderSyncRsp**](PoiOrderSyncRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

