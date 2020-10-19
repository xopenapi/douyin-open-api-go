# \SpuApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoiSpuQuery**](SpuApi.md#PoiSpuQuery) | **Get** /poi/spu/query | 查询SPU
[**PoiSpuSync**](SpuApi.md#PoiSpuSync) | **Post** /poi/spu/sync | SPU同步



## PoiSpuQuery

> PoiSpuQueryRsp PoiSpuQuery(ctx, accessToken, spuExtId)

查询SPU

查询SPU

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**spuExtId** | **string**|  | 

### Return type

[**PoiSpuQueryRsp**](PoiSpuQueryRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiSpuSync

> PoiSpuSyncRsp PoiSpuSync(ctx, accessToken, optional)

SPU同步

SPU同步

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***PoiSpuSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiSpuSyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PoiSpuSyncReq**](PoiSpuSyncReq.md)|  | 

### Return type

[**PoiSpuSyncRsp**](PoiSpuSyncRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

