# \SkuApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoiExtHotelSku**](SkuApi.md#PoiExtHotelSku) | **Get** /poi/ext/hotel/sku | sku拉取(该接口由接入方实现)
[**PoiSkuSync**](SkuApi.md#PoiSkuSync) | **Post** /poi/sku/sync | SKU同步



## PoiExtHotelSku

> PoiExtHotelSkuRsp PoiExtHotelSku(ctx, spuExtId, startDate, endDate)

sku拉取(该接口由接入方实现)

sku拉取(该接口由接入方实现)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spuExtId** | [**[]string**](string.md)| 接入方SPU ID 列表。 | 
**startDate** | **string**| 拉取价格时间区间[start_date, end_date)。 | 
**endDate** | **string**| 拉取价格时间区间[start_date, end_date)。 | 

### Return type

[**PoiExtHotelSkuRsp**](PoiExtHotelSkuRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiSkuSync

> PoiSkuSyncRsp PoiSkuSync(ctx, accessToken, optional)

SKU同步

SKU同步

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***PoiSkuSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiSkuSyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PoiSkuSyncReq**](PoiSkuSyncReq.md)|  | 

### Return type

[**PoiSkuSyncRsp**](PoiSkuSyncRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

