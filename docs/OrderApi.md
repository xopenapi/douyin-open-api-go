# \OrderApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoiExtHotelOrderCancel**](OrderApi.md#PoiExtHotelOrderCancel) | **Post** /poi/ext/hotel/order/cancel | 取消订单(该接口由接入方实现)
[**PoiExtHotelOrderCommit**](OrderApi.md#PoiExtHotelOrderCommit) | **Post** /poi/ext/hotel/order/commit | 下单接口(该接口由接入方实现)
[**PoiExtHotelOrderStatus**](OrderApi.md#PoiExtHotelOrderStatus) | **Post** /poi/ext/hotel/order/status | 支付状态通知(该接口由接入方实现)
[**PoiOrderStatus**](OrderApi.md#PoiOrderStatus) | **Post** /poi/order/status | 订单状态同步接口
[**PoiOrderSync**](OrderApi.md#PoiOrderSync) | **Post** /poi/order/sync | 订单同步



## PoiExtHotelOrderCancel

> PoiExtHotelOrderCancelRsp PoiExtHotelOrderCancel(ctx, optional)

取消订单(该接口由接入方实现)

取消订单(该接口由接入方实现)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoiExtHotelOrderCancelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiExtHotelOrderCancelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **optional.String**| 抖音订单号 | 
 **orderStatus** | **optional.Int64**| 订单状态。0 - 未支付, 1 - 已支付 | 
 **supplierExtId** | **optional.String**| 接入方商铺ID | 
 **orderExtId** | **optional.String**| 接入方订单号 | 
 **datePrice** | [**optional.Interface of []PoiExtHotelOrderCommitDatePrice**](PoiExtHotelOrderCommitDatePrice.md)|  | 

### Return type

[**PoiExtHotelOrderCancelRsp**](PoiExtHotelOrderCancelRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiExtHotelOrderCommit

> PoiExtHotelOrderCommitRsp PoiExtHotelOrderCommit(ctx, optional)

下单接口(该接口由接入方实现)

下单接口(该接口由接入方实现)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoiExtHotelOrderCommitOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiExtHotelOrderCommitOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.Int64**| 订单支付状态。0 - 未支付, 1 - 已支付 | 
 **checkIn** | **optional.String**| 入住时间 yyyyMMdd | 
 **checkOut** | **optional.String**| 离店时间 yyyyMMdd | 
 **customerName** | **optional.String**| 预订人姓名 | 
 **customerPhone** | **optional.Int64**| 预订人电话 | 
 **orderId** | **optional.String**| 抖音订单号 | 
 **remark** | **optional.String**| 备注 | 
 **totalPrice** | **optional.Int64**| 总价, 单位人民币分 | 
 **reserveAmount** | **optional.Int64**| 预定数量 | 
 **spuExtId** | **optional.String**| 接入方房型ID | 
 **datePrice** | [**optional.Interface of []PoiExtHotelOrderCommitDatePrice**](PoiExtHotelOrderCommitDatePrice.md)|  | 
 **customerList** | [**optional.Interface of []PoiExtHotelOrderCommitCustomerList**](PoiExtHotelOrderCommitCustomerList.md)|  | 

### Return type

[**PoiExtHotelOrderCommitRsp**](PoiExtHotelOrderCommitRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiExtHotelOrderStatus

> PoiExtHotelOrderStatusRsp PoiExtHotelOrderStatus(ctx, optional)

支付状态通知(该接口由接入方实现)

支付状态通知(该接口由接入方实现)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoiExtHotelOrderStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiExtHotelOrderStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.Int64**| 订单支付状态。0 - 未支付, 1 - 已支付 | 
 **orderId** | **optional.String**| 抖音订单号 | 
 **supplierExtId** | **optional.String**| 接入方商铺ID | 
 **orderExtId** | **optional.String**| 接入方订单号 | 
 **datePrice** | [**optional.Interface of []PoiExtHotelOrderCommitDatePrice**](PoiExtHotelOrderCommitDatePrice.md)|  | 

### Return type

[**PoiExtHotelOrderStatusRsp**](PoiExtHotelOrderStatusRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiOrderStatus

> PoiOrderStatus PoiOrderStatus(ctx, accessToken, optional)

订单状态同步接口

订单状态同步接口

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 通过/oauth/access_token/获取，用户唯一标志。 | 
 **optional** | ***PoiOrderStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiOrderStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderId** | **optional.String**| 抖音订单号 | 
 **status** | **optional.Int64**| 订单确认状态。0 - 订单确认, 1 - 价格变动, 2 - 库存不足, 3 - 确认中 | 
 **supplierExtId** | **optional.String**| 接入方商铺ID | 
 **orderExtId** | **optional.String**| 接入方订单号 | 

### Return type

[**PoiOrderStatus**](PoiOrderStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
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

 **body** | [**optional.Interface of InlineObject40**](InlineObject40.md)|  | 

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

